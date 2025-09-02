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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

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
	Domain *string `json:"domain,omitempty"`
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
	if o == nil || o.ErrorSummary == nil {
		var ret string
		return ret
	}
	return *o.ErrorSummary
}

// GetErrorSummaryOk returns a tuple with the ErrorSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetErrorSummaryOk() (*string, bool) {
	if o == nil || o.ErrorSummary == nil {
		return nil, false
	}
	return o.ErrorSummary, true
}

// HasErrorSummary returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasErrorSummary() bool {
	if o != nil && o.ErrorSummary != nil {
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
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasReason() bool {
	if o != nil && o.Reason != nil {
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
	if o == nil || o.LocationType == nil {
		var ret string
		return ret
	}
	return *o.LocationType
}

// GetLocationTypeOk returns a tuple with the LocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetLocationTypeOk() (*string, bool) {
	if o == nil || o.LocationType == nil {
		return nil, false
	}
	return o.LocationType, true
}

// HasLocationType returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasLocationType() bool {
	if o != nil && o.LocationType != nil {
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
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasLocation() bool {
	if o != nil && o.Location != nil {
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
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationResponseErrorErrorCausesInner) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RegistrationResponseErrorErrorCausesInner) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *RegistrationResponseErrorErrorCausesInner) SetDomain(v string) {
	o.Domain = &v
}

func (o RegistrationResponseErrorErrorCausesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorSummary != nil {
		toSerialize["errorSummary"] = o.ErrorSummary
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.LocationType != nil {
		toSerialize["locationType"] = o.LocationType
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RegistrationResponseErrorErrorCausesInner) UnmarshalJSON(bytes []byte) (err error) {
	varRegistrationResponseErrorErrorCausesInner := _RegistrationResponseErrorErrorCausesInner{}

	err = json.Unmarshal(bytes, &varRegistrationResponseErrorErrorCausesInner)
	if err == nil {
		*o = RegistrationResponseErrorErrorCausesInner(varRegistrationResponseErrorErrorCausesInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "errorSummary")
		delete(additionalProperties, "reason")
		delete(additionalProperties, "locationType")
		delete(additionalProperties, "location")
		delete(additionalProperties, "domain")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

