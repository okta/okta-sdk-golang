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

// SecurityEventTokenRequestJwtEvents A non-empty collection of events
type SecurityEventTokenRequestJwtEvents struct {
	HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange *OktaDeviceRiskChangeEvent `json:"https://schemas.okta.com/secevent/okta/event-type/device-risk-change,omitempty"`
	HttpsSchemasOktaComSeceventOktaEventTypeIpChange *OktaIpChangeEvent `json:"https://schemas.okta.com/secevent/okta/event-type/ip-change,omitempty"`
	HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange *OktaUserRiskChangeEvent `json:"https://schemas.okta.com/secevent/okta/event-type/user-risk-change,omitempty"`
	HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange *CaepDeviceComplianceChangeEvent `json:"https://schemas.openid.net/secevent/caep/event-type/device-compliance-change,omitempty"`
	HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked *CaepSessionRevokedEvent `json:"https://schemas.openid.net/secevent/caep/event-type/session-revoked,omitempty"`
	HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged *RiscIdentifierChangedEvent `json:"https://schemas.openid.net/secevent/risc/event-type/identifier-changed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventTokenRequestJwtEvents SecurityEventTokenRequestJwtEvents

// NewSecurityEventTokenRequestJwtEvents instantiates a new SecurityEventTokenRequestJwtEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventTokenRequestJwtEvents() *SecurityEventTokenRequestJwtEvents {
	this := SecurityEventTokenRequestJwtEvents{}
	return &this
}

// NewSecurityEventTokenRequestJwtEventsWithDefaults instantiates a new SecurityEventTokenRequestJwtEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventTokenRequestJwtEventsWithDefaults() *SecurityEventTokenRequestJwtEvents {
	this := SecurityEventTokenRequestJwtEvents{}
	return &this
}

// GetHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange returns the HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange field value if set, zero value otherwise.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange() OktaDeviceRiskChangeEvent {
	if o == nil || o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange == nil {
		var ret OktaDeviceRiskChangeEvent
		return ret
	}
	return *o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange
}

// GetHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChangeOk returns a tuple with the HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChangeOk() (*OktaDeviceRiskChangeEvent, bool) {
	if o == nil || o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange == nil {
		return nil, false
	}
	return o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange, true
}

// HasHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange returns a boolean if a field has been set.
func (o *SecurityEventTokenRequestJwtEvents) HasHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange() bool {
	if o != nil && o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange != nil {
		return true
	}

	return false
}

// SetHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange gets a reference to the given OktaDeviceRiskChangeEvent and assigns it to the HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange field.
func (o *SecurityEventTokenRequestJwtEvents) SetHttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange(v OktaDeviceRiskChangeEvent) {
	o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange = &v
}

// GetHttpsSchemasOktaComSeceventOktaEventTypeIpChange returns the HttpsSchemasOktaComSeceventOktaEventTypeIpChange field value if set, zero value otherwise.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOktaComSeceventOktaEventTypeIpChange() OktaIpChangeEvent {
	if o == nil || o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange == nil {
		var ret OktaIpChangeEvent
		return ret
	}
	return *o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange
}

// GetHttpsSchemasOktaComSeceventOktaEventTypeIpChangeOk returns a tuple with the HttpsSchemasOktaComSeceventOktaEventTypeIpChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOktaComSeceventOktaEventTypeIpChangeOk() (*OktaIpChangeEvent, bool) {
	if o == nil || o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange == nil {
		return nil, false
	}
	return o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange, true
}

// HasHttpsSchemasOktaComSeceventOktaEventTypeIpChange returns a boolean if a field has been set.
func (o *SecurityEventTokenRequestJwtEvents) HasHttpsSchemasOktaComSeceventOktaEventTypeIpChange() bool {
	if o != nil && o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange != nil {
		return true
	}

	return false
}

// SetHttpsSchemasOktaComSeceventOktaEventTypeIpChange gets a reference to the given OktaIpChangeEvent and assigns it to the HttpsSchemasOktaComSeceventOktaEventTypeIpChange field.
func (o *SecurityEventTokenRequestJwtEvents) SetHttpsSchemasOktaComSeceventOktaEventTypeIpChange(v OktaIpChangeEvent) {
	o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange = &v
}

// GetHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange returns the HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange field value if set, zero value otherwise.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange() OktaUserRiskChangeEvent {
	if o == nil || o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange == nil {
		var ret OktaUserRiskChangeEvent
		return ret
	}
	return *o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange
}

// GetHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChangeOk returns a tuple with the HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChangeOk() (*OktaUserRiskChangeEvent, bool) {
	if o == nil || o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange == nil {
		return nil, false
	}
	return o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange, true
}

// HasHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange returns a boolean if a field has been set.
func (o *SecurityEventTokenRequestJwtEvents) HasHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange() bool {
	if o != nil && o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange != nil {
		return true
	}

	return false
}

// SetHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange gets a reference to the given OktaUserRiskChangeEvent and assigns it to the HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange field.
func (o *SecurityEventTokenRequestJwtEvents) SetHttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange(v OktaUserRiskChangeEvent) {
	o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange = &v
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange returns the HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange field value if set, zero value otherwise.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange() CaepDeviceComplianceChangeEvent {
	if o == nil || o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange == nil {
		var ret CaepDeviceComplianceChangeEvent
		return ret
	}
	return *o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChangeOk returns a tuple with the HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChangeOk() (*CaepDeviceComplianceChangeEvent, bool) {
	if o == nil || o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange == nil {
		return nil, false
	}
	return o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange, true
}

// HasHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange returns a boolean if a field has been set.
func (o *SecurityEventTokenRequestJwtEvents) HasHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange() bool {
	if o != nil && o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange != nil {
		return true
	}

	return false
}

// SetHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange gets a reference to the given CaepDeviceComplianceChangeEvent and assigns it to the HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange field.
func (o *SecurityEventTokenRequestJwtEvents) SetHttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange(v CaepDeviceComplianceChangeEvent) {
	o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange = &v
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked returns the HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked field value if set, zero value otherwise.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked() CaepSessionRevokedEvent {
	if o == nil || o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked == nil {
		var ret CaepSessionRevokedEvent
		return ret
	}
	return *o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked
}

// GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevokedOk returns a tuple with the HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevokedOk() (*CaepSessionRevokedEvent, bool) {
	if o == nil || o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked == nil {
		return nil, false
	}
	return o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked, true
}

// HasHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked returns a boolean if a field has been set.
func (o *SecurityEventTokenRequestJwtEvents) HasHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked() bool {
	if o != nil && o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked != nil {
		return true
	}

	return false
}

// SetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked gets a reference to the given CaepSessionRevokedEvent and assigns it to the HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked field.
func (o *SecurityEventTokenRequestJwtEvents) SetHttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked(v CaepSessionRevokedEvent) {
	o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked = &v
}

// GetHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged returns the HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged field value if set, zero value otherwise.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged() RiscIdentifierChangedEvent {
	if o == nil || o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged == nil {
		var ret RiscIdentifierChangedEvent
		return ret
	}
	return *o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged
}

// GetHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChangedOk returns a tuple with the HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventTokenRequestJwtEvents) GetHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChangedOk() (*RiscIdentifierChangedEvent, bool) {
	if o == nil || o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged == nil {
		return nil, false
	}
	return o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged, true
}

// HasHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged returns a boolean if a field has been set.
func (o *SecurityEventTokenRequestJwtEvents) HasHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged() bool {
	if o != nil && o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged != nil {
		return true
	}

	return false
}

// SetHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged gets a reference to the given RiscIdentifierChangedEvent and assigns it to the HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged field.
func (o *SecurityEventTokenRequestJwtEvents) SetHttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged(v RiscIdentifierChangedEvent) {
	o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged = &v
}

func (o SecurityEventTokenRequestJwtEvents) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange != nil {
		toSerialize["https://schemas.okta.com/secevent/okta/event-type/device-risk-change"] = o.HttpsSchemasOktaComSeceventOktaEventTypeDeviceRiskChange
	}
	if o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange != nil {
		toSerialize["https://schemas.okta.com/secevent/okta/event-type/ip-change"] = o.HttpsSchemasOktaComSeceventOktaEventTypeIpChange
	}
	if o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange != nil {
		toSerialize["https://schemas.okta.com/secevent/okta/event-type/user-risk-change"] = o.HttpsSchemasOktaComSeceventOktaEventTypeUserRiskChange
	}
	if o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange != nil {
		toSerialize["https://schemas.openid.net/secevent/caep/event-type/device-compliance-change"] = o.HttpsSchemasOpenidNetSeceventCaepEventTypeDeviceComplianceChange
	}
	if o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked != nil {
		toSerialize["https://schemas.openid.net/secevent/caep/event-type/session-revoked"] = o.HttpsSchemasOpenidNetSeceventCaepEventTypeSessionRevoked
	}
	if o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged != nil {
		toSerialize["https://schemas.openid.net/secevent/risc/event-type/identifier-changed"] = o.HttpsSchemasOpenidNetSeceventRiscEventTypeIdentifierChanged
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityEventTokenRequestJwtEvents) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityEventTokenRequestJwtEvents := _SecurityEventTokenRequestJwtEvents{}

	err = json.Unmarshal(bytes, &varSecurityEventTokenRequestJwtEvents)
	if err == nil {
		*o = SecurityEventTokenRequestJwtEvents(varSecurityEventTokenRequestJwtEvents)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "https://schemas.okta.com/secevent/okta/event-type/device-risk-change")
		delete(additionalProperties, "https://schemas.okta.com/secevent/okta/event-type/ip-change")
		delete(additionalProperties, "https://schemas.okta.com/secevent/okta/event-type/user-risk-change")
		delete(additionalProperties, "https://schemas.openid.net/secevent/caep/event-type/device-compliance-change")
		delete(additionalProperties, "https://schemas.openid.net/secevent/caep/event-type/session-revoked")
		delete(additionalProperties, "https://schemas.openid.net/secevent/risc/event-type/identifier-changed")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityEventTokenRequestJwtEvents struct {
	value *SecurityEventTokenRequestJwtEvents
	isSet bool
}

func (v NullableSecurityEventTokenRequestJwtEvents) Get() *SecurityEventTokenRequestJwtEvents {
	return v.value
}

func (v *NullableSecurityEventTokenRequestJwtEvents) Set(val *SecurityEventTokenRequestJwtEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventTokenRequestJwtEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventTokenRequestJwtEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventTokenRequestJwtEvents(val *SecurityEventTokenRequestJwtEvents) *NullableSecurityEventTokenRequestJwtEvents {
	return &NullableSecurityEventTokenRequestJwtEvents{value: val, isSet: true}
}

func (v NullableSecurityEventTokenRequestJwtEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventTokenRequestJwtEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

