/*
Okta Admin Management API

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

// checks if the OktaManagedUserAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaManagedUserAccountRequest{}

// OktaManagedUserAccountRequest Request body for creating an Okta managed user account
type OktaManagedUserAccountRequest struct {
	// The description of the Okta managed user account
	Description *string `json:"description,omitempty"`
	// The user-defined name for the Okta managed user account
	Name string `json:"name" validate:"regexp=^[\\\\w\\\\-_. ]+$"`
	// The ID of the Okta user to manage as a service account. This must be an existing user in your Okta org.
	OktaUserId string `json:"oktaUserId" validate:"regexp=^[a-zA-Z0-9]+$"`
	// A list of IDs of the Okta groups who own the Okta managed user account
	OwnerGroupIds []string `json:"ownerGroupIds,omitempty"`
	// A list of IDs of the Okta users who own the Okta managed user account
	OwnerUserIds         []string `json:"ownerUserIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaManagedUserAccountRequest OktaManagedUserAccountRequest

// NewOktaManagedUserAccountRequest instantiates a new OktaManagedUserAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaManagedUserAccountRequest(name string, oktaUserId string) *OktaManagedUserAccountRequest {
	this := OktaManagedUserAccountRequest{}
	this.Name = name
	this.OktaUserId = oktaUserId
	return &this
}

// NewOktaManagedUserAccountRequestWithDefaults instantiates a new OktaManagedUserAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaManagedUserAccountRequestWithDefaults() *OktaManagedUserAccountRequest {
	this := OktaManagedUserAccountRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OktaManagedUserAccountRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OktaManagedUserAccountRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OktaManagedUserAccountRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *OktaManagedUserAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OktaManagedUserAccountRequest) SetName(v string) {
	o.Name = v
}

// GetOktaUserId returns the OktaUserId field value
func (o *OktaManagedUserAccountRequest) GetOktaUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaUserId
}

// GetOktaUserIdOk returns a tuple with the OktaUserId field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountRequest) GetOktaUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaUserId, true
}

// SetOktaUserId sets field value
func (o *OktaManagedUserAccountRequest) SetOktaUserId(v string) {
	o.OktaUserId = v
}

// GetOwnerGroupIds returns the OwnerGroupIds field value if set, zero value otherwise.
func (o *OktaManagedUserAccountRequest) GetOwnerGroupIds() []string {
	if o == nil || IsNil(o.OwnerGroupIds) {
		var ret []string
		return ret
	}
	return o.OwnerGroupIds
}

// GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountRequest) GetOwnerGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerGroupIds) {
		return nil, false
	}
	return o.OwnerGroupIds, true
}

// HasOwnerGroupIds returns a boolean if a field has been set.
func (o *OktaManagedUserAccountRequest) HasOwnerGroupIds() bool {
	if o != nil && !IsNil(o.OwnerGroupIds) {
		return true
	}

	return false
}

// SetOwnerGroupIds gets a reference to the given []string and assigns it to the OwnerGroupIds field.
func (o *OktaManagedUserAccountRequest) SetOwnerGroupIds(v []string) {
	o.OwnerGroupIds = v
}

// GetOwnerUserIds returns the OwnerUserIds field value if set, zero value otherwise.
func (o *OktaManagedUserAccountRequest) GetOwnerUserIds() []string {
	if o == nil || IsNil(o.OwnerUserIds) {
		var ret []string
		return ret
	}
	return o.OwnerUserIds
}

// GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountRequest) GetOwnerUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerUserIds) {
		return nil, false
	}
	return o.OwnerUserIds, true
}

// HasOwnerUserIds returns a boolean if a field has been set.
func (o *OktaManagedUserAccountRequest) HasOwnerUserIds() bool {
	if o != nil && !IsNil(o.OwnerUserIds) {
		return true
	}

	return false
}

// SetOwnerUserIds gets a reference to the given []string and assigns it to the OwnerUserIds field.
func (o *OktaManagedUserAccountRequest) SetOwnerUserIds(v []string) {
	o.OwnerUserIds = v
}

func (o OktaManagedUserAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaManagedUserAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["name"] = o.Name
	toSerialize["oktaUserId"] = o.OktaUserId
	if !IsNil(o.OwnerGroupIds) {
		toSerialize["ownerGroupIds"] = o.OwnerGroupIds
	}
	if !IsNil(o.OwnerUserIds) {
		toSerialize["ownerUserIds"] = o.OwnerUserIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaManagedUserAccountRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"oktaUserId",
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

	varOktaManagedUserAccountRequest := _OktaManagedUserAccountRequest{}

	err = json.Unmarshal(data, &varOktaManagedUserAccountRequest)

	if err != nil {
		return err
	}

	*o = OktaManagedUserAccountRequest(varOktaManagedUserAccountRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "oktaUserId")
		delete(additionalProperties, "ownerGroupIds")
		delete(additionalProperties, "ownerUserIds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaManagedUserAccountRequest struct {
	value *OktaManagedUserAccountRequest
	isSet bool
}

func (v NullableOktaManagedUserAccountRequest) Get() *OktaManagedUserAccountRequest {
	return v.value
}

func (v *NullableOktaManagedUserAccountRequest) Set(val *OktaManagedUserAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaManagedUserAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaManagedUserAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaManagedUserAccountRequest(val *OktaManagedUserAccountRequest) *NullableOktaManagedUserAccountRequest {
	return &NullableOktaManagedUserAccountRequest{value: val, isSet: true}
}

func (v NullableOktaManagedUserAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaManagedUserAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
