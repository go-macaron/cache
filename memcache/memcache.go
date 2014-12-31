// Copyright 2013 Beego Authors
// Copyright 2014 Unknwon
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
	"errors"

	"github.com/beego/memcache"

	"github.com/macaron-contrib/cache"
)

// Memcache adapter.
type MemcacheCache struct {
	c        *memcache.Connection
	conninfo string
}

// create new memcache adapter.
func NewMemCache() *MemcacheCache {
	return &MemcacheCache{}
}

// get value from memcache.
func (rc *MemcacheCache) Get(key string) interface{} {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	v, err := rc.c.Get(key)
	if err != nil {
		return nil
	}
	var contain interface{}
	if len(v) > 0 {
		contain = string(v[0].Value)
	} else {
		contain = nil
	}
	return contain
}

// put value to memcache. only support string.
func (rc *MemcacheCache) Put(key string, val interface{}, timeout int64) error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	v, ok := val.(string)
	if !ok {
		return errors.New("val must string")
	}
	stored, err := rc.c.Set(key, 0, uint64(timeout), []byte(v))
	if err == nil && stored == false {
		return errors.New("stored fail")
	}
	return err
}

// delete value in memcache.
func (rc *MemcacheCache) Delete(key string) error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	_, err := rc.c.Delete(key)
	return err
}

// [Not Support]
// increase counter.
func (rc *MemcacheCache) Incr(key string) error {
	return errors.New("not support in memcache")
}

// [Not Support]
// decrease counter.
func (rc *MemcacheCache) Decr(key string) error {
	return errors.New("not support in memcache")
}

// check value exists in memcache.
func (rc *MemcacheCache) IsExist(key string) bool {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return false
		}
	}
	v, err := rc.c.Get(key)
	if err != nil {
		return false
	}
	if len(v) == 0 {
		return false
	} else {
		return true
	}
}

// clear all cached in memcache.
func (rc *MemcacheCache) Flush() error {
	if rc.c == nil {
		var err error
		rc.c, err = rc.connectInit()
		if err != nil {
			return err
		}
	}
	err := rc.c.FlushAll()
	return err
}

// start memcache adapter.
// config string is like {"conn":"connection info"}.
// if connecting error, return.
func (rc *MemcacheCache) StartAndGC(opt cache.Options) error {
	rc.conninfo = opt.AdapterConfig
	var err error
	if rc.c != nil {
		rc.c, err = rc.connectInit()
		if err != nil {
			return errors.New("dial tcp conn error")
		}
	}
	return nil
}

// connect to memcache and keep the connection.
func (rc *MemcacheCache) connectInit() (*memcache.Connection, error) {
	c, err := memcache.Connect(rc.conninfo)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func init() {
	cache.Register("memcache", NewMemCache())
}
