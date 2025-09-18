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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the ResourceSetResourcePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceSetResourcePostRequest{}

// ResourceSetResourcePostRequest struct for ResourceSetResourcePostRequest
type ResourceSetResourcePostRequest struct {
	Conditions ResourceConditions `json:"conditions"`
	// Resource in ORN or REST API URL format
	ResourceOrnOrUrl     string `json:"resourceOrnOrUrl"`
	AdditionalProperties map[string]interface{}
}

type _ResourceSetResourcePostRequest ResourceSetResourcePostRequest

// NewResourceSetResourcePostRequest instantiates a new ResourceSetResourcePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSetResourcePostRequest(conditions ResourceConditions, resourceOrnOrUrl string) *ResourceSetResourcePostRequest {
	this := ResourceSetResourcePostRequest{}
	this.Conditions = conditions
	this.ResourceOrnOrUrl = resourceOrnOrUrl
	return &this
}

// NewResourceSetResourcePostRequestWithDefaults instantiates a new ResourceSetResourcePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSetResourcePostRequestWithDefaults() *ResourceSetResourcePostRequest {
	this := ResourceSetResourcePostRequest{}
	return &this
}

// GetConditions returns the Conditions field value
func (o *ResourceSetResourcePostRequest) GetConditions() ResourceConditions {
	if o == nil {
		var ret ResourceConditions
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *ResourceSetResourcePostRequest) GetConditionsOk() (*ResourceConditions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value
func (o *ResourceSetResourcePostRequest) SetConditions(v ResourceConditions) {
	o.Conditions = v
}

// GetResourceOrnOrUrl returns the ResourceOrnOrUrl field value
func (o *ResourceSetResourcePostRequest) GetResourceOrnOrUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrnOrUrl
}

// GetResourceOrnOrUrlOk returns a tuple with the ResourceOrnOrUrl field value
// and a boolean to check if the value has been set.
func (o *ResourceSetResourcePostRequest) GetResourceOrnOrUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrnOrUrl, true
}

// SetResourceOrnOrUrl sets field value
func (o *ResourceSetResourcePostRequest) SetResourceOrnOrUrl(v string) {
	o.ResourceOrnOrUrl = v
}

func (o ResourceSetResourcePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSetResourcePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditions"] = o.Conditions
	toSerialize["resourceOrnOrUrl"] = o.ResourceOrnOrUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceSetResourcePostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conditions",
		"resourceOrnOrUrl",
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

	varResourceSetResourcePostRequest := _ResourceSetResourcePostRequest{}

	err = json.Unmarshal(data, &varResourceSetResourcePostRequest)

	if err != nil {
		return err
	}

	*o = ResourceSetResourcePostRequest(varResourceSetResourcePostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "resourceOrnOrUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceSetResourcePostRequest struct {
	value *ResourceSetResourcePostRequest
	isSet bool
}

func (v NullableResourceSetResourcePostRequest) Get() *ResourceSetResourcePostRequest {
	return v.value
}

func (v *NullableResourceSetResourcePostRequest) Set(val *ResourceSetResourcePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSetResourcePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSetResourcePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSetResourcePostRequest(val *ResourceSetResourcePostRequest) *NullableResourceSetResourcePostRequest {
	return &NullableResourceSetResourcePostRequest{value: val, isSet: true}
}

func (v NullableResourceSetResourcePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSetResourcePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
