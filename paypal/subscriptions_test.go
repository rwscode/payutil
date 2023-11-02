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

package paypal

import (
	"testing"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xlog"
)

func TestClient_CreateBillingPlan(t *testing.T) {
	var billingCycles []*BillingCycles
	var billingCycle = &BillingCycles{
		Frequency: &Frequency{
			IntervalUnit:  "MONTH",
			IntervalCount: 1,
		},
		TenureType:  "REGULAR",
		Sequence:    1,
		TotalCycles: 0,
		PricingScheme: &PricingScheme{
			FixedPrice: &FixedPrice{
				Value:        "101",
				CurrencyCode: "USD",
			},
		},
	}
	billingCycles = append(billingCycles, billingCycle)

	// 创建 PayPal 支付订单
	bm := make(pay.BodyMap)
	bm.Set("product_id", "PROD-9TH539347F0791830").
		Set("name", "gopay").
		Set("billing_cycles", billingCycles).
		Set("description", "Monthly subscription for premium users").
		Set("payment_preferences", &PaymentPreferences{
			AutoBillOutstanding:     true,
			SetupFeeFailureAction:   "CONTINUE",
			PaymentFailureThreshold: 3,
		})

	xlog.Debug("bm：", bm.JsonBody())

	ppRsp, err := client.CreateBillingPlan(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	if ppRsp.Code != Success {
		xlog.Debugf("ppRsp.Code: %+v", ppRsp.Code)
		xlog.Debugf("ppRsp.Error: %+v", ppRsp.Error)
		xlog.Debugf("ppRsp.ErrorResponse: %+v", ppRsp.ErrorResponse)
		return
	}
	xlog.Debugf("ppRsp.Response: %+v", ppRsp.Response)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}
