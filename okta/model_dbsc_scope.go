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
	"fmt"
)

// checks if the DbscScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbscScope{}

// DbscScope struct for DbscScope
type DbscScope struct {
	// Whether to include subdomains in the binding scope (`false` = exact origin only, `true` = includes subdomains)
	IncludeSite bool `json:"include_site"`
	// The origin URL for which the DBSC binding is valid
	Origin               string `json:"origin"`
	AdditionalProperties map[string]interface{}
}

type _DbscScope DbscScope

// NewDbscScope instantiates a new DbscScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbscScope(includeSite bool, origin string) *DbscScope {
	this := DbscScope{}
	this.IncludeSite = includeSite
	this.Origin = origin
	return &this
}

// NewDbscScopeWithDefaults instantiates a new DbscScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbscScopeWithDefaults() *DbscScope {
	this := DbscScope{}
	return &this
}

// GetIncludeSite returns the IncludeSite field value
func (o *DbscScope) GetIncludeSite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeSite
}

// GetIncludeSiteOk returns a tuple with the IncludeSite field value
// and a boolean to check if the value has been set.
func (o *DbscScope) GetIncludeSiteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeSite, true
}

// SetIncludeSite sets field value
func (o *DbscScope) SetIncludeSite(v bool) {
	o.IncludeSite = v
}

// GetOrigin returns the Origin field value
func (o *DbscScope) GetOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *DbscScope) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *DbscScope) SetOrigin(v string) {
	o.Origin = v
}

func (o DbscScope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbscScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["include_site"] = o.IncludeSite
	toSerialize["origin"] = o.Origin

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DbscScope) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"include_site",
		"origin",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDbscScope := _DbscScope{}

	err = json.Unmarshal(data, &varDbscScope)

	if err != nil {
		return err
	}

	*o = DbscScope(varDbscScope)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "include_site")
		delete(additionalProperties, "origin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDbscScope struct {
	value *DbscScope
	isSet bool
}

func (v NullableDbscScope) Get() *DbscScope {
	return v.value
}

func (v *NullableDbscScope) Set(val *DbscScope) {
	v.value = val
	v.isSet = true
}

func (v NullableDbscScope) IsSet() bool {
	return v.isSet
}

func (v *NullableDbscScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbscScope(val *DbscScope) *NullableDbscScope {
	return &NullableDbscScope{value: val, isSet: true}
}

func (v NullableDbscScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbscScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
