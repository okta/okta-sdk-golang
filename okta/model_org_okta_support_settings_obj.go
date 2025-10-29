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

// checks if the OrgOktaSupportSettingsObj type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgOktaSupportSettingsObj{}

// OrgOktaSupportSettingsObj struct for OrgOktaSupportSettingsObj
type OrgOktaSupportSettingsObj struct {
	// Support case number for the Okta Support access grant
	CaseNumber NullableString `json:"caseNumber,omitempty"`
	// Expiration of Okta Support
	Expiration NullableTime `json:"expiration,omitempty"`
	// Status of Okta Support Settings
	Support              *string                         `json:"support,omitempty"`
	Links                *OrgOktaSupportSettingsObjLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgOktaSupportSettingsObj OrgOktaSupportSettingsObj

// NewOrgOktaSupportSettingsObj instantiates a new OrgOktaSupportSettingsObj object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgOktaSupportSettingsObj() *OrgOktaSupportSettingsObj {
	this := OrgOktaSupportSettingsObj{}
	return &this
}

// NewOrgOktaSupportSettingsObjWithDefaults instantiates a new OrgOktaSupportSettingsObj object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgOktaSupportSettingsObjWithDefaults() *OrgOktaSupportSettingsObj {
	this := OrgOktaSupportSettingsObj{}
	return &this
}

// GetCaseNumber returns the CaseNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgOktaSupportSettingsObj) GetCaseNumber() string {
	if o == nil || IsNil(o.CaseNumber.Get()) {
		var ret string
		return ret
	}
	return *o.CaseNumber.Get()
}

// GetCaseNumberOk returns a tuple with the CaseNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgOktaSupportSettingsObj) GetCaseNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseNumber.Get(), o.CaseNumber.IsSet()
}

// HasCaseNumber returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasCaseNumber() bool {
	if o != nil && o.CaseNumber.IsSet() {
		return true
	}

	return false
}

// SetCaseNumber gets a reference to the given NullableString and assigns it to the CaseNumber field.
func (o *OrgOktaSupportSettingsObj) SetCaseNumber(v string) {
	o.CaseNumber.Set(&v)
}

// SetCaseNumberNil sets the value for CaseNumber to be an explicit nil
func (o *OrgOktaSupportSettingsObj) SetCaseNumberNil() {
	o.CaseNumber.Set(nil)
}

// UnsetCaseNumber ensures that no value is present for CaseNumber, not even an explicit nil
func (o *OrgOktaSupportSettingsObj) UnsetCaseNumber() {
	o.CaseNumber.Unset()
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrgOktaSupportSettingsObj) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrgOktaSupportSettingsObj) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *OrgOktaSupportSettingsObj) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *OrgOktaSupportSettingsObj) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *OrgOktaSupportSettingsObj) UnsetExpiration() {
	o.Expiration.Unset()
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObj) GetSupport() string {
	if o == nil || IsNil(o.Support) {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObj) GetSupportOk() (*string, bool) {
	if o == nil || IsNil(o.Support) {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasSupport() bool {
	if o != nil && !IsNil(o.Support) {
		return true
	}

	return false
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *OrgOktaSupportSettingsObj) SetSupport(v string) {
	o.Support = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObj) GetLinks() OrgOktaSupportSettingsObjLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrgOktaSupportSettingsObjLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObj) GetLinksOk() (*OrgOktaSupportSettingsObjLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgOktaSupportSettingsObjLinks and assigns it to the Links field.
func (o *OrgOktaSupportSettingsObj) SetLinks(v OrgOktaSupportSettingsObjLinks) {
	o.Links = &v
}

func (o OrgOktaSupportSettingsObj) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgOktaSupportSettingsObj) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CaseNumber.IsSet() {
		toSerialize["caseNumber"] = o.CaseNumber.Get()
	}
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}
	if !IsNil(o.Support) {
		toSerialize["support"] = o.Support
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgOktaSupportSettingsObj) UnmarshalJSON(data []byte) (err error) {
	varOrgOktaSupportSettingsObj := _OrgOktaSupportSettingsObj{}

	err = json.Unmarshal(data, &varOrgOktaSupportSettingsObj)

	if err != nil {
		return err
	}

	*o = OrgOktaSupportSettingsObj(varOrgOktaSupportSettingsObj)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "caseNumber")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "support")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgOktaSupportSettingsObj struct {
	value *OrgOktaSupportSettingsObj
	isSet bool
}

func (v NullableOrgOktaSupportSettingsObj) Get() *OrgOktaSupportSettingsObj {
	return v.value
}

func (v *NullableOrgOktaSupportSettingsObj) Set(val *OrgOktaSupportSettingsObj) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOktaSupportSettingsObj) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOktaSupportSettingsObj) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOktaSupportSettingsObj(val *OrgOktaSupportSettingsObj) *NullableOrgOktaSupportSettingsObj {
	return &NullableOrgOktaSupportSettingsObj{value: val, isSet: true}
}

func (v NullableOrgOktaSupportSettingsObj) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOktaSupportSettingsObj) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
