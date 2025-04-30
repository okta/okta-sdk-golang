package okta

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Cache_Key(t *testing.T) {
	var buff io.ReadWriter
	request, _ := http.NewRequest("GET", "https://example.com/sample/cache-key/test+test@test."+"com?with=a&query=string", buff)
	cacheKey := CreateCacheKey(request)
	assert.Equal(t, "https://example.com/sample/cache-key/test+test@test.com?with=a&query=string", cacheKey, "The cache key was not created correctly.")
}

func Test_Item_Stored_Successful(t *testing.T) {
	var buff io.ReadWriter
	url := "https://example.com/sample/cache-key/"
	request, _ := http.NewRequest("GET", url, buff)
	cacheKey := CreateCacheKey(request)
	myCache := NewGoCache(30, 30)
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

func Test_Item_Delete_Successful(t *testing.T) {
	var buff io.ReadWriter
	url := "https://example.com/sample/cache-key/delete"
	request, _ := http.NewRequest("GET", url, buff)
	cacheKey := CreateCacheKey(request)
	myCache := NewGoCache(30, 30)
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

func Test_Cache_Cleared_Successful(t *testing.T) {
	var buff io.ReadWriter
	url := "https://example.com/sample/cache-key/clear"
	request, _ := http.NewRequest("GET", url, buff)
	cacheKey := CreateCacheKey(request)
	myCache := NewGoCache(30, 30)
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
