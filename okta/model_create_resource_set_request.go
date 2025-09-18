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
	"fmt"
)

// checks if the CreateResourceSetRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateResourceSetRequest{}

// CreateResourceSetRequest struct for CreateResourceSetRequest
type CreateResourceSetRequest struct {
	// Description of the resource set
	Description string `json:"description"`
	// Unique name for the resource set
	Label string `json:"label"`
	// The endpoint (URL) that references all resource objects included in the resource set. Resources are identified by either an Okta Resource Name (ORN) or by a REST URL format. See [Okta Resource Name](/openapi/okta-management/guides/roles/#okta-resource-name-orn).
	Resources            []string `json:"resources"`
	AdditionalProperties map[string]interface{}
}

type _CreateResourceSetRequest CreateResourceSetRequest

// NewCreateResourceSetRequest instantiates a new CreateResourceSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourceSetRequest(description string, label string, resources []string) *CreateResourceSetRequest {
	this := CreateResourceSetRequest{}
	this.Description = description
	this.Label = label
	this.Resources = resources
	return &this
}

// NewCreateResourceSetRequestWithDefaults instantiates a new CreateResourceSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourceSetRequestWithDefaults() *CreateResourceSetRequest {
	this := CreateResourceSetRequest{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateResourceSetRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateResourceSetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateResourceSetRequest) SetDescription(v string) {
	o.Description = v
}

// GetLabel returns the Label field value
func (o *CreateResourceSetRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CreateResourceSetRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CreateResourceSetRequest) SetLabel(v string) {
	o.Label = v
}

// GetResources returns the Resources field value
func (o *CreateResourceSetRequest) GetResources() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
func (o *CreateResourceSetRequest) GetResourcesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *CreateResourceSetRequest) SetResources(v []string) {
	o.Resources = v
}

func (o CreateResourceSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateResourceSetRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["label"] = o.Label
	toSerialize["resources"] = o.Resources

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateResourceSetRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"label",
		"resources",
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

	varCreateResourceSetRequest := _CreateResourceSetRequest{}

	err = json.Unmarshal(data, &varCreateResourceSetRequest)

	if err != nil {
		return err
	}

	*o = CreateResourceSetRequest(varCreateResourceSetRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateResourceSetRequest struct {
	value *CreateResourceSetRequest
	isSet bool
}

func (v NullableCreateResourceSetRequest) Get() *CreateResourceSetRequest {
	return v.value
}

func (v *NullableCreateResourceSetRequest) Set(val *CreateResourceSetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateResourceSetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateResourceSetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateResourceSetRequest(val *CreateResourceSetRequest) *NullableCreateResourceSetRequest {
	return &NullableCreateResourceSetRequest{value: val, isSet: true}
}

func (v NullableCreateResourceSetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateResourceSetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
