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
	"time"
)

// checks if the UserResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserResponseSchema{}

// UserResponseSchema struct for UserResponseSchema
type UserResponseSchema struct {
	// The timestamp when the user was created in the identity source
	Created *time.Time `json:"created,omitempty"`
	// The external ID of the user in the identity source
	ExternalId *string `json:"externalId,omitempty"`
	// The ID of the user in the identity source
	Id *string `json:"id,omitempty"`
	// The timestamp when the user was last updated in the identity source
	LastUpdated          *time.Time                          `json:"lastUpdated,omitempty"`
	Profile              *IdentitySourceUserProfileForUpsert `json:"profile,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserResponseSchema UserResponseSchema

// NewUserResponseSchema instantiates a new UserResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponseSchema() *UserResponseSchema {
	this := UserResponseSchema{}
	return &this
}

// NewUserResponseSchemaWithDefaults instantiates a new UserResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseSchemaWithDefaults() *UserResponseSchema {
	this := UserResponseSchema{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserResponseSchema) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseSchema) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserResponseSchema) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserResponseSchema) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *UserResponseSchema) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseSchema) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *UserResponseSchema) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *UserResponseSchema) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserResponseSchema) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseSchema) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserResponseSchema) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserResponseSchema) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UserResponseSchema) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseSchema) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserResponseSchema) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UserResponseSchema) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserResponseSchema) GetProfile() IdentitySourceUserProfileForUpsert {
	if o == nil || IsNil(o.Profile) {
		var ret IdentitySourceUserProfileForUpsert
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponseSchema) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserResponseSchema) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given IdentitySourceUserProfileForUpsert and assigns it to the Profile field.
func (o *UserResponseSchema) SetProfile(v IdentitySourceUserProfileForUpsert) {
	o.Profile = &v
}

func (o UserResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserResponseSchema) UnmarshalJSON(data []byte) (err error) {
	varUserResponseSchema := _UserResponseSchema{}

	err = json.Unmarshal(data, &varUserResponseSchema)

	if err != nil {
		return err
	}

	*o = UserResponseSchema(varUserResponseSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "profile")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserResponseSchema struct {
	value *UserResponseSchema
	isSet bool
}

func (v NullableUserResponseSchema) Get() *UserResponseSchema {
	return v.value
}

func (v *NullableUserResponseSchema) Set(val *UserResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponseSchema(val *UserResponseSchema) *NullableUserResponseSchema {
	return &NullableUserResponseSchema{value: val, isSet: true}
}

func (v NullableUserResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
