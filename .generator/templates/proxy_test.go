package okta

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Config_Proxy(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is proxy server end point"))
	})
	proxyServer := httptest.NewServer(mux)
	defer proxyServer.Close()
	configuration := NewConfiguration()
	configuration.Debug = false
	proxyURL, err := url.Parse(proxyServer.URL)
	require.NoError(t, err, "Parse url should not error")
	password, _ := proxyURL.User.Password()
	hp := strings.Split(proxyURL.Host, ":")
	require.Equal(t, 2, len(hp), "Host should only host and port")
	intVar, err := strconv.Atoi(hp[1])
	require.NoError(t, err, "Convert string to int should not error")
	configuration.Okta.Client.Proxy.Host = hp[0]
	configuration.Okta.Client.Proxy.Port = int32(intVar)
	configuration.Okta.Client.Proxy.Username = proxyURL.User.Username()
	configuration.Okta.Client.Proxy.Password = password
	proxyClient := NewAPIClient(configuration)
	req, err := http.NewRequest(http.MethodGet, "http://example.com", nil)
	require.NoError(t, err, "Create new http request should not error")
	resp, err := proxyClient.callAPI(req)
	require.NoError(t, err, "Make http request should not error")
	b, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Read http response should not error")
	assert.Equal(t, "This is proxy server end point", string(b))
}
