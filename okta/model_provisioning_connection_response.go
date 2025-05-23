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

// ProvisioningConnectionResponse struct for ProvisioningConnectionResponse
type ProvisioningConnectionResponse struct {
	// A token is used to authenticate with the app. This property is only returned for the `TOKEN` authentication scheme.
	AuthScheme *string `json:"authScheme,omitempty"`
	// Base URL
	BaseUrl *string `json:"baseUrl,omitempty"`
	Profile ProvisioningConnectionResponseProfile `json:"profile"`
	// Provisioning Connection status
	Status string `json:"status"`
	Links *LinksSelfLifecycleAndAuthorize `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProvisioningConnectionResponse ProvisioningConnectionResponse

// NewProvisioningConnectionResponse instantiates a new ProvisioningConnectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningConnectionResponse(profile ProvisioningConnectionResponseProfile, status string) *ProvisioningConnectionResponse {
	this := ProvisioningConnectionResponse{}
	this.Profile = profile
	this.Status = status
	return &this
}

// NewProvisioningConnectionResponseWithDefaults instantiates a new ProvisioningConnectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningConnectionResponseWithDefaults() *ProvisioningConnectionResponse {
	this := ProvisioningConnectionResponse{}
	var status string = "DISABLED"
	this.Status = status
	return &this
}

// GetAuthScheme returns the AuthScheme field value if set, zero value otherwise.
func (o *ProvisioningConnectionResponse) GetAuthScheme() string {
	if o == nil || o.AuthScheme == nil {
		var ret string
		return ret
	}
	return *o.AuthScheme
}

// GetAuthSchemeOk returns a tuple with the AuthScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponse) GetAuthSchemeOk() (*string, bool) {
	if o == nil || o.AuthScheme == nil {
		return nil, false
	}
	return o.AuthScheme, true
}

// HasAuthScheme returns a boolean if a field has been set.
func (o *ProvisioningConnectionResponse) HasAuthScheme() bool {
	if o != nil && o.AuthScheme != nil {
		return true
	}

	return false
}

// SetAuthScheme gets a reference to the given string and assigns it to the AuthScheme field.
func (o *ProvisioningConnectionResponse) SetAuthScheme(v string) {
	o.AuthScheme = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *ProvisioningConnectionResponse) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *ProvisioningConnectionResponse) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *ProvisioningConnectionResponse) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetProfile returns the Profile field value
func (o *ProvisioningConnectionResponse) GetProfile() ProvisioningConnectionResponseProfile {
	if o == nil {
		var ret ProvisioningConnectionResponseProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponse) GetProfileOk() (*ProvisioningConnectionResponseProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *ProvisioningConnectionResponse) SetProfile(v ProvisioningConnectionResponseProfile) {
	o.Profile = v
}

// GetStatus returns the Status field value
func (o *ProvisioningConnectionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ProvisioningConnectionResponse) SetStatus(v string) {
	o.Status = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProvisioningConnectionResponse) GetLinks() LinksSelfLifecycleAndAuthorize {
	if o == nil || o.Links == nil {
		var ret LinksSelfLifecycleAndAuthorize
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningConnectionResponse) GetLinksOk() (*LinksSelfLifecycleAndAuthorize, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProvisioningConnectionResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfLifecycleAndAuthorize and assigns it to the Links field.
func (o *ProvisioningConnectionResponse) SetLinks(v LinksSelfLifecycleAndAuthorize) {
	o.Links = &v
}

func (o ProvisioningConnectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthScheme != nil {
		toSerialize["authScheme"] = o.AuthScheme
	}
	if o.BaseUrl != nil {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if true {
		toSerialize["profile"] = o.Profile
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProvisioningConnectionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProvisioningConnectionResponse := _ProvisioningConnectionResponse{}

	err = json.Unmarshal(bytes, &varProvisioningConnectionResponse)
	if err == nil {
		*o = ProvisioningConnectionResponse(varProvisioningConnectionResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "authScheme")
		delete(additionalProperties, "baseUrl")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableProvisioningConnectionResponse struct {
	value *ProvisioningConnectionResponse
	isSet bool
}

func (v NullableProvisioningConnectionResponse) Get() *ProvisioningConnectionResponse {
	return v.value
}

func (v *NullableProvisioningConnectionResponse) Set(val *ProvisioningConnectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningConnectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningConnectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningConnectionResponse(val *ProvisioningConnectionResponse) *NullableProvisioningConnectionResponse {
	return &NullableProvisioningConnectionResponse{value: val, isSet: true}
}

func (v NullableProvisioningConnectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningConnectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

