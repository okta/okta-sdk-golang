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

// LinksCancel struct for LinksCancel
type LinksCancel struct {
	Cancel *LinksCancelCancel `json:"cancel,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksCancel LinksCancel

// NewLinksCancel instantiates a new LinksCancel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksCancel() *LinksCancel {
	this := LinksCancel{}
	return &this
}

// NewLinksCancelWithDefaults instantiates a new LinksCancel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksCancelWithDefaults() *LinksCancel {
	this := LinksCancel{}
	return &this
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *LinksCancel) GetCancel() LinksCancelCancel {
	if o == nil || o.Cancel == nil {
		var ret LinksCancelCancel
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksCancel) GetCancelOk() (*LinksCancelCancel, bool) {
	if o == nil || o.Cancel == nil {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *LinksCancel) HasCancel() bool {
	if o != nil && o.Cancel != nil {
		return true
	}

	return false
}

// SetCancel gets a reference to the given LinksCancelCancel and assigns it to the Cancel field.
func (o *LinksCancel) SetCancel(v LinksCancelCancel) {
	o.Cancel = &v
}

func (o LinksCancel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cancel != nil {
		toSerialize["cancel"] = o.Cancel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksCancel) UnmarshalJSON(bytes []byte) (err error) {
	varLinksCancel := _LinksCancel{}

	err = json.Unmarshal(bytes, &varLinksCancel)
	if err == nil {
		*o = LinksCancel(varLinksCancel)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "cancel")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksCancel struct {
	value *LinksCancel
	isSet bool
}

func (v NullableLinksCancel) Get() *LinksCancel {
	return v.value
}

func (v *NullableLinksCancel) Set(val *LinksCancel) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksCancel) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksCancel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksCancel(val *LinksCancel) *NullableLinksCancel {
	return &NullableLinksCancel{value: val, isSet: true}
}

func (v NullableLinksCancel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksCancel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

