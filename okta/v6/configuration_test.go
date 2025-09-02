package okta

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserAgent(t *testing.T) {
	configuration, err := NewConfiguration()
	require.NoError(t, err, "Creating a new config should not error")
	userAgent := "okta-sdk-golang/" + VERSION + " golang/" + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH
	require.Equal(t, userAgent, configuration.UserAgent)
}

func TestUserAgentWithExtra(t *testing.T) {
	configuration, err := NewConfiguration(
		WithUserAgentExtra("extra/info"),
	)
	require.NoError(t, err, "Creating a new config should not error")
	userAgent := "okta-sdk-golang/" + VERSION + " golang/" + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH + " extra/info"
	require.Equal(t, userAgent, configuration.UserAgent)
}
