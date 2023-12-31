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
	"net/http"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
)

// 合单APP下单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_1.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_1.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_3_1.shtml
func (c *ClientV3) V3CombineTransactionApp(ctx context.Context, bm pay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("combine_mchid") == util.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayApp, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombinePayApp, authorization)
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

// 合单JSAPI/小程序下单API
//
//	Code = 0 is success
//	商户JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_3.shtml
//	商户小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_4.shtml
//	服务商JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_3.shtml
//	服务商小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_4.shtml
//	电商JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_3_3.shtml
//	电商小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_3_4.shtml
func (c *ClientV3) V3CombineTransactionJsapi(ctx context.Context, bm pay.BodyMap) (wxRsp *PrepayRsp, err error) {
	if bm.GetString("combine_mchid") == util.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayJsapi, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombinePayJsapi, authorization)
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

// 合单Native下单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_5.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_5.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_3_5.shtml
func (c *ClientV3) V3CombineTransactionNative(ctx context.Context, bm pay.BodyMap) (wxRsp *NativeRsp, err error) {
	if bm.GetString("combine_mchid") == util.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombineNative, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombineNative, authorization)
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

// 合单H5下单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_2.shtml
func (c *ClientV3) V3CombineTransactionH5(ctx context.Context, bm pay.BodyMap) (wxRsp *H5Rsp, err error) {
	if bm.GetString("combine_mchid") == util.NULL {
		bm.Set("combine_mchid", c.Mchid)
	}
	authorization, err := c.authorization(MethodPost, v3CombinePayH5, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3CombinePayH5, authorization)
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

// 合单查询订单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_11.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_11.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_3_11.shtml
func (c *ClientV3) V3CombineQueryOrder(ctx context.Context, traderNo string) (wxRsp *CombineQueryOrderRsp, err error) {
	uri := fmt.Sprintf(v3CombineQuery, traderNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp = &CombineQueryOrderRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(CombineQueryOrder)
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

// 合单关闭订单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter5_1_12.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter5_1_12.shtml
//	电商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_3_12.shtml
func (c *ClientV3) V3CombineCloseOrder(ctx context.Context, tradeNo string, bm pay.BodyMap) (wxRsp *CloseOrderRsp, err error) {
	url := fmt.Sprintf(v3CombineClose, tradeNo)
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
