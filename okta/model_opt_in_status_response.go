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

// OptInStatusResponse struct for OptInStatusResponse
type OptInStatusResponse struct {
	OptInStatus *string `json:"optInStatus,omitempty"`
	Links *OptInStatusResponseLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OptInStatusResponse OptInStatusResponse

// NewOptInStatusResponse instantiates a new OptInStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptInStatusResponse() *OptInStatusResponse {
	this := OptInStatusResponse{}
	return &this
}

// NewOptInStatusResponseWithDefaults instantiates a new OptInStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptInStatusResponseWithDefaults() *OptInStatusResponse {
	this := OptInStatusResponse{}
	return &this
}

// GetOptInStatus returns the OptInStatus field value if set, zero value otherwise.
func (o *OptInStatusResponse) GetOptInStatus() string {
	if o == nil || o.OptInStatus == nil {
		var ret string
		return ret
	}
	return *o.OptInStatus
}

// GetOptInStatusOk returns a tuple with the OptInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptInStatusResponse) GetOptInStatusOk() (*string, bool) {
	if o == nil || o.OptInStatus == nil {
		return nil, false
	}
	return o.OptInStatus, true
}

// HasOptInStatus returns a boolean if a field has been set.
func (o *OptInStatusResponse) HasOptInStatus() bool {
	if o != nil && o.OptInStatus != nil {
		return true
	}

	return false
}

// SetOptInStatus gets a reference to the given string and assigns it to the OptInStatus field.
func (o *OptInStatusResponse) SetOptInStatus(v string) {
	o.OptInStatus = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OptInStatusResponse) GetLinks() OptInStatusResponseLinks {
	if o == nil || o.Links == nil {
		var ret OptInStatusResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptInStatusResponse) GetLinksOk() (*OptInStatusResponseLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OptInStatusResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OptInStatusResponseLinks and assigns it to the Links field.
func (o *OptInStatusResponse) SetLinks(v OptInStatusResponseLinks) {
	o.Links = &v
}

func (o OptInStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OptInStatus != nil {
		toSerialize["optInStatus"] = o.OptInStatus
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OptInStatusResponse) UnmarshalJSON(bytes []byte) (err error) {
	varOptInStatusResponse := _OptInStatusResponse{}

	err = json.Unmarshal(bytes, &varOptInStatusResponse)
	if err == nil {
		*o = OptInStatusResponse(varOptInStatusResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "optInStatus")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableOptInStatusResponse struct {
	value *OptInStatusResponse
	isSet bool
}

func (v NullableOptInStatusResponse) Get() *OptInStatusResponse {
	return v.value
}

func (v *NullableOptInStatusResponse) Set(val *OptInStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptInStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptInStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptInStatusResponse(val *OptInStatusResponse) *NullableOptInStatusResponse {
	return &NullableOptInStatusResponse{value: val, isSet: true}
}

func (v NullableOptInStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptInStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

