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

// checks if the OktaActiveDirectoryGroupProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaActiveDirectoryGroupProfile{}

// OktaActiveDirectoryGroupProfile Profile for a group that is imported from Active Directory.  The `objectClass` for such groups is `okta:windows_security_principal`.
type OktaActiveDirectoryGroupProfile struct {
	// Description of the Windows group
	Description *string `json:"description,omitempty"`
	// The distinguished name of the Windows group
	Dn *string `json:"dn,omitempty"`
	// Base-64 encoded GUID (`objectGUID`) of the Windows group
	ExternalId *string `json:"externalId,omitempty"`
	// Name of the Windows group
	Name *string `json:"name,omitempty"`
	// Pre-Windows 2000 name of the Windows group
	SamAccountName *string `json:"samAccountName,omitempty"`
	// Fully qualified name of the Windows group
	WindowsDomainQualifiedName *string `json:"windowsDomainQualifiedName,omitempty"`
	AdditionalProperties       map[string]interface{}
}

type _OktaActiveDirectoryGroupProfile OktaActiveDirectoryGroupProfile

// NewOktaActiveDirectoryGroupProfile instantiates a new OktaActiveDirectoryGroupProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaActiveDirectoryGroupProfile() *OktaActiveDirectoryGroupProfile {
	this := OktaActiveDirectoryGroupProfile{}
	return &this
}

// NewOktaActiveDirectoryGroupProfileWithDefaults instantiates a new OktaActiveDirectoryGroupProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaActiveDirectoryGroupProfileWithDefaults() *OktaActiveDirectoryGroupProfile {
	this := OktaActiveDirectoryGroupProfile{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OktaActiveDirectoryGroupProfile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaActiveDirectoryGroupProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OktaActiveDirectoryGroupProfile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OktaActiveDirectoryGroupProfile) SetDescription(v string) {
	o.Description = &v
}

// GetDn returns the Dn field value if set, zero value otherwise.
func (o *OktaActiveDirectoryGroupProfile) GetDn() string {
	if o == nil || IsNil(o.Dn) {
		var ret string
		return ret
	}
	return *o.Dn
}

// GetDnOk returns a tuple with the Dn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaActiveDirectoryGroupProfile) GetDnOk() (*string, bool) {
	if o == nil || IsNil(o.Dn) {
		return nil, false
	}
	return o.Dn, true
}

// HasDn returns a boolean if a field has been set.
func (o *OktaActiveDirectoryGroupProfile) HasDn() bool {
	if o != nil && !IsNil(o.Dn) {
		return true
	}

	return false
}

// SetDn gets a reference to the given string and assigns it to the Dn field.
func (o *OktaActiveDirectoryGroupProfile) SetDn(v string) {
	o.Dn = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *OktaActiveDirectoryGroupProfile) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaActiveDirectoryGroupProfile) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *OktaActiveDirectoryGroupProfile) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *OktaActiveDirectoryGroupProfile) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OktaActiveDirectoryGroupProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaActiveDirectoryGroupProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OktaActiveDirectoryGroupProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OktaActiveDirectoryGroupProfile) SetName(v string) {
	o.Name = &v
}

// GetSamAccountName returns the SamAccountName field value if set, zero value otherwise.
func (o *OktaActiveDirectoryGroupProfile) GetSamAccountName() string {
	if o == nil || IsNil(o.SamAccountName) {
		var ret string
		return ret
	}
	return *o.SamAccountName
}

// GetSamAccountNameOk returns a tuple with the SamAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaActiveDirectoryGroupProfile) GetSamAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.SamAccountName) {
		return nil, false
	}
	return o.SamAccountName, true
}

// HasSamAccountName returns a boolean if a field has been set.
func (o *OktaActiveDirectoryGroupProfile) HasSamAccountName() bool {
	if o != nil && !IsNil(o.SamAccountName) {
		return true
	}

	return false
}

// SetSamAccountName gets a reference to the given string and assigns it to the SamAccountName field.
func (o *OktaActiveDirectoryGroupProfile) SetSamAccountName(v string) {
	o.SamAccountName = &v
}

// GetWindowsDomainQualifiedName returns the WindowsDomainQualifiedName field value if set, zero value otherwise.
func (o *OktaActiveDirectoryGroupProfile) GetWindowsDomainQualifiedName() string {
	if o == nil || IsNil(o.WindowsDomainQualifiedName) {
		var ret string
		return ret
	}
	return *o.WindowsDomainQualifiedName
}

// GetWindowsDomainQualifiedNameOk returns a tuple with the WindowsDomainQualifiedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaActiveDirectoryGroupProfile) GetWindowsDomainQualifiedNameOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsDomainQualifiedName) {
		return nil, false
	}
	return o.WindowsDomainQualifiedName, true
}

// HasWindowsDomainQualifiedName returns a boolean if a field has been set.
func (o *OktaActiveDirectoryGroupProfile) HasWindowsDomainQualifiedName() bool {
	if o != nil && !IsNil(o.WindowsDomainQualifiedName) {
		return true
	}

	return false
}

// SetWindowsDomainQualifiedName gets a reference to the given string and assigns it to the WindowsDomainQualifiedName field.
func (o *OktaActiveDirectoryGroupProfile) SetWindowsDomainQualifiedName(v string) {
	o.WindowsDomainQualifiedName = &v
}

func (o OktaActiveDirectoryGroupProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaActiveDirectoryGroupProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Dn) {
		toSerialize["dn"] = o.Dn
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SamAccountName) {
		toSerialize["samAccountName"] = o.SamAccountName
	}
	if !IsNil(o.WindowsDomainQualifiedName) {
		toSerialize["windowsDomainQualifiedName"] = o.WindowsDomainQualifiedName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaActiveDirectoryGroupProfile) UnmarshalJSON(data []byte) (err error) {
	varOktaActiveDirectoryGroupProfile := _OktaActiveDirectoryGroupProfile{}

	err = json.Unmarshal(data, &varOktaActiveDirectoryGroupProfile)

	if err != nil {
		return err
	}

	*o = OktaActiveDirectoryGroupProfile(varOktaActiveDirectoryGroupProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "dn")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "name")
		delete(additionalProperties, "samAccountName")
		delete(additionalProperties, "windowsDomainQualifiedName")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaActiveDirectoryGroupProfile struct {
	value *OktaActiveDirectoryGroupProfile
	isSet bool
}

func (v NullableOktaActiveDirectoryGroupProfile) Get() *OktaActiveDirectoryGroupProfile {
	return v.value
}

func (v *NullableOktaActiveDirectoryGroupProfile) Set(val *OktaActiveDirectoryGroupProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaActiveDirectoryGroupProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaActiveDirectoryGroupProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaActiveDirectoryGroupProfile(val *OktaActiveDirectoryGroupProfile) *NullableOktaActiveDirectoryGroupProfile {
	return &NullableOktaActiveDirectoryGroupProfile{value: val, isSet: true}
}

func (v NullableOktaActiveDirectoryGroupProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaActiveDirectoryGroupProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
