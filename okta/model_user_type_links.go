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

// UserTypeLinks struct for UserTypeLinks
type UserTypeLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Schema *SourceLinksAllOfSchema `json:"schema,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserTypeLinks UserTypeLinks

// NewUserTypeLinks instantiates a new UserTypeLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserTypeLinks() *UserTypeLinks {
	this := UserTypeLinks{}
	return &this
}

// NewUserTypeLinksWithDefaults instantiates a new UserTypeLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTypeLinksWithDefaults() *UserTypeLinks {
	this := UserTypeLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *UserTypeLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypeLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *UserTypeLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *UserTypeLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *UserTypeLinks) GetSchema() SourceLinksAllOfSchema {
	if o == nil || o.Schema == nil {
		var ret SourceLinksAllOfSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserTypeLinks) GetSchemaOk() (*SourceLinksAllOfSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *UserTypeLinks) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given SourceLinksAllOfSchema and assigns it to the Schema field.
func (o *UserTypeLinks) SetSchema(v SourceLinksAllOfSchema) {
	o.Schema = &v
}

func (o UserTypeLinks) MarshalJSON() ([]byte, error) {
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

func (o *UserTypeLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserTypeLinks := _UserTypeLinks{}

	err = json.Unmarshal(bytes, &varUserTypeLinks)
	if err == nil {
		*o = UserTypeLinks(varUserTypeLinks)
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

type NullableUserTypeLinks struct {
	value *UserTypeLinks
	isSet bool
}

func (v NullableUserTypeLinks) Get() *UserTypeLinks {
	return v.value
}

func (v *NullableUserTypeLinks) Set(val *UserTypeLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypeLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypeLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypeLinks(val *UserTypeLinks) *NullableUserTypeLinks {
	return &NullableUserTypeLinks{value: val, isSet: true}
}

func (v NullableUserTypeLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypeLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

