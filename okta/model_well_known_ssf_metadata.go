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

// WellKnownSSFMetadata Metadata about Okta as a transmitter and relevant information for configuration.
type WellKnownSSFMetadata struct {
	// The URL of the SSF Stream configuration endpoint
	ConfigurationEndpoint *string `json:"configuration_endpoint,omitempty"`
	// An array of supported SET delivery methods
	DeliveryMethodsSupported []string `json:"delivery_methods_supported,omitempty"`
	// The issuer used in Security Event Tokens. This value is set as `iss` in the claim.
	Issuer *string `json:"issuer,omitempty"`
	// The URL of the JSON Web Key Set (JWKS) that contains the signing keys for validating the signatures of Security Event Tokens (SETs)
	JwksUri *string `json:"jwks_uri,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WellKnownSSFMetadata WellKnownSSFMetadata

// NewWellKnownSSFMetadata instantiates a new WellKnownSSFMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWellKnownSSFMetadata() *WellKnownSSFMetadata {
	this := WellKnownSSFMetadata{}
	return &this
}

// NewWellKnownSSFMetadataWithDefaults instantiates a new WellKnownSSFMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWellKnownSSFMetadataWithDefaults() *WellKnownSSFMetadata {
	this := WellKnownSSFMetadata{}
	return &this
}

// GetConfigurationEndpoint returns the ConfigurationEndpoint field value if set, zero value otherwise.
func (o *WellKnownSSFMetadata) GetConfigurationEndpoint() string {
	if o == nil || o.ConfigurationEndpoint == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationEndpoint
}

// GetConfigurationEndpointOk returns a tuple with the ConfigurationEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownSSFMetadata) GetConfigurationEndpointOk() (*string, bool) {
	if o == nil || o.ConfigurationEndpoint == nil {
		return nil, false
	}
	return o.ConfigurationEndpoint, true
}

// HasConfigurationEndpoint returns a boolean if a field has been set.
func (o *WellKnownSSFMetadata) HasConfigurationEndpoint() bool {
	if o != nil && o.ConfigurationEndpoint != nil {
		return true
	}

	return false
}

// SetConfigurationEndpoint gets a reference to the given string and assigns it to the ConfigurationEndpoint field.
func (o *WellKnownSSFMetadata) SetConfigurationEndpoint(v string) {
	o.ConfigurationEndpoint = &v
}

// GetDeliveryMethodsSupported returns the DeliveryMethodsSupported field value if set, zero value otherwise.
func (o *WellKnownSSFMetadata) GetDeliveryMethodsSupported() []string {
	if o == nil || o.DeliveryMethodsSupported == nil {
		var ret []string
		return ret
	}
	return o.DeliveryMethodsSupported
}

// GetDeliveryMethodsSupportedOk returns a tuple with the DeliveryMethodsSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownSSFMetadata) GetDeliveryMethodsSupportedOk() ([]string, bool) {
	if o == nil || o.DeliveryMethodsSupported == nil {
		return nil, false
	}
	return o.DeliveryMethodsSupported, true
}

// HasDeliveryMethodsSupported returns a boolean if a field has been set.
func (o *WellKnownSSFMetadata) HasDeliveryMethodsSupported() bool {
	if o != nil && o.DeliveryMethodsSupported != nil {
		return true
	}

	return false
}

// SetDeliveryMethodsSupported gets a reference to the given []string and assigns it to the DeliveryMethodsSupported field.
func (o *WellKnownSSFMetadata) SetDeliveryMethodsSupported(v []string) {
	o.DeliveryMethodsSupported = v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *WellKnownSSFMetadata) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownSSFMetadata) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *WellKnownSSFMetadata) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *WellKnownSSFMetadata) SetIssuer(v string) {
	o.Issuer = &v
}

// GetJwksUri returns the JwksUri field value if set, zero value otherwise.
func (o *WellKnownSSFMetadata) GetJwksUri() string {
	if o == nil || o.JwksUri == nil {
		var ret string
		return ret
	}
	return *o.JwksUri
}

// GetJwksUriOk returns a tuple with the JwksUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WellKnownSSFMetadata) GetJwksUriOk() (*string, bool) {
	if o == nil || o.JwksUri == nil {
		return nil, false
	}
	return o.JwksUri, true
}

// HasJwksUri returns a boolean if a field has been set.
func (o *WellKnownSSFMetadata) HasJwksUri() bool {
	if o != nil && o.JwksUri != nil {
		return true
	}

	return false
}

// SetJwksUri gets a reference to the given string and assigns it to the JwksUri field.
func (o *WellKnownSSFMetadata) SetJwksUri(v string) {
	o.JwksUri = &v
}

func (o WellKnownSSFMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigurationEndpoint != nil {
		toSerialize["configuration_endpoint"] = o.ConfigurationEndpoint
	}
	if o.DeliveryMethodsSupported != nil {
		toSerialize["delivery_methods_supported"] = o.DeliveryMethodsSupported
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.JwksUri != nil {
		toSerialize["jwks_uri"] = o.JwksUri
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WellKnownSSFMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varWellKnownSSFMetadata := _WellKnownSSFMetadata{}

	err = json.Unmarshal(bytes, &varWellKnownSSFMetadata)
	if err == nil {
		*o = WellKnownSSFMetadata(varWellKnownSSFMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "configuration_endpoint")
		delete(additionalProperties, "delivery_methods_supported")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "jwks_uri")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableWellKnownSSFMetadata struct {
	value *WellKnownSSFMetadata
	isSet bool
}

func (v NullableWellKnownSSFMetadata) Get() *WellKnownSSFMetadata {
	return v.value
}

func (v *NullableWellKnownSSFMetadata) Set(val *WellKnownSSFMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableWellKnownSSFMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableWellKnownSSFMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWellKnownSSFMetadata(val *WellKnownSSFMetadata) *NullableWellKnownSSFMetadata {
	return &NullableWellKnownSSFMetadata{value: val, isSet: true}
}

func (v NullableWellKnownSSFMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWellKnownSSFMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

