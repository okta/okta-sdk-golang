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
	configuration := NewConfiguration()
	configuration.Okta.Client.RateLimit.MaxRetries = 2
	configuration.Debug = true
	proxyClient := NewAPIClient(configuration)
	httpmock.RegisterResponder("GET", "/api/v1/users",
		MockResponse(
			Mock429Response(),
			MockValidResponse(),
		),
	)

	_, resp, err := proxyClient.UserApi.ListUsers(apiClient.cfg.Context).Execute()
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
