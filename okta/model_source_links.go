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

// SourceLinks struct for SourceLinks
type SourceLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Schema *SourceLinksAllOfSchema `json:"schema,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SourceLinks SourceLinks

// NewSourceLinks instantiates a new SourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceLinks() *SourceLinks {
	this := SourceLinks{}
	return &this
}

// NewSourceLinksWithDefaults instantiates a new SourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceLinksWithDefaults() *SourceLinks {
	this := SourceLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SourceLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SourceLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *SourceLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *SourceLinks) GetSchema() SourceLinksAllOfSchema {
	if o == nil || o.Schema == nil {
		var ret SourceLinksAllOfSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceLinks) GetSchemaOk() (*SourceLinksAllOfSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *SourceLinks) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given SourceLinksAllOfSchema and assigns it to the Schema field.
func (o *SourceLinks) SetSchema(v SourceLinksAllOfSchema) {
	o.Schema = &v
}

func (o SourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SourceLinks) UnmarshalJSON(bytes []byte) (err error) {
	varSourceLinks := _SourceLinks{}

	err = json.Unmarshal(bytes, &varSourceLinks)
	if err == nil {
		*o = SourceLinks(varSourceLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "schema")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSourceLinks struct {
	value *SourceLinks
	isSet bool
}

func (v NullableSourceLinks) Get() *SourceLinks {
	return v.value
}

func (v *NullableSourceLinks) Set(val *SourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceLinks(val *SourceLinks) *NullableSourceLinks {
	return &NullableSourceLinks{value: val, isSet: true}
}

func (v NullableSourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

