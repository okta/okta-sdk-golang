/*
Okta Admin Management API

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
)

// checks if the GracePeriodRequiredSoon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GracePeriodRequiredSoon{}

// GracePeriodRequiredSoon <x-lifecycle-container><x-lifecycle class=\"ea\"></x-lifecycle></x-lifecycle-container>Customizable strings to use with [grace periods](https://developer.okta.com/docs/concepts/policies/#authenticator-enrollment-policies) that are shown on the Sign-In Widget  You can use the `gracePeriodRequiredSoonDescription` string without using either of the custom link strings. However, if you use one of the custom link strings (`gracePeriodRequiredSoonCustomLinkLabel` or `gracePeriodRequiredSoonCustomLinkUrl`), then both are required.
type GracePeriodRequiredSoon struct {
	// The label of the custom link that's shown on the Sign-In Widget when users are prompted to enroll required authenticators before their grace period ends.
	GracePeriodRequiredSoonCustomLinkLabel *string `json:"gracePeriodRequiredSoonCustomLinkLabel,omitempty"`
	// The URL for the custom link that's shown on the Sign-In Widget when users are prompted to enroll required authenticators before their grace period ends.
	GracePeriodRequiredSoonCustomLinkUrl *string `json:"gracePeriodRequiredSoonCustomLinkUrl,omitempty"`
	// The description that's shown on the Sign-In Widget for users who are within an authenticator grace period. This description prompts users to enroll required authenticators before their grace period ends.
	GracePeriodRequiredSoonDescription *string `json:"gracePeriodRequiredSoonDescription,omitempty"`
	AdditionalProperties               map[string]interface{}
}

type _GracePeriodRequiredSoon GracePeriodRequiredSoon

// NewGracePeriodRequiredSoon instantiates a new GracePeriodRequiredSoon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGracePeriodRequiredSoon() *GracePeriodRequiredSoon {
	this := GracePeriodRequiredSoon{}
	return &this
}

// NewGracePeriodRequiredSoonWithDefaults instantiates a new GracePeriodRequiredSoon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGracePeriodRequiredSoonWithDefaults() *GracePeriodRequiredSoon {
	this := GracePeriodRequiredSoon{}
	return &this
}

// GetGracePeriodRequiredSoonCustomLinkLabel returns the GracePeriodRequiredSoonCustomLinkLabel field value if set, zero value otherwise.
func (o *GracePeriodRequiredSoon) GetGracePeriodRequiredSoonCustomLinkLabel() string {
	if o == nil || IsNil(o.GracePeriodRequiredSoonCustomLinkLabel) {
		var ret string
		return ret
	}
	return *o.GracePeriodRequiredSoonCustomLinkLabel
}

// GetGracePeriodRequiredSoonCustomLinkLabelOk returns a tuple with the GracePeriodRequiredSoonCustomLinkLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GracePeriodRequiredSoon) GetGracePeriodRequiredSoonCustomLinkLabelOk() (*string, bool) {
	if o == nil || IsNil(o.GracePeriodRequiredSoonCustomLinkLabel) {
		return nil, false
	}
	return o.GracePeriodRequiredSoonCustomLinkLabel, true
}

// HasGracePeriodRequiredSoonCustomLinkLabel returns a boolean if a field has been set.
func (o *GracePeriodRequiredSoon) HasGracePeriodRequiredSoonCustomLinkLabel() bool {
	if o != nil && !IsNil(o.GracePeriodRequiredSoonCustomLinkLabel) {
		return true
	}

	return false
}

// SetGracePeriodRequiredSoonCustomLinkLabel gets a reference to the given string and assigns it to the GracePeriodRequiredSoonCustomLinkLabel field.
func (o *GracePeriodRequiredSoon) SetGracePeriodRequiredSoonCustomLinkLabel(v string) {
	o.GracePeriodRequiredSoonCustomLinkLabel = &v
}

// GetGracePeriodRequiredSoonCustomLinkUrl returns the GracePeriodRequiredSoonCustomLinkUrl field value if set, zero value otherwise.
func (o *GracePeriodRequiredSoon) GetGracePeriodRequiredSoonCustomLinkUrl() string {
	if o == nil || IsNil(o.GracePeriodRequiredSoonCustomLinkUrl) {
		var ret string
		return ret
	}
	return *o.GracePeriodRequiredSoonCustomLinkUrl
}

// GetGracePeriodRequiredSoonCustomLinkUrlOk returns a tuple with the GracePeriodRequiredSoonCustomLinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GracePeriodRequiredSoon) GetGracePeriodRequiredSoonCustomLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.GracePeriodRequiredSoonCustomLinkUrl) {
		return nil, false
	}
	return o.GracePeriodRequiredSoonCustomLinkUrl, true
}

// HasGracePeriodRequiredSoonCustomLinkUrl returns a boolean if a field has been set.
func (o *GracePeriodRequiredSoon) HasGracePeriodRequiredSoonCustomLinkUrl() bool {
	if o != nil && !IsNil(o.GracePeriodRequiredSoonCustomLinkUrl) {
		return true
	}

	return false
}

// SetGracePeriodRequiredSoonCustomLinkUrl gets a reference to the given string and assigns it to the GracePeriodRequiredSoonCustomLinkUrl field.
func (o *GracePeriodRequiredSoon) SetGracePeriodRequiredSoonCustomLinkUrl(v string) {
	o.GracePeriodRequiredSoonCustomLinkUrl = &v
}

// GetGracePeriodRequiredSoonDescription returns the GracePeriodRequiredSoonDescription field value if set, zero value otherwise.
func (o *GracePeriodRequiredSoon) GetGracePeriodRequiredSoonDescription() string {
	if o == nil || IsNil(o.GracePeriodRequiredSoonDescription) {
		var ret string
		return ret
	}
	return *o.GracePeriodRequiredSoonDescription
}

// GetGracePeriodRequiredSoonDescriptionOk returns a tuple with the GracePeriodRequiredSoonDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GracePeriodRequiredSoon) GetGracePeriodRequiredSoonDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.GracePeriodRequiredSoonDescription) {
		return nil, false
	}
	return o.GracePeriodRequiredSoonDescription, true
}

// HasGracePeriodRequiredSoonDescription returns a boolean if a field has been set.
func (o *GracePeriodRequiredSoon) HasGracePeriodRequiredSoonDescription() bool {
	if o != nil && !IsNil(o.GracePeriodRequiredSoonDescription) {
		return true
	}

	return false
}

// SetGracePeriodRequiredSoonDescription gets a reference to the given string and assigns it to the GracePeriodRequiredSoonDescription field.
func (o *GracePeriodRequiredSoon) SetGracePeriodRequiredSoonDescription(v string) {
	o.GracePeriodRequiredSoonDescription = &v
}

func (o GracePeriodRequiredSoon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GracePeriodRequiredSoon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GracePeriodRequiredSoonCustomLinkLabel) {
		toSerialize["gracePeriodRequiredSoonCustomLinkLabel"] = o.GracePeriodRequiredSoonCustomLinkLabel
	}
	if !IsNil(o.GracePeriodRequiredSoonCustomLinkUrl) {
		toSerialize["gracePeriodRequiredSoonCustomLinkUrl"] = o.GracePeriodRequiredSoonCustomLinkUrl
	}
	if !IsNil(o.GracePeriodRequiredSoonDescription) {
		toSerialize["gracePeriodRequiredSoonDescription"] = o.GracePeriodRequiredSoonDescription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GracePeriodRequiredSoon) UnmarshalJSON(data []byte) (err error) {
	varGracePeriodRequiredSoon := _GracePeriodRequiredSoon{}

	err = json.Unmarshal(data, &varGracePeriodRequiredSoon)

	if err != nil {
		return err
	}

	*o = GracePeriodRequiredSoon(varGracePeriodRequiredSoon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gracePeriodRequiredSoonCustomLinkLabel")
		delete(additionalProperties, "gracePeriodRequiredSoonCustomLinkUrl")
		delete(additionalProperties, "gracePeriodRequiredSoonDescription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGracePeriodRequiredSoon struct {
	value *GracePeriodRequiredSoon
	isSet bool
}

func (v NullableGracePeriodRequiredSoon) Get() *GracePeriodRequiredSoon {
	return v.value
}

func (v *NullableGracePeriodRequiredSoon) Set(val *GracePeriodRequiredSoon) {
	v.value = val
	v.isSet = true
}

func (v NullableGracePeriodRequiredSoon) IsSet() bool {
	return v.isSet
}

func (v *NullableGracePeriodRequiredSoon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGracePeriodRequiredSoon(val *GracePeriodRequiredSoon) *NullableGracePeriodRequiredSoon {
	return &NullableGracePeriodRequiredSoon{value: val, isSet: true}
}

func (v NullableGracePeriodRequiredSoon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGracePeriodRequiredSoon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
