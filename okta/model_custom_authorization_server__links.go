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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the CustomAuthorizationServerLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAuthorizationServerLinks{}

// CustomAuthorizationServerLinks struct for CustomAuthorizationServerLinks
type CustomAuthorizationServerLinks struct {
	Self                 HrefObject  `json:"self"`
	Web                  *HrefObject `json:"web,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomAuthorizationServerLinks CustomAuthorizationServerLinks

// NewCustomAuthorizationServerLinks instantiates a new CustomAuthorizationServerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAuthorizationServerLinks(self HrefObject) *CustomAuthorizationServerLinks {
	this := CustomAuthorizationServerLinks{}
	this.Self = self
	return &this
}

// NewCustomAuthorizationServerLinksWithDefaults instantiates a new CustomAuthorizationServerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAuthorizationServerLinksWithDefaults() *CustomAuthorizationServerLinks {
	this := CustomAuthorizationServerLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *CustomAuthorizationServerLinks) GetSelf() HrefObject {
	if o == nil {
		var ret HrefObject
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServerLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *CustomAuthorizationServerLinks) SetSelf(v HrefObject) {
	o.Self = v
}

// GetWeb returns the Web field value if set, zero value otherwise.
func (o *CustomAuthorizationServerLinks) GetWeb() HrefObject {
	if o == nil || IsNil(o.Web) {
		var ret HrefObject
		return ret
	}
	return *o.Web
}

// GetWebOk returns a tuple with the Web field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServerLinks) GetWebOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Web) {
		return nil, false
	}
	return o.Web, true
}

// HasWeb returns a boolean if a field has been set.
func (o *CustomAuthorizationServerLinks) HasWeb() bool {
	if o != nil && !IsNil(o.Web) {
		return true
	}

	return false
}

// SetWeb gets a reference to the given HrefObject and assigns it to the Web field.
func (o *CustomAuthorizationServerLinks) SetWeb(v HrefObject) {
	o.Web = &v
}

func (o CustomAuthorizationServerLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAuthorizationServerLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	if !IsNil(o.Web) {
		toSerialize["web"] = o.Web
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomAuthorizationServerLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCustomAuthorizationServerLinks := _CustomAuthorizationServerLinks{}

	err = json.Unmarshal(data, &varCustomAuthorizationServerLinks)

	if err != nil {
		return err
	}

	*o = CustomAuthorizationServerLinks(varCustomAuthorizationServerLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "web")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomAuthorizationServerLinks struct {
	value *CustomAuthorizationServerLinks
	isSet bool
}

func (v NullableCustomAuthorizationServerLinks) Get() *CustomAuthorizationServerLinks {
	return v.value
}

func (v *NullableCustomAuthorizationServerLinks) Set(val *CustomAuthorizationServerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAuthorizationServerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAuthorizationServerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAuthorizationServerLinks(val *CustomAuthorizationServerLinks) *NullableCustomAuthorizationServerLinks {
	return &NullableCustomAuthorizationServerLinks{value: val, isSet: true}
}

func (v NullableCustomAuthorizationServerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAuthorizationServerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
