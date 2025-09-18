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

// checks if the SAMLPayLoadDataContextAllOfProtocolIssuer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SAMLPayLoadDataContextAllOfProtocolIssuer{}

// SAMLPayLoadDataContextAllOfProtocolIssuer struct for SAMLPayLoadDataContextAllOfProtocolIssuer
type SAMLPayLoadDataContextAllOfProtocolIssuer struct {
	// The unique identifier of the issuer that provided the SAML assertion
	Id *string `json:"id,omitempty"`
	// The name of the issuer that provided the SAML assertion
	Name *string `json:"name,omitempty"`
	// The base URI of the SAML endpoint that's used to assert the authorization
	Uri                  *string `json:"uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SAMLPayLoadDataContextAllOfProtocolIssuer SAMLPayLoadDataContextAllOfProtocolIssuer

// NewSAMLPayLoadDataContextAllOfProtocolIssuer instantiates a new SAMLPayLoadDataContextAllOfProtocolIssuer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLPayLoadDataContextAllOfProtocolIssuer() *SAMLPayLoadDataContextAllOfProtocolIssuer {
	this := SAMLPayLoadDataContextAllOfProtocolIssuer{}
	return &this
}

// NewSAMLPayLoadDataContextAllOfProtocolIssuerWithDefaults instantiates a new SAMLPayLoadDataContextAllOfProtocolIssuer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLPayLoadDataContextAllOfProtocolIssuerWithDefaults() *SAMLPayLoadDataContextAllOfProtocolIssuer {
	this := SAMLPayLoadDataContextAllOfProtocolIssuer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) SetName(v string) {
	o.Name = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) SetUri(v string) {
	o.Uri = &v
}

func (o SAMLPayLoadDataContextAllOfProtocolIssuer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SAMLPayLoadDataContextAllOfProtocolIssuer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SAMLPayLoadDataContextAllOfProtocolIssuer) UnmarshalJSON(data []byte) (err error) {
	varSAMLPayLoadDataContextAllOfProtocolIssuer := _SAMLPayLoadDataContextAllOfProtocolIssuer{}

	err = json.Unmarshal(data, &varSAMLPayLoadDataContextAllOfProtocolIssuer)

	if err != nil {
		return err
	}

	*o = SAMLPayLoadDataContextAllOfProtocolIssuer(varSAMLPayLoadDataContextAllOfProtocolIssuer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "uri")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSAMLPayLoadDataContextAllOfProtocolIssuer struct {
	value *SAMLPayLoadDataContextAllOfProtocolIssuer
	isSet bool
}

func (v NullableSAMLPayLoadDataContextAllOfProtocolIssuer) Get() *SAMLPayLoadDataContextAllOfProtocolIssuer {
	return v.value
}

func (v *NullableSAMLPayLoadDataContextAllOfProtocolIssuer) Set(val *SAMLPayLoadDataContextAllOfProtocolIssuer) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLPayLoadDataContextAllOfProtocolIssuer) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLPayLoadDataContextAllOfProtocolIssuer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLPayLoadDataContextAllOfProtocolIssuer(val *SAMLPayLoadDataContextAllOfProtocolIssuer) *NullableSAMLPayLoadDataContextAllOfProtocolIssuer {
	return &NullableSAMLPayLoadDataContextAllOfProtocolIssuer{value: val, isSet: true}
}

func (v NullableSAMLPayLoadDataContextAllOfProtocolIssuer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLPayLoadDataContextAllOfProtocolIssuer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
