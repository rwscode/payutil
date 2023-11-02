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

func Transfer() {
	// client只需要初始化一个，此处为了演示，每个方法都做了初始化
	// 初始化微信客户端
	//    appId：应用ID
	//    MchID：商户ID
	//    ApiKey：Key值
	//    isProd：是否是正式环境（企业转账到个人账户，默认正式环境）
	client := wechat.NewClient("wxdaa2ab9ef87b5497", "1368139502", "GFDS8j98rewnmgl45wHTt980jg543abc", false)

	err := client.AddCertPkcs12FileContent([]byte("apiclient_cert.p12 content"))
	if err != nil {
		xlog.Error(err)
		return
	}

	// 初始化参数结构体
	bm := make(pay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32)).
		Set("partner_trade_no", util.RandomString(32)).
		Set("openid", "o0Df70H2Q0fY8JXh1aFPIRyOBgu8").
		Set("check_name", "FORCE_CHECK"). // NO_CHECK：不校验真实姓名 , FORCE_CHECK：强校验真实姓名
		Set("re_user_name", "付明明").       // 收款用户真实姓名。 如果check_name设置为FORCE_CHECK，则必填用户真实姓名
		Set("amount", 30).                // 企业付款金额，单位为分
		Set("desc", "测试转账").              // 企业付款备注，必填。注意：备注中的敏感词会被转成字符*
		Set("spbill_create_ip", "127.0.0.1")

	// 企业向微信用户个人付款（不支持沙箱环境）
	//    body：参数Body
	wxRsp, err := client.Transfer(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)
}
