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
)

// checks if the UserImportRequestDataContextApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportRequestDataContextApplication{}

// UserImportRequestDataContextApplication Details of the app from which the user is being imported
type UserImportRequestDataContextApplication struct {
	// The app name
	Name *string `json:"name,omitempty"`
	// The app ID
	Id *string `json:"id,omitempty"`
	// The user-defined display name for the app
	Label *string `json:"label,omitempty"`
	// The status of the app
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserImportRequestDataContextApplication UserImportRequestDataContextApplication

// NewUserImportRequestDataContextApplication instantiates a new UserImportRequestDataContextApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportRequestDataContextApplication() *UserImportRequestDataContextApplication {
	this := UserImportRequestDataContextApplication{}
	return &this
}

// NewUserImportRequestDataContextApplicationWithDefaults instantiates a new UserImportRequestDataContextApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportRequestDataContextApplicationWithDefaults() *UserImportRequestDataContextApplication {
	this := UserImportRequestDataContextApplication{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserImportRequestDataContextApplication) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContextApplication) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserImportRequestDataContextApplication) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserImportRequestDataContextApplication) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserImportRequestDataContextApplication) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContextApplication) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserImportRequestDataContextApplication) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserImportRequestDataContextApplication) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *UserImportRequestDataContextApplication) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContextApplication) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *UserImportRequestDataContextApplication) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *UserImportRequestDataContextApplication) SetLabel(v string) {
	o.Label = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserImportRequestDataContextApplication) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportRequestDataContextApplication) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserImportRequestDataContextApplication) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserImportRequestDataContextApplication) SetStatus(v string) {
	o.Status = &v
}

func (o UserImportRequestDataContextApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportRequestDataContextApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserImportRequestDataContextApplication) UnmarshalJSON(data []byte) (err error) {
	varUserImportRequestDataContextApplication := _UserImportRequestDataContextApplication{}

	err = json.Unmarshal(data, &varUserImportRequestDataContextApplication)

	if err != nil {
		return err
	}

	*o = UserImportRequestDataContextApplication(varUserImportRequestDataContextApplication)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUserImportRequestDataContextApplication struct {
	value *UserImportRequestDataContextApplication
	isSet bool
}

func (v NullableUserImportRequestDataContextApplication) Get() *UserImportRequestDataContextApplication {
	return v.value
}

func (v *NullableUserImportRequestDataContextApplication) Set(val *UserImportRequestDataContextApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportRequestDataContextApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportRequestDataContextApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportRequestDataContextApplication(val *UserImportRequestDataContextApplication) *NullableUserImportRequestDataContextApplication {
	return &NullableUserImportRequestDataContextApplication{value: val, isSet: true}
}

func (v NullableUserImportRequestDataContextApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportRequestDataContextApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
