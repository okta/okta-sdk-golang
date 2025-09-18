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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the OrgCrossAppAccessConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgCrossAppAccessConnection{}

// OrgCrossAppAccessConnection Connection object for Cross App Access connections
type OrgCrossAppAccessConnection struct {
	// The ISO 8601 formatted date and time when the connection was created
	Created *time.Time `json:"created,omitempty"`
	// Unique identifier for the connection
	Id *string `json:"id,omitempty"`
	// The ISO 8601 formatted date and time when the connection was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// ID of the requesting app instance
	RequestingAppInstanceId *string `json:"requestingAppInstanceId,omitempty"`
	// ID of the resource app instance
	ResourceAppInstanceId *string `json:"resourceAppInstanceId,omitempty"`
	// Indicates if the Cross App Access connection is active or inactive
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgCrossAppAccessConnection OrgCrossAppAccessConnection

// NewOrgCrossAppAccessConnection instantiates a new OrgCrossAppAccessConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgCrossAppAccessConnection() *OrgCrossAppAccessConnection {
	this := OrgCrossAppAccessConnection{}
	return &this
}

// NewOrgCrossAppAccessConnectionWithDefaults instantiates a new OrgCrossAppAccessConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgCrossAppAccessConnectionWithDefaults() *OrgCrossAppAccessConnection {
	this := OrgCrossAppAccessConnection{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrgCrossAppAccessConnection) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnection) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrgCrossAppAccessConnection) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OrgCrossAppAccessConnection) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgCrossAppAccessConnection) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnection) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgCrossAppAccessConnection) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgCrossAppAccessConnection) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OrgCrossAppAccessConnection) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnection) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OrgCrossAppAccessConnection) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OrgCrossAppAccessConnection) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetRequestingAppInstanceId returns the RequestingAppInstanceId field value if set, zero value otherwise.
func (o *OrgCrossAppAccessConnection) GetRequestingAppInstanceId() string {
	if o == nil || IsNil(o.RequestingAppInstanceId) {
		var ret string
		return ret
	}
	return *o.RequestingAppInstanceId
}

// GetRequestingAppInstanceIdOk returns a tuple with the RequestingAppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnection) GetRequestingAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestingAppInstanceId) {
		return nil, false
	}
	return o.RequestingAppInstanceId, true
}

// HasRequestingAppInstanceId returns a boolean if a field has been set.
func (o *OrgCrossAppAccessConnection) HasRequestingAppInstanceId() bool {
	if o != nil && !IsNil(o.RequestingAppInstanceId) {
		return true
	}

	return false
}

// SetRequestingAppInstanceId gets a reference to the given string and assigns it to the RequestingAppInstanceId field.
func (o *OrgCrossAppAccessConnection) SetRequestingAppInstanceId(v string) {
	o.RequestingAppInstanceId = &v
}

// GetResourceAppInstanceId returns the ResourceAppInstanceId field value if set, zero value otherwise.
func (o *OrgCrossAppAccessConnection) GetResourceAppInstanceId() string {
	if o == nil || IsNil(o.ResourceAppInstanceId) {
		var ret string
		return ret
	}
	return *o.ResourceAppInstanceId
}

// GetResourceAppInstanceIdOk returns a tuple with the ResourceAppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnection) GetResourceAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceAppInstanceId) {
		return nil, false
	}
	return o.ResourceAppInstanceId, true
}

// HasResourceAppInstanceId returns a boolean if a field has been set.
func (o *OrgCrossAppAccessConnection) HasResourceAppInstanceId() bool {
	if o != nil && !IsNil(o.ResourceAppInstanceId) {
		return true
	}

	return false
}

// SetResourceAppInstanceId gets a reference to the given string and assigns it to the ResourceAppInstanceId field.
func (o *OrgCrossAppAccessConnection) SetResourceAppInstanceId(v string) {
	o.ResourceAppInstanceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrgCrossAppAccessConnection) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgCrossAppAccessConnection) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrgCrossAppAccessConnection) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrgCrossAppAccessConnection) SetStatus(v string) {
	o.Status = &v
}

func (o OrgCrossAppAccessConnection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgCrossAppAccessConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.RequestingAppInstanceId) {
		toSerialize["requestingAppInstanceId"] = o.RequestingAppInstanceId
	}
	if !IsNil(o.ResourceAppInstanceId) {
		toSerialize["resourceAppInstanceId"] = o.ResourceAppInstanceId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgCrossAppAccessConnection) UnmarshalJSON(data []byte) (err error) {
	varOrgCrossAppAccessConnection := _OrgCrossAppAccessConnection{}

	err = json.Unmarshal(data, &varOrgCrossAppAccessConnection)

	if err != nil {
		return err
	}

	*o = OrgCrossAppAccessConnection(varOrgCrossAppAccessConnection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "requestingAppInstanceId")
		delete(additionalProperties, "resourceAppInstanceId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgCrossAppAccessConnection struct {
	value *OrgCrossAppAccessConnection
	isSet bool
}

func (v NullableOrgCrossAppAccessConnection) Get() *OrgCrossAppAccessConnection {
	return v.value
}

func (v *NullableOrgCrossAppAccessConnection) Set(val *OrgCrossAppAccessConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgCrossAppAccessConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgCrossAppAccessConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgCrossAppAccessConnection(val *OrgCrossAppAccessConnection) *NullableOrgCrossAppAccessConnection {
	return &NullableOrgCrossAppAccessConnection{value: val, isSet: true}
}

func (v NullableOrgCrossAppAccessConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgCrossAppAccessConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
