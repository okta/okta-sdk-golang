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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
	"fmt"
)

// ServiceAccount struct for ServiceAccount
type ServiceAccount struct {
	// The type of service account
	AccountType string `json:"accountType"`
	// Timestamp when the service account was created
	Created *time.Time `json:"created,omitempty"`
	// The description of the service account
	Description *string `json:"description,omitempty"`
	// The UUID of the service account
	Id *string `json:"id,omitempty" validate:"regexp=(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"`
	// Timestamp when the service account was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The user-defined name for the service account
	Name string `json:"name" validate:"regexp=^[\\\\w\\\\-_. ]+$"`
	// A list of IDs of the Okta groups that own the service account
	OwnerGroupIds []string `json:"ownerGroupIds,omitempty"`
	// A list of IDs of the Okta users that own the service account
	OwnerUserIds []string `json:"ownerUserIds,omitempty"`
	// Describes the current status of an app service account
	Status *string `json:"status,omitempty"`
	// Describes the detailed status of an app service account
	StatusDetail *string `json:"statusDetail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccount ServiceAccount

// NewServiceAccount instantiates a new ServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccount(accountType string, name string) *ServiceAccount {
	this := ServiceAccount{}
	this.AccountType = accountType
	this.Name = name
	return &this
}

// NewServiceAccountWithDefaults instantiates a new ServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountWithDefaults() *ServiceAccount {
	this := ServiceAccount{}
	return &this
}

// GetAccountType returns the AccountType field value
func (o *ServiceAccount) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *ServiceAccount) SetAccountType(v string) {
	o.AccountType = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServiceAccount) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServiceAccount) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ServiceAccount) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceAccount) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceAccount) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ServiceAccount) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ServiceAccount) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ServiceAccount) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *ServiceAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceAccount) SetName(v string) {
	o.Name = v
}

// GetOwnerGroupIds returns the OwnerGroupIds field value if set, zero value otherwise.
func (o *ServiceAccount) GetOwnerGroupIds() []string {
	if o == nil || o.OwnerGroupIds == nil {
		var ret []string
		return ret
	}
	return o.OwnerGroupIds
}

// GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetOwnerGroupIdsOk() ([]string, bool) {
	if o == nil || o.OwnerGroupIds == nil {
		return nil, false
	}
	return o.OwnerGroupIds, true
}

// HasOwnerGroupIds returns a boolean if a field has been set.
func (o *ServiceAccount) HasOwnerGroupIds() bool {
	if o != nil && o.OwnerGroupIds != nil {
		return true
	}

	return false
}

// SetOwnerGroupIds gets a reference to the given []string and assigns it to the OwnerGroupIds field.
func (o *ServiceAccount) SetOwnerGroupIds(v []string) {
	o.OwnerGroupIds = v
}

// GetOwnerUserIds returns the OwnerUserIds field value if set, zero value otherwise.
func (o *ServiceAccount) GetOwnerUserIds() []string {
	if o == nil || o.OwnerUserIds == nil {
		var ret []string
		return ret
	}
	return o.OwnerUserIds
}

// GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetOwnerUserIdsOk() ([]string, bool) {
	if o == nil || o.OwnerUserIds == nil {
		return nil, false
	}
	return o.OwnerUserIds, true
}

// HasOwnerUserIds returns a boolean if a field has been set.
func (o *ServiceAccount) HasOwnerUserIds() bool {
	if o != nil && o.OwnerUserIds != nil {
		return true
	}

	return false
}

// SetOwnerUserIds gets a reference to the given []string and assigns it to the OwnerUserIds field.
func (o *ServiceAccount) SetOwnerUserIds(v []string) {
	o.OwnerUserIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServiceAccount) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServiceAccount) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServiceAccount) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *ServiceAccount) GetStatusDetail() string {
	if o == nil || o.StatusDetail == nil {
		var ret string
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccount) GetStatusDetailOk() (*string, bool) {
	if o == nil || o.StatusDetail == nil {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *ServiceAccount) HasStatusDetail() bool {
	if o != nil && o.StatusDetail != nil {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given string and assigns it to the StatusDetail field.
func (o *ServiceAccount) SetStatusDetail(v string) {
	o.StatusDetail = &v
}

func (o ServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountType"] = o.AccountType
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OwnerGroupIds != nil {
		toSerialize["ownerGroupIds"] = o.OwnerGroupIds
	}
	if o.OwnerUserIds != nil {
		toSerialize["ownerUserIds"] = o.OwnerUserIds
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetail != nil {
		toSerialize["statusDetail"] = o.StatusDetail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceAccount) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAccount := _ServiceAccount{}

	err = json.Unmarshal(bytes, &varServiceAccount)
	if err == nil {
		*o = ServiceAccount(varServiceAccount)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountType")
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ownerGroupIds")
		delete(additionalProperties, "ownerUserIds")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetail")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableServiceAccount struct {
	value *ServiceAccount
	isSet bool
}

func (v NullableServiceAccount) Get() *ServiceAccount {
	return v.value
}

func (v *NullableServiceAccount) Set(val *ServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccount(val *ServiceAccount) *NullableServiceAccount {
	return &NullableServiceAccount{value: val, isSet: true}
}

func (v NullableServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

