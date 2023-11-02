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
	"github.com/rwscode/payutil/alipay/cert"
	"github.com/rwscode/payutil/pkg/xlog"
)

func TestClient_SystemOauthToken(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("grant_type", "authorization_code")
	bm.Set("code", "3a06216ac8f84b8c93507bb9774bWX11")

	// 发起请求
	aliRsp, err := client.SystemOauthToken(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp:", aliRsp.Response.AccessToken)
	xlog.Debug("aliRsp:", aliRsp.SignData)
}

func TestClient_OpenAuthTokenApp(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("grant_type", "authorization_code").
		Set("code", "866185490c4e40efa9f71efea6766X02")

	// 发起请求
	aliRsp, err := client.OpenAuthTokenApp(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserInfoAuth(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	// 接口权限值，目前只支持auth_user和auth_base两个值。具体说明看文档介绍
	bm.Set("scopes", []string{"auth_user"}).
		Set("state", "init")

	// 发起请求
	aliRsp, err := client.UserInfoAuth(ctx, bm)
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserInfoShare(t *testing.T) {
	// 发起请求
	aliRsp, err := client.UserInfoShare(ctx, "auth_token")
	if err != nil {
		if bizErr, ok := IsBizError(err); ok {
			xlog.Errorf("%+v", bizErr)
			// do something
			return
		}
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

	// 同步返回验签
	ok, err := VerifySyncSignWithCert(cert.AlipayPublicContentRSA2, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		xlog.Error(err)
	}
	xlog.Debug("ok:", ok)
}

func TestClient_PublicCertDownload(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("alipay_cert_sn", "52c63ed47b57c049b4bc9bea9da02c2a")

	// 发起请求
	aliRsp, err := client.PublicCertDownload(ctx, bm)
	if err != nil {
		// xlog.Errorf("client.UserInfoShare(),error:%+v", err)
		return
	}
	xlog.Debugf("aliRsp.Response.AlipayCertContent:\n %s", aliRsp.Response.AlipayCertContent)
}
