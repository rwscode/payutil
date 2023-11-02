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

package wechat

import (
	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
	"github.com/rwscode/payutil/pkg/xlog"
	"github.com/rwscode/payutil/wechat"
)

func Micropay() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	number := util.RandomString(32)
	xlog.Debug("out_trade_no:", number)
	// 初始化参数Map
	bm := make(pay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", "扫用户付款码支付").
		Set("out_trade_no", number).
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("auth_code", "134595229789828537").
		Set("sign_type", wechat.SignType_MD5)

	// 请求支付，成功后得到结果
	wxRsp, err := client.Micropay(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)

	ok, err := wechat.VerifySign("GFDS8j98rewnmgl45wHTt980jg543abc", wechat.SignType_MD5, wxRsp)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("SignOk?：", ok)
}
