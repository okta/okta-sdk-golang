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

// checks if the UserFactorSupported type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFactorSupported{}

// UserFactorSupported struct for UserFactorSupported
type UserFactorSupported struct {
	// Indicates if the factor is required for the specified user
	Enrollment *string `json:"enrollment,omitempty"`
	// Type of factor
	FactorType *string `json:"factorType,omitempty"`
	Provider   *string `json:"provider,omitempty"`
	// Status of the factor
	Status *string `json:"status,omitempty"`
	// Name of the factor vendor. This is usually the same as the provider except for On-Prem MFA, which depends on admin settings.
	VendorName *string `json:"vendorName,omitempty"`
	// Embedded resources related to the factor
	Embedded             map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links                *UserFactorLinks                  `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorSupported UserFactorSupported

// NewUserFactorSupported instantiates a new UserFactorSupported object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorSupported() *UserFactorSupported {
	this := UserFactorSupported{}
	return &this
}

// NewUserFactorSupportedWithDefaults instantiates a new UserFactorSupported object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorSupportedWithDefaults() *UserFactorSupported {
	this := UserFactorSupported{}
	return &this
}

// GetEnrollment returns the Enrollment field value if set, zero value otherwise.
func (o *UserFactorSupported) GetEnrollment() string {
	if o == nil || IsNil(o.Enrollment) {
		var ret string
		return ret
	}
	return *o.Enrollment
}

// GetEnrollmentOk returns a tuple with the Enrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetEnrollmentOk() (*string, bool) {
	if o == nil || IsNil(o.Enrollment) {
		return nil, false
	}
	return o.Enrollment, true
}

// HasEnrollment returns a boolean if a field has been set.
func (o *UserFactorSupported) HasEnrollment() bool {
	if o != nil && !IsNil(o.Enrollment) {
		return true
	}

	return false
}

// SetEnrollment gets a reference to the given string and assigns it to the Enrollment field.
func (o *UserFactorSupported) SetEnrollment(v string) {
	o.Enrollment = &v
}

// GetFactorType returns the FactorType field value if set, zero value otherwise.
func (o *UserFactorSupported) GetFactorType() string {
	if o == nil || IsNil(o.FactorType) {
		var ret string
		return ret
	}
	return *o.FactorType
}

// GetFactorTypeOk returns a tuple with the FactorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetFactorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FactorType) {
		return nil, false
	}
	return o.FactorType, true
}

// HasFactorType returns a boolean if a field has been set.
func (o *UserFactorSupported) HasFactorType() bool {
	if o != nil && !IsNil(o.FactorType) {
		return true
	}

	return false
}

// SetFactorType gets a reference to the given string and assigns it to the FactorType field.
func (o *UserFactorSupported) SetFactorType(v string) {
	o.FactorType = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *UserFactorSupported) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *UserFactorSupported) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *UserFactorSupported) SetProvider(v string) {
	o.Provider = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserFactorSupported) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserFactorSupported) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserFactorSupported) SetStatus(v string) {
	o.Status = &v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *UserFactorSupported) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *UserFactorSupported) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *UserFactorSupported) SetVendorName(v string) {
	o.VendorName = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactorSupported) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactorSupported) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *UserFactorSupported) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorSupported) GetLinks() UserFactorLinks {
	if o == nil || IsNil(o.Links) {
		var ret UserFactorLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorSupported) GetLinksOk() (*UserFactorLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorSupported) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorLinks and assigns it to the Links field.
func (o *UserFactorSupported) SetLinks(v UserFactorLinks) {
	o.Links = &v
}

func (o UserFactorSupported) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFactorSupported) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enrollment) {
		toSerialize["enrollment"] = o.Enrollment
	}
	if !IsNil(o.FactorType) {
		toSerialize["factorType"] = o.FactorType
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserFactorSupported) UnmarshalJSON(data []byte) (err error) {
	varUserFactorSupported := _UserFactorSupported{}

	err = json.Unmarshal(data, &varUserFactorSupported)

	if err != nil {
		return err
	}

	*o = UserFactorSupported(varUserFactorSupported)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enrollment")
		delete(additionalProperties, "factorType")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "status")
		delete(additionalProperties, "vendorName")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserFactorSupported struct {
	value *UserFactorSupported
	isSet bool
}

func (v NullableUserFactorSupported) Get() *UserFactorSupported {
	return v.value
}

func (v *NullableUserFactorSupported) Set(val *UserFactorSupported) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorSupported) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorSupported) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorSupported(val *UserFactorSupported) *NullableUserFactorSupported {
	return &NullableUserFactorSupported{value: val, isSet: true}
}

func (v NullableUserFactorSupported) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorSupported) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
