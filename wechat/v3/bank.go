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
	"net/url"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
)

// 获取对私银行卡号开户银行
//
//	注意：accountNo 需此方法加密：client.V3EncryptText()
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_1.shtml
func (c *ClientV3) V3BankSearchBank(ctx context.Context, accountNo string) (wxRsp *BankSearchBankRsp, err error) {
	uri := v3BankSearchBank + "?account_number=" + url.QueryEscape(accountNo)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchBankRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchBank)
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

// 查询支持个人业务的银行列表
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_2.shtml
func (c *ClientV3) V3BankSearchPersonalList(ctx context.Context, limit, offset int) (wxRsp *BankSearchPersonalListRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := v3BankSearchPersonalList + "?limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchPersonalListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchList)
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

// 查询支持对公业务的银行列表
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_3.shtml
func (c *ClientV3) V3BankSearchCorporateList(ctx context.Context, limit, offset int) (wxRsp *BankSearchCorporateListRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := v3BankSearchCorporateList + "?limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchCorporateListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchList)
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

// 查询省份列表
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_4.shtml
func (c *ClientV3) V3BankSearchProvinceList(ctx context.Context) (wxRsp *BankSearchProvinceListRsp, err error) {
	authorization, err := c.authorization(MethodGet, v3BankSearchProvinceList, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, v3BankSearchProvinceList, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchProvinceListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchProvince)
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

// 查询城市列表
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_5.shtml
func (c *ClientV3) V3BankSearchCityList(ctx context.Context, provinceCode int) (wxRsp *BankSearchCityListRsp, err error) {
	url := fmt.Sprintf(v3BankSearchCityList, provinceCode)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, url, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchCityListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchCity)
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

// 查询支行列表
//
//	Code = 0 is success
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/Offline/apis/chapter11_2_6.shtml
func (c *ClientV3) V3BankSearchBranchList(ctx context.Context, bankAliasCode string, cityCode, limit, offset int) (wxRsp *BankSearchBranchListRsp, err error) {
	if limit == 0 {
		limit = 20
	}
	uri := fmt.Sprintf(v3BankSearchBranchList, bankAliasCode) + "?city_code=" + util.Int2String(cityCode) + "&limit=" + util.Int2String(limit) + "&offset=" + util.Int2String(offset)
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	wxRsp = &BankSearchBranchListRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(BankSearchBranch)
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
