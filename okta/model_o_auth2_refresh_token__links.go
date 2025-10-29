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
)

// checks if the OAuth2RefreshTokenLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2RefreshTokenLinks{}

// OAuth2RefreshTokenLinks struct for OAuth2RefreshTokenLinks
type OAuth2RefreshTokenLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the app resource
	App    *AppResourceHrefObject              `json:"app,omitempty"`
	Revoke *OAuth2RefreshTokenLinksAllOfRevoke `json:"revoke,omitempty"`
	// Link to the client resource
	Client *AppResourceHrefObject `json:"client,omitempty"`
	// Link to the user resource
	User *UserResourceHrefObject `json:"user,omitempty"`
	// Link to the Token authorization server resource
	AuthorizationServer  *AuthorizationServerResourceHrefObject `json:"authorizationServer,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OAuth2RefreshTokenLinks OAuth2RefreshTokenLinks

// NewOAuth2RefreshTokenLinks instantiates a new OAuth2RefreshTokenLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2RefreshTokenLinks() *OAuth2RefreshTokenLinks {
	this := OAuth2RefreshTokenLinks{}
	return &this
}

// NewOAuth2RefreshTokenLinksWithDefaults instantiates a new OAuth2RefreshTokenLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2RefreshTokenLinksWithDefaults() *OAuth2RefreshTokenLinks {
	this := OAuth2RefreshTokenLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *OAuth2RefreshTokenLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinks) GetApp() AppResourceHrefObject {
	if o == nil || IsNil(o.App) {
		var ret AppResourceHrefObject
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinks) GetAppOk() (*AppResourceHrefObject, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinks) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given AppResourceHrefObject and assigns it to the App field.
func (o *OAuth2RefreshTokenLinks) SetApp(v AppResourceHrefObject) {
	o.App = &v
}

// GetRevoke returns the Revoke field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinks) GetRevoke() OAuth2RefreshTokenLinksAllOfRevoke {
	if o == nil || IsNil(o.Revoke) {
		var ret OAuth2RefreshTokenLinksAllOfRevoke
		return ret
	}
	return *o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinks) GetRevokeOk() (*OAuth2RefreshTokenLinksAllOfRevoke, bool) {
	if o == nil || IsNil(o.Revoke) {
		return nil, false
	}
	return o.Revoke, true
}

// HasRevoke returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinks) HasRevoke() bool {
	if o != nil && !IsNil(o.Revoke) {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given OAuth2RefreshTokenLinksAllOfRevoke and assigns it to the Revoke field.
func (o *OAuth2RefreshTokenLinks) SetRevoke(v OAuth2RefreshTokenLinksAllOfRevoke) {
	o.Revoke = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinks) GetClient() AppResourceHrefObject {
	if o == nil || IsNil(o.Client) {
		var ret AppResourceHrefObject
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinks) GetClientOk() (*AppResourceHrefObject, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinks) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given AppResourceHrefObject and assigns it to the Client field.
func (o *OAuth2RefreshTokenLinks) SetClient(v AppResourceHrefObject) {
	o.Client = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinks) GetUser() UserResourceHrefObject {
	if o == nil || IsNil(o.User) {
		var ret UserResourceHrefObject
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinks) GetUserOk() (*UserResourceHrefObject, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinks) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserResourceHrefObject and assigns it to the User field.
func (o *OAuth2RefreshTokenLinks) SetUser(v UserResourceHrefObject) {
	o.User = &v
}

// GetAuthorizationServer returns the AuthorizationServer field value if set, zero value otherwise.
func (o *OAuth2RefreshTokenLinks) GetAuthorizationServer() AuthorizationServerResourceHrefObject {
	if o == nil || IsNil(o.AuthorizationServer) {
		var ret AuthorizationServerResourceHrefObject
		return ret
	}
	return *o.AuthorizationServer
}

// GetAuthorizationServerOk returns a tuple with the AuthorizationServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2RefreshTokenLinks) GetAuthorizationServerOk() (*AuthorizationServerResourceHrefObject, bool) {
	if o == nil || IsNil(o.AuthorizationServer) {
		return nil, false
	}
	return o.AuthorizationServer, true
}

// HasAuthorizationServer returns a boolean if a field has been set.
func (o *OAuth2RefreshTokenLinks) HasAuthorizationServer() bool {
	if o != nil && !IsNil(o.AuthorizationServer) {
		return true
	}

	return false
}

// SetAuthorizationServer gets a reference to the given AuthorizationServerResourceHrefObject and assigns it to the AuthorizationServer field.
func (o *OAuth2RefreshTokenLinks) SetAuthorizationServer(v AuthorizationServerResourceHrefObject) {
	o.AuthorizationServer = &v
}

func (o OAuth2RefreshTokenLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2RefreshTokenLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Revoke) {
		toSerialize["revoke"] = o.Revoke
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
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

func (o *OAuth2RefreshTokenLinks) UnmarshalJSON(data []byte) (err error) {
	varOAuth2RefreshTokenLinks := _OAuth2RefreshTokenLinks{}

	err = json.Unmarshal(data, &varOAuth2RefreshTokenLinks)

	if err != nil {
		return err
	}

	*o = OAuth2RefreshTokenLinks(varOAuth2RefreshTokenLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "app")
		delete(additionalProperties, "revoke")
		delete(additionalProperties, "client")
		delete(additionalProperties, "user")
		delete(additionalProperties, "authorizationServer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOAuth2RefreshTokenLinks struct {
	value *OAuth2RefreshTokenLinks
	isSet bool
}

func (v NullableOAuth2RefreshTokenLinks) Get() *OAuth2RefreshTokenLinks {
	return v.value
}

func (v *NullableOAuth2RefreshTokenLinks) Set(val *OAuth2RefreshTokenLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2RefreshTokenLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2RefreshTokenLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2RefreshTokenLinks(val *OAuth2RefreshTokenLinks) *NullableOAuth2RefreshTokenLinks {
	return &NullableOAuth2RefreshTokenLinks{value: val, isSet: true}
}

func (v NullableOAuth2RefreshTokenLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2RefreshTokenLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
