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
)

// HrefObjectAuthorizeLink Link to authorize scopes
type HrefObjectAuthorizeLink struct {
	Hints *HrefHintsGuidanceObject `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	AdditionalProperties map[string]interface{}
}

type _HrefObjectAuthorizeLink HrefObjectAuthorizeLink

// NewHrefObjectAuthorizeLink instantiates a new HrefObjectAuthorizeLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHrefObjectAuthorizeLink(href string) *HrefObjectAuthorizeLink {
	this := HrefObjectAuthorizeLink{}
	this.Href = href
	return &this
}

// NewHrefObjectAuthorizeLinkWithDefaults instantiates a new HrefObjectAuthorizeLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHrefObjectAuthorizeLinkWithDefaults() *HrefObjectAuthorizeLink {
	this := HrefObjectAuthorizeLink{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *HrefObjectAuthorizeLink) GetHints() HrefHintsGuidanceObject {
	if o == nil || o.Hints == nil {
		var ret HrefHintsGuidanceObject
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HrefObjectAuthorizeLink) GetHintsOk() (*HrefHintsGuidanceObject, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *HrefObjectAuthorizeLink) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given HrefHintsGuidanceObject and assigns it to the Hints field.
func (o *HrefObjectAuthorizeLink) SetHints(v HrefHintsGuidanceObject) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *HrefObjectAuthorizeLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *HrefObjectAuthorizeLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *HrefObjectAuthorizeLink) SetHref(v string) {
	o.Href = v
}

func (o HrefObjectAuthorizeLink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	if true {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HrefObjectAuthorizeLink) UnmarshalJSON(bytes []byte) (err error) {
	varHrefObjectAuthorizeLink := _HrefObjectAuthorizeLink{}

	err = json.Unmarshal(bytes, &varHrefObjectAuthorizeLink)
	if err == nil {
		*o = HrefObjectAuthorizeLink(varHrefObjectAuthorizeLink)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableHrefObjectAuthorizeLink struct {
	value *HrefObjectAuthorizeLink
	isSet bool
}

func (v NullableHrefObjectAuthorizeLink) Get() *HrefObjectAuthorizeLink {
	return v.value
}

func (v *NullableHrefObjectAuthorizeLink) Set(val *HrefObjectAuthorizeLink) {
	v.value = val
	v.isSet = true
}

func (v NullableHrefObjectAuthorizeLink) IsSet() bool {
	return v.isSet
}

func (v *NullableHrefObjectAuthorizeLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHrefObjectAuthorizeLink(val *HrefObjectAuthorizeLink) *NullableHrefObjectAuthorizeLink {
	return &NullableHrefObjectAuthorizeLink{value: val, isSet: true}
}

func (v NullableHrefObjectAuthorizeLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHrefObjectAuthorizeLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

