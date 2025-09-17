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

// checks if the CustomizablePage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomizablePage{}

// CustomizablePage struct for CustomizablePage
type CustomizablePage struct {
	// The HTML for the page
	PageContent          *string `json:"pageContent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomizablePage CustomizablePage

// NewCustomizablePage instantiates a new CustomizablePage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomizablePage() *CustomizablePage {
	this := CustomizablePage{}
	return &this
}

// NewCustomizablePageWithDefaults instantiates a new CustomizablePage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomizablePageWithDefaults() *CustomizablePage {
	this := CustomizablePage{}
	return &this
}

// GetPageContent returns the PageContent field value if set, zero value otherwise.
func (o *CustomizablePage) GetPageContent() string {
	if o == nil || IsNil(o.PageContent) {
		var ret string
		return ret
	}
	return *o.PageContent
}

// GetPageContentOk returns a tuple with the PageContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomizablePage) GetPageContentOk() (*string, bool) {
	if o == nil || IsNil(o.PageContent) {
		return nil, false
	}
	return o.PageContent, true
}

// HasPageContent returns a boolean if a field has been set.
func (o *CustomizablePage) HasPageContent() bool {
	if o != nil && !IsNil(o.PageContent) {
		return true
	}

	return false
}

// SetPageContent gets a reference to the given string and assigns it to the PageContent field.
func (o *CustomizablePage) SetPageContent(v string) {
	o.PageContent = &v
}

func (o CustomizablePage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomizablePage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageContent) {
		toSerialize["pageContent"] = o.PageContent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomizablePage) UnmarshalJSON(data []byte) (err error) {
	varCustomizablePage := _CustomizablePage{}

	err = json.Unmarshal(data, &varCustomizablePage)

	if err != nil {
		return err
	}

	*o = CustomizablePage(varCustomizablePage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pageContent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomizablePage struct {
	value *CustomizablePage
	isSet bool
}

func (v NullableCustomizablePage) Get() *CustomizablePage {
	return v.value
}

func (v *NullableCustomizablePage) Set(val *CustomizablePage) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomizablePage) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomizablePage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomizablePage(val *CustomizablePage) *NullableCustomizablePage {
	return &NullableCustomizablePage{value: val, isSet: true}
}

func (v NullableCustomizablePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomizablePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
