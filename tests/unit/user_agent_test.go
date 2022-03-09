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
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"
)

func TestUserAgent(t *testing.T) {
	_, client, _ := tests.NewClient(context.TODO(), okta.WithOrgUrl("https://example.com"))
	agent := okta.NewUserAgent(client.GetConfig())
	userAgent := "okta-sdk-golang/" + okta.Version + " golang/" + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH
	require.Equal(t, userAgent, agent.String())
}

func TestUserAgentWithExtra(t *testing.T) {
	_, client, _ := tests.NewClient(context.TODO(),
		okta.WithOrgUrl("https://example.com"),
		okta.WithUserAgentExtra("extra"),
	)
	agent := okta.NewUserAgent(client.GetConfig())
	userAgent := "okta-sdk-golang/" + okta.Version + " golang/" + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH + " extra"
	require.Equal(t, userAgent, agent.String())
}
