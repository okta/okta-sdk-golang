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

// checks if the SecurityEventsProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityEventsProviderResponse{}

// SecurityEventsProviderResponse The Security Events Provider response
type SecurityEventsProviderResponse struct {
	// The unique identifier of this instance
	Id *string `json:"id,omitempty"`
	// The name of the Security Events Provider instance
	Name     *string                                 `json:"name,omitempty"`
	Settings *SecurityEventsProviderSettingsResponse `json:"settings,omitempty"`
	// Indicates whether the Security Events Provider is active or not
	Status *string `json:"status,omitempty"`
	// The application type of the Security Events Provider
	Type                 *string                `json:"type,omitempty"`
	Links                *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityEventsProviderResponse SecurityEventsProviderResponse

// NewSecurityEventsProviderResponse instantiates a new SecurityEventsProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityEventsProviderResponse() *SecurityEventsProviderResponse {
	this := SecurityEventsProviderResponse{}
	return &this
}

// NewSecurityEventsProviderResponseWithDefaults instantiates a new SecurityEventsProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityEventsProviderResponseWithDefaults() *SecurityEventsProviderResponse {
	this := SecurityEventsProviderResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SecurityEventsProviderResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SecurityEventsProviderResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SecurityEventsProviderResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityEventsProviderResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityEventsProviderResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityEventsProviderResponse) SetName(v string) {
	o.Name = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *SecurityEventsProviderResponse) GetSettings() SecurityEventsProviderSettingsResponse {
	if o == nil || IsNil(o.Settings) {
		var ret SecurityEventsProviderSettingsResponse
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderResponse) GetSettingsOk() (*SecurityEventsProviderSettingsResponse, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *SecurityEventsProviderResponse) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given SecurityEventsProviderSettingsResponse and assigns it to the Settings field.
func (o *SecurityEventsProviderResponse) SetSettings(v SecurityEventsProviderSettingsResponse) {
	o.Settings = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SecurityEventsProviderResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SecurityEventsProviderResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SecurityEventsProviderResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityEventsProviderResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityEventsProviderResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SecurityEventsProviderResponse) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SecurityEventsProviderResponse) GetLinks() LinksSelfAndLifecycle {
	if o == nil || IsNil(o.Links) {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityEventsProviderResponse) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SecurityEventsProviderResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *SecurityEventsProviderResponse) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o SecurityEventsProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityEventsProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SecurityEventsProviderResponse) UnmarshalJSON(data []byte) (err error) {
	varSecurityEventsProviderResponse := _SecurityEventsProviderResponse{}

	err = json.Unmarshal(data, &varSecurityEventsProviderResponse)

	if err != nil {
		return err
	}

	*o = SecurityEventsProviderResponse(varSecurityEventsProviderResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSecurityEventsProviderResponse struct {
	value *SecurityEventsProviderResponse
	isSet bool
}

func (v NullableSecurityEventsProviderResponse) Get() *SecurityEventsProviderResponse {
	return v.value
}

func (v *NullableSecurityEventsProviderResponse) Set(val *SecurityEventsProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityEventsProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityEventsProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityEventsProviderResponse(val *SecurityEventsProviderResponse) *NullableSecurityEventsProviderResponse {
	return &NullableSecurityEventsProviderResponse{value: val, isSet: true}
}

func (v NullableSecurityEventsProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityEventsProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
