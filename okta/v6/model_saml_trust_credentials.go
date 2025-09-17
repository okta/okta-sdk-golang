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

// checks if the SamlTrustCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlTrustCredentials{}

// SamlTrustCredentials Federation Trust Credentials for verifying assertions from the IdP
type SamlTrustCredentials struct {
	// Additional IdP key credential reference to the Okta X.509 signature certificate
	AdditionalKids []string `json:"additionalKids,omitempty"`
	// URI that identifies the target Okta IdP instance (SP) for an `<Assertion>`
	Audience *string `json:"audience,omitempty"`
	// URI that identifies the issuer (IdP) of a `<SAMLResponse>` message `<Assertion>` element
	Issuer *string `json:"issuer,omitempty"`
	// IdP key credential reference to the Okta X.509 signature certificate
	Kid                  *string `json:"kid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SamlTrustCredentials SamlTrustCredentials

// NewSamlTrustCredentials instantiates a new SamlTrustCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlTrustCredentials() *SamlTrustCredentials {
	this := SamlTrustCredentials{}
	return &this
}

// NewSamlTrustCredentialsWithDefaults instantiates a new SamlTrustCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlTrustCredentialsWithDefaults() *SamlTrustCredentials {
	this := SamlTrustCredentials{}
	return &this
}

// GetAdditionalKids returns the AdditionalKids field value if set, zero value otherwise.
func (o *SamlTrustCredentials) GetAdditionalKids() []string {
	if o == nil || IsNil(o.AdditionalKids) {
		var ret []string
		return ret
	}
	return o.AdditionalKids
}

// GetAdditionalKidsOk returns a tuple with the AdditionalKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlTrustCredentials) GetAdditionalKidsOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalKids) {
		return nil, false
	}
	return o.AdditionalKids, true
}

// HasAdditionalKids returns a boolean if a field has been set.
func (o *SamlTrustCredentials) HasAdditionalKids() bool {
	if o != nil && !IsNil(o.AdditionalKids) {
		return true
	}

	return false
}

// SetAdditionalKids gets a reference to the given []string and assigns it to the AdditionalKids field.
func (o *SamlTrustCredentials) SetAdditionalKids(v []string) {
	o.AdditionalKids = v
}

// GetAudience returns the Audience field value if set, zero value otherwise.
func (o *SamlTrustCredentials) GetAudience() string {
	if o == nil || IsNil(o.Audience) {
		var ret string
		return ret
	}
	return *o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlTrustCredentials) GetAudienceOk() (*string, bool) {
	if o == nil || IsNil(o.Audience) {
		return nil, false
	}
	return o.Audience, true
}

// HasAudience returns a boolean if a field has been set.
func (o *SamlTrustCredentials) HasAudience() bool {
	if o != nil && !IsNil(o.Audience) {
		return true
	}

	return false
}

// SetAudience gets a reference to the given string and assigns it to the Audience field.
func (o *SamlTrustCredentials) SetAudience(v string) {
	o.Audience = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SamlTrustCredentials) GetIssuer() string {
	if o == nil || IsNil(o.Issuer) {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlTrustCredentials) GetIssuerOk() (*string, bool) {
	if o == nil || IsNil(o.Issuer) {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SamlTrustCredentials) HasIssuer() bool {
	if o != nil && !IsNil(o.Issuer) {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *SamlTrustCredentials) SetIssuer(v string) {
	o.Issuer = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *SamlTrustCredentials) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlTrustCredentials) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *SamlTrustCredentials) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *SamlTrustCredentials) SetKid(v string) {
	o.Kid = &v
}

func (o SamlTrustCredentials) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlTrustCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalKids) {
		toSerialize["additionalKids"] = o.AdditionalKids
	}
	if !IsNil(o.Audience) {
		toSerialize["audience"] = o.Audience
	}
	if !IsNil(o.Issuer) {
		toSerialize["issuer"] = o.Issuer
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SamlTrustCredentials) UnmarshalJSON(data []byte) (err error) {
	varSamlTrustCredentials := _SamlTrustCredentials{}

	err = json.Unmarshal(data, &varSamlTrustCredentials)

	if err != nil {
		return err
	}

	*o = SamlTrustCredentials(varSamlTrustCredentials)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "additionalKids")
		delete(additionalProperties, "audience")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "kid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSamlTrustCredentials struct {
	value *SamlTrustCredentials
	isSet bool
}

func (v NullableSamlTrustCredentials) Get() *SamlTrustCredentials {
	return v.value
}

func (v *NullableSamlTrustCredentials) Set(val *SamlTrustCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlTrustCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlTrustCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlTrustCredentials(val *SamlTrustCredentials) *NullableSamlTrustCredentials {
	return &NullableSamlTrustCredentials{value: val, isSet: true}
}

func (v NullableSamlTrustCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlTrustCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
