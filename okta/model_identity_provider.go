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
	"time"
)

// IdentityProvider struct for IdentityProvider
type IdentityProvider struct {
	Created NullableTime `json:"created,omitempty"`
	Id *string `json:"id,omitempty"`
	IssuerMode *string `json:"issuerMode,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	Policy *IdentityProviderPolicy `json:"policy,omitempty"`
	Properties NullableIdentityProviderProperties `json:"properties,omitempty"`
	Protocol *Protocol `json:"protocol,omitempty"`
	Status *string `json:"status,omitempty"`
	Type *string `json:"type,omitempty"`
	Links *IdentityProviderLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProvider IdentityProvider

// NewIdentityProvider instantiates a new IdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProvider() *IdentityProvider {
	this := IdentityProvider{}
	return &this
}

// NewIdentityProviderWithDefaults instantiates a new IdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderWithDefaults() *IdentityProvider {
	this := IdentityProvider{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProvider) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProvider) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *IdentityProvider) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *IdentityProvider) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *IdentityProvider) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *IdentityProvider) UnsetCreated() {
	o.Created.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProvider) SetId(v string) {
	o.Id = &v
}

// GetIssuerMode returns the IssuerMode field value if set, zero value otherwise.
func (o *IdentityProvider) GetIssuerMode() string {
	if o == nil || o.IssuerMode == nil {
		var ret string
		return ret
	}
	return *o.IssuerMode
}

// GetIssuerModeOk returns a tuple with the IssuerMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetIssuerModeOk() (*string, bool) {
	if o == nil || o.IssuerMode == nil {
		return nil, false
	}
	return o.IssuerMode, true
}

// HasIssuerMode returns a boolean if a field has been set.
func (o *IdentityProvider) HasIssuerMode() bool {
	if o != nil && o.IssuerMode != nil {
		return true
	}

	return false
}

// SetIssuerMode gets a reference to the given string and assigns it to the IssuerMode field.
func (o *IdentityProvider) SetIssuerMode(v string) {
	o.IssuerMode = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *IdentityProvider) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *IdentityProvider) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *IdentityProvider) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IdentityProvider) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IdentityProvider) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IdentityProvider) SetName(v string) {
	o.Name = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *IdentityProvider) GetPolicy() IdentityProviderPolicy {
	if o == nil || o.Policy == nil {
		var ret IdentityProviderPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetPolicyOk() (*IdentityProviderPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *IdentityProvider) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given IdentityProviderPolicy and assigns it to the Policy field.
func (o *IdentityProvider) SetPolicy(v IdentityProviderPolicy) {
	o.Policy = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityProvider) GetProperties() IdentityProviderProperties {
	if o == nil || o.Properties.Get() == nil {
		var ret IdentityProviderProperties
		return ret
	}
	return *o.Properties.Get()
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityProvider) GetPropertiesOk() (*IdentityProviderProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Get(), o.Properties.IsSet()
}

// HasProperties returns a boolean if a field has been set.
func (o *IdentityProvider) HasProperties() bool {
	if o != nil && o.Properties.IsSet() {
		return true
	}

	return false
}

// SetProperties gets a reference to the given NullableIdentityProviderProperties and assigns it to the Properties field.
func (o *IdentityProvider) SetProperties(v IdentityProviderProperties) {
	o.Properties.Set(&v)
}
// SetPropertiesNil sets the value for Properties to be an explicit nil
func (o *IdentityProvider) SetPropertiesNil() {
	o.Properties.Set(nil)
}

// UnsetProperties ensures that no value is present for Properties, not even an explicit nil
func (o *IdentityProvider) UnsetProperties() {
	o.Properties.Unset()
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *IdentityProvider) GetProtocol() Protocol {
	if o == nil || o.Protocol == nil {
		var ret Protocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetProtocolOk() (*Protocol, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *IdentityProvider) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given Protocol and assigns it to the Protocol field.
func (o *IdentityProvider) SetProtocol(v Protocol) {
	o.Protocol = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IdentityProvider) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IdentityProvider) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IdentityProvider) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityProvider) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityProvider) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityProvider) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentityProvider) GetLinks() IdentityProviderLinks {
	if o == nil || o.Links == nil {
		var ret IdentityProviderLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProvider) GetLinksOk() (*IdentityProviderLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentityProvider) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given IdentityProviderLinks and assigns it to the Links field.
func (o *IdentityProvider) SetLinks(v IdentityProviderLinks) {
	o.Links = &v
}

func (o IdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IssuerMode != nil {
		toSerialize["issuerMode"] = o.IssuerMode
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	if o.Properties.IsSet() {
		toSerialize["properties"] = o.Properties.Get()
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProvider) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProvider := _IdentityProvider{}

	err = json.Unmarshal(bytes, &varIdentityProvider)
	if err == nil {
		*o = IdentityProvider(varIdentityProvider)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "issuerMode")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "policy")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProvider struct {
	value *IdentityProvider
	isSet bool
}

func (v NullableIdentityProvider) Get() *IdentityProvider {
	return v.value
}

func (v *NullableIdentityProvider) Set(val *IdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProvider(val *IdentityProvider) *NullableIdentityProvider {
	return &NullableIdentityProvider{value: val, isSet: true}
}

func (v NullableIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

