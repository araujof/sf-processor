//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package engine

// Parsing constants.
const (
	LISTSEP string = ","
	EMPTY   string = ""
	QUOTE   string = "\""
	SPACE   string = " "
)

// SysFlow object types.
const (
	TyP      string = "P"
	TyF      string = "F"
	TyC      string = "C"
	TyH      string = "H"
	TyPE     string = "PE"
	TyFE     string = "FE"
	TyFF     string = "FF"
	TyNF     string = "NF"
	TyUnknow string = ""
)

// SysFlow attribute names.
const (
	SF_TYPE                 string = "sf.type"
	SF_OPFLAGS              string = "sf.opflags"
	SF_RET                  string = "sf.ret"
	SF_TS                   string = "sf.ts"
	SF_ENDTS                string = "sf.endts"
	SF_PROC_OID             string = "sf.proc.oid"
	SF_PROC_PID             string = "sf.proc.pid"
	SF_PROC_NAME            string = "sf.proc.name"
	SF_PROC_EXE             string = "sf.proc.exe"
	SF_PROC_ARGS            string = "sf.proc.args"
	SF_PROC_UID             string = "sf.proc.uid"
	SF_PROC_USER            string = "sf.proc.user"
	SF_PROC_TID             string = "sf.proc.tid"
	SF_PROC_GID             string = "sf.proc.gid"
	SF_PROC_GROUP           string = "sf.proc.group"
	SF_PROC_CREATETS        string = "sf.proc.createts"
	SF_PROC_TTY             string = "sf.proc.tty"
	SF_PROC_ENTRY           string = "sf.proc.entry"
	SF_PROC_CMDLINE         string = "sf.proc.cmdline"
	SF_PROC_ANAME           string = "sf.proc.aname"
	SF_PROC_AEXE            string = "sf.proc.aexe"
	SF_PROC_ACMDLINE        string = "sf.proc.acmdline"
	SF_PROC_APID            string = "sf.proc.apid"
	SF_PPROC_OID            string = "sf.pproc.oid"
	SF_PPROC_PID            string = "sf.pproc.pid"
	SF_PPROC_NAME           string = "sf.pproc.name"
	SF_PPROC_EXE            string = "sf.pproc.exe"
	SF_PPROC_ARGS           string = "sf.pproc.args"
	SF_PPROC_UID            string = "sf.pproc.uid"
	SF_PPROC_USER           string = "sf.pproc.user"
	SF_PPROC_GID            string = "sf.pproc.gid"
	SF_PPROC_GROUP          string = "sf.pproc.group"
	SF_PPROC_CREATETS       string = "sf.pproc.createts"
	SF_PPROC_TTY            string = "sf.pproc.tty"
	SF_PPROC_ENTRY          string = "sf.pproc.entry"
	SF_PPROC_CMDLINE        string = "sf.pproc.cmdline"
	SF_FILE_NAME            string = "sf.file.name"
	SF_FILE_PATH            string = "sf.file.path"
	SF_FILE_CANONICALPATH   string = "sf.file.canonicalpath"
	SF_FILE_DIRECTORY       string = "sf.file.directory"
	SF_FILE_NEWNAME         string = "sf.file.newname"
	SF_FILE_NEWPATH         string = "sf.file.newpath"
	SF_FILE_NEWDIRECTORY    string = "sf.file.newdirectory"
	SF_FILE_TYPE            string = "sf.file.type"
	SF_FILE_IS_OPEN_WRITE   string = "sf.file.is_open_write"
	SF_FILE_IS_OPEN_READ    string = "sf.file.is_open_read"
	SF_FILE_FD              string = "sf.file.fd"
	SF_FILE_OPENFLAGS       string = "sf.file.openflags"
	SF_NET_PROTO            string = "sf.net.proto"
	SF_NET_PROTONAME        string = "sf.net.protoname"
	SF_NET_SPORT            string = "sf.net.sport"
	SF_NET_DPORT            string = "sf.net.dport"
	SF_NET_PORT             string = "sf.net.port"
	SF_NET_SIP              string = "sf.net.sip"
	SF_NET_DIP              string = "sf.net.dip"
	SF_NET_IP               string = "sf.net.ip"
	SF_FLOW_RBYTES          string = "sf.flow.rbytes"
	SF_FLOW_ROPS            string = "sf.flow.rops"
	SF_FLOW_WBYTES          string = "sf.flow.wbytes"
	SF_FLOW_WOPS            string = "sf.flow.wops"
	SF_CONTAINER_ID         string = "sf.container.id"
	SF_CONTAINER_NAME       string = "sf.container.name"
	SF_CONTAINER_IMAGEID    string = "sf.container.imageid"
	SF_CONTAINER_IMAGE      string = "sf.container.image"
	SF_CONTAINER_TYPE       string = "sf.container.type"
	SF_CONTAINER_PRIVILEGED string = "sf.container.privileged"
	SF_NODE_ID              string = "sf.node.id"
	SF_NODE_IP              string = "sf.node.ip"
	SF_SCHEMA_VERSION       string = "sf.schema"
)

// extension proc attributes
const (
	EXT_PROC_GUID_STR                = "ext.proc.guid"
	EXT_PROC_IMAGE_STR               = "ext.proc.image"
	EXT_PROC_CURR_DIRECTORY_STR      = "ext.proc.curdir"
	EXT_PROC_LOGON_GUID_STR          = "ext.proc.logonguid"
	EXT_PROC_LOGON_ID_STR            = "ext.proc.logonid"
	EXT_PROC_TERMINAL_SESSION_ID_STR = "ext.proc.termsessid"
	EXT_PROC_INTEGRITY_LEVEL_STR     = "ext.proc.integrity"
	EXT_PROC_SIGNATURE_STR           = "ext.proc.signature"
	EXT_PROC_SIGNATURE_STATUS_STR    = "ext.proc.sigstatus"
	EXT_PROC_SHA1_HASH_STR           = "ext.proc.sha1"
	EXT_PROC_MD5_HASH_STR            = "ext.proc.md5"
	EXT_PROC_SHA256_HASH_STR         = "ext.proc.sha256"
	EXT_PROC_IMP_HASH_STR            = "ext.proc.imphash"
	EXT_PROC_SIGNED_INT              = "ext.proc.signed"
)

//extension file attributes
const (
	EXT_FILE_SHA1_HASH_STR        = "ext.file.sha1"
	EXT_FILE_MD5_HASH_STR         = "ext.file.md5"
	EXT_FILE_SHA256_HASH_STR      = "ext.file.sha256"
	EXT_FILE_IMP_HASH_STR         = "ext.file.imp"
	EXT_FILE_SIGNATURE_STR        = "ext.file.signature"
	EXT_FILE_SIGNATURE_STATUS_STR = "ext.file.sigstatus"
	EXT_FILE_DETAILS_STR          = "ext.registry.details"
	EXT_FILE_SIGNED_INT           = "ext.file.signed"
)

// extensions for network
const (
	EXT_NET_SOURCE_HOST_NAME_STR = "ext.net.srchostname"
	EXT_NET_SOURCE_PORT_NAME_STR = "ext.net.srcportname"
	EXT_NET_DEST_HOST_NAME_STR   = "ext.net.desthostname"
	EXT_NET_DEST_PORT_NAME_STR   = "ext.net.destportname"
)

// extensions for events
const (
	EXT_TARG_PROC_STATE_INT              = "ext.targetproc.state"
	EXT_TARG_PROC_OID_CREATETS_INT       = "ext.targetproc.createts"
	EXT_TARG_PROC_OID_HPID_INT           = "ext.targetproc.pid"
	EXT_TARG_PROC_POID_CREATETS_INT      = "ext.targetpproc.createts"
	EXT_TARG_PROC_POID_HPID_INT          = "ext.targetpproc.pid"
	EXT_TARG_PROC_TS_INT                 = "ext.targetproc.ts"
	EXT_TARG_PROC_EXE_STR                = "ext.targetproc.exe"
	EXT_TARG_PROC_EXEARGS_STR            = "ext.targetproc.args"
	EXT_TARG_PROC_UID_INT                = "ext.targetproc.uid"
	EXT_TARG_PROC_USERNAME_STR           = "ext.targetproc.user"
	EXT_TARG_PROC_GID_INT                = "ext.targetproc.gid"
	EXT_TARG_PROC_GROUPNAME_STR          = "ext.targetproc.group"
	EXT_TARG_PROC_TTY_INT                = "ext.targetproc.tty"
	EXT_TARG_PROC_CONTAINERID_STRING_STR = "ext.targetcontainer.id"
	EXT_TARG_PROC_ENTRY_INT              = "ext.targetproc.entry"

	EXT_TARG_PROC_GUID_STR                = "ext.targetproc.guid"
	EXT_TARG_PROC_IMAGE_STR               = "ext.targetproc.image"
	EXT_TARG_PROC_CURR_DIRECTORY_STR      = "ext.targetproc.curdir"
	EXT_TARG_PROC_LOGON_GUID_STR          = "ext.targetproc.logonguid"
	EXT_TARG_PROC_LOGON_ID_STR            = "ext.targetproc.logonid"
	EXT_TARG_PROC_TERMINAL_SESSION_ID_STR = "ext.targetproc.termsessid"
	EXT_TARG_PROC_INTEGRITY_LEVEL_STR     = "ext.targetproc.integrity"
	EXT_TARG_PROC_SIGNATURE_STR           = "ext.targetproc.signature"
	EXT_TARG_PROC_SIGNATURE_STATUS_STR    = "ext.targetproc.sigstatus"
	EXT_TARG_PROC_SHA1_HASH_STR           = "ext.targetproc.sha1"
	EXT_TARG_PROC_MD5_HASH_STR            = "ext.targetproc.md5"
	EXT_TARG_PROC_SHA256_HASH_STR         = "ext.targetproc.sha256"
	EXT_TARG_PROC_IMP_HASH_STR            = "ext.targetproc.imphash"
	EXT_TARG_PROC_START_ADDR_STR          = "ext.targetproc.startaddr"
	EXT_TARG_PROC_START_MODULE_STR        = "ext.targetproc.startmod"
	EXT_TARG_PROC_START_FUNCTION_STR      = "ext.targetproc.startfunc"
	EXT_TARG_PROC_GRANT_ACCESS_STR        = "ext.targetproc.grantaccess"
	EXT_TARG_PROC_CALL_TRACE_STR          = "ext.targetproc.calltrace"
	EXT_TARG_PROC_ACCESS_TYPE_STR         = "ext.targetproc.accesstype"

	EXT_TARG_PROC_SIGNED_INT        = "ext.targetproc.signed"
	EXT_TARG_PROC_NEW_THREAD_ID_INT = "ext.targetproc.newthreadid"
)
