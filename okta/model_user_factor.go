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

// UserFactor struct for UserFactor
type UserFactor struct {
	// Timestamp when the Factor was enrolled
	Created *time.Time `json:"created,omitempty"`
	// Type of Factor
	FactorType *string `json:"factorType,omitempty"`
	// ID of the Factor
	Id *string `json:"id,omitempty"`
	// Timestamp when the Factor was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Specific attributes related to the Factor
	Profile map[string]interface{} `json:"profile,omitempty"`
	// Provider for the Factor
	Provider *string `json:"provider,omitempty"`
	// Status of the Factor
	Status *string `json:"status,omitempty"`
	// Name of the Factor vendor. This is usually the same as the provider except for On-Prem MFA where it depends on administrator settings.
	VendorName *string `json:"vendorName,omitempty"`
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactor UserFactor

// NewUserFactor instantiates a new UserFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactor() *UserFactor {
	this := UserFactor{}
	return &this
}

// NewUserFactorWithDefaults instantiates a new UserFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorWithDefaults() *UserFactor {
	this := UserFactor{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserFactor) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserFactor) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserFactor) SetCreated(v time.Time) {
	o.Created = &v
}

// GetFactorType returns the FactorType field value if set, zero value otherwise.
func (o *UserFactor) GetFactorType() string {
	if o == nil || o.FactorType == nil {
		var ret string
		return ret
	}
	return *o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetFactorTypeOk() (*string, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactor) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given string and assigns it to the FactorType field.
func (o *UserFactor) SetFactorType(v string) {
	o.FactorType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserFactor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserFactor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserFactor) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UserFactor) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserFactor) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UserFactor) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactor) GetProfile() map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *UserFactor) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactor) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactor) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactor) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserFactor) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserFactor) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserFactor) SetStatus(v string) {
	o.Status = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *UserFactor) GetVendorName() string {
	if o == nil || o.VendorName == nil {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetVendorNameOk() (*string, bool) {
	if o == nil || o.VendorName == nil {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *UserFactor) HasVendorName() bool {
	if o != nil && o.VendorName != nil {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *UserFactor) SetVendorName(v string) {
	o.VendorName = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactor) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactor) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *UserFactor) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactor) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactor) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactor) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *UserFactor) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o UserFactor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.FactorType != nil {
		toSerialize["factorType"] = o.FactorType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.VendorName != nil {
		toSerialize["vendorName"] = o.VendorName
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactor) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactor := _UserFactor{}

	err = json.Unmarshal(bytes, &varUserFactor)
	if err == nil {
		*o = UserFactor(varUserFactor)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "factorType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "status")
		delete(additionalProperties, "vendorName")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactor struct {
	value *UserFactor
	isSet bool
}

func (v NullableUserFactor) Get() *UserFactor {
	return v.value
}

func (v *NullableUserFactor) Set(val *UserFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactor(val *UserFactor) *NullableUserFactor {
	return &NullableUserFactor{value: val, isSet: true}
}

func (v NullableUserFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

