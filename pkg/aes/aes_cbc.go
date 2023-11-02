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
	"errors"
)

// AES-CBC 加密数据
func CBCEncrypt(originData, key, iv []byte) ([]byte, error) {
	return cbcEncrypt(originData, key, iv)
}

// AES-CBC 解密数据
func CBCDecrypt(secretData, key, iv []byte) ([]byte, error) {
	return cbcDecrypt(secretData, key, iv)
}

func cbcEncrypt(originData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originData = PKCS7Padding(originData, block.BlockSize())
	secretData := make([]byte, len(originData))
	blockMode := cipher.NewCBCEncrypter(block, iv[:block.BlockSize()])
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}

func cbcDecrypt(secretData, key, iv []byte) (originByte []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originByte = make([]byte, len(secretData))
	blockMode := cipher.NewCBCDecrypter(block, iv[:block.BlockSize()])
	blockMode.CryptBlocks(originByte, secretData)
	if len(originByte) == 0 {
		return nil, errors.New("blockMode.CryptBlocks error")
	}
	return PKCS7UnPadding(originByte), nil
}
