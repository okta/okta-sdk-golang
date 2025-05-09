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

// ThreatInsightConfiguration struct for ThreatInsightConfiguration
type ThreatInsightConfiguration struct {
	// Specifies how Okta responds to authentication requests from suspicious IP addresses
	Action string `json:"action"`
	// Timestamp when the ThreatInsight Configuration object was created
	Created *time.Time `json:"created,omitempty"`
	// Accepts a list of [Network Zone](/openapi/okta-management/management/tag/NetworkZone/) IDs. IPs in the excluded network zones aren't logged or blocked. This ensures that traffic from known, trusted IPs isn't accidentally logged or blocked.
	ExcludeZones []string `json:"excludeZones,omitempty"`
	// Timestamp when the ThreatInsight Configuration object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThreatInsightConfiguration ThreatInsightConfiguration

// NewThreatInsightConfiguration instantiates a new ThreatInsightConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreatInsightConfiguration(action string) *ThreatInsightConfiguration {
	this := ThreatInsightConfiguration{}
	this.Action = action
	return &this
}

// NewThreatInsightConfigurationWithDefaults instantiates a new ThreatInsightConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreatInsightConfigurationWithDefaults() *ThreatInsightConfiguration {
	this := ThreatInsightConfiguration{}
	return &this
}

// GetAction returns the Action field value
func (o *ThreatInsightConfiguration) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ThreatInsightConfiguration) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ThreatInsightConfiguration) SetAction(v string) {
	o.Action = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ThreatInsightConfiguration) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatInsightConfiguration) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ThreatInsightConfiguration) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ThreatInsightConfiguration) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExcludeZones returns the ExcludeZones field value if set, zero value otherwise.
func (o *ThreatInsightConfiguration) GetExcludeZones() []string {
	if o == nil || o.ExcludeZones == nil {
		var ret []string
		return ret
	}
	return o.ExcludeZones
}

// GetExcludeZonesOk returns a tuple with the ExcludeZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatInsightConfiguration) GetExcludeZonesOk() ([]string, bool) {
	if o == nil || o.ExcludeZones == nil {
		return nil, false
	}
	return o.ExcludeZones, true
}

// HasExcludeZones returns a boolean if a field has been set.
func (o *ThreatInsightConfiguration) HasExcludeZones() bool {
	if o != nil && o.ExcludeZones != nil {
		return true
	}

	return false
}

// SetExcludeZones gets a reference to the given []string and assigns it to the ExcludeZones field.
func (o *ThreatInsightConfiguration) SetExcludeZones(v []string) {
	o.ExcludeZones = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ThreatInsightConfiguration) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatInsightConfiguration) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ThreatInsightConfiguration) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ThreatInsightConfiguration) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ThreatInsightConfiguration) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThreatInsightConfiguration) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ThreatInsightConfiguration) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *ThreatInsightConfiguration) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o ThreatInsightConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.ExcludeZones != nil {
		toSerialize["excludeZones"] = o.ExcludeZones
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ThreatInsightConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varThreatInsightConfiguration := _ThreatInsightConfiguration{}

	err = json.Unmarshal(bytes, &varThreatInsightConfiguration)
	if err == nil {
		*o = ThreatInsightConfiguration(varThreatInsightConfiguration)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "created")
		delete(additionalProperties, "excludeZones")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableThreatInsightConfiguration struct {
	value *ThreatInsightConfiguration
	isSet bool
}

func (v NullableThreatInsightConfiguration) Get() *ThreatInsightConfiguration {
	return v.value
}

func (v *NullableThreatInsightConfiguration) Set(val *ThreatInsightConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableThreatInsightConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableThreatInsightConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreatInsightConfiguration(val *ThreatInsightConfiguration) *NullableThreatInsightConfiguration {
	return &NullableThreatInsightConfiguration{value: val, isSet: true}
}

func (v NullableThreatInsightConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreatInsightConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

