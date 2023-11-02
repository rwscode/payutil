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

package alipay

import (
	"fmt"
)

// BizErr 用于判断支付宝的业务逻辑是否有错误
type BizErr struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

// bizErrCheck 检查业务码是否为10000 否则返回一个BizErr
func bizErrCheck(errRsp ErrorResponse) error {
	if errRsp.Code != "10000" {
		return &BizErr{
			Code:    errRsp.SubCode,
			Msg:     errRsp.Msg,
			SubCode: errRsp.SubCode,
			SubMsg:  errRsp.SubMsg,
		}
	}
	return nil
}

func (e *BizErr) Error() string {
	return fmt.Sprintf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, e.Code, e.Msg, e.SubCode, e.SubMsg)
}

func IsBizError(err error) (*BizErr, bool) {
	if bizErr, ok := err.(*BizErr); ok {
		return bizErr, true
	}
	return nil, false
}
