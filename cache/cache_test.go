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

package cache

import (
	"os"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	bm, err := NewCache("memory", `{"interval":20}`)
	if err != nil {
		t.Error("init err")
	}
	if err = bm.Put("titan-group", 1, 10); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("titan-group") {
		t.Error("check err")
	}

	if v := bm.Get("titan-group"); v.(int) != 1 {
		t.Error("get err")
	}

	time.Sleep(30 * time.Second)

	if bm.IsExist("titan-group") {
		t.Error("check err")
	}

	if err = bm.Put("titan-group", 1, 10); err != nil {
		t.Error("set Error", err)
	}

	if err = bm.Incr("titan-group"); err != nil {
		t.Error("Incr Error", err)
	}

	if v := bm.Get("titan-group"); v.(int) != 2 {
		t.Error("get err")
	}

	if err = bm.Decr("titan-group"); err != nil {
		t.Error("Decr Error", err)
	}

	if v := bm.Get("titan-group"); v.(int) != 1 {
		t.Error("get err")
	}
	bm.Delete("titan-group")
	if bm.IsExist("titan-group") {
		t.Error("delete err")
	}

	//test GetMulti
	if err = bm.Put("titan-group", "author", 10); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("titan-group") {
		t.Error("check err")
	}
	if v := bm.Get("titan-group"); v.(string) != "author" {
		t.Error("get err")
	}

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
}

func TestFileCache(t *testing.T) {
	bm, err := NewCache("file", `{"CachePath":"cache","FileSuffix":".bin","DirectoryLevel":2,"EmbedExpiry":0}`)
	if err != nil {
		t.Error("init err")
	}
	if err = bm.Put("titan-group", 1, 10); err != nil {
		t.Error("set Error", err)
	}
	if !bm.IsExist("titan-group") {
		t.Error("check err")
	}

	if v := bm.Get("titan-group"); v.(int) != 1 {
		t.Error("get err")
	}

	if err = bm.Incr("titan-group"); err != nil {
		t.Error("Incr Error", err)
	}

	if v := bm.Get("titan-group"); v.(int) != 2 {
		t.Error("get err")
	}

	if err = bm.Decr("titan-group"); err != nil {
		t.Error("Decr Error", err)
	}

	if v := bm.Get("titan-group"); v.(int) != 1 {
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
	if v := bm.Get("titan-group"); v.(string) != "author" {
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

	os.RemoveAll("cache")
}
