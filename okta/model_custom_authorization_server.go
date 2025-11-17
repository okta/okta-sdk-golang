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
	"fmt"
)

// checks if the CustomAuthorizationServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomAuthorizationServer{}

// CustomAuthorizationServer Custom authorization server for the managed connection
type CustomAuthorizationServer struct {
	// Issuer URL for the authorization server
	IssuerUrl string `json:"issuerUrl"`
	// Image URL for the authorization server
	Logo *string `json:"logo,omitempty"`
	// Display name of the authorization server
	Name string `json:"name"`
	// The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the authorization server
	Orn string `json:"orn"`
	// Audience value to use when requesting tokens
	ResourceIndicator    *string                        `json:"resourceIndicator,omitempty"`
	Links                CustomAuthorizationServerLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _CustomAuthorizationServer CustomAuthorizationServer

// NewCustomAuthorizationServer instantiates a new CustomAuthorizationServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAuthorizationServer(issuerUrl string, name string, orn string, links CustomAuthorizationServerLinks) *CustomAuthorizationServer {
	this := CustomAuthorizationServer{}
	this.IssuerUrl = issuerUrl
	this.Name = name
	this.Orn = orn
	this.Links = links
	return &this
}

// NewCustomAuthorizationServerWithDefaults instantiates a new CustomAuthorizationServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAuthorizationServerWithDefaults() *CustomAuthorizationServer {
	this := CustomAuthorizationServer{}
	return &this
}

// GetIssuerUrl returns the IssuerUrl field value
func (o *CustomAuthorizationServer) GetIssuerUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerUrl
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServer) GetIssuerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssuerUrl, true
}

// SetIssuerUrl sets field value
func (o *CustomAuthorizationServer) SetIssuerUrl(v string) {
	o.IssuerUrl = v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *CustomAuthorizationServer) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServer) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *CustomAuthorizationServer) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *CustomAuthorizationServer) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value
func (o *CustomAuthorizationServer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomAuthorizationServer) SetName(v string) {
	o.Name = v
}

// GetOrn returns the Orn field value
func (o *CustomAuthorizationServer) GetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Orn
}

// GetOrnOk returns a tuple with the Orn field value
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServer) GetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orn, true
}

// SetOrn sets field value
func (o *CustomAuthorizationServer) SetOrn(v string) {
	o.Orn = v
}

// GetResourceIndicator returns the ResourceIndicator field value if set, zero value otherwise.
func (o *CustomAuthorizationServer) GetResourceIndicator() string {
	if o == nil || IsNil(o.ResourceIndicator) {
		var ret string
		return ret
	}
	return *o.ResourceIndicator
}

// GetResourceIndicatorOk returns a tuple with the ResourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServer) GetResourceIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceIndicator) {
		return nil, false
	}
	return o.ResourceIndicator, true
}

// HasResourceIndicator returns a boolean if a field has been set.
func (o *CustomAuthorizationServer) HasResourceIndicator() bool {
	if o != nil && !IsNil(o.ResourceIndicator) {
		return true
	}

	return false
}

// SetResourceIndicator gets a reference to the given string and assigns it to the ResourceIndicator field.
func (o *CustomAuthorizationServer) SetResourceIndicator(v string) {
	o.ResourceIndicator = &v
}

// GetLinks returns the Links field value
func (o *CustomAuthorizationServer) GetLinks() CustomAuthorizationServerLinks {
	if o == nil {
		var ret CustomAuthorizationServerLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CustomAuthorizationServer) GetLinksOk() (*CustomAuthorizationServerLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CustomAuthorizationServer) SetLinks(v CustomAuthorizationServerLinks) {
	o.Links = v
}

func (o CustomAuthorizationServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomAuthorizationServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuerUrl"] = o.IssuerUrl
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	toSerialize["name"] = o.Name
	toSerialize["orn"] = o.Orn
	if !IsNil(o.ResourceIndicator) {
		toSerialize["resourceIndicator"] = o.ResourceIndicator
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomAuthorizationServer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"issuerUrl",
		"name",
		"orn",
		"_links",
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

	varCustomAuthorizationServer := _CustomAuthorizationServer{}

	err = json.Unmarshal(data, &varCustomAuthorizationServer)

	if err != nil {
		return err
	}

	*o = CustomAuthorizationServer(varCustomAuthorizationServer)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "issuerUrl")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "orn")
		delete(additionalProperties, "resourceIndicator")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomAuthorizationServer struct {
	value *CustomAuthorizationServer
	isSet bool
}

func (v NullableCustomAuthorizationServer) Get() *CustomAuthorizationServer {
	return v.value
}

func (v *NullableCustomAuthorizationServer) Set(val *CustomAuthorizationServer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAuthorizationServer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAuthorizationServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAuthorizationServer(val *CustomAuthorizationServer) *NullableCustomAuthorizationServer {
	return &NullableCustomAuthorizationServer{value: val, isSet: true}
}

func (v NullableCustomAuthorizationServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAuthorizationServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
