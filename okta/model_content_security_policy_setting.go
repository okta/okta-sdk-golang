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
)

// checks if the ContentSecurityPolicySetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentSecurityPolicySetting{}

// ContentSecurityPolicySetting struct for ContentSecurityPolicySetting
type ContentSecurityPolicySetting struct {
	Mode                 *string  `json:"mode,omitempty"`
	ReportUri            *string  `json:"reportUri,omitempty"`
	SrcList              []string `json:"srcList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContentSecurityPolicySetting ContentSecurityPolicySetting

// NewContentSecurityPolicySetting instantiates a new ContentSecurityPolicySetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentSecurityPolicySetting() *ContentSecurityPolicySetting {
	this := ContentSecurityPolicySetting{}
	return &this
}

// NewContentSecurityPolicySettingWithDefaults instantiates a new ContentSecurityPolicySetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentSecurityPolicySettingWithDefaults() *ContentSecurityPolicySetting {
	this := ContentSecurityPolicySetting{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ContentSecurityPolicySetting) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSecurityPolicySetting) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ContentSecurityPolicySetting) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *ContentSecurityPolicySetting) SetMode(v string) {
	o.Mode = &v
}

// GetReportUri returns the ReportUri field value if set, zero value otherwise.
func (o *ContentSecurityPolicySetting) GetReportUri() string {
	if o == nil || IsNil(o.ReportUri) {
		var ret string
		return ret
	}
	return *o.ReportUri
}

// GetReportUriOk returns a tuple with the ReportUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSecurityPolicySetting) GetReportUriOk() (*string, bool) {
	if o == nil || IsNil(o.ReportUri) {
		return nil, false
	}
	return o.ReportUri, true
}

// HasReportUri returns a boolean if a field has been set.
func (o *ContentSecurityPolicySetting) HasReportUri() bool {
	if o != nil && !IsNil(o.ReportUri) {
		return true
	}

	return false
}

// SetReportUri gets a reference to the given string and assigns it to the ReportUri field.
func (o *ContentSecurityPolicySetting) SetReportUri(v string) {
	o.ReportUri = &v
}

// GetSrcList returns the SrcList field value if set, zero value otherwise.
func (o *ContentSecurityPolicySetting) GetSrcList() []string {
	if o == nil || IsNil(o.SrcList) {
		var ret []string
		return ret
	}
	return o.SrcList
}

// GetSrcListOk returns a tuple with the SrcList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentSecurityPolicySetting) GetSrcListOk() ([]string, bool) {
	if o == nil || IsNil(o.SrcList) {
		return nil, false
	}
	return o.SrcList, true
}

// HasSrcList returns a boolean if a field has been set.
func (o *ContentSecurityPolicySetting) HasSrcList() bool {
	if o != nil && !IsNil(o.SrcList) {
		return true
	}

	return false
}

// SetSrcList gets a reference to the given []string and assigns it to the SrcList field.
func (o *ContentSecurityPolicySetting) SetSrcList(v []string) {
	o.SrcList = v
}

func (o ContentSecurityPolicySetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentSecurityPolicySetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.ReportUri) {
		toSerialize["reportUri"] = o.ReportUri
	}
	if !IsNil(o.SrcList) {
		toSerialize["srcList"] = o.SrcList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContentSecurityPolicySetting) UnmarshalJSON(data []byte) (err error) {
	varContentSecurityPolicySetting := _ContentSecurityPolicySetting{}

	err = json.Unmarshal(data, &varContentSecurityPolicySetting)

	if err != nil {
		return err
	}

	*o = ContentSecurityPolicySetting(varContentSecurityPolicySetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		delete(additionalProperties, "reportUri")
		delete(additionalProperties, "srcList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContentSecurityPolicySetting struct {
	value *ContentSecurityPolicySetting
	isSet bool
}

func (v NullableContentSecurityPolicySetting) Get() *ContentSecurityPolicySetting {
	return v.value
}

func (v *NullableContentSecurityPolicySetting) Set(val *ContentSecurityPolicySetting) {
	v.value = val
	v.isSet = true
}

func (v NullableContentSecurityPolicySetting) IsSet() bool {
	return v.isSet
}

func (v *NullableContentSecurityPolicySetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentSecurityPolicySetting(val *ContentSecurityPolicySetting) *NullableContentSecurityPolicySetting {
	return &NullableContentSecurityPolicySetting{value: val, isSet: true}
}

func (v NullableContentSecurityPolicySetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentSecurityPolicySetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
