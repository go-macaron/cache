// Copyright 2013 Beego Authors
// Copyright 2014 Unknown
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
	"fmt"

	"github.com/Unknwon/macaron"
)

// Cache interface contains all behaviors for cache adapter.
// usage:
//	cache.Register("file",cache.NewFileCache()) // this operation is run in init method of file.go.
//	c := cache.NewCache("file","{....}")
//	c.Put("key",value,3600)
//	v := c.Get("key")
//
//	c.Incr("counter")  // now is 1
//	c.Incr("counter")  // now is 2
//	count := c.Get("counter").(int)
type Cache interface {
	// get cached value by key.
	Get(key string) interface{}
	// set cached value with key and expire time.
	Put(key string, val interface{}, timeout int64) error
	// delete cached value by key.
	Delete(key string) error
	// increase cached int value by key, as a counter.
	Incr(key string) error
	// decrease cached int value by key, as a counter.
	Decr(key string) error
	// check if cached value exists or not.
	IsExist(key string) bool
	// clear all cache.
	ClearAll() error
	// start gc routine based on config string settings.
	StartAndGC(config string) error
}

type CacheOptions struct {
	// Name of adapter. Default is "memory".
	Adapter string
	// GC interval for memory adapter. Default is 60.
	Interval int
	// Connection string for non-memory adapter.
	Conn string
}

func prepareOptions(options []CacheOptions) CacheOptions {
	var opt CacheOptions
	if len(options) > 0 {
		opt = options[0]
	}

	// Defaults
	if len(opt.Adapter) == 0 {
		opt.Adapter = "memory"
	}

	if opt.Adapter == "memory" {
		if opt.Interval == 0 {
			opt.Interval = 60
		}
	} else {
		if len(opt.Conn) == 0 {
			panic("no connection string is given for non-memory cache adapter")
		}
	}

	return opt
}

// Cacher is a middleware that maps a cache.Cache service into the Macaron handler chain.
// An single variadic cache.CacheOptions struct can be optionally provided to configure.
func Cacher(options ...CacheOptions) macaron.Handler {
	opt := prepareOptions(options)
	return func(ctx *macaron.Context) {
		c, err := NewCache(opt.Adapter,
			fmt.Sprintf(`{"interval":%d,"conn":"%s"}`, opt.Interval, opt.Conn))
		if err != nil {
			panic(err)
		}
		ctx.Map(c)
	}
}

var adapters = make(map[string]Cache)

// Register makes a cache adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, adapter Cache) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// Create a new cache driver by adapter name and config string.
// config need to be correct JSON as string: {"interval":360}.
// it will start gc automatically.
func NewCache(adapterName, config string) (Cache, error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		return nil, fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
	}
	err := adapter.StartAndGC(config)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}
