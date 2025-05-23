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
	"time"
)

// IdentitySourceSession struct for IdentitySourceSession
type IdentitySourceSession struct {
	Created *time.Time `json:"created,omitempty"`
	Id *string `json:"id,omitempty"`
	IdentitySourceId *string `json:"identitySourceId,omitempty"`
	ImportType *string `json:"importType,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentitySourceSession IdentitySourceSession

// NewIdentitySourceSession instantiates a new IdentitySourceSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentitySourceSession() *IdentitySourceSession {
	this := IdentitySourceSession{}
	return &this
}

// NewIdentitySourceSessionWithDefaults instantiates a new IdentitySourceSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentitySourceSessionWithDefaults() *IdentitySourceSession {
	this := IdentitySourceSession{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *IdentitySourceSession) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceSession) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *IdentitySourceSession) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *IdentitySourceSession) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentitySourceSession) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceSession) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentitySourceSession) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentitySourceSession) SetId(v string) {
	o.Id = &v
}

// GetIdentitySourceId returns the IdentitySourceId field value if set, zero value otherwise.
func (o *IdentitySourceSession) GetIdentitySourceId() string {
	if o == nil || o.IdentitySourceId == nil {
		var ret string
		return ret
	}
	return *o.IdentitySourceId
}

// GetIdentitySourceIdOk returns a tuple with the IdentitySourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceSession) GetIdentitySourceIdOk() (*string, bool) {
	if o == nil || o.IdentitySourceId == nil {
		return nil, false
	}
	return o.IdentitySourceId, true
}

// HasIdentitySourceId returns a boolean if a field has been set.
func (o *IdentitySourceSession) HasIdentitySourceId() bool {
	if o != nil && o.IdentitySourceId != nil {
		return true
	}

	return false
}

// SetIdentitySourceId gets a reference to the given string and assigns it to the IdentitySourceId field.
func (o *IdentitySourceSession) SetIdentitySourceId(v string) {
	o.IdentitySourceId = &v
}

// GetImportType returns the ImportType field value if set, zero value otherwise.
func (o *IdentitySourceSession) GetImportType() string {
	if o == nil || o.ImportType == nil {
		var ret string
		return ret
	}
	return *o.ImportType
}

// GetImportTypeOk returns a tuple with the ImportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceSession) GetImportTypeOk() (*string, bool) {
	if o == nil || o.ImportType == nil {
		return nil, false
	}
	return o.ImportType, true
}

// HasImportType returns a boolean if a field has been set.
func (o *IdentitySourceSession) HasImportType() bool {
	if o != nil && o.ImportType != nil {
		return true
	}

	return false
}

// SetImportType gets a reference to the given string and assigns it to the ImportType field.
func (o *IdentitySourceSession) SetImportType(v string) {
	o.ImportType = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *IdentitySourceSession) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceSession) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *IdentitySourceSession) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *IdentitySourceSession) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentitySourceSession) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentitySourceSession) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentitySourceSession) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentitySourceSession) SetStatus(v string) {
	o.Status = &v
}

func (o IdentitySourceSession) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentitySourceId != nil {
		toSerialize["identitySourceId"] = o.IdentitySourceId
	}
	if o.ImportType != nil {
		toSerialize["importType"] = o.ImportType
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentitySourceSession) UnmarshalJSON(bytes []byte) (err error) {
	varIdentitySourceSession := _IdentitySourceSession{}

	err = json.Unmarshal(bytes, &varIdentitySourceSession)
	if err == nil {
		*o = IdentitySourceSession(varIdentitySourceSession)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "identitySourceId")
		delete(additionalProperties, "importType")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentitySourceSession struct {
	value *IdentitySourceSession
	isSet bool
}

func (v NullableIdentitySourceSession) Get() *IdentitySourceSession {
	return v.value
}

func (v *NullableIdentitySourceSession) Set(val *IdentitySourceSession) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentitySourceSession) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentitySourceSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentitySourceSession(val *IdentitySourceSession) *NullableIdentitySourceSession {
	return &NullableIdentitySourceSession{value: val, isSet: true}
}

func (v NullableIdentitySourceSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentitySourceSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

