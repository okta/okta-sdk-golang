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

// ApplicationLayoutsLinks struct for ApplicationLayoutsLinks
type ApplicationLayoutsLinks struct {
	General []HrefObject `json:"general,omitempty"`
	SignOn []HrefObject `json:"signOn,omitempty"`
	Provisioning []HrefObject `json:"provisioning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationLayoutsLinks ApplicationLayoutsLinks

// NewApplicationLayoutsLinks instantiates a new ApplicationLayoutsLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLayoutsLinks() *ApplicationLayoutsLinks {
	this := ApplicationLayoutsLinks{}
	return &this
}

// NewApplicationLayoutsLinksWithDefaults instantiates a new ApplicationLayoutsLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLayoutsLinksWithDefaults() *ApplicationLayoutsLinks {
	this := ApplicationLayoutsLinks{}
	return &this
}

// GetGeneral returns the General field value if set, zero value otherwise.
func (o *ApplicationLayoutsLinks) GetGeneral() []HrefObject {
	if o == nil || o.General == nil {
		var ret []HrefObject
		return ret
	}
	return o.General
}

// GetGeneralOk returns a tuple with the General field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutsLinks) GetGeneralOk() ([]HrefObject, bool) {
	if o == nil || o.General == nil {
		return nil, false
	}
	return o.General, true
}

// HasGeneral returns a boolean if a field has been set.
func (o *ApplicationLayoutsLinks) HasGeneral() bool {
	if o != nil && o.General != nil {
		return true
	}

	return false
}

// SetGeneral gets a reference to the given []HrefObject and assigns it to the General field.
func (o *ApplicationLayoutsLinks) SetGeneral(v []HrefObject) {
	o.General = v
}

// GetSignOn returns the SignOn field value if set, zero value otherwise.
func (o *ApplicationLayoutsLinks) GetSignOn() []HrefObject {
	if o == nil || o.SignOn == nil {
		var ret []HrefObject
		return ret
	}
	return o.SignOn
}

// GetSignOnOk returns a tuple with the SignOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutsLinks) GetSignOnOk() ([]HrefObject, bool) {
	if o == nil || o.SignOn == nil {
		return nil, false
	}
	return o.SignOn, true
}

// HasSignOn returns a boolean if a field has been set.
func (o *ApplicationLayoutsLinks) HasSignOn() bool {
	if o != nil && o.SignOn != nil {
		return true
	}

	return false
}

// SetSignOn gets a reference to the given []HrefObject and assigns it to the SignOn field.
func (o *ApplicationLayoutsLinks) SetSignOn(v []HrefObject) {
	o.SignOn = v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *ApplicationLayoutsLinks) GetProvisioning() []HrefObject {
	if o == nil || o.Provisioning == nil {
		var ret []HrefObject
		return ret
	}
	return o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLayoutsLinks) GetProvisioningOk() ([]HrefObject, bool) {
	if o == nil || o.Provisioning == nil {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *ApplicationLayoutsLinks) HasProvisioning() bool {
	if o != nil && o.Provisioning != nil {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given []HrefObject and assigns it to the Provisioning field.
func (o *ApplicationLayoutsLinks) SetProvisioning(v []HrefObject) {
	o.Provisioning = v
}

func (o ApplicationLayoutsLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.General != nil {
		toSerialize["general"] = o.General
	}
	if o.SignOn != nil {
		toSerialize["signOn"] = o.SignOn
	}
	if o.Provisioning != nil {
		toSerialize["provisioning"] = o.Provisioning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationLayoutsLinks) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationLayoutsLinks := _ApplicationLayoutsLinks{}

	err = json.Unmarshal(bytes, &varApplicationLayoutsLinks)
	if err == nil {
		*o = ApplicationLayoutsLinks(varApplicationLayoutsLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "general")
		delete(additionalProperties, "signOn")
		delete(additionalProperties, "provisioning")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationLayoutsLinks struct {
	value *ApplicationLayoutsLinks
	isSet bool
}

func (v NullableApplicationLayoutsLinks) Get() *ApplicationLayoutsLinks {
	return v.value
}

func (v *NullableApplicationLayoutsLinks) Set(val *ApplicationLayoutsLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLayoutsLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLayoutsLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLayoutsLinks(val *ApplicationLayoutsLinks) *NullableApplicationLayoutsLinks {
	return &NullableApplicationLayoutsLinks{value: val, isSet: true}
}

func (v NullableApplicationLayoutsLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLayoutsLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

