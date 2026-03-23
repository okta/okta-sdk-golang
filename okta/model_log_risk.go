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

// checks if the LogRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogRisk{}

// LogRisk Risk associated with the event
type LogRisk struct {
	// The name of the detection mechanism that identified the risk
	DetectionName NullableString `json:"detectionName,omitempty"`
	// The entity that issued the associated risk
	Issuer NullableString `json:"issuer,omitempty"`
	// The risk level associated with the request
	Level NullableString `json:"level,omitempty"`
	// The previous risk level (if any) associated with the user
	PreviousLevel NullableString `json:"previousLevel,omitempty"`
	// Reasons for the associated risk
	Reasons              []string `json:"reasons,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogRisk LogRisk

// NewLogRisk instantiates a new LogRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogRisk() *LogRisk {
	this := LogRisk{}
	return &this
}

// NewLogRiskWithDefaults instantiates a new LogRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogRiskWithDefaults() *LogRisk {
	this := LogRisk{}
	return &this
}

// GetDetectionName returns the DetectionName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogRisk) GetDetectionName() string {
	if o == nil || IsNil(o.DetectionName.Get()) {
		var ret string
		return ret
	}
	return *o.DetectionName.Get()
}

// GetDetectionNameOk returns a tuple with the DetectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogRisk) GetDetectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetectionName.Get(), o.DetectionName.IsSet()
}

// HasDetectionName returns a boolean if a field has been set.
func (o *LogRisk) HasDetectionName() bool {
	if o != nil && o.DetectionName.IsSet() {
		return true
	}

	return false
}

// SetDetectionName gets a reference to the given NullableString and assigns it to the DetectionName field.
func (o *LogRisk) SetDetectionName(v string) {
	o.DetectionName.Set(&v)
}

// SetDetectionNameNil sets the value for DetectionName to be an explicit nil
func (o *LogRisk) SetDetectionNameNil() {
	o.DetectionName.Set(nil)
}

// UnsetDetectionName ensures that no value is present for DetectionName, not even an explicit nil
func (o *LogRisk) UnsetDetectionName() {
	o.DetectionName.Unset()
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogRisk) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogRisk) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *LogRisk) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *LogRisk) SetIssuer(v string) {
	o.Issuer.Set(&v)
}

// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *LogRisk) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *LogRisk) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogRisk) GetLevel() string {
	if o == nil || IsNil(o.Level.Get()) {
		var ret string
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogRisk) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *LogRisk) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableString and assigns it to the Level field.
func (o *LogRisk) SetLevel(v string) {
	o.Level.Set(&v)
}

// SetLevelNil sets the value for Level to be an explicit nil
func (o *LogRisk) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *LogRisk) UnsetLevel() {
	o.Level.Unset()
}

// GetPreviousLevel returns the PreviousLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogRisk) GetPreviousLevel() string {
	if o == nil || IsNil(o.PreviousLevel.Get()) {
		var ret string
		return ret
	}
	return *o.PreviousLevel.Get()
}

// GetPreviousLevelOk returns a tuple with the PreviousLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogRisk) GetPreviousLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousLevel.Get(), o.PreviousLevel.IsSet()
}

// HasPreviousLevel returns a boolean if a field has been set.
func (o *LogRisk) HasPreviousLevel() bool {
	if o != nil && o.PreviousLevel.IsSet() {
		return true
	}

	return false
}

// SetPreviousLevel gets a reference to the given NullableString and assigns it to the PreviousLevel field.
func (o *LogRisk) SetPreviousLevel(v string) {
	o.PreviousLevel.Set(&v)
}

// SetPreviousLevelNil sets the value for PreviousLevel to be an explicit nil
func (o *LogRisk) SetPreviousLevelNil() {
	o.PreviousLevel.Set(nil)
}

// UnsetPreviousLevel ensures that no value is present for PreviousLevel, not even an explicit nil
func (o *LogRisk) UnsetPreviousLevel() {
	o.PreviousLevel.Unset()
}

// GetReasons returns the Reasons field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LogRisk) GetReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LogRisk) GetReasonsOk() ([]string, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *LogRisk) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []string and assigns it to the Reasons field.
func (o *LogRisk) SetReasons(v []string) {
	o.Reasons = v
}

func (o LogRisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DetectionName.IsSet() {
		toSerialize["detectionName"] = o.DetectionName.Get()
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.PreviousLevel.IsSet() {
		toSerialize["previousLevel"] = o.PreviousLevel.Get()
	}
	if o.Reasons != nil {
		toSerialize["reasons"] = o.Reasons
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogRisk) UnmarshalJSON(data []byte) (err error) {
	varLogRisk := _LogRisk{}

	err = json.Unmarshal(data, &varLogRisk)

	if err != nil {
		return err
	}

	*o = LogRisk(varLogRisk)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "detectionName")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "level")
		delete(additionalProperties, "previousLevel")
		delete(additionalProperties, "reasons")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogRisk struct {
	value *LogRisk
	isSet bool
}

func (v NullableLogRisk) Get() *LogRisk {
	return v.value
}

func (v *NullableLogRisk) Set(val *LogRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableLogRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableLogRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogRisk(val *LogRisk) *NullableLogRisk {
	return &NullableLogRisk{value: val, isSet: true}
}

func (v NullableLogRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
