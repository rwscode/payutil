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

package errgroup

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/rwscode/payutil/pkg/xlog"
)

func TestErrgroup(t *testing.T) {
	var count int64 = 1
	countBackup := count
	eg := WithContext(context.Background())

	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	// go 协程
	eg.Go(func(ctx context.Context) error {
		atomic.AddInt64(&count, 1)
		return errors.New("error ,reset count")
	})
	// wait 协程 Done
	if err := eg.Wait(); err != nil {
		// do some thing
		count = countBackup
		xlog.Error(err)
		// return
	}
	xlog.Debug(count)
}

func TestErrgroup1(t *testing.T) {
	var (
		count int64 = 1
		eg    Group
		goNum = 3 // every times run goNum goroutine
	)
	for i := 0; i < 10; i++ {
		eg.Go(func(ctx context.Context) error {
			atomic.AddInt64(&count, 1)
			xlog.Debug("count:", count)
			return nil
		})
		if eg.WorkNum() == goNum {
			if err := eg.Wait(); err != nil {
				xlog.Error(err)
				// to do something you need
			}
			xlog.Info("wat")
			time.Sleep(time.Second)
		}
	}
}
