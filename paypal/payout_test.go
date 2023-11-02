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
@Author: wzy
@Time: 2022/6/8
*/
package paypal

import (
	"encoding/json"
	"testing"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xlog"
)

func TestCreateBatchPayout(t *testing.T) {
	receiver := "test-email@testemail.com"
	bm := make(pay.BodyMap)
	bm.SetBodyMap("sender_batch_header", func(bm pay.BodyMap) {
		bm.Set("sender_batch_id", "2022060811140003").
			Set("email_subject", "You have a payout!").
			Set("email_message", "You have received a payout! Thanks for using our service!")
	}).Set("items", []map[string]interface{}{
		{
			"recipient_type": "EMAIL",
			"amount": map[string]string{
				"value":    "5",
				"currency": "USD",
			},
			"note":           "Thanks for your verify",
			"sender_item_id": "20220608111304",
			"receiver":       receiver,
		},
	})
	xlog.Debug("bm", bm.JsonBody())
	ppRsp, err := client.CreateBatchPayout(ctx, bm)
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
	_rspJson, _ := json.MarshalIndent(ppRsp.Response, "", "\t")
	xlog.Debugf("ppRsp.Response: %s", _rspJson)
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestShowPayoutBatchDetails(t *testing.T) {
	ppRsp, err := client.ShowPayoutBatchDetails(ctx, "YSESATZEDPRY6", nil)
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
	for _, v := range ppRsp.Response.Items {
		_item, _ := json.MarshalIndent(v, "", "\t")
		xlog.Debugf("ppRsp.Response.PayoutItemDetail: \n%s", _item)
	}
	for _, v := range ppRsp.Response.Links {
		xlog.Debugf("ppRsp.Response.Links: %+v", v)
	}
}

func TestShowPayoutItemDetails(t *testing.T) {
	ppRsp, err := client.ShowPayoutItemDetails(ctx, "HGYYMWW7PRJKW")
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
	_rspJson, _ := json.MarshalIndent(ppRsp.Response, "", "\t")
	xlog.Debugf("ppRsp.Response: %s", _rspJson)
}

func TestCancelUnclaimedPayoutItem(t *testing.T) {
	ppRsp, err := client.CancelUnclaimedPayoutItem(ctx, "HGYYMWW7PRJKW")
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
	_rspJson, _ := json.MarshalIndent(ppRsp.Response, "", "\t")
	xlog.Debugf("ppRsp.Response: %s", _rspJson)
}
