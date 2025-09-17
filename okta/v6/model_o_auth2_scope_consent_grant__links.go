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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the OAuth2ScopeConsentGrantLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2ScopeConsentGrantLinks{}

// OAuth2ScopeConsentGrantLinks struct for OAuth2ScopeConsentGrantLinks
type OAuth2ScopeConsentGrantLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the app resource
	App *AppResourceHrefObject `json:"app,omitempty"`
	// Link to the client resource
	Client *AppResourceHrefObject `json:"client,omitempty"`
	// Link to the scope resource
	Scope *ScopeResourceHrefObject `json:"scope,omitempty"`
	// Link to the user resource
	User *UserResourceHrefObject `json:"user,omitempty"`
	// Link to the authorization server resource
	AuthorizationServer  *AuthorizationServerResourceHrefObject `json:"authorizationServer,omitempty"`
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
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *OAuth2ScopeConsentGrantLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetApp() AppResourceHrefObject {
	if o == nil || IsNil(o.App) {
		var ret AppResourceHrefObject
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetAppOk() (*AppResourceHrefObject, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppResourceHrefObject and assigns it to the App field.
func (o *OAuth2ScopeConsentGrantLinks) SetApp(v AppResourceHrefObject) {
	o.App = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetClient() AppResourceHrefObject {
	if o == nil || IsNil(o.Client) {
		var ret AppResourceHrefObject
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetClientOk() (*AppResourceHrefObject, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given AppResourceHrefObject and assigns it to the Client field.
func (o *OAuth2ScopeConsentGrantLinks) SetClient(v AppResourceHrefObject) {
	o.Client = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetScope() ScopeResourceHrefObject {
	if o == nil || IsNil(o.Scope) {
		var ret ScopeResourceHrefObject
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetScopeOk() (*ScopeResourceHrefObject, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given ScopeResourceHrefObject and assigns it to the Scope field.
func (o *OAuth2ScopeConsentGrantLinks) SetScope(v ScopeResourceHrefObject) {
	o.Scope = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetUser() UserResourceHrefObject {
	if o == nil || IsNil(o.User) {
		var ret UserResourceHrefObject
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetUserOk() (*UserResourceHrefObject, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserResourceHrefObject and assigns it to the User field.
func (o *OAuth2ScopeConsentGrantLinks) SetUser(v UserResourceHrefObject) {
	o.User = &v
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *OAuth2ScopeConsentGrantLinks) GetAuthorizationServer() AuthorizationServerResourceHrefObject {
	if o == nil || IsNil(o.AuthorizationServer) {
		var ret AuthorizationServerResourceHrefObject
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2ScopeConsentGrantLinks) GetAuthorizationServerOk() (*AuthorizationServerResourceHrefObject, bool) {
	if o == nil || IsNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *OAuth2ScopeConsentGrantLinks) HasAuthorizationServer() bool {
	if o != nil && !IsNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given AuthorizationServerResourceHrefObject and assigns it to the AuthorizationServer field.
func (o *OAuth2ScopeConsentGrantLinks) SetAuthorizationServer(v AuthorizationServerResourceHrefObject) {
	o.AuthorizationServer = &v
}

func (o OAuth2ScopeConsentGrantLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2ScopeConsentGrantLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.AuthorizationServer) {
		toSerialize["authorizationServer"] = o.AuthorizationServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OAuth2ScopeConsentGrantLinks) UnmarshalJSON(data []byte) (err error) {
	varOAuth2ScopeConsentGrantLinks := _OAuth2ScopeConsentGrantLinks{}

	err = json.Unmarshal(data, &varOAuth2ScopeConsentGrantLinks)

	if err != nil {
		return err
	}

	*o = OAuth2ScopeConsentGrantLinks(varOAuth2ScopeConsentGrantLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "app")
		delete(additionalProperties, "client")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "user")
		delete(additionalProperties, "authorizationServer")
		o.AdditionalProperties = additionalProperties
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
