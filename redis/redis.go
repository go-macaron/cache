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

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/beego/redigo/redis"

	"github.com/macaron-contrib/cache"
)

var (
	// the collection name of redis for cache adapter.
	DefaultKey string = "MacaronRedis"
)

// Redis cache adapter.
type RedisCache struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	key      string
}

// create new redis cache with default collection name.
func NewRedisCache() *RedisCache {
	return &RedisCache{key: DefaultKey}
}

// actually do the redis cmds
func (rc *RedisCache) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

// Get cache from redis.
func (rc *RedisCache) Get(key string) interface{} {
	v, err := rc.do("GET", key)
	if err != nil {
		return nil
	}

	return v
}

// put cache to redis.
func (rc *RedisCache) Put(key string, val interface{}, timeout int64) error {
	_, err := rc.do("SETEX", key, timeout, val)
	if err != nil {
		return nil
	}
	_, err = rc.do("HSET", rc.key, key, true)
	return err
}

// delete cache in redis.
func (rc *RedisCache) Delete(key string) error {
	_, err := rc.do("DEL", key)
	if err != nil {
		return nil
	}
	_, err = rc.do("HDEL", rc.key, key)
	return err
}

// check cache's existence in redis.
func (rc *RedisCache) IsExist(key string) bool {
	v, err := redis.Bool(rc.do("EXISTS", key))
	if err != nil {
		return false
	}
	if v == false {
		_, err := rc.do("HDEL", rc.key, key)
		if err != nil {
			return false
		}
	}
	return v
}

// increase counter in redis.
func (rc *RedisCache) Incr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, 1))
	return err
}

// decrease counter in redis.
func (rc *RedisCache) Decr(key string) error {
	_, err := redis.Bool(rc.do("INCRBY", key, -1))
	return err
}

// clean all cache in redis. delete this redis collection.
func (rc *RedisCache) ClearAll() error {
	cachedKeys, err := redis.Strings(rc.do("HKEYS", rc.key))
	for _, str := range cachedKeys {
		_, err := rc.do("DEL", str)
		if err != nil {
			return nil
		}
	}
	_, err = rc.do("DEL", rc.key)
	return err
}

// start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *RedisCache) StartAndGC(config string) error {
	var cf map[string]interface{}
	if err := json.Unmarshal([]byte(config), &cf); err != nil {
		return err
	}

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}

	if _, ok := cf["conn"]; !ok {
		return errors.New("redis: config has no conn key")
	}

	rc.key = cf["key"].(string)
	rc.conninfo = cf["conn"].(string)
	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()
	if err := c.Err(); err != nil {
		return err
	}

	return nil
}

// connect to redis.
func (rc *RedisCache) connectInit() {
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", rc.conninfo)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	cache.Register("redis", NewRedisCache())
}
