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

// checks if the LogSecurityContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSecurityContext{}

// LogSecurityContext The `securityContext` object provides security information that is directly related to the evaluation of the event's IP reputation. IP reputation is a trustworthiness rating that evaluates how likely a sender is to be malicious and is based on the sender's IP address. As the name implies, the `securityContext` object is useful for security applications-flagging and inspecting suspicious events.
type LogSecurityContext struct {
	// The [Autonomous system](https://docs.telemetry.mozilla.org/datasets/other/asn_aggregates/reference) number that's associated with the autonomous system the event request was sourced to
	AsNumber NullableInt32 `json:"asNumber,omitempty"`
	// The organization that is associated with the autonomous system that the event request is sourced to
	AsOrg NullableString `json:"asOrg,omitempty"`
	// The domain name that's associated with the IP address of the inbound event request
	Domain NullableString `json:"domain,omitempty"`
	// The Internet service provider that's used to send the event's request
	Isp NullableString `json:"isp,omitempty"`
	// Specifies whether an event's request is from a known proxy
	IsProxy NullableBool `json:"isProxy,omitempty"`
	// The result of the user behavior detection models associated with the event
	UserBehaviors        []string `json:"userBehaviors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogSecurityContext LogSecurityContext

// NewLogSecurityContext instantiates a new LogSecurityContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogSecurityContext() *LogSecurityContext {
	this := LogSecurityContext{}
	return &this
}

// NewLogSecurityContextWithDefaults instantiates a new LogSecurityContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogSecurityContextWithDefaults() *LogSecurityContext {
	this := LogSecurityContext{}
	return &this
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogSecurityContext) GetAsNumber() int32 {
	if o == nil || IsNil(o.AsNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.AsNumber.Get()
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogSecurityContext) GetAsNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsNumber.Get(), o.AsNumber.IsSet()
}

// HasAsNumber returns a boolean if a field has been set.
func (o *LogSecurityContext) HasAsNumber() bool {
	if o != nil && o.AsNumber.IsSet() {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given NullableInt32 and assigns it to the AsNumber field.
func (o *LogSecurityContext) SetAsNumber(v int32) {
	o.AsNumber.Set(&v)
}

// SetAsNumberNil sets the value for AsNumber to be an explicit nil
func (o *LogSecurityContext) SetAsNumberNil() {
	o.AsNumber.Set(nil)
}

// UnsetAsNumber ensures that no value is present for AsNumber, not even an explicit nil
func (o *LogSecurityContext) UnsetAsNumber() {
	o.AsNumber.Unset()
}

// GetAsOrg returns the AsOrg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogSecurityContext) GetAsOrg() string {
	if o == nil || IsNil(o.AsOrg.Get()) {
		var ret string
		return ret
	}
	return *o.AsOrg.Get()
}

// GetAsOrgOk returns a tuple with the AsOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogSecurityContext) GetAsOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsOrg.Get(), o.AsOrg.IsSet()
}

// HasAsOrg returns a boolean if a field has been set.
func (o *LogSecurityContext) HasAsOrg() bool {
	if o != nil && o.AsOrg.IsSet() {
		return true
	}

	return false
}

// SetAsOrg gets a reference to the given NullableString and assigns it to the AsOrg field.
func (o *LogSecurityContext) SetAsOrg(v string) {
	o.AsOrg.Set(&v)
}

// SetAsOrgNil sets the value for AsOrg to be an explicit nil
func (o *LogSecurityContext) SetAsOrgNil() {
	o.AsOrg.Set(nil)
}

// UnsetAsOrg ensures that no value is present for AsOrg, not even an explicit nil
func (o *LogSecurityContext) UnsetAsOrg() {
	o.AsOrg.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogSecurityContext) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogSecurityContext) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *LogSecurityContext) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *LogSecurityContext) SetDomain(v string) {
	o.Domain.Set(&v)
}

// SetDomainNil sets the value for Domain to be an explicit nil
func (o *LogSecurityContext) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *LogSecurityContext) UnsetDomain() {
	o.Domain.Unset()
}

// GetIsp returns the Isp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogSecurityContext) GetIsp() string {
	if o == nil || IsNil(o.Isp.Get()) {
		var ret string
		return ret
	}
	return *o.Isp.Get()
}

// GetIspOk returns a tuple with the Isp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogSecurityContext) GetIspOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Isp.Get(), o.Isp.IsSet()
}

// HasIsp returns a boolean if a field has been set.
func (o *LogSecurityContext) HasIsp() bool {
	if o != nil && o.Isp.IsSet() {
		return true
	}

	return false
}

// SetIsp gets a reference to the given NullableString and assigns it to the Isp field.
func (o *LogSecurityContext) SetIsp(v string) {
	o.Isp.Set(&v)
}

// SetIspNil sets the value for Isp to be an explicit nil
func (o *LogSecurityContext) SetIspNil() {
	o.Isp.Set(nil)
}

// UnsetIsp ensures that no value is present for Isp, not even an explicit nil
func (o *LogSecurityContext) UnsetIsp() {
	o.Isp.Unset()
}

// GetIsProxy returns the IsProxy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogSecurityContext) GetIsProxy() bool {
	if o == nil || IsNil(o.IsProxy.Get()) {
		var ret bool
		return ret
	}
	return *o.IsProxy.Get()
}

// GetIsProxyOk returns a tuple with the IsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogSecurityContext) GetIsProxyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsProxy.Get(), o.IsProxy.IsSet()
}

// HasIsProxy returns a boolean if a field has been set.
func (o *LogSecurityContext) HasIsProxy() bool {
	if o != nil && o.IsProxy.IsSet() {
		return true
	}

	return false
}

// SetIsProxy gets a reference to the given NullableBool and assigns it to the IsProxy field.
func (o *LogSecurityContext) SetIsProxy(v bool) {
	o.IsProxy.Set(&v)
}

// SetIsProxyNil sets the value for IsProxy to be an explicit nil
func (o *LogSecurityContext) SetIsProxyNil() {
	o.IsProxy.Set(nil)
}

// UnsetIsProxy ensures that no value is present for IsProxy, not even an explicit nil
func (o *LogSecurityContext) UnsetIsProxy() {
	o.IsProxy.Unset()
}

// GetUserBehaviors returns the UserBehaviors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogSecurityContext) GetUserBehaviors() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.UserBehaviors
}

// GetUserBehaviorsOk returns a tuple with the UserBehaviors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogSecurityContext) GetUserBehaviorsOk() ([]string, bool) {
	if o == nil || IsNil(o.UserBehaviors) {
		return nil, false
	}
	return o.UserBehaviors, true
}

// HasUserBehaviors returns a boolean if a field has been set.
func (o *LogSecurityContext) HasUserBehaviors() bool {
	if o != nil && !IsNil(o.UserBehaviors) {
		return true
	}

	return false
}

// SetUserBehaviors gets a reference to the given []string and assigns it to the UserBehaviors field.
func (o *LogSecurityContext) SetUserBehaviors(v []string) {
	o.UserBehaviors = v
}

func (o LogSecurityContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogSecurityContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AsNumber.IsSet() {
		toSerialize["asNumber"] = o.AsNumber.Get()
	}
	if o.AsOrg.IsSet() {
		toSerialize["asOrg"] = o.AsOrg.Get()
	}
	if o.Domain.IsSet() {
		toSerialize["domain"] = o.Domain.Get()
	}
	if o.Isp.IsSet() {
		toSerialize["isp"] = o.Isp.Get()
	}
	if o.IsProxy.IsSet() {
		toSerialize["isProxy"] = o.IsProxy.Get()
	}
	if o.UserBehaviors != nil {
		toSerialize["userBehaviors"] = o.UserBehaviors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogSecurityContext) UnmarshalJSON(data []byte) (err error) {
	varLogSecurityContext := _LogSecurityContext{}

	err = json.Unmarshal(data, &varLogSecurityContext)

	if err != nil {
		return err
	}

	*o = LogSecurityContext(varLogSecurityContext)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asNumber")
		delete(additionalProperties, "asOrg")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "isp")
		delete(additionalProperties, "isProxy")
		delete(additionalProperties, "userBehaviors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogSecurityContext struct {
	value *LogSecurityContext
	isSet bool
}

func (v NullableLogSecurityContext) Get() *LogSecurityContext {
	return v.value
}

func (v *NullableLogSecurityContext) Set(val *LogSecurityContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLogSecurityContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLogSecurityContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogSecurityContext(val *LogSecurityContext) *NullableLogSecurityContext {
	return &NullableLogSecurityContext{value: val, isSet: true}
}

func (v NullableLogSecurityContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogSecurityContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
