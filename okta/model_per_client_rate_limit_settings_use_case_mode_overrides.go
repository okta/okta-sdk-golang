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

// checks if the PerClientRateLimitSettingsUseCaseModeOverrides type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerClientRateLimitSettingsUseCaseModeOverrides{}

// PerClientRateLimitSettingsUseCaseModeOverrides A map of Per-Client Rate Limit Use Case to the applicable PerClientRateLimitMode. Overrides the `defaultMode` property for the specified use cases.
type PerClientRateLimitSettingsUseCaseModeOverrides struct {
	LOGIN_PAGE           *string `json:"LOGIN_PAGE,omitempty"`
	OAUTH2AUTHORIZE      *string `json:"OAUTH2_AUTHORIZE,omitempty"`
	OIE_APP_INTENT       *string `json:"OIE_APP_INTENT,omitempty"`
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
	if o == nil || IsNil(o.LOGIN_PAGE) {
		var ret string
		return ret
	}
	return *o.LOGIN_PAGE
}

// GetLOGIN_PAGEOk returns a tuple with the LOGIN_PAGE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetLOGIN_PAGEOk() (*string, bool) {
	if o == nil || IsNil(o.LOGIN_PAGE) {
		return nil, false
	}
	return o.LOGIN_PAGE, true
}

// HasLOGIN_PAGE returns a boolean if a field has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) HasLOGIN_PAGE() bool {
	if o != nil && !IsNil(o.LOGIN_PAGE) {
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
	if o == nil || IsNil(o.OAUTH2AUTHORIZE) {
		var ret string
		return ret
	}
	return *o.OAUTH2AUTHORIZE
}

// GetOAUTH2AUTHORIZEOk returns a tuple with the OAUTH2AUTHORIZE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetOAUTH2AUTHORIZEOk() (*string, bool) {
	if o == nil || IsNil(o.OAUTH2AUTHORIZE) {
		return nil, false
	}
	return o.OAUTH2AUTHORIZE, true
}

// HasOAUTH2AUTHORIZE returns a boolean if a field has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) HasOAUTH2AUTHORIZE() bool {
	if o != nil && !IsNil(o.OAUTH2AUTHORIZE) {
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
	if o == nil || IsNil(o.OIE_APP_INTENT) {
		var ret string
		return ret
	}
	return *o.OIE_APP_INTENT
}

// GetOIE_APP_INTENTOk returns a tuple with the OIE_APP_INTENT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) GetOIE_APP_INTENTOk() (*string, bool) {
	if o == nil || IsNil(o.OIE_APP_INTENT) {
		return nil, false
	}
	return o.OIE_APP_INTENT, true
}

// HasOIE_APP_INTENT returns a boolean if a field has been set.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) HasOIE_APP_INTENT() bool {
	if o != nil && !IsNil(o.OIE_APP_INTENT) {
		return true
	}

	return false
}

// SetOIE_APP_INTENT gets a reference to the given string and assigns it to the OIE_APP_INTENT field.
func (o *PerClientRateLimitSettingsUseCaseModeOverrides) SetOIE_APP_INTENT(v string) {
	o.OIE_APP_INTENT = &v
}

func (o PerClientRateLimitSettingsUseCaseModeOverrides) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerClientRateLimitSettingsUseCaseModeOverrides) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LOGIN_PAGE) {
		toSerialize["LOGIN_PAGE"] = o.LOGIN_PAGE
	}
	if !IsNil(o.OAUTH2AUTHORIZE) {
		toSerialize["OAUTH2_AUTHORIZE"] = o.OAUTH2AUTHORIZE
	}
	if !IsNil(o.OIE_APP_INTENT) {
		toSerialize["OIE_APP_INTENT"] = o.OIE_APP_INTENT
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PerClientRateLimitSettingsUseCaseModeOverrides) UnmarshalJSON(data []byte) (err error) {
	varPerClientRateLimitSettingsUseCaseModeOverrides := _PerClientRateLimitSettingsUseCaseModeOverrides{}

	err = json.Unmarshal(data, &varPerClientRateLimitSettingsUseCaseModeOverrides)

	if err != nil {
		return err
	}

	*o = PerClientRateLimitSettingsUseCaseModeOverrides(varPerClientRateLimitSettingsUseCaseModeOverrides)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "LOGIN_PAGE")
		delete(additionalProperties, "OAUTH2_AUTHORIZE")
		delete(additionalProperties, "OIE_APP_INTENT")
		o.AdditionalProperties = additionalProperties
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
