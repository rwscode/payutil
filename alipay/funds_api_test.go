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
	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xlog"
	"testing"
)

func TestFundTransUniTransfer(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("out_biz_no", "201806300011232301").
		Set("trans_amount", "0.01").
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		SetBodyMap("payee_info", func(bm pay.BodyMap) {
			bm.Set("identity", "85411418@qq.com")
			bm.Set("identity_type", "ALIPAY_LOGON_ID")
		})

	aliRsp, err := client.FundTransUniTransfer(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundAccountQuery(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("alipay_user_id", "2088301409188095") /*.Set("account_type", "ACCTRANS_ACCOUNT")*/

	aliRsp, err := client.FundAccountQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundTransCommonQuery(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("order_id", "20190801110070000006380000250621")

	aliRsp, err := client.FundTransCommonQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundTransOrderQuery(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("out_biz_no", "201806300011232301")

	aliRsp, err := client.FundTransOrderQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundAuthOrderAppFreeze(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("out_order_no", "8077735255938023").
		Set("out_request_no", "8077735255938032").
		Set("order_title", "预授权冻结").
		Set("amount", "0.01").
		Set("product_code", "PRE_AUTH_ONLINE")

	aliRsp, err := client.FundAuthOrderAppFreeze(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", aliRsp)
}

func TestClient_FundTransPagePay(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("out_biz_no", "2018062800001").
		Set("trans_amount", "8.88").
		Set("product_code", "STD_APP_TRANSFER").
		Set("biz_scene", "PARTY_MEMBERSHIP_DUES").
		SetBodyMap("payee_info", func(b pay.BodyMap) {
			b.Set("identity", "208812*****41234").
				Set("identity_type", "ALIPAY_USER_ID").
				Set("name", "黄龙国际有限公司")
		})

	aliRsp, err := client.FundTransPagePay(ctx, bm)
	if err != nil {
		// xlog.Errorf("client.FundTransPagePay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
