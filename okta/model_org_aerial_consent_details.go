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
	"fmt"
)

// checks if the OrgAerialConsentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgAerialConsentDetails{}

// OrgAerialConsentDetails struct for OrgAerialConsentDetails
type OrgAerialConsentDetails struct {
	// The unique ID of the Aerial account
	AccountId string `json:"accountId"`
	// Principal ID of the user who granted the permission
	GrantedBy *string `json:"grantedBy,omitempty"`
	// Date when grant was created
	GrantedDate          *string                    `json:"grantedDate,omitempty"`
	Links                *LinksAerialConsentGranted `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgAerialConsentDetails OrgAerialConsentDetails

// NewOrgAerialConsentDetails instantiates a new OrgAerialConsentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgAerialConsentDetails(accountId string) *OrgAerialConsentDetails {
	this := OrgAerialConsentDetails{}
	this.AccountId = accountId
	return &this
}

// NewOrgAerialConsentDetailsWithDefaults instantiates a new OrgAerialConsentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgAerialConsentDetailsWithDefaults() *OrgAerialConsentDetails {
	this := OrgAerialConsentDetails{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *OrgAerialConsentDetails) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *OrgAerialConsentDetails) SetAccountId(v string) {
	o.AccountId = v
}

// GetGrantedBy returns the GrantedBy field value if set, zero value otherwise.
func (o *OrgAerialConsentDetails) GetGrantedBy() string {
	if o == nil || IsNil(o.GrantedBy) {
		var ret string
		return ret
	}
	return *o.GrantedBy
}

// GetGrantedByOk returns a tuple with the GrantedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetGrantedByOk() (*string, bool) {
	if o == nil || IsNil(o.GrantedBy) {
		return nil, false
	}
	return o.GrantedBy, true
}

// HasGrantedBy returns a boolean if a field has been set.
func (o *OrgAerialConsentDetails) HasGrantedBy() bool {
	if o != nil && !IsNil(o.GrantedBy) {
		return true
	}

	return false
}

// SetGrantedBy gets a reference to the given string and assigns it to the GrantedBy field.
func (o *OrgAerialConsentDetails) SetGrantedBy(v string) {
	o.GrantedBy = &v
}

// GetGrantedDate returns the GrantedDate field value if set, zero value otherwise.
func (o *OrgAerialConsentDetails) GetGrantedDate() string {
	if o == nil || IsNil(o.GrantedDate) {
		var ret string
		return ret
	}
	return *o.GrantedDate
}

// GetGrantedDateOk returns a tuple with the GrantedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetGrantedDateOk() (*string, bool) {
	if o == nil || IsNil(o.GrantedDate) {
		return nil, false
	}
	return o.GrantedDate, true
}

// HasGrantedDate returns a boolean if a field has been set.
func (o *OrgAerialConsentDetails) HasGrantedDate() bool {
	if o != nil && !IsNil(o.GrantedDate) {
		return true
	}

	return false
}

// SetGrantedDate gets a reference to the given string and assigns it to the GrantedDate field.
func (o *OrgAerialConsentDetails) SetGrantedDate(v string) {
	o.GrantedDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgAerialConsentDetails) GetLinks() LinksAerialConsentGranted {
	if o == nil || IsNil(o.Links) {
		var ret LinksAerialConsentGranted
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetLinksOk() (*LinksAerialConsentGranted, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgAerialConsentDetails) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksAerialConsentGranted and assigns it to the Links field.
func (o *OrgAerialConsentDetails) SetLinks(v LinksAerialConsentGranted) {
	o.Links = &v
}

func (o OrgAerialConsentDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgAerialConsentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	if !IsNil(o.GrantedBy) {
		toSerialize["grantedBy"] = o.GrantedBy
	}
	if !IsNil(o.GrantedDate) {
		toSerialize["grantedDate"] = o.GrantedDate
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgAerialConsentDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOrgAerialConsentDetails := _OrgAerialConsentDetails{}

	err = json.Unmarshal(data, &varOrgAerialConsentDetails)

	if err != nil {
		return err
	}

	*o = OrgAerialConsentDetails(varOrgAerialConsentDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "grantedBy")
		delete(additionalProperties, "grantedDate")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgAerialConsentDetails struct {
	value *OrgAerialConsentDetails
	isSet bool
}

func (v NullableOrgAerialConsentDetails) Get() *OrgAerialConsentDetails {
	return v.value
}

func (v *NullableOrgAerialConsentDetails) Set(val *OrgAerialConsentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgAerialConsentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgAerialConsentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgAerialConsentDetails(val *OrgAerialConsentDetails) *NullableOrgAerialConsentDetails {
	return &NullableOrgAerialConsentDetails{value: val, isSet: true}
}

func (v NullableOrgAerialConsentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgAerialConsentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
