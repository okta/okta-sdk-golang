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

package okta

import (
	"encoding/json"
)

// CSRLinks Specifies link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) available for the current status of a CSR object using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification. This object is used for dynamic discovery of related resources and lifecycle operations.
type CSRLinks struct {
	Publish *HrefCsrPublishLink `json:"publish,omitempty"`
	Self *HrefCsrSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CSRLinks CSRLinks

// NewCSRLinks instantiates a new CSRLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSRLinks() *CSRLinks {
	this := CSRLinks{}
	return &this
}

// NewCSRLinksWithDefaults instantiates a new CSRLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSRLinksWithDefaults() *CSRLinks {
	this := CSRLinks{}
	return &this
}

// GetPublish returns the Publish field value if set, zero value otherwise.
func (o *CSRLinks) GetPublish() HrefCsrPublishLink {
	if o == nil || o.Publish == nil {
		var ret HrefCsrPublishLink
		return ret
	}
	return *o.Publish
}

// GetPublishOk returns a tuple with the Publish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSRLinks) GetPublishOk() (*HrefCsrPublishLink, bool) {
	if o == nil || o.Publish == nil {
		return nil, false
	}
	return o.Publish, true
}

// HasPublish returns a boolean if a field has been set.
func (o *CSRLinks) HasPublish() bool {
	if o != nil && o.Publish != nil {
		return true
	}

	return false
}

// SetPublish gets a reference to the given HrefCsrPublishLink and assigns it to the Publish field.
func (o *CSRLinks) SetPublish(v HrefCsrPublishLink) {
	o.Publish = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *CSRLinks) GetSelf() HrefCsrSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefCsrSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSRLinks) GetSelfOk() (*HrefCsrSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *CSRLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefCsrSelfLink and assigns it to the Self field.
func (o *CSRLinks) SetSelf(v HrefCsrSelfLink) {
	o.Self = &v
}

func (o CSRLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Publish != nil {
		toSerialize["publish"] = o.Publish
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CSRLinks) UnmarshalJSON(bytes []byte) (err error) {
	varCSRLinks := _CSRLinks{}

	err = json.Unmarshal(bytes, &varCSRLinks)
	if err == nil {
		*o = CSRLinks(varCSRLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "publish")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCSRLinks struct {
	value *CSRLinks
	isSet bool
}

func (v NullableCSRLinks) Get() *CSRLinks {
	return v.value
}

func (v *NullableCSRLinks) Set(val *CSRLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableCSRLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableCSRLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSRLinks(val *CSRLinks) *NullableCSRLinks {
	return &NullableCSRLinks{value: val, isSet: true}
}

func (v NullableCSRLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSRLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

