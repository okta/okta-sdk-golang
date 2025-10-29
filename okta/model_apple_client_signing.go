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

// checks if the AppleClientSigning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppleClientSigning{}

// AppleClientSigning Information used to generate the secret JSON Web Token for the token requests to Apple IdP > **Note:** The `privateKey` property is required for a CREATE request. For an UPDATE request, it can be null and keeps the existing value if it's null. The `privateKey` property isn't returned for LIST and GET requests or UPDATE requests if it's null.
type AppleClientSigning struct {
	// The key ID that you obtained from Apple when you created the private key for the client
	Kid *string `json:"kid,omitempty"`
	// The PKCS \\#8 encoded private key that you created for the client and downloaded from Apple
	PrivateKey *string `json:"privateKey,omitempty"`
	// The Team ID associated with your Apple developer account
	TeamId               *string `json:"teamId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppleClientSigning AppleClientSigning

// NewAppleClientSigning instantiates a new AppleClientSigning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppleClientSigning() *AppleClientSigning {
	this := AppleClientSigning{}
	return &this
}

// NewAppleClientSigningWithDefaults instantiates a new AppleClientSigning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleClientSigningWithDefaults() *AppleClientSigning {
	this := AppleClientSigning{}
	return &this
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AppleClientSigning) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleClientSigning) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AppleClientSigning) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AppleClientSigning) SetKid(v string) {
	o.Kid = &v
}

// GetPrivateKey returns the PrivateKey field value if set, zero value otherwise.
func (o *AppleClientSigning) GetPrivateKey() string {
	if o == nil || IsNil(o.PrivateKey) {
		var ret string
		return ret
	}
	return *o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleClientSigning) GetPrivateKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateKey) {
		return nil, false
	}
	return o.PrivateKey, true
}

// HasPrivateKey returns a boolean if a field has been set.
func (o *AppleClientSigning) HasPrivateKey() bool {
	if o != nil && !IsNil(o.PrivateKey) {
		return true
	}

	return false
}

// SetPrivateKey gets a reference to the given string and assigns it to the PrivateKey field.
func (o *AppleClientSigning) SetPrivateKey(v string) {
	o.PrivateKey = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *AppleClientSigning) GetTeamId() string {
	if o == nil || IsNil(o.TeamId) {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppleClientSigning) GetTeamIdOk() (*string, bool) {
	if o == nil || IsNil(o.TeamId) {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *AppleClientSigning) HasTeamId() bool {
	if o != nil && !IsNil(o.TeamId) {
		return true
	}

	return false
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *AppleClientSigning) SetTeamId(v string) {
	o.TeamId = &v
}

func (o AppleClientSigning) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppleClientSigning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.PrivateKey) {
		toSerialize["privateKey"] = o.PrivateKey
	}
	if !IsNil(o.TeamId) {
		toSerialize["teamId"] = o.TeamId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppleClientSigning) UnmarshalJSON(data []byte) (err error) {
	varAppleClientSigning := _AppleClientSigning{}

	err = json.Unmarshal(data, &varAppleClientSigning)

	if err != nil {
		return err
	}

	*o = AppleClientSigning(varAppleClientSigning)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "kid")
		delete(additionalProperties, "privateKey")
		delete(additionalProperties, "teamId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppleClientSigning struct {
	value *AppleClientSigning
	isSet bool
}

func (v NullableAppleClientSigning) Get() *AppleClientSigning {
	return v.value
}

func (v *NullableAppleClientSigning) Set(val *AppleClientSigning) {
	v.value = val
	v.isSet = true
}

func (v NullableAppleClientSigning) IsSet() bool {
	return v.isSet
}

func (v *NullableAppleClientSigning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppleClientSigning(val *AppleClientSigning) *NullableAppleClientSigning {
	return &NullableAppleClientSigning{value: val, isSet: true}
}

func (v NullableAppleClientSigning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppleClientSigning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
