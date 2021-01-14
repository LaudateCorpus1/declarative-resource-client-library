// Copyright 2021 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package dcl

import (
	"encoding/json"
	"fmt"

	glog "github.com/golang/glog"
)

// MoveMapEntry moves the entry at `from` to `to`.  `from` and `to` are slices
// of string keys.  Each key except the last must refer to a map[string]interface{}
// in m - we will descend into m following those keys.  If the maps at the levels
// above the target are empty after the move, they will be deleted.  If there
// are no maps along the path to `to`, they will be created.  If a map above
// the level of the target is missing, nothing will be done.  If the map exists
// but `target` is not present, `nil` will be inserted at `to`.
func MoveMapEntry(m map[string]interface{}, from, to []string) error {
	fetch := m
	// All elements before the last must point to a map[string]interface{} -
	// this ranges over all those elements, so at the end of this loop, we have
	// the map which contains the actual final element to move.
	for _, idx := range from[:len(from)-1] {
		f, ok := fetch[idx]
		if !ok {
			// Nothing to move, so it's not an error not to move it.
			return nil
		}
		fetch, ok = f.(map[string]interface{})
		if !ok {
			return fmt.Errorf("could not fetch %q from %v", idx, fetch)
		}
	}
	value, ok := fetch[from[len(from)-1]]
	if !ok {
		value = nil
	}
	delete(fetch, from[len(from)-1])
	if len(to) > 0 {
		fetch = m
		for _, idx := range to[:len(to)-1] {
			f, ok := fetch[idx]
			if !ok {
				fetch[idx] = make(map[string]interface{})
				f = fetch[idx]
			}
			fetch, ok = f.(map[string]interface{})
			if !ok {
				return fmt.Errorf("%v is not map[string]interface{}", f)
			}
		}
		fetch[to[len(to)-1]] = value
	}
	return deleteIfEmpty(m, from)
}

// GetMapEntry returns the value at `path` from `m`, following the same rules as
// `MoveMapEntry` except that a missing map or value is an error.
func GetMapEntry(m map[string]interface{}, path []string) (interface{}, error) {
	if len(path) == 0 {
		return m, nil
	}
	fetch := m
	// All elements before the last must point to a map[string]interface{} -
	// this ranges over all those elements, so at the end of this loop, we have
	// the map which contains the element to fetch.
	for _, idx := range path[:len(path)-1] {
		f, ok := fetch[idx]
		if !ok {
			return nil, fmt.Errorf("could not find %q in %v", idx, fetch)
		}
		fetch, ok = f.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not fetch %q from %v", idx, fetch)
		}
	}
	value, ok := fetch[path[len(path)-1]]
	if !ok {
		return nil, fmt.Errorf("could not fetch %q from %v", path[len(path)-1], fetch)
	}
	return value, nil
}

func deleteIfEmpty(m map[string]interface{}, from []string) error {
	if len(from) > 1 {
		sub, ok := m[from[0]]
		if !ok {
			return fmt.Errorf("could not fetch %q from %v", from[0], m)
		}
		smap, ok := sub.(map[string]interface{})
		if !ok {
			glog.Warningf("In deleting empty map while marshalling, %v not map[string]interface{}", sub)
			return nil
		}
		deleteIfEmpty(smap, from[1:])
	}
	if len(from) >= 1 {
		if sub, ok := m[from[0]]; ok {
			if subm, ok := sub.(map[string]interface{}); ok && len(subm) == 0 {
				delete(m, from[0])
			}
		}
	}
	return nil
}

// PutMapEntry inserts `item` at `path` into `m` - the inverse of GetMapEntry.
func PutMapEntry(m map[string]interface{}, path []string, item interface{}) error {
	if len(path) == 0 {
		return fmt.Errorf("cannot insert value at empty path")
	}
	put := m
	// All elements before the last must point to a map[string]interface{} -
	// this ranges over all those elements, so at the end of this loop, we have
	// the map which contains the element to fetch.
	for _, idx := range path[:len(path)-1] {
		f, ok := put[idx]
		if !ok {
			f = make(map[string]interface{})
			put[idx] = f
		}
		put, ok = f.(map[string]interface{})
		if !ok {
			return fmt.Errorf("could not cast %q from %v as map[string]interface{}", idx, put)
		}
	}
	put[path[len(path)-1]] = item
	return nil
}

// MapFromListOfKeyValues turns a [{"key": k, "value": v}, ...] format-map into a normal string-string map.
// This is useful for a handful of GCP APIs which have chosen to represent maps this way.  We
// expect relatively few of these in newer APIs - it is explicitly against https://aip.dev/apps/2717 -
// ("such a map is represented by a normal JSON object").
// That AIP didn't exist at the time of development of, for instance, Compute v1.
func MapFromListOfKeyValues(rawFetch map[string]interface{}, path []string) (map[string]string, error) {
	i, err := GetMapEntry(rawFetch, path)
	if err != nil {
		// If there's nothing there, it's okay to ignore.
		glog.Warningf("In converting a map to [{\"key\": k, ...}, ...] format, no entry at %q in %v", path, rawFetch)
		return nil, nil
	}
	il, ok := i.([]interface{})
	if !ok {
		return nil, fmt.Errorf("could not cast %v to []interface{}", i)
	}
	var items []map[string]interface{}
	for _, it := range il {
		cast, ok := it.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("could not cast %v to map[string]interface{}", it)
		}
		items = append(items, cast)
	}

	m := make(map[string]string, len(items))
	for _, item := range items {
		key, ok := item["key"].(string)
		if !ok {
			return nil, fmt.Errorf("could not find 'key' in %v", item)
		}
		value, ok := item["value"].(string)
		if !ok {
			return nil, fmt.Errorf("could not find 'value' in %v", item)
		}
		m[key] = value
	}
	return m, nil
}

// ListOfKeyValuesFromMap is the opposite of MapFromListOfKeyValues, used in marshalling instead of unmarshalling.
func ListOfKeyValuesFromMap(m map[string]string) ([]map[string]string, error) {
	var items []map[string]string
	for k, v := range m {
		items = append(items, map[string]string{
			"key":   k,
			"value": v,
		})
	}
	return items, nil
}

// ListOfKeyValuesFromMapInItemsStruct returns the opposite of MapFromListOfKeyValues, except nested inside an struct under the key "items".
func ListOfKeyValuesFromMapInItemsStruct(m map[string]string) (map[string][]map[string]string, error) {
	maps, err := ListOfKeyValuesFromMap(m)
	if err != nil {
		return nil, err
	}
	return map[string][]map[string]string{
		"items": maps,
	}, nil
}

// ConvertToMap converts the specified object into the map[string]interface{} which can
// be serialized into the same json object as the input object.
func ConvertToMap(obj interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

// ValueOrEmptyString takes in the pointer to a string and returns either the empty string or its value.
func ValueOrEmptyString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// ValueOrEmptyInt64 returns the value or the default value if the pointer is nil.
func ValueOrEmptyInt64(s *int64) int64 {
	if s == nil {
		return 0
	}
	return *s
}

// ValueOrEmptyBool returns the value or the default value if the pointer is nil.
func ValueOrEmptyBool(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}

// ValueOrEmptyDouble returns the value or the default value if the pointer is nil.
func ValueOrEmptyDouble(s *float64) float64 {
	if s == nil {
		return 0.0
	}
	return *s
}

// FindStringInArray returns true if value found in array of strings
func FindStringInArray(s string, items []string) bool {
	for _, v := range items {
		if v == s {
			return true
		}
	}
	return false
}
