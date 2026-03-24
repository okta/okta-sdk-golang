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
)

// checks if the StartOrgFailback200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartOrgFailback200Response{}

// StartOrgFailback200Response struct for StartOrgFailback200Response
type StartOrgFailback200Response struct {
	// Results of the failback operation
	Results              []StartOrgFailback200ResponseResultsInner `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartOrgFailback200Response StartOrgFailback200Response

// NewStartOrgFailback200Response instantiates a new StartOrgFailback200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartOrgFailback200Response() *StartOrgFailback200Response {
	this := StartOrgFailback200Response{}
	return &this
}

// NewStartOrgFailback200ResponseWithDefaults instantiates a new StartOrgFailback200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartOrgFailback200ResponseWithDefaults() *StartOrgFailback200Response {
	this := StartOrgFailback200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *StartOrgFailback200Response) GetResults() []StartOrgFailback200ResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []StartOrgFailback200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartOrgFailback200Response) GetResultsOk() ([]StartOrgFailback200ResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *StartOrgFailback200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []StartOrgFailback200ResponseResultsInner and assigns it to the Results field.
func (o *StartOrgFailback200Response) SetResults(v []StartOrgFailback200ResponseResultsInner) {
	o.Results = v
}

func (o StartOrgFailback200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartOrgFailback200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartOrgFailback200Response) UnmarshalJSON(data []byte) (err error) {
	varStartOrgFailback200Response := _StartOrgFailback200Response{}

	err = json.Unmarshal(data, &varStartOrgFailback200Response)

	if err != nil {
		return err
	}

	*o = StartOrgFailback200Response(varStartOrgFailback200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartOrgFailback200Response struct {
	value *StartOrgFailback200Response
	isSet bool
}

func (v NullableStartOrgFailback200Response) Get() *StartOrgFailback200Response {
	return v.value
}

func (v *NullableStartOrgFailback200Response) Set(val *StartOrgFailback200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStartOrgFailback200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStartOrgFailback200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartOrgFailback200Response(val *StartOrgFailback200Response) *NullableStartOrgFailback200Response {
	return &NullableStartOrgFailback200Response{value: val, isSet: true}
}

func (v NullableStartOrgFailback200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartOrgFailback200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
