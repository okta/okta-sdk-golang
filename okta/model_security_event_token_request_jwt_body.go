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

// SecurityEventTokenRequestJwtBody JSON Web Token body payload for a Security Event Token
type SecurityEventTokenRequestJwtBody struct {
	// Audience
	Aud string `json:"aud"`
	Events SecurityEventTokenRequestJwtEvents `json:"events"`
	// Token issue time (UNIX timestamp)
	Iat int64 `json:"iat"`
	// Token issuer
	Iss string `json:"iss"`
	// Token ID
	Jti string `json:"jti"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventTokenRequestJwtBody SecurityEventTokenRequestJwtBody

// NewSecurityEventTokenRequestJwtBody instantiates a new SecurityEventTokenRequestJwtBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenRequestJwtBody(aud string, events SecurityEventTokenRequestJwtEvents, iat int64, iss string, jti string) *SecurityEventTokenRequestJwtBody {
	this := SecurityEventTokenRequestJwtBody{}
	this.Aud = aud
	this.Events = events
	this.Iat = iat
	this.Iss = iss
	this.Jti = jti
	return &this
}

// NewSecurityEventTokenRequestJwtBodyWithDefaults instantiates a new SecurityEventTokenRequestJwtBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenRequestJwtBodyWithDefaults() *SecurityEventTokenRequestJwtBody {
	this := SecurityEventTokenRequestJwtBody{}
	return &this
}

// GetAud returns the Aud field value
func (o *SecurityEventTokenRequestJwtBody) GetAud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtBody) GetAudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aud, true
}

// SetAud sets field value
func (o *SecurityEventTokenRequestJwtBody) SetAud(v string) {
	o.Aud = v
}

// GetEvents returns the Events field value
func (o *SecurityEventTokenRequestJwtBody) GetEvents() SecurityEventTokenRequestJwtEvents {
	if o == nil {
		var ret SecurityEventTokenRequestJwtEvents
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtBody) GetEventsOk() (*SecurityEventTokenRequestJwtEvents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *SecurityEventTokenRequestJwtBody) SetEvents(v SecurityEventTokenRequestJwtEvents) {
	o.Events = v
}

// GetIat returns the Iat field value
func (o *SecurityEventTokenRequestJwtBody) GetIat() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Iat
}

// GetIatOk returns a tuple with the Iat field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtBody) GetIatOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iat, true
}

// SetIat sets field value
func (o *SecurityEventTokenRequestJwtBody) SetIat(v int64) {
	o.Iat = v
}

// GetIss returns the Iss field value
func (o *SecurityEventTokenRequestJwtBody) GetIss() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iss
}

// GetIssOk returns a tuple with the Iss field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtBody) GetIssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iss, true
}

// SetIss sets field value
func (o *SecurityEventTokenRequestJwtBody) SetIss(v string) {
	o.Iss = v
}

// GetJti returns the Jti field value
func (o *SecurityEventTokenRequestJwtBody) GetJti() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Jti
}

// GetJtiOk returns a tuple with the Jti field value
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtBody) GetJtiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jti, true
}

// SetJti sets field value
func (o *SecurityEventTokenRequestJwtBody) SetJti(v string) {
	o.Jti = v
}

func (o SecurityEventTokenRequestJwtBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["aud"] = o.Aud
	}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["iat"] = o.Iat
	}
	if true {
		toSerialize["iss"] = o.Iss
	}
	if true {
		toSerialize["jti"] = o.Jti
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventTokenRequestJwtBody) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventTokenRequestJwtBody := _SecurityEventTokenRequestJwtBody{}

	err = json.Unmarshal(bytes, &varSecurityEventTokenRequestJwtBody)
	if err == nil {
		*o = SecurityEventTokenRequestJwtBody(varSecurityEventTokenRequestJwtBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "aud")
		delete(additionalProperties, "events")
		delete(additionalProperties, "iat")
		delete(additionalProperties, "iss")
		delete(additionalProperties, "jti")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventTokenRequestJwtBody struct {
	value *SecurityEventTokenRequestJwtBody
	isSet bool
}

func (v NullableSecurityEventTokenRequestJwtBody) Get() *SecurityEventTokenRequestJwtBody {
	return v.value
}

func (v *NullableSecurityEventTokenRequestJwtBody) Set(val *SecurityEventTokenRequestJwtBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenRequestJwtBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenRequestJwtBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenRequestJwtBody(val *SecurityEventTokenRequestJwtBody) *NullableSecurityEventTokenRequestJwtBody {
	return &NullableSecurityEventTokenRequestJwtBody{value: val, isSet: true}
}

func (v NullableSecurityEventTokenRequestJwtBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenRequestJwtBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

