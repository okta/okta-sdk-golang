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

// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	Credentials *UserCredentials `json:"credentials,omitempty"`
	GroupIds []string `json:"groupIds,omitempty"`
	Profile UserProfile `json:"profile"`
	// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle></div>The ID of the Realm in which the user is residing
	RealmId *string `json:"realmId,omitempty"`
	Type *CreateUserRequestType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateUserRequest CreateUserRequest

// NewCreateUserRequest instantiates a new CreateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserRequest(profile UserProfile) *CreateUserRequest {
	this := CreateUserRequest{}
	this.Profile = profile
	return &this
}

// NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserRequestWithDefaults() *CreateUserRequest {
	this := CreateUserRequest{}
	return &this
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *CreateUserRequest) GetCredentials() UserCredentials {
	if o == nil || o.Credentials == nil {
		var ret UserCredentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetCredentialsOk() (*UserCredentials, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *CreateUserRequest) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given UserCredentials and assigns it to the Credentials field.
func (o *CreateUserRequest) SetCredentials(v UserCredentials) {
	o.Credentials = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *CreateUserRequest) GetGroupIds() []string {
	if o == nil || o.GroupIds == nil {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetGroupIdsOk() ([]string, bool) {
	if o == nil || o.GroupIds == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *CreateUserRequest) HasGroupIds() bool {
	if o != nil && o.GroupIds != nil {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *CreateUserRequest) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetProfile returns the Profile field value
func (o *CreateUserRequest) GetProfile() UserProfile {
	if o == nil {
		var ret UserProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetProfileOk() (*UserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *CreateUserRequest) SetProfile(v UserProfile) {
	o.Profile = v
}

// GetRealmId returns the RealmId field value if set, zero value otherwise.
func (o *CreateUserRequest) GetRealmId() string {
	if o == nil || o.RealmId == nil {
		var ret string
		return ret
	}
	return *o.RealmId
}

// GetRealmIdOk returns a tuple with the RealmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetRealmIdOk() (*string, bool) {
	if o == nil || o.RealmId == nil {
		return nil, false
	}
	return o.RealmId, true
}

// HasRealmId returns a boolean if a field has been set.
func (o *CreateUserRequest) HasRealmId() bool {
	if o != nil && o.RealmId != nil {
		return true
	}

	return false
}

// SetRealmId gets a reference to the given string and assigns it to the RealmId field.
func (o *CreateUserRequest) SetRealmId(v string) {
	o.RealmId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateUserRequest) GetType() CreateUserRequestType {
	if o == nil || o.Type == nil {
		var ret CreateUserRequestType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUserRequest) GetTypeOk() (*CreateUserRequestType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateUserRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given CreateUserRequestType and assigns it to the Type field.
func (o *CreateUserRequest) SetType(v CreateUserRequestType) {
	o.Type = &v
}

func (o CreateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.GroupIds != nil {
		toSerialize["groupIds"] = o.GroupIds
	}
	if true {
		toSerialize["profile"] = o.Profile
	}
	if o.RealmId != nil {
		toSerialize["realmId"] = o.RealmId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateUserRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateUserRequest := _CreateUserRequest{}

	err = json.Unmarshal(bytes, &varCreateUserRequest)
	if err == nil {
		*o = CreateUserRequest(varCreateUserRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "groupIds")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "realmId")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCreateUserRequest struct {
	value *CreateUserRequest
	isSet bool
}

func (v NullableCreateUserRequest) Get() *CreateUserRequest {
	return v.value
}

func (v *NullableCreateUserRequest) Set(val *CreateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserRequest(val *CreateUserRequest) *NullableCreateUserRequest {
	return &NullableCreateUserRequest{value: val, isSet: true}
}

func (v NullableCreateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

