/*
Okta Admin Management API

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
	"fmt"
)

// checks if the ExpressConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpressConfiguration{}

// ExpressConfiguration Auth0 tenant details used to enable Express Configuration on an OIN Integration submission. Populates the submission's sign-on, authentication, and app settings so that admins can configure the app through the Express Configuration flow.
type ExpressConfiguration struct {
	// The client ID of the SaaS application that end users sign in to
	ApplicationClientId string `json:"applicationClientId"`
	// The Express Configuration capabilities to enable on the submission. If omitted, all capabilities are automatically configured based on the OIN integration's supported protocols.
	Capabilities []string `json:"capabilities,omitempty"`
	// The URL template that Okta uses to launch the app from the end-user dashboard. Supports template variables `{organization_name}`, `{organization_id}`, and `{connection_name}`.
	InitiateLoginUriTemplate *string `json:"initiateLoginUriTemplate,omitempty"`
	// The Auth0 admin login domain that Okta redirects to as part of the consent flow in a web browser. Use the custom domain name if one is configured in Auth0.
	LoginDomain string `json:"loginDomain"`
	// The client ID of the additional client application that Auth0 creates for the OIN administrator consent flow between Okta and Auth0
	OinClientId string `json:"oinClientId"`
	// The Auth0 tenant domain (for example, `example.auth0.com`)
	TenantDomain         string `json:"tenantDomain"`
	AdditionalProperties map[string]interface{}
}

type _ExpressConfiguration ExpressConfiguration

// NewExpressConfiguration instantiates a new ExpressConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpressConfiguration(applicationClientId string, loginDomain string, oinClientId string, tenantDomain string) *ExpressConfiguration {
	this := ExpressConfiguration{}
	this.ApplicationClientId = applicationClientId
	this.LoginDomain = loginDomain
	this.OinClientId = oinClientId
	this.TenantDomain = tenantDomain
	return &this
}

// NewExpressConfigurationWithDefaults instantiates a new ExpressConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpressConfigurationWithDefaults() *ExpressConfiguration {
	this := ExpressConfiguration{}
	return &this
}

// GetApplicationClientId returns the ApplicationClientId field value
func (o *ExpressConfiguration) GetApplicationClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationClientId
}

// GetApplicationClientIdOk returns a tuple with the ApplicationClientId field value
// and a boolean to check if the value has been set.
func (o *ExpressConfiguration) GetApplicationClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationClientId, true
}

// SetApplicationClientId sets field value
func (o *ExpressConfiguration) SetApplicationClientId(v string) {
	o.ApplicationClientId = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *ExpressConfiguration) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressConfiguration) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *ExpressConfiguration) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *ExpressConfiguration) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetInitiateLoginUriTemplate returns the InitiateLoginUriTemplate field value if set, zero value otherwise.
func (o *ExpressConfiguration) GetInitiateLoginUriTemplate() string {
	if o == nil || IsNil(o.InitiateLoginUriTemplate) {
		var ret string
		return ret
	}
	return *o.InitiateLoginUriTemplate
}

// GetInitiateLoginUriTemplateOk returns a tuple with the InitiateLoginUriTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpressConfiguration) GetInitiateLoginUriTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.InitiateLoginUriTemplate) {
		return nil, false
	}
	return o.InitiateLoginUriTemplate, true
}

// HasInitiateLoginUriTemplate returns a boolean if a field has been set.
func (o *ExpressConfiguration) HasInitiateLoginUriTemplate() bool {
	if o != nil && !IsNil(o.InitiateLoginUriTemplate) {
		return true
	}

	return false
}

// SetInitiateLoginUriTemplate gets a reference to the given string and assigns it to the InitiateLoginUriTemplate field.
func (o *ExpressConfiguration) SetInitiateLoginUriTemplate(v string) {
	o.InitiateLoginUriTemplate = &v
}

// GetLoginDomain returns the LoginDomain field value
func (o *ExpressConfiguration) GetLoginDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoginDomain
}

// GetLoginDomainOk returns a tuple with the LoginDomain field value
// and a boolean to check if the value has been set.
func (o *ExpressConfiguration) GetLoginDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoginDomain, true
}

// SetLoginDomain sets field value
func (o *ExpressConfiguration) SetLoginDomain(v string) {
	o.LoginDomain = v
}

// GetOinClientId returns the OinClientId field value
func (o *ExpressConfiguration) GetOinClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OinClientId
}

// GetOinClientIdOk returns a tuple with the OinClientId field value
// and a boolean to check if the value has been set.
func (o *ExpressConfiguration) GetOinClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OinClientId, true
}

// SetOinClientId sets field value
func (o *ExpressConfiguration) SetOinClientId(v string) {
	o.OinClientId = v
}

// GetTenantDomain returns the TenantDomain field value
func (o *ExpressConfiguration) GetTenantDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantDomain
}

// GetTenantDomainOk returns a tuple with the TenantDomain field value
// and a boolean to check if the value has been set.
func (o *ExpressConfiguration) GetTenantDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantDomain, true
}

// SetTenantDomain sets field value
func (o *ExpressConfiguration) SetTenantDomain(v string) {
	o.TenantDomain = v
}

func (o ExpressConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpressConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicationClientId"] = o.ApplicationClientId
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.InitiateLoginUriTemplate) {
		toSerialize["initiateLoginUriTemplate"] = o.InitiateLoginUriTemplate
	}
	toSerialize["loginDomain"] = o.LoginDomain
	toSerialize["oinClientId"] = o.OinClientId
	toSerialize["tenantDomain"] = o.TenantDomain

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExpressConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicationClientId",
		"loginDomain",
		"oinClientId",
		"tenantDomain",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varExpressConfiguration := _ExpressConfiguration{}

	err = json.Unmarshal(data, &varExpressConfiguration)

	if err != nil {
		return err
	}

	*o = ExpressConfiguration(varExpressConfiguration)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applicationClientId")
		delete(additionalProperties, "capabilities")
		delete(additionalProperties, "initiateLoginUriTemplate")
		delete(additionalProperties, "loginDomain")
		delete(additionalProperties, "oinClientId")
		delete(additionalProperties, "tenantDomain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpressConfiguration struct {
	value *ExpressConfiguration
	isSet bool
}

func (v NullableExpressConfiguration) Get() *ExpressConfiguration {
	return v.value
}

func (v *NullableExpressConfiguration) Set(val *ExpressConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressConfiguration(val *ExpressConfiguration) *NullableExpressConfiguration {
	return &NullableExpressConfiguration{value: val, isSet: true}
}

func (v NullableExpressConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
