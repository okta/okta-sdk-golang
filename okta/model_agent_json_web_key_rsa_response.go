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
	"fmt"
)

// checks if the AgentJsonWebKeyRsaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonWebKeyRsaResponse{}

// AgentJsonWebKeyRsaResponse An RSA signing key
type AgentJsonWebKeyRsaResponse struct {
	// The public exponent of the RSA key, represented as a Base64URL-encoded string.  This value is used in combination with the modulus (`n`) to verify signatures and encrypt data.
	E *string `json:"e,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty string `json:"kty"`
	// The modulus of the RSA public key, represented as a Base64URL-encoded string.  This is the primary component of the RSA key and, with the exponent (`e`), is used for cryptographic signature verification and encryption.
	N *string `json:"n,omitempty"`
	// Unique identifier of the JSON Web Key in the AI agent's JSON Web Key Set (JWKS)
	Kid *string `json:"kid,omitempty"`
	// Status of the AI agent JSON Web Key
	Status *string `json:"status,omitempty"`
	// Algorithm that's used in the JSON Web Key
	Alg *string `json:"alg,omitempty"`
	// Acceptable use of the JSON Web Key  You can only use signing keys for AI agents, so the value of `use` is always `sig`.
	Use *string `json:"use,omitempty"`
	// Timestamp of when the AI agent JSON Web Key was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the AI agent JSON Web Key
	Id *string `json:"id,omitempty"`
	// Timestamp of when the AI agent JSON Web Key was last updated
	LastUpdated          *string           `json:"lastUpdated,omitempty"`
	Links                *AgentSecretLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AgentJsonWebKeyRsaResponse AgentJsonWebKeyRsaResponse

// NewAgentJsonWebKeyRsaResponse instantiates a new AgentJsonWebKeyRsaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonWebKeyRsaResponse(kty string) *AgentJsonWebKeyRsaResponse {
	this := AgentJsonWebKeyRsaResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewAgentJsonWebKeyRsaResponseWithDefaults instantiates a new AgentJsonWebKeyRsaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonWebKeyRsaResponseWithDefaults() *AgentJsonWebKeyRsaResponse {
	this := AgentJsonWebKeyRsaResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetE returns the E field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetE() string {
	if o == nil || IsNil(o.E) {
		var ret string
		return ret
	}
	return *o.E
}

// GetEOk returns a tuple with the E field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetEOk() (*string, bool) {
	if o == nil || IsNil(o.E) {
		return nil, false
	}
	return o.E, true
}

// HasE returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasE() bool {
	if o != nil && !IsNil(o.E) {
		return true
	}

	return false
}

// SetE gets a reference to the given string and assigns it to the E field.
func (o *AgentJsonWebKeyRsaResponse) SetE(v string) {
	o.E = &v
}

// GetKty returns the Kty field value
func (o *AgentJsonWebKeyRsaResponse) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *AgentJsonWebKeyRsaResponse) SetKty(v string) {
	o.Kty = v
}

// GetN returns the N field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetN() string {
	if o == nil || IsNil(o.N) {
		var ret string
		return ret
	}
	return *o.N
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetNOk() (*string, bool) {
	if o == nil || IsNil(o.N) {
		return nil, false
	}
	return o.N, true
}

// HasN returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasN() bool {
	if o != nil && !IsNil(o.N) {
		return true
	}

	return false
}

// SetN gets a reference to the given string and assigns it to the N field.
func (o *AgentJsonWebKeyRsaResponse) SetN(v string) {
	o.N = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AgentJsonWebKeyRsaResponse) SetKid(v string) {
	o.Kid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentJsonWebKeyRsaResponse) SetStatus(v string) {
	o.Status = &v
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *AgentJsonWebKeyRsaResponse) SetAlg(v string) {
	o.Alg = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AgentJsonWebKeyRsaResponse) SetUse(v string) {
	o.Use = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *AgentJsonWebKeyRsaResponse) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentJsonWebKeyRsaResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *AgentJsonWebKeyRsaResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentJsonWebKeyRsaResponse) GetLinks() AgentSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret AgentSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyRsaResponse) GetLinksOk() (*AgentSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentJsonWebKeyRsaResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AgentSecretLinks and assigns it to the Links field.
func (o *AgentJsonWebKeyRsaResponse) SetLinks(v AgentSecretLinks) {
	o.Links = &v
}

func (o AgentJsonWebKeyRsaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonWebKeyRsaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.E) {
		toSerialize["e"] = o.E
	}
	toSerialize["kty"] = o.Kty
	if !IsNil(o.N) {
		toSerialize["n"] = o.N
	}
	if !IsNil(o.Kid) {
		toSerialize["kid"] = o.Kid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Alg) {
		toSerialize["alg"] = o.Alg
	}
	if !IsNil(o.Use) {
		toSerialize["use"] = o.Use
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AgentJsonWebKeyRsaResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kty",
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

	varAgentJsonWebKeyRsaResponse := _AgentJsonWebKeyRsaResponse{}

	err = json.Unmarshal(data, &varAgentJsonWebKeyRsaResponse)

	if err != nil {
		return err
	}

	*o = AgentJsonWebKeyRsaResponse(varAgentJsonWebKeyRsaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "e")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "n")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "alg")
		delete(additionalProperties, "use")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAgentJsonWebKeyRsaResponse struct {
	value *AgentJsonWebKeyRsaResponse
	isSet bool
}

func (v NullableAgentJsonWebKeyRsaResponse) Get() *AgentJsonWebKeyRsaResponse {
	return v.value
}

func (v *NullableAgentJsonWebKeyRsaResponse) Set(val *AgentJsonWebKeyRsaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonWebKeyRsaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonWebKeyRsaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonWebKeyRsaResponse(val *AgentJsonWebKeyRsaResponse) *NullableAgentJsonWebKeyRsaResponse {
	return &NullableAgentJsonWebKeyRsaResponse{value: val, isSet: true}
}

func (v NullableAgentJsonWebKeyRsaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonWebKeyRsaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
