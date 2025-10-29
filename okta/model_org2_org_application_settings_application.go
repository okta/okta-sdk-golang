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

// checks if the Org2OrgApplicationSettingsApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Org2OrgApplicationSettingsApplication{}

// Org2OrgApplicationSettingsApplication Org2Org app instance properties
type Org2OrgApplicationSettingsApplication struct {
	// The Assertion Consumer Service (ACS) URL of the source org (for `SAML_2_0` sign-on mode)
	AcsUrl *string `json:"acsUrl,omitempty"`
	// The entity ID of the SP (for `SAML_2_0` sign-on mode)
	AudRestriction *string `json:"audRestriction,omitempty"`
	// The base URL of the target Okta org (for `SAML_2_0` sign-on mode)
	BaseUrl string `json:"baseUrl"`
	// Used to track and manage the state of the app's creation or the provisioning process between two Okta orgs
	CreationState *string `json:"creationState,omitempty"`
	// Indicates that you don't want to use an email address as the username
	PreferUsernameOverEmail *bool `json:"preferUsernameOverEmail,omitempty"`
	// An API token from the target org that's used to secure the connection between the orgs
	Token *string `json:"token,omitempty"`
	// Encrypted token to enhance security
	TokenEncrypted       *string `json:"tokenEncrypted,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Org2OrgApplicationSettingsApplication Org2OrgApplicationSettingsApplication

// NewOrg2OrgApplicationSettingsApplication instantiates a new Org2OrgApplicationSettingsApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrg2OrgApplicationSettingsApplication(baseUrl string) *Org2OrgApplicationSettingsApplication {
	this := Org2OrgApplicationSettingsApplication{}
	this.BaseUrl = baseUrl
	return &this
}

// NewOrg2OrgApplicationSettingsApplicationWithDefaults instantiates a new Org2OrgApplicationSettingsApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrg2OrgApplicationSettingsApplicationWithDefaults() *Org2OrgApplicationSettingsApplication {
	this := Org2OrgApplicationSettingsApplication{}
	return &this
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettingsApplication) GetAcsUrl() string {
	if o == nil || IsNil(o.AcsUrl) {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetAcsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AcsUrl) {
		return nil, false
	}
	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettingsApplication) HasAcsUrl() bool {
	if o != nil && !IsNil(o.AcsUrl) {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *Org2OrgApplicationSettingsApplication) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAudRestriction returns the AudRestriction field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettingsApplication) GetAudRestriction() string {
	if o == nil || IsNil(o.AudRestriction) {
		var ret string
		return ret
	}
	return *o.AudRestriction
}

// GetAudRestrictionOk returns a tuple with the AudRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetAudRestrictionOk() (*string, bool) {
	if o == nil || IsNil(o.AudRestriction) {
		return nil, false
	}
	return o.AudRestriction, true
}

// HasAudRestriction returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettingsApplication) HasAudRestriction() bool {
	if o != nil && !IsNil(o.AudRestriction) {
		return true
	}

	return false
}

// SetAudRestriction gets a reference to the given string and assigns it to the AudRestriction field.
func (o *Org2OrgApplicationSettingsApplication) SetAudRestriction(v string) {
	o.AudRestriction = &v
}

// GetBaseUrl returns the BaseUrl field value
func (o *Org2OrgApplicationSettingsApplication) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *Org2OrgApplicationSettingsApplication) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetCreationState returns the CreationState field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettingsApplication) GetCreationState() string {
	if o == nil || IsNil(o.CreationState) {
		var ret string
		return ret
	}
	return *o.CreationState
}

// GetCreationStateOk returns a tuple with the CreationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetCreationStateOk() (*string, bool) {
	if o == nil || IsNil(o.CreationState) {
		return nil, false
	}
	return o.CreationState, true
}

// HasCreationState returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettingsApplication) HasCreationState() bool {
	if o != nil && !IsNil(o.CreationState) {
		return true
	}

	return false
}

// SetCreationState gets a reference to the given string and assigns it to the CreationState field.
func (o *Org2OrgApplicationSettingsApplication) SetCreationState(v string) {
	o.CreationState = &v
}

// GetPreferUsernameOverEmail returns the PreferUsernameOverEmail field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettingsApplication) GetPreferUsernameOverEmail() bool {
	if o == nil || IsNil(o.PreferUsernameOverEmail) {
		var ret bool
		return ret
	}
	return *o.PreferUsernameOverEmail
}

// GetPreferUsernameOverEmailOk returns a tuple with the PreferUsernameOverEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetPreferUsernameOverEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.PreferUsernameOverEmail) {
		return nil, false
	}
	return o.PreferUsernameOverEmail, true
}

// HasPreferUsernameOverEmail returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettingsApplication) HasPreferUsernameOverEmail() bool {
	if o != nil && !IsNil(o.PreferUsernameOverEmail) {
		return true
	}

	return false
}

// SetPreferUsernameOverEmail gets a reference to the given bool and assigns it to the PreferUsernameOverEmail field.
func (o *Org2OrgApplicationSettingsApplication) SetPreferUsernameOverEmail(v bool) {
	o.PreferUsernameOverEmail = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettingsApplication) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettingsApplication) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *Org2OrgApplicationSettingsApplication) SetToken(v string) {
	o.Token = &v
}

// GetTokenEncrypted returns the TokenEncrypted field value if set, zero value otherwise.
func (o *Org2OrgApplicationSettingsApplication) GetTokenEncrypted() string {
	if o == nil || IsNil(o.TokenEncrypted) {
		var ret string
		return ret
	}
	return *o.TokenEncrypted
}

// GetTokenEncryptedOk returns a tuple with the TokenEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Org2OrgApplicationSettingsApplication) GetTokenEncryptedOk() (*string, bool) {
	if o == nil || IsNil(o.TokenEncrypted) {
		return nil, false
	}
	return o.TokenEncrypted, true
}

// HasTokenEncrypted returns a boolean if a field has been set.
func (o *Org2OrgApplicationSettingsApplication) HasTokenEncrypted() bool {
	if o != nil && !IsNil(o.TokenEncrypted) {
		return true
	}

	return false
}

// SetTokenEncrypted gets a reference to the given string and assigns it to the TokenEncrypted field.
func (o *Org2OrgApplicationSettingsApplication) SetTokenEncrypted(v string) {
	o.TokenEncrypted = &v
}

func (o Org2OrgApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Org2OrgApplicationSettingsApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcsUrl) {
		toSerialize["acsUrl"] = o.AcsUrl
	}
	if !IsNil(o.AudRestriction) {
		toSerialize["audRestriction"] = o.AudRestriction
	}
	toSerialize["baseUrl"] = o.BaseUrl
	if !IsNil(o.CreationState) {
		toSerialize["creationState"] = o.CreationState
	}
	if !IsNil(o.PreferUsernameOverEmail) {
		toSerialize["preferUsernameOverEmail"] = o.PreferUsernameOverEmail
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TokenEncrypted) {
		toSerialize["tokenEncrypted"] = o.TokenEncrypted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Org2OrgApplicationSettingsApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"baseUrl",
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

	varOrg2OrgApplicationSettingsApplication := _Org2OrgApplicationSettingsApplication{}

	err = json.Unmarshal(data, &varOrg2OrgApplicationSettingsApplication)

	if err != nil {
		return err
	}

	*o = Org2OrgApplicationSettingsApplication(varOrg2OrgApplicationSettingsApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acsUrl")
		delete(additionalProperties, "audRestriction")
		delete(additionalProperties, "baseUrl")
		delete(additionalProperties, "creationState")
		delete(additionalProperties, "preferUsernameOverEmail")
		delete(additionalProperties, "token")
		delete(additionalProperties, "tokenEncrypted")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrg2OrgApplicationSettingsApplication struct {
	value *Org2OrgApplicationSettingsApplication
	isSet bool
}

func (v NullableOrg2OrgApplicationSettingsApplication) Get() *Org2OrgApplicationSettingsApplication {
	return v.value
}

func (v *NullableOrg2OrgApplicationSettingsApplication) Set(val *Org2OrgApplicationSettingsApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableOrg2OrgApplicationSettingsApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableOrg2OrgApplicationSettingsApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrg2OrgApplicationSettingsApplication(val *Org2OrgApplicationSettingsApplication) *NullableOrg2OrgApplicationSettingsApplication {
	return &NullableOrg2OrgApplicationSettingsApplication{value: val, isSet: true}
}

func (v NullableOrg2OrgApplicationSettingsApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrg2OrgApplicationSettingsApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
