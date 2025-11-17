/*
Okta Admin Management

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

// checks if the SubmissionResponseGlobalTokenRevocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmissionResponseGlobalTokenRevocation{}

// SubmissionResponseGlobalTokenRevocation struct for SubmissionResponseGlobalTokenRevocation
type SubmissionResponseGlobalTokenRevocation struct {
	// URL of the authorization server's global token revocation endpoint
	Endpoint string `json:"endpoint"`
	// The format of the subject
	SubjectFormat string `json:"subjectFormat"`
	// Authentication method <br> **Note:** Currently, only the `SIGNED_JWT` method is supported.
	AuthMethod string `json:"authMethod"`
	// Allow partial support for Universal Logout
	PartialLogout        *bool `json:"partialLogout,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionResponseGlobalTokenRevocation SubmissionResponseGlobalTokenRevocation

// NewSubmissionResponseGlobalTokenRevocation instantiates a new SubmissionResponseGlobalTokenRevocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionResponseGlobalTokenRevocation(endpoint string, subjectFormat string, authMethod string) *SubmissionResponseGlobalTokenRevocation {
	this := SubmissionResponseGlobalTokenRevocation{}
	this.Endpoint = endpoint
	this.SubjectFormat = subjectFormat
	this.AuthMethod = authMethod
	var partialLogout bool = false
	this.PartialLogout = &partialLogout
	return &this
}

// NewSubmissionResponseGlobalTokenRevocationWithDefaults instantiates a new SubmissionResponseGlobalTokenRevocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionResponseGlobalTokenRevocationWithDefaults() *SubmissionResponseGlobalTokenRevocation {
	this := SubmissionResponseGlobalTokenRevocation{}
	var partialLogout bool = false
	this.PartialLogout = &partialLogout
	return &this
}

// GetEndpoint returns the Endpoint field value
func (o *SubmissionResponseGlobalTokenRevocation) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *SubmissionResponseGlobalTokenRevocation) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *SubmissionResponseGlobalTokenRevocation) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetSubjectFormat returns the SubjectFormat field value
func (o *SubmissionResponseGlobalTokenRevocation) GetSubjectFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectFormat
}

// GetSubjectFormatOk returns a tuple with the SubjectFormat field value
// and a boolean to check if the value has been set.
func (o *SubmissionResponseGlobalTokenRevocation) GetSubjectFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectFormat, true
}

// SetSubjectFormat sets field value
func (o *SubmissionResponseGlobalTokenRevocation) SetSubjectFormat(v string) {
	o.SubjectFormat = v
}

// GetAuthMethod returns the AuthMethod field value
func (o *SubmissionResponseGlobalTokenRevocation) GetAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value
// and a boolean to check if the value has been set.
func (o *SubmissionResponseGlobalTokenRevocation) GetAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMethod, true
}

// SetAuthMethod sets field value
func (o *SubmissionResponseGlobalTokenRevocation) SetAuthMethod(v string) {
	o.AuthMethod = v
}

// GetPartialLogout returns the PartialLogout field value if set, zero value otherwise.
func (o *SubmissionResponseGlobalTokenRevocation) GetPartialLogout() bool {
	if o == nil || IsNil(o.PartialLogout) {
		var ret bool
		return ret
	}
	return *o.PartialLogout
}

// GetPartialLogoutOk returns a tuple with the PartialLogout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponseGlobalTokenRevocation) GetPartialLogoutOk() (*bool, bool) {
	if o == nil || IsNil(o.PartialLogout) {
		return nil, false
	}
	return o.PartialLogout, true
}

// HasPartialLogout returns a boolean if a field has been set.
func (o *SubmissionResponseGlobalTokenRevocation) HasPartialLogout() bool {
	if o != nil && !IsNil(o.PartialLogout) {
		return true
	}

	return false
}

// SetPartialLogout gets a reference to the given bool and assigns it to the PartialLogout field.
func (o *SubmissionResponseGlobalTokenRevocation) SetPartialLogout(v bool) {
	o.PartialLogout = &v
}

func (o SubmissionResponseGlobalTokenRevocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmissionResponseGlobalTokenRevocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["subjectFormat"] = o.SubjectFormat
	toSerialize["authMethod"] = o.AuthMethod
	if !IsNil(o.PartialLogout) {
		toSerialize["partialLogout"] = o.PartialLogout
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubmissionResponseGlobalTokenRevocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"endpoint",
		"subjectFormat",
		"authMethod",
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

	varSubmissionResponseGlobalTokenRevocation := _SubmissionResponseGlobalTokenRevocation{}

	err = json.Unmarshal(data, &varSubmissionResponseGlobalTokenRevocation)

	if err != nil {
		return err
	}

	*o = SubmissionResponseGlobalTokenRevocation(varSubmissionResponseGlobalTokenRevocation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "endpoint")
		delete(additionalProperties, "subjectFormat")
		delete(additionalProperties, "authMethod")
		delete(additionalProperties, "partialLogout")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubmissionResponseGlobalTokenRevocation struct {
	value *SubmissionResponseGlobalTokenRevocation
	isSet bool
}

func (v NullableSubmissionResponseGlobalTokenRevocation) Get() *SubmissionResponseGlobalTokenRevocation {
	return v.value
}

func (v *NullableSubmissionResponseGlobalTokenRevocation) Set(val *SubmissionResponseGlobalTokenRevocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionResponseGlobalTokenRevocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionResponseGlobalTokenRevocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionResponseGlobalTokenRevocation(val *SubmissionResponseGlobalTokenRevocation) *NullableSubmissionResponseGlobalTokenRevocation {
	return &NullableSubmissionResponseGlobalTokenRevocation{value: val, isSet: true}
}

func (v NullableSubmissionResponseGlobalTokenRevocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionResponseGlobalTokenRevocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
