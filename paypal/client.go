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
	"net/http"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
	"github.com/rwscode/payutil/pkg/xhttp"
	"github.com/rwscode/payutil/pkg/xlog"
)

// Client PayPal支付客
type Client struct {
	Clientid    string
	Secret      string
	Appid       string
	AccessToken string
	ExpiresIn   int
	bodySize    int // http response body size(MB), default is 10MB
	IsProd      bool
	ctx         context.Context
	DebugSwitch pay.DebugSwitch
}

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool) (client *Client, err error) {
	if clientid == util.NULL || secret == util.NULL {
		return nil, pay.MissPayPalInitParamErr
	}
	client = &Client{
		Clientid:    clientid,
		Secret:      secret,
		IsProd:      isProd,
		ctx:         context.Background(),
		DebugSwitch: pay.DebugOff,
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	return client, nil
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.bodySize = sizeMB
	}
}

func (c *Client) doPayPalGet(ctx context.Context, uri string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + uri
	if !c.IsProd {
		url = baseUrlSandbox + uri
	}
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_Url: %s", url)
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Get(url).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPost(ctx context.Context, bm pay.BodyMap, path string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + path
	if !c.IsProd {
		url = baseUrlSandbox + path
	}
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_RequestBody: %s", bm.JsonBody())
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Post(url).SendBodyMap(bm).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}

func (c *Client) doPayPalPatch(ctx context.Context, patchs []*Patch, path string) (res *http.Response, bs []byte, err error) {
	var url = baseUrlProd + path
	if !c.IsProd {
		url = baseUrlSandbox + path
	}
	httpClient := xhttp.NewClient()
	if c.bodySize > 0 {
		httpClient.SetBodySize(c.bodySize)
	}
	authHeader := AuthorizationPrefixBearer + c.AccessToken
	if c.DebugSwitch == pay.DebugOn {
		jb, _ := json.Marshal(patchs)
		xlog.Debugf("PayPal_RequestBody: %s", string(jb))
		xlog.Debugf("PayPal_Authorization: %s", authHeader)
	}
	httpClient.Header.Add(HeaderAuthorization, authHeader)
	httpClient.Header.Add("Accept", "*/*")
	res, bs, err = httpClient.Type(xhttp.TypeJSON).Patch(url).SendStruct(patchs).EndBytes(ctx)
	if err != nil {
		return nil, nil, err
	}
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("PayPal_Response: %d > %s", res.StatusCode, string(bs))
		xlog.Debugf("PayPal_Headers: %#v", res.Header)
	}
	return res, bs, nil
}
