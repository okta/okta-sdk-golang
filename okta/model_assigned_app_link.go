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

// checks if the AssignedAppLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignedAppLink{}

// AssignedAppLink struct for AssignedAppLink
type AssignedAppLink struct {
	AppAssignmentId      *string `json:"appAssignmentId,omitempty"`
	AppInstanceId        *string `json:"appInstanceId,omitempty"`
	AppName              *string `json:"appName,omitempty"`
	CredentialsSetup     *bool   `json:"credentialsSetup,omitempty"`
	Hidden               *bool   `json:"hidden,omitempty"`
	Id                   *string `json:"id,omitempty"`
	Label                *string `json:"label,omitempty"`
	LinkUrl              *string `json:"linkUrl,omitempty"`
	LogoUrl              *string `json:"logoUrl,omitempty"`
	SortOrder            *int32  `json:"sortOrder,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssignedAppLink AssignedAppLink

// NewAssignedAppLink instantiates a new AssignedAppLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedAppLink() *AssignedAppLink {
	this := AssignedAppLink{}
	return &this
}

// NewAssignedAppLinkWithDefaults instantiates a new AssignedAppLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedAppLinkWithDefaults() *AssignedAppLink {
	this := AssignedAppLink{}
	return &this
}

// GetAppAssignmentId returns the AppAssignmentId field value if set, zero value otherwise.
func (o *AssignedAppLink) GetAppAssignmentId() string {
	if o == nil || IsNil(o.AppAssignmentId) {
		var ret string
		return ret
	}
	return *o.AppAssignmentId
}

// GetAppAssignmentIdOk returns a tuple with the AppAssignmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetAppAssignmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppAssignmentId) {
		return nil, false
	}
	return o.AppAssignmentId, true
}

// HasAppAssignmentId returns a boolean if a field has been set.
func (o *AssignedAppLink) HasAppAssignmentId() bool {
	if o != nil && !IsNil(o.AppAssignmentId) {
		return true
	}

	return false
}

// SetAppAssignmentId gets a reference to the given string and assigns it to the AppAssignmentId field.
func (o *AssignedAppLink) SetAppAssignmentId(v string) {
	o.AppAssignmentId = &v
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *AssignedAppLink) GetAppInstanceId() string {
	if o == nil || IsNil(o.AppInstanceId) {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceId) {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *AssignedAppLink) HasAppInstanceId() bool {
	if o != nil && !IsNil(o.AppInstanceId) {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *AssignedAppLink) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *AssignedAppLink) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *AssignedAppLink) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *AssignedAppLink) SetAppName(v string) {
	o.AppName = &v
}

// GetCredentialsSetup returns the CredentialsSetup field value if set, zero value otherwise.
func (o *AssignedAppLink) GetCredentialsSetup() bool {
	if o == nil || IsNil(o.CredentialsSetup) {
		var ret bool
		return ret
	}
	return *o.CredentialsSetup
}

// GetCredentialsSetupOk returns a tuple with the CredentialsSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetCredentialsSetupOk() (*bool, bool) {
	if o == nil || IsNil(o.CredentialsSetup) {
		return nil, false
	}
	return o.CredentialsSetup, true
}

// HasCredentialsSetup returns a boolean if a field has been set.
func (o *AssignedAppLink) HasCredentialsSetup() bool {
	if o != nil && !IsNil(o.CredentialsSetup) {
		return true
	}

	return false
}

// SetCredentialsSetup gets a reference to the given bool and assigns it to the CredentialsSetup field.
func (o *AssignedAppLink) SetCredentialsSetup(v bool) {
	o.CredentialsSetup = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *AssignedAppLink) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *AssignedAppLink) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *AssignedAppLink) SetHidden(v bool) {
	o.Hidden = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssignedAppLink) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssignedAppLink) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssignedAppLink) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AssignedAppLink) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AssignedAppLink) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AssignedAppLink) SetLabel(v string) {
	o.Label = &v
}

// GetLinkUrl returns the LinkUrl field value if set, zero value otherwise.
func (o *AssignedAppLink) GetLinkUrl() string {
	if o == nil || IsNil(o.LinkUrl) {
		var ret string
		return ret
	}
	return *o.LinkUrl
}

// GetLinkUrlOk returns a tuple with the LinkUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetLinkUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LinkUrl) {
		return nil, false
	}
	return o.LinkUrl, true
}

// HasLinkUrl returns a boolean if a field has been set.
func (o *AssignedAppLink) HasLinkUrl() bool {
	if o != nil && !IsNil(o.LinkUrl) {
		return true
	}

	return false
}

// SetLinkUrl gets a reference to the given string and assigns it to the LinkUrl field.
func (o *AssignedAppLink) SetLinkUrl(v string) {
	o.LinkUrl = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *AssignedAppLink) GetLogoUrl() string {
	if o == nil || IsNil(o.LogoUrl) {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetLogoUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoUrl) {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *AssignedAppLink) HasLogoUrl() bool {
	if o != nil && !IsNil(o.LogoUrl) {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *AssignedAppLink) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *AssignedAppLink) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedAppLink) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *AssignedAppLink) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *AssignedAppLink) SetSortOrder(v int32) {
	o.SortOrder = &v
}

func (o AssignedAppLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignedAppLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppAssignmentId) {
		toSerialize["appAssignmentId"] = o.AppAssignmentId
	}
	if !IsNil(o.AppInstanceId) {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.CredentialsSetup) {
		toSerialize["credentialsSetup"] = o.CredentialsSetup
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.LinkUrl) {
		toSerialize["linkUrl"] = o.LinkUrl
	}
	if !IsNil(o.LogoUrl) {
		toSerialize["logoUrl"] = o.LogoUrl
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssignedAppLink) UnmarshalJSON(data []byte) (err error) {
	varAssignedAppLink := _AssignedAppLink{}

	err = json.Unmarshal(data, &varAssignedAppLink)

	if err != nil {
		return err
	}

	*o = AssignedAppLink(varAssignedAppLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appAssignmentId")
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "appName")
		delete(additionalProperties, "credentialsSetup")
		delete(additionalProperties, "hidden")
		delete(additionalProperties, "id")
		delete(additionalProperties, "label")
		delete(additionalProperties, "linkUrl")
		delete(additionalProperties, "logoUrl")
		delete(additionalProperties, "sortOrder")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssignedAppLink struct {
	value *AssignedAppLink
	isSet bool
}

func (v NullableAssignedAppLink) Get() *AssignedAppLink {
	return v.value
}

func (v *NullableAssignedAppLink) Set(val *AssignedAppLink) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedAppLink) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedAppLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedAppLink(val *AssignedAppLink) *NullableAssignedAppLink {
	return &NullableAssignedAppLink{value: val, isSet: true}
}

func (v NullableAssignedAppLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedAppLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
