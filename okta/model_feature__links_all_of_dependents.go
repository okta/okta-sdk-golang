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

// FeatureLinksAllOfDependents Link to feature dependents
type FeatureLinksAllOfDependents struct {
	// Link URI
	Href *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeatureLinksAllOfDependents FeatureLinksAllOfDependents

// NewFeatureLinksAllOfDependents instantiates a new FeatureLinksAllOfDependents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureLinksAllOfDependents() *FeatureLinksAllOfDependents {
	this := FeatureLinksAllOfDependents{}
	return &this
}

// NewFeatureLinksAllOfDependentsWithDefaults instantiates a new FeatureLinksAllOfDependents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureLinksAllOfDependentsWithDefaults() *FeatureLinksAllOfDependents {
	this := FeatureLinksAllOfDependents{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *FeatureLinksAllOfDependents) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureLinksAllOfDependents) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *FeatureLinksAllOfDependents) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *FeatureLinksAllOfDependents) SetHref(v string) {
	o.Href = &v
}

func (o FeatureLinksAllOfDependents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FeatureLinksAllOfDependents) UnmarshalJSON(bytes []byte) (err error) {
	varFeatureLinksAllOfDependents := _FeatureLinksAllOfDependents{}

	err = json.Unmarshal(bytes, &varFeatureLinksAllOfDependents)
	if err == nil {
		*o = FeatureLinksAllOfDependents(varFeatureLinksAllOfDependents)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFeatureLinksAllOfDependents struct {
	value *FeatureLinksAllOfDependents
	isSet bool
}

func (v NullableFeatureLinksAllOfDependents) Get() *FeatureLinksAllOfDependents {
	return v.value
}

func (v *NullableFeatureLinksAllOfDependents) Set(val *FeatureLinksAllOfDependents) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureLinksAllOfDependents) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureLinksAllOfDependents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureLinksAllOfDependents(val *FeatureLinksAllOfDependents) *NullableFeatureLinksAllOfDependents {
	return &NullableFeatureLinksAllOfDependents{value: val, isSet: true}
}

func (v NullableFeatureLinksAllOfDependents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureLinksAllOfDependents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

