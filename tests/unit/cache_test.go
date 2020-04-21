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

package unit

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta/cache"

	"github.com/stretchr/testify/assert"
)

func Test_cache_key_can_be_created_from_request_object(t *testing.T) {
	var buff io.ReadWriter
	request, _ := http.NewRequest("GET", "https://example.com/sample/cache-key/test+test@test."+
		"com?with=a&query=string",
		buff)

	cacheKey := cache.CreateCacheKey(request)

	assert.Equal(t, "https://example.com/sample/cache-key/test+test@test.com?with=a&query=string", cacheKey,
		"The cache key was not created correctly.")
}

func Test_an_item_can_be_stored_in_cache(t *testing.T) {
	var buff io.ReadWriter
	url := "https://example.com/sample/cache-key/"
	request, _ := http.NewRequest("GET", url, buff)

	cacheKey := cache.CreateCacheKey(request)

	myCache := cache.NewGoCache(30, 30)

	found := myCache.Has(cacheKey)
	assert.False(t, found, "item already existed in cache")

	toCache := "test Item"
	record := httptest.NewRecorder()
	record.WriteString(toCache)
	result := record.Result()

	myCache.Set(cacheKey, result)

	found = myCache.Has(cacheKey)
	assert.True(t, found, "item does not exist in cache")

	pulledFromCache := myCache.Get(cacheKey)

	assert.NotEqual(t, result, pulledFromCache, "Item pulled from cache was not a copy")
	cachedBody, _ := ioutil.ReadAll(pulledFromCache.Body)
	assert.Equal(t, toCache, string(cachedBody), "Item pulled from cache was not correct")

}

func Test_an_item_can_be_deleted_from_cache(t *testing.T) {
	var buff io.ReadWriter
	url := "https://example.com/sample/cache-key/delete"
	request, _ := http.NewRequest("GET", url, buff)

	cacheKey := cache.CreateCacheKey(request)

	myCache := cache.NewGoCache(30, 30)

	record := httptest.NewRecorder()
	record.WriteString("test Item")
	result := record.Result()

	myCache.Set(cacheKey, result)

	found := myCache.Has(cacheKey)
	assert.True(t, found, "item does not exist in cache")

	myCache.Delete(cacheKey)

	found = myCache.Has(cacheKey)
	assert.False(t, found, "item was not deleted from cache")

}

func Test_cache_can_be_cleared(t *testing.T) {
	var buff io.ReadWriter
	url := "https://example.com/sample/cache-key/clear"
	request, _ := http.NewRequest("GET", url, buff)

	cacheKey := cache.CreateCacheKey(request)

	myCache := cache.NewGoCache(30, 30)

	record := httptest.NewRecorder()
	record.WriteString("test Item")
	result := record.Result()

	myCache.Set(cacheKey, result)

	found := myCache.Has(cacheKey)
	assert.True(t, found, "item does not exist in cache")

	myCache.Clear()

	found = myCache.Has(cacheKey)
	assert.False(t, found, "cache was not cleared")
}
