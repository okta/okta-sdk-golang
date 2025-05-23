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
)

// PerClientRateLimitSettingsUseCaseModeOverrides A map of Per-Client Rate Limit Use Case to the applicable PerClientRateLimitMode. Overrides the `defaultMode` property for the specified use cases.
type PerClientRateLimitSettingsUseCaseModeOverrides struct {
	LOGIN_PAGE *string `json:"LOGIN_PAGE,omitempty"`
	OAUTH2AUTHORIZE *string `json:"OAUTH2_AUTHORIZE,omitempty"`
	OIE_APP_INTENT *string `json:"OIE_APP_INTENT,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PerClientRateLimitSettingsUseCaseModeOverrides PerClientRateLimitSettingsUseCaseModeOverrides

// NewPerClientRateLimitSettingsUseCaseModeOverrides instantiates a new PerClientRateLimitSettingsUseCaseModeOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerClientRateLimitSettingsUseCaseModeOverrides() *PerClientRateLimitSettingsUseCaseModeOverrides {
	this := PerClientRateLimitSettingsUseCaseModeOverrides{}
	return &this
}

// NewPerClientRateLimitSettingsUseCaseModeOverridesWithDefaults instantiates a new PerClientRateLimitSettingsUseCaseModeOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerClientRateLimitSettingsUseCaseModeOverridesWithDefaults() *PerClientRateLimitSettingsUseCaseModeOverrides {
	this := PerClientRateLimitSettingsUseCaseModeOverrides{}
	return &this
}

// GetLOGIN_PAGE returns the LOGIN_PAGE field value if set, zero value otherwise.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetLOGIN_PAGE() string {
	if o == nil || o.LOGIN_PAGE == nil {
		var ret string
		return ret
	}
	return *o.LOGIN_PAGE
}

// GetLOGIN_PAGEOk returns a tuple with the LOGIN_PAGE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetLOGIN_PAGEOk() (*string, bool) {
	if o == nil || o.LOGIN_PAGE == nil {
		return nil, false
	}
	return o.LOGIN_PAGE, true
}

// HasLOGIN_PAGE returns a boolean if a field has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) HasLOGIN_PAGE() bool {
	if o != nil && o.LOGIN_PAGE != nil {
		return true
	}

	return false
}

// SetLOGIN_PAGE gets a reference to the given string and assigns it to the LOGIN_PAGE field.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) SetLOGIN_PAGE(v string) {
	o.LOGIN_PAGE = &v
}

// GetOAUTH2AUTHORIZE returns the OAUTH2AUTHORIZE field value if set, zero value otherwise.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetOAUTH2AUTHORIZE() string {
	if o == nil || o.OAUTH2AUTHORIZE == nil {
		var ret string
		return ret
	}
	return *o.OAUTH2AUTHORIZE
}

// GetOAUTH2AUTHORIZEOk returns a tuple with the OAUTH2AUTHORIZE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetOAUTH2AUTHORIZEOk() (*string, bool) {
	if o == nil || o.OAUTH2AUTHORIZE == nil {
		return nil, false
	}
	return o.OAUTH2AUTHORIZE, true
}

// HasOAUTH2AUTHORIZE returns a boolean if a field has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) HasOAUTH2AUTHORIZE() bool {
	if o != nil && o.OAUTH2AUTHORIZE != nil {
		return true
	}

	return false
}

// SetOAUTH2AUTHORIZE gets a reference to the given string and assigns it to the OAUTH2AUTHORIZE field.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) SetOAUTH2AUTHORIZE(v string) {
	o.OAUTH2AUTHORIZE = &v
}

// GetOIE_APP_INTENT returns the OIE_APP_INTENT field value if set, zero value otherwise.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetOIE_APP_INTENT() string {
	if o == nil || o.OIE_APP_INTENT == nil {
		var ret string
		return ret
	}
	return *o.OIE_APP_INTENT
}

// GetOIE_APP_INTENTOk returns a tuple with the OIE_APP_INTENT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetOIE_APP_INTENTOk() (*string, bool) {
	if o == nil || o.OIE_APP_INTENT == nil {
		return nil, false
	}
	return o.OIE_APP_INTENT, true
}

// HasOIE_APP_INTENT returns a boolean if a field has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) HasOIE_APP_INTENT() bool {
	if o != nil && o.OIE_APP_INTENT != nil {
		return true
	}

	return false
}

// SetOIE_APP_INTENT gets a reference to the given string and assigns it to the OIE_APP_INTENT field.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) SetOIE_APP_INTENT(v string) {
	o.OIE_APP_INTENT = &v
}

func (o PerClientRateLimitSettingsUseCaseModeOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LOGIN_PAGE != nil {
		toSerialize["LOGIN_PAGE"] = o.LOGIN_PAGE
	}
	if o.OAUTH2AUTHORIZE != nil {
		toSerialize["OAUTH2_AUTHORIZE"] = o.OAUTH2AUTHORIZE
	}
	if o.OIE_APP_INTENT != nil {
		toSerialize["OIE_APP_INTENT"] = o.OIE_APP_INTENT
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PerClientRateLimitSettingsUseCaseModeOverrides) UnmarshalJSON(bytes []byte) (err error) {
	varPerClientRateLimitSettingsUseCaseModeOverrides := _PerClientRateLimitSettingsUseCaseModeOverrides{}

	err = json.Unmarshal(bytes, &varPerClientRateLimitSettingsUseCaseModeOverrides)
	if err == nil {
		*o = PerClientRateLimitSettingsUseCaseModeOverrides(varPerClientRateLimitSettingsUseCaseModeOverrides)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "LOGIN_PAGE")
		delete(additionalProperties, "OAUTH2_AUTHORIZE")
		delete(additionalProperties, "OIE_APP_INTENT")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePerClientRateLimitSettingsUseCaseModeOverrides struct {
	value *PerClientRateLimitSettingsUseCaseModeOverrides
	isSet bool
}

func (v NullablePerClientRateLimitSettingsUseCaseModeOverrides) Get() *PerClientRateLimitSettingsUseCaseModeOverrides {
	return v.value
}

func (v *NullablePerClientRateLimitSettingsUseCaseModeOverrides) Set(val *PerClientRateLimitSettingsUseCaseModeOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullablePerClientRateLimitSettingsUseCaseModeOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullablePerClientRateLimitSettingsUseCaseModeOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerClientRateLimitSettingsUseCaseModeOverrides(val *PerClientRateLimitSettingsUseCaseModeOverrides) *NullablePerClientRateLimitSettingsUseCaseModeOverrides {
	return &NullablePerClientRateLimitSettingsUseCaseModeOverrides{value: val, isSet: true}
}

func (v NullablePerClientRateLimitSettingsUseCaseModeOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerClientRateLimitSettingsUseCaseModeOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

