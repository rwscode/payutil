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

// 图片上传（营销专用）
//
//	注意：图片不能超过2MB
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter9_0_1.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter9_0_1.shtml
func (c *ClientV3) V3FavorMediaUploadImage(ctx context.Context, fileName, fileSha256 string, img *util.File) (wxRsp *MarketMediaUploadRsp, err error) {
	bmFile := make(pay.BodyMap)
	bmFile.Set("filename", fileName).Set("sha256", fileSha256)
	authorization, err := c.authorization(MethodPost, v3FavorMediaUploadImage, bmFile)
	if err != nil {
		return nil, err
	}

	bm := make(pay.BodyMap)
	bm.SetBodyMap("meta", func(bm pay.BodyMap) {
		bm.Set("filename", fileName).Set("sha256", fileSha256)
	}).SetFormFile("file", img)
	res, si, bs, err := c.doProdPostFile(ctx, bm, v3FavorMediaUploadImage, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &MarketMediaUploadRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MarketMediaUpload)
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
