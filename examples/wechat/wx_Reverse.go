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

func Reverse() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	// 初始化参数Map
	bm := make(pay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("out_trade_no", "6aDCor1nUcAihrV5JBlI09tLvXbUp02B").
		Set("sign_type", wechat.SignType_MD5)

	// 请求撤销订单，成功后得到结果，沙箱环境下，证书路径参数可传空
	wxRsp, err := client.Reverse(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("wxRsp：", wxRsp)
}
