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
	"time"
	"fmt"
)

// AppServiceAccount struct for AppServiceAccount
type AppServiceAccount struct {
	// The key name of the app in the Okta Integration Network (OIN)
	ContainerGlobalName *string `json:"containerGlobalName,omitempty"`
	// The app instance label
	ContainerInstanceName *string `json:"containerInstanceName,omitempty"`
	// The [ORN](/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the relevant resource.  Use the specific app ORN format (`orn:{partition}:idp:{yourOrgId}:apps:{appType}:{appId}`) to identify an Okta app instance in your org.
	ContainerOrn string `json:"containerOrn"`
	// Timestamp when the app service account was created
	Created *time.Time `json:"created,omitempty"`
	// The description of the app service account
	Description *string `json:"description,omitempty"`
	// The UUID of the app service account
	Id *string `json:"id,omitempty" validate:"regexp=(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"`
	// Timestamp when the app service account was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The user-defined name for the app service account
	Name string `json:"name" validate:"regexp=^[\\\\w\\\\-_. ]+$"`
	// A list of IDs of the Okta groups who own the app service account
	OwnerGroupIds []string `json:"ownerGroupIds,omitempty"`
	// A list of IDs of the Okta users who own the app service account
	OwnerUserIds []string `json:"ownerUserIds,omitempty"`
	// The app service account password. Required for apps that don't have provisioning enabled or don't support password synchronization.
	Password *string `json:"password,omitempty"`
	// Describes the current status of an app service account
	Status *string `json:"status,omitempty"`
	// Describes the detailed status of an app service account
	StatusDetail *string `json:"statusDetail,omitempty"`
	// The username that serves as the direct link to your managed app account. Ensure that this value precisely matches the identifier of the target app account.
	Username string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _AppServiceAccount AppServiceAccount

// NewAppServiceAccount instantiates a new AppServiceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppServiceAccount(containerOrn string, name string, username string) *AppServiceAccount {
	this := AppServiceAccount{}
	this.ContainerOrn = containerOrn
	this.Name = name
	this.Username = username
	return &this
}

// NewAppServiceAccountWithDefaults instantiates a new AppServiceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppServiceAccountWithDefaults() *AppServiceAccount {
	this := AppServiceAccount{}
	return &this
}

// GetContainerGlobalName returns the ContainerGlobalName field value if set, zero value otherwise.
func (o *AppServiceAccount) GetContainerGlobalName() string {
	if o == nil || o.ContainerGlobalName == nil {
		var ret string
		return ret
	}
	return *o.ContainerGlobalName
}

// GetContainerGlobalNameOk returns a tuple with the ContainerGlobalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetContainerGlobalNameOk() (*string, bool) {
	if o == nil || o.ContainerGlobalName == nil {
		return nil, false
	}
	return o.ContainerGlobalName, true
}

// HasContainerGlobalName returns a boolean if a field has been set.
func (o *AppServiceAccount) HasContainerGlobalName() bool {
	if o != nil && o.ContainerGlobalName != nil {
		return true
	}

	return false
}

// SetContainerGlobalName gets a reference to the given string and assigns it to the ContainerGlobalName field.
func (o *AppServiceAccount) SetContainerGlobalName(v string) {
	o.ContainerGlobalName = &v
}

// GetContainerInstanceName returns the ContainerInstanceName field value if set, zero value otherwise.
func (o *AppServiceAccount) GetContainerInstanceName() string {
	if o == nil || o.ContainerInstanceName == nil {
		var ret string
		return ret
	}
	return *o.ContainerInstanceName
}

// GetContainerInstanceNameOk returns a tuple with the ContainerInstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetContainerInstanceNameOk() (*string, bool) {
	if o == nil || o.ContainerInstanceName == nil {
		return nil, false
	}
	return o.ContainerInstanceName, true
}

// HasContainerInstanceName returns a boolean if a field has been set.
func (o *AppServiceAccount) HasContainerInstanceName() bool {
	if o != nil && o.ContainerInstanceName != nil {
		return true
	}

	return false
}

// SetContainerInstanceName gets a reference to the given string and assigns it to the ContainerInstanceName field.
func (o *AppServiceAccount) SetContainerInstanceName(v string) {
	o.ContainerInstanceName = &v
}

// GetContainerOrn returns the ContainerOrn field value
func (o *AppServiceAccount) GetContainerOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContainerOrn
}

// GetContainerOrnOk returns a tuple with the ContainerOrn field value
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetContainerOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainerOrn, true
}

// SetContainerOrn sets field value
func (o *AppServiceAccount) SetContainerOrn(v string) {
	o.ContainerOrn = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *AppServiceAccount) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *AppServiceAccount) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *AppServiceAccount) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppServiceAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppServiceAccount) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppServiceAccount) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppServiceAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppServiceAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppServiceAccount) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AppServiceAccount) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AppServiceAccount) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AppServiceAccount) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *AppServiceAccount) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppServiceAccount) SetName(v string) {
	o.Name = v
}

// GetOwnerGroupIds returns the OwnerGroupIds field value if set, zero value otherwise.
func (o *AppServiceAccount) GetOwnerGroupIds() []string {
	if o == nil || o.OwnerGroupIds == nil {
		var ret []string
		return ret
	}
	return o.OwnerGroupIds
}

// GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetOwnerGroupIdsOk() ([]string, bool) {
	if o == nil || o.OwnerGroupIds == nil {
		return nil, false
	}
	return o.OwnerGroupIds, true
}

// HasOwnerGroupIds returns a boolean if a field has been set.
func (o *AppServiceAccount) HasOwnerGroupIds() bool {
	if o != nil && o.OwnerGroupIds != nil {
		return true
	}

	return false
}

// SetOwnerGroupIds gets a reference to the given []string and assigns it to the OwnerGroupIds field.
func (o *AppServiceAccount) SetOwnerGroupIds(v []string) {
	o.OwnerGroupIds = v
}

// GetOwnerUserIds returns the OwnerUserIds field value if set, zero value otherwise.
func (o *AppServiceAccount) GetOwnerUserIds() []string {
	if o == nil || o.OwnerUserIds == nil {
		var ret []string
		return ret
	}
	return o.OwnerUserIds
}

// GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetOwnerUserIdsOk() ([]string, bool) {
	if o == nil || o.OwnerUserIds == nil {
		return nil, false
	}
	return o.OwnerUserIds, true
}

// HasOwnerUserIds returns a boolean if a field has been set.
func (o *AppServiceAccount) HasOwnerUserIds() bool {
	if o != nil && o.OwnerUserIds != nil {
		return true
	}

	return false
}

// SetOwnerUserIds gets a reference to the given []string and assigns it to the OwnerUserIds field.
func (o *AppServiceAccount) SetOwnerUserIds(v []string) {
	o.OwnerUserIds = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *AppServiceAccount) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *AppServiceAccount) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *AppServiceAccount) SetPassword(v string) {
	o.Password = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AppServiceAccount) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AppServiceAccount) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AppServiceAccount) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *AppServiceAccount) GetStatusDetail() string {
	if o == nil || o.StatusDetail == nil {
		var ret string
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetStatusDetailOk() (*string, bool) {
	if o == nil || o.StatusDetail == nil {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *AppServiceAccount) HasStatusDetail() bool {
	if o != nil && o.StatusDetail != nil {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given string and assigns it to the StatusDetail field.
func (o *AppServiceAccount) SetStatusDetail(v string) {
	o.StatusDetail = &v
}

// GetUsername returns the Username field value
func (o *AppServiceAccount) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AppServiceAccount) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AppServiceAccount) SetUsername(v string) {
	o.Username = v
}

func (o AppServiceAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerGlobalName != nil {
		toSerialize["containerGlobalName"] = o.ContainerGlobalName
	}
	if o.ContainerInstanceName != nil {
		toSerialize["containerInstanceName"] = o.ContainerInstanceName
	}
	if true {
		toSerialize["containerOrn"] = o.ContainerOrn
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OwnerGroupIds != nil {
		toSerialize["ownerGroupIds"] = o.OwnerGroupIds
	}
	if o.OwnerUserIds != nil {
		toSerialize["ownerUserIds"] = o.OwnerUserIds
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusDetail != nil {
		toSerialize["statusDetail"] = o.StatusDetail
	}
	if true {
		toSerialize["username"] = o.Username
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AppServiceAccount) UnmarshalJSON(bytes []byte) (err error) {
	varAppServiceAccount := _AppServiceAccount{}

	err = json.Unmarshal(bytes, &varAppServiceAccount)
	if err == nil {
		*o = AppServiceAccount(varAppServiceAccount)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "containerGlobalName")
		delete(additionalProperties, "containerInstanceName")
		delete(additionalProperties, "containerOrn")
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ownerGroupIds")
		delete(additionalProperties, "ownerUserIds")
		delete(additionalProperties, "password")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetail")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAppServiceAccount struct {
	value *AppServiceAccount
	isSet bool
}

func (v NullableAppServiceAccount) Get() *AppServiceAccount {
	return v.value
}

func (v *NullableAppServiceAccount) Set(val *AppServiceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAppServiceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAppServiceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppServiceAccount(val *AppServiceAccount) *NullableAppServiceAccount {
	return &NullableAppServiceAccount{value: val, isSet: true}
}

func (v NullableAppServiceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppServiceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

