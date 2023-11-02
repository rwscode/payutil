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
	"context"
	"encoding/json"
	"fmt"
	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
)

// alipay.merchant.item.file.upload(商品文件上传接口)
// 文档地址：https://opendocs.alipay.com/apis/api_4/alipay.merchant.item.file.upload
func (a *Client) MerchantItemFileUpload(ctx context.Context, file *util.File) (aliRsp *MerchantItemFileUploadRsp, err error) {
	bm := make(pay.BodyMap)
	bm.Set("scene", "SYNC_ORDER") // 素材固定值

	var bs []byte
	if bs, err = a.FileRequest(ctx, bm, file, "alipay.merchant.item.file.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(MerchantItemFileUploadRsp)
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, err
	}
	if aliRsp.Response != nil && aliRsp.Response.Code != "10000" {
		info := aliRsp.Response
		return aliRsp, fmt.Errorf(`{"code":"%s","msg":"%s","sub_code":"%s","sub_msg":"%s"}`, info.Code, info.Msg, info.SubCode, info.SubMsg)
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
