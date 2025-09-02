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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// MtlsTrustCredentials struct for MtlsTrustCredentials
type MtlsTrustCredentials struct {
	// Not used
	Audience *string `json:"audience,omitempty"`
	// Description of the certificate issuer
	Issuer *string `json:"issuer,omitempty"`
	// IdP key credential reference to the Okta X.509 signature certificate
	Kid *string `json:"kid,omitempty"`
	// Mechanism to validate the certificate
	Revocation *string `json:"revocation,omitempty"`
	// Time in minutes to cache the certificate revocation information
	RevocationCacheLifetime *float32 `json:"revocationCacheLifetime,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MtlsTrustCredentials MtlsTrustCredentials

// NewMtlsTrustCredentials instantiates a new MtlsTrustCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtlsTrustCredentials() *MtlsTrustCredentials {
	this := MtlsTrustCredentials{}
	return &this
}

// NewMtlsTrustCredentialsWithDefaults instantiates a new MtlsTrustCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtlsTrustCredentialsWithDefaults() *MtlsTrustCredentials {
	this := MtlsTrustCredentials{}
	return &this
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *MtlsTrustCredentials) GetAudience() string {
	if o == nil || o.Audience == nil {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsTrustCredentials) GetAudienceOk() (*string, bool) {
	if o == nil || o.Audience == nil {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *MtlsTrustCredentials) HasAudience() bool {
	if o != nil && o.Audience != nil {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *MtlsTrustCredentials) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *MtlsTrustCredentials) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsTrustCredentials) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *MtlsTrustCredentials) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *MtlsTrustCredentials) SetIssuer(v string) {
	o.Issuer = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *MtlsTrustCredentials) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsTrustCredentials) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *MtlsTrustCredentials) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *MtlsTrustCredentials) SetKid(v string) {
	o.Kid = &v
}

// GetRevocation returns the Revocation field value if set, zero value otherwise.
func (o *MtlsTrustCredentials) GetRevocation() string {
	if o == nil || o.Revocation == nil {
		var ret string
		return ret
	}
	return *o.Revocation
}

// GetRevocationOk returns a tuple with the Revocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsTrustCredentials) GetRevocationOk() (*string, bool) {
	if o == nil || o.Revocation == nil {
		return nil, false
	}
	return o.Revocation, true
}

// HasRevocation returns a boolean if a field has been set.
func (o *MtlsTrustCredentials) HasRevocation() bool {
	if o != nil && o.Revocation != nil {
		return true
	}

	return false
}

// SetRevocation gets a reference to the given string and assigns it to the Revocation field.
func (o *MtlsTrustCredentials) SetRevocation(v string) {
	o.Revocation = &v
}

// GetRevocationCacheLifetime returns the RevocationCacheLifetime field value if set, zero value otherwise.
func (o *MtlsTrustCredentials) GetRevocationCacheLifetime() float32 {
	if o == nil || o.RevocationCacheLifetime == nil {
		var ret float32
		return ret
	}
	return *o.RevocationCacheLifetime
}

// GetRevocationCacheLifetimeOk returns a tuple with the RevocationCacheLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtlsTrustCredentials) GetRevocationCacheLifetimeOk() (*float32, bool) {
	if o == nil || o.RevocationCacheLifetime == nil {
		return nil, false
	}
	return o.RevocationCacheLifetime, true
}

// HasRevocationCacheLifetime returns a boolean if a field has been set.
func (o *MtlsTrustCredentials) HasRevocationCacheLifetime() bool {
	if o != nil && o.RevocationCacheLifetime != nil {
		return true
	}

	return false
}

// SetRevocationCacheLifetime gets a reference to the given float32 and assigns it to the RevocationCacheLifetime field.
func (o *MtlsTrustCredentials) SetRevocationCacheLifetime(v float32) {
	o.RevocationCacheLifetime = &v
}

func (o MtlsTrustCredentials) MarshalJSON() ([]byte, error) {
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

func (o *MtlsTrustCredentials) UnmarshalJSON(bytes []byte) (err error) {
	varMtlsTrustCredentials := _MtlsTrustCredentials{}

	err = json.Unmarshal(bytes, &varMtlsTrustCredentials)
	if err == nil {
		*o = MtlsTrustCredentials(varMtlsTrustCredentials)
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

type NullableMtlsTrustCredentials struct {
	value *MtlsTrustCredentials
	isSet bool
}

func (v NullableMtlsTrustCredentials) Get() *MtlsTrustCredentials {
	return v.value
}

func (v *NullableMtlsTrustCredentials) Set(val *MtlsTrustCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableMtlsTrustCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableMtlsTrustCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtlsTrustCredentials(val *MtlsTrustCredentials) *NullableMtlsTrustCredentials {
	return &NullableMtlsTrustCredentials{value: val, isSet: true}
}

func (v NullableMtlsTrustCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtlsTrustCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

