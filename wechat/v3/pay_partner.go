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
	"errors"
	"fmt"
	"net/http"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
)

// （服务商、电商模式）APP下单API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_2_1.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_2_1.shtml
func (c *ClientV3) V3PartnerTransactionApp(ctx context.Context, bm pay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerPayApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerPayApp, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// （服务商、电商模式）JSAPI/小程序下单API
//
//	Code = 0 is success
//	服务商JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_1.shtml
//	服务商小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_1.shtml
//	电商JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_2_2.shtml
//	电商小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_2_3.shtml
func (c *ClientV3) V3PartnerTransactionJsapi(ctx context.Context, bm pay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerJsapi, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &PrepayRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Prepay)
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

// （服务商、电商模式）Native下单API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_1.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_2_12.shtml
func (c *ClientV3) V3PartnerTransactionNative(ctx context.Context, bm pay.BodyMap) (wxRsp *NativeRsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerNative, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerNative, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &NativeRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Native)
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

// （服务商模式）H5下单API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_3_1.shtml
func (c *ClientV3) V3PartnerTransactionH5(ctx context.Context, bm pay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("sp_mchid") == util.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3ApiPartnerH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3ApiPartnerH5, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &H5Rsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(H5Url)
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

// （服务商、电商模式）查询订单API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_2.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_2_5.shtml
func (c *ClientV3) V3PartnerQueryOrder(ctx context.Context, orderNoType OrderNoType, orderNo string, bm pay.BodyMap) (wxRsp *PartnerQueryOrderRsp, err error) {
	var uri string
	if bm.GetString("sp_mchid") == pay.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	switch orderNoType {
	case TransactionId:
		uri = fmt.Sprintf(v3ApiPartnerQueryOrderTransactionId, orderNo) + "?" + bm.EncodeURLParams()
	case OutTradeNo:
		uri = fmt.Sprintf(v3ApiPartnerQueryOrderOutTradeNo, orderNo) + "?" + bm.EncodeURLParams()
	default:
		return nil, errors.New("unsupported order number type")
	}
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &PartnerQueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerQueryOrder)
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

// （服务商、电商模式）关单API
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_3.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_2_6.shtml
func (c *ClientV3) V3PartnerCloseOrder(ctx context.Context, tradeNo string, bm pay.BodyMap) (wxRsp *CloseOrderRsp, err error) {
	url := fmt.Sprintf(v3ApiPartnerCloseOrder, tradeNo)
	if bm.GetString("sp_mchid") == pay.NULL {
		bm.Set("sp_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, url, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CloseOrderRsp{Code: Success, SignInfo: si}
	if res.StatusCode != http.StatusNoContent {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
