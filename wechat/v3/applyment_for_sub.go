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
	"encoding/json"
	"fmt"
	pay "github.com/rwscode/payutil"
	"net/http"
)

// 提交申请单API
// 注意：本接口会提交一些敏感信息，需调用 client.V3EncryptText() 进行加密
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_1.shtml
func (c *ClientV3) V3Apply4SubSubmit(ctx context.Context, bm pay.BodyMap) (*Apply4SubSubmitRsp, error) {
	if err := bm.CheckEmptyError("business_code", "contact_info", "subject_info", "business_info", "settlement_info", "bank_account_info"); err != nil {
		return nil, err
	}
	authorization, err := c.authorization(MethodPost, v3Apply4SubSubmit, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3Apply4SubSubmit, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &Apply4SubSubmitRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubSubmit)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 通过业务申请编号查询申请状态API
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (c *ClientV3) V3Apply4SubQueryByBusinessCode(ctx context.Context, businessCode string) (*Apply4SubQueryRsp, error) {
	uri := fmt.Sprintf(v3Apply4SubQueryByBusinessCode, businessCode)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &Apply4SubQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 通过申请单号查询申请状态API
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_2.shtml
func (c *ClientV3) V3Apply4SubQueryByApplyId(ctx context.Context, applyId string) (*Apply4SubQueryRsp, error) {
	uri := fmt.Sprintf(v3Apply4SubQueryByApplyId, applyId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &Apply4SubQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubQuery)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 修改结算账号 API
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_3.shtml
// 电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_1_4.shtml
func (c *ClientV3) V3Apply4SubModifySettlement(ctx context.Context, bm pay.BodyMap) (*EmptyRsp, error) {
	if err := bm.CheckEmptyError("sub_mchid", "account_type", "account_bank", "account_number"); err != nil {
		return nil, err
	}
	postUrl := fmt.Sprintf(v3Apply4SubModifySettlement, bm["sub_mchid"])
	bm.Remove("sub_mchid")
	authorization, err := c.authorization(MethodPost, postUrl, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, postUrl, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询结算账户 API
// Code = 0 is success
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter11_1_4.shtml
// 电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_1_5.shtml
func (c *ClientV3) V3Apply4SubQuerySettlement(ctx context.Context, subMchId string) (*Apply4SubQuerySettlementRsp, error) {
	uri := fmt.Sprintf(v3Apply4SubQuerySettlement, subMchId)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &Apply4SubQuerySettlementRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Apply4SubQuerySettlement)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
