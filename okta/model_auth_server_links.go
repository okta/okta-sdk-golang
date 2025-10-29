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

// checks if the AuthServerLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthServerLinks{}

// AuthServerLinks struct for AuthServerLinks
type AuthServerLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	// Link to the authorization server claims
	Claims     *HrefObject               `json:"claims,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	// Link to the authorization server metadata
	Metadata []HrefObject `json:"metadata,omitempty"`
	// Link to the authorization server policies
	Policies *HrefObject `json:"policies,omitempty"`
	// Link to the authorization server key rotation
	RotateKey *HrefObject `json:"rotateKey,omitempty"`
	// Link to the authorization server scopes
	Scopes               *HrefObject `json:"scopes,omitempty"`
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
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *AuthServerLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *AuthServerLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *AuthServerLinks) GetClaims() HrefObject {
	if o == nil || IsNil(o.Claims) {
		var ret HrefObject
		return ret
	}
	return *o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetClaimsOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *AuthServerLinks) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given HrefObject and assigns it to the Claims field.
func (o *AuthServerLinks) SetClaims(v HrefObject) {
	o.Claims = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *AuthServerLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *AuthServerLinks) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
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
	if o == nil || IsNil(o.Metadata) {
		var ret []HrefObject
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetMetadataOk() ([]HrefObject, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AuthServerLinks) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []HrefObject and assigns it to the Metadata field.
func (o *AuthServerLinks) SetMetadata(v []HrefObject) {
	o.Metadata = v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *AuthServerLinks) GetPolicies() HrefObject {
	if o == nil || IsNil(o.Policies) {
		var ret HrefObject
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetPoliciesOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *AuthServerLinks) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given HrefObject and assigns it to the Policies field.
func (o *AuthServerLinks) SetPolicies(v HrefObject) {
	o.Policies = &v
}

// GetRotateKey returns the RotateKey field value if set, zero value otherwise.
func (o *AuthServerLinks) GetRotateKey() HrefObject {
	if o == nil || IsNil(o.RotateKey) {
		var ret HrefObject
		return ret
	}
	return *o.RotateKey
}

// GetRotateKeyOk returns a tuple with the RotateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetRotateKeyOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.RotateKey) {
		return nil, false
	}
	return o.RotateKey, true
}

// HasRotateKey returns a boolean if a field has been set.
func (o *AuthServerLinks) HasRotateKey() bool {
	if o != nil && !IsNil(o.RotateKey) {
		return true
	}

	return false
}

// SetRotateKey gets a reference to the given HrefObject and assigns it to the RotateKey field.
func (o *AuthServerLinks) SetRotateKey(v HrefObject) {
	o.RotateKey = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *AuthServerLinks) GetScopes() HrefObject {
	if o == nil || IsNil(o.Scopes) {
		var ret HrefObject
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthServerLinks) GetScopesOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *AuthServerLinks) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given HrefObject and assigns it to the Scopes field.
func (o *AuthServerLinks) SetScopes(v HrefObject) {
	o.Scopes = &v
}

func (o AuthServerLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthServerLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.RotateKey) {
		toSerialize["rotateKey"] = o.RotateKey
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthServerLinks) UnmarshalJSON(data []byte) (err error) {
	varAuthServerLinks := _AuthServerLinks{}

	err = json.Unmarshal(data, &varAuthServerLinks)

	if err != nil {
		return err
	}

	*o = AuthServerLinks(varAuthServerLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "claims")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "policies")
		delete(additionalProperties, "rotateKey")
		delete(additionalProperties, "scopes")
		o.AdditionalProperties = additionalProperties
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
