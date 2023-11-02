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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xhttp"
	"github.com/rwscode/payutil/pkg/xlog"
)

// 获取AccessToken（Get an access token）
// 文档：https://developer.paypal.com/docs/api/reference/get-an-access-token
func (c *Client) GetAccessToken() (token *AccessToken, err error) {
	var (
		baseUrl = baseUrlProd
		url     string
	)
	if !c.IsProd {
		baseUrl = baseUrlSandbox
	}
	url = baseUrl + getAccessToken
	// Authorization
	authHeader := AuthorizationPrefixBasic + base64.StdEncoding.EncodeToString([]byte(c.Clientid+":"+c.Secret))
	// Request
	httpClient := xhttp.NewClient()
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	// Body
	bm := make(pay.BodyMap)
	bm.Set("grant_type", "client_credentials")
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	res, bs, err := httpClient.Type(xhttp.TypeForm).Post(url).SendBodyMap(bm).EndBytes(c.ctx)
	if err != nil {
		return nil, err
	}
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	token = new(AccessToken)
	if err = json.Unmarshal(bs, token); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	c.Appid = token.Appid
	c.AccessToken = token.AccessToken
	c.ExpiresIn = token.ExpiresIn
	return token, nil
}
