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
	"fmt"
)

// checks if the Saml type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Saml{}

// Saml SAML configuration details
type Saml struct {
	// List of Assertion Consumer Service (ACS) URLs. The default ACS URL is required and is indicated by a null `index` value. You can use the org-level variables you defined in the `config` array in the URL. For example: `https://${org.subdomain}.example.com/saml/login`
	Acs []SamlAcsInner `json:"acs"`
	// Attribute statements to appear in the Okta SAML assertion
	Claims []SamlClaimsInner `json:"claims,omitempty"`
	// The URL to your customer-facing instructions for configuring your SAML integration. See [Customer configuration document guidelines](https://developer.okta.com/docs/guides/submit-app-prereq/main/#customer-configuration-document-guidelines).
	Doc string `json:"doc"`
	// Globally unique name for your SAML entity. For instance, your Identity Provider (IdP) or Service Provider (SP) URL.
	EntityId string `json:"entityId"`
	// Defines the group attribute names for the SAML assertion statement. Okta inserts the list of Okta user groups into the attribute names in the statement.
	Groups               []string `json:"groups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Saml Saml

// NewSaml instantiates a new Saml object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaml(acs []SamlAcsInner, doc string, entityId string) *Saml {
	this := Saml{}
	this.Acs = acs
	this.Doc = doc
	this.EntityId = entityId
	return &this
}

// NewSamlWithDefaults instantiates a new Saml object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlWithDefaults() *Saml {
	this := Saml{}
	return &this
}

// GetAcs returns the Acs field value
func (o *Saml) GetAcs() []SamlAcsInner {
	if o == nil {
		var ret []SamlAcsInner
		return ret
	}

	return o.Acs
}

// GetAcsOk returns a tuple with the Acs field value
// and a boolean to check if the value has been set.
func (o *Saml) GetAcsOk() ([]SamlAcsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Acs, true
}

// SetAcs sets field value
func (o *Saml) SetAcs(v []SamlAcsInner) {
	o.Acs = v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *Saml) GetClaims() []SamlClaimsInner {
	if o == nil || IsNil(o.Claims) {
		var ret []SamlClaimsInner
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml) GetClaimsOk() ([]SamlClaimsInner, bool) {
	if o == nil || IsNil(o.Claims) {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *Saml) HasClaims() bool {
	if o != nil && !IsNil(o.Claims) {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []SamlClaimsInner and assigns it to the Claims field.
func (o *Saml) SetClaims(v []SamlClaimsInner) {
	o.Claims = v
}

// GetDoc returns the Doc field value
func (o *Saml) GetDoc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Doc
}

// GetDocOk returns a tuple with the Doc field value
// and a boolean to check if the value has been set.
func (o *Saml) GetDocOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Doc, true
}

// SetDoc sets field value
func (o *Saml) SetDoc(v string) {
	o.Doc = v
}

// GetEntityId returns the EntityId field value
func (o *Saml) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *Saml) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *Saml) SetEntityId(v string) {
	o.EntityId = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *Saml) GetGroups() []string {
	if o == nil || IsNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Saml) GetGroupsOk() ([]string, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *Saml) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *Saml) SetGroups(v []string) {
	o.Groups = v
}

func (o Saml) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Saml) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acs"] = o.Acs
	if !IsNil(o.Claims) {
		toSerialize["claims"] = o.Claims
	}
	toSerialize["doc"] = o.Doc
	toSerialize["entityId"] = o.EntityId
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Saml) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"acs",
		"doc",
		"entityId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSaml := _Saml{}

	err = json.Unmarshal(data, &varSaml)

	if err != nil {
		return err
	}

	*o = Saml(varSaml)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "acs")
		delete(additionalProperties, "claims")
		delete(additionalProperties, "doc")
		delete(additionalProperties, "entityId")
		delete(additionalProperties, "groups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSaml struct {
	value *Saml
	isSet bool
}

func (v NullableSaml) Get() *Saml {
	return v.value
}

func (v *NullableSaml) Set(val *Saml) {
	v.value = val
	v.isSet = true
}

func (v NullableSaml) IsSet() bool {
	return v.isSet
}

func (v *NullableSaml) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaml(val *Saml) *NullableSaml {
	return &NullableSaml{value: val, isSet: true}
}

func (v NullableSaml) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaml) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
