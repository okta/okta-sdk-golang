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

// PossessionConstraint struct for PossessionConstraint
type PossessionConstraint struct {
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
	// Indicates if device-bound Factors are required. This property is only set for `POSSESSION` constraints.
	DeviceBound *string `json:"deviceBound,omitempty"`
	// Indicates if any secrets or private keys used during authentication must be hardware protected and not exportable. This property is only set for `POSSESSION` constraints.
	HardwareProtection *string `json:"hardwareProtection,omitempty"`
	// Indicates if phishing-resistant Factors are required. This property is only set for `POSSESSION` constraints.
	PhishingResistant *string `json:"phishingResistant,omitempty"`
	// Indicates if the user needs to approve an Okta Verify prompt or provide biometrics (meets NIST AAL2 requirements). This property is only set for `POSSESSION` constraints.
	UserPresence *string `json:"userPresence,omitempty"`
	// Indicates the user interaction requirement (PIN or biometrics) to ensure verification of a possession factor
	UserVerification *string `json:"userVerification,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PossessionConstraint PossessionConstraint

// NewPossessionConstraint instantiates a new PossessionConstraint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPossessionConstraint() *PossessionConstraint {
	this := PossessionConstraint{}
	var deviceBound string = "OPTIONAL"
	this.DeviceBound = &deviceBound
	var hardwareProtection string = "OPTIONAL"
	this.HardwareProtection = &hardwareProtection
	var phishingResistant string = "OPTIONAL"
	this.PhishingResistant = &phishingResistant
	var userPresence string = "REQUIRED"
	this.UserPresence = &userPresence
	var userVerification string = "OPTIONAL"
	this.UserVerification = &userVerification
	return &this
}

// NewPossessionConstraintWithDefaults instantiates a new PossessionConstraint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPossessionConstraintWithDefaults() *PossessionConstraint {
	this := PossessionConstraint{}
	var deviceBound string = "OPTIONAL"
	this.DeviceBound = &deviceBound
	var hardwareProtection string = "OPTIONAL"
	this.HardwareProtection = &hardwareProtection
	var phishingResistant string = "OPTIONAL"
	this.PhishingResistant = &phishingResistant
	var userPresence string = "REQUIRED"
	this.UserPresence = &userPresence
	var userVerification string = "OPTIONAL"
	this.UserVerification = &userVerification
	return &this
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *PossessionConstraint) GetAuthenticationMethods() []AuthenticationMethodObject {
	if o == nil || o.AuthenticationMethods == nil {
		var ret []AuthenticationMethodObject
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetAuthenticationMethodsOk() ([]AuthenticationMethodObject, bool) {
	if o == nil || o.AuthenticationMethods == nil {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *PossessionConstraint) HasAuthenticationMethods() bool {
	if o != nil && o.AuthenticationMethods != nil {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []AuthenticationMethodObject and assigns it to the AuthenticationMethods field.
func (o *PossessionConstraint) SetAuthenticationMethods(v []AuthenticationMethodObject) {
	o.AuthenticationMethods = v
}

// GetExcludedAuthenticationMethods returns the ExcludedAuthenticationMethods field value if set, zero value otherwise.
func (o *PossessionConstraint) GetExcludedAuthenticationMethods() []AuthenticationMethodObject {
	if o == nil || o.ExcludedAuthenticationMethods == nil {
		var ret []AuthenticationMethodObject
		return ret
	}
	return o.ExcludedAuthenticationMethods
}

// GetExcludedAuthenticationMethodsOk returns a tuple with the ExcludedAuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetExcludedAuthenticationMethodsOk() ([]AuthenticationMethodObject, bool) {
	if o == nil || o.ExcludedAuthenticationMethods == nil {
		return nil, false
	}
	return o.ExcludedAuthenticationMethods, true
}

// HasExcludedAuthenticationMethods returns a boolean if a field has been set.
func (o *PossessionConstraint) HasExcludedAuthenticationMethods() bool {
	if o != nil && o.ExcludedAuthenticationMethods != nil {
		return true
	}

	return false
}

// SetExcludedAuthenticationMethods gets a reference to the given []AuthenticationMethodObject and assigns it to the ExcludedAuthenticationMethods field.
func (o *PossessionConstraint) SetExcludedAuthenticationMethods(v []AuthenticationMethodObject) {
	o.ExcludedAuthenticationMethods = v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *PossessionConstraint) GetMethods() []string {
	if o == nil || o.Methods == nil {
		var ret []string
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetMethodsOk() ([]string, bool) {
	if o == nil || o.Methods == nil {
		return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *PossessionConstraint) HasMethods() bool {
	if o != nil && o.Methods != nil {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []string and assigns it to the Methods field.
func (o *PossessionConstraint) SetMethods(v []string) {
	o.Methods = v
}

// GetReauthenticateIn returns the ReauthenticateIn field value if set, zero value otherwise.
func (o *PossessionConstraint) GetReauthenticateIn() string {
	if o == nil || o.ReauthenticateIn == nil {
		var ret string
		return ret
	}
	return *o.ReauthenticateIn
}

// GetReauthenticateInOk returns a tuple with the ReauthenticateIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetReauthenticateInOk() (*string, bool) {
	if o == nil || o.ReauthenticateIn == nil {
		return nil, false
	}
	return o.ReauthenticateIn, true
}

// HasReauthenticateIn returns a boolean if a field has been set.
func (o *PossessionConstraint) HasReauthenticateIn() bool {
	if o != nil && o.ReauthenticateIn != nil {
		return true
	}

	return false
}

// SetReauthenticateIn gets a reference to the given string and assigns it to the ReauthenticateIn field.
func (o *PossessionConstraint) SetReauthenticateIn(v string) {
	o.ReauthenticateIn = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *PossessionConstraint) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *PossessionConstraint) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *PossessionConstraint) SetRequired(v bool) {
	o.Required = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *PossessionConstraint) GetTypes() []string {
	if o == nil || o.Types == nil {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetTypesOk() ([]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *PossessionConstraint) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *PossessionConstraint) SetTypes(v []string) {
	o.Types = v
}

// GetDeviceBound returns the DeviceBound field value if set, zero value otherwise.
func (o *PossessionConstraint) GetDeviceBound() string {
	if o == nil || o.DeviceBound == nil {
		var ret string
		return ret
	}
	return *o.DeviceBound
}

// GetDeviceBoundOk returns a tuple with the DeviceBound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetDeviceBoundOk() (*string, bool) {
	if o == nil || o.DeviceBound == nil {
		return nil, false
	}
	return o.DeviceBound, true
}

// HasDeviceBound returns a boolean if a field has been set.
func (o *PossessionConstraint) HasDeviceBound() bool {
	if o != nil && o.DeviceBound != nil {
		return true
	}

	return false
}

// SetDeviceBound gets a reference to the given string and assigns it to the DeviceBound field.
func (o *PossessionConstraint) SetDeviceBound(v string) {
	o.DeviceBound = &v
}

// GetHardwareProtection returns the HardwareProtection field value if set, zero value otherwise.
func (o *PossessionConstraint) GetHardwareProtection() string {
	if o == nil || o.HardwareProtection == nil {
		var ret string
		return ret
	}
	return *o.HardwareProtection
}

// GetHardwareProtectionOk returns a tuple with the HardwareProtection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetHardwareProtectionOk() (*string, bool) {
	if o == nil || o.HardwareProtection == nil {
		return nil, false
	}
	return o.HardwareProtection, true
}

// HasHardwareProtection returns a boolean if a field has been set.
func (o *PossessionConstraint) HasHardwareProtection() bool {
	if o != nil && o.HardwareProtection != nil {
		return true
	}

	return false
}

// SetHardwareProtection gets a reference to the given string and assigns it to the HardwareProtection field.
func (o *PossessionConstraint) SetHardwareProtection(v string) {
	o.HardwareProtection = &v
}

// GetPhishingResistant returns the PhishingResistant field value if set, zero value otherwise.
func (o *PossessionConstraint) GetPhishingResistant() string {
	if o == nil || o.PhishingResistant == nil {
		var ret string
		return ret
	}
	return *o.PhishingResistant
}

// GetPhishingResistantOk returns a tuple with the PhishingResistant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetPhishingResistantOk() (*string, bool) {
	if o == nil || o.PhishingResistant == nil {
		return nil, false
	}
	return o.PhishingResistant, true
}

// HasPhishingResistant returns a boolean if a field has been set.
func (o *PossessionConstraint) HasPhishingResistant() bool {
	if o != nil && o.PhishingResistant != nil {
		return true
	}

	return false
}

// SetPhishingResistant gets a reference to the given string and assigns it to the PhishingResistant field.
func (o *PossessionConstraint) SetPhishingResistant(v string) {
	o.PhishingResistant = &v
}

// GetUserPresence returns the UserPresence field value if set, zero value otherwise.
func (o *PossessionConstraint) GetUserPresence() string {
	if o == nil || o.UserPresence == nil {
		var ret string
		return ret
	}
	return *o.UserPresence
}

// GetUserPresenceOk returns a tuple with the UserPresence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetUserPresenceOk() (*string, bool) {
	if o == nil || o.UserPresence == nil {
		return nil, false
	}
	return o.UserPresence, true
}

// HasUserPresence returns a boolean if a field has been set.
func (o *PossessionConstraint) HasUserPresence() bool {
	if o != nil && o.UserPresence != nil {
		return true
	}

	return false
}

// SetUserPresence gets a reference to the given string and assigns it to the UserPresence field.
func (o *PossessionConstraint) SetUserPresence(v string) {
	o.UserPresence = &v
}

// GetUserVerification returns the UserVerification field value if set, zero value otherwise.
func (o *PossessionConstraint) GetUserVerification() string {
	if o == nil || o.UserVerification == nil {
		var ret string
		return ret
	}
	return *o.UserVerification
}

// GetUserVerificationOk returns a tuple with the UserVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PossessionConstraint) GetUserVerificationOk() (*string, bool) {
	if o == nil || o.UserVerification == nil {
		return nil, false
	}
	return o.UserVerification, true
}

// HasUserVerification returns a boolean if a field has been set.
func (o *PossessionConstraint) HasUserVerification() bool {
	if o != nil && o.UserVerification != nil {
		return true
	}

	return false
}

// SetUserVerification gets a reference to the given string and assigns it to the UserVerification field.
func (o *PossessionConstraint) SetUserVerification(v string) {
	o.UserVerification = &v
}

func (o PossessionConstraint) MarshalJSON() ([]byte, error) {
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
	if o.DeviceBound != nil {
		toSerialize["deviceBound"] = o.DeviceBound
	}
	if o.HardwareProtection != nil {
		toSerialize["hardwareProtection"] = o.HardwareProtection
	}
	if o.PhishingResistant != nil {
		toSerialize["phishingResistant"] = o.PhishingResistant
	}
	if o.UserPresence != nil {
		toSerialize["userPresence"] = o.UserPresence
	}
	if o.UserVerification != nil {
		toSerialize["userVerification"] = o.UserVerification
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PossessionConstraint) UnmarshalJSON(bytes []byte) (err error) {
	varPossessionConstraint := _PossessionConstraint{}

	err = json.Unmarshal(bytes, &varPossessionConstraint)
	if err == nil {
		*o = PossessionConstraint(varPossessionConstraint)
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
		delete(additionalProperties, "deviceBound")
		delete(additionalProperties, "hardwareProtection")
		delete(additionalProperties, "phishingResistant")
		delete(additionalProperties, "userPresence")
		delete(additionalProperties, "userVerification")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePossessionConstraint struct {
	value *PossessionConstraint
	isSet bool
}

func (v NullablePossessionConstraint) Get() *PossessionConstraint {
	return v.value
}

func (v *NullablePossessionConstraint) Set(val *PossessionConstraint) {
	v.value = val
	v.isSet = true
}

func (v NullablePossessionConstraint) IsSet() bool {
	return v.isSet
}

func (v *NullablePossessionConstraint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePossessionConstraint(val *PossessionConstraint) *NullablePossessionConstraint {
	return &NullablePossessionConstraint{value: val, isSet: true}
}

func (v NullablePossessionConstraint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePossessionConstraint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

