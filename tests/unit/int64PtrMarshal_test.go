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
	"encoding/json"
	"testing"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/stretchr/testify/require"
)

// TestInt64PtrMarshal tests the generated code marshaling for the int64
// pointers using PasswordPolicy as test case.
func TestInt64PtrMarshal(t *testing.T) {
	// Example mainting old behavior, Priority is 0 so it doesn't get
	// marshalled/unmarshalled at all and the pointer is not utilized. This is
	// to keep the code backwards compatible.
	example := okta.PasswordPolicy{Priority: 0}
	_json, _ := json.Marshal(&example)
	var passwordPolicy okta.PasswordPolicy
	_ = json.Unmarshal(_json, &passwordPolicy)
	require.Equal(t, int64(0), passwordPolicy.Priority)
	require.Nil(t, passwordPolicy.PriorityPtr)

	// Example mainting old behavior, Priority is 1 so it gets
	// marshalled/marshalled untilizing the int pointer. This is to keep the
	// code backwards compatible.
	example = okta.PasswordPolicy{Priority: 1}
	_json, _ = json.Marshal(&example)
	passwordPolicy = okta.PasswordPolicy{}
	_ = json.Unmarshal(_json, &passwordPolicy)
	require.Equal(t, int64(1), passwordPolicy.Priority)
	require.NotNil(t, passwordPolicy.PriorityPtr)
	require.Equal(t, int64(1), *passwordPolicy.PriorityPtr)

	// Exmaple new behavior, int pointer maintains value during
	// marshal/unmarshal
	example = okta.PasswordPolicy{PriorityPtr: okta.Int64Ptr(9)}
	_json, _ = json.Marshal(&example)
	passwordPolicy = okta.PasswordPolicy{}
	_ = json.Unmarshal(_json, &passwordPolicy)
	require.Equal(t, int64(9), passwordPolicy.Priority)
	require.NotNil(t, passwordPolicy.PriorityPtr)
	require.Equal(t, int64(9), *passwordPolicy.PriorityPtr)
}
