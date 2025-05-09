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

// LogAuthenticationContext struct for LogAuthenticationContext
type LogAuthenticationContext struct {
	AuthenticationProvider *string `json:"authenticationProvider,omitempty"`
	AuthenticationStep *int32 `json:"authenticationStep,omitempty"`
	CredentialProvider *string `json:"credentialProvider,omitempty"`
	CredentialType *string `json:"credentialType,omitempty"`
	ExternalSessionId *string `json:"externalSessionId,omitempty"`
	Interface *string `json:"interface,omitempty"`
	Issuer *LogIssuer `json:"issuer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogAuthenticationContext LogAuthenticationContext

// NewLogAuthenticationContext instantiates a new LogAuthenticationContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogAuthenticationContext() *LogAuthenticationContext {
	this := LogAuthenticationContext{}
	return &this
}

// NewLogAuthenticationContextWithDefaults instantiates a new LogAuthenticationContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogAuthenticationContextWithDefaults() *LogAuthenticationContext {
	this := LogAuthenticationContext{}
	return &this
}

// GetAuthenticationProvider returns the AuthenticationProvider field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetAuthenticationProvider() string {
	if o == nil || o.AuthenticationProvider == nil {
		var ret string
		return ret
	}
	return *o.AuthenticationProvider
}

// GetAuthenticationProviderOk returns a tuple with the AuthenticationProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetAuthenticationProviderOk() (*string, bool) {
	if o == nil || o.AuthenticationProvider == nil {
		return nil, false
	}
	return o.AuthenticationProvider, true
}

// HasAuthenticationProvider returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasAuthenticationProvider() bool {
	if o != nil && o.AuthenticationProvider != nil {
		return true
	}

	return false
}

// SetAuthenticationProvider gets a reference to the given string and assigns it to the AuthenticationProvider field.
func (o *LogAuthenticationContext) SetAuthenticationProvider(v string) {
	o.AuthenticationProvider = &v
}

// GetAuthenticationStep returns the AuthenticationStep field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetAuthenticationStep() int32 {
	if o == nil || o.AuthenticationStep == nil {
		var ret int32
		return ret
	}
	return *o.AuthenticationStep
}

// GetAuthenticationStepOk returns a tuple with the AuthenticationStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetAuthenticationStepOk() (*int32, bool) {
	if o == nil || o.AuthenticationStep == nil {
		return nil, false
	}
	return o.AuthenticationStep, true
}

// HasAuthenticationStep returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasAuthenticationStep() bool {
	if o != nil && o.AuthenticationStep != nil {
		return true
	}

	return false
}

// SetAuthenticationStep gets a reference to the given int32 and assigns it to the AuthenticationStep field.
func (o *LogAuthenticationContext) SetAuthenticationStep(v int32) {
	o.AuthenticationStep = &v
}

// GetCredentialProvider returns the CredentialProvider field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetCredentialProvider() string {
	if o == nil || o.CredentialProvider == nil {
		var ret string
		return ret
	}
	return *o.CredentialProvider
}

// GetCredentialProviderOk returns a tuple with the CredentialProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetCredentialProviderOk() (*string, bool) {
	if o == nil || o.CredentialProvider == nil {
		return nil, false
	}
	return o.CredentialProvider, true
}

// HasCredentialProvider returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasCredentialProvider() bool {
	if o != nil && o.CredentialProvider != nil {
		return true
	}

	return false
}

// SetCredentialProvider gets a reference to the given string and assigns it to the CredentialProvider field.
func (o *LogAuthenticationContext) SetCredentialProvider(v string) {
	o.CredentialProvider = &v
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetCredentialType() string {
	if o == nil || o.CredentialType == nil {
		var ret string
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetCredentialTypeOk() (*string, bool) {
	if o == nil || o.CredentialType == nil {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasCredentialType() bool {
	if o != nil && o.CredentialType != nil {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given string and assigns it to the CredentialType field.
func (o *LogAuthenticationContext) SetCredentialType(v string) {
	o.CredentialType = &v
}

// GetExternalSessionId returns the ExternalSessionId field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetExternalSessionId() string {
	if o == nil || o.ExternalSessionId == nil {
		var ret string
		return ret
	}
	return *o.ExternalSessionId
}

// GetExternalSessionIdOk returns a tuple with the ExternalSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetExternalSessionIdOk() (*string, bool) {
	if o == nil || o.ExternalSessionId == nil {
		return nil, false
	}
	return o.ExternalSessionId, true
}

// HasExternalSessionId returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasExternalSessionId() bool {
	if o != nil && o.ExternalSessionId != nil {
		return true
	}

	return false
}

// SetExternalSessionId gets a reference to the given string and assigns it to the ExternalSessionId field.
func (o *LogAuthenticationContext) SetExternalSessionId(v string) {
	o.ExternalSessionId = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetInterface() string {
	if o == nil || o.Interface == nil {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetInterfaceOk() (*string, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *LogAuthenticationContext) SetInterface(v string) {
	o.Interface = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *LogAuthenticationContext) GetIssuer() LogIssuer {
	if o == nil || o.Issuer == nil {
		var ret LogIssuer
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogAuthenticationContext) GetIssuerOk() (*LogIssuer, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *LogAuthenticationContext) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given LogIssuer and assigns it to the Issuer field.
func (o *LogAuthenticationContext) SetIssuer(v LogIssuer) {
	o.Issuer = &v
}

func (o LogAuthenticationContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthenticationProvider != nil {
		toSerialize["authenticationProvider"] = o.AuthenticationProvider
	}
	if o.AuthenticationStep != nil {
		toSerialize["authenticationStep"] = o.AuthenticationStep
	}
	if o.CredentialProvider != nil {
		toSerialize["credentialProvider"] = o.CredentialProvider
	}
	if o.CredentialType != nil {
		toSerialize["credentialType"] = o.CredentialType
	}
	if o.ExternalSessionId != nil {
		toSerialize["externalSessionId"] = o.ExternalSessionId
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogAuthenticationContext) UnmarshalJSON(bytes []byte) (err error) {
	varLogAuthenticationContext := _LogAuthenticationContext{}

	err = json.Unmarshal(bytes, &varLogAuthenticationContext)
	if err == nil {
		*o = LogAuthenticationContext(varLogAuthenticationContext)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authenticationProvider")
		delete(additionalProperties, "authenticationStep")
		delete(additionalProperties, "credentialProvider")
		delete(additionalProperties, "credentialType")
		delete(additionalProperties, "externalSessionId")
		delete(additionalProperties, "interface")
		delete(additionalProperties, "issuer")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogAuthenticationContext struct {
	value *LogAuthenticationContext
	isSet bool
}

func (v NullableLogAuthenticationContext) Get() *LogAuthenticationContext {
	return v.value
}

func (v *NullableLogAuthenticationContext) Set(val *LogAuthenticationContext) {
	v.value = val
	v.isSet = true
}

func (v NullableLogAuthenticationContext) IsSet() bool {
	return v.isSet
}

func (v *NullableLogAuthenticationContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogAuthenticationContext(val *LogAuthenticationContext) *NullableLogAuthenticationContext {
	return &NullableLogAuthenticationContext{value: val, isSet: true}
}

func (v NullableLogAuthenticationContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogAuthenticationContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

