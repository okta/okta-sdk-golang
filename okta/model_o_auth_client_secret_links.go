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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OAuthClientSecretLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthClientSecretLinks{}

// OAuthClientSecretLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type OAuthClientSecretLinks struct {
	Activate             *HrefObjectActivateLink   `json:"activate,omitempty"`
	Deactivate           *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Delete               *HrefObjectDeleteLink     `json:"delete,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuthClientSecretLinks OAuthClientSecretLinks

// NewOAuthClientSecretLinks instantiates a new OAuthClientSecretLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthClientSecretLinks() *OAuthClientSecretLinks {
	this := OAuthClientSecretLinks{}
	return &this
}

// NewOAuthClientSecretLinksWithDefaults instantiates a new OAuthClientSecretLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthClientSecretLinksWithDefaults() *OAuthClientSecretLinks {
	this := OAuthClientSecretLinks{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *OAuthClientSecretLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || IsNil(o.Activate) {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientSecretLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *OAuthClientSecretLinks) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *OAuthClientSecretLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *OAuthClientSecretLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientSecretLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *OAuthClientSecretLinks) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *OAuthClientSecretLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *OAuthClientSecretLinks) GetDelete() HrefObjectDeleteLink {
	if o == nil || IsNil(o.Delete) {
		var ret HrefObjectDeleteLink
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthClientSecretLinks) GetDeleteOk() (*HrefObjectDeleteLink, bool) {
	if o == nil || IsNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *OAuthClientSecretLinks) HasDelete() bool {
	if o != nil && !IsNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given HrefObjectDeleteLink and assigns it to the Delete field.
func (o *OAuthClientSecretLinks) SetDelete(v HrefObjectDeleteLink) {
	o.Delete = &v
}

func (o OAuthClientSecretLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthClientSecretLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuthClientSecretLinks) UnmarshalJSON(data []byte) (err error) {
	varOAuthClientSecretLinks := _OAuthClientSecretLinks{}

	err = json.Unmarshal(data, &varOAuthClientSecretLinks)

	if err != nil {
		return err
	}

	*o = OAuthClientSecretLinks(varOAuthClientSecretLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "delete")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuthClientSecretLinks struct {
	value *OAuthClientSecretLinks
	isSet bool
}

func (v NullableOAuthClientSecretLinks) Get() *OAuthClientSecretLinks {
	return v.value
}

func (v *NullableOAuthClientSecretLinks) Set(val *OAuthClientSecretLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthClientSecretLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthClientSecretLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthClientSecretLinks(val *OAuthClientSecretLinks) *NullableOAuthClientSecretLinks {
	return &NullableOAuthClientSecretLinks{value: val, isSet: true}
}

func (v NullableOAuthClientSecretLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthClientSecretLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
