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
	"testing"
)

func TestBizErr_BizErrCheck(t *testing.T) {
	bizErrRsp := ErrorResponse{
		Code: "40004",
		Msg:  "NOT_FOUND",
	}
	if bizErrCheck(bizErrRsp) == nil {
		t.Fail()
	}

	noBizErrRsp := ErrorResponse{
		Code: "10000",
		Msg:  "SUCCEED",
	}

	if bizErrCheck(noBizErrRsp) != nil {
		t.Fail()
	}
}

func TestBizErr_AsBizError(t *testing.T) {
	bizErrRsp := ErrorResponse{
		Code: "40004",
		Msg:  "NOT_FOUND",
	}
	noBizErrRsp := ErrorResponse{
		Code: "10000",
		Msg:  "SUCCEED",
	}
	var err error
	err = bizErrCheck(bizErrRsp)
	if _, ok := IsBizError(err); !ok {
		t.Fail()
	}

	err = bizErrCheck(noBizErrRsp)
	if _, ok := IsBizError(err); !ok {
		t.Fail()
	}
}
