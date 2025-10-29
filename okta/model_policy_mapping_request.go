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
)

// checks if the PolicyMappingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyMappingRequest{}

// PolicyMappingRequest struct for PolicyMappingRequest
type PolicyMappingRequest struct {
	// [Policy ID](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/Policy/#tag/Policy/operation/listPolicies!c=200&path=0/id&t=response) of the app sign-in policy that you want to map
	ResourceId *string `json:"resourceId,omitempty"`
	// Specifies the type of resource to map. You can only map an app sign-in policy to a device signal collection policy (the `policyId` path parameter).
	ResourceType         *string `json:"resourceType,omitempty"`
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
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingRequest) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *PolicyMappingRequest) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
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
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyMappingRequest) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *PolicyMappingRequest) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *PolicyMappingRequest) SetResourceType(v string) {
	o.ResourceType = &v
}

func (o PolicyMappingRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyMappingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyMappingRequest) UnmarshalJSON(data []byte) (err error) {
	varPolicyMappingRequest := _PolicyMappingRequest{}

	err = json.Unmarshal(data, &varPolicyMappingRequest)

	if err != nil {
		return err
	}

	*o = PolicyMappingRequest(varPolicyMappingRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceId")
		delete(additionalProperties, "resourceType")
		o.AdditionalProperties = additionalProperties
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
