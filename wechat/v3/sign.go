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
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
	"github.com/rwscode/payutil/pkg/xlog"
	"github.com/rwscode/payutil/pkg/xpem"
)

// Deprecated
// 推荐使用 wechat.V3VerifySignByPK()
func V3VerifySign(timestamp, nonce, signBody, sign, wxPubKeyContent string) (err error) {
	publicKey, err := xpem.DecodePublicKey([]byte(wxPubKeyContent))
	if err != nil {
		return err
	}
	str := timestamp + "\n" + nonce + "\n" + signBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", pay.VerifySignatureErr, err)
	}
	return nil
}

// 推荐直接开启自动同步验签功能
// 微信V3 版本验签（同步）
// wxPublicKey：微信平台证书公钥内容，通过 client.WxPublicKeyMap() 获取，然后根据 signInfo.HeaderSerial 获取相应的公钥
func V3VerifySignByPK(timestamp, nonce, signBody, sign string, wxPublicKey *rsa.PublicKey) (err error) {
	str := timestamp + "\n" + nonce + "\n" + signBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(sign)

	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(wxPublicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", pay.VerifySignatureErr, err)
	}
	return nil
}

// PaySignOfJSAPI 获取 JSAPI paySign
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml
func (c *ClientV3) PaySignOfJSAPI(appid, prepayid string) (jsapi *JSAPIPayParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.RandomString(32)
	pkg := "prepay_id=" + prepayid

	_str := appid + "\n" + ts + "\n" + nonceStr + "\n" + pkg + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	jsapi = &JSAPIPayParams{
		AppId:     appid,
		TimeStamp: ts,
		NonceStr:  nonceStr,
		Package:   pkg,
		SignType:  SignTypeRSA,
		PaySign:   sign,
	}
	return jsapi, nil
}

// PaySignOfApp 获取 App sign
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_4.shtml
func (c *ClientV3) PaySignOfApp(appid, prepayid string) (app *AppPayParams, err error) {
	ts := util.Int642String(time.Now().Unix())
	nonceStr := util.RandomString(32)

	_str := appid + "\n" + ts + "\n" + nonceStr + "\n" + prepayid + "\n"
	sign, err := c.rsaSign(_str)
	if err != nil {
		return nil, err
	}

	app = &AppPayParams{
		Appid:     appid,
		Partnerid: c.Mchid,
		Prepayid:  prepayid,
		Package:   "Sign=WXPay",
		Noncestr:  nonceStr,
		Timestamp: ts,
		Sign:      sign,
	}
	return app, nil
}

// PaySignOfApplet 获取 小程序 paySign
// 文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_4.shtml
func (c *ClientV3) PaySignOfApplet(appid, prepayid string) (applet *AppletParams, err error) {
	jsapi, err := c.PaySignOfJSAPI(appid, prepayid)
	if err != nil {
		return nil, err
	}
	applet = &AppletParams{
		AppId:     jsapi.AppId,
		TimeStamp: jsapi.TimeStamp,
		NonceStr:  jsapi.NonceStr,
		Package:   jsapi.Package,
		SignType:  jsapi.SignType,
		PaySign:   jsapi.PaySign,
	}
	return applet, nil
}

// v3 鉴权请求Header
func (c *ClientV3) authorization(method, path string, bm pay.BodyMap) (string, error) {
	var (
		jb        = ""
		timestamp = time.Now().Unix()
		nonceStr  = util.RandomString(32)
	)
	if bm != nil {
		jb = bm.JsonBody()
	}
	if strings.HasSuffix(path, "?") {
		path = path[:len(path)-1]
	}
	ts := util.Int642String(timestamp)
	_str := method + "\n" + path + "\n" + ts + "\n" + nonceStr + "\n" + jb + "\n"
	if c.DebugSwitch == pay.DebugOn {
		xlog.Debugf("Wechat_V3_SignString:\n%s", _str)
	}
	sign, err := c.rsaSign(_str)
	if err != nil {
		return "", err
	}
	return Authorization + ` mchid="` + c.Mchid + `",nonce_str="` + nonceStr + `",timestamp="` + ts + `",serial_no="` + c.SerialNo + `",signature="` + sign + `"`, nil
}

func (c *ClientV3) rsaSign(str string) (string, error) {
	if c.privateKey == nil {
		return "", errors.New("privateKey can't be nil")
	}
	h := sha256.New()
	h.Write([]byte(str))
	result, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return util.NULL, fmt.Errorf("[%w]: %+v", pay.SignatureErr, err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// 自动同步请求验签
func (c *ClientV3) verifySyncSign(si *SignInfo) (err error) {
	if !c.autoSign {
		return nil
	}
	if si == nil {
		return errors.New("auto verify sign, but SignInfo is nil")
	}
	c.rwMu.RLock()
	wxPublicKey, exist := c.SnCertMap[si.HeaderSerial]
	c.rwMu.RUnlock()
	if !exist {
		err = c.AutoVerifySign(false)
		if err != nil {
			return fmt.Errorf("[get all public key err]: %v", err)
		}
		c.rwMu.RLock()
		wxPublicKey, exist = c.SnCertMap[si.HeaderSerial]
		c.rwMu.RUnlock()
		if !exist {
			return errors.New("auto verify sign, but public key not found")
		}
	}
	str := si.HeaderTimestamp + "\n" + si.HeaderNonce + "\n" + si.SignBody + "\n"
	signBytes, _ := base64.StdEncoding.DecodeString(si.HeaderSignature)
	h := sha256.New()
	h.Write([]byte(str))
	if err = rsa.VerifyPKCS1v15(wxPublicKey, crypto.SHA256, h.Sum(nil), signBytes); err != nil {
		return fmt.Errorf("[%w]: %v", pay.VerifySignatureErr, err)
	}
	return nil
}
