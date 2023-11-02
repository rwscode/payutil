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

// 商圈积分同步
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_2.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_6_2.shtml
func (c *ClientV3) V3BusinessPointsSync(ctx context.Context, bm pay.BodyMap) (wxRsp *EmptyRsp, err error) {
	authorization, err := c.authorization(MethodPost, v3BusinessPointsSync, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3BusinessPointsSync, authorization)
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

// 商圈积分授权查询
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter8_6_4.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_6_4.shtml
func (c *ClientV3) V3BusinessAuthPointsQuery(ctx context.Context, appid, openid string) (*BusinessAuthPointsQueryRsp, error) {
	uri := fmt.Sprintf(v3BusinessAuthPointsQuery, openid) + "?appid=" + appid
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp := &BusinessAuthPointsQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BusinessAuthPointsQuery)
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
