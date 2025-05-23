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

// GovernanceBundleCreateRequest struct for GovernanceBundleCreateRequest
type GovernanceBundleCreateRequest struct {
	Description *string `json:"description,omitempty"`
	Entitlements []IAMBundleEntitlement `json:"entitlements,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GovernanceBundleCreateRequest GovernanceBundleCreateRequest

// NewGovernanceBundleCreateRequest instantiates a new GovernanceBundleCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGovernanceBundleCreateRequest() *GovernanceBundleCreateRequest {
	this := GovernanceBundleCreateRequest{}
	return &this
}

// NewGovernanceBundleCreateRequestWithDefaults instantiates a new GovernanceBundleCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGovernanceBundleCreateRequestWithDefaults() *GovernanceBundleCreateRequest {
	this := GovernanceBundleCreateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GovernanceBundleCreateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundleCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GovernanceBundleCreateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GovernanceBundleCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *GovernanceBundleCreateRequest) GetEntitlements() []IAMBundleEntitlement {
	if o == nil || o.Entitlements == nil {
		var ret []IAMBundleEntitlement
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundleCreateRequest) GetEntitlementsOk() ([]IAMBundleEntitlement, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *GovernanceBundleCreateRequest) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []IAMBundleEntitlement and assigns it to the Entitlements field.
func (o *GovernanceBundleCreateRequest) SetEntitlements(v []IAMBundleEntitlement) {
	o.Entitlements = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GovernanceBundleCreateRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceBundleCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GovernanceBundleCreateRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GovernanceBundleCreateRequest) SetName(v string) {
	o.Name = &v
}

func (o GovernanceBundleCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *GovernanceBundleCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varGovernanceBundleCreateRequest := _GovernanceBundleCreateRequest{}

	err = json.Unmarshal(bytes, &varGovernanceBundleCreateRequest)
	if err == nil {
		*o = GovernanceBundleCreateRequest(varGovernanceBundleCreateRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableGovernanceBundleCreateRequest struct {
	value *GovernanceBundleCreateRequest
	isSet bool
}

func (v NullableGovernanceBundleCreateRequest) Get() *GovernanceBundleCreateRequest {
	return v.value
}

func (v *NullableGovernanceBundleCreateRequest) Set(val *GovernanceBundleCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGovernanceBundleCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGovernanceBundleCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGovernanceBundleCreateRequest(val *GovernanceBundleCreateRequest) *NullableGovernanceBundleCreateRequest {
	return &NullableGovernanceBundleCreateRequest{value: val, isSet: true}
}

func (v NullableGovernanceBundleCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGovernanceBundleCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

