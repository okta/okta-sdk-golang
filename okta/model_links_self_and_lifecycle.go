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

// LinksSelfAndLifecycle struct for LinksSelfAndLifecycle
type LinksSelfAndLifecycle struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSelfAndLifecycle LinksSelfAndLifecycle

// NewLinksSelfAndLifecycle instantiates a new LinksSelfAndLifecycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSelfAndLifecycle() *LinksSelfAndLifecycle {
	this := LinksSelfAndLifecycle{}
	return &this
}

// NewLinksSelfAndLifecycleWithDefaults instantiates a new LinksSelfAndLifecycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSelfAndLifecycleWithDefaults() *LinksSelfAndLifecycle {
	this := LinksSelfAndLifecycle{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksSelfAndLifecycle) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndLifecycle) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelfAndLifecycle) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksSelfAndLifecycle) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LinksSelfAndLifecycle) GetActivate() HrefObjectActivateLink {
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndLifecycle) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LinksSelfAndLifecycle) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *LinksSelfAndLifecycle) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LinksSelfAndLifecycle) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfAndLifecycle) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LinksSelfAndLifecycle) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *LinksSelfAndLifecycle) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

func (o LinksSelfAndLifecycle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksSelfAndLifecycle) UnmarshalJSON(bytes []byte) (err error) {
	varLinksSelfAndLifecycle := _LinksSelfAndLifecycle{}

	err = json.Unmarshal(bytes, &varLinksSelfAndLifecycle)
	if err == nil {
		*o = LinksSelfAndLifecycle(varLinksSelfAndLifecycle)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksSelfAndLifecycle struct {
	value *LinksSelfAndLifecycle
	isSet bool
}

func (v NullableLinksSelfAndLifecycle) Get() *LinksSelfAndLifecycle {
	return v.value
}

func (v *NullableLinksSelfAndLifecycle) Set(val *LinksSelfAndLifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSelfAndLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSelfAndLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSelfAndLifecycle(val *LinksSelfAndLifecycle) *NullableLinksSelfAndLifecycle {
	return &NullableLinksSelfAndLifecycle{value: val, isSet: true}
}

func (v NullableLinksSelfAndLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSelfAndLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

