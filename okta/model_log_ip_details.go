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
)

// checks if the LogIpDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogIpDetails{}

// LogIpDetails Details about the associated IP address
type LogIpDetails struct {
	// The [Autonomous system](https://docs.telemetry.mozilla.org/datasets/other/asn_aggregates/reference) number that's associated with the IP address
	AsNumber NullableInt32 `json:"asNumber,omitempty"`
	// The name associated with the Autonomous System Number (ASN)
	AsOrg NullableString `json:"asOrg,omitempty"`
	// The domain name associated with the IP address
	Domain NullableString `json:"domain,omitempty"`
	// The associated IP service categories for the IP address
	IpServiceCategories []LogIpServiceCategory `json:"ipServiceCategories,omitempty"`
	// The internet service provider associated with the IP address
	Isp                  NullableString `json:"isp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogIpDetails LogIpDetails

// NewLogIpDetails instantiates a new LogIpDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIpDetails() *LogIpDetails {
	this := LogIpDetails{}
	return &this
}

// NewLogIpDetailsWithDefaults instantiates a new LogIpDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogIpDetailsWithDefaults() *LogIpDetails {
	this := LogIpDetails{}
	return &this
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpDetails) GetAsNumber() int32 {
	if o == nil || IsNil(o.AsNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.AsNumber.Get()
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpDetails) GetAsNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsNumber.Get(), o.AsNumber.IsSet()
}

// HasAsNumber returns a boolean if a field has been set.
func (o *LogIpDetails) HasAsNumber() bool {
	if o != nil && o.AsNumber.IsSet() {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given NullableInt32 and assigns it to the AsNumber field.
func (o *LogIpDetails) SetAsNumber(v int32) {
	o.AsNumber.Set(&v)
}

// SetAsNumberNil sets the value for AsNumber to be an explicit nil
func (o *LogIpDetails) SetAsNumberNil() {
	o.AsNumber.Set(nil)
}

// UnsetAsNumber ensures that no value is present for AsNumber, not even an explicit nil
func (o *LogIpDetails) UnsetAsNumber() {
	o.AsNumber.Unset()
}

// GetAsOrg returns the AsOrg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpDetails) GetAsOrg() string {
	if o == nil || IsNil(o.AsOrg.Get()) {
		var ret string
		return ret
	}
	return *o.AsOrg.Get()
}

// GetAsOrgOk returns a tuple with the AsOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpDetails) GetAsOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AsOrg.Get(), o.AsOrg.IsSet()
}

// HasAsOrg returns a boolean if a field has been set.
func (o *LogIpDetails) HasAsOrg() bool {
	if o != nil && o.AsOrg.IsSet() {
		return true
	}

	return false
}

// SetAsOrg gets a reference to the given NullableString and assigns it to the AsOrg field.
func (o *LogIpDetails) SetAsOrg(v string) {
	o.AsOrg.Set(&v)
}

// SetAsOrgNil sets the value for AsOrg to be an explicit nil
func (o *LogIpDetails) SetAsOrgNil() {
	o.AsOrg.Set(nil)
}

// UnsetAsOrg ensures that no value is present for AsOrg, not even an explicit nil
func (o *LogIpDetails) UnsetAsOrg() {
	o.AsOrg.Unset()
}

// GetDomain returns the Domain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpDetails) GetDomain() string {
	if o == nil || IsNil(o.Domain.Get()) {
		var ret string
		return ret
	}
	return *o.Domain.Get()
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpDetails) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Domain.Get(), o.Domain.IsSet()
}

// HasDomain returns a boolean if a field has been set.
func (o *LogIpDetails) HasDomain() bool {
	if o != nil && o.Domain.IsSet() {
		return true
	}

	return false
}

// SetDomain gets a reference to the given NullableString and assigns it to the Domain field.
func (o *LogIpDetails) SetDomain(v string) {
	o.Domain.Set(&v)
}

// SetDomainNil sets the value for Domain to be an explicit nil
func (o *LogIpDetails) SetDomainNil() {
	o.Domain.Set(nil)
}

// UnsetDomain ensures that no value is present for Domain, not even an explicit nil
func (o *LogIpDetails) UnsetDomain() {
	o.Domain.Unset()
}

// GetIpServiceCategories returns the IpServiceCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpDetails) GetIpServiceCategories() []LogIpServiceCategory {
	if o == nil {
		var ret []LogIpServiceCategory
		return ret
	}
	return o.IpServiceCategories
}

// GetIpServiceCategoriesOk returns a tuple with the IpServiceCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpDetails) GetIpServiceCategoriesOk() ([]LogIpServiceCategory, bool) {
	if o == nil || IsNil(o.IpServiceCategories) {
		return nil, false
	}
	return o.IpServiceCategories, true
}

// HasIpServiceCategories returns a boolean if a field has been set.
func (o *LogIpDetails) HasIpServiceCategories() bool {
	if o != nil && !IsNil(o.IpServiceCategories) {
		return true
	}

	return false
}

// SetIpServiceCategories gets a reference to the given []LogIpServiceCategory and assigns it to the IpServiceCategories field.
func (o *LogIpDetails) SetIpServiceCategories(v []LogIpServiceCategory) {
	o.IpServiceCategories = v
}

// GetIsp returns the Isp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogIpDetails) GetIsp() string {
	if o == nil || IsNil(o.Isp.Get()) {
		var ret string
		return ret
	}
	return *o.Isp.Get()
}

// GetIspOk returns a tuple with the Isp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogIpDetails) GetIspOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Isp.Get(), o.Isp.IsSet()
}

// HasIsp returns a boolean if a field has been set.
func (o *LogIpDetails) HasIsp() bool {
	if o != nil && o.Isp.IsSet() {
		return true
	}

	return false
}

// SetIsp gets a reference to the given NullableString and assigns it to the Isp field.
func (o *LogIpDetails) SetIsp(v string) {
	o.Isp.Set(&v)
}

// SetIspNil sets the value for Isp to be an explicit nil
func (o *LogIpDetails) SetIspNil() {
	o.Isp.Set(nil)
}

// UnsetIsp ensures that no value is present for Isp, not even an explicit nil
func (o *LogIpDetails) UnsetIsp() {
	o.Isp.Unset()
}

func (o LogIpDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogIpDetails) ToMap() (map[string]interface{}, error) {
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
	if o.IpServiceCategories != nil {
		toSerialize["ipServiceCategories"] = o.IpServiceCategories
	}
	if o.Isp.IsSet() {
		toSerialize["isp"] = o.Isp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogIpDetails) UnmarshalJSON(data []byte) (err error) {
	varLogIpDetails := _LogIpDetails{}

	err = json.Unmarshal(data, &varLogIpDetails)

	if err != nil {
		return err
	}

	*o = LogIpDetails(varLogIpDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asNumber")
		delete(additionalProperties, "asOrg")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "ipServiceCategories")
		delete(additionalProperties, "isp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogIpDetails struct {
	value *LogIpDetails
	isSet bool
}

func (v NullableLogIpDetails) Get() *LogIpDetails {
	return v.value
}

func (v *NullableLogIpDetails) Set(val *LogIpDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableLogIpDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableLogIpDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogIpDetails(val *LogIpDetails) *NullableLogIpDetails {
	return &NullableLogIpDetails{value: val, isSet: true}
}

func (v NullableLogIpDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogIpDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
