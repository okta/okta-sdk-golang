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

// TokenAuthorizationServerPolicyRuleAction struct for TokenAuthorizationServerPolicyRuleAction
type TokenAuthorizationServerPolicyRuleAction struct {
	// Lifetime of the access token in minutes. The minimum is five minutes. The maximum is one day.
	AccessTokenLifetimeMinutes *int32 `json:"accessTokenLifetimeMinutes,omitempty"`
	InlineHook *TokenAuthorizationServerPolicyRuleActionInlineHook `json:"inlineHook,omitempty"`
	// Lifetime of the refresh token is the minimum access token lifetime.
	RefreshTokenLifetimeMinutes *int32 `json:"refreshTokenLifetimeMinutes,omitempty"`
	// Timeframe when the refresh token is valid. The minimum is 10 minutes. The maximum is five years (2,628,000 minutes).
	RefreshTokenWindowMinutes *int32 `json:"refreshTokenWindowMinutes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TokenAuthorizationServerPolicyRuleAction TokenAuthorizationServerPolicyRuleAction

// NewTokenAuthorizationServerPolicyRuleAction instantiates a new TokenAuthorizationServerPolicyRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAuthorizationServerPolicyRuleAction() *TokenAuthorizationServerPolicyRuleAction {
	this := TokenAuthorizationServerPolicyRuleAction{}
	return &this
}

// NewTokenAuthorizationServerPolicyRuleActionWithDefaults instantiates a new TokenAuthorizationServerPolicyRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAuthorizationServerPolicyRuleActionWithDefaults() *TokenAuthorizationServerPolicyRuleAction {
	this := TokenAuthorizationServerPolicyRuleAction{}
	return &this
}

// GetAccessTokenLifetimeMinutes returns the AccessTokenLifetimeMinutes field value if set, zero value otherwise.
func (o *TokenAuthorizationServerPolicyRuleAction) GetAccessTokenLifetimeMinutes() int32 {
	if o == nil || o.AccessTokenLifetimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.AccessTokenLifetimeMinutes
}

// GetAccessTokenLifetimeMinutesOk returns a tuple with the AccessTokenLifetimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) GetAccessTokenLifetimeMinutesOk() (*int32, bool) {
	if o == nil || o.AccessTokenLifetimeMinutes == nil {
		return nil, false
	}
	return o.AccessTokenLifetimeMinutes, true
}

// HasAccessTokenLifetimeMinutes returns a boolean if a field has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) HasAccessTokenLifetimeMinutes() bool {
	if o != nil && o.AccessTokenLifetimeMinutes != nil {
		return true
	}

	return false
}

// SetAccessTokenLifetimeMinutes gets a reference to the given int32 and assigns it to the AccessTokenLifetimeMinutes field.
func (o *TokenAuthorizationServerPolicyRuleAction) SetAccessTokenLifetimeMinutes(v int32) {
	o.AccessTokenLifetimeMinutes = &v
}

// GetInlineHook returns the InlineHook field value if set, zero value otherwise.
func (o *TokenAuthorizationServerPolicyRuleAction) GetInlineHook() TokenAuthorizationServerPolicyRuleActionInlineHook {
	if o == nil || o.InlineHook == nil {
		var ret TokenAuthorizationServerPolicyRuleActionInlineHook
		return ret
	}
	return *o.InlineHook
}

// GetInlineHookOk returns a tuple with the InlineHook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) GetInlineHookOk() (*TokenAuthorizationServerPolicyRuleActionInlineHook, bool) {
	if o == nil || o.InlineHook == nil {
		return nil, false
	}
	return o.InlineHook, true
}

// HasInlineHook returns a boolean if a field has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) HasInlineHook() bool {
	if o != nil && o.InlineHook != nil {
		return true
	}

	return false
}

// SetInlineHook gets a reference to the given TokenAuthorizationServerPolicyRuleActionInlineHook and assigns it to the InlineHook field.
func (o *TokenAuthorizationServerPolicyRuleAction) SetInlineHook(v TokenAuthorizationServerPolicyRuleActionInlineHook) {
	o.InlineHook = &v
}

// GetRefreshTokenLifetimeMinutes returns the RefreshTokenLifetimeMinutes field value if set, zero value otherwise.
func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenLifetimeMinutes() int32 {
	if o == nil || o.RefreshTokenLifetimeMinutes == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenLifetimeMinutes
}

// GetRefreshTokenLifetimeMinutesOk returns a tuple with the RefreshTokenLifetimeMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenLifetimeMinutesOk() (*int32, bool) {
	if o == nil || o.RefreshTokenLifetimeMinutes == nil {
		return nil, false
	}
	return o.RefreshTokenLifetimeMinutes, true
}

// HasRefreshTokenLifetimeMinutes returns a boolean if a field has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) HasRefreshTokenLifetimeMinutes() bool {
	if o != nil && o.RefreshTokenLifetimeMinutes != nil {
		return true
	}

	return false
}

// SetRefreshTokenLifetimeMinutes gets a reference to the given int32 and assigns it to the RefreshTokenLifetimeMinutes field.
func (o *TokenAuthorizationServerPolicyRuleAction) SetRefreshTokenLifetimeMinutes(v int32) {
	o.RefreshTokenLifetimeMinutes = &v
}

// GetRefreshTokenWindowMinutes returns the RefreshTokenWindowMinutes field value if set, zero value otherwise.
func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenWindowMinutes() int32 {
	if o == nil || o.RefreshTokenWindowMinutes == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokenWindowMinutes
}

// GetRefreshTokenWindowMinutesOk returns a tuple with the RefreshTokenWindowMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) GetRefreshTokenWindowMinutesOk() (*int32, bool) {
	if o == nil || o.RefreshTokenWindowMinutes == nil {
		return nil, false
	}
	return o.RefreshTokenWindowMinutes, true
}

// HasRefreshTokenWindowMinutes returns a boolean if a field has been set.
func (o *TokenAuthorizationServerPolicyRuleAction) HasRefreshTokenWindowMinutes() bool {
	if o != nil && o.RefreshTokenWindowMinutes != nil {
		return true
	}

	return false
}

// SetRefreshTokenWindowMinutes gets a reference to the given int32 and assigns it to the RefreshTokenWindowMinutes field.
func (o *TokenAuthorizationServerPolicyRuleAction) SetRefreshTokenWindowMinutes(v int32) {
	o.RefreshTokenWindowMinutes = &v
}

func (o TokenAuthorizationServerPolicyRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessTokenLifetimeMinutes != nil {
		toSerialize["accessTokenLifetimeMinutes"] = o.AccessTokenLifetimeMinutes
	}
	if o.InlineHook != nil {
		toSerialize["inlineHook"] = o.InlineHook
	}
	if o.RefreshTokenLifetimeMinutes != nil {
		toSerialize["refreshTokenLifetimeMinutes"] = o.RefreshTokenLifetimeMinutes
	}
	if o.RefreshTokenWindowMinutes != nil {
		toSerialize["refreshTokenWindowMinutes"] = o.RefreshTokenWindowMinutes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TokenAuthorizationServerPolicyRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	varTokenAuthorizationServerPolicyRuleAction := _TokenAuthorizationServerPolicyRuleAction{}

	err = json.Unmarshal(bytes, &varTokenAuthorizationServerPolicyRuleAction)
	if err == nil {
		*o = TokenAuthorizationServerPolicyRuleAction(varTokenAuthorizationServerPolicyRuleAction)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accessTokenLifetimeMinutes")
		delete(additionalProperties, "inlineHook")
		delete(additionalProperties, "refreshTokenLifetimeMinutes")
		delete(additionalProperties, "refreshTokenWindowMinutes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTokenAuthorizationServerPolicyRuleAction struct {
	value *TokenAuthorizationServerPolicyRuleAction
	isSet bool
}

func (v NullableTokenAuthorizationServerPolicyRuleAction) Get() *TokenAuthorizationServerPolicyRuleAction {
	return v.value
}

func (v *NullableTokenAuthorizationServerPolicyRuleAction) Set(val *TokenAuthorizationServerPolicyRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAuthorizationServerPolicyRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAuthorizationServerPolicyRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAuthorizationServerPolicyRuleAction(val *TokenAuthorizationServerPolicyRuleAction) *NullableTokenAuthorizationServerPolicyRuleAction {
	return &NullableTokenAuthorizationServerPolicyRuleAction{value: val, isSet: true}
}

func (v NullableTokenAuthorizationServerPolicyRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAuthorizationServerPolicyRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

