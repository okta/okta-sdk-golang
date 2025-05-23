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

// SubscriptionLinks Discoverable resources related to the subscription
type SubscriptionLinks struct {
	Self *HrefObject `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionLinks SubscriptionLinks

// NewSubscriptionLinks instantiates a new SubscriptionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionLinks() *SubscriptionLinks {
	this := SubscriptionLinks{}
	return &this
}

// NewSubscriptionLinksWithDefaults instantiates a new SubscriptionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionLinksWithDefaults() *SubscriptionLinks {
	this := SubscriptionLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SubscriptionLinks) GetSelf() HrefObject {
	if o == nil || o.Self == nil {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SubscriptionLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *SubscriptionLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

func (o SubscriptionLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SubscriptionLinks) UnmarshalJSON(bytes []byte) (err error) {
	varSubscriptionLinks := _SubscriptionLinks{}

	err = json.Unmarshal(bytes, &varSubscriptionLinks)
	if err == nil {
		*o = SubscriptionLinks(varSubscriptionLinks)
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

type NullableSubscriptionLinks struct {
	value *SubscriptionLinks
	isSet bool
}

func (v NullableSubscriptionLinks) Get() *SubscriptionLinks {
	return v.value
}

func (v *NullableSubscriptionLinks) Set(val *SubscriptionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionLinks(val *SubscriptionLinks) *NullableSubscriptionLinks {
	return &NullableSubscriptionLinks{value: val, isSet: true}
}

func (v NullableSubscriptionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

