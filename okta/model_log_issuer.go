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

// checks if the LogIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogIssuer{}

// LogIssuer Describes the issuer of the authorization server when the authentication is performed through OAuth. This is the location where well-known resources regarding the details of the authorization servers are published.
type LogIssuer struct {
	// Varies depending on the type of authentication. If authentication is SAML 2.0, `id` is the issuer in the SAML assertion. For social login, `id` is the issuer of the token.
	Id *string `json:"id,omitempty"`
	// Information on the `issuer` and source of the SAML assertion or token
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogIssuer LogIssuer

// NewLogIssuer instantiates a new LogIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIssuer() *LogIssuer {
	this := LogIssuer{}
	return &this
}

// NewLogIssuerWithDefaults instantiates a new LogIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogIssuerWithDefaults() *LogIssuer {
	this := LogIssuer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogIssuer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIssuer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogIssuer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogIssuer) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogIssuer) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIssuer) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogIssuer) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogIssuer) SetType(v string) {
	o.Type = &v
}

func (o LogIssuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogIssuer) UnmarshalJSON(data []byte) (err error) {
	varLogIssuer := _LogIssuer{}

	err = json.Unmarshal(data, &varLogIssuer)

	if err != nil {
		return err
	}

	*o = LogIssuer(varLogIssuer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogIssuer struct {
	value *LogIssuer
	isSet bool
}

func (v NullableLogIssuer) Get() *LogIssuer {
	return v.value
}

func (v *NullableLogIssuer) Set(val *LogIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableLogIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableLogIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogIssuer(val *LogIssuer) *NullableLogIssuer {
	return &NullableLogIssuer{value: val, isSet: true}
}

func (v NullableLogIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
