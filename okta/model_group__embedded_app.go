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

// checks if the GroupEmbeddedApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupEmbeddedApp{}

// GroupEmbeddedApp If the group is sourced from an app, this object contains information about that app
type GroupEmbeddedApp struct {
	// The ID of the `AppInstance`
	Id *string `json:"id,omitempty"`
	// The name of the `AppInstance`
	Name *string `json:"name,omitempty"`
	// The user-facing display name of the `AppInstance`
	Label *string `json:"label,omitempty"`
	// The configured sign-on mode for the `AppInstance`
	SignOnMode           *string `json:"signOnMode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupEmbeddedApp GroupEmbeddedApp

// NewGroupEmbeddedApp instantiates a new GroupEmbeddedApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupEmbeddedApp() *GroupEmbeddedApp {
	this := GroupEmbeddedApp{}
	return &this
}

// NewGroupEmbeddedAppWithDefaults instantiates a new GroupEmbeddedApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupEmbeddedAppWithDefaults() *GroupEmbeddedApp {
	this := GroupEmbeddedApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupEmbeddedApp) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedApp) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupEmbeddedApp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupEmbeddedApp) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupEmbeddedApp) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedApp) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupEmbeddedApp) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupEmbeddedApp) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GroupEmbeddedApp) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedApp) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GroupEmbeddedApp) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GroupEmbeddedApp) SetLabel(v string) {
	o.Label = &v
}

// GetSignOnMode returns the SignOnMode field value if set, zero value otherwise.
func (o *GroupEmbeddedApp) GetSignOnMode() string {
	if o == nil || IsNil(o.SignOnMode) {
		var ret string
		return ret
	}
	return *o.SignOnMode
}

// GetSignOnModeOk returns a tuple with the SignOnMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupEmbeddedApp) GetSignOnModeOk() (*string, bool) {
	if o == nil || IsNil(o.SignOnMode) {
		return nil, false
	}
	return o.SignOnMode, true
}

// HasSignOnMode returns a boolean if a field has been set.
func (o *GroupEmbeddedApp) HasSignOnMode() bool {
	if o != nil && !IsNil(o.SignOnMode) {
		return true
	}

	return false
}

// SetSignOnMode gets a reference to the given string and assigns it to the SignOnMode field.
func (o *GroupEmbeddedApp) SetSignOnMode(v string) {
	o.SignOnMode = &v
}

func (o GroupEmbeddedApp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupEmbeddedApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.SignOnMode) {
		toSerialize["signOnMode"] = o.SignOnMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupEmbeddedApp) UnmarshalJSON(data []byte) (err error) {
	varGroupEmbeddedApp := _GroupEmbeddedApp{}

	err = json.Unmarshal(data, &varGroupEmbeddedApp)

	if err != nil {
		return err
	}

	*o = GroupEmbeddedApp(varGroupEmbeddedApp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "signOnMode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupEmbeddedApp struct {
	value *GroupEmbeddedApp
	isSet bool
}

func (v NullableGroupEmbeddedApp) Get() *GroupEmbeddedApp {
	return v.value
}

func (v *NullableGroupEmbeddedApp) Set(val *GroupEmbeddedApp) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupEmbeddedApp) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupEmbeddedApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupEmbeddedApp(val *GroupEmbeddedApp) *NullableGroupEmbeddedApp {
	return &NullableGroupEmbeddedApp{value: val, isSet: true}
}

func (v NullableGroupEmbeddedApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupEmbeddedApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
