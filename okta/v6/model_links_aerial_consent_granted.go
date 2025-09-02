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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// LinksAerialConsentGranted struct for LinksAerialConsentGranted
type LinksAerialConsentGranted struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Revoke *HrefObjectRevokeAerialConsent `json:"revoke,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinksAerialConsentGranted LinksAerialConsentGranted

// NewLinksAerialConsentGranted instantiates a new LinksAerialConsentGranted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinksAerialConsentGranted() *LinksAerialConsentGranted {
	this := LinksAerialConsentGranted{}
	return &this
}

// NewLinksAerialConsentGrantedWithDefaults instantiates a new LinksAerialConsentGranted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksAerialConsentGrantedWithDefaults() *LinksAerialConsentGranted {
	this := LinksAerialConsentGranted{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *LinksAerialConsentGranted) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksAerialConsentGranted) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *LinksAerialConsentGranted) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *LinksAerialConsentGranted) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetRevoke returns the Revoke field value if set, zero value otherwise.
func (o *LinksAerialConsentGranted) GetRevoke() HrefObjectRevokeAerialConsent {
	if o == nil || o.Revoke == nil {
		var ret HrefObjectRevokeAerialConsent
		return ret
	}
	return *o.Revoke
}

// GetRevokeOk returns a tuple with the Revoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinksAerialConsentGranted) GetRevokeOk() (*HrefObjectRevokeAerialConsent, bool) {
	if o == nil || o.Revoke == nil {
		return nil, false
	}
	return o.Revoke, true
}

// HasRevoke returns a boolean if a field has been set.
func (o *LinksAerialConsentGranted) HasRevoke() bool {
	if o != nil && o.Revoke != nil {
		return true
	}

	return false
}

// SetRevoke gets a reference to the given HrefObjectRevokeAerialConsent and assigns it to the Revoke field.
func (o *LinksAerialConsentGranted) SetRevoke(v HrefObjectRevokeAerialConsent) {
	o.Revoke = &v
}

func (o LinksAerialConsentGranted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Revoke != nil {
		toSerialize["revoke"] = o.Revoke
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinksAerialConsentGranted) UnmarshalJSON(bytes []byte) (err error) {
	varLinksAerialConsentGranted := _LinksAerialConsentGranted{}

	err = json.Unmarshal(bytes, &varLinksAerialConsentGranted)
	if err == nil {
		*o = LinksAerialConsentGranted(varLinksAerialConsentGranted)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "revoke")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableLinksAerialConsentGranted struct {
	value *LinksAerialConsentGranted
	isSet bool
}

func (v NullableLinksAerialConsentGranted) Get() *LinksAerialConsentGranted {
	return v.value
}

func (v *NullableLinksAerialConsentGranted) Set(val *LinksAerialConsentGranted) {
	v.value = val
	v.isSet = true
}

func (v NullableLinksAerialConsentGranted) IsSet() bool {
	return v.isSet
}

func (v *NullableLinksAerialConsentGranted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinksAerialConsentGranted(val *LinksAerialConsentGranted) *NullableLinksAerialConsentGranted {
	return &NullableLinksAerialConsentGranted{value: val, isSet: true}
}

func (v NullableLinksAerialConsentGranted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinksAerialConsentGranted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

