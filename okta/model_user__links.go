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

// UserLinks Specifies link relations (see [Web Linking](https://datatracker.ietf.org/doc/html/rfc8288) available for the current status of a user. The Links object is used for dynamic discovery of related resources, lifecycle operations, and credential operations. The Links object is read-only.  For an individual user result, the Links object contains a full set of link relations available for that user as determined by your policies. For a collection of users, the Links object contains only the `self` link. Operations that return a collection of Users include List Users and List Group Members.
type UserLinks struct {
	Self *HrefObject `json:"self,omitempty"`
	Activate *HrefObject `json:"activate,omitempty"`
	ResetPassword *HrefObject `json:"resetPassword,omitempty"`
	ResetFactors *HrefObject `json:"resetFactors,omitempty"`
	ExpirePassword *HrefObject `json:"expirePassword,omitempty"`
	ForgotPassword *HrefObject `json:"forgotPassword,omitempty"`
	ChangeRecoveryQuestion *HrefObject `json:"changeRecoveryQuestion,omitempty"`
	Deactivate *HrefObject `json:"deactivate,omitempty"`
	Reactivate *HrefObject `json:"reactivate,omitempty"`
	ChangePassword *HrefObject `json:"changePassword,omitempty"`
	Schema *HrefObject `json:"schema,omitempty"`
	Suspend *HrefObject `json:"suspend,omitempty"`
	Unsuspend *HrefObject `json:"unsuspend,omitempty"`
	Unlock *HrefObject `json:"unlock,omitempty"`
	Type *HrefObject `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserLinks UserLinks

// NewUserLinks instantiates a new UserLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserLinks() *UserLinks {
	this := UserLinks{}
	return &this
}

// NewUserLinksWithDefaults instantiates a new UserLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserLinksWithDefaults() *UserLinks {
	this := UserLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *UserLinks) GetSelf() HrefObject {
	if o == nil || o.Self == nil {
		var ret HrefObject
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetSelfOk() (*HrefObject, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *UserLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObject and assigns it to the Self field.
func (o *UserLinks) SetSelf(v HrefObject) {
	o.Self = &v
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *UserLinks) GetActivate() HrefObject {
	if o == nil || o.Activate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetActivateOk() (*HrefObject, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *UserLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given HrefObject and assigns it to the Activate field.
func (o *UserLinks) SetActivate(v HrefObject) {
	o.Activate = &v
}

// GetResetPassword returns the ResetPassword field value if set, zero value otherwise.
func (o *UserLinks) GetResetPassword() HrefObject {
	if o == nil || o.ResetPassword == nil {
		var ret HrefObject
		return ret
	}
	return *o.ResetPassword
}

// GetResetPasswordOk returns a tuple with the ResetPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetResetPasswordOk() (*HrefObject, bool) {
	if o == nil || o.ResetPassword == nil {
		return nil, false
	}
	return o.ResetPassword, true
}

// HasResetPassword returns a boolean if a field has been set.
func (o *UserLinks) HasResetPassword() bool {
	if o != nil && o.ResetPassword != nil {
		return true
	}

	return false
}

// SetResetPassword gets a reference to the given HrefObject and assigns it to the ResetPassword field.
func (o *UserLinks) SetResetPassword(v HrefObject) {
	o.ResetPassword = &v
}

// GetResetFactors returns the ResetFactors field value if set, zero value otherwise.
func (o *UserLinks) GetResetFactors() HrefObject {
	if o == nil || o.ResetFactors == nil {
		var ret HrefObject
		return ret
	}
	return *o.ResetFactors
}

// GetResetFactorsOk returns a tuple with the ResetFactors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetResetFactorsOk() (*HrefObject, bool) {
	if o == nil || o.ResetFactors == nil {
		return nil, false
	}
	return o.ResetFactors, true
}

// HasResetFactors returns a boolean if a field has been set.
func (o *UserLinks) HasResetFactors() bool {
	if o != nil && o.ResetFactors != nil {
		return true
	}

	return false
}

// SetResetFactors gets a reference to the given HrefObject and assigns it to the ResetFactors field.
func (o *UserLinks) SetResetFactors(v HrefObject) {
	o.ResetFactors = &v
}

// GetExpirePassword returns the ExpirePassword field value if set, zero value otherwise.
func (o *UserLinks) GetExpirePassword() HrefObject {
	if o == nil || o.ExpirePassword == nil {
		var ret HrefObject
		return ret
	}
	return *o.ExpirePassword
}

// GetExpirePasswordOk returns a tuple with the ExpirePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetExpirePasswordOk() (*HrefObject, bool) {
	if o == nil || o.ExpirePassword == nil {
		return nil, false
	}
	return o.ExpirePassword, true
}

// HasExpirePassword returns a boolean if a field has been set.
func (o *UserLinks) HasExpirePassword() bool {
	if o != nil && o.ExpirePassword != nil {
		return true
	}

	return false
}

// SetExpirePassword gets a reference to the given HrefObject and assigns it to the ExpirePassword field.
func (o *UserLinks) SetExpirePassword(v HrefObject) {
	o.ExpirePassword = &v
}

// GetForgotPassword returns the ForgotPassword field value if set, zero value otherwise.
func (o *UserLinks) GetForgotPassword() HrefObject {
	if o == nil || o.ForgotPassword == nil {
		var ret HrefObject
		return ret
	}
	return *o.ForgotPassword
}

// GetForgotPasswordOk returns a tuple with the ForgotPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetForgotPasswordOk() (*HrefObject, bool) {
	if o == nil || o.ForgotPassword == nil {
		return nil, false
	}
	return o.ForgotPassword, true
}

// HasForgotPassword returns a boolean if a field has been set.
func (o *UserLinks) HasForgotPassword() bool {
	if o != nil && o.ForgotPassword != nil {
		return true
	}

	return false
}

// SetForgotPassword gets a reference to the given HrefObject and assigns it to the ForgotPassword field.
func (o *UserLinks) SetForgotPassword(v HrefObject) {
	o.ForgotPassword = &v
}

// GetChangeRecoveryQuestion returns the ChangeRecoveryQuestion field value if set, zero value otherwise.
func (o *UserLinks) GetChangeRecoveryQuestion() HrefObject {
	if o == nil || o.ChangeRecoveryQuestion == nil {
		var ret HrefObject
		return ret
	}
	return *o.ChangeRecoveryQuestion
}

// GetChangeRecoveryQuestionOk returns a tuple with the ChangeRecoveryQuestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetChangeRecoveryQuestionOk() (*HrefObject, bool) {
	if o == nil || o.ChangeRecoveryQuestion == nil {
		return nil, false
	}
	return o.ChangeRecoveryQuestion, true
}

// HasChangeRecoveryQuestion returns a boolean if a field has been set.
func (o *UserLinks) HasChangeRecoveryQuestion() bool {
	if o != nil && o.ChangeRecoveryQuestion != nil {
		return true
	}

	return false
}

// SetChangeRecoveryQuestion gets a reference to the given HrefObject and assigns it to the ChangeRecoveryQuestion field.
func (o *UserLinks) SetChangeRecoveryQuestion(v HrefObject) {
	o.ChangeRecoveryQuestion = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *UserLinks) GetDeactivate() HrefObject {
	if o == nil || o.Deactivate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetDeactivateOk() (*HrefObject, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *UserLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given HrefObject and assigns it to the Deactivate field.
func (o *UserLinks) SetDeactivate(v HrefObject) {
	o.Deactivate = &v
}

// GetReactivate returns the Reactivate field value if set, zero value otherwise.
func (o *UserLinks) GetReactivate() HrefObject {
	if o == nil || o.Reactivate == nil {
		var ret HrefObject
		return ret
	}
	return *o.Reactivate
}

// GetReactivateOk returns a tuple with the Reactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetReactivateOk() (*HrefObject, bool) {
	if o == nil || o.Reactivate == nil {
		return nil, false
	}
	return o.Reactivate, true
}

// HasReactivate returns a boolean if a field has been set.
func (o *UserLinks) HasReactivate() bool {
	if o != nil && o.Reactivate != nil {
		return true
	}

	return false
}

// SetReactivate gets a reference to the given HrefObject and assigns it to the Reactivate field.
func (o *UserLinks) SetReactivate(v HrefObject) {
	o.Reactivate = &v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *UserLinks) GetChangePassword() HrefObject {
	if o == nil || o.ChangePassword == nil {
		var ret HrefObject
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetChangePasswordOk() (*HrefObject, bool) {
	if o == nil || o.ChangePassword == nil {
		return nil, false
	}
	return o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *UserLinks) HasChangePassword() bool {
	if o != nil && o.ChangePassword != nil {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given HrefObject and assigns it to the ChangePassword field.
func (o *UserLinks) SetChangePassword(v HrefObject) {
	o.ChangePassword = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *UserLinks) GetSchema() HrefObject {
	if o == nil || o.Schema == nil {
		var ret HrefObject
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetSchemaOk() (*HrefObject, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *UserLinks) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given HrefObject and assigns it to the Schema field.
func (o *UserLinks) SetSchema(v HrefObject) {
	o.Schema = &v
}

// GetSuspend returns the Suspend field value if set, zero value otherwise.
func (o *UserLinks) GetSuspend() HrefObject {
	if o == nil || o.Suspend == nil {
		var ret HrefObject
		return ret
	}
	return *o.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetSuspendOk() (*HrefObject, bool) {
	if o == nil || o.Suspend == nil {
		return nil, false
	}
	return o.Suspend, true
}

// HasSuspend returns a boolean if a field has been set.
func (o *UserLinks) HasSuspend() bool {
	if o != nil && o.Suspend != nil {
		return true
	}

	return false
}

// SetSuspend gets a reference to the given HrefObject and assigns it to the Suspend field.
func (o *UserLinks) SetSuspend(v HrefObject) {
	o.Suspend = &v
}

// GetUnsuspend returns the Unsuspend field value if set, zero value otherwise.
func (o *UserLinks) GetUnsuspend() HrefObject {
	if o == nil || o.Unsuspend == nil {
		var ret HrefObject
		return ret
	}
	return *o.Unsuspend
}

// GetUnsuspendOk returns a tuple with the Unsuspend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetUnsuspendOk() (*HrefObject, bool) {
	if o == nil || o.Unsuspend == nil {
		return nil, false
	}
	return o.Unsuspend, true
}

// HasUnsuspend returns a boolean if a field has been set.
func (o *UserLinks) HasUnsuspend() bool {
	if o != nil && o.Unsuspend != nil {
		return true
	}

	return false
}

// SetUnsuspend gets a reference to the given HrefObject and assigns it to the Unsuspend field.
func (o *UserLinks) SetUnsuspend(v HrefObject) {
	o.Unsuspend = &v
}

// GetUnlock returns the Unlock field value if set, zero value otherwise.
func (o *UserLinks) GetUnlock() HrefObject {
	if o == nil || o.Unlock == nil {
		var ret HrefObject
		return ret
	}
	return *o.Unlock
}

// GetUnlockOk returns a tuple with the Unlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetUnlockOk() (*HrefObject, bool) {
	if o == nil || o.Unlock == nil {
		return nil, false
	}
	return o.Unlock, true
}

// HasUnlock returns a boolean if a field has been set.
func (o *UserLinks) HasUnlock() bool {
	if o != nil && o.Unlock != nil {
		return true
	}

	return false
}

// SetUnlock gets a reference to the given HrefObject and assigns it to the Unlock field.
func (o *UserLinks) SetUnlock(v HrefObject) {
	o.Unlock = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserLinks) GetType() HrefObject {
	if o == nil || o.Type == nil {
		var ret HrefObject
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserLinks) GetTypeOk() (*HrefObject, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserLinks) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given HrefObject and assigns it to the Type field.
func (o *UserLinks) SetType(v HrefObject) {
	o.Type = &v
}

func (o UserLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.ResetPassword != nil {
		toSerialize["resetPassword"] = o.ResetPassword
	}
	if o.ResetFactors != nil {
		toSerialize["resetFactors"] = o.ResetFactors
	}
	if o.ExpirePassword != nil {
		toSerialize["expirePassword"] = o.ExpirePassword
	}
	if o.ForgotPassword != nil {
		toSerialize["forgotPassword"] = o.ForgotPassword
	}
	if o.ChangeRecoveryQuestion != nil {
		toSerialize["changeRecoveryQuestion"] = o.ChangeRecoveryQuestion
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Reactivate != nil {
		toSerialize["reactivate"] = o.Reactivate
	}
	if o.ChangePassword != nil {
		toSerialize["changePassword"] = o.ChangePassword
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Suspend != nil {
		toSerialize["suspend"] = o.Suspend
	}
	if o.Unsuspend != nil {
		toSerialize["unsuspend"] = o.Unsuspend
	}
	if o.Unlock != nil {
		toSerialize["unlock"] = o.Unlock
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserLinks := _UserLinks{}

	err = json.Unmarshal(bytes, &varUserLinks)
	if err == nil {
		*o = UserLinks(varUserLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "activate")
		delete(additionalProperties, "resetPassword")
		delete(additionalProperties, "resetFactors")
		delete(additionalProperties, "expirePassword")
		delete(additionalProperties, "forgotPassword")
		delete(additionalProperties, "changeRecoveryQuestion")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "reactivate")
		delete(additionalProperties, "changePassword")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "suspend")
		delete(additionalProperties, "unsuspend")
		delete(additionalProperties, "unlock")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserLinks struct {
	value *UserLinks
	isSet bool
}

func (v NullableUserLinks) Get() *UserLinks {
	return v.value
}

func (v *NullableUserLinks) Set(val *UserLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserLinks(val *UserLinks) *NullableUserLinks {
	return &NullableUserLinks{value: val, isSet: true}
}

func (v NullableUserLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

