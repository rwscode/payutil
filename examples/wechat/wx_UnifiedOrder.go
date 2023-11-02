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
	"strconv"
	"time"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
	"github.com/rwscode/payutil/pkg/xlog"
	"github.com/rwscode/payutil/wechat"
)

func UnifiedOrder() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 设置国家
	client.SetCountry(wechat.China)

	number := util.RandomString(32)
	xlog.Debug("out_trade_no:", number)

	// 初始化参数Map
	bm := make(pay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("body", "H5支付").
		Set("out_trade_no", number).
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("notify_url", "https://www.fmm.ink").
		Set("trade_type", wechat.TradeType_H5).
		Set("device_info", "WEB").
		Set("sign_type", wechat.SignType_MD5).
		SetBodyMap("scene_info", func(bm pay.BodyMap) {
			bm.SetBodyMap("h5_info", func(bm pay.BodyMap) {
				bm.Set("type", "Wap")
				bm.Set("wap_url", "https://www.fmm.ink")
				bm.Set("wap_name", "H5测试支付")
			})
		}) /*.Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8")*/

	// 请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)
	// xlog.Debug("wxRsp.MwebUrl:", wxRsp.MwebUrl)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 获取小程序支付需要的paySign
	// pac := "prepay_id=" + wxRsp.PrepayId
	// paySign := wechat.GetMiniPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, wechat.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	// xlog.Debug("paySign:", paySign)

	// 获取Jsapi支付需要的paySign
	pac := "prepay_id=" + wxRsp.PrepayId
	paySign := wechat.GetJsapiPaySign("wxdaa2ab9ef87b5497", wxRsp.NonceStr, pac, wechat.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	xlog.Debug("paySign:", paySign)

	// 获取App支付需要的paySign
	// paySign := wechat.GetAppPaySign("wxdaa2ab9ef87b5497","", wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, "GFDS8j98rewnmgl45wHTt980jg543abc")
	// xlog.Debug("paySign:", paySign)
}
