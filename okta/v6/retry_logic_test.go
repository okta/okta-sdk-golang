package okta

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func Test_429_Will_Automatically_Retry(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	configuration, err := NewConfiguration()
	require.NoError(t, err, "Creating a new config should not error")
	configuration.Okta.Client.RateLimit.MaxRetries = 2
	configuration.Debug = true
	proxyClient := NewAPIClient(configuration)
	httpmock.RegisterResponder("GET", "/api/v1/users",
		MockResponse(
			Mock429Response(),
			MockValidResponse(),
		),
	)

	_, resp, err := proxyClient.UserAPI.ListUsers(apiClient.cfg.Context).Execute()
	require.Nil(t, err, "Error should have been nil")
	require.NotNil(t, resp, "Response was nil")

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()
	require.Equal(t, 2, info["GET /api/v1/users"], "did not make exactly 2 call to /api/v1/users")
}

func MockResponse(responses ...*http.Response) httpmock.Responder {
	return func(req *http.Request) (*http.Response, error) {
		httpmock.GetTotalCallCount()
		info := httpmock.GetCallCountInfo()
		count := info[req.Method+" "+req.URL.Path]

		if len(responses) >= count {
			return responses[count-1], nil
		}

		return nil, fmt.Errorf("no response found for call %v to %s", count, req.URL.Path)
	}
}

func Mock429Response() *http.Response {
	loc, _ := time.LoadLocation("UTC")
	zulu := time.Now().In(loc)
	header := http.Header{}
	header.Add("X-Okta-Now", strconv.FormatInt(zulu.Unix(), 10))
	header.Add("X-Rate-Limit-Reset", strconv.FormatInt(time.Now().Unix()+1, 10))
	header.Add("X-Okta-Request-id", "a-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Date", zulu.Format("Mon, 02 Jan 2006 15:04:05 GMT"))

	return &http.Response{
		Status:        strconv.Itoa(429),
		StatusCode:    429,
		Body:          httpmock.NewRespBodyFromString("{}"),
		Header:        header,
		ContentLength: -1,
	}
}

func MockValidResponse() *http.Response {
	header := http.Header{}
	header.Add("X-Okta-Request-id", "another-request-id")
	header.Add("Content-Type", "application/json")
	header.Add("Date", time.Now().Add(time.Second*10).Format(time.RFC3339))

	return &http.Response{
		Status:        strconv.Itoa(200),
		StatusCode:    200,
		Body:          httpmock.NewRespBodyFromString("[]"),
		Header:        header,
		ContentLength: -1,
	}
}

func Test_RateLimitPrevent_Blocks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	configuration, err := NewConfiguration(WithRateLimitPrevent(true))
	require.NoError(t, err, "Creating a new config should not error")
	configuration.Debug = true

	proxyClient := NewAPIClient(configuration)
	resetTime := time.Now().Add(5 * time.Second)
	count := 0

	// Register mock response
	httpmock.RegisterResponder("GET", "/api/v1/users", func(*http.Request) (*http.Response, error) {
		if count <= 1 {
			count++
			now := time.Now().UTC()
			dateString := now.Format("Mon, 02 Jan 2006 15:04:05") + " GMT"
			return &http.Response{
				Status:     strconv.Itoa(200),
				StatusCode: 200,
				Body:       httpmock.NewRespBodyFromString("[]"),
				Header: http.Header{
					"X-Rate-Limit-Limit":     []string{"50"},
					"X-Rate-Limit-Remaining": []string{"0"},
					"X-Rate-Limit-Reset":     []string{strconv.FormatInt(resetTime.Unix(), 10)},
					"Content-Type":           []string{"application/json"},
					"X-Okta-Request-id":      []string{"a-request-id"},
					"Date":                   []string{dateString},
				},
				ContentLength: -1,
			}, nil

		}
		return nil, fmt.Errorf("more than 2 calls")
	})

	// First request (normal)
	_, resp, err := proxyClient.UserAPI.ListUsers(apiClient.cfg.Context).Execute()
	require.NoError(t, err, "should have been able to parse the date")
	require.NotNil(t, resp, "Response was nil")

	// Measure time before second request
	start := time.Now()

	// Second request should block due to rate limiting
	_, resp, err = proxyClient.UserAPI.ListUsers(apiClient.cfg.Context).Execute()
	require.NoError(t, err, "should have been able to parse the date")
	require.NotNil(t, resp, "Second response was nil")

	// Validate time waited
	duration := time.Since(start)
	require.GreaterOrEqual(t, duration, 5*time.Second, "Second request should have blocked for at least 5 seconds")

	// Ensure only two calls were made
	info := httpmock.GetCallCountInfo()
	require.Equal(t, 2, info["GET /api/v1/users"], "Expected exactly 2 calls to /api/v1/users")
}
