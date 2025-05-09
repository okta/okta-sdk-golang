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

// OAuth2Scope struct for OAuth2Scope
type OAuth2Scope struct {
	// Indicates whether a consent dialog is needed for the Scope
	Consent *string `json:"consent,omitempty"`
	// Indicates if this Scope is a default scope
	Default *bool `json:"default,omitempty"`
	// Description of the Scope
	Description *string `json:"description,omitempty"`
	// Name of the end user displayed in a consent dialog
	DisplayName *string `json:"displayName,omitempty"`
	// Scope object ID
	Id *string `json:"id,omitempty"`
	// Indicates whether the Scope is included in the metadata
	MetadataPublish *string `json:"metadataPublish,omitempty"`
	// Scope name
	Name *string `json:"name,omitempty"`
	Optional *bool `json:"optional,omitempty"`
	// Indicates if Okta created the Scope
	System *bool `json:"system,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2Scope OAuth2Scope

// NewOAuth2Scope instantiates a new OAuth2Scope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Scope() *OAuth2Scope {
	this := OAuth2Scope{}
	var consent string = "IMPLICIT"
	this.Consent = &consent
	var metadataPublish string = "NO_CLIENTS"
	this.MetadataPublish = &metadataPublish
	return &this
}

// NewOAuth2ScopeWithDefaults instantiates a new OAuth2Scope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ScopeWithDefaults() *OAuth2Scope {
	this := OAuth2Scope{}
	var consent string = "IMPLICIT"
	this.Consent = &consent
	var metadataPublish string = "NO_CLIENTS"
	this.MetadataPublish = &metadataPublish
	return &this
}

// GetConsent returns the Consent field value if set, zero value otherwise.
func (o *OAuth2Scope) GetConsent() string {
	if o == nil || o.Consent == nil {
		var ret string
		return ret
	}
	return *o.Consent
}

// GetConsentOk returns a tuple with the Consent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetConsentOk() (*string, bool) {
	if o == nil || o.Consent == nil {
		return nil, false
	}
	return o.Consent, true
}

// HasConsent returns a boolean if a field has been set.
func (o *OAuth2Scope) HasConsent() bool {
	if o != nil && o.Consent != nil {
		return true
	}

	return false
}

// SetConsent gets a reference to the given string and assigns it to the Consent field.
func (o *OAuth2Scope) SetConsent(v string) {
	o.Consent = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *OAuth2Scope) GetDefault() bool {
	if o == nil || o.Default == nil {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetDefaultOk() (*bool, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *OAuth2Scope) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *OAuth2Scope) SetDefault(v bool) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OAuth2Scope) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OAuth2Scope) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OAuth2Scope) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *OAuth2Scope) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *OAuth2Scope) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *OAuth2Scope) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OAuth2Scope) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OAuth2Scope) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OAuth2Scope) SetId(v string) {
	o.Id = &v
}

// GetMetadataPublish returns the MetadataPublish field value if set, zero value otherwise.
func (o *OAuth2Scope) GetMetadataPublish() string {
	if o == nil || o.MetadataPublish == nil {
		var ret string
		return ret
	}
	return *o.MetadataPublish
}

// GetMetadataPublishOk returns a tuple with the MetadataPublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetMetadataPublishOk() (*string, bool) {
	if o == nil || o.MetadataPublish == nil {
		return nil, false
	}
	return o.MetadataPublish, true
}

// HasMetadataPublish returns a boolean if a field has been set.
func (o *OAuth2Scope) HasMetadataPublish() bool {
	if o != nil && o.MetadataPublish != nil {
		return true
	}

	return false
}

// SetMetadataPublish gets a reference to the given string and assigns it to the MetadataPublish field.
func (o *OAuth2Scope) SetMetadataPublish(v string) {
	o.MetadataPublish = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OAuth2Scope) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OAuth2Scope) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OAuth2Scope) SetName(v string) {
	o.Name = &v
}

// GetOptional returns the Optional field value if set, zero value otherwise.
func (o *OAuth2Scope) GetOptional() bool {
	if o == nil || o.Optional == nil {
		var ret bool
		return ret
	}
	return *o.Optional
}

// GetOptionalOk returns a tuple with the Optional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetOptionalOk() (*bool, bool) {
	if o == nil || o.Optional == nil {
		return nil, false
	}
	return o.Optional, true
}

// HasOptional returns a boolean if a field has been set.
func (o *OAuth2Scope) HasOptional() bool {
	if o != nil && o.Optional != nil {
		return true
	}

	return false
}

// SetOptional gets a reference to the given bool and assigns it to the Optional field.
func (o *OAuth2Scope) SetOptional(v bool) {
	o.Optional = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *OAuth2Scope) GetSystem() bool {
	if o == nil || o.System == nil {
		var ret bool
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Scope) GetSystemOk() (*bool, bool) {
	if o == nil || o.System == nil {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *OAuth2Scope) HasSystem() bool {
	if o != nil && o.System != nil {
		return true
	}

	return false
}

// SetSystem gets a reference to the given bool and assigns it to the System field.
func (o *OAuth2Scope) SetSystem(v bool) {
	o.System = &v
}

func (o OAuth2Scope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Consent != nil {
		toSerialize["consent"] = o.Consent
	}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MetadataPublish != nil {
		toSerialize["metadataPublish"] = o.MetadataPublish
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Optional != nil {
		toSerialize["optional"] = o.Optional
	}
	if o.System != nil {
		toSerialize["system"] = o.System
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2Scope) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2Scope := _OAuth2Scope{}

	err = json.Unmarshal(bytes, &varOAuth2Scope)
	if err == nil {
		*o = OAuth2Scope(varOAuth2Scope)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "consent")
		delete(additionalProperties, "default")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metadataPublish")
		delete(additionalProperties, "name")
		delete(additionalProperties, "optional")
		delete(additionalProperties, "system")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2Scope struct {
	value *OAuth2Scope
	isSet bool
}

func (v NullableOAuth2Scope) Get() *OAuth2Scope {
	return v.value
}

func (v *NullableOAuth2Scope) Set(val *OAuth2Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Scope) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Scope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Scope(val *OAuth2Scope) *NullableOAuth2Scope {
	return &NullableOAuth2Scope{value: val, isSet: true}
}

func (v NullableOAuth2Scope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Scope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

