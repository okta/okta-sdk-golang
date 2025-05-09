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

// BouncesRemoveListObj struct for BouncesRemoveListObj
type BouncesRemoveListObj struct {
	EmailAddresses []string `json:"emailAddresses,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BouncesRemoveListObj BouncesRemoveListObj

// NewBouncesRemoveListObj instantiates a new BouncesRemoveListObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBouncesRemoveListObj() *BouncesRemoveListObj {
	this := BouncesRemoveListObj{}
	return &this
}

// NewBouncesRemoveListObjWithDefaults instantiates a new BouncesRemoveListObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBouncesRemoveListObjWithDefaults() *BouncesRemoveListObj {
	this := BouncesRemoveListObj{}
	return &this
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *BouncesRemoveListObj) GetEmailAddresses() []string {
	if o == nil || o.EmailAddresses == nil {
		var ret []string
		return ret
	}
	return o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BouncesRemoveListObj) GetEmailAddressesOk() ([]string, bool) {
	if o == nil || o.EmailAddresses == nil {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *BouncesRemoveListObj) HasEmailAddresses() bool {
	if o != nil && o.EmailAddresses != nil {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []string and assigns it to the EmailAddresses field.
func (o *BouncesRemoveListObj) SetEmailAddresses(v []string) {
	o.EmailAddresses = v
}

func (o BouncesRemoveListObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddresses != nil {
		toSerialize["emailAddresses"] = o.EmailAddresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BouncesRemoveListObj) UnmarshalJSON(bytes []byte) (err error) {
	varBouncesRemoveListObj := _BouncesRemoveListObj{}

	err = json.Unmarshal(bytes, &varBouncesRemoveListObj)
	if err == nil {
		*o = BouncesRemoveListObj(varBouncesRemoveListObj)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "emailAddresses")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBouncesRemoveListObj struct {
	value *BouncesRemoveListObj
	isSet bool
}

func (v NullableBouncesRemoveListObj) Get() *BouncesRemoveListObj {
	return v.value
}

func (v *NullableBouncesRemoveListObj) Set(val *BouncesRemoveListObj) {
	v.value = val
	v.isSet = true
}

func (v NullableBouncesRemoveListObj) IsSet() bool {
	return v.isSet
}

func (v *NullableBouncesRemoveListObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBouncesRemoveListObj(val *BouncesRemoveListObj) *NullableBouncesRemoveListObj {
	return &NullableBouncesRemoveListObj{value: val, isSet: true}
}

func (v NullableBouncesRemoveListObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBouncesRemoveListObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

