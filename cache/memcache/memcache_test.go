// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package memcache

import (
	_ "github.com/bradfitz/gomemcache/memcache"

	"github.com/titan-group/beego/cache"
	"strconv"
	"testing"
	"time"
)

func TestRedisCache(t *testing.T) {
	bm, err := cache.NewCache("memcache", `{"conn": "127.0.0.1:11211"}`)
	if err != nil {
		t.Error("init err")
	}
	if err = bm.Put("titan-group", "1", 10); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("titan-group") {
		t.Error("check err")
	}

	time.Sleep(10 * time.Second)

	if bm.IsExist("titan-group") {
		t.Error("check err")
	}
	if err = bm.Put("titan-group", "1", 10); err != nil {
		t.Error("set Error", err)
	}

	if v, err := strconv.Atoi(bm.Get("titan-group").(string)); err != nil || v != 1 {
		t.Error("get err")
	}

	if err = bm.Incr("titan-group"); err != nil {
		t.Error("Incr Error", err)
	}

	if v, err := strconv.Atoi(bm.Get("titan-group").(string)); err != nil || v != 2 {
		t.Error("get err")
	}

	if err = bm.Decr("titan-group"); err != nil {
		t.Error("Decr Error", err)
	}

	if v, err := strconv.Atoi(bm.Get("titan-group").(string)); err != nil || v != 1 {
		t.Error("get err")
	}
	bm.Delete("titan-group")
	if bm.IsExist("titan-group") {
		t.Error("delete err")
	}

	//test string
	if err = bm.Put("titan-group", "author", 10); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("titan-group") {
		t.Error("check err")
	}

	if v := bm.Get("titan-group").(string); v != "author" {
		t.Error("get err")
	}

	//test GetMulti
	if err = bm.Put("titan-group1", "author1", 10); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("titan-group1") {
		t.Error("check err")
	}

	vv := bm.GetMulti([]string{"titan-group", "titan-group1"})
	if len(vv) != 2 {
		t.Error("GetMulti ERROR")
	}
	if vv[0].(string) != "author" {
		t.Error("GetMulti ERROR")
	}
	if vv[1].(string) != "author1" {
		t.Error("GetMulti ERROR")
	}

	// test clear all
	if err = bm.ClearAll(); err != nil {
		t.Error("clear all err")
	}
}
