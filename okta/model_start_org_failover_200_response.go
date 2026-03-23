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

// checks if the StartOrgFailover200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartOrgFailover200Response{}

// StartOrgFailover200Response struct for StartOrgFailover200Response
type StartOrgFailover200Response struct {
	// Results of the failover operation
	Results              []StartOrgFailover200ResponseResultsInner `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StartOrgFailover200Response StartOrgFailover200Response

// NewStartOrgFailover200Response instantiates a new StartOrgFailover200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartOrgFailover200Response() *StartOrgFailover200Response {
	this := StartOrgFailover200Response{}
	return &this
}

// NewStartOrgFailover200ResponseWithDefaults instantiates a new StartOrgFailover200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartOrgFailover200ResponseWithDefaults() *StartOrgFailover200Response {
	this := StartOrgFailover200Response{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *StartOrgFailover200Response) GetResults() []StartOrgFailover200ResponseResultsInner {
	if o == nil || IsNil(o.Results) {
		var ret []StartOrgFailover200ResponseResultsInner
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartOrgFailover200Response) GetResultsOk() ([]StartOrgFailover200ResponseResultsInner, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *StartOrgFailover200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []StartOrgFailover200ResponseResultsInner and assigns it to the Results field.
func (o *StartOrgFailover200Response) SetResults(v []StartOrgFailover200ResponseResultsInner) {
	o.Results = v
}

func (o StartOrgFailover200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartOrgFailover200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *StartOrgFailover200Response) UnmarshalJSON(data []byte) (err error) {
	varStartOrgFailover200Response := _StartOrgFailover200Response{}

	err = json.Unmarshal(data, &varStartOrgFailover200Response)

	if err != nil {
		return err
	}

	*o = StartOrgFailover200Response(varStartOrgFailover200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStartOrgFailover200Response struct {
	value *StartOrgFailover200Response
	isSet bool
}

func (v NullableStartOrgFailover200Response) Get() *StartOrgFailover200Response {
	return v.value
}

func (v *NullableStartOrgFailover200Response) Set(val *StartOrgFailover200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableStartOrgFailover200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableStartOrgFailover200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartOrgFailover200Response(val *StartOrgFailover200Response) *NullableStartOrgFailover200Response {
	return &NullableStartOrgFailover200Response{value: val, isSet: true}
}

func (v NullableStartOrgFailover200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartOrgFailover200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
