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

// PostAuthSessionFailureActionsObject struct for PostAuthSessionFailureActionsObject
type PostAuthSessionFailureActionsObject struct {
	Action *string `json:"action,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostAuthSessionFailureActionsObject PostAuthSessionFailureActionsObject

// NewPostAuthSessionFailureActionsObject instantiates a new PostAuthSessionFailureActionsObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAuthSessionFailureActionsObject() *PostAuthSessionFailureActionsObject {
	this := PostAuthSessionFailureActionsObject{}
	return &this
}

// NewPostAuthSessionFailureActionsObjectWithDefaults instantiates a new PostAuthSessionFailureActionsObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAuthSessionFailureActionsObjectWithDefaults() *PostAuthSessionFailureActionsObject {
	this := PostAuthSessionFailureActionsObject{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PostAuthSessionFailureActionsObject) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAuthSessionFailureActionsObject) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PostAuthSessionFailureActionsObject) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PostAuthSessionFailureActionsObject) SetAction(v string) {
	o.Action = &v
}

func (o PostAuthSessionFailureActionsObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PostAuthSessionFailureActionsObject) UnmarshalJSON(bytes []byte) (err error) {
	varPostAuthSessionFailureActionsObject := _PostAuthSessionFailureActionsObject{}

	err = json.Unmarshal(bytes, &varPostAuthSessionFailureActionsObject)
	if err == nil {
		*o = PostAuthSessionFailureActionsObject(varPostAuthSessionFailureActionsObject)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePostAuthSessionFailureActionsObject struct {
	value *PostAuthSessionFailureActionsObject
	isSet bool
}

func (v NullablePostAuthSessionFailureActionsObject) Get() *PostAuthSessionFailureActionsObject {
	return v.value
}

func (v *NullablePostAuthSessionFailureActionsObject) Set(val *PostAuthSessionFailureActionsObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAuthSessionFailureActionsObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAuthSessionFailureActionsObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAuthSessionFailureActionsObject(val *PostAuthSessionFailureActionsObject) *NullablePostAuthSessionFailureActionsObject {
	return &NullablePostAuthSessionFailureActionsObject{value: val, isSet: true}
}

func (v NullablePostAuthSessionFailureActionsObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAuthSessionFailureActionsObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

