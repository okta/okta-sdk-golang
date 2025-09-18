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
	"time"
)

// checks if the InlineHookCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineHookCreateResponse{}

// InlineHookCreateResponse An inline hook object that specifies the details of the inline hook
type InlineHookCreateResponse struct {
	Channel *InlineHookChannel `json:"channel,omitempty"`
	// Date of the inline hook creation
	Created *time.Time `json:"created,omitempty"`
	// The unique identifier for the inline hook
	Id *string `json:"id,omitempty"`
	// Date of the last inline hook update
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The display name of the inline hook
	Name   *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
	// One of the inline hook types
	Type *string `json:"type,omitempty"`
	// Version of the inline hook type. The currently supported version is `1.0.0`.
	Version              *string                `json:"version,omitempty"`
	Links                *InlineHookLinksCreate `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineHookCreateResponse InlineHookCreateResponse

// NewInlineHookCreateResponse instantiates a new InlineHookCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineHookCreateResponse() *InlineHookCreateResponse {
	this := InlineHookCreateResponse{}
	return &this
}

// NewInlineHookCreateResponseWithDefaults instantiates a new InlineHookCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineHookCreateResponseWithDefaults() *InlineHookCreateResponse {
	this := InlineHookCreateResponse{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetChannel() InlineHookChannel {
	if o == nil || IsNil(o.Channel) {
		var ret InlineHookChannel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetChannelOk() (*InlineHookChannel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given InlineHookChannel and assigns it to the Channel field.
func (o *InlineHookCreateResponse) SetChannel(v InlineHookChannel) {
	o.Channel = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *InlineHookCreateResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineHookCreateResponse) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *InlineHookCreateResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineHookCreateResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineHookCreateResponse) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineHookCreateResponse) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *InlineHookCreateResponse) SetVersion(v string) {
	o.Version = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InlineHookCreateResponse) GetLinks() InlineHookLinksCreate {
	if o == nil || IsNil(o.Links) {
		var ret InlineHookLinksCreate
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineHookCreateResponse) GetLinksOk() (*InlineHookLinksCreate, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InlineHookCreateResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given InlineHookLinksCreate and assigns it to the Links field.
func (o *InlineHookCreateResponse) SetLinks(v InlineHookLinksCreate) {
	o.Links = &v
}

func (o InlineHookCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineHookCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InlineHookCreateResponse) UnmarshalJSON(data []byte) (err error) {
	varInlineHookCreateResponse := _InlineHookCreateResponse{}

	err = json.Unmarshal(data, &varInlineHookCreateResponse)

	if err != nil {
		return err
	}

	*o = InlineHookCreateResponse(varInlineHookCreateResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "channel")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "version")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineHookCreateResponse struct {
	value *InlineHookCreateResponse
	isSet bool
}

func (v NullableInlineHookCreateResponse) Get() *InlineHookCreateResponse {
	return v.value
}

func (v *NullableInlineHookCreateResponse) Set(val *InlineHookCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineHookCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineHookCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineHookCreateResponse(val *InlineHookCreateResponse) *NullableInlineHookCreateResponse {
	return &NullableInlineHookCreateResponse{value: val, isSet: true}
}

func (v NullableInlineHookCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineHookCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
