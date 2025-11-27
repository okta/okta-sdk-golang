/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2025 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// GroupProfile Specifies required and optional properties for a group. The `objectClass` of a group determines which additional properties are available.  You can extend group profiles with custom properties, but you must first add the properties to the group profile schema before you can reference them. Use the Profile Editor in the Admin Console or the [Schemas API](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Schema/) to manage schema extensions.  Custom properties can contain HTML tags. It is the client's responsibility to escape or encode this data before displaying it. Use [best-practices](https://cheatsheetseries.owasp.org/cheatsheets/Cross_Site_Scripting_Prevention_Cheat_Sheet.html) to prevent cross-site scripting.
type GroupProfile struct {
	OktaActiveDirectoryGroupProfile *OktaActiveDirectoryGroupProfile
	OktaUserGroupProfile            *OktaUserGroupProfile
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GroupProfile) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OktaActiveDirectoryGroupProfile
	err = json.Unmarshal(data, &dst.OktaActiveDirectoryGroupProfile)
	if err == nil {
		jsonOktaActiveDirectoryGroupProfile, _ := json.Marshal(dst.OktaActiveDirectoryGroupProfile)
		if string(jsonOktaActiveDirectoryGroupProfile) == "{}" { // empty struct
			dst.OktaActiveDirectoryGroupProfile = nil
		} else {
			return nil // data stored in dst.OktaActiveDirectoryGroupProfile, return on the first match
		}
	} else {
		dst.OktaActiveDirectoryGroupProfile = nil
	}

	// try to unmarshal JSON data into OktaUserGroupProfile
	err = json.Unmarshal(data, &dst.OktaUserGroupProfile)
	if err == nil {
		jsonOktaUserGroupProfile, _ := json.Marshal(dst.OktaUserGroupProfile)
		if string(jsonOktaUserGroupProfile) == "{}" { // empty struct
			dst.OktaUserGroupProfile = nil
		} else {
			return nil // data stored in dst.OktaUserGroupProfile, return on the first match
		}
	} else {
		dst.OktaUserGroupProfile = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GroupProfile)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GroupProfile) MarshalJSON() ([]byte, error) {
	if src.OktaActiveDirectoryGroupProfile != nil {
		return json.Marshal(&src.OktaActiveDirectoryGroupProfile)
	}

	if src.OktaUserGroupProfile != nil {
		return json.Marshal(&src.OktaUserGroupProfile)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGroupProfile struct {
	value *GroupProfile
	isSet bool
}

func (v NullableGroupProfile) Get() *GroupProfile {
	return v.value
}

func (v *NullableGroupProfile) Set(val *GroupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupProfile(val *GroupProfile) *NullableGroupProfile {
	return &NullableGroupProfile{value: val, isSet: true}
}

func (v NullableGroupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
