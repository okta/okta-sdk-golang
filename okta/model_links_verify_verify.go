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

// LinksVerifyVerify struct for LinksVerifyVerify
type LinksVerifyVerify struct {
	Hints *HrefHints `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	// Link name
	Name *string `json:"name,omitempty"`
	// Indicates whether the Link Object's `href` property is a URI template.
	Templated *bool `json:"templated,omitempty"`
	// The media type of the link. If omitted, it is implicitly `application/json`.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksVerifyVerify LinksVerifyVerify

// NewLinksVerifyVerify instantiates a new LinksVerifyVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksVerifyVerify(href string) *LinksVerifyVerify {
	this := LinksVerifyVerify{}
	this.Href = href
	return &this
}

// NewLinksVerifyVerifyWithDefaults instantiates a new LinksVerifyVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksVerifyVerifyWithDefaults() *LinksVerifyVerify {
	this := LinksVerifyVerify{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *LinksVerifyVerify) GetHints() HrefHints {
	if o == nil || o.Hints == nil {
		var ret HrefHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksVerifyVerify) GetHintsOk() (*HrefHints, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *LinksVerifyVerify) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given HrefHints and assigns it to the Hints field.
func (o *LinksVerifyVerify) SetHints(v HrefHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *LinksVerifyVerify) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *LinksVerifyVerify) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *LinksVerifyVerify) SetHref(v string) {
	o.Href = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LinksVerifyVerify) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksVerifyVerify) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LinksVerifyVerify) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LinksVerifyVerify) SetName(v string) {
	o.Name = &v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *LinksVerifyVerify) GetTemplated() bool {
	if o == nil || o.Templated == nil {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksVerifyVerify) GetTemplatedOk() (*bool, bool) {
	if o == nil || o.Templated == nil {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *LinksVerifyVerify) HasTemplated() bool {
	if o != nil && o.Templated != nil {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *LinksVerifyVerify) SetTemplated(v bool) {
	o.Templated = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LinksVerifyVerify) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksVerifyVerify) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LinksVerifyVerify) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LinksVerifyVerify) SetType(v string) {
	o.Type = &v
}

func (o LinksVerifyVerify) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hints != nil {
		toSerialize["hints"] = o.Hints
	}
	if true {
		toSerialize["href"] = o.Href
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Templated != nil {
		toSerialize["templated"] = o.Templated
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksVerifyVerify) UnmarshalJSON(bytes []byte) (err error) {
	varLinksVerifyVerify := _LinksVerifyVerify{}

	err = json.Unmarshal(bytes, &varLinksVerifyVerify)
	if err == nil {
		*o = LinksVerifyVerify(varLinksVerifyVerify)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		delete(additionalProperties, "name")
		delete(additionalProperties, "templated")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksVerifyVerify struct {
	value *LinksVerifyVerify
	isSet bool
}

func (v NullableLinksVerifyVerify) Get() *LinksVerifyVerify {
	return v.value
}

func (v *NullableLinksVerifyVerify) Set(val *LinksVerifyVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksVerifyVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksVerifyVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksVerifyVerify(val *LinksVerifyVerify) *NullableLinksVerifyVerify {
	return &NullableLinksVerifyVerify{value: val, isSet: true}
}

func (v NullableLinksVerifyVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksVerifyVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

