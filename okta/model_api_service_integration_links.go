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

// APIServiceIntegrationLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of an application using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type APIServiceIntegrationLinks struct {
	Client *HrefObjectClientLink `json:"client,omitempty"`
	Logo *HrefObjectLogoLink `json:"logo,omitempty"`
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APIServiceIntegrationLinks APIServiceIntegrationLinks

// NewAPIServiceIntegrationLinks instantiates a new APIServiceIntegrationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServiceIntegrationLinks() *APIServiceIntegrationLinks {
	this := APIServiceIntegrationLinks{}
	return &this
}

// NewAPIServiceIntegrationLinksWithDefaults instantiates a new APIServiceIntegrationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServiceIntegrationLinksWithDefaults() *APIServiceIntegrationLinks {
	this := APIServiceIntegrationLinks{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *APIServiceIntegrationLinks) GetClient() HrefObjectClientLink {
	if o == nil || o.Client == nil {
		var ret HrefObjectClientLink
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationLinks) GetClientOk() (*HrefObjectClientLink, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *APIServiceIntegrationLinks) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given HrefObjectClientLink and assigns it to the Client field.
func (o *APIServiceIntegrationLinks) SetClient(v HrefObjectClientLink) {
	o.Client = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *APIServiceIntegrationLinks) GetLogo() HrefObjectLogoLink {
	if o == nil || o.Logo == nil {
		var ret HrefObjectLogoLink
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationLinks) GetLogoOk() (*HrefObjectLogoLink, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *APIServiceIntegrationLinks) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given HrefObjectLogoLink and assigns it to the Logo field.
func (o *APIServiceIntegrationLinks) SetLogo(v HrefObjectLogoLink) {
	o.Logo = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *APIServiceIntegrationLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *APIServiceIntegrationLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *APIServiceIntegrationLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o APIServiceIntegrationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APIServiceIntegrationLinks) UnmarshalJSON(bytes []byte) (err error) {
	varAPIServiceIntegrationLinks := _APIServiceIntegrationLinks{}

	err = json.Unmarshal(bytes, &varAPIServiceIntegrationLinks)
	if err == nil {
		*o = APIServiceIntegrationLinks(varAPIServiceIntegrationLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "client")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAPIServiceIntegrationLinks struct {
	value *APIServiceIntegrationLinks
	isSet bool
}

func (v NullableAPIServiceIntegrationLinks) Get() *APIServiceIntegrationLinks {
	return v.value
}

func (v *NullableAPIServiceIntegrationLinks) Set(val *APIServiceIntegrationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServiceIntegrationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServiceIntegrationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServiceIntegrationLinks(val *APIServiceIntegrationLinks) *NullableAPIServiceIntegrationLinks {
	return &NullableAPIServiceIntegrationLinks{value: val, isSet: true}
}

func (v NullableAPIServiceIntegrationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServiceIntegrationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

