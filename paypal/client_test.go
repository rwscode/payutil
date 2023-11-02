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

package paypal

import (
	"context"
	"encoding/base64"
	"os"
	"testing"

	pay "github.com/rwscode/payutil"
	"github.com/rwscode/payutil/pkg/xlog"
)

var (
	client   *Client
	ctx      = context.Background()
	err      error
	Clientid = ""
	Secret   = ""
)

func TestMain(m *testing.M) {
	client, err = NewClient(Clientid, Secret, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = pay.DebugOff

	xlog.Debugf("Appid: %s", client.Appid)
	xlog.Debugf("AccessToken: %s", client.AccessToken)
	xlog.Debugf("ExpiresIn: %d", client.ExpiresIn)
	os.Exit(m.Run())
}

func TestBasicAuth(t *testing.T) {
	uname := "jerry"
	passwd := "12346"
	auth := base64.StdEncoding.EncodeToString([]byte(uname + ":" + passwd))
	xlog.Debugf("Basic %s", auth)
}
