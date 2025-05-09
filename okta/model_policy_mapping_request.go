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

// PolicyMappingRequest struct for PolicyMappingRequest
type PolicyMappingRequest struct {
	ResourceId *string `json:"resourceId,omitempty"`
	ResourceType *string `json:"resourceType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyMappingRequest PolicyMappingRequest

// NewPolicyMappingRequest instantiates a new PolicyMappingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyMappingRequest() *PolicyMappingRequest {
	this := PolicyMappingRequest{}
	return &this
}

// NewPolicyMappingRequestWithDefaults instantiates a new PolicyMappingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyMappingRequestWithDefaults() *PolicyMappingRequest {
	this := PolicyMappingRequest{}
	return &this
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *PolicyMappingRequest) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingRequest) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *PolicyMappingRequest) HasResourceId() bool {
	if o != nil && o.ResourceId != nil {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *PolicyMappingRequest) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *PolicyMappingRequest) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *PolicyMappingRequest) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *PolicyMappingRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o PolicyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceId != nil {
		toSerialize["resourceId"] = o.ResourceId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PolicyMappingRequest) UnmarshalJSON(bytes []byte) (err error) {
	varPolicyMappingRequest := _PolicyMappingRequest{}

	err = json.Unmarshal(bytes, &varPolicyMappingRequest)
	if err == nil {
		*o = PolicyMappingRequest(varPolicyMappingRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePolicyMappingRequest struct {
	value *PolicyMappingRequest
	isSet bool
}

func (v NullablePolicyMappingRequest) Get() *PolicyMappingRequest {
	return v.value
}

func (v *NullablePolicyMappingRequest) Set(val *PolicyMappingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyMappingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyMappingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyMappingRequest(val *PolicyMappingRequest) *NullablePolicyMappingRequest {
	return &NullablePolicyMappingRequest{value: val, isSet: true}
}

func (v NullablePolicyMappingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyMappingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

