/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)


// EntitlementValuesResponseLinks struct for EntitlementValuesResponseLinks
type EntitlementValuesResponseLinks struct {
	EntitlementValuesResponseLinksAnyOf *EntitlementValuesResponseLinksAnyOf
	LinksNext *LinksNext
	LinksSelf *LinksSelf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntitlementValuesResponseLinks) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EntitlementValuesResponseLinksAnyOf
	err = json.Unmarshal(data, &dst.EntitlementValuesResponseLinksAnyOf);
	if err == nil {
		jsonEntitlementValuesResponseLinksAnyOf, _ := json.Marshal(dst.EntitlementValuesResponseLinksAnyOf)
		if string(jsonEntitlementValuesResponseLinksAnyOf) == "{}" { // empty struct
			dst.EntitlementValuesResponseLinksAnyOf = nil
		} else {
			return nil // data stored in dst.EntitlementValuesResponseLinksAnyOf, return on the first match
		}
	} else {
		dst.EntitlementValuesResponseLinksAnyOf = nil
	}

	// try to unmarshal JSON data into LinksNext
	err = json.Unmarshal(data, &dst.LinksNext);
	if err == nil {
		jsonLinksNext, _ := json.Marshal(dst.LinksNext)
		if string(jsonLinksNext) == "{}" { // empty struct
			dst.LinksNext = nil
		} else {
			return nil // data stored in dst.LinksNext, return on the first match
		}
	} else {
		dst.LinksNext = nil
	}

	// try to unmarshal JSON data into LinksSelf
	err = json.Unmarshal(data, &dst.LinksSelf);
	if err == nil {
		jsonLinksSelf, _ := json.Marshal(dst.LinksSelf)
		if string(jsonLinksSelf) == "{}" { // empty struct
			dst.LinksSelf = nil
		} else {
			return nil // data stored in dst.LinksSelf, return on the first match
		}
	} else {
		dst.LinksSelf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EntitlementValuesResponseLinks)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntitlementValuesResponseLinks) MarshalJSON() ([]byte, error) {
	if src.EntitlementValuesResponseLinksAnyOf != nil {
		return json.Marshal(&src.EntitlementValuesResponseLinksAnyOf)
	}

	if src.LinksNext != nil {
		return json.Marshal(&src.LinksNext)
	}

	if src.LinksSelf != nil {
		return json.Marshal(&src.LinksSelf)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableEntitlementValuesResponseLinks struct {
	value *EntitlementValuesResponseLinks
	isSet bool
}

func (v NullableEntitlementValuesResponseLinks) Get() *EntitlementValuesResponseLinks {
	return v.value
}

func (v *NullableEntitlementValuesResponseLinks) Set(val *EntitlementValuesResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValuesResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValuesResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValuesResponseLinks(val *EntitlementValuesResponseLinks) *NullableEntitlementValuesResponseLinks {
	return &NullableEntitlementValuesResponseLinks{value: val, isSet: true}
}

func (v NullableEntitlementValuesResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValuesResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


