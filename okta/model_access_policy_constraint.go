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

// AccessPolicyConstraint struct for AccessPolicyConstraint
type AccessPolicyConstraint struct {
	// This property specifies the precise authenticator and method for authentication. <x-lifecycle class=\"oie\"></x-lifecycle>
	AuthenticationMethods []AuthenticationMethodObject `json:"authenticationMethods,omitempty"`
	// This property specifies the precise authenticator and method to exclude from authentication. <x-lifecycle class=\"oie\"></x-lifecycle>
	ExcludedAuthenticationMethods []AuthenticationMethodObject `json:"excludedAuthenticationMethods,omitempty"`
	// The Authenticator methods that are permitted
	Methods []string `json:"methods,omitempty"`
	// The duration after which the user must re-authenticate regardless of user activity. This re-authentication interval overrides the Verification Method object's `reauthenticateIn` interval. The supported values use ISO 8601 period format for recurring time intervals (for example, `PT1H`).
	ReauthenticateIn *string `json:"reauthenticateIn,omitempty"`
	// This property indicates whether the knowledge or possession factor is required by the assurance. It's optional in the request, but is always returned in the response. By default, this field is `true`. If the knowledge or possession constraint has values for `excludedAuthenticationMethods` the `required` value is false. <x-lifecycle class=\"oie\"></x-lifecycle>
	Required *bool `json:"required,omitempty"`
	// The Authenticator types that are permitted
	Types []string `json:"types,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessPolicyConstraint AccessPolicyConstraint

// NewAccessPolicyConstraint instantiates a new AccessPolicyConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessPolicyConstraint() *AccessPolicyConstraint {
	this := AccessPolicyConstraint{}
	return &this
}

// NewAccessPolicyConstraintWithDefaults instantiates a new AccessPolicyConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessPolicyConstraintWithDefaults() *AccessPolicyConstraint {
	this := AccessPolicyConstraint{}
	return &this
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *AccessPolicyConstraint) GetAuthenticationMethods() []AuthenticationMethodObject {
	if o == nil || o.AuthenticationMethods == nil {
		var ret []AuthenticationMethodObject
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetAuthenticationMethodsOk() ([]AuthenticationMethodObject, bool) {
	if o == nil || o.AuthenticationMethods == nil {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasAuthenticationMethods() bool {
	if o != nil && o.AuthenticationMethods != nil {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []AuthenticationMethodObject and assigns it to the AuthenticationMethods field.
func (o *AccessPolicyConstraint) SetAuthenticationMethods(v []AuthenticationMethodObject) {
	o.AuthenticationMethods = v
}

// GetExcludedAuthenticationMethods returns the ExcludedAuthenticationMethods field value if set, zero value otherwise.
func (o *AccessPolicyConstraint) GetExcludedAuthenticationMethods() []AuthenticationMethodObject {
	if o == nil || o.ExcludedAuthenticationMethods == nil {
		var ret []AuthenticationMethodObject
		return ret
	}
	return o.ExcludedAuthenticationMethods
}

// GetExcludedAuthenticationMethodsOk returns a tuple with the ExcludedAuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetExcludedAuthenticationMethodsOk() ([]AuthenticationMethodObject, bool) {
	if o == nil || o.ExcludedAuthenticationMethods == nil {
		return nil, false
	}
	return o.ExcludedAuthenticationMethods, true
}

// HasExcludedAuthenticationMethods returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasExcludedAuthenticationMethods() bool {
	if o != nil && o.ExcludedAuthenticationMethods != nil {
		return true
	}

	return false
}

// SetExcludedAuthenticationMethods gets a reference to the given []AuthenticationMethodObject and assigns it to the ExcludedAuthenticationMethods field.
func (o *AccessPolicyConstraint) SetExcludedAuthenticationMethods(v []AuthenticationMethodObject) {
	o.ExcludedAuthenticationMethods = v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *AccessPolicyConstraint) GetMethods() []string {
	if o == nil || o.Methods == nil {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetMethodsOk() ([]string, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []string and assigns it to the Methods field.
func (o *AccessPolicyConstraint) SetMethods(v []string) {
	o.Methods = v
}

// GetReauthenticateIn returns the ReauthenticateIn field value if set, zero value otherwise.
func (o *AccessPolicyConstraint) GetReauthenticateIn() string {
	if o == nil || o.ReauthenticateIn == nil {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetReauthenticateInOk() (*string, bool) {
	if o == nil || o.ReauthenticateIn == nil {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasReauthenticateIn() bool {
	if o != nil && o.ReauthenticateIn != nil {
		return true
	}

	return false
}

// SetReauthenticateIn gets a reference to the given string and assigns it to the ReauthenticateIn field.
func (o *AccessPolicyConstraint) SetReauthenticateIn(v string) {
	o.ReauthenticateIn = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *AccessPolicyConstraint) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *AccessPolicyConstraint) SetRequired(v bool) {
	o.Required = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *AccessPolicyConstraint) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetTypesOk() ([]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *AccessPolicyConstraint) SetTypes(v []string) {
	o.Types = v
}

func (o AccessPolicyConstraint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationMethods != nil {
		toSerialize["authenticationMethods"] = o.AuthenticationMethods
	}
	if o.ExcludedAuthenticationMethods != nil {
		toSerialize["excludedAuthenticationMethods"] = o.ExcludedAuthenticationMethods
	}
	if o.Methods != nil {
		toSerialize["methods"] = o.Methods
	}
	if o.ReauthenticateIn != nil {
		toSerialize["reauthenticateIn"] = o.ReauthenticateIn
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessPolicyConstraint) UnmarshalJSON(bytes []byte) (err error) {
	varAccessPolicyConstraint := _AccessPolicyConstraint{}

	err = json.Unmarshal(bytes, &varAccessPolicyConstraint)
	if err == nil {
		*o = AccessPolicyConstraint(varAccessPolicyConstraint)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticationMethods")
		delete(additionalProperties, "excludedAuthenticationMethods")
		delete(additionalProperties, "methods")
		delete(additionalProperties, "reauthenticateIn")
		delete(additionalProperties, "required")
		delete(additionalProperties, "types")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAccessPolicyConstraint struct {
	value *AccessPolicyConstraint
	isSet bool
}

func (v NullableAccessPolicyConstraint) Get() *AccessPolicyConstraint {
	return v.value
}

func (v *NullableAccessPolicyConstraint) Set(val *AccessPolicyConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessPolicyConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessPolicyConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessPolicyConstraint(val *AccessPolicyConstraint) *NullableAccessPolicyConstraint {
	return &NullableAccessPolicyConstraint{value: val, isSet: true}
}

func (v NullableAccessPolicyConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessPolicyConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

