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

package apple

import (
	"context"

	"github.com/rwscode/payutil/pkg/xhttp"
)

const (
	// is the URL when testing your app in the sandbox and while your application is in review
	UrlSandbox = "https://sandbox.itunes.apple.com/verifyReceipt"

	// is the URL when your app is live in the App Store
	UrlProd = "https://buy.itunes.apple.com/verifyReceipt"
)

// VerifyReceipt 请求APP Store 校验支付请求,实际测试时发现这个文档介绍的返回信息只有那个status==0表示成功可以用，其他的返回信息跟文档对不上
// url：取 UrlProd 或 UrlSandbox
// pwd：苹果APP秘钥，https://help.apple.com/app-store-connect/#/devf341c0f01
// 文档：https://developer.apple.com/documentation/appstorereceipts/verifyreceipt
func VerifyReceipt(ctx context.Context, url, pwd, receipt string) (*VerifyResponse, error) {
	req := &VerifyRequest{Receipt: receipt, Password: pwd}
	vr := new(VerifyResponse)
	_, err := xhttp.NewClient().Type(xhttp.TypeJSON).Post(url).SendStruct(req).EndStruct(ctx, vr)
	if err != nil {
		return nil, err
	}
	return vr, nil
}
