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

// AccessPolicyAllOfEmbedded struct for AccessPolicyAllOfEmbedded
type AccessPolicyAllOfEmbedded struct {
	// The resource that this policy controls. For the [Okta account management policy](https://developer.okta.com/docs/guides/okta-account-management-policy/main/#example-response), `END_USER_ACCOUNT_MANAGEMENT` is returned. For other policies, `APP` is returned.
	ResourceType *string `json:"resourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyAllOfEmbedded AccessPolicyAllOfEmbedded

// NewAccessPolicyAllOfEmbedded instantiates a new AccessPolicyAllOfEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyAllOfEmbedded() *AccessPolicyAllOfEmbedded {
	this := AccessPolicyAllOfEmbedded{}
	return &this
}

// NewAccessPolicyAllOfEmbeddedWithDefaults instantiates a new AccessPolicyAllOfEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyAllOfEmbeddedWithDefaults() *AccessPolicyAllOfEmbedded {
	this := AccessPolicyAllOfEmbedded{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AccessPolicyAllOfEmbedded) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyAllOfEmbedded) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AccessPolicyAllOfEmbedded) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *AccessPolicyAllOfEmbedded) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o AccessPolicyAllOfEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyAllOfEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyAllOfEmbedded := _AccessPolicyAllOfEmbedded{}

	err = json.Unmarshal(bytes, &varAccessPolicyAllOfEmbedded)
	if err == nil {
		*o = AccessPolicyAllOfEmbedded(varAccessPolicyAllOfEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyAllOfEmbedded struct {
	value *AccessPolicyAllOfEmbedded
	isSet bool
}

func (v NullableAccessPolicyAllOfEmbedded) Get() *AccessPolicyAllOfEmbedded {
	return v.value
}

func (v *NullableAccessPolicyAllOfEmbedded) Set(val *AccessPolicyAllOfEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyAllOfEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyAllOfEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyAllOfEmbedded(val *AccessPolicyAllOfEmbedded) *NullableAccessPolicyAllOfEmbedded {
	return &NullableAccessPolicyAllOfEmbedded{value: val, isSet: true}
}

func (v NullableAccessPolicyAllOfEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyAllOfEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

