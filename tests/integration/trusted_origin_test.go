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

package integration

import (
	"context"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/tests"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_trusted_origin_for_iframe_embedding(t *testing.T) {
	ctx, client, err := tests.NewClient(context.TODO())
	require.NoError(t, err)

	scopeCors := &okta.Scope{
		Type: "CORS",
	}
	scopeIframe := &okta.Scope{
		Type: "IFRAME_EMBED",
	}
	scopes := make([]*okta.Scope, 2)
	scopes[0] = scopeCors
	scopes[1] = scopeIframe

	trustedOrigin := &okta.TrustedOrigin{
		Name:   "Trusted Origin",
		Origin: "http://example.com",
		Scopes: scopes,
	}

	to, _, err := client.TrustedOrigin.CreateOrigin(ctx, *trustedOrigin)

	require.NoError(t, err)

	assert.IsType(t, okta.TrustedOrigin{}, to)

	client.TrustedOrigin.DeleteOrigin(ctx, to.Id)
}
