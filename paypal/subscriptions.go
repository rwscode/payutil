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
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// 创建订阅计划（CreateBillingPlan）
// Code = 0 is success
// 文档：https://developer.paypal.com/docs/api/subscriptions/v1/#plans_create
func (c *Client) CreateBillingPlan(ctx context.Context, bm pay.BodyMap) (ppRsp *CreateBillingRsp, err error) {
	if err = bm.CheckEmptyError("product_id", "billing_cycles"); err != nil {
		return nil, err
	}
	res, bs, err := c.doPayPalPost(ctx, bm, subscriptionCreate)
	if err != nil {
		return nil, err
	}
	ppRsp = &CreateBillingRsp{Code: Success}
	ppRsp.Response = new(BillingDetail)
	if err = json.Unmarshal(bs, ppRsp.Response); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	if res.StatusCode != http.StatusCreated {
		ppRsp.Code = res.StatusCode
		ppRsp.Error = string(bs)
		ppRsp.ErrorResponse = new(ErrorResponse)
		_ = json.Unmarshal(bs, ppRsp.ErrorResponse)
	}
	return ppRsp, nil
}
