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

// checks if the ApplicationUniversalLogout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationUniversalLogout{}

// ApplicationUniversalLogout <div class=\"x-lifecycle-container\"><x-lifecycle class=\"oie\"></x-lifecycle></div> Universal Logout properties for the app. These properties are only returned and can't be updated.
type ApplicationUniversalLogout struct {
	// Indicates whether the app uses a shared identity stack that may cause the user to sign out of other apps by the same company
	IdentityStack *string `json:"identityStack,omitempty"`
	// The protocol used for Universal Logout
	Protocol *string `json:"protocol,omitempty"`
	// Universal Logout status for the app instance
	Status *string `json:"status,omitempty"`
	// Indicates whether the app supports full or partial Universal Logout (UL).
	SupportType          *string `json:"supportType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationUniversalLogout ApplicationUniversalLogout

// NewApplicationUniversalLogout instantiates a new ApplicationUniversalLogout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationUniversalLogout() *ApplicationUniversalLogout {
	this := ApplicationUniversalLogout{}
	return &this
}

// NewApplicationUniversalLogoutWithDefaults instantiates a new ApplicationUniversalLogout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationUniversalLogoutWithDefaults() *ApplicationUniversalLogout {
	this := ApplicationUniversalLogout{}
	return &this
}

// GetIdentityStack returns the IdentityStack field value if set, zero value otherwise.
func (o *ApplicationUniversalLogout) GetIdentityStack() string {
	if o == nil || IsNil(o.IdentityStack) {
		var ret string
		return ret
	}
	return *o.IdentityStack
}

// GetIdentityStackOk returns a tuple with the IdentityStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUniversalLogout) GetIdentityStackOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityStack) {
		return nil, false
	}
	return o.IdentityStack, true
}

// HasIdentityStack returns a boolean if a field has been set.
func (o *ApplicationUniversalLogout) HasIdentityStack() bool {
	if o != nil && !IsNil(o.IdentityStack) {
		return true
	}

	return false
}

// SetIdentityStack gets a reference to the given string and assigns it to the IdentityStack field.
func (o *ApplicationUniversalLogout) SetIdentityStack(v string) {
	o.IdentityStack = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApplicationUniversalLogout) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUniversalLogout) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApplicationUniversalLogout) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApplicationUniversalLogout) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplicationUniversalLogout) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUniversalLogout) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplicationUniversalLogout) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplicationUniversalLogout) SetStatus(v string) {
	o.Status = &v
}

// GetSupportType returns the SupportType field value if set, zero value otherwise.
func (o *ApplicationUniversalLogout) GetSupportType() string {
	if o == nil || IsNil(o.SupportType) {
		var ret string
		return ret
	}
	return *o.SupportType
}

// GetSupportTypeOk returns a tuple with the SupportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationUniversalLogout) GetSupportTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SupportType) {
		return nil, false
	}
	return o.SupportType, true
}

// HasSupportType returns a boolean if a field has been set.
func (o *ApplicationUniversalLogout) HasSupportType() bool {
	if o != nil && !IsNil(o.SupportType) {
		return true
	}

	return false
}

// SetSupportType gets a reference to the given string and assigns it to the SupportType field.
func (o *ApplicationUniversalLogout) SetSupportType(v string) {
	o.SupportType = &v
}

func (o ApplicationUniversalLogout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationUniversalLogout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdentityStack) {
		toSerialize["identityStack"] = o.IdentityStack
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SupportType) {
		toSerialize["supportType"] = o.SupportType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationUniversalLogout) UnmarshalJSON(data []byte) (err error) {
	varApplicationUniversalLogout := _ApplicationUniversalLogout{}

	err = json.Unmarshal(data, &varApplicationUniversalLogout)

	if err != nil {
		return err
	}

	*o = ApplicationUniversalLogout(varApplicationUniversalLogout)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "identityStack")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "status")
		delete(additionalProperties, "supportType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationUniversalLogout struct {
	value *ApplicationUniversalLogout
	isSet bool
}

func (v NullableApplicationUniversalLogout) Get() *ApplicationUniversalLogout {
	return v.value
}

func (v *NullableApplicationUniversalLogout) Set(val *ApplicationUniversalLogout) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationUniversalLogout) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationUniversalLogout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationUniversalLogout(val *ApplicationUniversalLogout) *NullableApplicationUniversalLogout {
	return &NullableApplicationUniversalLogout{value: val, isSet: true}
}

func (v NullableApplicationUniversalLogout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationUniversalLogout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
