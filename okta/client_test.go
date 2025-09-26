package okta

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockRoundTripper struct {
	Responses []http.Response
	CallCount int
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// simulate responses for each call
	if m.CallCount < len(m.Responses) {
		resp := m.Responses[m.CallCount]
		m.CallCount++
		return &resp, nil
	}
	return &http.Response{StatusCode: http.StatusInternalServerError}, nil
}

func TestCreateRetryableHTTPClient_RetryOn500(t *testing.T) {
	bodyString := "this is a test body"
	body := io.NopCloser(strings.NewReader(bodyString))
	mockResponses := []http.Response{
		{StatusCode: 500},             // should retry
		{StatusCode: 500},             // should retry
		{StatusCode: 500},             // should retry
		{StatusCode: 500},             // should retry
		{StatusCode: 500, Body: body}, // should retry
	}

	mockTransport := &mockRoundTripper{Responses: mockResponses}

	// base client with the mock transport
	baseClient := &http.Client{
		Transport: mockTransport,
		Timeout:   10 * time.Second,
	}

	configuration, err := NewConfiguration(WithAuthorizationMode("Bearer"), WithHttpClientPtr(baseClient), WithRateLimitMaxBackOff(1), WithRateLimitMaxRetries(4))
	require.NoError(t, err, "Creating a new config should not error")
	client := NewAPIClient(configuration)

	req, err := http.NewRequest("GET", "http://foo.com", nil)
	require.NoError(t, err)

	resp, err := client.doWithRetries(client.cfg.Context, req)

	require.Error(t, err)
	require.Equal(t, "unexpected status code 500, giving up after 4 retries", err.Error())
	require.True(t, resp != nil)
	require.Equal(t, resp.StatusCode, 500)
	message, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, bodyString, string(message))
	require.Equal(t, 5, mockTransport.CallCount)
}
