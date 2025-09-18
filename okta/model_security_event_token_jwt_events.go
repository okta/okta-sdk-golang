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
)

// checks if the SecurityEventTokenJwtEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityEventTokenJwtEvents{}

// SecurityEventTokenJwtEvents A non-empty set of events. Expected size is 1 for each SET
type SecurityEventTokenJwtEvents struct {
	HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange *CaepCredentialChangeEvent `json:"https://schemas.openid.net/secevent/caep/event-type/credential-change,omitempty"`
	HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked   *CaepSessionRevokedEvent   `json:"https://schemas.openid.net/secevent/caep/event-type/session-revoked,omitempty"`
	AdditionalProperties                                       map[string]interface{}
}

type _SecurityEventTokenJwtEvents SecurityEventTokenJwtEvents

// NewSecurityEventTokenJwtEvents instantiates a new SecurityEventTokenJwtEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenJwtEvents() *SecurityEventTokenJwtEvents {
	this := SecurityEventTokenJwtEvents{}
	return &this
}

// NewSecurityEventTokenJwtEventsWithDefaults instantiates a new SecurityEventTokenJwtEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenJwtEventsWithDefaults() *SecurityEventTokenJwtEvents {
	this := SecurityEventTokenJwtEvents{}
	return &this
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange returns the HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange field value if set, zero value otherwise.
func (o *SecurityEventTokenJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange() CaepCredentialChangeEvent {
	if o == nil || IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange) {
		var ret CaepCredentialChangeEvent
		return ret
	}
	return *o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChangeOk returns a tuple with the HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChangeOk() (*CaepCredentialChangeEvent, bool) {
	if o == nil || IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange) {
		return nil, false
	}
	return o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange, true
}

// HasHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange returns a boolean if a field has been set.
func (o *SecurityEventTokenJwtEvents) HasHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange() bool {
	if o != nil && !IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange) {
		return true
	}

	return false
}

// SetHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange gets a reference to the given CaepCredentialChangeEvent and assigns it to the HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange field.
func (o *SecurityEventTokenJwtEvents) SetHttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange(v CaepCredentialChangeEvent) {
	o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange = &v
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked returns the HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked field value if set, zero value otherwise.
func (o *SecurityEventTokenJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked() CaepSessionRevokedEvent {
	if o == nil || IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked) {
		var ret CaepSessionRevokedEvent
		return ret
	}
	return *o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevokedOk returns a tuple with the HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevokedOk() (*CaepSessionRevokedEvent, bool) {
	if o == nil || IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked) {
		return nil, false
	}
	return o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked, true
}

// HasHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked returns a boolean if a field has been set.
func (o *SecurityEventTokenJwtEvents) HasHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked() bool {
	if o != nil && !IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked) {
		return true
	}

	return false
}

// SetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked gets a reference to the given CaepSessionRevokedEvent and assigns it to the HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked field.
func (o *SecurityEventTokenJwtEvents) SetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked(v CaepSessionRevokedEvent) {
	o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked = &v
}

func (o SecurityEventTokenJwtEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityEventTokenJwtEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange) {
		toSerialize["https://schemas.openid.net/secevent/caep/event-type/credential-change"] = o.HttpsSchemasOpenidNetSeceventCaepEventTypeCredentialChange
	}
	if !IsNil(o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked) {
		toSerialize["https://schemas.openid.net/secevent/caep/event-type/session-revoked"] = o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityEventTokenJwtEvents) UnmarshalJSON(data []byte) (err error) {
	varSecurityEventTokenJwtEvents := _SecurityEventTokenJwtEvents{}

	err = json.Unmarshal(data, &varSecurityEventTokenJwtEvents)

	if err != nil {
		return err
	}

	*o = SecurityEventTokenJwtEvents(varSecurityEventTokenJwtEvents)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "https://schemas.openid.net/secevent/caep/event-type/credential-change")
		delete(additionalProperties, "https://schemas.openid.net/secevent/caep/event-type/session-revoked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityEventTokenJwtEvents struct {
	value *SecurityEventTokenJwtEvents
	isSet bool
}

func (v NullableSecurityEventTokenJwtEvents) Get() *SecurityEventTokenJwtEvents {
	return v.value
}

func (v *NullableSecurityEventTokenJwtEvents) Set(val *SecurityEventTokenJwtEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenJwtEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenJwtEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenJwtEvents(val *SecurityEventTokenJwtEvents) *NullableSecurityEventTokenJwtEvents {
	return &NullableSecurityEventTokenJwtEvents{value: val, isSet: true}
}

func (v NullableSecurityEventTokenJwtEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenJwtEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
