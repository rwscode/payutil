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

package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/rwscode/payutil/pkg/util"
)

// AES-GCM 加密数据
func GCMEncrypt(originText, additional, key []byte) (nonce []byte, cipherText []byte, err error) {
	return gcmEncrypt(originText, additional, key)
}

// AES-GCM 解密数据
func GCMDecrypt(cipherText, nonce, additional, key []byte) ([]byte, error) {
	return gcmDecrypt(cipherText, nonce, additional, key)
}

func gcmDecrypt(secretData, nonce, additional, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("cipher.NewGCM(),error:%w", err)
	}
	originByte, err := gcm.Open(nil, nonce, secretData, additional)
	if err != nil {
		return nil, err
	}
	return originByte, nil
}

func gcmEncrypt(originText, additional, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	nonce := []byte(util.RandomString(12))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, fmt.Errorf("cipher.NewGCM(),error:%w", err)
	}
	cipherBytes := gcm.Seal(nil, nonce, originText, additional)
	return nonce, cipherBytes, nil
}
