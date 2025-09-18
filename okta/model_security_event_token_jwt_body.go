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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the SecurityEventTokenJwtBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityEventTokenJwtBody{}

// SecurityEventTokenJwtBody JSON Web Token body payload for a Security Event Token sent by the SSF Transmitter. For examples and more information, see [SSF Transmitter SET payload structures](https://developer.okta.com/docs/reference/ssf-transmitter-sets).
type SecurityEventTokenJwtBody struct {
	// Audience
	Aud    string                      `json:"aud"`
	Events SecurityEventTokenJwtEvents `json:"events"`
	// Token issue time (UNIX timestamp)
	Iat int64 `json:"iat"`
	// Token issuer
	Iss string `json:"iss"`
	// Token ID
	Jti                  string `json:"jti"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventTokenJwtBody SecurityEventTokenJwtBody

// NewSecurityEventTokenJwtBody instantiates a new SecurityEventTokenJwtBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenJwtBody(aud string, events SecurityEventTokenJwtEvents, iat int64, iss string, jti string) *SecurityEventTokenJwtBody {
	this := SecurityEventTokenJwtBody{}
	this.Aud = aud
	this.Events = events
	this.Iat = iat
	this.Iss = iss
	this.Jti = jti
	return &this
}

// NewSecurityEventTokenJwtBodyWithDefaults instantiates a new SecurityEventTokenJwtBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenJwtBodyWithDefaults() *SecurityEventTokenJwtBody {
	this := SecurityEventTokenJwtBody{}
	return &this
}

// GetAud returns the Aud field value
func (o *SecurityEventTokenJwtBody) GetAud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtBody) GetAudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aud, true
}

// SetAud sets field value
func (o *SecurityEventTokenJwtBody) SetAud(v string) {
	o.Aud = v
}

// GetEvents returns the Events field value
func (o *SecurityEventTokenJwtBody) GetEvents() SecurityEventTokenJwtEvents {
	if o == nil {
		var ret SecurityEventTokenJwtEvents
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtBody) GetEventsOk() (*SecurityEventTokenJwtEvents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *SecurityEventTokenJwtBody) SetEvents(v SecurityEventTokenJwtEvents) {
	o.Events = v
}

// GetIat returns the Iat field value
func (o *SecurityEventTokenJwtBody) GetIat() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Iat
}

// GetIatOk returns a tuple with the Iat field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtBody) GetIatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iat, true
}

// SetIat sets field value
func (o *SecurityEventTokenJwtBody) SetIat(v int64) {
	o.Iat = v
}

// GetIss returns the Iss field value
func (o *SecurityEventTokenJwtBody) GetIss() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iss
}

// GetIssOk returns a tuple with the Iss field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtBody) GetIssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iss, true
}

// SetIss sets field value
func (o *SecurityEventTokenJwtBody) SetIss(v string) {
	o.Iss = v
}

// GetJti returns the Jti field value
func (o *SecurityEventTokenJwtBody) GetJti() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jti
}

// GetJtiOk returns a tuple with the Jti field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtBody) GetJtiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jti, true
}

// SetJti sets field value
func (o *SecurityEventTokenJwtBody) SetJti(v string) {
	o.Jti = v
}

func (o SecurityEventTokenJwtBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityEventTokenJwtBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aud"] = o.Aud
	toSerialize["events"] = o.Events
	toSerialize["iat"] = o.Iat
	toSerialize["iss"] = o.Iss
	toSerialize["jti"] = o.Jti

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityEventTokenJwtBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"aud",
		"events",
		"iat",
		"iss",
		"jti",
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

	varSecurityEventTokenJwtBody := _SecurityEventTokenJwtBody{}

	err = json.Unmarshal(data, &varSecurityEventTokenJwtBody)

	if err != nil {
		return err
	}

	*o = SecurityEventTokenJwtBody(varSecurityEventTokenJwtBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aud")
		delete(additionalProperties, "events")
		delete(additionalProperties, "iat")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "jti")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityEventTokenJwtBody struct {
	value *SecurityEventTokenJwtBody
	isSet bool
}

func (v NullableSecurityEventTokenJwtBody) Get() *SecurityEventTokenJwtBody {
	return v.value
}

func (v *NullableSecurityEventTokenJwtBody) Set(val *SecurityEventTokenJwtBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenJwtBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenJwtBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenJwtBody(val *SecurityEventTokenJwtBody) *NullableSecurityEventTokenJwtBody {
	return &NullableSecurityEventTokenJwtBody{value: val, isSet: true}
}

func (v NullableSecurityEventTokenJwtBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenJwtBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
