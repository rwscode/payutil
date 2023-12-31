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

// alipay.trade.royalty.relation.bind(分账关系绑定)
// 文档地址：https://opendocs.alipay.com/open/02c7hq
func (a *Client) TradeRelationBind(ctx context.Context, bm pay.BodyMap) (aliRsp *TradeRelationBindResponse, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.bind"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationBindResponse)
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

// alipay.trade.royalty.relation.unbind(分账关系解绑)
// 文档地址：https://opendocs.alipay.com/open/02c7hr
func (a *Client) TradeRelationUnbind(ctx context.Context, bm pay.BodyMap) (aliRsp *TradeRelationUnbindResponse, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.unbind"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationUnbindResponse)
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

// alipay.trade.royalty.relation.batchquery(分账关系查询)
// 文档地址：https://opendocs.alipay.com/open/02c7hs
func (a *Client) TradeRelationBatchQuery(ctx context.Context, bm pay.BodyMap) (aliRsp *TradeRelationBatchQueryResponse, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationBatchQueryResponse)
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

// alipay.trade.order.settle(统一收单交易结算接口)
// 文档地址：https://opendocs.alipay.com/open/02j2bt
func (a *Client) TradeOrderSettle(ctx context.Context, bm pay.BodyMap) (aliRsp *TradeOrderSettleResponse, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "royalty_parameters")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.settle"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleResponse)
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

// alipay.trade.order.settle.query(交易分账查询接口)
// 文档地址：https://opendocs.alipay.com/open/02pj6l
func (a *Client) TradeOrderSettleQuery(ctx context.Context, bm pay.BodyMap) (aliRsp *TradeOrderSettleQueryResponse, err error) {
	err = bm.CheckEmptyError("settle_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.settle.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleQueryResponse)
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
