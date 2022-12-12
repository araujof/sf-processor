//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package sysflow implements pluggable drivers for SysFlow ingestion.
package sysflow

import (
	"bufio"
	"errors"
	"io/ioutil"
	"os"

	"github.com/linkedin/goavro"
	"github.com/sysflow-telemetry/sf-apis/go/converter"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/plugins"
)

const (
	fileDriverName = "file"
)

func getFiles(filename string) ([]string, error) {
	var fls []string
	if fi, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	} else if fi.IsDir() {
		logger.Trace.Println("File is a directory")
		var files []os.FileInfo
		var err error
		if files, err = ioutil.ReadDir(filename); err != nil {
			return nil, err
		}
		for _, file := range files {
			f := filename + "/" + file.Name()
			logger.Trace.Println("File in Directory: " + f)
			fls = append(fls, f)
		}
		if len(fls) == 0 {
			return nil, errors.New("No files present in directory: " + filename)
		}

	} else {
		fls = append(fls, filename)
	}
	logger.Trace.Printf("Number of files in list: %d\n", len(fls))
	return fls, nil
}

// FileDriver represents reading a sysflow file from source
type FileDriver struct {
	pipeline plugins.SFPipeline
	config   map[string]interface{}
	file     *os.File
}

// NewFileDriver creates a new file driver object
func NewFileDriver() plugins.SFDriver {
	return &FileDriver{}
}

// GetName returns the driver name.
func (s *FileDriver) GetName() string {
	return fileDriverName
}

// Register registers driver to plugin cache
func (s *FileDriver) Register(pc plugins.SFPluginCache) {
	pc.AddDriver(fileDriverName, NewFileDriver)
}

// Init initializes the file driver with the pipeline
func (s *FileDriver) Init(pipeline plugins.SFPipeline, config map[string]interface{}) error {
	s.pipeline = pipeline
	s.config = config
	return nil
}

// Run runs the file driver
func (s *FileDriver) Run(path string, running *bool) error {
	var channel interface{}
	configpath := path
	if s.config == nil {
		channel = s.pipeline.GetRootChannel()
	} else {
		if v, o := s.config[OutChanConfig].(string); o {
			ch, err := s.pipeline.GetChannel(v)
			if err != nil {
				return err
			}
			channel = ch
		} else {
			return errors.New("out tag does not exist in driver configuration for driver " + fileDriverName)
		}
		if v, o := s.config[PathConfig].(string); o {
			configpath = v
		}
	}
	sfChannel := channel.(*plugins.SFChannel)
	records := sfChannel.In

	logger.Trace.Println("Loading file: ", path)

	sfobjcvter := converter.NewSFObjectConverter()

	files, err := getFiles(configpath)
	if err != nil {
		logger.Error.Println("Files error: ", err)
		return err
	}
	for _, fn := range files {
		logger.Trace.Println("Loading file: " + fn)
		s.file, err = os.Open(fn)
		if err != nil {
			logger.Error.Println("File open error: ", err)
			return err
		}
		reader := bufio.NewReader(s.file)
		sreader, err := goavro.NewOCFReader(reader)
		if err != nil {
			logger.Error.Println("Reader error: ", err)
			return err
		}
		for sreader.Scan() {
			if !*running {
				break
			}
			datum, err := sreader.Read()
			if err != nil {
				logger.Error.Println("Datum reading error: ", err)
				break
			}
			records <- sfobjcvter.ConvertToSysFlow(datum)
		}
		s.file.Close()
		if !*running {
			break
		}
	}
	logger.Trace.Println("Closing main channel filedriver")
	close(records)
	s.pipeline.Wait()
	logger.Trace.Println("Exiting Process() function filedriver")
	return nil
}

// Cleanup tears down the driver resources.
func (s *FileDriver) Cleanup() {
	logger.Trace.Println("Exiting ", fileDriverName)
	if s.file != nil {
		s.file.Close()
	}
}
