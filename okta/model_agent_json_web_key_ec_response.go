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

// checks if the AgentJsonWebKeyECResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentJsonWebKeyECResponse{}

// AgentJsonWebKeyECResponse An EC signing key
type AgentJsonWebKeyECResponse struct {
	// The cryptographic curve that's used for the key pair
	Crv *string `json:"crv,omitempty"`
	// Cryptographic algorithm family for the certificate's key pair
	Kty string `json:"kty"`
	// The public x coordinate for the elliptic curve point
	X *string `json:"x,omitempty"`
	// The public y coordinate for the elliptic curve point
	Y *string `json:"y,omitempty"`
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

type _AgentJsonWebKeyECResponse AgentJsonWebKeyECResponse

// NewAgentJsonWebKeyECResponse instantiates a new AgentJsonWebKeyECResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentJsonWebKeyECResponse(kty string) *AgentJsonWebKeyECResponse {
	this := AgentJsonWebKeyECResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// NewAgentJsonWebKeyECResponseWithDefaults instantiates a new AgentJsonWebKeyECResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentJsonWebKeyECResponseWithDefaults() *AgentJsonWebKeyECResponse {
	this := AgentJsonWebKeyECResponse{}
	var status string = "ACTIVE"
	this.Status = &status
	return &this
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetCrv() string {
	if o == nil || IsNil(o.Crv) {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetCrvOk() (*string, bool) {
	if o == nil || IsNil(o.Crv) {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasCrv() bool {
	if o != nil && !IsNil(o.Crv) {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *AgentJsonWebKeyECResponse) SetCrv(v string) {
	o.Crv = &v
}

// GetKty returns the Kty field value
func (o *AgentJsonWebKeyECResponse) GetKty() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kty
}

// GetKtyOk returns a tuple with the Kty field value
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetKtyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kty, true
}

// SetKty sets field value
func (o *AgentJsonWebKeyECResponse) SetKty(v string) {
	o.Kty = v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetX() string {
	if o == nil || IsNil(o.X) {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetXOk() (*string, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *AgentJsonWebKeyECResponse) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetY() string {
	if o == nil || IsNil(o.Y) {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetYOk() (*string, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *AgentJsonWebKeyECResponse) SetY(v string) {
	o.Y = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetKid() string {
	if o == nil || IsNil(o.Kid) {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetKidOk() (*string, bool) {
	if o == nil || IsNil(o.Kid) {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasKid() bool {
	if o != nil && !IsNil(o.Kid) {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *AgentJsonWebKeyECResponse) SetKid(v string) {
	o.Kid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentJsonWebKeyECResponse) SetStatus(v string) {
	o.Status = &v
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetAlg() string {
	if o == nil || IsNil(o.Alg) {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetAlgOk() (*string, bool) {
	if o == nil || IsNil(o.Alg) {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasAlg() bool {
	if o != nil && !IsNil(o.Alg) {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *AgentJsonWebKeyECResponse) SetAlg(v string) {
	o.Alg = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetUse() string {
	if o == nil || IsNil(o.Use) {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetUseOk() (*string, bool) {
	if o == nil || IsNil(o.Use) {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasUse() bool {
	if o != nil && !IsNil(o.Use) {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *AgentJsonWebKeyECResponse) SetUse(v string) {
	o.Use = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *AgentJsonWebKeyECResponse) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentJsonWebKeyECResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *AgentJsonWebKeyECResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentJsonWebKeyECResponse) GetLinks() AgentSecretLinks {
	if o == nil || IsNil(o.Links) {
		var ret AgentSecretLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentJsonWebKeyECResponse) GetLinksOk() (*AgentSecretLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentJsonWebKeyECResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AgentSecretLinks and assigns it to the Links field.
func (o *AgentJsonWebKeyECResponse) SetLinks(v AgentSecretLinks) {
	o.Links = &v
}

func (o AgentJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentJsonWebKeyECResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Crv) {
		toSerialize["crv"] = o.Crv
	}
	toSerialize["kty"] = o.Kty
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
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

func (o *AgentJsonWebKeyECResponse) UnmarshalJSON(data []byte) (err error) {
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

	varAgentJsonWebKeyECResponse := _AgentJsonWebKeyECResponse{}

	err = json.Unmarshal(data, &varAgentJsonWebKeyECResponse)

	if err != nil {
		return err
	}

	*o = AgentJsonWebKeyECResponse(varAgentJsonWebKeyECResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "crv")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
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

type NullableAgentJsonWebKeyECResponse struct {
	value *AgentJsonWebKeyECResponse
	isSet bool
}

func (v NullableAgentJsonWebKeyECResponse) Get() *AgentJsonWebKeyECResponse {
	return v.value
}

func (v *NullableAgentJsonWebKeyECResponse) Set(val *AgentJsonWebKeyECResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentJsonWebKeyECResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentJsonWebKeyECResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentJsonWebKeyECResponse(val *AgentJsonWebKeyECResponse) *NullableAgentJsonWebKeyECResponse {
	return &NullableAgentJsonWebKeyECResponse{value: val, isSet: true}
}

func (v NullableAgentJsonWebKeyECResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentJsonWebKeyECResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
