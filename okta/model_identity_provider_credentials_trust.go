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

// IdentityProviderCredentialsTrust struct for IdentityProviderCredentialsTrust
type IdentityProviderCredentialsTrust struct {
	Audience *string `json:"audience,omitempty"`
	Issuer *string `json:"issuer,omitempty"`
	Kid *string `json:"kid,omitempty"`
	Revocation *string `json:"revocation,omitempty"`
	RevocationCacheLifetime *int32 `json:"revocationCacheLifetime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderCredentialsTrust IdentityProviderCredentialsTrust

// NewIdentityProviderCredentialsTrust instantiates a new IdentityProviderCredentialsTrust object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCredentialsTrust() *IdentityProviderCredentialsTrust {
	this := IdentityProviderCredentialsTrust{}
	return &this
}

// NewIdentityProviderCredentialsTrustWithDefaults instantiates a new IdentityProviderCredentialsTrust object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCredentialsTrustWithDefaults() *IdentityProviderCredentialsTrust {
	this := IdentityProviderCredentialsTrust{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsTrust) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsTrust) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsTrust) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *IdentityProviderCredentialsTrust) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsTrust) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsTrust) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsTrust) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *IdentityProviderCredentialsTrust) SetIssuer(v string) {
	o.Issuer = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsTrust) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsTrust) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsTrust) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *IdentityProviderCredentialsTrust) SetKid(v string) {
	o.Kid = &v
}

// GetRevocation returns the Revocation field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsTrust) GetRevocation() string {
	if o == nil || o.Revocation == nil {
		var ret string
		return ret
	}
	return *o.Revocation
}

// GetRevocationOk returns a tuple with the Revocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsTrust) GetRevocationOk() (*string, bool) {
	if o == nil || o.Revocation == nil {
		return nil, false
	}
	return o.Revocation, true
}

// HasRevocation returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsTrust) HasRevocation() bool {
	if o != nil && o.Revocation != nil {
		return true
	}

	return false
}

// SetRevocation gets a reference to the given string and assigns it to the Revocation field.
func (o *IdentityProviderCredentialsTrust) SetRevocation(v string) {
	o.Revocation = &v
}

// GetRevocationCacheLifetime returns the RevocationCacheLifetime field value if set, zero value otherwise.
func (o *IdentityProviderCredentialsTrust) GetRevocationCacheLifetime() int32 {
	if o == nil || o.RevocationCacheLifetime == nil {
		var ret int32
		return ret
	}
	return *o.RevocationCacheLifetime
}

// GetRevocationCacheLifetimeOk returns a tuple with the RevocationCacheLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCredentialsTrust) GetRevocationCacheLifetimeOk() (*int32, bool) {
	if o == nil || o.RevocationCacheLifetime == nil {
		return nil, false
	}
	return o.RevocationCacheLifetime, true
}

// HasRevocationCacheLifetime returns a boolean if a field has been set.
func (o *IdentityProviderCredentialsTrust) HasRevocationCacheLifetime() bool {
	if o != nil && o.RevocationCacheLifetime != nil {
		return true
	}

	return false
}

// SetRevocationCacheLifetime gets a reference to the given int32 and assigns it to the RevocationCacheLifetime field.
func (o *IdentityProviderCredentialsTrust) SetRevocationCacheLifetime(v int32) {
	o.RevocationCacheLifetime = &v
}

func (o IdentityProviderCredentialsTrust) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Audience != nil {
		toSerialize["audience"] = o.Audience
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.Revocation != nil {
		toSerialize["revocation"] = o.Revocation
	}
	if o.RevocationCacheLifetime != nil {
		toSerialize["revocationCacheLifetime"] = o.RevocationCacheLifetime
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderCredentialsTrust) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderCredentialsTrust := _IdentityProviderCredentialsTrust{}

	err = json.Unmarshal(bytes, &varIdentityProviderCredentialsTrust)
	if err == nil {
		*o = IdentityProviderCredentialsTrust(varIdentityProviderCredentialsTrust)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "audience")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "revocation")
		delete(additionalProperties, "revocationCacheLifetime")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderCredentialsTrust struct {
	value *IdentityProviderCredentialsTrust
	isSet bool
}

func (v NullableIdentityProviderCredentialsTrust) Get() *IdentityProviderCredentialsTrust {
	return v.value
}

func (v *NullableIdentityProviderCredentialsTrust) Set(val *IdentityProviderCredentialsTrust) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCredentialsTrust) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCredentialsTrust) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCredentialsTrust(val *IdentityProviderCredentialsTrust) *NullableIdentityProviderCredentialsTrust {
	return &NullableIdentityProviderCredentialsTrust{value: val, isSet: true}
}

func (v NullableIdentityProviderCredentialsTrust) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCredentialsTrust) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

