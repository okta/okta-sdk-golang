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

// IdentityProviderApplicationUserLinks struct for IdentityProviderApplicationUserLinks
type IdentityProviderApplicationUserLinks struct {
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Next *HrefObject `json:"next,omitempty"`
	// The IdP instance
	Idp *HrefObject `json:"idp,omitempty"`
	// The linked Okta user
	User *HrefObject `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityProviderApplicationUserLinks IdentityProviderApplicationUserLinks

// NewIdentityProviderApplicationUserLinks instantiates a new IdentityProviderApplicationUserLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderApplicationUserLinks() *IdentityProviderApplicationUserLinks {
	this := IdentityProviderApplicationUserLinks{}
	return &this
}

// NewIdentityProviderApplicationUserLinksWithDefaults instantiates a new IdentityProviderApplicationUserLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderApplicationUserLinksWithDefaults() *IdentityProviderApplicationUserLinks {
	this := IdentityProviderApplicationUserLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *IdentityProviderApplicationUserLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApplicationUserLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *IdentityProviderApplicationUserLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *IdentityProviderApplicationUserLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *IdentityProviderApplicationUserLinks) GetNext() HrefObject {
	if o == nil || o.Next == nil {
		var ret HrefObject
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApplicationUserLinks) GetNextOk() (*HrefObject, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *IdentityProviderApplicationUserLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given HrefObject and assigns it to the Next field.
func (o *IdentityProviderApplicationUserLinks) SetNext(v HrefObject) {
	o.Next = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *IdentityProviderApplicationUserLinks) GetIdp() HrefObject {
	if o == nil || o.Idp == nil {
		var ret HrefObject
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApplicationUserLinks) GetIdpOk() (*HrefObject, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *IdentityProviderApplicationUserLinks) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given HrefObject and assigns it to the Idp field.
func (o *IdentityProviderApplicationUserLinks) SetIdp(v HrefObject) {
	o.Idp = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *IdentityProviderApplicationUserLinks) GetUser() HrefObject {
	if o == nil || o.User == nil {
		var ret HrefObject
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApplicationUserLinks) GetUserOk() (*HrefObject, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *IdentityProviderApplicationUserLinks) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given HrefObject and assigns it to the User field.
func (o *IdentityProviderApplicationUserLinks) SetUser(v HrefObject) {
	o.User = &v
}

func (o IdentityProviderApplicationUserLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityProviderApplicationUserLinks) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityProviderApplicationUserLinks := _IdentityProviderApplicationUserLinks{}

	err = json.Unmarshal(bytes, &varIdentityProviderApplicationUserLinks)
	if err == nil {
		*o = IdentityProviderApplicationUserLinks(varIdentityProviderApplicationUserLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		delete(additionalProperties, "idp")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableIdentityProviderApplicationUserLinks struct {
	value *IdentityProviderApplicationUserLinks
	isSet bool
}

func (v NullableIdentityProviderApplicationUserLinks) Get() *IdentityProviderApplicationUserLinks {
	return v.value
}

func (v *NullableIdentityProviderApplicationUserLinks) Set(val *IdentityProviderApplicationUserLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderApplicationUserLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderApplicationUserLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderApplicationUserLinks(val *IdentityProviderApplicationUserLinks) *NullableIdentityProviderApplicationUserLinks {
	return &NullableIdentityProviderApplicationUserLinks{value: val, isSet: true}
}

func (v NullableIdentityProviderApplicationUserLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderApplicationUserLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

