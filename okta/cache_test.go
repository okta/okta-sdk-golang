package okta

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestOAuthTokensAlwaysCached(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	WithCache(false)
	cfg := NewConfiguration(
		WithCache(false),
		WithOrgUrl("https://testing.oktapreview.com"),
		WithAuthorizationMode("PrivateKey"),
		WithClientId("abc"),
		WithPrivateKey(`
-----BEGIN RSA PRIVATE KEY-----
MIIBOgIBAAJBAKj34GkxFhD90vcNLYLInFEX6Ppy1tPf9Cnzj4p4WGeKLs1Pt8Qu
KUpRKfFLfRYC9AIKjbJTWit+CqvjWYzvQwECAwEAAQJAIJLixBy2qpFoS4DSmoEm
o3qGy0t6z09AIJtH+5OeRV1be+N4cDYJKffGzDa88vQENZiRm0GRq6a+HPGQMd2k
TQIhAKMSvzIBnni7ot/OSie2TmJLY4SwTQAevXysE2RbFDYdAiEBCUEaRQnMnbp7
9mxDXDf6AU0cN/RPBjb9qSHDcWZHGzUCIG2Es59z8ugGrDY+pxLQnwfotadxd+Uy
v/Ow5T0q5gIJAiEAyS4RaI9YG8EWx/2w0T67ZUVAw8eOMB6BIUg0Xcu+3okCIBOs
/5OiPgoTdSy7bcF9IGpSE8ZgGKzgYQVZeN97YE00
-----END RSA PRIVATE KEY-----
		`),
		WithScopes(([]string{"okta.users.read"})),
	)

	client := NewAPIClient(cfg)

	accessToken := RequestAccessToken{
		TokenType:   "Bearer",
		ExpiresIn:   3600,
		AccessToken: "xyz",
		Scope:       "okta.users.read",
	}
	httpmockTokenURLRegex := `=~^https://testing\.oktapreview\.com/oauth2/v1/token\?client_assertion=.*\z`
	jsonResp, err := httpmock.NewJsonResponder(200, accessToken)
	require.NoError(t, err)
	httpmock.RegisterResponder("POST", httpmockTokenURLRegex, jsonResp)

	adminConsole := Application{}
	adminConsole.SetId("abc123")
	adminConsole.SetStatus("ACTIVE")
	adminConsole.SetLabel("Okta Admin Console")
	apps1 := []*Application{
		&adminConsole,
	}
	jsonResp, err = httpmock.NewJsonResponder(200, apps1)
	require.NoError(t, err)
	httpmockAdminConsoleRegex := `=~^https://testing\.oktapreview\.com/api/v1/apps?.*q\=Okta\+Admin\+Console.*\z`
	httpmock.RegisterResponder("GET", httpmockAdminConsoleRegex, jsonResp)

	dashboard := Application{}
	adminConsole.SetId("def456")
	adminConsole.SetStatus("ACTIVE")
	adminConsole.SetLabel("Okta Dashboard")
	apps2 := []*Application{
		&dashboard,
	}
	jsonResp, err = httpmock.NewJsonResponder(200, apps2)
	require.NoError(t, err)
	httpmockDashboardRegex := `=~^https://testing\.oktapreview\.com/api/v1/apps?.*q\=Okta\+Dashboard.*\z`
	httpmock.RegisterResponder("GET", httpmockDashboardRegex, jsonResp)

	_, _, err = client.ApplicationApi.ListApplications(cfg.Context).Limit(1).Filter("status eq ACTIVE").Q("Okta Admin Console").Execute()
	require.NoError(t, err)
	_, _, err = client.ApplicationApi.ListApplications(cfg.Context).Limit(1).Filter("status eq ACTIVE").Q("Okta Admin Console").Execute()
	require.NoError(t, err)

	_, _, err = client.ApplicationApi.ListApplications(cfg.Context).Limit(1).Filter("status eq ACTIVE").Q("Okta Dashboard").Execute()
	require.NoError(t, err)
	_, _, err = client.ApplicationApi.ListApplications(cfg.Context).Limit(1).Filter("status eq ACTIVE").Q("Okta Dashboard").Execute()
	require.NoError(t, err)

	info := httpmock.GetCallCountInfo()
	totalCalls := httpmock.GetTotalCallCount()

	assert.Equal(t, 5, totalCalls, fmt.Sprintf("there should only be 5 API calls in this test but there were %d calls", totalCalls))
	// Tokens from requests should be cached.
	require.True(t, info[fmt.Sprintf("POST %s", httpmockTokenURLRegex)] == 1, "tokens endpoint should only be called once")

	// But all other requests should not be cached.
	require.True(t, info[fmt.Sprintf("GET %s", httpmockAdminConsoleRegex)] == 2)
	require.True(t, info[fmt.Sprintf("GET %s", httpmockDashboardRegex)] == 2)
}
