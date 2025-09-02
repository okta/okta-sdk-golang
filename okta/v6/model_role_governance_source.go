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

// RoleGovernanceSource User role governance source
type RoleGovernanceSource struct {
	// `id` of the entitlement bundle
	BundleId *string `json:"bundleId,omitempty"`
	// The expiration date of the entitlement bundle
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	// `id` of the grant
	GrantId string `json:"grantId"`
	// The grant type
	Type string `json:"type"`
	Links *RoleGovernanceSourceLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoleGovernanceSource RoleGovernanceSource

// NewRoleGovernanceSource instantiates a new RoleGovernanceSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleGovernanceSource(grantId string, type_ string) *RoleGovernanceSource {
	this := RoleGovernanceSource{}
	this.GrantId = grantId
	this.Type = type_
	return &this
}

// NewRoleGovernanceSourceWithDefaults instantiates a new RoleGovernanceSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleGovernanceSourceWithDefaults() *RoleGovernanceSource {
	this := RoleGovernanceSource{}
	return &this
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *RoleGovernanceSource) GetBundleId() string {
	if o == nil || o.BundleId == nil {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSource) GetBundleIdOk() (*string, bool) {
	if o == nil || o.BundleId == nil {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *RoleGovernanceSource) HasBundleId() bool {
	if o != nil && o.BundleId != nil {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *RoleGovernanceSource) SetBundleId(v string) {
	o.BundleId = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *RoleGovernanceSource) GetExpirationDate() time.Time {
	if o == nil || o.ExpirationDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSource) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *RoleGovernanceSource) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate != nil {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *RoleGovernanceSource) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetGrantId returns the GrantId field value
func (o *RoleGovernanceSource) GetGrantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GrantId
}

// GetGrantIdOk returns a tuple with the GrantId field value
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSource) GetGrantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GrantId, true
}

// SetGrantId sets field value
func (o *RoleGovernanceSource) SetGrantId(v string) {
	o.GrantId = v
}

// GetType returns the Type field value
func (o *RoleGovernanceSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSource) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoleGovernanceSource) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RoleGovernanceSource) GetLinks() RoleGovernanceSourceLinks {
	if o == nil || o.Links == nil {
		var ret RoleGovernanceSourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleGovernanceSource) GetLinksOk() (*RoleGovernanceSourceLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RoleGovernanceSource) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given RoleGovernanceSourceLinks and assigns it to the Links field.
func (o *RoleGovernanceSource) SetLinks(v RoleGovernanceSourceLinks) {
	o.Links = &v
}

func (o RoleGovernanceSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BundleId != nil {
		toSerialize["bundleId"] = o.BundleId
	}
	if o.ExpirationDate != nil {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if true {
		toSerialize["grantId"] = o.GrantId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RoleGovernanceSource) UnmarshalJSON(bytes []byte) (err error) {
	varRoleGovernanceSource := _RoleGovernanceSource{}

	err = json.Unmarshal(bytes, &varRoleGovernanceSource)
	if err == nil {
		*o = RoleGovernanceSource(varRoleGovernanceSource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "bundleId")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "grantId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRoleGovernanceSource struct {
	value *RoleGovernanceSource
	isSet bool
}

func (v NullableRoleGovernanceSource) Get() *RoleGovernanceSource {
	return v.value
}

func (v *NullableRoleGovernanceSource) Set(val *RoleGovernanceSource) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleGovernanceSource) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleGovernanceSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleGovernanceSource(val *RoleGovernanceSource) *NullableRoleGovernanceSource {
	return &NullableRoleGovernanceSource{value: val, isSet: true}
}

func (v NullableRoleGovernanceSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleGovernanceSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

