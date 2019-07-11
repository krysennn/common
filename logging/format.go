package logging

// Copy-pasted from prometheus/common/promlog.
// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"flag"

	"github.com/pkg/errors"
)

// Format is a settable identifier to format a log entry
// must be have.
type Format struct {
	s      string
}

// RegisterFlags adds the log format flag to the provided flagset.
func (l *Format) RegisterFlags(f *flag.FlagSet) {
	l.Set("logfmt")
	f.Var(l, "log.format", "Format log message. Valid format: [logfmt, json]")
}

func (l *Format) String() string {
	return l.s
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (l *Format) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var format string
	if err := unmarshal(&format); err != nil {
		return err
	}
	return l.Set(format)
}

// MarshalYAML implements yaml.Marshaler.
func (l Format) MarshalYAML() (interface{}, error) {
	return l.String(), nil
}

// Set updates the value of the allowed format.  Implments flag.Value.
func (l *Format) Set(s string) error {
	switch s {
	case "logfmt":
		l.s = "logfmt"
	case "json":
		l.s = "json"
	default:
		return errors.Errorf("unrecognized log format %q", s)
	}

	return nil
}
