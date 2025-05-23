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

// GoogleApplicationSettingsApplication Google app instance properties
type GoogleApplicationSettingsApplication struct {
	// Your Google company domain
	Domain string `json:"domain"`
	// RPID
	RpId *string `json:"rpId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GoogleApplicationSettingsApplication GoogleApplicationSettingsApplication

// NewGoogleApplicationSettingsApplication instantiates a new GoogleApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleApplicationSettingsApplication(domain string) *GoogleApplicationSettingsApplication {
	this := GoogleApplicationSettingsApplication{}
	this.Domain = domain
	return &this
}

// NewGoogleApplicationSettingsApplicationWithDefaults instantiates a new GoogleApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleApplicationSettingsApplicationWithDefaults() *GoogleApplicationSettingsApplication {
	this := GoogleApplicationSettingsApplication{}
	return &this
}

// GetDomain returns the Domain field value
func (o *GoogleApplicationSettingsApplication) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *GoogleApplicationSettingsApplication) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *GoogleApplicationSettingsApplication) SetDomain(v string) {
	o.Domain = v
}

// GetRpId returns the RpId field value if set, zero value otherwise.
func (o *GoogleApplicationSettingsApplication) GetRpId() string {
	if o == nil || o.RpId == nil {
		var ret string
		return ret
	}
	return *o.RpId
}

// GetRpIdOk returns a tuple with the RpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleApplicationSettingsApplication) GetRpIdOk() (*string, bool) {
	if o == nil || o.RpId == nil {
		return nil, false
	}
	return o.RpId, true
}

// HasRpId returns a boolean if a field has been set.
func (o *GoogleApplicationSettingsApplication) HasRpId() bool {
	if o != nil && o.RpId != nil {
		return true
	}

	return false
}

// SetRpId gets a reference to the given string and assigns it to the RpId field.
func (o *GoogleApplicationSettingsApplication) SetRpId(v string) {
	o.RpId = &v
}

func (o GoogleApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if o.RpId != nil {
		toSerialize["rpId"] = o.RpId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GoogleApplicationSettingsApplication) UnmarshalJSON(bytes []byte) (err error) {
	varGoogleApplicationSettingsApplication := _GoogleApplicationSettingsApplication{}

	err = json.Unmarshal(bytes, &varGoogleApplicationSettingsApplication)
	if err == nil {
		*o = GoogleApplicationSettingsApplication(varGoogleApplicationSettingsApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "domain")
		delete(additionalProperties, "rpId")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGoogleApplicationSettingsApplication struct {
	value *GoogleApplicationSettingsApplication
	isSet bool
}

func (v NullableGoogleApplicationSettingsApplication) Get() *GoogleApplicationSettingsApplication {
	return v.value
}

func (v *NullableGoogleApplicationSettingsApplication) Set(val *GoogleApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleApplicationSettingsApplication(val *GoogleApplicationSettingsApplication) *NullableGoogleApplicationSettingsApplication {
	return &NullableGoogleApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableGoogleApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

