/*
 * Copyright 2018 - Present Okta, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cache

import (
	"net/http"
)

type MapCache struct {
	ttl   int32
	tti   int32
	cache map[string]*http.Response
}

func NewMapCache(ttl int32, tti int32) MapCache {
	return MapCache{
		ttl:   ttl,
		tti:   tti,
		cache: make(map[string]*http.Response),
	}
}

func (c MapCache) Get(key string) *http.Response {
	return c.cache[key]

}

func (c MapCache) Set(key string, value *http.Response) {
	c.cache[key] = value
}

func (c MapCache) Delete(key string) {
	delete(c.cache, key)
}

func (c MapCache) Clear() {
	for k := range c.cache {
		delete(c.cache, k)
	}
}

func (c MapCache) Has(key string) bool {
	_, found := c.cache[key]

	return found
}
