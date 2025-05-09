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

// WebAuthnPreregistrationFactor User Factor variant used for WebAuthn Preregistration Factors
type WebAuthnPreregistrationFactor struct {
	// Timestamp indicating when the Factor was enrolled
	Created *time.Time `json:"created,omitempty"`
	// Type of Factor
	FactorType *string `json:"factorType,omitempty"`
	// ID of the Factor
	Id *string `json:"id,omitempty"`
	// Timestamp indicating when the Factor was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Specific attributes related to the Factor
	Profile map[string]interface{} `json:"profile,omitempty"`
	// Provider for the Factor
	Provider *string `json:"provider,omitempty"`
	// Status of the Factor
	Status *string `json:"status,omitempty"`
	// Name of the Factor vendor. This is usually the same as the provider.
	VendorName *string `json:"vendorName,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebAuthnPreregistrationFactor WebAuthnPreregistrationFactor

// NewWebAuthnPreregistrationFactor instantiates a new WebAuthnPreregistrationFactor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebAuthnPreregistrationFactor() *WebAuthnPreregistrationFactor {
	this := WebAuthnPreregistrationFactor{}
	return &this
}

// NewWebAuthnPreregistrationFactorWithDefaults instantiates a new WebAuthnPreregistrationFactor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebAuthnPreregistrationFactorWithDefaults() *WebAuthnPreregistrationFactor {
	this := WebAuthnPreregistrationFactor{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *WebAuthnPreregistrationFactor) SetCreated(v time.Time) {
	o.Created = &v
}

// GetFactorType returns the FactorType field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetFactorType() string {
	if o == nil || o.FactorType == nil {
		var ret string
		return ret
	}
	return *o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetFactorTypeOk() (*string, bool) {
	if o == nil || o.FactorType == nil {
		return nil, false
	}
	return o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasFactorType() bool {
	if o != nil && o.FactorType != nil {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given string and assigns it to the FactorType field.
func (o *WebAuthnPreregistrationFactor) SetFactorType(v string) {
	o.FactorType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WebAuthnPreregistrationFactor) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *WebAuthnPreregistrationFactor) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetProfile() map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetProfileOk() (map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]interface{} and assigns it to the Profile field.
func (o *WebAuthnPreregistrationFactor) SetProfile(v map[string]interface{}) {
	o.Profile = v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *WebAuthnPreregistrationFactor) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebAuthnPreregistrationFactor) SetStatus(v string) {
	o.Status = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetVendorName() string {
	if o == nil || o.VendorName == nil {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetVendorNameOk() (*string, bool) {
	if o == nil || o.VendorName == nil {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasVendorName() bool {
	if o != nil && o.VendorName != nil {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *WebAuthnPreregistrationFactor) SetVendorName(v string) {
	o.VendorName = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebAuthnPreregistrationFactor) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebAuthnPreregistrationFactor) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebAuthnPreregistrationFactor) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *WebAuthnPreregistrationFactor) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o WebAuthnPreregistrationFactor) MarshalJSON() ([]byte, error) {
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
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WebAuthnPreregistrationFactor) UnmarshalJSON(bytes []byte) (err error) {
	varWebAuthnPreregistrationFactor := _WebAuthnPreregistrationFactor{}

	err = json.Unmarshal(bytes, &varWebAuthnPreregistrationFactor)
	if err == nil {
		*o = WebAuthnPreregistrationFactor(varWebAuthnPreregistrationFactor)
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
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWebAuthnPreregistrationFactor struct {
	value *WebAuthnPreregistrationFactor
	isSet bool
}

func (v NullableWebAuthnPreregistrationFactor) Get() *WebAuthnPreregistrationFactor {
	return v.value
}

func (v *NullableWebAuthnPreregistrationFactor) Set(val *WebAuthnPreregistrationFactor) {
	v.value = val
	v.isSet = true
}

func (v NullableWebAuthnPreregistrationFactor) IsSet() bool {
	return v.isSet
}

func (v *NullableWebAuthnPreregistrationFactor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebAuthnPreregistrationFactor(val *WebAuthnPreregistrationFactor) *NullableWebAuthnPreregistrationFactor {
	return &NullableWebAuthnPreregistrationFactor{value: val, isSet: true}
}

func (v NullableWebAuthnPreregistrationFactor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebAuthnPreregistrationFactor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

