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

// StandardRoleEmbeddedTargetsCatalog App targets
type StandardRoleEmbeddedTargetsCatalog struct {
	Apps []CatalogApplication `json:"apps,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StandardRoleEmbeddedTargetsCatalog StandardRoleEmbeddedTargetsCatalog

// NewStandardRoleEmbeddedTargetsCatalog instantiates a new StandardRoleEmbeddedTargetsCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStandardRoleEmbeddedTargetsCatalog() *StandardRoleEmbeddedTargetsCatalog {
	this := StandardRoleEmbeddedTargetsCatalog{}
	return &this
}

// NewStandardRoleEmbeddedTargetsCatalogWithDefaults instantiates a new StandardRoleEmbeddedTargetsCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStandardRoleEmbeddedTargetsCatalogWithDefaults() *StandardRoleEmbeddedTargetsCatalog {
	this := StandardRoleEmbeddedTargetsCatalog{}
	return &this
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *StandardRoleEmbeddedTargetsCatalog) GetApps() []CatalogApplication {
	if o == nil || o.Apps == nil {
		var ret []CatalogApplication
		return ret
	}
	return o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StandardRoleEmbeddedTargetsCatalog) GetAppsOk() ([]CatalogApplication, bool) {
	if o == nil || o.Apps == nil {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *StandardRoleEmbeddedTargetsCatalog) HasApps() bool {
	if o != nil && o.Apps != nil {
		return true
	}

	return false
}

// SetApps gets a reference to the given []CatalogApplication and assigns it to the Apps field.
func (o *StandardRoleEmbeddedTargetsCatalog) SetApps(v []CatalogApplication) {
	o.Apps = v
}

func (o StandardRoleEmbeddedTargetsCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Apps != nil {
		toSerialize["apps"] = o.Apps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StandardRoleEmbeddedTargetsCatalog) UnmarshalJSON(bytes []byte) (err error) {
	varStandardRoleEmbeddedTargetsCatalog := _StandardRoleEmbeddedTargetsCatalog{}

	err = json.Unmarshal(bytes, &varStandardRoleEmbeddedTargetsCatalog)
	if err == nil {
		*o = StandardRoleEmbeddedTargetsCatalog(varStandardRoleEmbeddedTargetsCatalog)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "apps")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableStandardRoleEmbeddedTargetsCatalog struct {
	value *StandardRoleEmbeddedTargetsCatalog
	isSet bool
}

func (v NullableStandardRoleEmbeddedTargetsCatalog) Get() *StandardRoleEmbeddedTargetsCatalog {
	return v.value
}

func (v *NullableStandardRoleEmbeddedTargetsCatalog) Set(val *StandardRoleEmbeddedTargetsCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableStandardRoleEmbeddedTargetsCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableStandardRoleEmbeddedTargetsCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStandardRoleEmbeddedTargetsCatalog(val *StandardRoleEmbeddedTargetsCatalog) *NullableStandardRoleEmbeddedTargetsCatalog {
	return &NullableStandardRoleEmbeddedTargetsCatalog{value: val, isSet: true}
}

func (v NullableStandardRoleEmbeddedTargetsCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStandardRoleEmbeddedTargetsCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

