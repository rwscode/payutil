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

type WarnLogger struct {
	logger *log.Logger
	once   sync.Once
}

func (i *WarnLogger) LogOut(col *ColorType, format *string, v ...interface{}) {
	i.once.Do(func() {
		i.init()
	})
	if Level >= WarnLevel {
		if col != nil {
			if format != nil {
				i.logger.Output(3, string(*col)+fmt.Sprintf(*format, v...)+string(Reset))
				return
			}
			i.logger.Output(3, string(*col)+fmt.Sprintln(v...)+string(Reset))
			return
		}
		if format != nil {
			i.logger.Output(3, fmt.Sprintf(*format, v...))
			return
		}
		i.logger.Output(3, fmt.Sprintln(v...))
	}
}

func (i *WarnLogger) init() {
	if Level == 0 {
		Level = DebugLevel
	}
	i.logger = log.New(os.Stdout, "[WARN] >> ", log.Lmsgprefix|log.Lshortfile|log.Ldate|log.Lmicroseconds)
}
