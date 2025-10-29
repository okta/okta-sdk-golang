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

API version: 2025.10.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the PageRootLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRootLinks{}

// PageRootLinks struct for PageRootLinks
type PageRootLinks struct {
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Default              *HrefObject         `json:"default,omitempty"`
	Customized           *HrefObject         `json:"customized,omitempty"`
	Preview              *HrefObject         `json:"preview,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PageRootLinks PageRootLinks

// NewPageRootLinks instantiates a new PageRootLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRootLinks() *PageRootLinks {
	this := PageRootLinks{}
	return &this
}

// NewPageRootLinksWithDefaults instantiates a new PageRootLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRootLinksWithDefaults() *PageRootLinks {
	this := PageRootLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PageRootLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PageRootLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *PageRootLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PageRootLinks) GetDefault() HrefObject {
	if o == nil || IsNil(o.Default) {
		var ret HrefObject
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootLinks) GetDefaultOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PageRootLinks) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given HrefObject and assigns it to the Default field.
func (o *PageRootLinks) SetDefault(v HrefObject) {
	o.Default = &v
}

// GetCustomized returns the Customized field value if set, zero value otherwise.
func (o *PageRootLinks) GetCustomized() HrefObject {
	if o == nil || IsNil(o.Customized) {
		var ret HrefObject
		return ret
	}
	return *o.Customized
}

// GetCustomizedOk returns a tuple with the Customized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootLinks) GetCustomizedOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Customized) {
		return nil, false
	}
	return o.Customized, true
}

// HasCustomized returns a boolean if a field has been set.
func (o *PageRootLinks) HasCustomized() bool {
	if o != nil && !IsNil(o.Customized) {
		return true
	}

	return false
}

// SetCustomized gets a reference to the given HrefObject and assigns it to the Customized field.
func (o *PageRootLinks) SetCustomized(v HrefObject) {
	o.Customized = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *PageRootLinks) GetPreview() HrefObject {
	if o == nil || IsNil(o.Preview) {
		var ret HrefObject
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRootLinks) GetPreviewOk() (*HrefObject, bool) {
	if o == nil || IsNil(o.Preview) {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *PageRootLinks) HasPreview() bool {
	if o != nil && !IsNil(o.Preview) {
		return true
	}

	return false
}

// SetPreview gets a reference to the given HrefObject and assigns it to the Preview field.
func (o *PageRootLinks) SetPreview(v HrefObject) {
	o.Preview = &v
}

func (o PageRootLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRootLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Customized) {
		toSerialize["customized"] = o.Customized
	}
	if !IsNil(o.Preview) {
		toSerialize["preview"] = o.Preview
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PageRootLinks) UnmarshalJSON(data []byte) (err error) {
	varPageRootLinks := _PageRootLinks{}

	err = json.Unmarshal(data, &varPageRootLinks)

	if err != nil {
		return err
	}

	*o = PageRootLinks(varPageRootLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "default")
		delete(additionalProperties, "customized")
		delete(additionalProperties, "preview")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePageRootLinks struct {
	value *PageRootLinks
	isSet bool
}

func (v NullablePageRootLinks) Get() *PageRootLinks {
	return v.value
}

func (v *NullablePageRootLinks) Set(val *PageRootLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRootLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRootLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRootLinks(val *PageRootLinks) *NullablePageRootLinks {
	return &NullablePageRootLinks{value: val, isSet: true}
}

func (v NullablePageRootLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRootLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
