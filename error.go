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

package pay

import "errors"

var (
	MissWechatInitParamErr = errors.New("missing wechat init parameter")
	MissAlipayInitParamErr = errors.New("missing alipay init parameter")
	MissPayPalInitParamErr = errors.New("missing paypal init parameter")
	MissParamErr           = errors.New("missing required parameter")
	MarshalErr             = errors.New("marshal error")
	UnmarshalErr           = errors.New("unmarshal error")
	SignatureErr           = errors.New("signature error")
	VerifySignatureErr     = errors.New("verify signature error")
	CertNotMatchErr        = errors.New("cert not match error")
	GetSignDataErr         = errors.New("get signature data error")
)
