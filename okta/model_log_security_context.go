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

// LogSecurityContext struct for LogSecurityContext
type LogSecurityContext struct {
	AsNumber *int32 `json:"asNumber,omitempty"`
	AsOrg *string `json:"asOrg,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Isp *string `json:"isp,omitempty"`
	IsProxy *bool `json:"isProxy,omitempty"`
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
	if o == nil || o.AsNumber == nil {
		var ret int32
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetAsNumberOk() (*int32, bool) {
	if o == nil || o.AsNumber == nil {
		return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *LogSecurityContext) HasAsNumber() bool {
	if o != nil && o.AsNumber != nil {
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
	if o == nil || o.AsOrg == nil {
		var ret string
		return ret
	}
	return *o.AsOrg
}

// GetAsOrgOk returns a tuple with the AsOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetAsOrgOk() (*string, bool) {
	if o == nil || o.AsOrg == nil {
		return nil, false
	}
	return o.AsOrg, true
}

// HasAsOrg returns a boolean if a field has been set.
func (o *LogSecurityContext) HasAsOrg() bool {
	if o != nil && o.AsOrg != nil {
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
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *LogSecurityContext) HasDomain() bool {
	if o != nil && o.Domain != nil {
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
	if o == nil || o.Isp == nil {
		var ret string
		return ret
	}
	return *o.Isp
}

// GetIspOk returns a tuple with the Isp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetIspOk() (*string, bool) {
	if o == nil || o.Isp == nil {
		return nil, false
	}
	return o.Isp, true
}

// HasIsp returns a boolean if a field has been set.
func (o *LogSecurityContext) HasIsp() bool {
	if o != nil && o.Isp != nil {
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
	if o == nil || o.IsProxy == nil {
		var ret bool
		return ret
	}
	return *o.IsProxy
}

// GetIsProxyOk returns a tuple with the IsProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogSecurityContext) GetIsProxyOk() (*bool, bool) {
	if o == nil || o.IsProxy == nil {
		return nil, false
	}
	return o.IsProxy, true
}

// HasIsProxy returns a boolean if a field has been set.
func (o *LogSecurityContext) HasIsProxy() bool {
	if o != nil && o.IsProxy != nil {
		return true
	}

	return false
}

// SetIsProxy gets a reference to the given bool and assigns it to the IsProxy field.
func (o *LogSecurityContext) SetIsProxy(v bool) {
	o.IsProxy = &v
}

func (o LogSecurityContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AsNumber != nil {
		toSerialize["asNumber"] = o.AsNumber
	}
	if o.AsOrg != nil {
		toSerialize["asOrg"] = o.AsOrg
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Isp != nil {
		toSerialize["isp"] = o.Isp
	}
	if o.IsProxy != nil {
		toSerialize["isProxy"] = o.IsProxy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogSecurityContext) UnmarshalJSON(bytes []byte) (err error) {
	varLogSecurityContext := _LogSecurityContext{}

	err = json.Unmarshal(bytes, &varLogSecurityContext)
	if err == nil {
		*o = LogSecurityContext(varLogSecurityContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "asNumber")
		delete(additionalProperties, "asOrg")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "isp")
		delete(additionalProperties, "isProxy")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

