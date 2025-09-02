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
	"fmt"
)

// ResourceSetResourcePostRequest struct for ResourceSetResourcePostRequest
type ResourceSetResourcePostRequest struct {
	Conditions ResourceConditions `json:"conditions"`
	// Resource in ORN or REST API URL format
	ResourceOrnOrUrl string `json:"resourceOrnOrUrl"`
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
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["conditions"] = o.Conditions
	}
	if true {
		toSerialize["resourceOrnOrUrl"] = o.ResourceOrnOrUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceSetResourcePostRequest) UnmarshalJSON(bytes []byte) (err error) {
	varResourceSetResourcePostRequest := _ResourceSetResourcePostRequest{}

	err = json.Unmarshal(bytes, &varResourceSetResourcePostRequest)
	if err == nil {
		*o = ResourceSetResourcePostRequest(varResourceSetResourcePostRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "resourceOrnOrUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

