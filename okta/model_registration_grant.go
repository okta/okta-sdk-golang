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
	"time"
)

// checks if the RegistrationGrant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationGrant{}

// RegistrationGrant struct for RegistrationGrant
type RegistrationGrant struct {
	// Timestamp when the grant was consumed by successful device registration. `null` while `PENDING`.
	ConsumedAt NullableTime `json:"consumedAt,omitempty"`
	// Timestamp when the grant was created
	Created *time.Time `json:"created,omitempty"`
	// Type of registration grant
	GrantType *string `json:"grantType,omitempty"`
	// Unique identifier for the registration grant (prefix `dpr`)
	Id *string `json:"id,omitempty"`
	// ID of the registration secret used for this grant
	RegistrationSecretId *string `json:"registrationSecretId,omitempty"`
	// Status of the registration grant
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationGrant RegistrationGrant

// NewRegistrationGrant instantiates a new RegistrationGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationGrant() *RegistrationGrant {
	this := RegistrationGrant{}
	return &this
}

// NewRegistrationGrantWithDefaults instantiates a new RegistrationGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationGrantWithDefaults() *RegistrationGrant {
	this := RegistrationGrant{}
	return &this
}

// GetConsumedAt returns the ConsumedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegistrationGrant) GetConsumedAt() time.Time {
	if o == nil || IsNil(o.ConsumedAt.Get()) {
		var ret time.Time
		return ret
	}
	return *o.ConsumedAt.Get()
}

// GetConsumedAtOk returns a tuple with the ConsumedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegistrationGrant) GetConsumedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConsumedAt.Get(), o.ConsumedAt.IsSet()
}

// HasConsumedAt returns a boolean if a field has been set.
func (o *RegistrationGrant) HasConsumedAt() bool {
	if o != nil && o.ConsumedAt.IsSet() {
		return true
	}

	return false
}

// SetConsumedAt gets a reference to the given NullableTime and assigns it to the ConsumedAt field.
func (o *RegistrationGrant) SetConsumedAt(v time.Time) {
	o.ConsumedAt.Set(&v)
}

// SetConsumedAtNil sets the value for ConsumedAt to be an explicit nil
func (o *RegistrationGrant) SetConsumedAtNil() {
	o.ConsumedAt.Set(nil)
}

// UnsetConsumedAt ensures that no value is present for ConsumedAt, not even an explicit nil
func (o *RegistrationGrant) UnsetConsumedAt() {
	o.ConsumedAt.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RegistrationGrant) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationGrant) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RegistrationGrant) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RegistrationGrant) SetCreated(v time.Time) {
	o.Created = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *RegistrationGrant) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationGrant) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *RegistrationGrant) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *RegistrationGrant) SetGrantType(v string) {
	o.GrantType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrationGrant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationGrant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrationGrant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrationGrant) SetId(v string) {
	o.Id = &v
}

// GetRegistrationSecretId returns the RegistrationSecretId field value if set, zero value otherwise.
func (o *RegistrationGrant) GetRegistrationSecretId() string {
	if o == nil || IsNil(o.RegistrationSecretId) {
		var ret string
		return ret
	}
	return *o.RegistrationSecretId
}

// GetRegistrationSecretIdOk returns a tuple with the RegistrationSecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationGrant) GetRegistrationSecretIdOk() (*string, bool) {
	if o == nil || IsNil(o.RegistrationSecretId) {
		return nil, false
	}
	return o.RegistrationSecretId, true
}

// HasRegistrationSecretId returns a boolean if a field has been set.
func (o *RegistrationGrant) HasRegistrationSecretId() bool {
	if o != nil && !IsNil(o.RegistrationSecretId) {
		return true
	}

	return false
}

// SetRegistrationSecretId gets a reference to the given string and assigns it to the RegistrationSecretId field.
func (o *RegistrationGrant) SetRegistrationSecretId(v string) {
	o.RegistrationSecretId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegistrationGrant) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationGrant) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegistrationGrant) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RegistrationGrant) SetStatus(v string) {
	o.Status = &v
}

func (o RegistrationGrant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationGrant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConsumedAt.IsSet() {
		toSerialize["consumedAt"] = o.ConsumedAt.Get()
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.GrantType) {
		toSerialize["grantType"] = o.GrantType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RegistrationSecretId) {
		toSerialize["registrationSecretId"] = o.RegistrationSecretId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationGrant) UnmarshalJSON(data []byte) (err error) {
	varRegistrationGrant := _RegistrationGrant{}

	err = json.Unmarshal(data, &varRegistrationGrant)

	if err != nil {
		return err
	}

	*o = RegistrationGrant(varRegistrationGrant)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "consumedAt")
		delete(additionalProperties, "created")
		delete(additionalProperties, "grantType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "registrationSecretId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationGrant struct {
	value *RegistrationGrant
	isSet bool
}

func (v NullableRegistrationGrant) Get() *RegistrationGrant {
	return v.value
}

func (v *NullableRegistrationGrant) Set(val *RegistrationGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationGrant(val *RegistrationGrant) *NullableRegistrationGrant {
	return &NullableRegistrationGrant{value: val, isSet: true}
}

func (v NullableRegistrationGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
