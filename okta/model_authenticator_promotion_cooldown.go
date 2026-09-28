/*
Okta Admin Management API

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

// checks if the AuthenticatorPromotionCooldown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorPromotionCooldown{}

// AuthenticatorPromotionCooldown <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Specifies the cadence between prompts within the promotion period.
type AuthenticatorPromotionCooldown struct {
	// The ISO 8601 period between prompts. Only month (`M`), week (`W`), and day (`D`) components are allowed (for example, `P30D`, `P2W`, or `P1M`). Year components and time components (hours, minutes, seconds) aren't permitted: a period that includes a year or time section is rejected even when that section is zero (for example, `P0Y`, `PT0H`, and `P1DT0H` are all invalid). The period must be positive and non-zero (`P0D` is invalid); to prompt on every eligible sign-in, use `type: BY_SIGN_IN` instead. Required when `type` is `BY_DURATION` and must be omitted when `type` is `BY_SIGN_IN`.
	Duration *string `json:"duration,omitempty"`
	// The cadence driver between prompts.  * `BY_DURATION`: Re-prompts after `duration` has elapsed since the user's last skip. * `BY_SIGN_IN`: Re-prompts on every eligible sign-in. No `duration` is used.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorPromotionCooldown AuthenticatorPromotionCooldown

// NewAuthenticatorPromotionCooldown instantiates a new AuthenticatorPromotionCooldown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorPromotionCooldown(type_ string) *AuthenticatorPromotionCooldown {
	this := AuthenticatorPromotionCooldown{}
	this.Type = type_
	return &this
}

// NewAuthenticatorPromotionCooldownWithDefaults instantiates a new AuthenticatorPromotionCooldown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorPromotionCooldownWithDefaults() *AuthenticatorPromotionCooldown {
	this := AuthenticatorPromotionCooldown{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *AuthenticatorPromotionCooldown) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorPromotionCooldown) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *AuthenticatorPromotionCooldown) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *AuthenticatorPromotionCooldown) SetDuration(v string) {
	o.Duration = &v
}

// GetType returns the Type field value
func (o *AuthenticatorPromotionCooldown) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AuthenticatorPromotionCooldown) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AuthenticatorPromotionCooldown) SetType(v string) {
	o.Type = v
}

func (o AuthenticatorPromotionCooldown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorPromotionCooldown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorPromotionCooldown) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varAuthenticatorPromotionCooldown := _AuthenticatorPromotionCooldown{}

	err = json.Unmarshal(data, &varAuthenticatorPromotionCooldown)

	if err != nil {
		return err
	}

	*o = AuthenticatorPromotionCooldown(varAuthenticatorPromotionCooldown)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "duration")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorPromotionCooldown struct {
	value *AuthenticatorPromotionCooldown
	isSet bool
}

func (v NullableAuthenticatorPromotionCooldown) Get() *AuthenticatorPromotionCooldown {
	return v.value
}

func (v *NullableAuthenticatorPromotionCooldown) Set(val *AuthenticatorPromotionCooldown) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorPromotionCooldown) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorPromotionCooldown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorPromotionCooldown(val *AuthenticatorPromotionCooldown) *NullableAuthenticatorPromotionCooldown {
	return &NullableAuthenticatorPromotionCooldown{value: val, isSet: true}
}

func (v NullableAuthenticatorPromotionCooldown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorPromotionCooldown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
