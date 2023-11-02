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

// 点金计划管理API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_1.shtml
func (c *ClientV3) V3GoldPlanManage(ctx context.Context, bm pay.BodyMap) (wxRsp *GoldPlanManageRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GoldPlanManage, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanManage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &GoldPlanManageRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(GoldPlanManage)
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

// 商家小票管理API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_2.shtml
func (c *ClientV3) V3GoldPlanBillManage(ctx context.Context, bm pay.BodyMap) (wxRsp *GoldPlanManageRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GoldPlanBillManage, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanBillManage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &GoldPlanManageRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(GoldPlanManage)
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

// 同业过滤标签管理API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_3.shtml
func (c *ClientV3) V3GoldPlanFilterManage(ctx context.Context, bm pay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3GoldPlanFilterManage, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanFilterManage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 开通广告展示API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_4.shtml
func (c *ClientV3) V3GoldPlanOpenAdShow(ctx context.Context, bm pay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPATCH, v3GoldPlanOpenAdShow, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPatch(ctx, bm, v3GoldPlanOpenAdShow, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 关闭广告展示API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_5_5.shtml
func (c *ClientV3) V3GoldPlanCloseAdShow(ctx context.Context, bm pay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPATCH, v3GoldPlanCloseAdShow, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3GoldPlanCloseAdShow, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &EmptyRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}