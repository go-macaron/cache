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
	"fmt"
	"strconv"
)

// convert interface to string.
func GetString(v interface{}) string {
	switch result := v.(type) {
	case string:
		return result
	case []byte:
		return string(result)
	default:
		if v == nil {
			return ""
		} else {
			return fmt.Sprintf("%v", result)
		}
	}
}

// convert interface to int.
func GetInt(v interface{}) int {
	switch result := v.(type) {
	case int:
		return result
	case int32:
		return int(result)
	case int64:
		return int(result)
	default:
		d := GetString(v)
		if d != "" {
			value, err := strconv.Atoi(d)
			if err == nil {
				return value
			}
		}
	}
	return 0
}

// convert interface to int64.
func GetInt64(v interface{}) int64 {
	switch result := v.(type) {
	case int:
		return int64(result)
	case int32:
		return int64(result)
	case int64:
		return result
	default:
		d := GetString(v)
		if d != "" {
			result, err := strconv.ParseInt(d, 10, 64)
			if err == nil {
				return result
			}
		}
	}
	return 0
}

// convert interface to float64.
func GetFloat64(v interface{}) float64 {
	switch result := v.(type) {
	case float64:
		return result
	default:
		d := GetString(v)
		if d != "" {
			value, err := strconv.ParseFloat(d, 64)
			if err == nil {
				return value
			}
		}
	}
	return 0
}

// convert interface to bool.
func GetBool(v interface{}) bool {
	switch result := v.(type) {
	case bool:
		return result
	default:
		d := GetString(v)
		if d != "" {
			result, err := strconv.ParseBool(d)
			if err == nil {
				return result
			}
		}
	}
	return false
}

// convert interface to byte slice.
func getByteArray(v interface{}) []byte {
	switch result := v.(type) {
	case []byte:
		return result
	case string:
		return []byte(result)
	default:
		return nil
	}
}
