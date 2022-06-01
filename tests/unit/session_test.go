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
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/okta/okta-sdk-golang/v2/okta"

	"github.com/jarcoal/httpmock"
	"github.com/okta/okta-sdk-golang/v2/tests"
)

func Test_Create_Session(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	ctx, client, err := tests.NewClient(context.TODO(), okta.WithCache(false))
	require.NoError(t, err, "failed to create client")

	httpmock.RegisterResponder("POST", "/api/v1/sessions",
		tests.MockResponse(
			tests.MockSessionCreateResponse(),
		),
	)

	csr := okta.CreateSessionRequest{
		SessionToken: "abc123",
	}

	_, resp, err := client.Session.CreateSession(ctx, csr)
	require.Nil(t, err, "Error should have been nil")
	require.NotNil(t, resp, "Response was nil")

	httpmock.GetTotalCallCount()
	info := httpmock.GetCallCountInfo()
	require.Equal(t, 1, info["POST /api/v1/sessions"], "did not make exactly 1 call to /api/v1/sessions")
}
