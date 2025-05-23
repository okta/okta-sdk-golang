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

// OrgOktaSupportSettingsObj struct for OrgOktaSupportSettingsObj
type OrgOktaSupportSettingsObj struct {
	Expiration *time.Time `json:"expiration,omitempty"`
	Support *string `json:"support,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
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

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObj) GetExpiration() time.Time {
	if o == nil || o.Expiration == nil {
		var ret time.Time
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObj) GetExpirationOk() (*time.Time, bool) {
	if o == nil || o.Expiration == nil {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasExpiration() bool {
	if o != nil && o.Expiration != nil {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given time.Time and assigns it to the Expiration field.
func (o *OrgOktaSupportSettingsObj) SetExpiration(v time.Time) {
	o.Expiration = &v
}

// GetSupport returns the Support field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObj) GetSupport() string {
	if o == nil || o.Support == nil {
		var ret string
		return ret
	}
	return *o.Support
}

// GetSupportOk returns a tuple with the Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObj) GetSupportOk() (*string, bool) {
	if o == nil || o.Support == nil {
		return nil, false
	}
	return o.Support, true
}

// HasSupport returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasSupport() bool {
	if o != nil && o.Support != nil {
		return true
	}

	return false
}

// SetSupport gets a reference to the given string and assigns it to the Support field.
func (o *OrgOktaSupportSettingsObj) SetSupport(v string) {
	o.Support = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObj) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObj) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObj) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *OrgOktaSupportSettingsObj) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o OrgOktaSupportSettingsObj) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiration != nil {
		toSerialize["expiration"] = o.Expiration
	}
	if o.Support != nil {
		toSerialize["support"] = o.Support
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgOktaSupportSettingsObj) UnmarshalJSON(bytes []byte) (err error) {
	varOrgOktaSupportSettingsObj := _OrgOktaSupportSettingsObj{}

	err = json.Unmarshal(bytes, &varOrgOktaSupportSettingsObj)
	if err == nil {
		*o = OrgOktaSupportSettingsObj(varOrgOktaSupportSettingsObj)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "support")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

