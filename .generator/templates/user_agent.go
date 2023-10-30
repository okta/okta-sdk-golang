package okta

import "runtime"

const Version = "3.0.16"

type UserAgent struct {
	goVersion string

	osName string

	osVersion string

	config *Configuration
}

func NewUserAgent(config *Configuration) UserAgent {
	ua := UserAgent{}
	ua.config = config
	ua.goVersion = runtime.Version()
	ua.osName = runtime.GOOS
	ua.osVersion = runtime.GOARCH

	return ua
}

func (ua UserAgent) String() string {
	userAgentString := "okta-sdk-golang/" + Version + " "
	userAgentString += "golang/" + ua.goVersion + " "
	userAgentString += ua.osName + "/" + ua.osVersion

	if ua.config.UserAgentExtra != "" {
		userAgentString += " " + ua.config.UserAgentExtra
	}

	return userAgentString
}
