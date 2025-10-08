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
	"time"
)

// checks if the OktaSupportCaseImpersonation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaSupportCaseImpersonation{}

// OktaSupportCaseImpersonation Allows the Okta Support team to sign in to your org as an admin and troubleshoot issues
type OktaSupportCaseImpersonation struct {
	// Status of Okta Support access
	Status *string `json:"status,omitempty"`
	// Expiration date of Okta Support access
	Expiration           NullableTime `json:"expiration,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OktaSupportCaseImpersonation OktaSupportCaseImpersonation

// NewOktaSupportCaseImpersonation instantiates a new OktaSupportCaseImpersonation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaSupportCaseImpersonation() *OktaSupportCaseImpersonation {
	this := OktaSupportCaseImpersonation{}
	return &this
}

// NewOktaSupportCaseImpersonationWithDefaults instantiates a new OktaSupportCaseImpersonation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaSupportCaseImpersonationWithDefaults() *OktaSupportCaseImpersonation {
	this := OktaSupportCaseImpersonation{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OktaSupportCaseImpersonation) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaSupportCaseImpersonation) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OktaSupportCaseImpersonation) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OktaSupportCaseImpersonation) SetStatus(v string) {
	o.Status = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OktaSupportCaseImpersonation) GetExpiration() time.Time {
	if o == nil || IsNil(o.Expiration.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Expiration.Get()
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OktaSupportCaseImpersonation) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Expiration.Get(), o.Expiration.IsSet()
}

// HasExpiration returns a boolean if a field has been set.
func (o *OktaSupportCaseImpersonation) HasExpiration() bool {
	if o != nil && o.Expiration.IsSet() {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given NullableTime and assigns it to the Expiration field.
func (o *OktaSupportCaseImpersonation) SetExpiration(v time.Time) {
	o.Expiration.Set(&v)
}

// SetExpirationNil sets the value for Expiration to be an explicit nil
func (o *OktaSupportCaseImpersonation) SetExpirationNil() {
	o.Expiration.Set(nil)
}

// UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
func (o *OktaSupportCaseImpersonation) UnsetExpiration() {
	o.Expiration.Unset()
}

func (o OktaSupportCaseImpersonation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaSupportCaseImpersonation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Expiration.IsSet() {
		toSerialize["expiration"] = o.Expiration.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaSupportCaseImpersonation) UnmarshalJSON(data []byte) (err error) {
	varOktaSupportCaseImpersonation := _OktaSupportCaseImpersonation{}

	err = json.Unmarshal(data, &varOktaSupportCaseImpersonation)

	if err != nil {
		return err
	}

	*o = OktaSupportCaseImpersonation(varOktaSupportCaseImpersonation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "expiration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaSupportCaseImpersonation struct {
	value *OktaSupportCaseImpersonation
	isSet bool
}

func (v NullableOktaSupportCaseImpersonation) Get() *OktaSupportCaseImpersonation {
	return v.value
}

func (v *NullableOktaSupportCaseImpersonation) Set(val *OktaSupportCaseImpersonation) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaSupportCaseImpersonation) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaSupportCaseImpersonation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaSupportCaseImpersonation(val *OktaSupportCaseImpersonation) *NullableOktaSupportCaseImpersonation {
	return &NullableOktaSupportCaseImpersonation{value: val, isSet: true}
}

func (v NullableOktaSupportCaseImpersonation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaSupportCaseImpersonation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
