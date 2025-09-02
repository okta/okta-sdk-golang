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
)

// CsrPublishHrefHints Describes allowed HTTP verbs for the `href`
type CsrPublishHrefHints struct {
	Allow []string `json:"allow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CsrPublishHrefHints CsrPublishHrefHints

// NewCsrPublishHrefHints instantiates a new CsrPublishHrefHints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCsrPublishHrefHints() *CsrPublishHrefHints {
	this := CsrPublishHrefHints{}
	return &this
}

// NewCsrPublishHrefHintsWithDefaults instantiates a new CsrPublishHrefHints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCsrPublishHrefHintsWithDefaults() *CsrPublishHrefHints {
	this := CsrPublishHrefHints{}
	return &this
}

// GetAllow returns the Allow field value if set, zero value otherwise.
func (o *CsrPublishHrefHints) GetAllow() []string {
	if o == nil || o.Allow == nil {
		var ret []string
		return ret
	}
	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsrPublishHrefHints) GetAllowOk() ([]string, bool) {
	if o == nil || o.Allow == nil {
		return nil, false
	}
	return o.Allow, true
}

// HasAllow returns a boolean if a field has been set.
func (o *CsrPublishHrefHints) HasAllow() bool {
	if o != nil && o.Allow != nil {
		return true
	}

	return false
}

// SetAllow gets a reference to the given []string and assigns it to the Allow field.
func (o *CsrPublishHrefHints) SetAllow(v []string) {
	o.Allow = v
}

func (o CsrPublishHrefHints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allow != nil {
		toSerialize["allow"] = o.Allow
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CsrPublishHrefHints) UnmarshalJSON(bytes []byte) (err error) {
	varCsrPublishHrefHints := _CsrPublishHrefHints{}

	err = json.Unmarshal(bytes, &varCsrPublishHrefHints)
	if err == nil {
		*o = CsrPublishHrefHints(varCsrPublishHrefHints)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "allow")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCsrPublishHrefHints struct {
	value *CsrPublishHrefHints
	isSet bool
}

func (v NullableCsrPublishHrefHints) Get() *CsrPublishHrefHints {
	return v.value
}

func (v *NullableCsrPublishHrefHints) Set(val *CsrPublishHrefHints) {
	v.value = val
	v.isSet = true
}

func (v NullableCsrPublishHrefHints) IsSet() bool {
	return v.isSet
}

func (v *NullableCsrPublishHrefHints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCsrPublishHrefHints(val *CsrPublishHrefHints) *NullableCsrPublishHrefHints {
	return &NullableCsrPublishHrefHints{value: val, isSet: true}
}

func (v NullableCsrPublishHrefHints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCsrPublishHrefHints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

