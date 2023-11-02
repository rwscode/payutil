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

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xlog"
)

func TestClient_UserCertifyOpenInit(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("outer_order_no", "ZGYD201809132323000001234")
	// 认证场景码：FACE：多因子人脸认证，CERT_PHOTO：多因子证照认证，CERT_PHOTO_FACE ：多因子证照和人脸认证，SMART_FACE：多因子快捷认证
	bm.Set("biz_code", "FACE")
	// 需要验证的身份信息参数，格式为json
	identity := make(map[string]string)
	identity["identity_type"] = "CERT_INFO"
	identity["cert_type"] = "IDENTITY_CARD"
	identity["cert_name"] = "张三"
	identity["cert_no"] = "310123199012301234"
	bm.Set("identity_param", identity)
	// 商户个性化配置，格式为json
	merchant := make(map[string]string)
	merchant["return_url"] = "https://www.fmm.ink"
	bm.Set("merchant_config", merchant)

	// 发起请求
	aliRsp, err := client.UserCertifyOpenInit(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserCertifyOpenCertify(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	// 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
	bm.Set("certify_id", "53827f9d085b3ce43938c6e5915b4729")

	// 发起请求
	certifyUrl, err := client.UserCertifyOpenCertify(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("certifyUrl:", certifyUrl)
}

func TestClient_UserCertifyOpenQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	// 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
	bm.Set("certify_id", "OC201809253000000393900404029253")

	// 发起请求
	aliRsp, err := client.UserCertifyOpenQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response.Passed:", aliRsp.Response.Passed)
}

func TestUserAgreementExecutionplanModify(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("agreement_no", "20185909000458725113")
	bm.Set("deduct_time", "2019-05-12")
	bm.Set("memo", "用户已购买半年包，需延期扣款时间")

	// 发起请求
	aliRsp, err := client.UserAgreementExecutionplanModify(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAgreementTransfer(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("agreement_no", "20170322450983769228")
	bm.Set("target_product_code", "CYCLE_PAY_AUTH_P")
	bm.SetBodyMap("period_rule_params", func(bm pay.BodyMap) {
		bm.Set("period_type", "DAY")
		bm.Set("period", "3")
		bm.Set("execute_time", "20190-01-23")
		bm.Set("single_amount", "10.99")
	})

	// 发起请求
	aliRsp, err := client.UserAgreementTransfer(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserTwostageCommonUse(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("dynamic_id", "286861260475412123")
	bm.Set("sence_no", "20170718xxxxxxxx")
	bm.Set("pay_pid", "2088702093900999")

	// 发起请求
	aliRsp, err := client.UserTwostageCommonUse(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAuthZhimaorgIdentityApply(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("cert_type", "NATIONAL_LEGAL")
	bm.Set("cert_no", "330701199901011311")
	bm.Set("name", "中国移动有限公司")

	// 发起请求
	aliRsp, err := client.UserAuthZhimaorgIdentityApply(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserCharityRecordexistQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("partner_id", "2088111122223333")
	bm.Set("user_id", "2088111122223333")

	// 发起请求
	aliRsp, err := client.UserCharityRecordexistQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAlipaypointSend(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("budget_code", "02559A591572")
	bm.Set("partner_biz_no", "011022222222212")
	bm.Set("point_amount", "1")

	// 发起请求
	aliRsp, err := client.UserAlipaypointSend(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestMemberDataIsvCreate(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("member_card_id", "2014323100009")
	bm.Set("member_source", "alipay")
	bm.Set("member_status", "1")
	bm.Set("gmt_merber_card_create", "2017-02-17 20:11:54")
	bm.Set("parter_id", "2088902248579233")

	// 发起请求
	aliRsp, err := client.MemberDataIsvCreate(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserFamilyArchiveQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("archive_token", "2020050200286001170017000004861")

	// 发起请求
	aliRsp, err := client.UserFamilyArchiveQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserFamilyArchiveInitialize(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("out_biz_no", "d0f003fdf57b4983bae5a0d1af2e7744")
	bm.Set("template_id", "2020050200286001170017000004861")
	bm.Set("redirect_uri", "https://www.alipay.com")

	// 发起请求
	aliRsp, err := client.UserFamilyArchiveInitialize(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserCertdocCertverifyPreconsult(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("user_name", "张三")
	bm.Set("cert_type", "IDENTITY_CARD")
	bm.Set("cert_no", "230100199901010001")

	// 发起请求
	aliRsp, err := client.UserCertdocCertverifyPreconsult(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserCertdocCertverifyConsult(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("verify_id", "671ffcda5447bc87e9ed2f669eb143d4")
	// 发起请求
	aliRsp, err := client.UserCertdocCertverifyConsult(ctx, bm, "auth_token")
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserFamilyShareZmgoInitialize(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("user_id", "2088161820676973")
	bm.Set("scene_id", "family_health_card")
	bm.Set("template_id", "2019112500020903940000454087")
	bm.Set("out_request_no", "d0f003fdf57b4983bae5a0d1af2e7744")
	// 发起请求
	aliRsp, err := client.UserFamilyShareZmgoInitialize(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserDtbankQrcodedataQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("data_date", "20210106")
	bm.Set("qrcode_id", "QRC884QRC00014990")
	bm.Set("qrcode_out_id", "18448-000006")
	// 发起请求
	aliRsp, err := client.UserDtbankQrcodedataQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAlipaypointBudgetlibQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("budget_code", "20201107050844")
	// 发起请求
	aliRsp, err := client.UserAlipaypointBudgetlibQuery(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}