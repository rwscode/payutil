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

package xpem

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func DecodePublicKey(pemContent []byte) (publicKey *rsa.PublicKey, err error) {
	block, _ := pem.Decode(pemContent)
	if block == nil {
		return nil, fmt.Errorf("pem.Decode(%s)：pemContent decode error", pemContent)
	}
	switch block.Type {
	case "CERTIFICATE":
		pubKeyCert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("x509.ParseCertificate(%s)：%w", pemContent, err)
		}
		pubKey, ok := pubKeyCert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("公钥证书提取公钥出错 [%s]", pemContent)
		}
		publicKey = pubKey
	case "PUBLIC KEY":
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("x509.ParsePKIXPublicKey(%s),err:%w", pemContent, err)
		}
		pubKey, ok := pub.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("公钥解析出错 [%s]", pemContent)
		}
		publicKey = pubKey
	case "RSA PUBLIC KEY":
		pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("x509.ParsePKCS1PublicKey(%s)：%w", pemContent, err)
		}
		publicKey = pubKey
	}
	return publicKey, nil
}

func DecodePrivateKey(pemContent []byte) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(pemContent)
	if block == nil {
		return nil, fmt.Errorf("pem.Decode(%s)：pemContent decode error", pemContent)
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("私钥解析出错 [%s]", pemContent)
		}
		var ok bool
		privateKey, ok = pk8.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("私钥解析出错 [%s]", pemContent)
		}
	}
	return privateKey, nil
}
