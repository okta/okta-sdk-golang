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

API version: 5.1.0
Contact: devex-public@okta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
package okta

import (
	"encoding/json"
)

// DomainLinksAllOfBrand struct for DomainLinksAllOfBrand
type DomainLinksAllOfBrand struct {
	Hints *HrefObjectHints `json:"hints,omitempty"`
	// Link URI
	Href string `json:"href"`
	// Link name
	Name *string `json:"name,omitempty"`
	// The media type of the link. If omitted, it is implicitly `application/json`.
	Type                 *string `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DomainLinksAllOfBrand DomainLinksAllOfBrand

// NewDomainLinksAllOfBrand instantiates a new DomainLinksAllOfBrand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainLinksAllOfBrand(href string) *DomainLinksAllOfBrand {
	this := DomainLinksAllOfBrand{}
	this.Href = href
	return &this
}

// NewDomainLinksAllOfBrandWithDefaults instantiates a new DomainLinksAllOfBrand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainLinksAllOfBrandWithDefaults() *DomainLinksAllOfBrand {
	this := DomainLinksAllOfBrand{}
	return &this
}

// GetHints returns the Hints field value if set, zero value otherwise.
func (o *DomainLinksAllOfBrand) GetHints() HrefObjectHints {
	if o == nil || o.Hints == nil {
		var ret HrefObjectHints
		return ret
	}
	return *o.Hints
}

// GetHintsOk returns a tuple with the Hints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinksAllOfBrand) GetHintsOk() (*HrefObjectHints, bool) {
	if o == nil || o.Hints == nil {
		return nil, false
	}
	return o.Hints, true
}

// HasHints returns a boolean if a field has been set.
func (o *DomainLinksAllOfBrand) HasHints() bool {
	if o != nil && o.Hints != nil {
		return true
	}

	return false
}

// SetHints gets a reference to the given HrefObjectHints and assigns it to the Hints field.
func (o *DomainLinksAllOfBrand) SetHints(v HrefObjectHints) {
	o.Hints = &v
}

// GetHref returns the Href field value
func (o *DomainLinksAllOfBrand) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *DomainLinksAllOfBrand) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *DomainLinksAllOfBrand) SetHref(v string) {
	o.Href = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DomainLinksAllOfBrand) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinksAllOfBrand) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DomainLinksAllOfBrand) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DomainLinksAllOfBrand) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DomainLinksAllOfBrand) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainLinksAllOfBrand) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DomainLinksAllOfBrand) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DomainLinksAllOfBrand) SetType(v string) {
	o.Type = &v
}

func (o DomainLinksAllOfBrand) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DomainLinksAllOfBrand) UnmarshalJSON(bytes []byte) (err error) {
	varDomainLinksAllOfBrand := _DomainLinksAllOfBrand{}

	err = json.Unmarshal(bytes, &varDomainLinksAllOfBrand)
	if err == nil {
		*o = DomainLinksAllOfBrand(varDomainLinksAllOfBrand)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "hints")
		delete(additionalProperties, "href")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableDomainLinksAllOfBrand struct {
	value *DomainLinksAllOfBrand
	isSet bool
}

func (v NullableDomainLinksAllOfBrand) Get() *DomainLinksAllOfBrand {
	return v.value
}

func (v *NullableDomainLinksAllOfBrand) Set(val *DomainLinksAllOfBrand) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainLinksAllOfBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainLinksAllOfBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainLinksAllOfBrand(val *DomainLinksAllOfBrand) *NullableDomainLinksAllOfBrand {
	return &NullableDomainLinksAllOfBrand{value: val, isSet: true}
}

func (v NullableDomainLinksAllOfBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainLinksAllOfBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
