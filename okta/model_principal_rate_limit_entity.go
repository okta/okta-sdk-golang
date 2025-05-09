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

// PrincipalRateLimitEntity 
type PrincipalRateLimitEntity struct {
	CreatedBy *string `json:"createdBy,omitempty"`
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	DefaultConcurrencyPercentage *int32 `json:"defaultConcurrencyPercentage,omitempty"`
	DefaultPercentage *int32 `json:"defaultPercentage,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdate *time.Time `json:"lastUpdate,omitempty"`
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	OrgId *string `json:"orgId,omitempty"`
	PrincipalId string `json:"principalId"`
	PrincipalType string `json:"principalType"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalRateLimitEntity PrincipalRateLimitEntity

// NewPrincipalRateLimitEntity instantiates a new PrincipalRateLimitEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalRateLimitEntity(principalId string, principalType string) *PrincipalRateLimitEntity {
	this := PrincipalRateLimitEntity{}
	this.PrincipalId = principalId
	this.PrincipalType = principalType
	return &this
}

// NewPrincipalRateLimitEntityWithDefaults instantiates a new PrincipalRateLimitEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalRateLimitEntityWithDefaults() *PrincipalRateLimitEntity {
	this := PrincipalRateLimitEntity{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *PrincipalRateLimitEntity) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetCreatedDate() time.Time {
	if o == nil || o.CreatedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || o.CreatedDate == nil {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *PrincipalRateLimitEntity) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDefaultConcurrencyPercentage returns the DefaultConcurrencyPercentage field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetDefaultConcurrencyPercentage() int32 {
	if o == nil || o.DefaultConcurrencyPercentage == nil {
		var ret int32
		return ret
	}
	return *o.DefaultConcurrencyPercentage
}

// GetDefaultConcurrencyPercentageOk returns a tuple with the DefaultConcurrencyPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetDefaultConcurrencyPercentageOk() (*int32, bool) {
	if o == nil || o.DefaultConcurrencyPercentage == nil {
		return nil, false
	}
	return o.DefaultConcurrencyPercentage, true
}

// HasDefaultConcurrencyPercentage returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasDefaultConcurrencyPercentage() bool {
	if o != nil && o.DefaultConcurrencyPercentage != nil {
		return true
	}

	return false
}

// SetDefaultConcurrencyPercentage gets a reference to the given int32 and assigns it to the DefaultConcurrencyPercentage field.
func (o *PrincipalRateLimitEntity) SetDefaultConcurrencyPercentage(v int32) {
	o.DefaultConcurrencyPercentage = &v
}

// GetDefaultPercentage returns the DefaultPercentage field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetDefaultPercentage() int32 {
	if o == nil || o.DefaultPercentage == nil {
		var ret int32
		return ret
	}
	return *o.DefaultPercentage
}

// GetDefaultPercentageOk returns a tuple with the DefaultPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetDefaultPercentageOk() (*int32, bool) {
	if o == nil || o.DefaultPercentage == nil {
		return nil, false
	}
	return o.DefaultPercentage, true
}

// HasDefaultPercentage returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasDefaultPercentage() bool {
	if o != nil && o.DefaultPercentage != nil {
		return true
	}

	return false
}

// SetDefaultPercentage gets a reference to the given int32 and assigns it to the DefaultPercentage field.
func (o *PrincipalRateLimitEntity) SetDefaultPercentage(v int32) {
	o.DefaultPercentage = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PrincipalRateLimitEntity) SetId(v string) {
	o.Id = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetLastUpdate() time.Time {
	if o == nil || o.LastUpdate == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil || o.LastUpdate == nil {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given time.Time and assigns it to the LastUpdate field.
func (o *PrincipalRateLimitEntity) SetLastUpdate(v time.Time) {
	o.LastUpdate = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *PrincipalRateLimitEntity) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *PrincipalRateLimitEntity) GetOrgId() string {
	if o == nil || o.OrgId == nil {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetOrgIdOk() (*string, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *PrincipalRateLimitEntity) HasOrgId() bool {
	if o != nil && o.OrgId != nil {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *PrincipalRateLimitEntity) SetOrgId(v string) {
	o.OrgId = &v
}

// GetPrincipalId returns the PrincipalId field value
func (o *PrincipalRateLimitEntity) GetPrincipalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetPrincipalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalId, true
}

// SetPrincipalId sets field value
func (o *PrincipalRateLimitEntity) SetPrincipalId(v string) {
	o.PrincipalId = v
}

// GetPrincipalType returns the PrincipalType field value
func (o *PrincipalRateLimitEntity) GetPrincipalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalType
}

// GetPrincipalTypeOk returns a tuple with the PrincipalType field value
// and a boolean to check if the value has been set.
func (o *PrincipalRateLimitEntity) GetPrincipalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalType, true
}

// SetPrincipalType sets field value
func (o *PrincipalRateLimitEntity) SetPrincipalType(v string) {
	o.PrincipalType = v
}

func (o PrincipalRateLimitEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if o.DefaultConcurrencyPercentage != nil {
		toSerialize["defaultConcurrencyPercentage"] = o.DefaultConcurrencyPercentage
	}
	if o.DefaultPercentage != nil {
		toSerialize["defaultPercentage"] = o.DefaultPercentage
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdate != nil {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if o.OrgId != nil {
		toSerialize["orgId"] = o.OrgId
	}
	if true {
		toSerialize["principalId"] = o.PrincipalId
	}
	if true {
		toSerialize["principalType"] = o.PrincipalType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalRateLimitEntity) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalRateLimitEntity := _PrincipalRateLimitEntity{}

	err = json.Unmarshal(bytes, &varPrincipalRateLimitEntity)
	if err == nil {
		*o = PrincipalRateLimitEntity(varPrincipalRateLimitEntity)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "defaultConcurrencyPercentage")
		delete(additionalProperties, "defaultPercentage")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdate")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "principalId")
		delete(additionalProperties, "principalType")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrincipalRateLimitEntity struct {
	value *PrincipalRateLimitEntity
	isSet bool
}

func (v NullablePrincipalRateLimitEntity) Get() *PrincipalRateLimitEntity {
	return v.value
}

func (v *NullablePrincipalRateLimitEntity) Set(val *PrincipalRateLimitEntity) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalRateLimitEntity) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalRateLimitEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalRateLimitEntity(val *PrincipalRateLimitEntity) *NullablePrincipalRateLimitEntity {
	return &NullablePrincipalRateLimitEntity{value: val, isSet: true}
}

func (v NullablePrincipalRateLimitEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalRateLimitEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

