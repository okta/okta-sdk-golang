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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"fmt"
)

// checks if the ServiceAccountDetailsAppAccountSub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccountDetailsAppAccountSub{}

// ServiceAccountDetailsAppAccountSub Details for a SaaS app account, which will be managed as a service account
type ServiceAccountDetailsAppAccountSub struct {
	// The name of the SaaS app in the Okta Integration Network catalog
	AppGlobalName *string `json:"appGlobalName,omitempty"`
	// The instance name of the SaaS app
	AppInstanceName *string                      `json:"appInstanceName,omitempty"`
	Credentials     AppServiceAccountCredentials `json:"credentials"`
	// The Okta app instance ID of the SaaS app
	OktaApplicationId    string `json:"oktaApplicationId"`
	AdditionalProperties map[string]interface{}
}

type _ServiceAccountDetailsAppAccountSub ServiceAccountDetailsAppAccountSub

// NewServiceAccountDetailsAppAccountSub instantiates a new ServiceAccountDetailsAppAccountSub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccountDetailsAppAccountSub(credentials AppServiceAccountCredentials, oktaApplicationId string) *ServiceAccountDetailsAppAccountSub {
	this := ServiceAccountDetailsAppAccountSub{}
	this.Credentials = credentials
	this.OktaApplicationId = oktaApplicationId
	return &this
}

// NewServiceAccountDetailsAppAccountSubWithDefaults instantiates a new ServiceAccountDetailsAppAccountSub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccountDetailsAppAccountSubWithDefaults() *ServiceAccountDetailsAppAccountSub {
	this := ServiceAccountDetailsAppAccountSub{}
	return &this
}

// GetAppGlobalName returns the AppGlobalName field value if set, zero value otherwise.
func (o *ServiceAccountDetailsAppAccountSub) GetAppGlobalName() string {
	if o == nil || IsNil(o.AppGlobalName) {
		var ret string
		return ret
	}
	return *o.AppGlobalName
}

// GetAppGlobalNameOk returns a tuple with the AppGlobalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsAppAccountSub) GetAppGlobalNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppGlobalName) {
		return nil, false
	}
	return o.AppGlobalName, true
}

// HasAppGlobalName returns a boolean if a field has been set.
func (o *ServiceAccountDetailsAppAccountSub) HasAppGlobalName() bool {
	if o != nil && !IsNil(o.AppGlobalName) {
		return true
	}

	return false
}

// SetAppGlobalName gets a reference to the given string and assigns it to the AppGlobalName field.
func (o *ServiceAccountDetailsAppAccountSub) SetAppGlobalName(v string) {
	o.AppGlobalName = &v
}

// GetAppInstanceName returns the AppInstanceName field value if set, zero value otherwise.
func (o *ServiceAccountDetailsAppAccountSub) GetAppInstanceName() string {
	if o == nil || IsNil(o.AppInstanceName) {
		var ret string
		return ret
	}
	return *o.AppInstanceName
}

// GetAppInstanceNameOk returns a tuple with the AppInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsAppAccountSub) GetAppInstanceNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceName) {
		return nil, false
	}
	return o.AppInstanceName, true
}

// HasAppInstanceName returns a boolean if a field has been set.
func (o *ServiceAccountDetailsAppAccountSub) HasAppInstanceName() bool {
	if o != nil && !IsNil(o.AppInstanceName) {
		return true
	}

	return false
}

// SetAppInstanceName gets a reference to the given string and assigns it to the AppInstanceName field.
func (o *ServiceAccountDetailsAppAccountSub) SetAppInstanceName(v string) {
	o.AppInstanceName = &v
}

// GetCredentials returns the Credentials field value
func (o *ServiceAccountDetailsAppAccountSub) GetCredentials() AppServiceAccountCredentials {
	if o == nil {
		var ret AppServiceAccountCredentials
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsAppAccountSub) GetCredentialsOk() (*AppServiceAccountCredentials, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Credentials, true
}

// SetCredentials sets field value
func (o *ServiceAccountDetailsAppAccountSub) SetCredentials(v AppServiceAccountCredentials) {
	o.Credentials = v
}

// GetOktaApplicationId returns the OktaApplicationId field value
func (o *ServiceAccountDetailsAppAccountSub) GetOktaApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaApplicationId
}

// GetOktaApplicationIdOk returns a tuple with the OktaApplicationId field value
// and a boolean to check if the value has been set.
func (o *ServiceAccountDetailsAppAccountSub) GetOktaApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaApplicationId, true
}

// SetOktaApplicationId sets field value
func (o *ServiceAccountDetailsAppAccountSub) SetOktaApplicationId(v string) {
	o.OktaApplicationId = v
}

func (o ServiceAccountDetailsAppAccountSub) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccountDetailsAppAccountSub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppGlobalName) {
		toSerialize["appGlobalName"] = o.AppGlobalName
	}
	if !IsNil(o.AppInstanceName) {
		toSerialize["appInstanceName"] = o.AppInstanceName
	}
	toSerialize["credentials"] = o.Credentials
	toSerialize["oktaApplicationId"] = o.OktaApplicationId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceAccountDetailsAppAccountSub) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"credentials",
		"oktaApplicationId",
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

	varServiceAccountDetailsAppAccountSub := _ServiceAccountDetailsAppAccountSub{}

	err = json.Unmarshal(data, &varServiceAccountDetailsAppAccountSub)

	if err != nil {
		return err
	}

	*o = ServiceAccountDetailsAppAccountSub(varServiceAccountDetailsAppAccountSub)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appGlobalName")
		delete(additionalProperties, "appInstanceName")
		delete(additionalProperties, "credentials")
		delete(additionalProperties, "oktaApplicationId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceAccountDetailsAppAccountSub struct {
	value *ServiceAccountDetailsAppAccountSub
	isSet bool
}

func (v NullableServiceAccountDetailsAppAccountSub) Get() *ServiceAccountDetailsAppAccountSub {
	return v.value
}

func (v *NullableServiceAccountDetailsAppAccountSub) Set(val *ServiceAccountDetailsAppAccountSub) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccountDetailsAppAccountSub) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccountDetailsAppAccountSub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccountDetailsAppAccountSub(val *ServiceAccountDetailsAppAccountSub) *NullableServiceAccountDetailsAppAccountSub {
	return &NullableServiceAccountDetailsAppAccountSub{value: val, isSet: true}
}

func (v NullableServiceAccountDetailsAppAccountSub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccountDetailsAppAccountSub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
