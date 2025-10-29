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

// checks if the LogUserAgent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogUserAgent{}

// LogUserAgent \"A user agent is software (a software agent) that is acting on behalf of a user.\" ([Definition of User Agent](https://developer.mozilla.org/en-US/docs/Glossary/User_agent))  In the Okta event data object, the `UserAgent` object provides specifications about the client software that makes event-triggering HTTP requests. User agent identification is often useful for identifying interoperability problems between servers and clients, and also for browser and operating system usage analytics.
type LogUserAgent struct {
	// If the client is a web browser, this field identifies the type of web browser (for example, CHROME, FIREFOX)
	Browser *string `json:"browser,omitempty"`
	// The operating system that the client runs on (for example, Windows 10)
	Os *string `json:"os,omitempty"`
	// A raw string representation of the user agent that is formatted according to [section 5.5.3 of HTTP/1.1 Semantics and Content](https://datatracker.ietf.org/doc/html/rfc7231#section-5.5.3). Both the `browser` and the `OS` fields can be derived from this field.
	RawUserAgent         *string `json:"rawUserAgent,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LogUserAgent LogUserAgent

// NewLogUserAgent instantiates a new LogUserAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogUserAgent() *LogUserAgent {
	this := LogUserAgent{}
	return &this
}

// NewLogUserAgentWithDefaults instantiates a new LogUserAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogUserAgentWithDefaults() *LogUserAgent {
	this := LogUserAgent{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *LogUserAgent) GetBrowser() string {
	if o == nil || IsNil(o.Browser) {
		var ret string
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogUserAgent) GetBrowserOk() (*string, bool) {
	if o == nil || IsNil(o.Browser) {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *LogUserAgent) HasBrowser() bool {
	if o != nil && !IsNil(o.Browser) {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given string and assigns it to the Browser field.
func (o *LogUserAgent) SetBrowser(v string) {
	o.Browser = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *LogUserAgent) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogUserAgent) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *LogUserAgent) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *LogUserAgent) SetOs(v string) {
	o.Os = &v
}

// GetRawUserAgent returns the RawUserAgent field value if set, zero value otherwise.
func (o *LogUserAgent) GetRawUserAgent() string {
	if o == nil || IsNil(o.RawUserAgent) {
		var ret string
		return ret
	}
	return *o.RawUserAgent
}

// GetRawUserAgentOk returns a tuple with the RawUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogUserAgent) GetRawUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.RawUserAgent) {
		return nil, false
	}
	return o.RawUserAgent, true
}

// HasRawUserAgent returns a boolean if a field has been set.
func (o *LogUserAgent) HasRawUserAgent() bool {
	if o != nil && !IsNil(o.RawUserAgent) {
		return true
	}

	return false
}

// SetRawUserAgent gets a reference to the given string and assigns it to the RawUserAgent field.
func (o *LogUserAgent) SetRawUserAgent(v string) {
	o.RawUserAgent = &v
}

func (o LogUserAgent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogUserAgent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Browser) {
		toSerialize["browser"] = o.Browser
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.RawUserAgent) {
		toSerialize["rawUserAgent"] = o.RawUserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LogUserAgent) UnmarshalJSON(data []byte) (err error) {
	varLogUserAgent := _LogUserAgent{}

	err = json.Unmarshal(data, &varLogUserAgent)

	if err != nil {
		return err
	}

	*o = LogUserAgent(varLogUserAgent)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "browser")
		delete(additionalProperties, "os")
		delete(additionalProperties, "rawUserAgent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLogUserAgent struct {
	value *LogUserAgent
	isSet bool
}

func (v NullableLogUserAgent) Get() *LogUserAgent {
	return v.value
}

func (v *NullableLogUserAgent) Set(val *LogUserAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableLogUserAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableLogUserAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogUserAgent(val *LogUserAgent) *NullableLogUserAgent {
	return &NullableLogUserAgent{value: val, isSet: true}
}

func (v NullableLogUserAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogUserAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
