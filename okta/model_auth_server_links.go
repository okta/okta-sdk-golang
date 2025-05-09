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

// AuthServerLinks struct for AuthServerLinks
type AuthServerLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Claims *AuthServerLinksAllOfClaims `json:"claims,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	// Link to the authorization server metadata
	Metadata []HrefObject `json:"metadata,omitempty"`
	Policies *AuthServerLinksAllOfPolicies `json:"policies,omitempty"`
	RotateKey *AuthServerLinksAllOfRotateKey `json:"rotateKey,omitempty"`
	Scopes *AuthServerLinksAllOfScopes `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthServerLinks AuthServerLinks

// NewAuthServerLinks instantiates a new AuthServerLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthServerLinks() *AuthServerLinks {
	this := AuthServerLinks{}
	return &this
}

// NewAuthServerLinksWithDefaults instantiates a new AuthServerLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthServerLinksWithDefaults() *AuthServerLinks {
	this := AuthServerLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *AuthServerLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthServerLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *AuthServerLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *AuthServerLinks) GetClaims() AuthServerLinksAllOfClaims {
	if o == nil || o.Claims == nil {
		var ret AuthServerLinksAllOfClaims
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetClaimsOk() (*AuthServerLinksAllOfClaims, bool) {
	if o == nil || o.Claims == nil {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *AuthServerLinks) HasClaims() bool {
	if o != nil && o.Claims != nil {
		return true
	}

	return false
}

// SetClaims gets a reference to the given AuthServerLinksAllOfClaims and assigns it to the Claims field.
func (o *AuthServerLinks) SetClaims(v AuthServerLinksAllOfClaims) {
	o.Claims = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *AuthServerLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *AuthServerLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *AuthServerLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AuthServerLinks) GetMetadata() []HrefObject {
	if o == nil || o.Metadata == nil {
		var ret []HrefObject
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetMetadataOk() ([]HrefObject, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AuthServerLinks) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []HrefObject and assigns it to the Metadata field.
func (o *AuthServerLinks) SetMetadata(v []HrefObject) {
	o.Metadata = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *AuthServerLinks) GetPolicies() AuthServerLinksAllOfPolicies {
	if o == nil || o.Policies == nil {
		var ret AuthServerLinksAllOfPolicies
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetPoliciesOk() (*AuthServerLinksAllOfPolicies, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AuthServerLinks) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given AuthServerLinksAllOfPolicies and assigns it to the Policies field.
func (o *AuthServerLinks) SetPolicies(v AuthServerLinksAllOfPolicies) {
	o.Policies = &v
}

// GetRotateKey returns the RotateKey field value if set, zero value otherwise.
func (o *AuthServerLinks) GetRotateKey() AuthServerLinksAllOfRotateKey {
	if o == nil || o.RotateKey == nil {
		var ret AuthServerLinksAllOfRotateKey
		return ret
	}
	return *o.RotateKey
}

// GetRotateKeyOk returns a tuple with the RotateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetRotateKeyOk() (*AuthServerLinksAllOfRotateKey, bool) {
	if o == nil || o.RotateKey == nil {
		return nil, false
	}
	return o.RotateKey, true
}

// HasRotateKey returns a boolean if a field has been set.
func (o *AuthServerLinks) HasRotateKey() bool {
	if o != nil && o.RotateKey != nil {
		return true
	}

	return false
}

// SetRotateKey gets a reference to the given AuthServerLinksAllOfRotateKey and assigns it to the RotateKey field.
func (o *AuthServerLinks) SetRotateKey(v AuthServerLinksAllOfRotateKey) {
	o.RotateKey = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthServerLinks) GetScopes() AuthServerLinksAllOfScopes {
	if o == nil || o.Scopes == nil {
		var ret AuthServerLinksAllOfScopes
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetScopesOk() (*AuthServerLinksAllOfScopes, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthServerLinks) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given AuthServerLinksAllOfScopes and assigns it to the Scopes field.
func (o *AuthServerLinks) SetScopes(v AuthServerLinksAllOfScopes) {
	o.Scopes = &v
}

func (o AuthServerLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Claims != nil {
		toSerialize["claims"] = o.Claims
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Policies != nil {
		toSerialize["policies"] = o.Policies
	}
	if o.RotateKey != nil {
		toSerialize["rotateKey"] = o.RotateKey
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AuthServerLinks) UnmarshalJSON(bytes []byte) (err error) {
	varAuthServerLinks := _AuthServerLinks{}

	err = json.Unmarshal(bytes, &varAuthServerLinks)
	if err == nil {
		*o = AuthServerLinks(varAuthServerLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "claims")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "policies")
		delete(additionalProperties, "rotateKey")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAuthServerLinks struct {
	value *AuthServerLinks
	isSet bool
}

func (v NullableAuthServerLinks) Get() *AuthServerLinks {
	return v.value
}

func (v *NullableAuthServerLinks) Set(val *AuthServerLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthServerLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthServerLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthServerLinks(val *AuthServerLinks) *NullableAuthServerLinks {
	return &NullableAuthServerLinks{value: val, isSet: true}
}

func (v NullableAuthServerLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthServerLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

