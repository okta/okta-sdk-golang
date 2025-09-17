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
	"time"
)

// checks if the LogStream type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogStream{}

// LogStream struct for LogStream
type LogStream struct {
	// Timestamp when the log stream object was created
	Created time.Time `json:"created"`
	// Unique identifier for the log stream
	Id string `json:"id"`
	// Timestamp when the log stream object was last updated
	LastUpdated time.Time `json:"lastUpdated"`
	// Unique name for the log stream object
	Name string `json:"name"`
	// Lifecycle status of the log stream object
	Status string `json:"status"`
	// Specifies the streaming provider used  Supported providers:   * `aws_eventbridge` ([AWS EventBridge](https://aws.amazon.com/eventbridge))   * `splunk_cloud_logstreaming` ([Splunk Cloud](https://www.splunk.com/en_us/software/splunk-cloud-platform.html))  Select the provider type to see provider-specific configurations in the `settings` property:
	Type                 string                         `json:"type"`
	Links                LogStreamLinksSelfAndLifecycle `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _LogStream LogStream

// NewLogStream instantiates a new LogStream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogStream(created time.Time, id string, lastUpdated time.Time, name string, status string, type_ string, links LogStreamLinksSelfAndLifecycle) *LogStream {
	this := LogStream{}
	this.Created = created
	this.Id = id
	this.LastUpdated = lastUpdated
	this.Name = name
	this.Status = status
	this.Type = type_
	this.Links = links
	return &this
}

// NewLogStreamWithDefaults instantiates a new LogStream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogStreamWithDefaults() *LogStream {
	this := LogStream{}
	return &this
}

// GetCreated returns the Created field value
func (o *LogStream) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *LogStream) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *LogStream) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LogStream) SetId(v string) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *LogStream) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *LogStream) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetName returns the Name field value
func (o *LogStream) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LogStream) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *LogStream) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LogStream) SetStatus(v string) {
	o.Status = v
}

// GetType returns the Type field value
func (o *LogStream) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LogStream) SetType(v string) {
	o.Type = v
}

// GetLinks returns the Links field value
func (o *LogStream) GetLinks() LogStreamLinksSelfAndLifecycle {
	if o == nil {
		var ret LogStreamLinksSelfAndLifecycle
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LogStream) GetLinksOk() (*LogStreamLinksSelfAndLifecycle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *LogStream) SetLinks(v LogStreamLinksSelfAndLifecycle) {
	o.Links = v
}

func (o LogStream) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogStream) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	toSerialize["lastUpdated"] = o.LastUpdated
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["type"] = o.Type
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogStream) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"id",
		"lastUpdated",
		"name",
		"status",
		"type",
		"_links",
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

	varLogStream := _LogStream{}

	err = json.Unmarshal(data, &varLogStream)

	if err != nil {
		return err
	}

	*o = LogStream(varLogStream)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogStream struct {
	value *LogStream
	isSet bool
}

func (v NullableLogStream) Get() *LogStream {
	return v.value
}

func (v *NullableLogStream) Set(val *LogStream) {
	v.value = val
	v.isSet = true
}

func (v NullableLogStream) IsSet() bool {
	return v.isSet
}

func (v *NullableLogStream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogStream(val *LogStream) *NullableLogStream {
	return &NullableLogStream{value: val, isSet: true}
}

func (v NullableLogStream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogStream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
