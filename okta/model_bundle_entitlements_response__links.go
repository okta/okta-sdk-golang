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

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// BundleEntitlementsResponseLinks struct for BundleEntitlementsResponseLinks
type BundleEntitlementsResponseLinks struct {
	LinksNext *LinksNext
	LinksSelf *LinksSelf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BundleEntitlementsResponseLinks) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(BundleEntitlementsResponseLinks)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BundleEntitlementsResponseLinks) MarshalJSON() ([]byte, error) {
	if src.LinksNext != nil {
		return json.Marshal(&src.LinksNext)
	}

	if src.LinksSelf != nil {
		return json.Marshal(&src.LinksSelf)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBundleEntitlementsResponseLinks struct {
	value *BundleEntitlementsResponseLinks
	isSet bool
}

func (v NullableBundleEntitlementsResponseLinks) Get() *BundleEntitlementsResponseLinks {
	return v.value
}

func (v *NullableBundleEntitlementsResponseLinks) Set(val *BundleEntitlementsResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleEntitlementsResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleEntitlementsResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleEntitlementsResponseLinks(val *BundleEntitlementsResponseLinks) *NullableBundleEntitlementsResponseLinks {
	return &NullableBundleEntitlementsResponseLinks{value: val, isSet: true}
}

func (v NullableBundleEntitlementsResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleEntitlementsResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


