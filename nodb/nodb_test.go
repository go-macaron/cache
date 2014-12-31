// Copyright 2014 lunny
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package cache

// import (
// 	"strconv"
// 	"testing"
// 	"time"

// 	"github.com/macaron-contrib/cache"
// )

// func toInt(v string) (int64, error) {
// 	return strconv.ParseInt(v, 10, 64)
// }

// func TestNodbCache(t *testing.T) {
// 	bm, err := cache.NewCacher("nodb", cache.Options{AdapterConfig: "./cache"})
// 	if err != nil {
// 		t.Error("init err", err)
// 	}
// 	if err = bm.Put("lunny", 1, 10); err != nil {
// 		t.Error("set Error", err)
// 	}
// 	if !bm.IsExist("lunny") {
// 		t.Error("check err")
// 	}

// 	time.Sleep(11 * time.Second)

// 	if bm.IsExist("lunny") {
// 		t.Error("check err")
// 	}
// 	if err = bm.Put("lunny", 1, 10); err != nil {
// 		t.Error("set Error", err)
// 	}

// 	if v, _ := toInt(bm.Get("lunny").(string)); v != 1 {
// 		t.Error("get err")
// 	}

// 	if err = bm.Incr("lunny"); err != nil {
// 		t.Error("Incr Error", err)
// 	}

// 	if v, _ := toInt(bm.Get("lunny").(string)); v != 2 {
// 		t.Error("get err")
// 	}

// 	if err = bm.Decr("lunny"); err != nil {
// 		t.Error("Decr Error", err)
// 	}

// 	if v, _ := toInt(bm.Get("lunny").(string)); v != 1 {
// 		t.Error("get err")
// 	}

// 	bm.Delete("lunny")
// 	if bm.IsExist("lunny") {
// 		t.Error("delete err")
// 	}
// 	//test string
// 	if err = bm.Put("lunny", "author", 10); err != nil {
// 		t.Error("set Error", err)
// 	}
// 	if !bm.IsExist("lunny") {
// 		t.Error("check err")
// 	}

// 	if v := string(bm.Get("lunny").(string)); v != "author" {
// 		t.Error("get err")
// 	}
// 	// test clear all
// 	if err = bm.Flush(); err != nil {
// 		t.Error("clear all err")
// 	}
// }
