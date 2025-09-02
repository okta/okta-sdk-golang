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

// ResourceConditions Conditions for further restricting a resource.
type ResourceConditions struct {
	Exclude *ResourceConditionsExclude `json:"Exclude,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceConditions ResourceConditions

// NewResourceConditions instantiates a new ResourceConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceConditions() *ResourceConditions {
	this := ResourceConditions{}
	return &this
}

// NewResourceConditionsWithDefaults instantiates a new ResourceConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceConditionsWithDefaults() *ResourceConditions {
	this := ResourceConditions{}
	return &this
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *ResourceConditions) GetExclude() ResourceConditionsExclude {
	if o == nil || o.Exclude == nil {
		var ret ResourceConditionsExclude
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceConditions) GetExcludeOk() (*ResourceConditionsExclude, bool) {
	if o == nil || o.Exclude == nil {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ResourceConditions) HasExclude() bool {
	if o != nil && o.Exclude != nil {
		return true
	}

	return false
}

// SetExclude gets a reference to the given ResourceConditionsExclude and assigns it to the Exclude field.
func (o *ResourceConditions) SetExclude(v ResourceConditionsExclude) {
	o.Exclude = &v
}

func (o ResourceConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Exclude != nil {
		toSerialize["Exclude"] = o.Exclude
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceConditions) UnmarshalJSON(bytes []byte) (err error) {
	varResourceConditions := _ResourceConditions{}

	err = json.Unmarshal(bytes, &varResourceConditions)
	if err == nil {
		*o = ResourceConditions(varResourceConditions)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "Exclude")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableResourceConditions struct {
	value *ResourceConditions
	isSet bool
}

func (v NullableResourceConditions) Get() *ResourceConditions {
	return v.value
}

func (v *NullableResourceConditions) Set(val *ResourceConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceConditions(val *ResourceConditions) *NullableResourceConditions {
	return &NullableResourceConditions{value: val, isSet: true}
}

func (v NullableResourceConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

