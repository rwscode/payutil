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

// RSA加密数据
//
//	t：PKCS1 或 PKCS8
//	originData：原始字符串byte数组
//	publicKey：公钥
func RsaEncryptData(t PKCSType, originData []byte, publicKey string) (cipherData []byte, err error) {
	var (
		key *rsa.PublicKey
	)

	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("publicKey decode error")
	}

	switch t {
	case PKCS1:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	case PKCS8:
		pkcs8Key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pk8, ok := pkcs8Key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
		key = pk8
	default:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	}

	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, key, originData)
	if err != nil {
		return nil, fmt.Errorf("xrsa.EncryptPKCS1v15：%w", err)
	}
	return cipherBytes, nil
}

// RSA加密数据
//
//	OAEPWithSHA-256AndMGF1Padding
func RsaEncryptOAEPData(h hash.Hash, t PKCSType, publicKey string, originData, label []byte) (cipherData []byte, err error) {
	var (
		key *rsa.PublicKey
	)
	if len(originData) > 190 {
		return nil, errors.New("message too long for RSA public key size")
	}
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("publicKey decode error")
	}

	switch t {
	case PKCS1:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	case PKCS8:
		pkcs8Key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		pk8, ok := pkcs8Key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("parse PKCS8 key error")
		}
		key = pk8
	default:
		pkcs1Key, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		key = pkcs1Key
	}

	cipherBytes, err := rsa.EncryptOAEP(h, rand.Reader, key, originData, label)
	if err != nil {
		return nil, err
	}
	return cipherBytes, nil
}
