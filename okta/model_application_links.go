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

// checks if the ApplicationLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationLinks{}

// ApplicationLinks Discoverable resources related to the app
type ApplicationLinks struct {
	AccessPolicy *AccessPolicyLink       `json:"accessPolicy,omitempty"`
	Activate     *HrefObjectActivateLink `json:"activate,omitempty"`
	// List of app link resources
	AppLinks   []HrefObject              `json:"appLinks,omitempty"`
	Deactivate *HrefObjectDeactivateLink `json:"deactivate,omitempty"`
	Groups     *GroupsLink               `json:"groups,omitempty"`
	Help       *HelpLink                 `json:"help,omitempty"`
	// List of app logo resources
	Logo                 []HrefObject        `json:"logo,omitempty"`
	Metadata             *MetadataLink       `json:"metadata,omitempty"`
	Self                 *HrefObjectSelfLink `json:"self,omitempty"`
	Users                *UsersLink          `json:"users,omitempty"`
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
func (o *ApplicationLinks) GetAccessPolicy() AccessPolicyLink {
	if o == nil || IsNil(o.AccessPolicy) {
		var ret AccessPolicyLink
		return ret
	}
	return *o.AccessPolicy
}

// GetAccessPolicyOk returns a tuple with the AccessPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetAccessPolicyOk() (*AccessPolicyLink, bool) {
	if o == nil || IsNil(o.AccessPolicy) {
		return nil, false
	}
	return o.AccessPolicy, true
}

// HasAccessPolicy returns a boolean if a field has been set.
func (o *ApplicationLinks) HasAccessPolicy() bool {
	if o != nil && !IsNil(o.AccessPolicy) {
		return true
	}

	return false
}

// SetAccessPolicy gets a reference to the given AccessPolicyLink and assigns it to the AccessPolicy field.
func (o *ApplicationLinks) SetAccessPolicy(v AccessPolicyLink) {
	o.AccessPolicy = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *ApplicationLinks) GetActivate() HrefObjectActivateLink {
	if o == nil || IsNil(o.Activate) {
		var ret HrefObjectActivateLink
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetActivateOk() (*HrefObjectActivateLink, bool) {
	if o == nil || IsNil(o.Activate) {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *ApplicationLinks) HasActivate() bool {
	if o != nil && !IsNil(o.Activate) {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObjectActivateLink and assigns it to the Activate field.
func (o *ApplicationLinks) SetActivate(v HrefObjectActivateLink) {
	o.Activate = &v
}

// GetAppLinks returns the AppLinks field value if set, zero value otherwise.
func (o *ApplicationLinks) GetAppLinks() []HrefObject {
	if o == nil || IsNil(o.AppLinks) {
		var ret []HrefObject
		return ret
	}
	return o.AppLinks
}

// GetAppLinksOk returns a tuple with the AppLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetAppLinksOk() ([]HrefObject, bool) {
	if o == nil || IsNil(o.AppLinks) {
		return nil, false
	}
	return o.AppLinks, true
}

// HasAppLinks returns a boolean if a field has been set.
func (o *ApplicationLinks) HasAppLinks() bool {
	if o != nil && !IsNil(o.AppLinks) {
		return true
	}

	return false
}

// SetAppLinks gets a reference to the given []HrefObject and assigns it to the AppLinks field.
func (o *ApplicationLinks) SetAppLinks(v []HrefObject) {
	o.AppLinks = v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *ApplicationLinks) GetDeactivate() HrefObjectDeactivateLink {
	if o == nil || IsNil(o.Deactivate) {
		var ret HrefObjectDeactivateLink
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetDeactivateOk() (*HrefObjectDeactivateLink, bool) {
	if o == nil || IsNil(o.Deactivate) {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *ApplicationLinks) HasDeactivate() bool {
	if o != nil && !IsNil(o.Deactivate) {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObjectDeactivateLink and assigns it to the Deactivate field.
func (o *ApplicationLinks) SetDeactivate(v HrefObjectDeactivateLink) {
	o.Deactivate = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ApplicationLinks) GetGroups() GroupsLink {
	if o == nil || IsNil(o.Groups) {
		var ret GroupsLink
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetGroupsOk() (*GroupsLink, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ApplicationLinks) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given GroupsLink and assigns it to the Groups field.
func (o *ApplicationLinks) SetGroups(v GroupsLink) {
	o.Groups = &v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *ApplicationLinks) GetHelp() HelpLink {
	if o == nil || IsNil(o.Help) {
		var ret HelpLink
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetHelpOk() (*HelpLink, bool) {
	if o == nil || IsNil(o.Help) {
		return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *ApplicationLinks) HasHelp() bool {
	if o != nil && !IsNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given HelpLink and assigns it to the Help field.
func (o *ApplicationLinks) SetHelp(v HelpLink) {
	o.Help = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ApplicationLinks) GetLogo() []HrefObject {
	if o == nil || IsNil(o.Logo) {
		var ret []HrefObject
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetLogoOk() ([]HrefObject, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ApplicationLinks) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []HrefObject and assigns it to the Logo field.
func (o *ApplicationLinks) SetLogo(v []HrefObject) {
	o.Logo = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationLinks) GetMetadata() MetadataLink {
	if o == nil || IsNil(o.Metadata) {
		var ret MetadataLink
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetMetadataOk() (*MetadataLink, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationLinks) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given MetadataLink and assigns it to the Metadata field.
func (o *ApplicationLinks) SetMetadata(v MetadataLink) {
	o.Metadata = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ApplicationLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || IsNil(o.Self) {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ApplicationLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *ApplicationLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ApplicationLinks) GetUsers() UsersLink {
	if o == nil || IsNil(o.Users) {
		var ret UsersLink
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationLinks) GetUsersOk() (*UsersLink, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ApplicationLinks) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given UsersLink and assigns it to the Users field.
func (o *ApplicationLinks) SetUsers(v UsersLink) {
	o.Users = &v
}

func (o ApplicationLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPolicy) {
		toSerialize["accessPolicy"] = o.AccessPolicy
	}
	if !IsNil(o.Activate) {
		toSerialize["activate"] = o.Activate
	}
	if !IsNil(o.AppLinks) {
		toSerialize["appLinks"] = o.AppLinks
	}
	if !IsNil(o.Deactivate) {
		toSerialize["deactivate"] = o.Deactivate
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Help) {
		toSerialize["help"] = o.Help
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationLinks) UnmarshalJSON(data []byte) (err error) {
	varApplicationLinks := _ApplicationLinks{}

	err = json.Unmarshal(data, &varApplicationLinks)

	if err != nil {
		return err
	}

	*o = ApplicationLinks(varApplicationLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessPolicy")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "appLinks")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "help")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "self")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
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
