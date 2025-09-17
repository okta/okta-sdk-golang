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

// checks if the AuthenticatorLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticatorLinks{}

// AuthenticatorLinks struct for AuthenticatorLinks
type AuthenticatorLinks struct {
	Self       *HrefObjectSelfLink       `json:"self,omitempty"`
	Activate   *HrefObjectActivateLink   `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	// Link to authenticator methods
	Methods              *HrefObject `json:"methods,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthenticatorLinks AuthenticatorLinks

// NewAuthenticatorLinks instantiates a new AuthenticatorLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticatorLinks() *AuthenticatorLinks {
	this := AuthenticatorLinks{}
	return &this
}

// NewAuthenticatorLinksWithDefaults instantiates a new AuthenticatorLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticatorLinksWithDefaults() *AuthenticatorLinks {
	this := AuthenticatorLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthenticatorLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthenticatorLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *AuthenticatorLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *AuthenticatorLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || IsNil(o.Activate) {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *AuthenticatorLinks) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *AuthenticatorLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *AuthenticatorLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *AuthenticatorLinks) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *AuthenticatorLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *AuthenticatorLinks) GetMethods() HrefObject {
	if o == nil || IsNil(o.Methods) {
		var ret HrefObject
		return ret
	}
	return *o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticatorLinks) GetMethodsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *AuthenticatorLinks) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given HrefObject and assigns it to the Methods field.
func (o *AuthenticatorLinks) SetMethods(v HrefObject) {
	o.Methods = &v
}

func (o AuthenticatorLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticatorLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthenticatorLinks) UnmarshalJSON(data []byte) (err error) {
	varAuthenticatorLinks := _AuthenticatorLinks{}

	err = json.Unmarshal(data, &varAuthenticatorLinks)

	if err != nil {
		return err
	}

	*o = AuthenticatorLinks(varAuthenticatorLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "methods")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthenticatorLinks struct {
	value *AuthenticatorLinks
	isSet bool
}

func (v NullableAuthenticatorLinks) Get() *AuthenticatorLinks {
	return v.value
}

func (v *NullableAuthenticatorLinks) Set(val *AuthenticatorLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticatorLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticatorLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticatorLinks(val *AuthenticatorLinks) *NullableAuthenticatorLinks {
	return &NullableAuthenticatorLinks{value: val, isSet: true}
}

func (v NullableAuthenticatorLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticatorLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
