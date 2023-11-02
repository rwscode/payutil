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

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type ErrorLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (e *ErrorLogger) LogOut(col *ColorType, format *string, v ...interface{}) {
	e.once.Do(func() {
		e.init()
	})
	if Level >= ErrorLevel {
		if col != nil {
			if format != nil {
				e.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
				return
			}
			e.logger.Output(3, string(*col)+fmt.Sprintln(v...)+string(Reset))
			return
		}
		if format != nil {
			e.logger.Output(3, fmt.Sprintf(*format, v...))
			return
		}
		e.logger.Output(3, fmt.Sprintln(v...))
	}
}

func (e *ErrorLogger) init() {
	if Level == 0 {
		Level = DebugLevel
	}
	e.logger = log.New(os.Stdout, "[ERROR] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
