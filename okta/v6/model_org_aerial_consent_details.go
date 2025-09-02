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
	"fmt"
)

// OrgAerialConsentDetails struct for OrgAerialConsentDetails
type OrgAerialConsentDetails struct {
	// The unique ID of the Aerial account
	AccountId string `json:"accountId"`
	// Principal ID of the user who granted the permission
	GrantedBy *string `json:"grantedBy,omitempty"`
	// Date when grant was created
	GrantedDate *string `json:"grantedDate,omitempty"`
	Links *LinksAerialConsentGranted `json:"_links,omitempty"`
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
	if o == nil || o.GrantedBy == nil {
		var ret string
		return ret
	}
	return *o.GrantedBy
}

// GetGrantedByOk returns a tuple with the GrantedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetGrantedByOk() (*string, bool) {
	if o == nil || o.GrantedBy == nil {
		return nil, false
	}
	return o.GrantedBy, true
}

// HasGrantedBy returns a boolean if a field has been set.
func (o *OrgAerialConsentDetails) HasGrantedBy() bool {
	if o != nil && o.GrantedBy != nil {
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
	if o == nil || o.GrantedDate == nil {
		var ret string
		return ret
	}
	return *o.GrantedDate
}

// GetGrantedDateOk returns a tuple with the GrantedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetGrantedDateOk() (*string, bool) {
	if o == nil || o.GrantedDate == nil {
		return nil, false
	}
	return o.GrantedDate, true
}

// HasGrantedDate returns a boolean if a field has been set.
func (o *OrgAerialConsentDetails) HasGrantedDate() bool {
	if o != nil && o.GrantedDate != nil {
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
	if o == nil || o.Links == nil {
		var ret LinksAerialConsentGranted
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgAerialConsentDetails) GetLinksOk() (*LinksAerialConsentGranted, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgAerialConsentDetails) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksAerialConsentGranted and assigns it to the Links field.
func (o *OrgAerialConsentDetails) SetLinks(v LinksAerialConsentGranted) {
	o.Links = &v
}

func (o OrgAerialConsentDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if o.GrantedBy != nil {
		toSerialize["grantedBy"] = o.GrantedBy
	}
	if o.GrantedDate != nil {
		toSerialize["grantedDate"] = o.GrantedDate
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgAerialConsentDetails) UnmarshalJSON(bytes []byte) (err error) {
	varOrgAerialConsentDetails := _OrgAerialConsentDetails{}

	err = json.Unmarshal(bytes, &varOrgAerialConsentDetails)
	if err == nil {
		*o = OrgAerialConsentDetails(varOrgAerialConsentDetails)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountId")
		delete(additionalProperties, "grantedBy")
		delete(additionalProperties, "grantedDate")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

