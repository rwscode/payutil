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

package apple

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rwscode/payutil/pkg/xhttp"
)

// GetTransactionHistory
// Doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
func GetTransactionHistory(ctx context.Context, signConfig *SignConfig, originalTransactionId string, bm pay.BodyMap, sandbox bool) (rsp *TransactionHistoryRsp, err error) {
	uri := hostUrl + fmt.Sprintf(getTransactionHistory, originalTransactionId) + "?" + bm.EncodeURLParams()
	if sandbox {
		uri = sandBoxHostUrl + fmt.Sprintf(getTransactionHistory, originalTransactionId) + "?" + bm.EncodeURLParams()
	}
	token, err := generatingToken(ctx, signConfig)
	if err != nil {
		return nil, err
	}
	cli := xhttp.NewClient()
	cli.Header.Set("Authorization", "Bearer "+token)
	res, bs, err := cli.Type(xhttp.TypeJSON).Get(uri).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http.stauts_coud = %d", res.StatusCode)
	}
	rsp = &TransactionHistoryRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", pay.UnmarshalErr, err, string(bs))
	}
	return
}
