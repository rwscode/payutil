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
	"testing"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xlog"
)

func TestAntMerchantShopModify(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)

	aliRsp, err := client.AntMerchantShopModify(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopModify(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopCreate(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.SetBodyMap("business_address", func(bm pay.BodyMap) {
		bm.Set("city_code", "37100")
		bm.Set("district_code", "371002")
		bm.Set("address", "万塘路18号黄龙时代广场B座")
		bm.Set("province_code", "310000")
	})
	bm.Set("shop_category", "B0001")
	bm.Set("store_id", "NO0001")
	bm.Set("shop_type", "01")
	bm.Set("ip_role_id", "2088301155943087")
	bm.Set("shop_name", "肯德基中关村店")

	aliRsp, err := client.AntMerchantShopCreate(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		xlog.Errorf("client.AntMerchantShopCreate(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopConsult(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.SetBodyMap("business_address", func(bm pay.BodyMap) {
		bm.Set("city_code", "37100")
		bm.Set("district_code", "371002")
		bm.Set("address", "万塘路18号黄龙时代广场B座")
		bm.Set("province_code", "310000")
	})
	bm.Set("shop_category", "B0001")
	bm.Set("store_id", "NO0001")
	bm.Set("shop_type", "01")
	bm.Set("ip_role_id", "2088301155943087")
	bm.Set("shop_name", "肯德基中关村店")

	aliRsp, err := client.AntMerchantShopConsult(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopConsult(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantOrderQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("order_id", "2017112200502000000004754299")

	aliRsp, err := client.AntMerchantOrderQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantOrderQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopQuery(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("shop_id", "2018011900502000000005124744")
	bm.Set("store_id", "NO0001")
	bm.Set("ip_role_id", "2088301155943087")

	aliRsp, err := client.AntMerchantShopQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestAntMerchantShopClose(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("shop_id", "2018011900502000000005124744")
	bm.Set("store_id", "NO0001")
	bm.Set("ip_role_id", "2088301155943087")

	aliRsp, err := client.AntMerchantShopClose(ctx, bm)
	if err != nil {
		xlog.Errorf("client.AntMerchantShopClose(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
