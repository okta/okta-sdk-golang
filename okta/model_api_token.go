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
	"time"
)

// checks if the ApiToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiToken{}

// ApiToken An API token for an Okta User. This token is NOT scoped any further and can be used for any API the user has permissions to call.
type ApiToken struct {
	ClientName  *string          `json:"clientName,omitempty"`
	Created     *time.Time       `json:"created,omitempty"`
	ExpiresAt   *time.Time       `json:"expiresAt,omitempty"`
	Id          *string          `json:"id,omitempty"`
	LastUpdated *time.Time       `json:"lastUpdated,omitempty"`
	Name        string           `json:"name"`
	Network     *ApiTokenNetwork `json:"network,omitempty"`
	// A time duration specified as an [ISO 8601 duration](https://en.wikipedia.org/wiki/ISO_8601#Durations).
	TokenWindow          *string    `json:"tokenWindow,omitempty" validate:"regexp=^P(?:$)(\\\\d+Y)?(\\\\d+M)?(\\\\d+W)?(\\\\d+D)?(T(?:\\\\d)(\\\\d+H)?(\\\\d+M)?(\\\\d+S)?)?$"`
	UserId               *string    `json:"userId,omitempty"`
	Link                 *LinksSelf `json:"_link,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApiToken ApiToken

// NewApiToken instantiates a new ApiToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiToken(name string) *ApiToken {
	this := ApiToken{}
	this.Name = name
	return &this
}

// NewApiTokenWithDefaults instantiates a new ApiToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenWithDefaults() *ApiToken {
	this := ApiToken{}
	return &this
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *ApiToken) GetClientName() string {
	if o == nil || IsNil(o.ClientName) {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetClientNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClientName) {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *ApiToken) HasClientName() bool {
	if o != nil && !IsNil(o.ClientName) {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *ApiToken) SetClientName(v string) {
	o.ClientName = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ApiToken) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ApiToken) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ApiToken) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ApiToken) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApiToken) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ApiToken) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiToken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiToken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiToken) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ApiToken) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ApiToken) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ApiToken) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *ApiToken) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiToken) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiToken) SetName(v string) {
	o.Name = v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ApiToken) GetNetwork() ApiTokenNetwork {
	if o == nil || IsNil(o.Network) {
		var ret ApiTokenNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetNetworkOk() (*ApiTokenNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ApiToken) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ApiTokenNetwork and assigns it to the Network field.
func (o *ApiToken) SetNetwork(v ApiTokenNetwork) {
	o.Network = &v
}

// GetTokenWindow returns the TokenWindow field value if set, zero value otherwise.
func (o *ApiToken) GetTokenWindow() string {
	if o == nil || IsNil(o.TokenWindow) {
		var ret string
		return ret
	}
	return *o.TokenWindow
}

// GetTokenWindowOk returns a tuple with the TokenWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetTokenWindowOk() (*string, bool) {
	if o == nil || IsNil(o.TokenWindow) {
		return nil, false
	}
	return o.TokenWindow, true
}

// HasTokenWindow returns a boolean if a field has been set.
func (o *ApiToken) HasTokenWindow() bool {
	if o != nil && !IsNil(o.TokenWindow) {
		return true
	}

	return false
}

// SetTokenWindow gets a reference to the given string and assigns it to the TokenWindow field.
func (o *ApiToken) SetTokenWindow(v string) {
	o.TokenWindow = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiToken) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiToken) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiToken) SetUserId(v string) {
	o.UserId = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ApiToken) GetLink() LinksSelf {
	if o == nil || IsNil(o.Link) {
		var ret LinksSelf
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetLinkOk() (*LinksSelf, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *ApiToken) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given LinksSelf and assigns it to the Link field.
func (o *ApiToken) SetLink(v LinksSelf) {
	o.Link = &v
}

func (o ApiToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientName) {
		toSerialize["clientName"] = o.ClientName
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.TokenWindow) {
		toSerialize["tokenWindow"] = o.TokenWindow
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Link) {
		toSerialize["_link"] = o.Link
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApiToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varApiToken := _ApiToken{}

	err = json.Unmarshal(data, &varApiToken)

	if err != nil {
		return err
	}

	*o = ApiToken(varApiToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "clientName")
		delete(additionalProperties, "created")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "network")
		delete(additionalProperties, "tokenWindow")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "_link")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApiToken struct {
	value *ApiToken
	isSet bool
}

func (v NullableApiToken) Get() *ApiToken {
	return v.value
}

func (v *NullableApiToken) Set(val *ApiToken) {
	v.value = val
	v.isSet = true
}

func (v NullableApiToken) IsSet() bool {
	return v.isSet
}

func (v *NullableApiToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiToken(val *ApiToken) *NullableApiToken {
	return &NullableApiToken{value: val, isSet: true}
}

func (v NullableApiToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
