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
	"context"
	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
	"github.com/rwscode/payutil/pkg/xlog"
	"github.com/rwscode/payutil/wechat"
)

var ctx = context.Background()

func Refund() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	s := util.RandomString(64)
	xlog.Debug("out_refund_no:", s)
	// 初始化参数结构体
	bm := make(pay.BodyMap)
	bm.Set("out_trade_no", "SdZBAqJHBQGKVwb7aMR2mUwC588NG2Sd").
		Set("nonce_str", util.RandomString(32)).
		Set("sign_type", wechat.SignType_MD5).
		Set("out_refund_no", s).
		Set("total_fee", 1).
		Set("refund_fee", 1).
		Set("notify_url", "https://www.fmm.ink")

	// 请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	wxRsp, resBm, err := client.Refund(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
	xlog.Debug("resBm:", resBm)
}
