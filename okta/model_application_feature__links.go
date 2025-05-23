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

// ApplicationFeatureLinks struct for ApplicationFeatureLinks
type ApplicationFeatureLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationFeatureLinks ApplicationFeatureLinks

// NewApplicationFeatureLinks instantiates a new ApplicationFeatureLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFeatureLinks() *ApplicationFeatureLinks {
	this := ApplicationFeatureLinks{}
	return &this
}

// NewApplicationFeatureLinksWithDefaults instantiates a new ApplicationFeatureLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFeatureLinksWithDefaults() *ApplicationFeatureLinks {
	this := ApplicationFeatureLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ApplicationFeatureLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFeatureLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ApplicationFeatureLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ApplicationFeatureLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

func (o ApplicationFeatureLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationFeatureLinks) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationFeatureLinks := _ApplicationFeatureLinks{}

	err = json.Unmarshal(bytes, &varApplicationFeatureLinks)
	if err == nil {
		*o = ApplicationFeatureLinks(varApplicationFeatureLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationFeatureLinks struct {
	value *ApplicationFeatureLinks
	isSet bool
}

func (v NullableApplicationFeatureLinks) Get() *ApplicationFeatureLinks {
	return v.value
}

func (v *NullableApplicationFeatureLinks) Set(val *ApplicationFeatureLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFeatureLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFeatureLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFeatureLinks(val *ApplicationFeatureLinks) *NullableApplicationFeatureLinks {
	return &NullableApplicationFeatureLinks{value: val, isSet: true}
}

func (v NullableApplicationFeatureLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFeatureLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

