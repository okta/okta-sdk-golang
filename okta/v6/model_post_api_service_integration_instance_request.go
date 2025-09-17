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

// checks if the PostAPIServiceIntegrationInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAPIServiceIntegrationInstanceRequest{}

// PostAPIServiceIntegrationInstanceRequest struct for PostAPIServiceIntegrationInstanceRequest
type PostAPIServiceIntegrationInstanceRequest struct {
	// The list of Okta management scopes granted to the API Service Integration instance. See [Okta management OAuth 2.0 scopes](/oauth2/#okta-admin-management).
	GrantedScopes []string `json:"grantedScopes"`
	// App instance properties
	Properties *map[string]AppPropertiesValue `json:"properties,omitempty"`
	// The type of the API service integration. This string is an underscore-concatenated, lowercased API service integration name. For example, `my_api_log_integration`.
	Type                 string `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _PostAPIServiceIntegrationInstanceRequest PostAPIServiceIntegrationInstanceRequest

// NewPostAPIServiceIntegrationInstanceRequest instantiates a new PostAPIServiceIntegrationInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAPIServiceIntegrationInstanceRequest(grantedScopes []string, type_ string) *PostAPIServiceIntegrationInstanceRequest {
	this := PostAPIServiceIntegrationInstanceRequest{}
	this.GrantedScopes = grantedScopes
	this.Type = type_
	return &this
}

// NewPostAPIServiceIntegrationInstanceRequestWithDefaults instantiates a new PostAPIServiceIntegrationInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAPIServiceIntegrationInstanceRequestWithDefaults() *PostAPIServiceIntegrationInstanceRequest {
	this := PostAPIServiceIntegrationInstanceRequest{}
	return &this
}

// GetGrantedScopes returns the GrantedScopes field value
func (o *PostAPIServiceIntegrationInstanceRequest) GetGrantedScopes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GrantedScopes
}

// GetGrantedScopesOk returns a tuple with the GrantedScopes field value
// and a boolean to check if the value has been set.
func (o *PostAPIServiceIntegrationInstanceRequest) GetGrantedScopesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantedScopes, true
}

// SetGrantedScopes sets field value
func (o *PostAPIServiceIntegrationInstanceRequest) SetGrantedScopes(v []string) {
	o.GrantedScopes = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *PostAPIServiceIntegrationInstanceRequest) GetProperties() map[string]AppPropertiesValue {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]AppPropertiesValue
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAPIServiceIntegrationInstanceRequest) GetPropertiesOk() (*map[string]AppPropertiesValue, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *PostAPIServiceIntegrationInstanceRequest) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]AppPropertiesValue and assigns it to the Properties field.
func (o *PostAPIServiceIntegrationInstanceRequest) SetProperties(v map[string]AppPropertiesValue) {
	o.Properties = &v
}

// GetType returns the Type field value
func (o *PostAPIServiceIntegrationInstanceRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostAPIServiceIntegrationInstanceRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostAPIServiceIntegrationInstanceRequest) SetType(v string) {
	o.Type = v
}

func (o PostAPIServiceIntegrationInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAPIServiceIntegrationInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["grantedScopes"] = o.GrantedScopes
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostAPIServiceIntegrationInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"grantedScopes",
		"type",
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

	varPostAPIServiceIntegrationInstanceRequest := _PostAPIServiceIntegrationInstanceRequest{}

	err = json.Unmarshal(data, &varPostAPIServiceIntegrationInstanceRequest)

	if err != nil {
		return err
	}

	*o = PostAPIServiceIntegrationInstanceRequest(varPostAPIServiceIntegrationInstanceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "grantedScopes")
		delete(additionalProperties, "properties")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostAPIServiceIntegrationInstanceRequest struct {
	value *PostAPIServiceIntegrationInstanceRequest
	isSet bool
}

func (v NullablePostAPIServiceIntegrationInstanceRequest) Get() *PostAPIServiceIntegrationInstanceRequest {
	return v.value
}

func (v *NullablePostAPIServiceIntegrationInstanceRequest) Set(val *PostAPIServiceIntegrationInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAPIServiceIntegrationInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAPIServiceIntegrationInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAPIServiceIntegrationInstanceRequest(val *PostAPIServiceIntegrationInstanceRequest) *NullablePostAPIServiceIntegrationInstanceRequest {
	return &NullablePostAPIServiceIntegrationInstanceRequest{value: val, isSet: true}
}

func (v NullablePostAPIServiceIntegrationInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAPIServiceIntegrationInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
