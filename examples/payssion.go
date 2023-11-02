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

package main

import (
	"github.com/rwscode/payutil/payssion"
	"log"
	"net/url"
	"os"
)

func Test() {
	c := payssion.NewClient("", "")
	c.SetLive(true)
	c.Debug(os.Stdout)
	data := url.Values{}
	data.Set("pm_id", "alipay_cn") // alipay_cn  , tenpay_cn
	data.Set("amount", "1")
	data.Set("currency", "CNY")
	data.Set("description", "")
	data.Set("order_id", "test00000002")
	data.Set("return_url", "https://www.baidu.com")
	rsp, err := c.Create(data)
	if err != nil {
		log.Println(err)
	}
	log.Println(rsp)
}
