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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the SamlSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlSettings{}

// SamlSettings Advanced settings for the SAML 2.0 protocol
type SamlSettings struct {
	// Determines if the IdP should persist account linking when the incoming assertion NameID format is `urn:oasis:names:tc:SAML:2.0:nameid-format:persistent`
	HonorPersistentNameId *bool `json:"honorPersistentNameId,omitempty"`
	// SAML 2.0 Name Identifier formats
	NameFormat *string `json:"nameFormat,omitempty"`
	// Set to `true` to have Okta send a logout request to the upstream IdP when a user signs out of Okta or a downstream app.
	ParticipateSlo *bool `json:"participateSlo,omitempty"`
	// Determines if the IdP should send the application context as `<OktaAppInstanceId>` and `<OktaAppName>` in the `<saml2p:Extensions>` element of the `<AuthnRequest>` message
	SendApplicationContext *bool `json:"sendApplicationContext,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _SamlSettings SamlSettings

// NewSamlSettings instantiates a new SamlSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlSettings() *SamlSettings {
	this := SamlSettings{}
	var honorPersistentNameId bool = true
	this.HonorPersistentNameId = &honorPersistentNameId
	var nameFormat string = "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"
	this.NameFormat = &nameFormat
	var sendApplicationContext bool = false
	this.SendApplicationContext = &sendApplicationContext
	return &this
}

// NewSamlSettingsWithDefaults instantiates a new SamlSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlSettingsWithDefaults() *SamlSettings {
	this := SamlSettings{}
	var honorPersistentNameId bool = true
	this.HonorPersistentNameId = &honorPersistentNameId
	var nameFormat string = "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"
	this.NameFormat = &nameFormat
	var sendApplicationContext bool = false
	this.SendApplicationContext = &sendApplicationContext
	return &this
}

// GetHonorPersistentNameId returns the HonorPersistentNameId field value if set, zero value otherwise.
func (o *SamlSettings) GetHonorPersistentNameId() bool {
	if o == nil || IsNil(o.HonorPersistentNameId) {
		var ret bool
		return ret
	}
	return *o.HonorPersistentNameId
}

// GetHonorPersistentNameIdOk returns a tuple with the HonorPersistentNameId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettings) GetHonorPersistentNameIdOk() (*bool, bool) {
	if o == nil || IsNil(o.HonorPersistentNameId) {
		return nil, false
	}
	return o.HonorPersistentNameId, true
}

// HasHonorPersistentNameId returns a boolean if a field has been set.
func (o *SamlSettings) HasHonorPersistentNameId() bool {
	if o != nil && !IsNil(o.HonorPersistentNameId) {
		return true
	}

	return false
}

// SetHonorPersistentNameId gets a reference to the given bool and assigns it to the HonorPersistentNameId field.
func (o *SamlSettings) SetHonorPersistentNameId(v bool) {
	o.HonorPersistentNameId = &v
}

// GetNameFormat returns the NameFormat field value if set, zero value otherwise.
func (o *SamlSettings) GetNameFormat() string {
	if o == nil || IsNil(o.NameFormat) {
		var ret string
		return ret
	}
	return *o.NameFormat
}

// GetNameFormatOk returns a tuple with the NameFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettings) GetNameFormatOk() (*string, bool) {
	if o == nil || IsNil(o.NameFormat) {
		return nil, false
	}
	return o.NameFormat, true
}

// HasNameFormat returns a boolean if a field has been set.
func (o *SamlSettings) HasNameFormat() bool {
	if o != nil && !IsNil(o.NameFormat) {
		return true
	}

	return false
}

// SetNameFormat gets a reference to the given string and assigns it to the NameFormat field.
func (o *SamlSettings) SetNameFormat(v string) {
	o.NameFormat = &v
}

// GetParticipateSlo returns the ParticipateSlo field value if set, zero value otherwise.
func (o *SamlSettings) GetParticipateSlo() bool {
	if o == nil || IsNil(o.ParticipateSlo) {
		var ret bool
		return ret
	}
	return *o.ParticipateSlo
}

// GetParticipateSloOk returns a tuple with the ParticipateSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettings) GetParticipateSloOk() (*bool, bool) {
	if o == nil || IsNil(o.ParticipateSlo) {
		return nil, false
	}
	return o.ParticipateSlo, true
}

// HasParticipateSlo returns a boolean if a field has been set.
func (o *SamlSettings) HasParticipateSlo() bool {
	if o != nil && !IsNil(o.ParticipateSlo) {
		return true
	}

	return false
}

// SetParticipateSlo gets a reference to the given bool and assigns it to the ParticipateSlo field.
func (o *SamlSettings) SetParticipateSlo(v bool) {
	o.ParticipateSlo = &v
}

// GetSendApplicationContext returns the SendApplicationContext field value if set, zero value otherwise.
func (o *SamlSettings) GetSendApplicationContext() bool {
	if o == nil || IsNil(o.SendApplicationContext) {
		var ret bool
		return ret
	}
	return *o.SendApplicationContext
}

// GetSendApplicationContextOk returns a tuple with the SendApplicationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSettings) GetSendApplicationContextOk() (*bool, bool) {
	if o == nil || IsNil(o.SendApplicationContext) {
		return nil, false
	}
	return o.SendApplicationContext, true
}

// HasSendApplicationContext returns a boolean if a field has been set.
func (o *SamlSettings) HasSendApplicationContext() bool {
	if o != nil && !IsNil(o.SendApplicationContext) {
		return true
	}

	return false
}

// SetSendApplicationContext gets a reference to the given bool and assigns it to the SendApplicationContext field.
func (o *SamlSettings) SetSendApplicationContext(v bool) {
	o.SendApplicationContext = &v
}

func (o SamlSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HonorPersistentNameId) {
		toSerialize["honorPersistentNameId"] = o.HonorPersistentNameId
	}
	if !IsNil(o.NameFormat) {
		toSerialize["nameFormat"] = o.NameFormat
	}
	if !IsNil(o.ParticipateSlo) {
		toSerialize["participateSlo"] = o.ParticipateSlo
	}
	if !IsNil(o.SendApplicationContext) {
		toSerialize["sendApplicationContext"] = o.SendApplicationContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlSettings) UnmarshalJSON(data []byte) (err error) {
	varSamlSettings := _SamlSettings{}

	err = json.Unmarshal(data, &varSamlSettings)

	if err != nil {
		return err
	}

	*o = SamlSettings(varSamlSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "honorPersistentNameId")
		delete(additionalProperties, "nameFormat")
		delete(additionalProperties, "participateSlo")
		delete(additionalProperties, "sendApplicationContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlSettings struct {
	value *SamlSettings
	isSet bool
}

func (v NullableSamlSettings) Get() *SamlSettings {
	return v.value
}

func (v *NullableSamlSettings) Set(val *SamlSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlSettings(val *SamlSettings) *NullableSamlSettings {
	return &NullableSamlSettings{value: val, isSet: true}
}

func (v NullableSamlSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
