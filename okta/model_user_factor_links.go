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

// UserFactorLinks struct for UserFactorLinks
type UserFactorLinks struct {
	Activate *LinksActivateActivate `json:"activate,omitempty"`
	Cancel *LinksCancelCancel `json:"cancel,omitempty"`
	Deactivate *LinksDeactivateDeactivate `json:"deactivate,omitempty"`
	Enroll *LinksEnrollEnroll `json:"enroll,omitempty"`
	Factor *LinksFactorFactor `json:"factor,omitempty"`
	Poll *LinksPollPoll `json:"poll,omitempty"`
	Qrcode *LinksQrcodeQrcode `json:"qrcode,omitempty"`
	Question *LinksQuestionsQuestion `json:"question,omitempty"`
	Resend *LinksResendResend `json:"resend,omitempty"`
	Send *LinksSendSend `json:"send,omitempty"`
	Self *HrefObjectSelfLink `json:"self,omitempty"`
	User *LinksUserUser `json:"user,omitempty"`
	Verify *LinksVerifyVerify `json:"verify,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorLinks UserFactorLinks

// NewUserFactorLinks instantiates a new UserFactorLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorLinks() *UserFactorLinks {
	this := UserFactorLinks{}
	return &this
}

// NewUserFactorLinksWithDefaults instantiates a new UserFactorLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorLinksWithDefaults() *UserFactorLinks {
	this := UserFactorLinks{}
	return &this
}

// GetActivate returns the Activate field value if set, zero value otherwise.
func (o *UserFactorLinks) GetActivate() LinksActivateActivate {
	if o == nil || o.Activate == nil {
		var ret LinksActivateActivate
		return ret
	}
	return *o.Activate
}

// GetActivateOk returns a tuple with the Activate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetActivateOk() (*LinksActivateActivate, bool) {
	if o == nil || o.Activate == nil {
		return nil, false
	}
	return o.Activate, true
}

// HasActivate returns a boolean if a field has been set.
func (o *UserFactorLinks) HasActivate() bool {
	if o != nil && o.Activate != nil {
		return true
	}

	return false
}

// SetActivate gets a reference to the given LinksActivateActivate and assigns it to the Activate field.
func (o *UserFactorLinks) SetActivate(v LinksActivateActivate) {
	o.Activate = &v
}

// GetCancel returns the Cancel field value if set, zero value otherwise.
func (o *UserFactorLinks) GetCancel() LinksCancelCancel {
	if o == nil || o.Cancel == nil {
		var ret LinksCancelCancel
		return ret
	}
	return *o.Cancel
}

// GetCancelOk returns a tuple with the Cancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetCancelOk() (*LinksCancelCancel, bool) {
	if o == nil || o.Cancel == nil {
		return nil, false
	}
	return o.Cancel, true
}

// HasCancel returns a boolean if a field has been set.
func (o *UserFactorLinks) HasCancel() bool {
	if o != nil && o.Cancel != nil {
		return true
	}

	return false
}

// SetCancel gets a reference to the given LinksCancelCancel and assigns it to the Cancel field.
func (o *UserFactorLinks) SetCancel(v LinksCancelCancel) {
	o.Cancel = &v
}

// GetDeactivate returns the Deactivate field value if set, zero value otherwise.
func (o *UserFactorLinks) GetDeactivate() LinksDeactivateDeactivate {
	if o == nil || o.Deactivate == nil {
		var ret LinksDeactivateDeactivate
		return ret
	}
	return *o.Deactivate
}

// GetDeactivateOk returns a tuple with the Deactivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetDeactivateOk() (*LinksDeactivateDeactivate, bool) {
	if o == nil || o.Deactivate == nil {
		return nil, false
	}
	return o.Deactivate, true
}

// HasDeactivate returns a boolean if a field has been set.
func (o *UserFactorLinks) HasDeactivate() bool {
	if o != nil && o.Deactivate != nil {
		return true
	}

	return false
}

// SetDeactivate gets a reference to the given LinksDeactivateDeactivate and assigns it to the Deactivate field.
func (o *UserFactorLinks) SetDeactivate(v LinksDeactivateDeactivate) {
	o.Deactivate = &v
}

// GetEnroll returns the Enroll field value if set, zero value otherwise.
func (o *UserFactorLinks) GetEnroll() LinksEnrollEnroll {
	if o == nil || o.Enroll == nil {
		var ret LinksEnrollEnroll
		return ret
	}
	return *o.Enroll
}

// GetEnrollOk returns a tuple with the Enroll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetEnrollOk() (*LinksEnrollEnroll, bool) {
	if o == nil || o.Enroll == nil {
		return nil, false
	}
	return o.Enroll, true
}

// HasEnroll returns a boolean if a field has been set.
func (o *UserFactorLinks) HasEnroll() bool {
	if o != nil && o.Enroll != nil {
		return true
	}

	return false
}

// SetEnroll gets a reference to the given LinksEnrollEnroll and assigns it to the Enroll field.
func (o *UserFactorLinks) SetEnroll(v LinksEnrollEnroll) {
	o.Enroll = &v
}

// GetFactor returns the Factor field value if set, zero value otherwise.
func (o *UserFactorLinks) GetFactor() LinksFactorFactor {
	if o == nil || o.Factor == nil {
		var ret LinksFactorFactor
		return ret
	}
	return *o.Factor
}

// GetFactorOk returns a tuple with the Factor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetFactorOk() (*LinksFactorFactor, bool) {
	if o == nil || o.Factor == nil {
		return nil, false
	}
	return o.Factor, true
}

// HasFactor returns a boolean if a field has been set.
func (o *UserFactorLinks) HasFactor() bool {
	if o != nil && o.Factor != nil {
		return true
	}

	return false
}

// SetFactor gets a reference to the given LinksFactorFactor and assigns it to the Factor field.
func (o *UserFactorLinks) SetFactor(v LinksFactorFactor) {
	o.Factor = &v
}

// GetPoll returns the Poll field value if set, zero value otherwise.
func (o *UserFactorLinks) GetPoll() LinksPollPoll {
	if o == nil || o.Poll == nil {
		var ret LinksPollPoll
		return ret
	}
	return *o.Poll
}

// GetPollOk returns a tuple with the Poll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetPollOk() (*LinksPollPoll, bool) {
	if o == nil || o.Poll == nil {
		return nil, false
	}
	return o.Poll, true
}

// HasPoll returns a boolean if a field has been set.
func (o *UserFactorLinks) HasPoll() bool {
	if o != nil && o.Poll != nil {
		return true
	}

	return false
}

// SetPoll gets a reference to the given LinksPollPoll and assigns it to the Poll field.
func (o *UserFactorLinks) SetPoll(v LinksPollPoll) {
	o.Poll = &v
}

// GetQrcode returns the Qrcode field value if set, zero value otherwise.
func (o *UserFactorLinks) GetQrcode() LinksQrcodeQrcode {
	if o == nil || o.Qrcode == nil {
		var ret LinksQrcodeQrcode
		return ret
	}
	return *o.Qrcode
}

// GetQrcodeOk returns a tuple with the Qrcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetQrcodeOk() (*LinksQrcodeQrcode, bool) {
	if o == nil || o.Qrcode == nil {
		return nil, false
	}
	return o.Qrcode, true
}

// HasQrcode returns a boolean if a field has been set.
func (o *UserFactorLinks) HasQrcode() bool {
	if o != nil && o.Qrcode != nil {
		return true
	}

	return false
}

// SetQrcode gets a reference to the given LinksQrcodeQrcode and assigns it to the Qrcode field.
func (o *UserFactorLinks) SetQrcode(v LinksQrcodeQrcode) {
	o.Qrcode = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *UserFactorLinks) GetQuestion() LinksQuestionsQuestion {
	if o == nil || o.Question == nil {
		var ret LinksQuestionsQuestion
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetQuestionOk() (*LinksQuestionsQuestion, bool) {
	if o == nil || o.Question == nil {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *UserFactorLinks) HasQuestion() bool {
	if o != nil && o.Question != nil {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given LinksQuestionsQuestion and assigns it to the Question field.
func (o *UserFactorLinks) SetQuestion(v LinksQuestionsQuestion) {
	o.Question = &v
}

// GetResend returns the Resend field value if set, zero value otherwise.
func (o *UserFactorLinks) GetResend() LinksResendResend {
	if o == nil || o.Resend == nil {
		var ret LinksResendResend
		return ret
	}
	return *o.Resend
}

// GetResendOk returns a tuple with the Resend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetResendOk() (*LinksResendResend, bool) {
	if o == nil || o.Resend == nil {
		return nil, false
	}
	return o.Resend, true
}

// HasResend returns a boolean if a field has been set.
func (o *UserFactorLinks) HasResend() bool {
	if o != nil && o.Resend != nil {
		return true
	}

	return false
}

// SetResend gets a reference to the given LinksResendResend and assigns it to the Resend field.
func (o *UserFactorLinks) SetResend(v LinksResendResend) {
	o.Resend = &v
}

// GetSend returns the Send field value if set, zero value otherwise.
func (o *UserFactorLinks) GetSend() LinksSendSend {
	if o == nil || o.Send == nil {
		var ret LinksSendSend
		return ret
	}
	return *o.Send
}

// GetSendOk returns a tuple with the Send field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetSendOk() (*LinksSendSend, bool) {
	if o == nil || o.Send == nil {
		return nil, false
	}
	return o.Send, true
}

// HasSend returns a boolean if a field has been set.
func (o *UserFactorLinks) HasSend() bool {
	if o != nil && o.Send != nil {
		return true
	}

	return false
}

// SetSend gets a reference to the given LinksSendSend and assigns it to the Send field.
func (o *UserFactorLinks) SetSend(v LinksSendSend) {
	o.Send = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *UserFactorLinks) GetSelf() HrefObjectSelfLink {
	if o == nil || o.Self == nil {
		var ret HrefObjectSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetSelfOk() (*HrefObjectSelfLink, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *UserFactorLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given HrefObjectSelfLink and assigns it to the Self field.
func (o *UserFactorLinks) SetSelf(v HrefObjectSelfLink) {
	o.Self = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserFactorLinks) GetUser() LinksUserUser {
	if o == nil || o.User == nil {
		var ret LinksUserUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetUserOk() (*LinksUserUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserFactorLinks) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given LinksUserUser and assigns it to the User field.
func (o *UserFactorLinks) SetUser(v LinksUserUser) {
	o.User = &v
}

// GetVerify returns the Verify field value if set, zero value otherwise.
func (o *UserFactorLinks) GetVerify() LinksVerifyVerify {
	if o == nil || o.Verify == nil {
		var ret LinksVerifyVerify
		return ret
	}
	return *o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorLinks) GetVerifyOk() (*LinksVerifyVerify, bool) {
	if o == nil || o.Verify == nil {
		return nil, false
	}
	return o.Verify, true
}

// HasVerify returns a boolean if a field has been set.
func (o *UserFactorLinks) HasVerify() bool {
	if o != nil && o.Verify != nil {
		return true
	}

	return false
}

// SetVerify gets a reference to the given LinksVerifyVerify and assigns it to the Verify field.
func (o *UserFactorLinks) SetVerify(v LinksVerifyVerify) {
	o.Verify = &v
}

func (o UserFactorLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activate != nil {
		toSerialize["activate"] = o.Activate
	}
	if o.Cancel != nil {
		toSerialize["cancel"] = o.Cancel
	}
	if o.Deactivate != nil {
		toSerialize["deactivate"] = o.Deactivate
	}
	if o.Enroll != nil {
		toSerialize["enroll"] = o.Enroll
	}
	if o.Factor != nil {
		toSerialize["factor"] = o.Factor
	}
	if o.Poll != nil {
		toSerialize["poll"] = o.Poll
	}
	if o.Qrcode != nil {
		toSerialize["qrcode"] = o.Qrcode
	}
	if o.Question != nil {
		toSerialize["question"] = o.Question
	}
	if o.Resend != nil {
		toSerialize["resend"] = o.Resend
	}
	if o.Send != nil {
		toSerialize["send"] = o.Send
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Verify != nil {
		toSerialize["verify"] = o.Verify
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorLinks) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorLinks := _UserFactorLinks{}

	err = json.Unmarshal(bytes, &varUserFactorLinks)
	if err == nil {
		*o = UserFactorLinks(varUserFactorLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "activate")
		delete(additionalProperties, "cancel")
		delete(additionalProperties, "deactivate")
		delete(additionalProperties, "enroll")
		delete(additionalProperties, "factor")
		delete(additionalProperties, "poll")
		delete(additionalProperties, "qrcode")
		delete(additionalProperties, "question")
		delete(additionalProperties, "resend")
		delete(additionalProperties, "send")
		delete(additionalProperties, "self")
		delete(additionalProperties, "user")
		delete(additionalProperties, "verify")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorLinks struct {
	value *UserFactorLinks
	isSet bool
}

func (v NullableUserFactorLinks) Get() *UserFactorLinks {
	return v.value
}

func (v *NullableUserFactorLinks) Set(val *UserFactorLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorLinks(val *UserFactorLinks) *NullableUserFactorLinks {
	return &NullableUserFactorLinks{value: val, isSet: true}
}

func (v NullableUserFactorLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

