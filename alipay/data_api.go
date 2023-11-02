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
)

// Deprecated
// 支付宝已不再支持
// alipay.data.bill.balance.query(支付宝商家账户当前余额查询)
// 文档地址：https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
func (a *Client) DataBillBalanceQuery(ctx context.Context, bm pay.BodyMap) (aliRsp *DataBillBalanceQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.bill.balance.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillBalanceQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.data.dataservice.bill.downloadurl.query(查询对账单下载地址)
// 文档地址：https://opendocs.alipay.com/open/02e7gr
func (a *Client) DataBillDownloadUrlQuery(ctx context.Context, bm pay.BodyMap) (aliRsp *DataBillDownloadUrlQueryResponse, err error) {
	err = bm.CheckEmptyError("bill_type", "bill_date")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.data.dataservice.bill.downloadurl.query"); err != nil {
		return nil, err
	}
	aliRsp = new(DataBillDownloadUrlQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", pay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
