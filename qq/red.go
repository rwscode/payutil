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

/*
	QQ 现金红包
	文档：https://qpay.qq.com/buss/wiki/221/1219
*/

package qq

import (
	"context"
	"encoding/xml"
	"fmt"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
)

// SendCashRed 创建现金红包
// 注意：如已使用client.AddCertFilePath()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
// 文档：https://qpay.qq.com/buss/wiki/221/1220
func (q *Client) SendCashRed(ctx context.Context, bm pay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (qqRsp *SendCashRedResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("charset", "nonce_str", "mch_billno", "mch_name", "re_openid",
		"total_amount", "total_num", "wishing", "act_name", "icon_id", "min_value", "max_value")
	if err != nil {
		return nil, err
	}
	tlsConfig, err := q.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQRed(ctx, bm, createCashRed, tlsConfig)
	if err != nil {
		return nil, err
	}
	qqRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}

// DownloadRedListFile 对账单下载
//
//	注意：data类型为int类型，例如：date=20200909，2020年9月9日
//	文档：https://qpay.qq.com/buss/wiki/221/1224
func (q *Client) DownloadRedListFile(ctx context.Context, bm pay.BodyMap) (qqRsp string, err error) {
	err = bm.CheckEmptyError("date")
	if err != nil {
		return pay.NULL, err
	}
	bs, err := q.doQQGet(ctx, bm, redFileDown, SignType_MD5)
	if err != nil {
		return util.NULL, err
	}
	return string(bs), nil
}

// QueryRedInfo 查询红包详情
//
//	文档：https://qpay.qq.com/buss/wiki/221/2174
func (q *Client) QueryRedInfo(ctx context.Context, bm pay.BodyMap) (qqRsp *QueryRedInfoResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "listid")
	if err != nil {
		return nil, err
	}
	bs, err := q.doQQRed(ctx, bm, queryRedInfo, nil)
	if err != nil {
		return nil, err
	}
	qqRsp = new(QueryRedInfoResponse)
	if err = xml.Unmarshal(bs, qqRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	return qqRsp, nil
}
