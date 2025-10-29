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

// checks if the OidcSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OidcSettings{}

// OidcSettings Advanced settings for the OpenID Connect protocol
type OidcSettings struct {
	// Set to `true` to have Okta send a logout request to the upstream IdP when a user signs out of Okta or a downstream app.
	ParticipateSlo *bool `json:"participateSlo,omitempty"`
	// Determines if the IdP should send the application context as `OktaAppInstanceId` and `OktaAppName` params in the request
	SendApplicationContext *bool `json:"sendApplicationContext,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _OidcSettings OidcSettings

// NewOidcSettings instantiates a new OidcSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOidcSettings() *OidcSettings {
	this := OidcSettings{}
	var sendApplicationContext bool = false
	this.SendApplicationContext = &sendApplicationContext
	return &this
}

// NewOidcSettingsWithDefaults instantiates a new OidcSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOidcSettingsWithDefaults() *OidcSettings {
	this := OidcSettings{}
	var sendApplicationContext bool = false
	this.SendApplicationContext = &sendApplicationContext
	return &this
}

// GetParticipateSlo returns the ParticipateSlo field value if set, zero value otherwise.
func (o *OidcSettings) GetParticipateSlo() bool {
	if o == nil || IsNil(o.ParticipateSlo) {
		var ret bool
		return ret
	}
	return *o.ParticipateSlo
}

// GetParticipateSloOk returns a tuple with the ParticipateSlo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettings) GetParticipateSloOk() (*bool, bool) {
	if o == nil || IsNil(o.ParticipateSlo) {
		return nil, false
	}
	return o.ParticipateSlo, true
}

// HasParticipateSlo returns a boolean if a field has been set.
func (o *OidcSettings) HasParticipateSlo() bool {
	if o != nil && !IsNil(o.ParticipateSlo) {
		return true
	}

	return false
}

// SetParticipateSlo gets a reference to the given bool and assigns it to the ParticipateSlo field.
func (o *OidcSettings) SetParticipateSlo(v bool) {
	o.ParticipateSlo = &v
}

// GetSendApplicationContext returns the SendApplicationContext field value if set, zero value otherwise.
func (o *OidcSettings) GetSendApplicationContext() bool {
	if o == nil || IsNil(o.SendApplicationContext) {
		var ret bool
		return ret
	}
	return *o.SendApplicationContext
}

// GetSendApplicationContextOk returns a tuple with the SendApplicationContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OidcSettings) GetSendApplicationContextOk() (*bool, bool) {
	if o == nil || IsNil(o.SendApplicationContext) {
		return nil, false
	}
	return o.SendApplicationContext, true
}

// HasSendApplicationContext returns a boolean if a field has been set.
func (o *OidcSettings) HasSendApplicationContext() bool {
	if o != nil && !IsNil(o.SendApplicationContext) {
		return true
	}

	return false
}

// SetSendApplicationContext gets a reference to the given bool and assigns it to the SendApplicationContext field.
func (o *OidcSettings) SetSendApplicationContext(v bool) {
	o.SendApplicationContext = &v
}

func (o OidcSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OidcSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *OidcSettings) UnmarshalJSON(data []byte) (err error) {
	varOidcSettings := _OidcSettings{}

	err = json.Unmarshal(data, &varOidcSettings)

	if err != nil {
		return err
	}

	*o = OidcSettings(varOidcSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "participateSlo")
		delete(additionalProperties, "sendApplicationContext")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOidcSettings struct {
	value *OidcSettings
	isSet bool
}

func (v NullableOidcSettings) Get() *OidcSettings {
	return v.value
}

func (v *NullableOidcSettings) Set(val *OidcSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOidcSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOidcSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOidcSettings(val *OidcSettings) *NullableOidcSettings {
	return &NullableOidcSettings{value: val, isSet: true}
}

func (v NullableOidcSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOidcSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
