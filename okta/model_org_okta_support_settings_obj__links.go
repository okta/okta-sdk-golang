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

// checks if the OrgOktaSupportSettingsObjLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgOktaSupportSettingsObjLinks{}

// OrgOktaSupportSettingsObjLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the Okta Support Settings object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgOktaSupportSettingsObjLinks struct {
	// Link to [extend Okta Support Access](/openapi/okta-management/management/tag/OrgSettingSupport/#tag/OrgSettingSupport/operation/extendOktaSupport)
	Extend *HrefObject `json:"extend,omitempty"`
	// Link to [revoke Okta Support Access](/openapi/okta-management/management/tag/OrgSettingSupport/#tag/OrgSettingSupport/operation/revokeOktaSupport)
	Revoke *HrefObject `json:"revoke,omitempty"`
	// Link to [grant Okta Support Access](/openapi/okta-management/management/tag/OrgSettingSupport/#tag/OrgSettingSupport/operation/grantOktaSupport)
	Grant *HrefObject `json:"grant,omitempty"`
	// Link to [update an Okta Support case](/openapi/okta-management/management/tag/OrgSettingSupport/#tag/OrgSettingSupport/operation/updateOktaSupportCase)
	Case *HrefObject `json:"case,omitempty"`
	// Link to [List all Okta Support cases](/openapi/okta-management/management/tag/OrgSettingSupport/#tag/OrgSettingSupport/operation/listOktaSupportCases)
	Cases                *HrefObject `json:"cases,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgOktaSupportSettingsObjLinks OrgOktaSupportSettingsObjLinks

// NewOrgOktaSupportSettingsObjLinks instantiates a new OrgOktaSupportSettingsObjLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgOktaSupportSettingsObjLinks() *OrgOktaSupportSettingsObjLinks {
	this := OrgOktaSupportSettingsObjLinks{}
	return &this
}

// NewOrgOktaSupportSettingsObjLinksWithDefaults instantiates a new OrgOktaSupportSettingsObjLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgOktaSupportSettingsObjLinksWithDefaults() *OrgOktaSupportSettingsObjLinks {
	this := OrgOktaSupportSettingsObjLinks{}
	return &this
}

// GetExtend returns the Extend field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObjLinks) GetExtend() HrefObject {
	if o == nil || IsNil(o.Extend) {
		var ret HrefObject
		return ret
	}
	return *o.Extend
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObjLinks) GetExtendOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Extend) {
		return nil, false
	}
	return o.Extend, true
}

// HasExtend returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObjLinks) HasExtend() bool {
	if o != nil && !IsNil(o.Extend) {
		return true
	}

	return false
}

// SetExtend gets a reference to the given HrefObject and assigns it to the Extend field.
func (o *OrgOktaSupportSettingsObjLinks) SetExtend(v HrefObject) {
	o.Extend = &v
}

// GetRevoke returns the Revoke field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObjLinks) GetRevoke() HrefObject {
	if o == nil || IsNil(o.Revoke) {
		var ret HrefObject
		return ret
	}
	return *o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObjLinks) GetRevokeOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Revoke) {
		return nil, false
	}
	return o.Revoke, true
}

// HasRevoke returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObjLinks) HasRevoke() bool {
	if o != nil && !IsNil(o.Revoke) {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given HrefObject and assigns it to the Revoke field.
func (o *OrgOktaSupportSettingsObjLinks) SetRevoke(v HrefObject) {
	o.Revoke = &v
}

// GetGrant returns the Grant field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObjLinks) GetGrant() HrefObject {
	if o == nil || IsNil(o.Grant) {
		var ret HrefObject
		return ret
	}
	return *o.Grant
}

// GetGrantOk returns a tuple with the Grant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObjLinks) GetGrantOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Grant) {
		return nil, false
	}
	return o.Grant, true
}

// HasGrant returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObjLinks) HasGrant() bool {
	if o != nil && !IsNil(o.Grant) {
		return true
	}

	return false
}

// SetGrant gets a reference to the given HrefObject and assigns it to the Grant field.
func (o *OrgOktaSupportSettingsObjLinks) SetGrant(v HrefObject) {
	o.Grant = &v
}

// GetCase returns the Case field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObjLinks) GetCase() HrefObject {
	if o == nil || IsNil(o.Case) {
		var ret HrefObject
		return ret
	}
	return *o.Case
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObjLinks) GetCaseOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Case) {
		return nil, false
	}
	return o.Case, true
}

// HasCase returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObjLinks) HasCase() bool {
	if o != nil && !IsNil(o.Case) {
		return true
	}

	return false
}

// SetCase gets a reference to the given HrefObject and assigns it to the Case field.
func (o *OrgOktaSupportSettingsObjLinks) SetCase(v HrefObject) {
	o.Case = &v
}

// GetCases returns the Cases field value if set, zero value otherwise.
func (o *OrgOktaSupportSettingsObjLinks) GetCases() HrefObject {
	if o == nil || IsNil(o.Cases) {
		var ret HrefObject
		return ret
	}
	return *o.Cases
}

// GetCasesOk returns a tuple with the Cases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgOktaSupportSettingsObjLinks) GetCasesOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Cases) {
		return nil, false
	}
	return o.Cases, true
}

// HasCases returns a boolean if a field has been set.
func (o *OrgOktaSupportSettingsObjLinks) HasCases() bool {
	if o != nil && !IsNil(o.Cases) {
		return true
	}

	return false
}

// SetCases gets a reference to the given HrefObject and assigns it to the Cases field.
func (o *OrgOktaSupportSettingsObjLinks) SetCases(v HrefObject) {
	o.Cases = &v
}

func (o OrgOktaSupportSettingsObjLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgOktaSupportSettingsObjLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Extend) {
		toSerialize["extend"] = o.Extend
	}
	if !IsNil(o.Revoke) {
		toSerialize["revoke"] = o.Revoke
	}
	if !IsNil(o.Grant) {
		toSerialize["grant"] = o.Grant
	}
	if !IsNil(o.Case) {
		toSerialize["case"] = o.Case
	}
	if !IsNil(o.Cases) {
		toSerialize["cases"] = o.Cases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgOktaSupportSettingsObjLinks) UnmarshalJSON(data []byte) (err error) {
	varOrgOktaSupportSettingsObjLinks := _OrgOktaSupportSettingsObjLinks{}

	err = json.Unmarshal(data, &varOrgOktaSupportSettingsObjLinks)

	if err != nil {
		return err
	}

	*o = OrgOktaSupportSettingsObjLinks(varOrgOktaSupportSettingsObjLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "extend")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "grant")
		delete(additionalProperties, "case")
		delete(additionalProperties, "cases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgOktaSupportSettingsObjLinks struct {
	value *OrgOktaSupportSettingsObjLinks
	isSet bool
}

func (v NullableOrgOktaSupportSettingsObjLinks) Get() *OrgOktaSupportSettingsObjLinks {
	return v.value
}

func (v *NullableOrgOktaSupportSettingsObjLinks) Set(val *OrgOktaSupportSettingsObjLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgOktaSupportSettingsObjLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgOktaSupportSettingsObjLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgOktaSupportSettingsObjLinks(val *OrgOktaSupportSettingsObjLinks) *NullableOrgOktaSupportSettingsObjLinks {
	return &NullableOrgOktaSupportSettingsObjLinks{value: val, isSet: true}
}

func (v NullableOrgOktaSupportSettingsObjLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgOktaSupportSettingsObjLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
