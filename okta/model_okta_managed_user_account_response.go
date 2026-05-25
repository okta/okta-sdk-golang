/*
Okta Admin Management API

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
	"fmt"
	"time"
)

// checks if the OktaManagedUserAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OktaManagedUserAccountResponse{}

// OktaManagedUserAccountResponse An Okta managed user account representing a Universal Directory user managed as a service account
type OktaManagedUserAccountResponse struct {
	// Timestamp when the Okta managed user account was created
	Created *time.Time `json:"created,omitempty"`
	// The description of the Okta managed user account
	Description *string `json:"description,omitempty"`
	// The email address associated with the Okta user. This parameter is read-only, and it is derived from the Okta user profile.
	Email string `json:"email"`
	// The UUID of the Okta managed user account
	Id string `json:"id" validate:"regexp=(?i)^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"`
	// Timestamp when the Okta managed user account was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The user-defined name for the Okta managed user account
	Name string `json:"name" validate:"regexp=^[\\\\w\\\\-_. ]+$"`
	// The ID of the Okta user being managed as a service account
	OktaUserId string `json:"oktaUserId"`
	// A list of IDs of the Okta groups who own the Okta managed user account
	OwnerGroupIds []string `json:"ownerGroupIds,omitempty"`
	// A list of IDs of the Okta users who own the Okta managed user account
	OwnerUserIds []string `json:"ownerUserIds,omitempty"`
	// Describes the current status of a service account
	Status *string `json:"status,omitempty"`
	// Describes the detailed status of a service account
	StatusDetail *string `json:"statusDetail,omitempty"`
	// The username associated with the Okta user. This parameter is read-only, and it is derived from the Okta user profile.
	Username             string `json:"username"`
	AdditionalProperties map[string]interface{}
}

type _OktaManagedUserAccountResponse OktaManagedUserAccountResponse

// NewOktaManagedUserAccountResponse instantiates a new OktaManagedUserAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOktaManagedUserAccountResponse(email string, id string, name string, oktaUserId string, username string) *OktaManagedUserAccountResponse {
	this := OktaManagedUserAccountResponse{}
	this.Email = email
	this.Id = id
	this.Name = name
	this.OktaUserId = oktaUserId
	this.Username = username
	return &this
}

// NewOktaManagedUserAccountResponseWithDefaults instantiates a new OktaManagedUserAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOktaManagedUserAccountResponseWithDefaults() *OktaManagedUserAccountResponse {
	this := OktaManagedUserAccountResponse{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OktaManagedUserAccountResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OktaManagedUserAccountResponse) SetDescription(v string) {
	o.Description = &v
}

// GetEmail returns the Email field value
func (o *OktaManagedUserAccountResponse) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *OktaManagedUserAccountResponse) SetEmail(v string) {
	o.Email = v
}

// GetId returns the Id field value
func (o *OktaManagedUserAccountResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OktaManagedUserAccountResponse) SetId(v string) {
	o.Id = v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OktaManagedUserAccountResponse) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *OktaManagedUserAccountResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OktaManagedUserAccountResponse) SetName(v string) {
	o.Name = v
}

// GetOktaUserId returns the OktaUserId field value
func (o *OktaManagedUserAccountResponse) GetOktaUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OktaUserId
}

// GetOktaUserIdOk returns a tuple with the OktaUserId field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetOktaUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OktaUserId, true
}

// SetOktaUserId sets field value
func (o *OktaManagedUserAccountResponse) SetOktaUserId(v string) {
	o.OktaUserId = v
}

// GetOwnerGroupIds returns the OwnerGroupIds field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetOwnerGroupIds() []string {
	if o == nil || IsNil(o.OwnerGroupIds) {
		var ret []string
		return ret
	}
	return o.OwnerGroupIds
}

// GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetOwnerGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerGroupIds) {
		return nil, false
	}
	return o.OwnerGroupIds, true
}

// HasOwnerGroupIds returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasOwnerGroupIds() bool {
	if o != nil && !IsNil(o.OwnerGroupIds) {
		return true
	}

	return false
}

// SetOwnerGroupIds gets a reference to the given []string and assigns it to the OwnerGroupIds field.
func (o *OktaManagedUserAccountResponse) SetOwnerGroupIds(v []string) {
	o.OwnerGroupIds = v
}

// GetOwnerUserIds returns the OwnerUserIds field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetOwnerUserIds() []string {
	if o == nil || IsNil(o.OwnerUserIds) {
		var ret []string
		return ret
	}
	return o.OwnerUserIds
}

// GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetOwnerUserIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OwnerUserIds) {
		return nil, false
	}
	return o.OwnerUserIds, true
}

// HasOwnerUserIds returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasOwnerUserIds() bool {
	if o != nil && !IsNil(o.OwnerUserIds) {
		return true
	}

	return false
}

// SetOwnerUserIds gets a reference to the given []string and assigns it to the OwnerUserIds field.
func (o *OktaManagedUserAccountResponse) SetOwnerUserIds(v []string) {
	o.OwnerUserIds = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OktaManagedUserAccountResponse) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDetail returns the StatusDetail field value if set, zero value otherwise.
func (o *OktaManagedUserAccountResponse) GetStatusDetail() string {
	if o == nil || IsNil(o.StatusDetail) {
		var ret string
		return ret
	}
	return *o.StatusDetail
}

// GetStatusDetailOk returns a tuple with the StatusDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetStatusDetailOk() (*string, bool) {
	if o == nil || IsNil(o.StatusDetail) {
		return nil, false
	}
	return o.StatusDetail, true
}

// HasStatusDetail returns a boolean if a field has been set.
func (o *OktaManagedUserAccountResponse) HasStatusDetail() bool {
	if o != nil && !IsNil(o.StatusDetail) {
		return true
	}

	return false
}

// SetStatusDetail gets a reference to the given string and assigns it to the StatusDetail field.
func (o *OktaManagedUserAccountResponse) SetStatusDetail(v string) {
	o.StatusDetail = &v
}

// GetUsername returns the Username field value
func (o *OktaManagedUserAccountResponse) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *OktaManagedUserAccountResponse) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *OktaManagedUserAccountResponse) SetUsername(v string) {
	o.Username = v
}

func (o OktaManagedUserAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OktaManagedUserAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["email"] = o.Email
	toSerialize["id"] = o.Id
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	toSerialize["name"] = o.Name
	toSerialize["oktaUserId"] = o.OktaUserId
	if !IsNil(o.OwnerGroupIds) {
		toSerialize["ownerGroupIds"] = o.OwnerGroupIds
	}
	if !IsNil(o.OwnerUserIds) {
		toSerialize["ownerUserIds"] = o.OwnerUserIds
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusDetail) {
		toSerialize["statusDetail"] = o.StatusDetail
	}
	toSerialize["username"] = o.Username

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OktaManagedUserAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"id",
		"name",
		"oktaUserId",
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOktaManagedUserAccountResponse := _OktaManagedUserAccountResponse{}

	err = json.Unmarshal(data, &varOktaManagedUserAccountResponse)

	if err != nil {
		return err
	}

	*o = OktaManagedUserAccountResponse(varOktaManagedUserAccountResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "description")
		delete(additionalProperties, "email")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "oktaUserId")
		delete(additionalProperties, "ownerGroupIds")
		delete(additionalProperties, "ownerUserIds")
		delete(additionalProperties, "status")
		delete(additionalProperties, "statusDetail")
		delete(additionalProperties, "username")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOktaManagedUserAccountResponse struct {
	value *OktaManagedUserAccountResponse
	isSet bool
}

func (v NullableOktaManagedUserAccountResponse) Get() *OktaManagedUserAccountResponse {
	return v.value
}

func (v *NullableOktaManagedUserAccountResponse) Set(val *OktaManagedUserAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOktaManagedUserAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOktaManagedUserAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOktaManagedUserAccountResponse(val *OktaManagedUserAccountResponse) *NullableOktaManagedUserAccountResponse {
	return &NullableOktaManagedUserAccountResponse{value: val, isSet: true}
}

func (v NullableOktaManagedUserAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOktaManagedUserAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
