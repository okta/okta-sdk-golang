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

// OAuth2ScopeConsentGrantLinks struct for OAuth2ScopeConsentGrantLinks
type OAuth2ScopeConsentGrantLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	App *AppCustomHrefObject `json:"app,omitempty"`
	Client *AppCustomHrefObject `json:"client,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2ScopeConsentGrantLinks OAuth2ScopeConsentGrantLinks

// NewOAuth2ScopeConsentGrantLinks instantiates a new OAuth2ScopeConsentGrantLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2ScopeConsentGrantLinks() *OAuth2ScopeConsentGrantLinks {
	this := OAuth2ScopeConsentGrantLinks{}
	return &this
}

// NewOAuth2ScopeConsentGrantLinksWithDefaults instantiates a new OAuth2ScopeConsentGrantLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ScopeConsentGrantLinksWithDefaults() *OAuth2ScopeConsentGrantLinks {
	this := OAuth2ScopeConsentGrantLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *OAuth2ScopeConsentGrantLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetApp() AppCustomHrefObject {
	if o == nil || o.App == nil {
		var ret AppCustomHrefObject
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetAppOk() (*AppCustomHrefObject, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppCustomHrefObject and assigns it to the App field.
func (o *OAuth2ScopeConsentGrantLinks) SetApp(v AppCustomHrefObject) {
	o.App = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetClient() AppCustomHrefObject {
	if o == nil || o.Client == nil {
		var ret AppCustomHrefObject
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetClientOk() (*AppCustomHrefObject, bool) {
	if o == nil || o.Client == nil {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasClient() bool {
	if o != nil && o.Client != nil {
		return true
	}

	return false
}

// SetClient gets a reference to the given AppCustomHrefObject and assigns it to the Client field.
func (o *OAuth2ScopeConsentGrantLinks) SetClient(v AppCustomHrefObject) {
	o.Client = &v
}

func (o OAuth2ScopeConsentGrantLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.Client != nil {
		toSerialize["client"] = o.Client
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OAuth2ScopeConsentGrantLinks) UnmarshalJSON(bytes []byte) (err error) {
	varOAuth2ScopeConsentGrantLinks := _OAuth2ScopeConsentGrantLinks{}

	err = json.Unmarshal(bytes, &varOAuth2ScopeConsentGrantLinks)
	if err == nil {
		*o = OAuth2ScopeConsentGrantLinks(varOAuth2ScopeConsentGrantLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "app")
		delete(additionalProperties, "client")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOAuth2ScopeConsentGrantLinks struct {
	value *OAuth2ScopeConsentGrantLinks
	isSet bool
}

func (v NullableOAuth2ScopeConsentGrantLinks) Get() *OAuth2ScopeConsentGrantLinks {
	return v.value
}

func (v *NullableOAuth2ScopeConsentGrantLinks) Set(val *OAuth2ScopeConsentGrantLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2ScopeConsentGrantLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2ScopeConsentGrantLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2ScopeConsentGrantLinks(val *OAuth2ScopeConsentGrantLinks) *NullableOAuth2ScopeConsentGrantLinks {
	return &NullableOAuth2ScopeConsentGrantLinks{value: val, isSet: true}
}

func (v NullableOAuth2ScopeConsentGrantLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2ScopeConsentGrantLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

