// Copyright 2023 payutil Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package xlog

const (
	ErrorLevel LogLevel = iota + 1
	WarnLevel
	InfoLevel
	DebugLevel
)

type LogLevel int

var (
	debugLog XLogger = &DebugLogger{}
	infoLog  XLogger = &InfoLogger{}
	warnLog  XLogger = &WarnLogger{}
	errLog   XLogger = &ErrorLogger{}

	Level LogLevel
)

type XLogger interface {
	LogOut(col *ColorType, format *string, args ...interface{})
}

func Info(args ...interface{}) {
	infoLog.LogOut(nil, nil, args...)
}

func Infof(format string, args ...interface{}) {
	infoLog.LogOut(nil, &format, args...)
}

func Debug(args ...interface{}) {
	debugLog.LogOut(nil, nil, args...)
}

func Debugf(format string, args ...interface{}) {
	debugLog.LogOut(nil, &format, args...)
}

func Warn(args ...interface{}) {
	warnLog.LogOut(nil, nil, args...)
}

func Warnf(format string, args ...interface{}) {
	warnLog.LogOut(nil, &format, args...)
}

func Error(args ...interface{}) {
	errLog.LogOut(nil, nil, args...)
}

func Errorf(format string, args ...interface{}) {
	errLog.LogOut(nil, &format, args...)
}

func SetDebugLog(logger XLogger) {
	debugLog = logger
}

func SetInfoLog(logger XLogger) {
	infoLog = logger
}

func SetWarnLog(logger XLogger) {
	warnLog = logger
}

func SetErrLog(logger XLogger) {
	errLog = logger
}
