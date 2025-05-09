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

// PageRootEmbedded struct for PageRootEmbedded
type PageRootEmbedded struct {
	Default *CustomizablePage `json:"default,omitempty"`
	Customized *CustomizablePage `json:"customized,omitempty"`
	CustomizedUrl *string `json:"customizedUrl,omitempty"`
	Preview *CustomizablePage `json:"preview,omitempty"`
	PreviewUrl *string `json:"previewUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PageRootEmbedded PageRootEmbedded

// NewPageRootEmbedded instantiates a new PageRootEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRootEmbedded() *PageRootEmbedded {
	this := PageRootEmbedded{}
	return &this
}

// NewPageRootEmbeddedWithDefaults instantiates a new PageRootEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRootEmbeddedWithDefaults() *PageRootEmbedded {
	this := PageRootEmbedded{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PageRootEmbedded) GetDefault() CustomizablePage {
	if o == nil || o.Default == nil {
		var ret CustomizablePage
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootEmbedded) GetDefaultOk() (*CustomizablePage, bool) {
	if o == nil || o.Default == nil {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PageRootEmbedded) HasDefault() bool {
	if o != nil && o.Default != nil {
		return true
	}

	return false
}

// SetDefault gets a reference to the given CustomizablePage and assigns it to the Default field.
func (o *PageRootEmbedded) SetDefault(v CustomizablePage) {
	o.Default = &v
}

// GetCustomized returns the Customized field value if set, zero value otherwise.
func (o *PageRootEmbedded) GetCustomized() CustomizablePage {
	if o == nil || o.Customized == nil {
		var ret CustomizablePage
		return ret
	}
	return *o.Customized
}

// GetCustomizedOk returns a tuple with the Customized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootEmbedded) GetCustomizedOk() (*CustomizablePage, bool) {
	if o == nil || o.Customized == nil {
		return nil, false
	}
	return o.Customized, true
}

// HasCustomized returns a boolean if a field has been set.
func (o *PageRootEmbedded) HasCustomized() bool {
	if o != nil && o.Customized != nil {
		return true
	}

	return false
}

// SetCustomized gets a reference to the given CustomizablePage and assigns it to the Customized field.
func (o *PageRootEmbedded) SetCustomized(v CustomizablePage) {
	o.Customized = &v
}

// GetCustomizedUrl returns the CustomizedUrl field value if set, zero value otherwise.
func (o *PageRootEmbedded) GetCustomizedUrl() string {
	if o == nil || o.CustomizedUrl == nil {
		var ret string
		return ret
	}
	return *o.CustomizedUrl
}

// GetCustomizedUrlOk returns a tuple with the CustomizedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootEmbedded) GetCustomizedUrlOk() (*string, bool) {
	if o == nil || o.CustomizedUrl == nil {
		return nil, false
	}
	return o.CustomizedUrl, true
}

// HasCustomizedUrl returns a boolean if a field has been set.
func (o *PageRootEmbedded) HasCustomizedUrl() bool {
	if o != nil && o.CustomizedUrl != nil {
		return true
	}

	return false
}

// SetCustomizedUrl gets a reference to the given string and assigns it to the CustomizedUrl field.
func (o *PageRootEmbedded) SetCustomizedUrl(v string) {
	o.CustomizedUrl = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *PageRootEmbedded) GetPreview() CustomizablePage {
	if o == nil || o.Preview == nil {
		var ret CustomizablePage
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootEmbedded) GetPreviewOk() (*CustomizablePage, bool) {
	if o == nil || o.Preview == nil {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *PageRootEmbedded) HasPreview() bool {
	if o != nil && o.Preview != nil {
		return true
	}

	return false
}

// SetPreview gets a reference to the given CustomizablePage and assigns it to the Preview field.
func (o *PageRootEmbedded) SetPreview(v CustomizablePage) {
	o.Preview = &v
}

// GetPreviewUrl returns the PreviewUrl field value if set, zero value otherwise.
func (o *PageRootEmbedded) GetPreviewUrl() string {
	if o == nil || o.PreviewUrl == nil {
		var ret string
		return ret
	}
	return *o.PreviewUrl
}

// GetPreviewUrlOk returns a tuple with the PreviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootEmbedded) GetPreviewUrlOk() (*string, bool) {
	if o == nil || o.PreviewUrl == nil {
		return nil, false
	}
	return o.PreviewUrl, true
}

// HasPreviewUrl returns a boolean if a field has been set.
func (o *PageRootEmbedded) HasPreviewUrl() bool {
	if o != nil && o.PreviewUrl != nil {
		return true
	}

	return false
}

// SetPreviewUrl gets a reference to the given string and assigns it to the PreviewUrl field.
func (o *PageRootEmbedded) SetPreviewUrl(v string) {
	o.PreviewUrl = &v
}

func (o PageRootEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Default != nil {
		toSerialize["default"] = o.Default
	}
	if o.Customized != nil {
		toSerialize["customized"] = o.Customized
	}
	if o.CustomizedUrl != nil {
		toSerialize["customizedUrl"] = o.CustomizedUrl
	}
	if o.Preview != nil {
		toSerialize["preview"] = o.Preview
	}
	if o.PreviewUrl != nil {
		toSerialize["previewUrl"] = o.PreviewUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PageRootEmbedded) UnmarshalJSON(bytes []byte) (err error) {
	varPageRootEmbedded := _PageRootEmbedded{}

	err = json.Unmarshal(bytes, &varPageRootEmbedded)
	if err == nil {
		*o = PageRootEmbedded(varPageRootEmbedded)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "default")
		delete(additionalProperties, "customized")
		delete(additionalProperties, "customizedUrl")
		delete(additionalProperties, "preview")
		delete(additionalProperties, "previewUrl")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePageRootEmbedded struct {
	value *PageRootEmbedded
	isSet bool
}

func (v NullablePageRootEmbedded) Get() *PageRootEmbedded {
	return v.value
}

func (v *NullablePageRootEmbedded) Set(val *PageRootEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRootEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRootEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRootEmbedded(val *PageRootEmbedded) *NullablePageRootEmbedded {
	return &NullablePageRootEmbedded{value: val, isSet: true}
}

func (v NullablePageRootEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRootEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

