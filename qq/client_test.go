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

package qq

import (
	"context"
	"os"
	"testing"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/util"
	"github.com/rwscode/payutil/pkg/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	mchId  = "1368139502"
	apiKey = "GFDS8j98rewnmgl45wHTt980jg543abc"
)

func TestMain(m *testing.M) {

	// 初始化QQ客户端
	//    mchId：商户ID
	//    apiKey：API秘钥值
	client = NewClient(mchId, apiKey)

	// 打开Debug开关，输出日志
	client.DebugSwitch = pay.DebugOn

	// err := client.AddCertFilePath(nil, nil, nil)
	// if err != nil {
	//	panic(err)
	// }
	os.Exit(m.Run())
}

func TestClient_MicroPay(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("nonce_str", util.RandomString(32))

	qqRsp, err := client.MicroPay(ctx, bm)
	if err != nil {
		xlog.Errorf("client.Micropay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("qqRsp:", *qqRsp)
}

func TestNotifyResponse_ToXmlString(t *testing.T) {
	n := new(NotifyResponse)
	n.ReturnCode = "SUCCESS"
	xlog.Info(n.ToXmlString())

	n.ReturnCode = "FAIL"
	n.ReturnMsg = "abc"
	xlog.Info(n.ToXmlString())
}

func TestClient_DownloadRedListFile(t *testing.T) {
	bm := make(pay.BodyMap)
	bm.Set("date", 20160803)
	file, err := client.DownloadRedListFile(ctx, bm)
	if err != nil {
		xlog.Errorf("client.DownloadRedListFile(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("qqRsp:", file)
}
