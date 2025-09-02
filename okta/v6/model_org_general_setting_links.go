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
)

// OrgGeneralSettingLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the org using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
type OrgGeneralSettingLinks struct {
	// Link to the [Org Contacts](/openapi/okta-management/management/tag/OrgSettingContact/) resource
	Contacts *HrefObject `json:"contacts,omitempty"`
	// Link to the org logo
	Logo *HrefObject `json:"logo,omitempty"`
	// Link to the [Org Communication Settings](/openapi/okta-management/management/tag/OrgSettingCommunication/) resource
	OktaCommunication *HrefObject `json:"oktaCommunication,omitempty"`
	// Link to the [Org Support Settings](/openapi/okta-management/management/tag/OrgSettingSupport/) resource
	OktaSupport *HrefObject `json:"oktaSupport,omitempty"`
	// Link to the [Org Preferences](/openapi/okta-management/management/tag/OrgSettingCustomization/#tag/OrgSettingCustomization/operation/getOrgPreferences) resource
	Preferences *HrefObject `json:"preferences,omitempty"`
	// Link to the [Upload Org Logo](/openapi/okta-management/management/tag/OrgSettingCustomization/#tag/OrgSettingCustomization/operation/uploadOrgLogo) resource
	UploadLogo *HrefObject `json:"uploadLogo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgGeneralSettingLinks OrgGeneralSettingLinks

// NewOrgGeneralSettingLinks instantiates a new OrgGeneralSettingLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgGeneralSettingLinks() *OrgGeneralSettingLinks {
	this := OrgGeneralSettingLinks{}
	return &this
}

// NewOrgGeneralSettingLinksWithDefaults instantiates a new OrgGeneralSettingLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgGeneralSettingLinksWithDefaults() *OrgGeneralSettingLinks {
	this := OrgGeneralSettingLinks{}
	return &this
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *OrgGeneralSettingLinks) GetContacts() HrefObject {
	if o == nil || o.Contacts == nil {
		var ret HrefObject
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGeneralSettingLinks) GetContactsOk() (*HrefObject, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *OrgGeneralSettingLinks) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given HrefObject and assigns it to the Contacts field.
func (o *OrgGeneralSettingLinks) SetContacts(v HrefObject) {
	o.Contacts = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *OrgGeneralSettingLinks) GetLogo() HrefObject {
	if o == nil || o.Logo == nil {
		var ret HrefObject
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGeneralSettingLinks) GetLogoOk() (*HrefObject, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *OrgGeneralSettingLinks) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given HrefObject and assigns it to the Logo field.
func (o *OrgGeneralSettingLinks) SetLogo(v HrefObject) {
	o.Logo = &v
}

// GetOktaCommunication returns the OktaCommunication field value if set, zero value otherwise.
func (o *OrgGeneralSettingLinks) GetOktaCommunication() HrefObject {
	if o == nil || o.OktaCommunication == nil {
		var ret HrefObject
		return ret
	}
	return *o.OktaCommunication
}

// GetOktaCommunicationOk returns a tuple with the OktaCommunication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGeneralSettingLinks) GetOktaCommunicationOk() (*HrefObject, bool) {
	if o == nil || o.OktaCommunication == nil {
		return nil, false
	}
	return o.OktaCommunication, true
}

// HasOktaCommunication returns a boolean if a field has been set.
func (o *OrgGeneralSettingLinks) HasOktaCommunication() bool {
	if o != nil && o.OktaCommunication != nil {
		return true
	}

	return false
}

// SetOktaCommunication gets a reference to the given HrefObject and assigns it to the OktaCommunication field.
func (o *OrgGeneralSettingLinks) SetOktaCommunication(v HrefObject) {
	o.OktaCommunication = &v
}

// GetOktaSupport returns the OktaSupport field value if set, zero value otherwise.
func (o *OrgGeneralSettingLinks) GetOktaSupport() HrefObject {
	if o == nil || o.OktaSupport == nil {
		var ret HrefObject
		return ret
	}
	return *o.OktaSupport
}

// GetOktaSupportOk returns a tuple with the OktaSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGeneralSettingLinks) GetOktaSupportOk() (*HrefObject, bool) {
	if o == nil || o.OktaSupport == nil {
		return nil, false
	}
	return o.OktaSupport, true
}

// HasOktaSupport returns a boolean if a field has been set.
func (o *OrgGeneralSettingLinks) HasOktaSupport() bool {
	if o != nil && o.OktaSupport != nil {
		return true
	}

	return false
}

// SetOktaSupport gets a reference to the given HrefObject and assigns it to the OktaSupport field.
func (o *OrgGeneralSettingLinks) SetOktaSupport(v HrefObject) {
	o.OktaSupport = &v
}

// GetPreferences returns the Preferences field value if set, zero value otherwise.
func (o *OrgGeneralSettingLinks) GetPreferences() HrefObject {
	if o == nil || o.Preferences == nil {
		var ret HrefObject
		return ret
	}
	return *o.Preferences
}

// GetPreferencesOk returns a tuple with the Preferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGeneralSettingLinks) GetPreferencesOk() (*HrefObject, bool) {
	if o == nil || o.Preferences == nil {
		return nil, false
	}
	return o.Preferences, true
}

// HasPreferences returns a boolean if a field has been set.
func (o *OrgGeneralSettingLinks) HasPreferences() bool {
	if o != nil && o.Preferences != nil {
		return true
	}

	return false
}

// SetPreferences gets a reference to the given HrefObject and assigns it to the Preferences field.
func (o *OrgGeneralSettingLinks) SetPreferences(v HrefObject) {
	o.Preferences = &v
}

// GetUploadLogo returns the UploadLogo field value if set, zero value otherwise.
func (o *OrgGeneralSettingLinks) GetUploadLogo() HrefObject {
	if o == nil || o.UploadLogo == nil {
		var ret HrefObject
		return ret
	}
	return *o.UploadLogo
}

// GetUploadLogoOk returns a tuple with the UploadLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgGeneralSettingLinks) GetUploadLogoOk() (*HrefObject, bool) {
	if o == nil || o.UploadLogo == nil {
		return nil, false
	}
	return o.UploadLogo, true
}

// HasUploadLogo returns a boolean if a field has been set.
func (o *OrgGeneralSettingLinks) HasUploadLogo() bool {
	if o != nil && o.UploadLogo != nil {
		return true
	}

	return false
}

// SetUploadLogo gets a reference to the given HrefObject and assigns it to the UploadLogo field.
func (o *OrgGeneralSettingLinks) SetUploadLogo(v HrefObject) {
	o.UploadLogo = &v
}

func (o OrgGeneralSettingLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.OktaCommunication != nil {
		toSerialize["oktaCommunication"] = o.OktaCommunication
	}
	if o.OktaSupport != nil {
		toSerialize["oktaSupport"] = o.OktaSupport
	}
	if o.Preferences != nil {
		toSerialize["preferences"] = o.Preferences
	}
	if o.UploadLogo != nil {
		toSerialize["uploadLogo"] = o.UploadLogo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgGeneralSettingLinks) UnmarshalJSON(bytes []byte) (err error) {
	varOrgGeneralSettingLinks := _OrgGeneralSettingLinks{}

	err = json.Unmarshal(bytes, &varOrgGeneralSettingLinks)
	if err == nil {
		*o = OrgGeneralSettingLinks(varOrgGeneralSettingLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "contacts")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "oktaCommunication")
		delete(additionalProperties, "oktaSupport")
		delete(additionalProperties, "preferences")
		delete(additionalProperties, "uploadLogo")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOrgGeneralSettingLinks struct {
	value *OrgGeneralSettingLinks
	isSet bool
}

func (v NullableOrgGeneralSettingLinks) Get() *OrgGeneralSettingLinks {
	return v.value
}

func (v *NullableOrgGeneralSettingLinks) Set(val *OrgGeneralSettingLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgGeneralSettingLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgGeneralSettingLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgGeneralSettingLinks(val *OrgGeneralSettingLinks) *NullableOrgGeneralSettingLinks {
	return &NullableOrgGeneralSettingLinks{value: val, isSet: true}
}

func (v NullableOrgGeneralSettingLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgGeneralSettingLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

