/*
Okta Admin Management API

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

// checks if the SignInPageRootAllOfEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignInPageRootAllOfEmbedded{}

// SignInPageRootAllOfEmbedded struct for SignInPageRootAllOfEmbedded
type SignInPageRootAllOfEmbedded struct {
	// <x-lifecycle-container><x-lifecycle class=\"oie\"></x-lifecycle></x-lifecycle-container>The schema used to validate `widgetConfiguration`. Returned as a JSON-serialized string (a JSON Schema document encoded as a string), not a JSON object, so clients must parse it to obtain the schema.
	WidgetConfigurationSchema *string `json:"widgetConfigurationSchema,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _SignInPageRootAllOfEmbedded SignInPageRootAllOfEmbedded

// NewSignInPageRootAllOfEmbedded instantiates a new SignInPageRootAllOfEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignInPageRootAllOfEmbedded() *SignInPageRootAllOfEmbedded {
	this := SignInPageRootAllOfEmbedded{}
	return &this
}

// NewSignInPageRootAllOfEmbeddedWithDefaults instantiates a new SignInPageRootAllOfEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignInPageRootAllOfEmbeddedWithDefaults() *SignInPageRootAllOfEmbedded {
	this := SignInPageRootAllOfEmbedded{}
	return &this
}

// GetWidgetConfigurationSchema returns the WidgetConfigurationSchema field value if set, zero value otherwise.
func (o *SignInPageRootAllOfEmbedded) GetWidgetConfigurationSchema() string {
	if o == nil || IsNil(o.WidgetConfigurationSchema) {
		var ret string
		return ret
	}
	return *o.WidgetConfigurationSchema
}

// GetWidgetConfigurationSchemaOk returns a tuple with the WidgetConfigurationSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignInPageRootAllOfEmbedded) GetWidgetConfigurationSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.WidgetConfigurationSchema) {
		return nil, false
	}
	return o.WidgetConfigurationSchema, true
}

// HasWidgetConfigurationSchema returns a boolean if a field has been set.
func (o *SignInPageRootAllOfEmbedded) HasWidgetConfigurationSchema() bool {
	if o != nil && !IsNil(o.WidgetConfigurationSchema) {
		return true
	}

	return false
}

// SetWidgetConfigurationSchema gets a reference to the given string and assigns it to the WidgetConfigurationSchema field.
func (o *SignInPageRootAllOfEmbedded) SetWidgetConfigurationSchema(v string) {
	o.WidgetConfigurationSchema = &v
}

func (o SignInPageRootAllOfEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignInPageRootAllOfEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WidgetConfigurationSchema) {
		toSerialize["widgetConfigurationSchema"] = o.WidgetConfigurationSchema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignInPageRootAllOfEmbedded) UnmarshalJSON(data []byte) (err error) {
	varSignInPageRootAllOfEmbedded := _SignInPageRootAllOfEmbedded{}

	err = json.Unmarshal(data, &varSignInPageRootAllOfEmbedded)

	if err != nil {
		return err
	}

	*o = SignInPageRootAllOfEmbedded(varSignInPageRootAllOfEmbedded)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "widgetConfigurationSchema")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignInPageRootAllOfEmbedded struct {
	value *SignInPageRootAllOfEmbedded
	isSet bool
}

func (v NullableSignInPageRootAllOfEmbedded) Get() *SignInPageRootAllOfEmbedded {
	return v.value
}

func (v *NullableSignInPageRootAllOfEmbedded) Set(val *SignInPageRootAllOfEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableSignInPageRootAllOfEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableSignInPageRootAllOfEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignInPageRootAllOfEmbedded(val *SignInPageRootAllOfEmbedded) *NullableSignInPageRootAllOfEmbedded {
	return &NullableSignInPageRootAllOfEmbedded{value: val, isSet: true}
}

func (v NullableSignInPageRootAllOfEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignInPageRootAllOfEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
