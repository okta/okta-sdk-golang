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

// IdentityProviderPolicy struct for IdentityProviderPolicy
type IdentityProviderPolicy struct {
	AccountLink *PolicyAccountLink `json:"accountLink,omitempty"`
	// <div class=\"x-lifecycle-container\"><x-lifecycle class=\"ea\"></x-lifecycle> <x-lifecycle class=\"oie\"></x-lifecycle></div>Enable mapping AMR from IdP to Okta to downstream apps
	MapAMRClaims *bool `json:"mapAMRClaims,omitempty"`
	MaxClockSkew *int32 `json:"maxClockSkew,omitempty"`
	Provisioning *Provisioning `json:"provisioning,omitempty"`
	Subject *PolicySubject `json:"subject,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderPolicy IdentityProviderPolicy

// NewIdentityProviderPolicy instantiates a new IdentityProviderPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderPolicy() *IdentityProviderPolicy {
	this := IdentityProviderPolicy{}
	var mapAMRClaims bool = false
	this.MapAMRClaims = &mapAMRClaims
	return &this
}

// NewIdentityProviderPolicyWithDefaults instantiates a new IdentityProviderPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderPolicyWithDefaults() *IdentityProviderPolicy {
	this := IdentityProviderPolicy{}
	var mapAMRClaims bool = false
	this.MapAMRClaims = &mapAMRClaims
	return &this
}

// GetAccountLink returns the AccountLink field value if set, zero value otherwise.
func (o *IdentityProviderPolicy) GetAccountLink() PolicyAccountLink {
	if o == nil || o.AccountLink == nil {
		var ret PolicyAccountLink
		return ret
	}
	return *o.AccountLink
}

// GetAccountLinkOk returns a tuple with the AccountLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicy) GetAccountLinkOk() (*PolicyAccountLink, bool) {
	if o == nil || o.AccountLink == nil {
		return nil, false
	}
	return o.AccountLink, true
}

// HasAccountLink returns a boolean if a field has been set.
func (o *IdentityProviderPolicy) HasAccountLink() bool {
	if o != nil && o.AccountLink != nil {
		return true
	}

	return false
}

// SetAccountLink gets a reference to the given PolicyAccountLink and assigns it to the AccountLink field.
func (o *IdentityProviderPolicy) SetAccountLink(v PolicyAccountLink) {
	o.AccountLink = &v
}

// GetMapAMRClaims returns the MapAMRClaims field value if set, zero value otherwise.
func (o *IdentityProviderPolicy) GetMapAMRClaims() bool {
	if o == nil || o.MapAMRClaims == nil {
		var ret bool
		return ret
	}
	return *o.MapAMRClaims
}

// GetMapAMRClaimsOk returns a tuple with the MapAMRClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicy) GetMapAMRClaimsOk() (*bool, bool) {
	if o == nil || o.MapAMRClaims == nil {
		return nil, false
	}
	return o.MapAMRClaims, true
}

// HasMapAMRClaims returns a boolean if a field has been set.
func (o *IdentityProviderPolicy) HasMapAMRClaims() bool {
	if o != nil && o.MapAMRClaims != nil {
		return true
	}

	return false
}

// SetMapAMRClaims gets a reference to the given bool and assigns it to the MapAMRClaims field.
func (o *IdentityProviderPolicy) SetMapAMRClaims(v bool) {
	o.MapAMRClaims = &v
}

// GetMaxClockSkew returns the MaxClockSkew field value if set, zero value otherwise.
func (o *IdentityProviderPolicy) GetMaxClockSkew() int32 {
	if o == nil || o.MaxClockSkew == nil {
		var ret int32
		return ret
	}
	return *o.MaxClockSkew
}

// GetMaxClockSkewOk returns a tuple with the MaxClockSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicy) GetMaxClockSkewOk() (*int32, bool) {
	if o == nil || o.MaxClockSkew == nil {
		return nil, false
	}
	return o.MaxClockSkew, true
}

// HasMaxClockSkew returns a boolean if a field has been set.
func (o *IdentityProviderPolicy) HasMaxClockSkew() bool {
	if o != nil && o.MaxClockSkew != nil {
		return true
	}

	return false
}

// SetMaxClockSkew gets a reference to the given int32 and assigns it to the MaxClockSkew field.
func (o *IdentityProviderPolicy) SetMaxClockSkew(v int32) {
	o.MaxClockSkew = &v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *IdentityProviderPolicy) GetProvisioning() Provisioning {
	if o == nil || o.Provisioning == nil {
		var ret Provisioning
		return ret
	}
	return *o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicy) GetProvisioningOk() (*Provisioning, bool) {
	if o == nil || o.Provisioning == nil {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *IdentityProviderPolicy) HasProvisioning() bool {
	if o != nil && o.Provisioning != nil {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given Provisioning and assigns it to the Provisioning field.
func (o *IdentityProviderPolicy) SetProvisioning(v Provisioning) {
	o.Provisioning = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IdentityProviderPolicy) GetSubject() PolicySubject {
	if o == nil || o.Subject == nil {
		var ret PolicySubject
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderPolicy) GetSubjectOk() (*PolicySubject, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IdentityProviderPolicy) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PolicySubject and assigns it to the Subject field.
func (o *IdentityProviderPolicy) SetSubject(v PolicySubject) {
	o.Subject = &v
}

func (o IdentityProviderPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountLink != nil {
		toSerialize["accountLink"] = o.AccountLink
	}
	if o.MapAMRClaims != nil {
		toSerialize["mapAMRClaims"] = o.MapAMRClaims
	}
	if o.MaxClockSkew != nil {
		toSerialize["maxClockSkew"] = o.MaxClockSkew
	}
	if o.Provisioning != nil {
		toSerialize["provisioning"] = o.Provisioning
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderPolicy := _IdentityProviderPolicy{}

	err = json.Unmarshal(bytes, &varIdentityProviderPolicy)
	if err == nil {
		*o = IdentityProviderPolicy(varIdentityProviderPolicy)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accountLink")
		delete(additionalProperties, "mapAMRClaims")
		delete(additionalProperties, "maxClockSkew")
		delete(additionalProperties, "provisioning")
		delete(additionalProperties, "subject")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderPolicy struct {
	value *IdentityProviderPolicy
	isSet bool
}

func (v NullableIdentityProviderPolicy) Get() *IdentityProviderPolicy {
	return v.value
}

func (v *NullableIdentityProviderPolicy) Set(val *IdentityProviderPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderPolicy(val *IdentityProviderPolicy) *NullableIdentityProviderPolicy {
	return &NullableIdentityProviderPolicy{value: val, isSet: true}
}

func (v NullableIdentityProviderPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

