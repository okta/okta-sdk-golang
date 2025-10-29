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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the RegistrationResponseErrorErrorCausesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationResponseErrorErrorCausesInner{}

// RegistrationResponseErrorErrorCausesInner struct for RegistrationResponseErrorErrorCausesInner
type RegistrationResponseErrorErrorCausesInner struct {
	// Human-readable summary of the error.
	ErrorSummary *string `json:"errorSummary,omitempty"`
	// A brief, enum-like string that indicates the nature of the error. For example, `UNIQUE_CONSTRAINT` for a property uniqueness violation.
	Reason *string `json:"reason,omitempty"`
	// Where in the request the error was found (`body`, `header`, `url`, or `query`).
	LocationType *string `json:"locationType,omitempty"`
	// The valid JSON path to the location of the error. For example, if there was an error in the user's `login` field, the `location` might be `data.userProfile.login`.
	Location *string `json:"location,omitempty"`
	// Indicates the source of the error. If the error was in the user's profile, for example, you might use `end-user`. If the error occurred in the external service, you might use `external-service`.
	Domain               *string `json:"domain,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationResponseErrorErrorCausesInner RegistrationResponseErrorErrorCausesInner

// NewRegistrationResponseErrorErrorCausesInner instantiates a new RegistrationResponseErrorErrorCausesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationResponseErrorErrorCausesInner() *RegistrationResponseErrorErrorCausesInner {
	this := RegistrationResponseErrorErrorCausesInner{}
	return &this
}

// NewRegistrationResponseErrorErrorCausesInnerWithDefaults instantiates a new RegistrationResponseErrorErrorCausesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationResponseErrorErrorCausesInnerWithDefaults() *RegistrationResponseErrorErrorCausesInner {
	this := RegistrationResponseErrorErrorCausesInner{}
	return &this
}

// GetErrorSummary returns the ErrorSummary field value if set, zero value otherwise.
func (o *RegistrationResponseErrorErrorCausesInner) GetErrorSummary() string {
	if o == nil || IsNil(o.ErrorSummary) {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetErrorSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorSummary) {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasErrorSummary() bool {
	if o != nil && !IsNil(o.ErrorSummary) {
		return true
	}

	return false
}

// SetErrorSummary gets a reference to the given string and assigns it to the ErrorSummary field.
func (o *RegistrationResponseErrorErrorCausesInner) SetErrorSummary(v string) {
	o.ErrorSummary = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RegistrationResponseErrorErrorCausesInner) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RegistrationResponseErrorErrorCausesInner) SetReason(v string) {
	o.Reason = &v
}

// GetLocationType returns the LocationType field value if set, zero value otherwise.
func (o *RegistrationResponseErrorErrorCausesInner) GetLocationType() string {
	if o == nil || IsNil(o.LocationType) {
		var ret string
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetLocationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LocationType) {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasLocationType() bool {
	if o != nil && !IsNil(o.LocationType) {
		return true
	}

	return false
}

// SetLocationType gets a reference to the given string and assigns it to the LocationType field.
func (o *RegistrationResponseErrorErrorCausesInner) SetLocationType(v string) {
	o.LocationType = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *RegistrationResponseErrorErrorCausesInner) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *RegistrationResponseErrorErrorCausesInner) SetLocation(v string) {
	o.Location = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *RegistrationResponseErrorErrorCausesInner) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *RegistrationResponseErrorErrorCausesInner) SetDomain(v string) {
	o.Domain = &v
}

func (o RegistrationResponseErrorErrorCausesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationResponseErrorErrorCausesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorSummary) {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.LocationType) {
		toSerialize["locationType"] = o.LocationType
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationResponseErrorErrorCausesInner) UnmarshalJSON(data []byte) (err error) {
	varRegistrationResponseErrorErrorCausesInner := _RegistrationResponseErrorErrorCausesInner{}

	err = json.Unmarshal(data, &varRegistrationResponseErrorErrorCausesInner)

	if err != nil {
		return err
	}

	*o = RegistrationResponseErrorErrorCausesInner(varRegistrationResponseErrorErrorCausesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "locationType")
		delete(additionalProperties, "location")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationResponseErrorErrorCausesInner struct {
	value *RegistrationResponseErrorErrorCausesInner
	isSet bool
}

func (v NullableRegistrationResponseErrorErrorCausesInner) Get() *RegistrationResponseErrorErrorCausesInner {
	return v.value
}

func (v *NullableRegistrationResponseErrorErrorCausesInner) Set(val *RegistrationResponseErrorErrorCausesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationResponseErrorErrorCausesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationResponseErrorErrorCausesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationResponseErrorErrorCausesInner(val *RegistrationResponseErrorErrorCausesInner) *NullableRegistrationResponseErrorErrorCausesInner {
	return &NullableRegistrationResponseErrorErrorCausesInner{value: val, isSet: true}
}

func (v NullableRegistrationResponseErrorErrorCausesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationResponseErrorErrorCausesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
