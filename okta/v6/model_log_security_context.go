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

// checks if the LogSecurityContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogSecurityContext{}

// LogSecurityContext The `securityContext` object provides security information that is directly related to the evaluation of the event's IP reputation. IP reputation is a trustworthiness rating that evaluates how likely a sender is to be malicious and is based on the sender's IP address. As the name implies, the `securityContext` object is useful for security applications-flagging and inspecting suspicious events.
type LogSecurityContext struct {
	// The [Autonomous system](https://docs.telemetry.mozilla.org/datasets/other/asn_aggregates/reference) number that's associated with the autonomous system the event request was sourced to
	AsNumber *int32 `json:"asNumber,omitempty"`
	// The organization that is associated with the autonomous system that the event request is sourced to
	AsOrg *string `json:"asOrg,omitempty"`
	// The domain name that's associated with the IP address of the inbound event request
	Domain *string `json:"domain,omitempty"`
	// The Internet service provider that's used to send the event's request
	Isp *string `json:"isp,omitempty"`
	// Specifies whether an event's request is from a known proxy
	IsProxy              *bool `json:"isProxy,omitempty"`
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

// GetAsNumber returns the AsNumber field value if set, zero value otherwise.
func (o *LogSecurityContext) GetAsNumber() int32 {
	if o == nil || IsNil(o.AsNumber) {
		var ret int32
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetAsNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.AsNumber) {
		return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *LogSecurityContext) HasAsNumber() bool {
	if o != nil && !IsNil(o.AsNumber) {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given int32 and assigns it to the AsNumber field.
func (o *LogSecurityContext) SetAsNumber(v int32) {
	o.AsNumber = &v
}

// GetAsOrg returns the AsOrg field value if set, zero value otherwise.
func (o *LogSecurityContext) GetAsOrg() string {
	if o == nil || IsNil(o.AsOrg) {
		var ret string
		return ret
	}
	return *o.AsOrg
}

// GetAsOrgOk returns a tuple with the AsOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetAsOrgOk() (*string, bool) {
	if o == nil || IsNil(o.AsOrg) {
		return nil, false
	}
	return o.AsOrg, true
}

// HasAsOrg returns a boolean if a field has been set.
func (o *LogSecurityContext) HasAsOrg() bool {
	if o != nil && !IsNil(o.AsOrg) {
		return true
	}

	return false
}

// SetAsOrg gets a reference to the given string and assigns it to the AsOrg field.
func (o *LogSecurityContext) SetAsOrg(v string) {
	o.AsOrg = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *LogSecurityContext) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LogSecurityContext) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *LogSecurityContext) SetDomain(v string) {
	o.Domain = &v
}

// GetIsp returns the Isp field value if set, zero value otherwise.
func (o *LogSecurityContext) GetIsp() string {
	if o == nil || IsNil(o.Isp) {
		var ret string
		return ret
	}
	return *o.Isp
}

// GetIspOk returns a tuple with the Isp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetIspOk() (*string, bool) {
	if o == nil || IsNil(o.Isp) {
		return nil, false
	}
	return o.Isp, true
}

// HasIsp returns a boolean if a field has been set.
func (o *LogSecurityContext) HasIsp() bool {
	if o != nil && !IsNil(o.Isp) {
		return true
	}

	return false
}

// SetIsp gets a reference to the given string and assigns it to the Isp field.
func (o *LogSecurityContext) SetIsp(v string) {
	o.Isp = &v
}

// GetIsProxy returns the IsProxy field value if set, zero value otherwise.
func (o *LogSecurityContext) GetIsProxy() bool {
	if o == nil || IsNil(o.IsProxy) {
		var ret bool
		return ret
	}
	return *o.IsProxy
}

// GetIsProxyOk returns a tuple with the IsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetIsProxyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsProxy) {
		return nil, false
	}
	return o.IsProxy, true
}

// HasIsProxy returns a boolean if a field has been set.
func (o *LogSecurityContext) HasIsProxy() bool {
	if o != nil && !IsNil(o.IsProxy) {
		return true
	}

	return false
}

// SetIsProxy gets a reference to the given bool and assigns it to the IsProxy field.
func (o *LogSecurityContext) SetIsProxy(v bool) {
	o.IsProxy = &v
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
	if !IsNil(o.AsNumber) {
		toSerialize["asNumber"] = o.AsNumber
	}
	if !IsNil(o.AsOrg) {
		toSerialize["asOrg"] = o.AsOrg
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Isp) {
		toSerialize["isp"] = o.Isp
	}
	if !IsNil(o.IsProxy) {
		toSerialize["isProxy"] = o.IsProxy
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
