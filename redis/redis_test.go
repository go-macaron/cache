// Copyright 2013 Beego Authors
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
// 	"testing"
// 	"time"

// 	"github.com/beego/redigo/redis"

// 	"github.com/macaron-contrib/cache"
// )

// func TestRedisCache(t *testing.T) {
// 	bm, err := cache.NewCacher("redis", cache.Options{AdapterConfig: "127.0.0.1:6379"})
// 	if err != nil {
// 		t.Error("init err")
// 	}
// 	if err = bm.Put("astaxie", 1, 10); err != nil {
// 		t.Error("set Error", err)
// 	}
// 	if !bm.IsExist("astaxie") {
// 		t.Error("check err")
// 	}

// 	time.Sleep(10 * time.Second)

// 	if bm.IsExist("astaxie") {
// 		t.Error("check err")
// 	}
// 	if err = bm.Put("astaxie", 1, 10); err != nil {
// 		t.Error("set Error", err)
// 	}

// 	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
// 		t.Error("get err")
// 	}

// 	if err = bm.Incr("astaxie"); err != nil {
// 		t.Error("Incr Error", err)
// 	}

// 	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 2 {
// 		t.Error("get err")
// 	}

// 	if err = bm.Decr("astaxie"); err != nil {
// 		t.Error("Decr Error", err)
// 	}

// 	if v, _ := redis.Int(bm.Get("astaxie"), err); v != 1 {
// 		t.Error("get err")
// 	}
// 	bm.Delete("astaxie")
// 	if bm.IsExist("astaxie") {
// 		t.Error("delete err")
// 	}
// 	//test string
// 	if err = bm.Put("astaxie", "author", 10); err != nil {
// 		t.Error("set Error", err)
// 	}
// 	if !bm.IsExist("astaxie") {
// 		t.Error("check err")
// 	}

// 	if v, _ := redis.String(bm.Get("astaxie"), err); v != "author" {
// 		t.Error("get err")
// 	}
// 	// test clear all
// 	if err = bm.Flush(); err != nil {
// 		t.Error("clear all err")
// 	}
// }
