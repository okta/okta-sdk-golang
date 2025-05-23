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

// IdentityProviderLinks struct for IdentityProviderLinks
type IdentityProviderLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Acs *HrefObject `json:"acs,omitempty"`
	Authorize *HrefObject `json:"authorize,omitempty"`
	ClientRedirectUri *HrefObject `json:"clientRedirectUri,omitempty"`
	Metadata *HrefObject `json:"metadata,omitempty"`
	Users *HrefObject `json:"users,omitempty"`
	Deactivate *HrefObject `json:"deactivate,omitempty"`
	Activate *HrefObject `json:"activate,omitempty"`
	Keys *HrefObject `json:"keys,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderLinks IdentityProviderLinks

// NewIdentityProviderLinks instantiates a new IdentityProviderLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderLinks() *IdentityProviderLinks {
	this := IdentityProviderLinks{}
	return &this
}

// NewIdentityProviderLinksWithDefaults instantiates a new IdentityProviderLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderLinksWithDefaults() *IdentityProviderLinks {
	this := IdentityProviderLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *IdentityProviderLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetAcs returns the Acs field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetAcs() HrefObject {
	if o == nil || o.Acs == nil {
		var ret HrefObject
		return ret
	}
	return *o.Acs
}

// GetAcsOk returns a tuple with the Acs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetAcsOk() (*HrefObject, bool) {
	if o == nil || o.Acs == nil {
		return nil, false
	}
	return o.Acs, true
}

// HasAcs returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasAcs() bool {
	if o != nil && o.Acs != nil {
		return true
	}

	return false
}

// SetAcs gets a reference to the given HrefObject and assigns it to the Acs field.
func (o *IdentityProviderLinks) SetAcs(v HrefObject) {
	o.Acs = &v
}

// GetAuthorize returns the Authorize field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetAuthorize() HrefObject {
	if o == nil || o.Authorize == nil {
		var ret HrefObject
		return ret
	}
	return *o.Authorize
}

// GetAuthorizeOk returns a tuple with the Authorize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetAuthorizeOk() (*HrefObject, bool) {
	if o == nil || o.Authorize == nil {
		return nil, false
	}
	return o.Authorize, true
}

// HasAuthorize returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasAuthorize() bool {
	if o != nil && o.Authorize != nil {
		return true
	}

	return false
}

// SetAuthorize gets a reference to the given HrefObject and assigns it to the Authorize field.
func (o *IdentityProviderLinks) SetAuthorize(v HrefObject) {
	o.Authorize = &v
}

// GetClientRedirectUri returns the ClientRedirectUri field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetClientRedirectUri() HrefObject {
	if o == nil || o.ClientRedirectUri == nil {
		var ret HrefObject
		return ret
	}
	return *o.ClientRedirectUri
}

// GetClientRedirectUriOk returns a tuple with the ClientRedirectUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetClientRedirectUriOk() (*HrefObject, bool) {
	if o == nil || o.ClientRedirectUri == nil {
		return nil, false
	}
	return o.ClientRedirectUri, true
}

// HasClientRedirectUri returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasClientRedirectUri() bool {
	if o != nil && o.ClientRedirectUri != nil {
		return true
	}

	return false
}

// SetClientRedirectUri gets a reference to the given HrefObject and assigns it to the ClientRedirectUri field.
func (o *IdentityProviderLinks) SetClientRedirectUri(v HrefObject) {
	o.ClientRedirectUri = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetMetadata() HrefObject {
	if o == nil || o.Metadata == nil {
		var ret HrefObject
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetMetadataOk() (*HrefObject, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given HrefObject and assigns it to the Metadata field.
func (o *IdentityProviderLinks) SetMetadata(v HrefObject) {
	o.Metadata = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetUsers() HrefObject {
	if o == nil || o.Users == nil {
		var ret HrefObject
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetUsersOk() (*HrefObject, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given HrefObject and assigns it to the Users field.
func (o *IdentityProviderLinks) SetUsers(v HrefObject) {
	o.Users = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetDeactivate() HrefObject {
	if o == nil || o.Deactivate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *IdentityProviderLinks) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetActivate() HrefObject {
	if o == nil || o.Activate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetActivateOk() (*HrefObject, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObject and assigns it to the Activate field.
func (o *IdentityProviderLinks) SetActivate(v HrefObject) {
	o.Activate = &v
}

// GetKeys returns the Keys field value if set, zero value otherwise.
func (o *IdentityProviderLinks) GetKeys() HrefObject {
	if o == nil || o.Keys == nil {
		var ret HrefObject
		return ret
	}
	return *o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderLinks) GetKeysOk() (*HrefObject, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *IdentityProviderLinks) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given HrefObject and assigns it to the Keys field.
func (o *IdentityProviderLinks) SetKeys(v HrefObject) {
	o.Keys = &v
}

func (o IdentityProviderLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Acs != nil {
		toSerialize["acs"] = o.Acs
	}
	if o.Authorize != nil {
		toSerialize["authorize"] = o.Authorize
	}
	if o.ClientRedirectUri != nil {
		toSerialize["clientRedirectUri"] = o.ClientRedirectUri
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderLinks) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderLinks := _IdentityProviderLinks{}

	err = json.Unmarshal(bytes, &varIdentityProviderLinks)
	if err == nil {
		*o = IdentityProviderLinks(varIdentityProviderLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "acs")
		delete(additionalProperties, "authorize")
		delete(additionalProperties, "clientRedirectUri")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "users")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "keys")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderLinks struct {
	value *IdentityProviderLinks
	isSet bool
}

func (v NullableIdentityProviderLinks) Get() *IdentityProviderLinks {
	return v.value
}

func (v *NullableIdentityProviderLinks) Set(val *IdentityProviderLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderLinks(val *IdentityProviderLinks) *NullableIdentityProviderLinks {
	return &NullableIdentityProviderLinks{value: val, isSet: true}
}

func (v NullableIdentityProviderLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

