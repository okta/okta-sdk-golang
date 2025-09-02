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

// IdentitySourceGroupProfileForUpsert Contains a set of external group attributes and their values that are mapped to Okta standard properties. See the group [`profile` object](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Group/#tag/Group/operation/getGroup!c=200&path=profile&t=response) and Declaration of a Custom Identity Source Schema in [Using anything as a source](https://help.okta.com/okta_help.htm?type=oie&id=ext-anything-as-a-source). > **Note:** Profile attributes can only be of the string type.
type IdentitySourceGroupProfileForUpsert struct {
	// Description of the group
	Description NullableString `json:"description,omitempty"`
	// Name of the group
	DisplayName *string `json:"displayName,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySourceGroupProfileForUpsert IdentitySourceGroupProfileForUpsert

// NewIdentitySourceGroupProfileForUpsert instantiates a new IdentitySourceGroupProfileForUpsert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourceGroupProfileForUpsert() *IdentitySourceGroupProfileForUpsert {
	this := IdentitySourceGroupProfileForUpsert{}
	return &this
}

// NewIdentitySourceGroupProfileForUpsertWithDefaults instantiates a new IdentitySourceGroupProfileForUpsert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourceGroupProfileForUpsertWithDefaults() *IdentitySourceGroupProfileForUpsert {
	this := IdentitySourceGroupProfileForUpsert{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentitySourceGroupProfileForUpsert) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentitySourceGroupProfileForUpsert) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentitySourceGroupProfileForUpsert) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *IdentitySourceGroupProfileForUpsert) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *IdentitySourceGroupProfileForUpsert) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *IdentitySourceGroupProfileForUpsert) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IdentitySourceGroupProfileForUpsert) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceGroupProfileForUpsert) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IdentitySourceGroupProfileForUpsert) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IdentitySourceGroupProfileForUpsert) SetDisplayName(v string) {
	o.DisplayName = &v
}

func (o IdentitySourceGroupProfileForUpsert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentitySourceGroupProfileForUpsert) UnmarshalJSON(bytes []byte) (err error) {
	varIdentitySourceGroupProfileForUpsert := _IdentitySourceGroupProfileForUpsert{}

	err = json.Unmarshal(bytes, &varIdentitySourceGroupProfileForUpsert)
	if err == nil {
		*o = IdentitySourceGroupProfileForUpsert(varIdentitySourceGroupProfileForUpsert)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentitySourceGroupProfileForUpsert struct {
	value *IdentitySourceGroupProfileForUpsert
	isSet bool
}

func (v NullableIdentitySourceGroupProfileForUpsert) Get() *IdentitySourceGroupProfileForUpsert {
	return v.value
}

func (v *NullableIdentitySourceGroupProfileForUpsert) Set(val *IdentitySourceGroupProfileForUpsert) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourceGroupProfileForUpsert) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourceGroupProfileForUpsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourceGroupProfileForUpsert(val *IdentitySourceGroupProfileForUpsert) *NullableIdentitySourceGroupProfileForUpsert {
	return &NullableIdentitySourceGroupProfileForUpsert{value: val, isSet: true}
}

func (v NullableIdentitySourceGroupProfileForUpsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourceGroupProfileForUpsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

