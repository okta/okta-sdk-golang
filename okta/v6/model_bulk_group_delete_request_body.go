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

// BulkGroupDeleteRequestBody struct for BulkGroupDeleteRequestBody
type BulkGroupDeleteRequestBody struct {
	// Array of external IDs of groups that need to be deleted in Okta
	ExternalIds []string `json:"externalIds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkGroupDeleteRequestBody BulkGroupDeleteRequestBody

// NewBulkGroupDeleteRequestBody instantiates a new BulkGroupDeleteRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkGroupDeleteRequestBody() *BulkGroupDeleteRequestBody {
	this := BulkGroupDeleteRequestBody{}
	return &this
}

// NewBulkGroupDeleteRequestBodyWithDefaults instantiates a new BulkGroupDeleteRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkGroupDeleteRequestBodyWithDefaults() *BulkGroupDeleteRequestBody {
	this := BulkGroupDeleteRequestBody{}
	return &this
}

// GetExternalIds returns the ExternalIds field value if set, zero value otherwise.
func (o *BulkGroupDeleteRequestBody) GetExternalIds() []string {
	if o == nil || o.ExternalIds == nil {
		var ret []string
		return ret
	}
	return o.ExternalIds
}

// GetExternalIdsOk returns a tuple with the ExternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGroupDeleteRequestBody) GetExternalIdsOk() ([]string, bool) {
	if o == nil || o.ExternalIds == nil {
		return nil, false
	}
	return o.ExternalIds, true
}

// HasExternalIds returns a boolean if a field has been set.
func (o *BulkGroupDeleteRequestBody) HasExternalIds() bool {
	if o != nil && o.ExternalIds != nil {
		return true
	}

	return false
}

// SetExternalIds gets a reference to the given []string and assigns it to the ExternalIds field.
func (o *BulkGroupDeleteRequestBody) SetExternalIds(v []string) {
	o.ExternalIds = v
}

func (o BulkGroupDeleteRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalIds != nil {
		toSerialize["externalIds"] = o.ExternalIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BulkGroupDeleteRequestBody) UnmarshalJSON(bytes []byte) (err error) {
	varBulkGroupDeleteRequestBody := _BulkGroupDeleteRequestBody{}

	err = json.Unmarshal(bytes, &varBulkGroupDeleteRequestBody)
	if err == nil {
		*o = BulkGroupDeleteRequestBody(varBulkGroupDeleteRequestBody)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "externalIds")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBulkGroupDeleteRequestBody struct {
	value *BulkGroupDeleteRequestBody
	isSet bool
}

func (v NullableBulkGroupDeleteRequestBody) Get() *BulkGroupDeleteRequestBody {
	return v.value
}

func (v *NullableBulkGroupDeleteRequestBody) Set(val *BulkGroupDeleteRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkGroupDeleteRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkGroupDeleteRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkGroupDeleteRequestBody(val *BulkGroupDeleteRequestBody) *NullableBulkGroupDeleteRequestBody {
	return &NullableBulkGroupDeleteRequestBody{value: val, isSet: true}
}

func (v NullableBulkGroupDeleteRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkGroupDeleteRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

