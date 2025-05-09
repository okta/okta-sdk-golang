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
	"time"
)

// RiskProvider struct for RiskProvider
type RiskProvider struct {
	// Action taken by Okta during authentication attempts based on the risk events sent by this provider
	Action string `json:"action"`
	// The ID of the [OAuth service app](https://developer.okta.com/docs/guides/implement-oauth-for-okta-serviceapp/main/#create-a-service-app-and-grant-scopes) that is used to send risk events to Okta
	ClientId string `json:"clientId"`
	// Timestamp when the Risk Provider object was created
	Created *time.Time `json:"created,omitempty"`
	// The ID of the Risk Provider object
	Id string `json:"id"`
	// Timestamp when the Risk Provider object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Name of the risk provider
	Name string `json:"name"`
	Links LinksSelf `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RiskProvider RiskProvider

// NewRiskProvider instantiates a new RiskProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskProvider(action string, clientId string, id string, name string, links LinksSelf) *RiskProvider {
	this := RiskProvider{}
	this.Action = action
	this.ClientId = clientId
	this.Id = id
	this.Name = name
	this.Links = links
	return &this
}

// NewRiskProviderWithDefaults instantiates a new RiskProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskProviderWithDefaults() *RiskProvider {
	this := RiskProvider{}
	var action string = "log_only"
	this.Action = action
	return &this
}

// GetAction returns the Action field value
func (o *RiskProvider) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RiskProvider) SetAction(v string) {
	o.Action = v
}

// GetClientId returns the ClientId field value
func (o *RiskProvider) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *RiskProvider) SetClientId(v string) {
	o.ClientId = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RiskProvider) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RiskProvider) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RiskProvider) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value
func (o *RiskProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RiskProvider) SetId(v string) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *RiskProvider) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *RiskProvider) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *RiskProvider) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *RiskProvider) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskProvider) SetName(v string) {
	o.Name = v
}

// GetLinks returns the Links field value
func (o *RiskProvider) GetLinks() LinksSelf {
	if o == nil {
		var ret LinksSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RiskProvider) GetLinksOk() (*LinksSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RiskProvider) SetLinks(v LinksSelf) {
	o.Links = v
}

func (o RiskProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["clientId"] = o.ClientId
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskProvider) UnmarshalJSON(bytes []byte) (err error) {
	varRiskProvider := _RiskProvider{}

	err = json.Unmarshal(bytes, &varRiskProvider)
	if err == nil {
		*o = RiskProvider(varRiskProvider)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "clientId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskProvider struct {
	value *RiskProvider
	isSet bool
}

func (v NullableRiskProvider) Get() *RiskProvider {
	return v.value
}

func (v *NullableRiskProvider) Set(val *RiskProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskProvider(val *RiskProvider) *NullableRiskProvider {
	return &NullableRiskProvider{value: val, isSet: true}
}

func (v NullableRiskProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

