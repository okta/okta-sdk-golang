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

// AAGUIDGroupObject <x-lifecycle class=\"ea\"></x-lifecycle> The AAGUID Group object supports the Early Access (Self-Service) Allow List for FIDO2 (WebAuthn) Authenticators feature. Enable the feature for your org from the **Settings** > **Features** page in the Admin Console.  This feature has several limitations when enrolling a security key:   - Enrollment is currently unsupported on Firefox.   - Enrollment is currently unsupported on Chrome if User Verification is set to DISCOURAGED and a PIN is set on the security key.   - If prompted during enrollment, users must allow Okta to see the make and model of the security key. 
type AAGUIDGroupObject struct {
	// A list of YubiKey hardware FIDO2 Authenticator Attestation Global Unique Identifiers (AAGUIDs). The available [AAGUIDs](https://support.yubico.com/hc/en-us/articles/360016648959-YubiKey-Hardware-FIDO2-AAGUIDs) (opens new window) are provided by the FIDO Alliance Metadata Service.
	Aaguids []string `json:"aaguids,omitempty"`
	// A name to identify the group of YubiKey hardware FIDO2 AAGUIDs
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AAGUIDGroupObject AAGUIDGroupObject

// NewAAGUIDGroupObject instantiates a new AAGUIDGroupObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAAGUIDGroupObject() *AAGUIDGroupObject {
	this := AAGUIDGroupObject{}
	return &this
}

// NewAAGUIDGroupObjectWithDefaults instantiates a new AAGUIDGroupObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAAGUIDGroupObjectWithDefaults() *AAGUIDGroupObject {
	this := AAGUIDGroupObject{}
	return &this
}

// GetAaguids returns the Aaguids field value if set, zero value otherwise.
func (o *AAGUIDGroupObject) GetAaguids() []string {
	if o == nil || o.Aaguids == nil {
		var ret []string
		return ret
	}
	return o.Aaguids
}

// GetAaguidsOk returns a tuple with the Aaguids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AAGUIDGroupObject) GetAaguidsOk() ([]string, bool) {
	if o == nil || o.Aaguids == nil {
		return nil, false
	}
	return o.Aaguids, true
}

// HasAaguids returns a boolean if a field has been set.
func (o *AAGUIDGroupObject) HasAaguids() bool {
	if o != nil && o.Aaguids != nil {
		return true
	}

	return false
}

// SetAaguids gets a reference to the given []string and assigns it to the Aaguids field.
func (o *AAGUIDGroupObject) SetAaguids(v []string) {
	o.Aaguids = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AAGUIDGroupObject) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AAGUIDGroupObject) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AAGUIDGroupObject) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AAGUIDGroupObject) SetName(v string) {
	o.Name = &v
}

func (o AAGUIDGroupObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aaguids != nil {
		toSerialize["aaguids"] = o.Aaguids
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AAGUIDGroupObject) UnmarshalJSON(bytes []byte) (err error) {
	varAAGUIDGroupObject := _AAGUIDGroupObject{}

	err = json.Unmarshal(bytes, &varAAGUIDGroupObject)
	if err == nil {
		*o = AAGUIDGroupObject(varAAGUIDGroupObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "aaguids")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAAGUIDGroupObject struct {
	value *AAGUIDGroupObject
	isSet bool
}

func (v NullableAAGUIDGroupObject) Get() *AAGUIDGroupObject {
	return v.value
}

func (v *NullableAAGUIDGroupObject) Set(val *AAGUIDGroupObject) {
	v.value = val
	v.isSet = true
}

func (v NullableAAGUIDGroupObject) IsSet() bool {
	return v.isSet
}

func (v *NullableAAGUIDGroupObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAAGUIDGroupObject(val *AAGUIDGroupObject) *NullableAAGUIDGroupObject {
	return &NullableAAGUIDGroupObject{value: val, isSet: true}
}

func (v NullableAAGUIDGroupObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAAGUIDGroupObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

