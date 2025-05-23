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

// LinksSelfLifecycleAndAuthorize struct for LinksSelfLifecycleAndAuthorize
type LinksSelfLifecycleAndAuthorize struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Authorize *HrefObjectAuthorizeLink `json:"authorize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksSelfLifecycleAndAuthorize LinksSelfLifecycleAndAuthorize

// NewLinksSelfLifecycleAndAuthorize instantiates a new LinksSelfLifecycleAndAuthorize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksSelfLifecycleAndAuthorize() *LinksSelfLifecycleAndAuthorize {
	this := LinksSelfLifecycleAndAuthorize{}
	return &this
}

// NewLinksSelfLifecycleAndAuthorizeWithDefaults instantiates a new LinksSelfLifecycleAndAuthorize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksSelfLifecycleAndAuthorizeWithDefaults() *LinksSelfLifecycleAndAuthorize {
	this := LinksSelfLifecycleAndAuthorize{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksSelfLifecycleAndAuthorize) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfLifecycleAndAuthorize) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksSelfLifecycleAndAuthorize) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksSelfLifecycleAndAuthorize) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *LinksSelfLifecycleAndAuthorize) GetActivate() HrefObjectActivateLink {
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfLifecycleAndAuthorize) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *LinksSelfLifecycleAndAuthorize) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *LinksSelfLifecycleAndAuthorize) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *LinksSelfLifecycleAndAuthorize) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfLifecycleAndAuthorize) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *LinksSelfLifecycleAndAuthorize) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *LinksSelfLifecycleAndAuthorize) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetAuthorize returns the Authorize field value if set, zero value otherwise.
func (o *LinksSelfLifecycleAndAuthorize) GetAuthorize() HrefObjectAuthorizeLink {
	if o == nil || o.Authorize == nil {
		var ret HrefObjectAuthorizeLink
		return ret
	}
	return *o.Authorize
}

// GetAuthorizeOk returns a tuple with the Authorize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksSelfLifecycleAndAuthorize) GetAuthorizeOk() (*HrefObjectAuthorizeLink, bool) {
	if o == nil || o.Authorize == nil {
		return nil, false
	}
	return o.Authorize, true
}

// HasAuthorize returns a boolean if a field has been set.
func (o *LinksSelfLifecycleAndAuthorize) HasAuthorize() bool {
	if o != nil && o.Authorize != nil {
		return true
	}

	return false
}

// SetAuthorize gets a reference to the given HrefObjectAuthorizeLink and assigns it to the Authorize field.
func (o *LinksSelfLifecycleAndAuthorize) SetAuthorize(v HrefObjectAuthorizeLink) {
	o.Authorize = &v
}

func (o LinksSelfLifecycleAndAuthorize) MarshalJSON() ([]byte, error) {
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
	if o.Authorize != nil {
		toSerialize["authorize"] = o.Authorize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksSelfLifecycleAndAuthorize) UnmarshalJSON(bytes []byte) (err error) {
	varLinksSelfLifecycleAndAuthorize := _LinksSelfLifecycleAndAuthorize{}

	err = json.Unmarshal(bytes, &varLinksSelfLifecycleAndAuthorize)
	if err == nil {
		*o = LinksSelfLifecycleAndAuthorize(varLinksSelfLifecycleAndAuthorize)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "authorize")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksSelfLifecycleAndAuthorize struct {
	value *LinksSelfLifecycleAndAuthorize
	isSet bool
}

func (v NullableLinksSelfLifecycleAndAuthorize) Get() *LinksSelfLifecycleAndAuthorize {
	return v.value
}

func (v *NullableLinksSelfLifecycleAndAuthorize) Set(val *LinksSelfLifecycleAndAuthorize) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksSelfLifecycleAndAuthorize) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksSelfLifecycleAndAuthorize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksSelfLifecycleAndAuthorize(val *LinksSelfLifecycleAndAuthorize) *NullableLinksSelfLifecycleAndAuthorize {
	return &NullableLinksSelfLifecycleAndAuthorize{value: val, isSet: true}
}

func (v NullableLinksSelfLifecycleAndAuthorize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksSelfLifecycleAndAuthorize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

