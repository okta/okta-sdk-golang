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

// checks if the AccessPolicyConstraint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessPolicyConstraint{}

// AccessPolicyConstraint Consists of a `POSSESSION` constraint, a `KNOWLEDGE` constraint, or both. You can't configure an `INHERENCE` constraint, but an inherence factor can satisfy the second part of a 2FA assurance if no other constraints are specified. Constraints are logically evaluated such that only one `constraint` object needs to be satisfied, but within a `constraint` object, each `constraint` property must be satisfied.
type AccessPolicyConstraint struct {
	// This property specifies the precise authenticator and method for authentication. <x-lifecycle class=\"oie\"></x-lifecycle>
	AuthenticationMethods []AuthenticationMethodObject `json:"authenticationMethods,omitempty"`
	// This property specifies the precise authenticator and method to exclude from authentication. <x-lifecycle class=\"oie\"></x-lifecycle>
	ExcludedAuthenticationMethods []AuthenticationMethodObject `json:"excludedAuthenticationMethods,omitempty"`
	// The authenticator methods that are permitted
	Methods []string `json:"methods,omitempty"`
	// The duration after which the user must re-authenticate regardless of user activity. This re-authentication interval overrides the Verification Method object's `reauthenticateIn` interval. The supported values use ISO 8601 period format for recurring time intervals (for example, `PT1H`).
	ReauthenticateIn *string `json:"reauthenticateIn,omitempty"`
	// This property indicates whether the knowledge or possession factor is required by the assurance. It's optional in the request, but is always returned in the response. By default, this field is `true`. If the knowledge or possession constraint has values for `excludedAuthenticationMethods` the `required` value is false. <x-lifecycle class=\"oie\"></x-lifecycle>
	Required *bool `json:"required,omitempty"`
	// The authenticator types that are permitted
	Types                []string `json:"types,omitempty"`
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
	if o == nil || IsNil(o.AuthenticationMethods) {
		var ret []AuthenticationMethodObject
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetAuthenticationMethodsOk() ([]AuthenticationMethodObject, bool) {
	if o == nil || IsNil(o.AuthenticationMethods) {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasAuthenticationMethods() bool {
	if o != nil && !IsNil(o.AuthenticationMethods) {
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
	if o == nil || IsNil(o.ExcludedAuthenticationMethods) {
		var ret []AuthenticationMethodObject
		return ret
	}
	return o.ExcludedAuthenticationMethods
}

// GetExcludedAuthenticationMethodsOk returns a tuple with the ExcludedAuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetExcludedAuthenticationMethodsOk() ([]AuthenticationMethodObject, bool) {
	if o == nil || IsNil(o.ExcludedAuthenticationMethods) {
		return nil, false
	}
	return o.ExcludedAuthenticationMethods, true
}

// HasExcludedAuthenticationMethods returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasExcludedAuthenticationMethods() bool {
	if o != nil && !IsNil(o.ExcludedAuthenticationMethods) {
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
	if o == nil || IsNil(o.Methods) {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.Methods) {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasMethods() bool {
	if o != nil && !IsNil(o.Methods) {
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
	if o == nil || IsNil(o.ReauthenticateIn) {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetReauthenticateInOk() (*string, bool) {
	if o == nil || IsNil(o.ReauthenticateIn) {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasReauthenticateIn() bool {
	if o != nil && !IsNil(o.ReauthenticateIn) {
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
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
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
	if o == nil || IsNil(o.Types) {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessPolicyConstraint) GetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *AccessPolicyConstraint) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *AccessPolicyConstraint) SetTypes(v []string) {
	o.Types = v
}

func (o AccessPolicyConstraint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessPolicyConstraint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationMethods) {
		toSerialize["authenticationMethods"] = o.AuthenticationMethods
	}
	if !IsNil(o.ExcludedAuthenticationMethods) {
		toSerialize["excludedAuthenticationMethods"] = o.ExcludedAuthenticationMethods
	}
	if !IsNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}
	if !IsNil(o.ReauthenticateIn) {
		toSerialize["reauthenticateIn"] = o.ReauthenticateIn
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AccessPolicyConstraint) UnmarshalJSON(data []byte) (err error) {
	varAccessPolicyConstraint := _AccessPolicyConstraint{}

	err = json.Unmarshal(data, &varAccessPolicyConstraint)

	if err != nil {
		return err
	}

	*o = AccessPolicyConstraint(varAccessPolicyConstraint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "authenticationMethods")
		delete(additionalProperties, "excludedAuthenticationMethods")
		delete(additionalProperties, "methods")
		delete(additionalProperties, "reauthenticateIn")
		delete(additionalProperties, "required")
		delete(additionalProperties, "types")
		o.AdditionalProperties = additionalProperties
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
