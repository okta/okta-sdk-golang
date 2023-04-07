package unit

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"
	"github.com/stretchr/testify/assert"
)

// readerFun makes it easier to implement an inline reader as a closure function.
type readerFun func(p []byte) (n int, err error)

// Read, part of io.Reader interface.
func (r readerFun) Read(p []byte) (n int, err error) { return r(p) }

// slowTransport provides a dummy http-like transport serving fixed content, but slowly.
type slowTransport struct{}

// RoundTrip, part of http.Transport interface. This servers 42 as a JSON response, but slowly.
// In particular, we serve the response immediately, but getting the body takes a second.
func (t slowTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	realBody := strings.NewReader("42")
	// This body takes 1 second to read. It also needs a valid context for the whole duration.
	slowBody := func(p []byte) (n int, err error) {
		select {
		case <-req.Context().Done():
			return 0, req.Context().Err()
		case <-time.After(1 * time.Second):
			return realBody.Read(p)
		}
	}

	rsp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(readerFun(slowBody)),
		Header:     http.Header{},
		Request:    req,
	}
	rsp.Header.Set("Content-Type", "application/json")
	return rsp, nil
}

// TestExecuteRequest tests that the request executor can handle a slow response.
// In particular, we want to make sure that the context is properly passed through
// and not canceled too early.
func TestExecuteRequest(t *testing.T) {
	cfg := []okta.ConfigSetter{
		okta.WithOrgUrl("https://fakeurl"),                               // This is ignored, but required for validator.
		okta.WithToken("foo"),                                            // ditto.
		okta.WithHttpClientPtr(&http.Client{Transport: slowTransport{}}), // Use our more realistic transport.
		okta.WithRequestTimeout(10),                                      // The context issues are gated with actually having a timeout.
	}
	ctx, cl, err := tests.NewClient(context.Background(), cfg...)
	assert.NoError(t, err, "Basic client errors")
	req, err := http.NewRequest("GET", "https://fakeurl", http.NoBody)
	assert.NoError(t, err, "Request building")
	var out int
	rs, err := cl.GetRequestExecutor().Do(ctx, req, &out)
	assert.NoError(t, err, "Request execution")
	if rs.StatusCode != 200 || out != 42 {
		t.Errorf("Got val=%d status=%d, want 42 status=200", out, rs.StatusCode)
	}
}
