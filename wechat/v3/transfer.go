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

// 发起商家转账API
//
//	注意：入参加密字段数据加密：client.V3EncryptText()
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_1.shtml
func (c *ClientV3) V3Transfer(ctx context.Context, bm pay.BodyMap) (*TransferRsp, error) {
	authorization, err := c.authorization(MethodPost, v3Transfer, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3Transfer, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Transfer)
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

// 发起批量转账API（服务商）
//
//	注意：入参加密字段数据加密：client.V3EncryptText()
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_1.shtml
func (c *ClientV3) V3PartnerTransfer(ctx context.Context, bm pay.BodyMap) (*TransferRsp, error) {
	authorization, err := c.authorization(MethodPost, v3PartnerTransfer, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3PartnerTransfer, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(Transfer)
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

// 微信批次单号查询批次单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_2.shtml
func (c *ClientV3) V3TransferQuery(ctx context.Context, batchId string, bm pay.BodyMap) (*TransferQueryRsp, error) {
	url := fmt.Sprintf(v3TransferQuery, batchId)
	bm.Remove("batch_id")
	uri := url + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferQuery)
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

// 微信批次单号查询批次单API（服务商）
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_2.shtml
func (c *ClientV3) V3PartnerTransferQuery(ctx context.Context, batchId string, bm pay.BodyMap) (*PartnerTransferQueryRsp, error) {
	url := fmt.Sprintf(v3PartnerTransferQuery, batchId)
	bm.Remove("batch_id")
	uri := url + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &PartnerTransferQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerTransferQuery)
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

// 微信明细单号查询明细单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_3.shtml
func (c *ClientV3) V3TransferDetail(ctx context.Context, batchId, detailId string) (*TransferDetailRsp, error) {
	url := fmt.Sprintf(v3TransferDetail, batchId, detailId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferDetailQuery)
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

// 微信明细单号查询明细单API（服务商）
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_3.shtml
func (c *ClientV3) V3PartnerTransferDetail(ctx context.Context, batchId, detailId string) (*PartnerTransferDetailRsp, error) {
	url := fmt.Sprintf(v3PartnerTransferDetail, batchId, detailId)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &PartnerTransferDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerTransferDetail)
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

// Deprecated
// 推荐直接使用 client.V3TransferDetail() 方法
func (c *ClientV3) V3TransferDetailQuery(ctx context.Context, batchId, detailId string) (*TransferDetailRsp, error) {
	return c.V3TransferDetail(ctx, batchId, detailId)
}

// 商家批次单号查询批次单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_5.shtml
func (c *ClientV3) V3TransferMerchantQuery(ctx context.Context, outBatchNo string, bm pay.BodyMap) (*TransferMerchantQueryRsp, error) {
	url := fmt.Sprintf(v3TransferMerchantQuery, outBatchNo)
	bm.Remove("out_batch_no")
	uri := url + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferMerchantQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferMerchantQuery)
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

// 商家批次单号查询批次单API（服务商）
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_4.shtml
func (c *ClientV3) V3PartnerTransferMerchantQuery(ctx context.Context, outBatchNo string, bm pay.BodyMap) (*PartnerTransferMerchantQueryRsp, error) {
	url := fmt.Sprintf(v3PartnerTransferMerchantQuery, outBatchNo)
	bm.Remove("out_batch_no")
	uri := url + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &PartnerTransferMerchantQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerTransferMerchantQuery)
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

// 商家明细单号查询明细单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_6.shtml
func (c *ClientV3) V3TransferMerchantDetail(ctx context.Context, outBatchNo, outDetailNo string) (*TransferMerchantDetailRsp, error) {
	url := fmt.Sprintf(v3TransferMerchantDetail, outBatchNo, outDetailNo)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferMerchantDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferMerchantDetail)
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

// 商家明细单号查询明细单API（服务商）
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_5.shtml
func (c *ClientV3) V3PartnerTransferMerchantDetail(ctx context.Context, outBatchNo, outDetailNo string) (*PartnerTransferMerchantDetailRsp, error) {
	url := fmt.Sprintf(v3PartnerTransferMerchantDetail, outBatchNo, outDetailNo)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &PartnerTransferMerchantDetailRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerTransferMerchantDetail)
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

// Deprecated
// 推荐直接使用 client.V3TransferMerchantDetail() 方法
func (c *ClientV3) V3TransferMerchantDetailQuery(ctx context.Context, outBatchNo, outDetailNo string) (*TransferMerchantDetailRsp, error) {
	return c.V3TransferMerchantDetail(ctx, outBatchNo, outDetailNo)
}

// 转账电子回单申请受理API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_7.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_6.shtml
func (c *ClientV3) V3TransferReceipt(ctx context.Context, outBatchNo string) (*TransferReceiptRsp, error) {
	bm := make(pay.BodyMap)
	bm.Set("out_batch_no", outBatchNo)

	authorization, err := c.authorization(MethodPost, v3TransferReceipt, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3TransferReceipt, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferReceiptRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferReceipt)
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

// 查询转账电子回单API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_8.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_7.shtml
func (c *ClientV3) V3TransferReceiptQuery(ctx context.Context, outBatchNo string) (*TransferReceiptQueryRsp, error) {
	url := fmt.Sprintf(v3TransferReceiptQuery, outBatchNo)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferReceiptQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferReceiptQuery)
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

// 转账明细电子回单受理API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_9.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_8.shtml
func (c *ClientV3) V3TransferDetailReceipt(ctx context.Context, bm pay.BodyMap) (*TransferDetailReceiptRsp, error) {
	authorization, err := c.authorization(MethodPost, v3TransferDetailReceipt, bm)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdPost(ctx, bm, v3TransferDetailReceipt, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferDetailReceiptRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferDetailReceipt)
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

// 查询转账明细电子回单受理结果API
//
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_10.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter4_3_9.shtml
func (c *ClientV3) V3TransferDetailReceiptQuery(ctx context.Context, bm pay.BodyMap) (*TransferDetailReceiptQueryRsp, error) {
	uri := v3TransferDetailReceiptQuery + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &TransferDetailReceiptQueryRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(TransferDetailReceiptQuery)
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
