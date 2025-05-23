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

// ApplicationLinks Discoverable resources related to the app
type ApplicationLinks struct {
	AccessPolicy *HrefObject `json:"accessPolicy,omitempty"`
	Activate *HrefObjectActivateLink `json:"activate,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Groups *HrefObject `json:"groups,omitempty"`
	Logo []HrefObject `json:"logo,omitempty"`
	Metadata *HrefObject `json:"metadata,omitempty"`
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	Users *HrefObject `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationLinks ApplicationLinks

// NewApplicationLinks instantiates a new ApplicationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationLinks() *ApplicationLinks {
	this := ApplicationLinks{}
	return &this
}

// NewApplicationLinksWithDefaults instantiates a new ApplicationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationLinksWithDefaults() *ApplicationLinks {
	this := ApplicationLinks{}
	return &this
}

// GetAccessPolicy returns the AccessPolicy field value if set, zero value otherwise.
func (o *ApplicationLinks) GetAccessPolicy() HrefObject {
	if o == nil || o.AccessPolicy == nil {
		var ret HrefObject
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetAccessPolicyOk() (*HrefObject, bool) {
	if o == nil || o.AccessPolicy == nil {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *ApplicationLinks) HasAccessPolicy() bool {
	if o != nil && o.AccessPolicy != nil {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given HrefObject and assigns it to the AccessPolicy field.
func (o *ApplicationLinks) SetAccessPolicy(v HrefObject) {
	o.AccessPolicy = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *ApplicationLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || o.Activate == nil {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *ApplicationLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *ApplicationLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *ApplicationLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || o.Deactivate == nil {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *ApplicationLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *ApplicationLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ApplicationLinks) GetGroups() HrefObject {
	if o == nil || o.Groups == nil {
		var ret HrefObject
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetGroupsOk() (*HrefObject, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ApplicationLinks) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given HrefObject and assigns it to the Groups field.
func (o *ApplicationLinks) SetGroups(v HrefObject) {
	o.Groups = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ApplicationLinks) GetLogo() []HrefObject {
	if o == nil || o.Logo == nil {
		var ret []HrefObject
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetLogoOk() ([]HrefObject, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ApplicationLinks) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []HrefObject and assigns it to the Logo field.
func (o *ApplicationLinks) SetLogo(v []HrefObject) {
	o.Logo = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationLinks) GetMetadata() HrefObject {
	if o == nil || o.Metadata == nil {
		var ret HrefObject
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetMetadataOk() (*HrefObject, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationLinks) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given HrefObject and assigns it to the Metadata field.
func (o *ApplicationLinks) SetMetadata(v HrefObject) {
	o.Metadata = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ApplicationLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ApplicationLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ApplicationLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ApplicationLinks) GetUsers() HrefObject {
	if o == nil || o.Users == nil {
		var ret HrefObject
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetUsersOk() (*HrefObject, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ApplicationLinks) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given HrefObject and assigns it to the Users field.
func (o *ApplicationLinks) SetUsers(v HrefObject) {
	o.Users = &v
}

func (o ApplicationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessPolicy != nil {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplicationLinks) UnmarshalJSON(bytes []byte) (err error) {
	varApplicationLinks := _ApplicationLinks{}

	err = json.Unmarshal(bytes, &varApplicationLinks)
	if err == nil {
		*o = ApplicationLinks(varApplicationLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "accessPolicy")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "self")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableApplicationLinks struct {
	value *ApplicationLinks
	isSet bool
}

func (v NullableApplicationLinks) Get() *ApplicationLinks {
	return v.value
}

func (v *NullableApplicationLinks) Set(val *ApplicationLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationLinks(val *ApplicationLinks) *NullableApplicationLinks {
	return &NullableApplicationLinks{value: val, isSet: true}
}

func (v NullableApplicationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

