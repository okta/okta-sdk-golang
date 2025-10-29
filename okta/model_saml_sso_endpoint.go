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

// checks if the SamlSsoEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlSsoEndpoint{}

// SamlSsoEndpoint IdP's `SingleSignOnService` endpoint where Okta sends an `<AuthnRequest>` message
type SamlSsoEndpoint struct {
	Binding *string `json:"binding,omitempty"`
	// URI reference that indicates the address to which the `<AuthnRequest>` message is sent. The `destination` property is required if request signatures are specified. See [SAML 2.0 Request Algorithm object](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path=protocol/0/algorithms/request&t=request).
	Destination *string `json:"destination,omitempty"`
	// URL of the binding-specific endpoint to send an `<AuthnRequest>` message to the IdP. The value of `url` defaults to the same value as the `sso` endpoint if omitted during creation of a new IdP instance. The `url` should be the same value as the `Location` attribute for a published binding in the IdP's SAML Metadata `IDPSSODescriptor`.
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlSsoEndpoint SamlSsoEndpoint

// NewSamlSsoEndpoint instantiates a new SamlSsoEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlSsoEndpoint() *SamlSsoEndpoint {
	this := SamlSsoEndpoint{}
	return &this
}

// NewSamlSsoEndpointWithDefaults instantiates a new SamlSsoEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlSsoEndpointWithDefaults() *SamlSsoEndpoint {
	this := SamlSsoEndpoint{}
	return &this
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *SamlSsoEndpoint) GetBinding() string {
	if o == nil || IsNil(o.Binding) {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSsoEndpoint) GetBindingOk() (*string, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *SamlSsoEndpoint) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *SamlSsoEndpoint) SetBinding(v string) {
	o.Binding = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *SamlSsoEndpoint) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSsoEndpoint) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *SamlSsoEndpoint) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *SamlSsoEndpoint) SetDestination(v string) {
	o.Destination = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SamlSsoEndpoint) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlSsoEndpoint) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SamlSsoEndpoint) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SamlSsoEndpoint) SetUrl(v string) {
	o.Url = &v
}

func (o SamlSsoEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlSsoEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlSsoEndpoint) UnmarshalJSON(data []byte) (err error) {
	varSamlSsoEndpoint := _SamlSsoEndpoint{}

	err = json.Unmarshal(data, &varSamlSsoEndpoint)

	if err != nil {
		return err
	}

	*o = SamlSsoEndpoint(varSamlSsoEndpoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "binding")
		delete(additionalProperties, "destination")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlSsoEndpoint struct {
	value *SamlSsoEndpoint
	isSet bool
}

func (v NullableSamlSsoEndpoint) Get() *SamlSsoEndpoint {
	return v.value
}

func (v *NullableSamlSsoEndpoint) Set(val *SamlSsoEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlSsoEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlSsoEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlSsoEndpoint(val *SamlSsoEndpoint) *NullableSamlSsoEndpoint {
	return &NullableSamlSsoEndpoint{value: val, isSet: true}
}

func (v NullableSamlSsoEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlSsoEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
