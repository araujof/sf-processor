package falco

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/common"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/falco/lang/errorhandler"
	"github.com/sysflow-telemetry/sf-processor/core/policyengine/policy/falco/lang/parser"
)

// Regular expression for parsing lists.
var itemsre = regexp.MustCompile(`(^\[)(.*)(\]$?)`)

// PolicyCompiler defines a compiler for extended Falco rules.
type PolicyCompiler[R any] struct {
	*parser.BaseSfplListener

	// Operations
	ops policy.Operations[R]

	// Parsed rule and filter object maps
	rules   []policy.Rule[R]
	filters []policy.Filter[R]

	// Accessory parsing maps
	lists     map[string][]string
	macroCtxs map[string]parser.IExpressionContext
}

// NewPolicyCompiler constructs a new compiler instance.
func NewPolicyCompiler[R any](ops policy.Operations[R]) policy.PolicyCompiler[R] {
	pc := new(PolicyCompiler[R])
	pc.ops = ops
	pc.rules = make([]policy.Rule[R], 0)
	pc.filters = make([]policy.Filter[R], 0)
	pc.lists = make(map[string][]string)
	pc.macroCtxs = make(map[string]parser.IExpressionContext)
	return pc
}

// Compile parses and interprets an input policy defined in path.
func (pc *PolicyCompiler[R]) compile(path string) error {
	// Setup the input
	is, err := antlr.NewFileStream(path)
	if err != nil {
		logger.Error.Println("Error reading policy from path", path)
		return err
	}

	// Create the Lexer
	lexerErrors := &errorhandler.SfplErrorListener{}
	lexer := parser.NewSfplLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	parserErrors := &errorhandler.SfplErrorListener{}
	p := parser.NewSfplParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	// Pre-processing (to deal with usage before definitions of macros and lists)
	antlr.ParseTreeWalkerDefault.Walk(pc, p.Defs())
	p.GetInputStream().Seek(0)

	// Parse the policy
	antlr.ParseTreeWalkerDefault.Walk(pc, p.Policy())

	errFound := false
	if len(lexerErrors.Errors) > 0 {
		logger.Error.Printf("Lexer %d errors found\n", len(lexerErrors.Errors))
		for _, e := range lexerErrors.Errors {
			logger.Error.Println("\t", e.Error())
		}
		errFound = true
	}
	if len(parserErrors.Errors) > 0 {
		logger.Error.Printf("Parser %d errors found\n", len(parserErrors.Errors))
		for _, e := range parserErrors.Errors {
			logger.Error.Println("\t", e.Error())
		}
		errFound = true
	}

	if errFound {
		return errors.New("errors found during compilation of policies. check logs for detail")
	}

	return nil
}

// Compile parses a set of input policies defined in paths.
func (pc *PolicyCompiler[R]) Compile(paths ...string) ([]policy.Rule[R], []policy.Filter[R], error) {
	for _, path := range paths {
		logger.Trace.Println("Parsing policy file ", path)
		if err := pc.compile(path); err != nil {
			return nil, nil, err
		}
	}
	return pc.rules, pc.filters, nil
}

// ExitList is called when production list is exited.
func (pc *PolicyCompiler[R]) ExitPlist(ctx *parser.PlistContext) {
	logger.Trace.Println("Parsing list ", ctx.GetText())
	pc.lists[ctx.ID().GetText()] = pc.extractListFromItems(ctx.Items())
}

// ExitMacro is called when production macro is exited.
func (pc *PolicyCompiler[R]) ExitPmacro(ctx *parser.PmacroContext) {
	logger.Trace.Println("Parsing macro ", ctx.GetText())
	pc.macroCtxs[ctx.ID().GetText()] = ctx.Expression()
}

// ExitFilter is called when production filter is exited.
func (pc *PolicyCompiler[R]) ExitPfilter(ctx *parser.PfilterContext) {
	logger.Trace.Println("Parsing filter ", ctx.GetText())
	f := policy.Filter[R]{
		Name:      ctx.ID().GetText(),
		Condition: pc.visitExpression(ctx.Expression()),
		Enabled:   ctx.ENABLED() == nil || pc.getEnabledFlag(ctx.Enabled()),
	}
	pc.filters = append(pc.filters, f)
}

// ExitFilter is called when production filter is exited.
func (pc *PolicyCompiler[R]) ExitPrule(ctx *parser.PruleContext) {
	logger.Trace.Println("Parsing rule ", ctx.GetText())
	r := policy.Rule[R]{
		Name:      pc.getOffChannelText(ctx.Text(0)),
		Desc:      pc.getOffChannelText(ctx.Text(1)),
		Condition: pc.visitExpression(ctx.Expression()),
		Actions:   pc.getActions(ctx),
		Tags:      pc.getTags(ctx),
		Priority:  pc.getPriority(ctx),
		Prefilter: pc.getPrefilter(ctx),
		Enabled:   ctx.ENABLED(0) == nil || pc.getEnabledFlag(ctx.Enabled(0)),
	}
	pc.rules = append(pc.rules, r)
}

func (pc *PolicyCompiler[R]) getEnabledFlag(ctx parser.IEnabledContext) bool {
	flag := common.TrimBoundingQuotes(ctx.GetText())
	if b, err := strconv.ParseBool(flag); err == nil {
		return b
	}
	logger.Warn.Println("Unrecognized enabled flag: ", flag)
	return true
}

func (pc *PolicyCompiler[R]) getOffChannelText(ctx parser.ITextContext) string {
	a := ctx.GetStart().GetStart()
	b := ctx.GetStop().GetStop()
	interval := antlr.Interval{Start: a, Stop: b}
	return ctx.GetStart().GetInputStream().GetTextFromInterval(&interval)
}

func (pc *PolicyCompiler[R]) getTags(ctx *parser.PruleContext) []policy.EnrichmentTag {
	var tags = make([]policy.EnrichmentTag, 0)
	ictx := ctx.Tags(0)
	if ictx != nil {
		return append(tags, pc.extractTags(ictx))
	}
	return tags
}

func (pc *PolicyCompiler[R]) getPrefilter(ctx *parser.PruleContext) []string {
	var pfs = make([]string, 0)
	ictx := ctx.Prefilter(0)
	if ictx != nil {
		return append(pfs, pc.extractList(ictx.GetText())...)
	}
	return pfs
}

// Fix: fix handling of priority levels.
func (pc *PolicyCompiler[R]) getPriority(ctx *parser.PruleContext) policy.Priority {
	ictx := ctx.Severity(0)
	if ictx != nil {
		p := ictx.GetText()
		switch strings.ToLower(p) {
		case policy.Low.String():
			return policy.Low
		case policy.Medium.String():
			return policy.Medium
		case policy.High.String():
			return policy.High
		case FPriorityDebug:
			return policy.Low
		case FPriorityInfo:
			return policy.Low
		case FPriorityInformational:
			return policy.Low
		case FPriorityNotice:
			return policy.Low
		case FPriorityWarning:
			return policy.Medium
		case FPriorityError:
			return policy.High
		case FPriorityCritical:
			return policy.High
		case FPriorityEmergency:
			return policy.High
		default:
			logger.Warn.Printf("Unrecognized priority value %s. Deferring to %s\n", p, policy.Low.String())
		}
	}
	return policy.Low
}

func (pc *PolicyCompiler[R]) getActions(ctx *parser.PruleContext) []string {
	var actions []string
	ictx := ctx.Actions(0)
	if ictx != nil {
		return append(actions, pc.extractActions(ictx)...)
	}
	return actions
}

func (pc *PolicyCompiler[R]) extractList(str string) []string {
	var items []string
	for _, i := range strings.Split(itemsre.ReplaceAllString(str, "$2"), common.LISTSEP) {
		items = append(items, common.TrimBoundingQuotes(i))
	}
	return items
}

func (pc *PolicyCompiler[R]) extractListFromItems(ctx parser.IItemsContext) []string {
	if ctx != nil {
		return pc.extractList(ctx.GetText())
	}
	return []string{}
}

func (pc *PolicyCompiler[R]) extractTags(ctx parser.ITagsContext) []string {
	if ctx != nil {
		return pc.extractList(ctx.GetText())
	}
	return []string{}
}

func (pc *PolicyCompiler[R]) extractActions(ctx parser.IActionsContext) []string {
	if ctx != nil {
		return pc.extractList(ctx.GetText())
	}
	return []string{}
}

func (pc *PolicyCompiler[R]) extractListFromAtoms(ctxs []parser.IAtomContext) []string {
	s := []string{}
	for _, v := range ctxs {
		s = append(s, pc.reduceList(v.GetText())...)
	}
	return s
}

func (pc *PolicyCompiler[R]) reduceList(sl string) []string {
	s := []string{}
	if l, ok := pc.lists[sl]; ok {
		for _, v := range l {
			s = append(s, pc.reduceList(v)...)
		}
	} else {
		s = append(s, common.TrimBoundingQuotes(sl))
	}
	return s
}

func (pc *PolicyCompiler[R]) visitExpression(ctx parser.IExpressionContext) policy.Criterion[R] {
	orCtx := ctx.GetChild(0).(parser.IOr_expressionContext)
	orPreds := make([]policy.Criterion[R], 0)
	for _, andCtx := range orCtx.GetChildren() {
		if andCtx.GetChildCount() > 0 {
			andPreds := make([]policy.Criterion[R], 0)
			for _, termCtx := range andCtx.GetChildren() {
				t, isTermCtx := termCtx.(parser.ITermContext)
				if isTermCtx {
					c := pc.visitTerm(t)
					andPreds = append(andPreds, c)
				}
			}
			orPreds = append(orPreds, policy.All(andPreds))
		}
	}
	return policy.Any(orPreds)
}

func (pc *PolicyCompiler[R]) visitTerm(ctx parser.ITermContext) policy.Criterion[R] {
	termCtx := ctx.(*parser.TermContext)
	if termCtx.Variable() != nil {
		if m, ok := pc.macroCtxs[termCtx.GetText()]; ok {
			return pc.visitExpression(m)
		}
		logger.Error.Println("Unrecognized reference ", termCtx.GetText())
	} else if termCtx.NOT() != nil {
		return pc.visitTerm(termCtx.GetChild(1).(parser.ITermContext)).Not()
	} else if opCtx, ok := termCtx.Unary_operator().(*parser.Unary_operatorContext); ok {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		if opCtx.EXISTS() != nil {
			return pc.ops.Exists(lop)
		}
		logger.Error.Println("Unrecognized unary operator ", opCtx.GetText())
	} else if opCtx, ok := termCtx.Binary_operator().(*parser.Binary_operatorContext); ok {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.Atom(1).(*parser.AtomContext).GetText()
		if opCtx.CONTAINS() != nil {
			return pc.ops.Contains(lop, rop)
		} else if opCtx.ICONTAINS() != nil {
			return pc.ops.IContains(lop, rop)
		} else if opCtx.STARTSWITH() != nil {
			return pc.ops.StartsWith(lop, rop)
		} else if opCtx.ENDSWITH() != nil {
			return pc.ops.EndsWith(lop, rop)
		} else if opCtx.EQ() != nil {
			return pc.ops.Eq(lop, rop)
		} else if opCtx.NEQ() != nil {
			return pc.ops.NEq(lop, rop)
		} else if opCtx.GT() != nil {
			return pc.ops.Gt(lop, rop)
		} else if opCtx.GE() != nil {
			return pc.ops.GEq(lop, rop)
		} else if opCtx.LT() != nil {
			return pc.ops.Lt(lop, rop)
		} else if opCtx.LE() != nil {
			return pc.ops.LEq(lop, rop)
		}
		logger.Error.Println("Unrecognized binary operator ", opCtx.GetText())
	} else if termCtx.Expression() != nil {
		return pc.visitExpression(termCtx.Expression())
	} else if termCtx.IN() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return pc.ops.In(lop, pc.extractListFromAtoms(rop))
	} else if termCtx.PMATCH() != nil {
		lop := termCtx.Atom(0).(*parser.AtomContext).GetText()
		rop := termCtx.AllAtom()[1:]
		return pc.ops.PMatch(lop, pc.extractListFromAtoms(rop))
	} else {
		logger.Warn.Println("Unrecognized term ", termCtx.GetText())
	}
	return policy.False[R]()
}
