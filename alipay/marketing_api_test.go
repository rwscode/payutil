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
	"github.com/rwscode/payutil/pkg/xlog"
)

func TestOpenAppQrcodeCreate(t *testing.T) {
	// 请求参数
	bm := make(pay.BodyMap)
	bm.Set("url_param", "page/component/component-pages/view/view").
		Set("query_param", "x=1").
		Set("describe", "二维码描述")

	// 发起请求
	aliRsp, err := client.OpenAppQrcodeCreate(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}
