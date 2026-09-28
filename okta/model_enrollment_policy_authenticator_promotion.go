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

// checks if the EnrollmentPolicyAuthenticatorPromotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollmentPolicyAuthenticatorPromotion{}

// EnrollmentPolicyAuthenticatorPromotion <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>Configures a non-blocking passkey enrollment promotion (a soft nudge) for the authenticator. The promotion prompts eligible users to enroll a passkey on sign-in but never forces enrollment. It's only valid on the `webauthn` authenticator and only when [`self`](/openapi/okta-management/management/tags/policy/other/createpolicy#other/createpolicy/t=request&path=&d=1/settings/authenticators/enroll/self) is `OPTIONAL`.
type EnrollmentPolicyAuthenticatorPromotion struct {
	Cooldown AuthenticatorPromotionCooldown `json:"cooldown"`
	// The number of times the user can skip the nudge before prompting stops permanently. A value of `0` means the nudge is never limited by skip count.
	SkipCount *int32 `json:"skipCount,omitempty"`
	// The promotion period type. Defines when prompting stops permanently for a user.  * `BY_SKIP_COUNT`: Prompting stops after the user skips the nudge `skipCount` times.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _EnrollmentPolicyAuthenticatorPromotion EnrollmentPolicyAuthenticatorPromotion

// NewEnrollmentPolicyAuthenticatorPromotion instantiates a new EnrollmentPolicyAuthenticatorPromotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollmentPolicyAuthenticatorPromotion(cooldown AuthenticatorPromotionCooldown, type_ string) *EnrollmentPolicyAuthenticatorPromotion {
	this := EnrollmentPolicyAuthenticatorPromotion{}
	this.Cooldown = cooldown
	this.Type = type_
	return &this
}

// NewEnrollmentPolicyAuthenticatorPromotionWithDefaults instantiates a new EnrollmentPolicyAuthenticatorPromotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollmentPolicyAuthenticatorPromotionWithDefaults() *EnrollmentPolicyAuthenticatorPromotion {
	this := EnrollmentPolicyAuthenticatorPromotion{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *EnrollmentPolicyAuthenticatorPromotion) GetCooldown() AuthenticatorPromotionCooldown {
	if o == nil {
		var ret AuthenticatorPromotionCooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *EnrollmentPolicyAuthenticatorPromotion) GetCooldownOk() (*AuthenticatorPromotionCooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *EnrollmentPolicyAuthenticatorPromotion) SetCooldown(v AuthenticatorPromotionCooldown) {
	o.Cooldown = v
}

// GetSkipCount returns the SkipCount field value if set, zero value otherwise.
func (o *EnrollmentPolicyAuthenticatorPromotion) GetSkipCount() int32 {
	if o == nil || IsNil(o.SkipCount) {
		var ret int32
		return ret
	}
	return *o.SkipCount
}

// GetSkipCountOk returns a tuple with the SkipCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollmentPolicyAuthenticatorPromotion) GetSkipCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SkipCount) {
		return nil, false
	}
	return o.SkipCount, true
}

// HasSkipCount returns a boolean if a field has been set.
func (o *EnrollmentPolicyAuthenticatorPromotion) HasSkipCount() bool {
	if o != nil && !IsNil(o.SkipCount) {
		return true
	}

	return false
}

// SetSkipCount gets a reference to the given int32 and assigns it to the SkipCount field.
func (o *EnrollmentPolicyAuthenticatorPromotion) SetSkipCount(v int32) {
	o.SkipCount = &v
}

// GetType returns the Type field value
func (o *EnrollmentPolicyAuthenticatorPromotion) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EnrollmentPolicyAuthenticatorPromotion) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EnrollmentPolicyAuthenticatorPromotion) SetType(v string) {
	o.Type = v
}

func (o EnrollmentPolicyAuthenticatorPromotion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollmentPolicyAuthenticatorPromotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	if !IsNil(o.SkipCount) {
		toSerialize["skipCount"] = o.SkipCount
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnrollmentPolicyAuthenticatorPromotion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
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

	varEnrollmentPolicyAuthenticatorPromotion := _EnrollmentPolicyAuthenticatorPromotion{}

	err = json.Unmarshal(data, &varEnrollmentPolicyAuthenticatorPromotion)

	if err != nil {
		return err
	}

	*o = EnrollmentPolicyAuthenticatorPromotion(varEnrollmentPolicyAuthenticatorPromotion)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cooldown")
		delete(additionalProperties, "skipCount")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEnrollmentPolicyAuthenticatorPromotion struct {
	value *EnrollmentPolicyAuthenticatorPromotion
	isSet bool
}

func (v NullableEnrollmentPolicyAuthenticatorPromotion) Get() *EnrollmentPolicyAuthenticatorPromotion {
	return v.value
}

func (v *NullableEnrollmentPolicyAuthenticatorPromotion) Set(val *EnrollmentPolicyAuthenticatorPromotion) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollmentPolicyAuthenticatorPromotion) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollmentPolicyAuthenticatorPromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollmentPolicyAuthenticatorPromotion(val *EnrollmentPolicyAuthenticatorPromotion) *NullableEnrollmentPolicyAuthenticatorPromotion {
	return &NullableEnrollmentPolicyAuthenticatorPromotion{value: val, isSet: true}
}

func (v NullableEnrollmentPolicyAuthenticatorPromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollmentPolicyAuthenticatorPromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
