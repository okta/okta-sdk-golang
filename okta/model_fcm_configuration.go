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

// FCMConfiguration struct for FCMConfiguration
type FCMConfiguration struct {
	// (Optional) File name for Admin Console display
	FileName *string `json:"fileName,omitempty"`
	// Project ID of FCM configuration
	ProjectId *string `json:"projectId,omitempty"`
	// JSON containing the private service account key and service account details. See [Creating and managing service account keys](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) for more information on creating service account keys in JSON.
	ServiceAccountJson map[string]interface{} `json:"serviceAccountJson,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FCMConfiguration FCMConfiguration

// NewFCMConfiguration instantiates a new FCMConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFCMConfiguration() *FCMConfiguration {
	this := FCMConfiguration{}
	return &this
}

// NewFCMConfigurationWithDefaults instantiates a new FCMConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFCMConfigurationWithDefaults() *FCMConfiguration {
	this := FCMConfiguration{}
	return &this
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *FCMConfiguration) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCMConfiguration) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *FCMConfiguration) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *FCMConfiguration) SetFileName(v string) {
	o.FileName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *FCMConfiguration) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCMConfiguration) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *FCMConfiguration) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *FCMConfiguration) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetServiceAccountJson returns the ServiceAccountJson field value if set, zero value otherwise.
func (o *FCMConfiguration) GetServiceAccountJson() map[string]interface{} {
	if o == nil || o.ServiceAccountJson == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ServiceAccountJson
}

// GetServiceAccountJsonOk returns a tuple with the ServiceAccountJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FCMConfiguration) GetServiceAccountJsonOk() (map[string]interface{}, bool) {
	if o == nil || o.ServiceAccountJson == nil {
		return nil, false
	}
	return o.ServiceAccountJson, true
}

// HasServiceAccountJson returns a boolean if a field has been set.
func (o *FCMConfiguration) HasServiceAccountJson() bool {
	if o != nil && o.ServiceAccountJson != nil {
		return true
	}

	return false
}

// SetServiceAccountJson gets a reference to the given map[string]interface{} and assigns it to the ServiceAccountJson field.
func (o *FCMConfiguration) SetServiceAccountJson(v map[string]interface{}) {
	o.ServiceAccountJson = v
}

func (o FCMConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ServiceAccountJson != nil {
		toSerialize["serviceAccountJson"] = o.ServiceAccountJson
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FCMConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varFCMConfiguration := _FCMConfiguration{}

	err = json.Unmarshal(bytes, &varFCMConfiguration)
	if err == nil {
		*o = FCMConfiguration(varFCMConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "fileName")
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "serviceAccountJson")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableFCMConfiguration struct {
	value *FCMConfiguration
	isSet bool
}

func (v NullableFCMConfiguration) Get() *FCMConfiguration {
	return v.value
}

func (v *NullableFCMConfiguration) Set(val *FCMConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableFCMConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableFCMConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFCMConfiguration(val *FCMConfiguration) *NullableFCMConfiguration {
	return &NullableFCMConfiguration{value: val, isSet: true}
}

func (v NullableFCMConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFCMConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

