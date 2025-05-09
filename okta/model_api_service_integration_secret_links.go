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

// APIServiceIntegrationSecretLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type APIServiceIntegrationSecretLinks struct {
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Delete *HrefObjectDeleteLink `json:"delete,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APIServiceIntegrationSecretLinks APIServiceIntegrationSecretLinks

// NewAPIServiceIntegrationSecretLinks instantiates a new APIServiceIntegrationSecretLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServiceIntegrationSecretLinks() *APIServiceIntegrationSecretLinks {
	this := APIServiceIntegrationSecretLinks{}
	return &this
}

// NewAPIServiceIntegrationSecretLinksWithDefaults instantiates a new APIServiceIntegrationSecretLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServiceIntegrationSecretLinksWithDefaults() *APIServiceIntegrationSecretLinks {
	this := APIServiceIntegrationSecretLinks{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *APIServiceIntegrationSecretLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationSecretLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *APIServiceIntegrationSecretLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *APIServiceIntegrationSecretLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *APIServiceIntegrationSecretLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationSecretLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *APIServiceIntegrationSecretLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *APIServiceIntegrationSecretLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *APIServiceIntegrationSecretLinks) GetDelete() HrefObjectDeleteLink {
	if o == nil || o.Delete == nil {
		var ret HrefObjectDeleteLink
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationSecretLinks) GetDeleteOk() (*HrefObjectDeleteLink, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *APIServiceIntegrationSecretLinks) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given HrefObjectDeleteLink and assigns it to the Delete field.
func (o *APIServiceIntegrationSecretLinks) SetDelete(v HrefObjectDeleteLink) {
	o.Delete = &v
}

func (o APIServiceIntegrationSecretLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APIServiceIntegrationSecretLinks) UnmarshalJSON(bytes []byte) (err error) {
	varAPIServiceIntegrationSecretLinks := _APIServiceIntegrationSecretLinks{}

	err = json.Unmarshal(bytes, &varAPIServiceIntegrationSecretLinks)
	if err == nil {
		*o = APIServiceIntegrationSecretLinks(varAPIServiceIntegrationSecretLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "delete")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAPIServiceIntegrationSecretLinks struct {
	value *APIServiceIntegrationSecretLinks
	isSet bool
}

func (v NullableAPIServiceIntegrationSecretLinks) Get() *APIServiceIntegrationSecretLinks {
	return v.value
}

func (v *NullableAPIServiceIntegrationSecretLinks) Set(val *APIServiceIntegrationSecretLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServiceIntegrationSecretLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServiceIntegrationSecretLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServiceIntegrationSecretLinks(val *APIServiceIntegrationSecretLinks) *NullableAPIServiceIntegrationSecretLinks {
	return &NullableAPIServiceIntegrationSecretLinks{value: val, isSet: true}
}

func (v NullableAPIServiceIntegrationSecretLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServiceIntegrationSecretLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

