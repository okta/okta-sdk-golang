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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/okta/okta-sdk-golang/okta/cache"

	"github.com/stretchr/testify/assert"
)

func Test_cache_key_can_be_created_from_request_object(t *testing.T) {
	var buff io.ReadWriter
	request, _ := http.NewRequest("GET", "https://okta.com/sample/cache-key/test+test@test.com?with=a&query=string", buff)

	cacheKey := cache.CreateCacheKey(request)

	assert.Equal(t, "okta_com_sample_cache_key_test_test_test_com", cacheKey, "The cache key was not created correctly.")
}

func Test_an_item_can_be_stored_in_cache(t *testing.T) {
	var buff io.ReadWriter
	url := "https://okta.com/sample/cache-key/"
	request, _ := http.NewRequest("GET", url, buff)

	cacheKey := cache.CreateCacheKey(request)

	myCache := cache.NewMapCache(30, 30)

	found := myCache.Has(cacheKey)
	assert.False(t, found, "item already existed in cache")

	record := httptest.NewRecorder()
	record.WriteString("test Item")
	result := record.Result()

	myCache.Set(cacheKey, result)

	found = myCache.Has(cacheKey)
	assert.True(t, found, "item does not exist in cache")

	pulledFromCache := myCache.Get(cacheKey)

	assert.Equal(t, result, pulledFromCache, "Item pulled from cache was not the same")

}
