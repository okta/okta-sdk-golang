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

API version: 2025.08.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the OrgCrossAppAccessConnectionPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCrossAppAccessConnectionPatchRequest{}

// OrgCrossAppAccessConnectionPatchRequest Patch request object for Cross App Access Connections
type OrgCrossAppAccessConnectionPatchRequest struct {
	// Requested value of Cross App Access connection status
	Status               string `json:"status"`
	AdditionalProperties map[string]interface{}
}

type _OrgCrossAppAccessConnectionPatchRequest OrgCrossAppAccessConnectionPatchRequest

// NewOrgCrossAppAccessConnectionPatchRequest instantiates a new OrgCrossAppAccessConnectionPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCrossAppAccessConnectionPatchRequest(status string) *OrgCrossAppAccessConnectionPatchRequest {
	this := OrgCrossAppAccessConnectionPatchRequest{}
	this.Status = status
	return &this
}

// NewOrgCrossAppAccessConnectionPatchRequestWithDefaults instantiates a new OrgCrossAppAccessConnectionPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCrossAppAccessConnectionPatchRequestWithDefaults() *OrgCrossAppAccessConnectionPatchRequest {
	this := OrgCrossAppAccessConnectionPatchRequest{}
	return &this
}

// GetStatus returns the Status field value
func (o *OrgCrossAppAccessConnectionPatchRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnectionPatchRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OrgCrossAppAccessConnectionPatchRequest) SetStatus(v string) {
	o.Status = v
}

func (o OrgCrossAppAccessConnectionPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCrossAppAccessConnectionPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCrossAppAccessConnectionPatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varOrgCrossAppAccessConnectionPatchRequest := _OrgCrossAppAccessConnectionPatchRequest{}

	err = json.Unmarshal(data, &varOrgCrossAppAccessConnectionPatchRequest)

	if err != nil {
		return err
	}

	*o = OrgCrossAppAccessConnectionPatchRequest(varOrgCrossAppAccessConnectionPatchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCrossAppAccessConnectionPatchRequest struct {
	value *OrgCrossAppAccessConnectionPatchRequest
	isSet bool
}

func (v NullableOrgCrossAppAccessConnectionPatchRequest) Get() *OrgCrossAppAccessConnectionPatchRequest {
	return v.value
}

func (v *NullableOrgCrossAppAccessConnectionPatchRequest) Set(val *OrgCrossAppAccessConnectionPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCrossAppAccessConnectionPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCrossAppAccessConnectionPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCrossAppAccessConnectionPatchRequest(val *OrgCrossAppAccessConnectionPatchRequest) *NullableOrgCrossAppAccessConnectionPatchRequest {
	return &NullableOrgCrossAppAccessConnectionPatchRequest{value: val, isSet: true}
}

func (v NullableOrgCrossAppAccessConnectionPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCrossAppAccessConnectionPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
