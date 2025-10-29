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
)

// checks if the FederatedClaimRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FederatedClaimRequestBody{}

// FederatedClaimRequestBody struct for FederatedClaimRequestBody
type FederatedClaimRequestBody struct {
	// The Okta Expression Language expression to be evaluated at runtime
	Expression *string `json:"expression,omitempty"`
	// The name of the claim to be used in the produced token
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FederatedClaimRequestBody FederatedClaimRequestBody

// NewFederatedClaimRequestBody instantiates a new FederatedClaimRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFederatedClaimRequestBody() *FederatedClaimRequestBody {
	this := FederatedClaimRequestBody{}
	return &this
}

// NewFederatedClaimRequestBodyWithDefaults instantiates a new FederatedClaimRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFederatedClaimRequestBodyWithDefaults() *FederatedClaimRequestBody {
	this := FederatedClaimRequestBody{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *FederatedClaimRequestBody) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaimRequestBody) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *FederatedClaimRequestBody) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *FederatedClaimRequestBody) SetExpression(v string) {
	o.Expression = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FederatedClaimRequestBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FederatedClaimRequestBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FederatedClaimRequestBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FederatedClaimRequestBody) SetName(v string) {
	o.Name = &v
}

func (o FederatedClaimRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FederatedClaimRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FederatedClaimRequestBody) UnmarshalJSON(data []byte) (err error) {
	varFederatedClaimRequestBody := _FederatedClaimRequestBody{}

	err = json.Unmarshal(data, &varFederatedClaimRequestBody)

	if err != nil {
		return err
	}

	*o = FederatedClaimRequestBody(varFederatedClaimRequestBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "expression")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFederatedClaimRequestBody struct {
	value *FederatedClaimRequestBody
	isSet bool
}

func (v NullableFederatedClaimRequestBody) Get() *FederatedClaimRequestBody {
	return v.value
}

func (v *NullableFederatedClaimRequestBody) Set(val *FederatedClaimRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFederatedClaimRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFederatedClaimRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFederatedClaimRequestBody(val *FederatedClaimRequestBody) *NullableFederatedClaimRequestBody {
	return &NullableFederatedClaimRequestBody{value: val, isSet: true}
}

func (v NullableFederatedClaimRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFederatedClaimRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
