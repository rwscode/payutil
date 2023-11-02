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
	"testing"

	"github.com/rwscode/payutil/pkg/xlog"
)

func TestPaySignOfJSAPIp(t *testing.T) {
	jsapi, err := client.PaySignOfJSAPI("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("jsapi:%#v", jsapi)
}

func TestPaySignOfApp(t *testing.T) {
	app, err := client.PaySignOfApp("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("app:%#v", app)
}

func TestPaySignOfApplet(t *testing.T) {
	applet, err := client.PaySignOfApplet("appid", "prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("applet:%#v", applet)
}
