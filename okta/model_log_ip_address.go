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
)

// checks if the LogIpAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogIpAddress{}

// LogIpAddress struct for LogIpAddress
type LogIpAddress struct {
	GeographicalContext *LogGeographicalContext `json:"geographicalContext,omitempty"`
	// IP address
	Ip *string `json:"ip,omitempty"`
	// Details regarding the source
	Source *string `json:"source,omitempty"`
	// IP address version
	Version              *string `json:"version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogIpAddress LogIpAddress

// NewLogIpAddress instantiates a new LogIpAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogIpAddress() *LogIpAddress {
	this := LogIpAddress{}
	return &this
}

// NewLogIpAddressWithDefaults instantiates a new LogIpAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogIpAddressWithDefaults() *LogIpAddress {
	this := LogIpAddress{}
	return &this
}

// GetGeographicalContext returns the GeographicalContext field value if set, zero value otherwise.
func (o *LogIpAddress) GetGeographicalContext() LogGeographicalContext {
	if o == nil || IsNil(o.GeographicalContext) {
		var ret LogGeographicalContext
		return ret
	}
	return *o.GeographicalContext
}

// GetGeographicalContextOk returns a tuple with the GeographicalContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIpAddress) GetGeographicalContextOk() (*LogGeographicalContext, bool) {
	if o == nil || IsNil(o.GeographicalContext) {
		return nil, false
	}
	return o.GeographicalContext, true
}

// HasGeographicalContext returns a boolean if a field has been set.
func (o *LogIpAddress) HasGeographicalContext() bool {
	if o != nil && !IsNil(o.GeographicalContext) {
		return true
	}

	return false
}

// SetGeographicalContext gets a reference to the given LogGeographicalContext and assigns it to the GeographicalContext field.
func (o *LogIpAddress) SetGeographicalContext(v LogGeographicalContext) {
	o.GeographicalContext = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *LogIpAddress) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIpAddress) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *LogIpAddress) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *LogIpAddress) SetIp(v string) {
	o.Ip = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *LogIpAddress) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIpAddress) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *LogIpAddress) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *LogIpAddress) SetSource(v string) {
	o.Source = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *LogIpAddress) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogIpAddress) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *LogIpAddress) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *LogIpAddress) SetVersion(v string) {
	o.Version = &v
}

func (o LogIpAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogIpAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeographicalContext) {
		toSerialize["geographicalContext"] = o.GeographicalContext
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogIpAddress) UnmarshalJSON(data []byte) (err error) {
	varLogIpAddress := _LogIpAddress{}

	err = json.Unmarshal(data, &varLogIpAddress)

	if err != nil {
		return err
	}

	*o = LogIpAddress(varLogIpAddress)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "geographicalContext")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "source")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogIpAddress struct {
	value *LogIpAddress
	isSet bool
}

func (v NullableLogIpAddress) Get() *LogIpAddress {
	return v.value
}

func (v *NullableLogIpAddress) Set(val *LogIpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableLogIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableLogIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogIpAddress(val *LogIpAddress) *NullableLogIpAddress {
	return &NullableLogIpAddress{value: val, isSet: true}
}

func (v NullableLogIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
