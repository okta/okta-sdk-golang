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

// OAuth2RefreshTokenScope struct for OAuth2RefreshTokenScope
type OAuth2RefreshTokenScope struct {
	// Description of the Scope
	Description *string `json:"description,omitempty"`
	// Name of the end user displayed in a consent dialog
	DisplayName *string `json:"displayName,omitempty"`
	// Scope object ID
	Id *string `json:"id,omitempty"`
	// Scope name
	Name *string `json:"name,omitempty"`
	Links *OAuth2RefreshTokenScopeLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenScope OAuth2RefreshTokenScope

// NewOAuth2RefreshTokenScope instantiates a new OAuth2RefreshTokenScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenScope() *OAuth2RefreshTokenScope {
	this := OAuth2RefreshTokenScope{}
	return &this
}

// NewOAuth2RefreshTokenScopeWithDefaults instantiates a new OAuth2RefreshTokenScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenScopeWithDefaults() *OAuth2RefreshTokenScope {
	this := OAuth2RefreshTokenScope{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenScope) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenScope) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenScope) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OAuth2RefreshTokenScope) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenScope) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenScope) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenScope) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OAuth2RefreshTokenScope) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenScope) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenScope) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenScope) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2RefreshTokenScope) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenScope) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenScope) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenScope) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OAuth2RefreshTokenScope) SetName(v string) {
	o.Name = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenScope) GetLinks() OAuth2RefreshTokenScopeLinks {
	if o == nil || o.Links == nil {
		var ret OAuth2RefreshTokenScopeLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenScope) GetLinksOk() (*OAuth2RefreshTokenScopeLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenScope) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OAuth2RefreshTokenScopeLinks and assigns it to the Links field.
func (o *OAuth2RefreshTokenScope) SetLinks(v OAuth2RefreshTokenScopeLinks) {
	o.Links = &v
}

func (o OAuth2RefreshTokenScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2RefreshTokenScope) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2RefreshTokenScope := _OAuth2RefreshTokenScope{}

	err = json.Unmarshal(bytes, &varOAuth2RefreshTokenScope)
	if err == nil {
		*o = OAuth2RefreshTokenScope(varOAuth2RefreshTokenScope)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2RefreshTokenScope struct {
	value *OAuth2RefreshTokenScope
	isSet bool
}

func (v NullableOAuth2RefreshTokenScope) Get() *OAuth2RefreshTokenScope {
	return v.value
}

func (v *NullableOAuth2RefreshTokenScope) Set(val *OAuth2RefreshTokenScope) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RefreshTokenScope) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RefreshTokenScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RefreshTokenScope(val *OAuth2RefreshTokenScope) *NullableOAuth2RefreshTokenScope {
	return &NullableOAuth2RefreshTokenScope{value: val, isSet: true}
}

func (v NullableOAuth2RefreshTokenScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RefreshTokenScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

