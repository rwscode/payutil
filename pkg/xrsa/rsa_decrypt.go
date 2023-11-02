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

package xrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
)

// RSA解密数据
//
//	t：PKCS1 或 PKCS8
//	cipherData：加密字符串byte数组
//	privateKey：私钥
func RsaDecryptData(t PKCSType, cipherData []byte, privateKey string) (originData []byte, err error) {
	var (
		key *rsa.PrivateKey
	)

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("privateKey decode error")
	}

	switch t {
	case PKCS1:
		if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	case PKCS8:
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pk8, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
		key = pk8
	default:
		if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	originBytes, err := rsa.DecryptPKCS1v15(rand.Reader, key, cipherData)
	if err != nil {
		return nil, fmt.Errorf("xrsa.DecryptPKCS1v15：%w", err)
	}
	return originBytes, nil
}

// RSA解密数据
//
//	OAEPWithSHA-256AndMGF1Padding
func RsaDecryptOAEPData(h hash.Hash, t PKCSType, privateKey string, ciphertext, label []byte) (originData []byte, err error) {
	var (
		key *rsa.PrivateKey
	)

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("privateKey decode error")
	}

	switch t {
	case PKCS1:
		if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	case PKCS8:
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pk8, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
		key = pk8
	default:
		if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return nil, err
		}
	}

	originBytes, err := rsa.DecryptOAEP(h, rand.Reader, key, ciphertext, label)
	if err != nil {
		return nil, err
	}
	return originBytes, nil
}
