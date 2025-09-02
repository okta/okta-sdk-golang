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

// FederatedClaim struct for FederatedClaim
type FederatedClaim struct {
	// Timestamp when the federated claim was created
	Created *string `json:"created,omitempty"`
	// The Okta Expression Language expression to be evaluated at runtime
	Expression *string `json:"expression,omitempty"`
	// The unique ID of the federated claim
	Id *string `json:"id,omitempty"`
	// Timestamp when the federated claim was updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// The name of the claim to be used in the produced token
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederatedClaim FederatedClaim

// NewFederatedClaim instantiates a new FederatedClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederatedClaim() *FederatedClaim {
	this := FederatedClaim{}
	return &this
}

// NewFederatedClaimWithDefaults instantiates a new FederatedClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederatedClaimWithDefaults() *FederatedClaim {
	this := FederatedClaim{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *FederatedClaim) GetCreated() string {
	if o == nil || o.Created == nil {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaim) GetCreatedOk() (*string, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *FederatedClaim) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *FederatedClaim) SetCreated(v string) {
	o.Created = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *FederatedClaim) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaim) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *FederatedClaim) HasExpression() bool {
	if o != nil && o.Expression != nil {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *FederatedClaim) SetExpression(v string) {
	o.Expression = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FederatedClaim) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaim) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FederatedClaim) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FederatedClaim) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *FederatedClaim) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaim) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *FederatedClaim) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *FederatedClaim) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FederatedClaim) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaim) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FederatedClaim) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FederatedClaim) SetName(v string) {
	o.Name = &v
}

func (o FederatedClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FederatedClaim) UnmarshalJSON(bytes []byte) (err error) {
	varFederatedClaim := _FederatedClaim{}

	err = json.Unmarshal(bytes, &varFederatedClaim)
	if err == nil {
		*o = FederatedClaim(varFederatedClaim)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "expression")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFederatedClaim struct {
	value *FederatedClaim
	isSet bool
}

func (v NullableFederatedClaim) Get() *FederatedClaim {
	return v.value
}

func (v *NullableFederatedClaim) Set(val *FederatedClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableFederatedClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableFederatedClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederatedClaim(val *FederatedClaim) *NullableFederatedClaim {
	return &NullableFederatedClaim{value: val, isSet: true}
}

func (v NullableFederatedClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederatedClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

