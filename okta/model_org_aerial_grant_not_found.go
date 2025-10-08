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
)

// checks if the OrgAerialGrantNotFound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgAerialGrantNotFound{}

// OrgAerialGrantNotFound struct for OrgAerialGrantNotFound
type OrgAerialGrantNotFound struct {
	// The unique ID of the Aerial account
	AccountId *string `json:"accountId,omitempty"`
	// Principal ID of the user who granted the permission
	GrantedBy *string `json:"grantedBy,omitempty"`
	// Date when grant was created
	GrantedDate          *string                    `json:"grantedDate,omitempty"`
	Links                *LinksAerialConsentGranted `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgAerialGrantNotFound OrgAerialGrantNotFound

// NewOrgAerialGrantNotFound instantiates a new OrgAerialGrantNotFound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgAerialGrantNotFound() *OrgAerialGrantNotFound {
	this := OrgAerialGrantNotFound{}
	return &this
}

// NewOrgAerialGrantNotFoundWithDefaults instantiates a new OrgAerialGrantNotFound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgAerialGrantNotFoundWithDefaults() *OrgAerialGrantNotFound {
	this := OrgAerialGrantNotFound{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *OrgAerialGrantNotFound) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialGrantNotFound) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *OrgAerialGrantNotFound) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *OrgAerialGrantNotFound) SetAccountId(v string) {
	o.AccountId = &v
}

// GetGrantedBy returns the GrantedBy field value if set, zero value otherwise.
func (o *OrgAerialGrantNotFound) GetGrantedBy() string {
	if o == nil || IsNil(o.GrantedBy) {
		var ret string
		return ret
	}
	return *o.GrantedBy
}

// GetGrantedByOk returns a tuple with the GrantedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialGrantNotFound) GetGrantedByOk() (*string, bool) {
	if o == nil || IsNil(o.GrantedBy) {
		return nil, false
	}
	return o.GrantedBy, true
}

// HasGrantedBy returns a boolean if a field has been set.
func (o *OrgAerialGrantNotFound) HasGrantedBy() bool {
	if o != nil && !IsNil(o.GrantedBy) {
		return true
	}

	return false
}

// SetGrantedBy gets a reference to the given string and assigns it to the GrantedBy field.
func (o *OrgAerialGrantNotFound) SetGrantedBy(v string) {
	o.GrantedBy = &v
}

// GetGrantedDate returns the GrantedDate field value if set, zero value otherwise.
func (o *OrgAerialGrantNotFound) GetGrantedDate() string {
	if o == nil || IsNil(o.GrantedDate) {
		var ret string
		return ret
	}
	return *o.GrantedDate
}

// GetGrantedDateOk returns a tuple with the GrantedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialGrantNotFound) GetGrantedDateOk() (*string, bool) {
	if o == nil || IsNil(o.GrantedDate) {
		return nil, false
	}
	return o.GrantedDate, true
}

// HasGrantedDate returns a boolean if a field has been set.
func (o *OrgAerialGrantNotFound) HasGrantedDate() bool {
	if o != nil && !IsNil(o.GrantedDate) {
		return true
	}

	return false
}

// SetGrantedDate gets a reference to the given string and assigns it to the GrantedDate field.
func (o *OrgAerialGrantNotFound) SetGrantedDate(v string) {
	o.GrantedDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgAerialGrantNotFound) GetLinks() LinksAerialConsentGranted {
	if o == nil || IsNil(o.Links) {
		var ret LinksAerialConsentGranted
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialGrantNotFound) GetLinksOk() (*LinksAerialConsentGranted, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgAerialGrantNotFound) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksAerialConsentGranted and assigns it to the Links field.
func (o *OrgAerialGrantNotFound) SetLinks(v LinksAerialConsentGranted) {
	o.Links = &v
}

func (o OrgAerialGrantNotFound) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgAerialGrantNotFound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
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

func (o *OrgAerialGrantNotFound) UnmarshalJSON(data []byte) (err error) {
	varOrgAerialGrantNotFound := _OrgAerialGrantNotFound{}

	err = json.Unmarshal(data, &varOrgAerialGrantNotFound)

	if err != nil {
		return err
	}

	*o = OrgAerialGrantNotFound(varOrgAerialGrantNotFound)

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

type NullableOrgAerialGrantNotFound struct {
	value *OrgAerialGrantNotFound
	isSet bool
}

func (v NullableOrgAerialGrantNotFound) Get() *OrgAerialGrantNotFound {
	return v.value
}

func (v *NullableOrgAerialGrantNotFound) Set(val *OrgAerialGrantNotFound) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgAerialGrantNotFound) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgAerialGrantNotFound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgAerialGrantNotFound(val *OrgAerialGrantNotFound) *NullableOrgAerialGrantNotFound {
	return &NullableOrgAerialGrantNotFound{value: val, isSet: true}
}

func (v NullableOrgAerialGrantNotFound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgAerialGrantNotFound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
