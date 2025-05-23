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

// LogClient struct for LogClient
type LogClient struct {
	Device *string `json:"device,omitempty"`
	GeographicalContext *LogGeographicalContext `json:"geographicalContext,omitempty"`
	Id *string `json:"id,omitempty"`
	IpAddress *string `json:"ipAddress,omitempty"`
	UserAgent *LogUserAgent `json:"userAgent,omitempty"`
	Zone *string `json:"zone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogClient LogClient

// NewLogClient instantiates a new LogClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogClient() *LogClient {
	this := LogClient{}
	return &this
}

// NewLogClientWithDefaults instantiates a new LogClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogClientWithDefaults() *LogClient {
	this := LogClient{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *LogClient) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogClient) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *LogClient) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *LogClient) SetDevice(v string) {
	o.Device = &v
}

// GetGeographicalContext returns the GeographicalContext field value if set, zero value otherwise.
func (o *LogClient) GetGeographicalContext() LogGeographicalContext {
	if o == nil || o.GeographicalContext == nil {
		var ret LogGeographicalContext
		return ret
	}
	return *o.GeographicalContext
}

// GetGeographicalContextOk returns a tuple with the GeographicalContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogClient) GetGeographicalContextOk() (*LogGeographicalContext, bool) {
	if o == nil || o.GeographicalContext == nil {
		return nil, false
	}
	return o.GeographicalContext, true
}

// HasGeographicalContext returns a boolean if a field has been set.
func (o *LogClient) HasGeographicalContext() bool {
	if o != nil && o.GeographicalContext != nil {
		return true
	}

	return false
}

// SetGeographicalContext gets a reference to the given LogGeographicalContext and assigns it to the GeographicalContext field.
func (o *LogClient) SetGeographicalContext(v LogGeographicalContext) {
	o.GeographicalContext = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogClient) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogClient) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogClient) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogClient) SetId(v string) {
	o.Id = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *LogClient) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogClient) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *LogClient) HasIpAddress() bool {
	if o != nil && o.IpAddress != nil {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *LogClient) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *LogClient) GetUserAgent() LogUserAgent {
	if o == nil || o.UserAgent == nil {
		var ret LogUserAgent
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogClient) GetUserAgentOk() (*LogUserAgent, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *LogClient) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given LogUserAgent and assigns it to the UserAgent field.
func (o *LogClient) SetUserAgent(v LogUserAgent) {
	o.UserAgent = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *LogClient) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogClient) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *LogClient) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *LogClient) SetZone(v string) {
	o.Zone = &v
}

func (o LogClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}
	if o.GeographicalContext != nil {
		toSerialize["geographicalContext"] = o.GeographicalContext
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IpAddress != nil {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if o.UserAgent != nil {
		toSerialize["userAgent"] = o.UserAgent
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LogClient) UnmarshalJSON(bytes []byte) (err error) {
	varLogClient := _LogClient{}

	err = json.Unmarshal(bytes, &varLogClient)
	if err == nil {
		*o = LogClient(varLogClient)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "device")
		delete(additionalProperties, "geographicalContext")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ipAddress")
		delete(additionalProperties, "userAgent")
		delete(additionalProperties, "zone")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLogClient struct {
	value *LogClient
	isSet bool
}

func (v NullableLogClient) Get() *LogClient {
	return v.value
}

func (v *NullableLogClient) Set(val *LogClient) {
	v.value = val
	v.isSet = true
}

func (v NullableLogClient) IsSet() bool {
	return v.isSet
}

func (v *NullableLogClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogClient(val *LogClient) *NullableLogClient {
	return &NullableLogClient{value: val, isSet: true}
}

func (v NullableLogClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

