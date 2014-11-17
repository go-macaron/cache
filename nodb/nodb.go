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

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/lunny/nodb"
	"github.com/lunny/nodb/config"

	"github.com/macaron-contrib/cache"
)

var (
	ErrDBExists = errors.New("db is exsit")
)

// Memcache adapter.
type NodbCache struct {
	dbs      *nodb.Nodb
	db       *nodb.DB
	filepath string
}

// create new nodb adapter.
func NewNodbCache() *NodbCache {
	return &NodbCache{}
}

// get value from nodb.
func (rc *NodbCache) Get(key string) interface{} {
	v, err := rc.db.Get([]byte(key))
	if err != nil {
		return nil
	}
	var contain interface{}
	if len(v) > 0 {
		contain = string(v)
	} else {
		contain = nil
	}
	return contain
}

// put value to nodb. only support string.
func (rc *NodbCache) Put(key string, val interface{}, timeout int64) error {
	var content []byte
	switch val.(type) {
	case string:
		content = []byte(val.(string))
	case []byte:
		content = val.([]byte)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		content = []byte(fmt.Sprintf("%v", val))
	default:
		return errors.New("val must string")
	}

	err := rc.db.Set([]byte(key), content)
	if err != nil {
		return err
	}

	_, err = rc.db.Expire([]byte(key), timeout)
	return err
}

// delete value in memcache.
func (rc *NodbCache) Delete(key string) error {
	_, err := rc.db.Del([]byte(key))
	return err
}

// increase counter.
func (rc *NodbCache) Incr(key string) error {
	_, err := rc.db.Incr([]byte(key))
	return err
}

// decrease counter.
func (rc *NodbCache) Decr(key string) error {
	_, err := rc.db.Decr([]byte(key))
	return err
}

// check value exists in memcache.
func (rc *NodbCache) IsExist(key string) bool {
	v, err := rc.db.Exists([]byte(key))
	if err != nil || v == 0 {
		return false
	}

	return true
}

// clear all cached in nodb.
func (rc *NodbCache) ClearAll() error {
	os.RemoveAll(rc.filepath)

	rc.dbs.Close()
	rc.db = nil
	rc.dbs = nil

	return rc.new()
}

func (rc *NodbCache) new() error {
	var err error
	if rc.db == nil {
		cfg := new(config.Config)
		cfg.DataDir = rc.filepath
		rc.dbs, err = nodb.Open(cfg)
		if err != nil {
			return err
		}

		rc.db, err = rc.dbs.Select(0)
		return err
	}
	return ErrDBExists
}

// start nodbcache adapter.
// config string is like {"path":"./cur.db", "interval":10}.//seconds
func (rc *NodbCache) StartAndGC(config string) error {
	var cf map[string]interface{}
	if err := json.Unmarshal([]byte(config), &cf); err != nil {
		return err
	}

	if _, ok := cf["path"]; !ok {
		return errors.New("nodbcache: config has no path key")
	}
	rc.filepath = cf["path"].(string)

	return rc.new()
}

func init() {
	cache.Register("nodb", NewNodbCache())
}
