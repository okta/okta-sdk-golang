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

// CreateResourceSetRequest struct for CreateResourceSetRequest
type CreateResourceSetRequest struct {
	// Description of the Resource Set
	Description *string `json:"description,omitempty"`
	// Unique name for the Resource Set
	Label *string `json:"label,omitempty"`
	// The endpoint (URL) that references all resource objects included in the Resource Set. Resources are identified by either an Okta Resource Name (ORN) or by a REST URL format. See [Okta Resource Name](/openapi/okta-management/guides/roles/#okta-resource-name-orn).
	Resources []string `json:"resources,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateResourceSetRequest CreateResourceSetRequest

// NewCreateResourceSetRequest instantiates a new CreateResourceSetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateResourceSetRequest() *CreateResourceSetRequest {
	this := CreateResourceSetRequest{}
	return &this
}

// NewCreateResourceSetRequestWithDefaults instantiates a new CreateResourceSetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateResourceSetRequestWithDefaults() *CreateResourceSetRequest {
	this := CreateResourceSetRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateResourceSetRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceSetRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateResourceSetRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateResourceSetRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateResourceSetRequest) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceSetRequest) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateResourceSetRequest) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateResourceSetRequest) SetLabel(v string) {
	o.Label = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CreateResourceSetRequest) GetResources() []string {
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateResourceSetRequest) GetResourcesOk() ([]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *CreateResourceSetRequest) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *CreateResourceSetRequest) SetResources(v []string) {
	o.Resources = v
}

func (o CreateResourceSetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateResourceSetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateResourceSetRequest := _CreateResourceSetRequest{}

	err = json.Unmarshal(bytes, &varCreateResourceSetRequest)
	if err == nil {
		*o = CreateResourceSetRequest(varCreateResourceSetRequest)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		delete(additionalProperties, "resources")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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

