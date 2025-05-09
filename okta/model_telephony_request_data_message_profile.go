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

// TelephonyRequestDataMessageProfile Message profile specifies information about the telephony (sms/voice) message to be sent to the Okta user
type TelephonyRequestDataMessageProfile struct {
	// Default or Okta org configured sms or voice message template
	MsgTemplate *string `json:"msgTemplate,omitempty"`
	// The Okta's user's phone number
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The time when OTP expires
	OtpExpires *string `json:"otpExpires,omitempty"`
	// The channel for OTP delivery - SMS or voice
	DeliveryChannel *string `json:"deliveryChannel,omitempty"`
	// The OTP code requested by the Okta user
	OtpCode *string `json:"otpCode,omitempty"`
	// The locale associated with the Okta user
	Locale *string `json:"locale,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyRequestDataMessageProfile TelephonyRequestDataMessageProfile

// NewTelephonyRequestDataMessageProfile instantiates a new TelephonyRequestDataMessageProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyRequestDataMessageProfile() *TelephonyRequestDataMessageProfile {
	this := TelephonyRequestDataMessageProfile{}
	return &this
}

// NewTelephonyRequestDataMessageProfileWithDefaults instantiates a new TelephonyRequestDataMessageProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyRequestDataMessageProfileWithDefaults() *TelephonyRequestDataMessageProfile {
	this := TelephonyRequestDataMessageProfile{}
	return &this
}

// GetMsgTemplate returns the MsgTemplate field value if set, zero value otherwise.
func (o *TelephonyRequestDataMessageProfile) GetMsgTemplate() string {
	if o == nil || o.MsgTemplate == nil {
		var ret string
		return ret
	}
	return *o.MsgTemplate
}

// GetMsgTemplateOk returns a tuple with the MsgTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataMessageProfile) GetMsgTemplateOk() (*string, bool) {
	if o == nil || o.MsgTemplate == nil {
		return nil, false
	}
	return o.MsgTemplate, true
}

// HasMsgTemplate returns a boolean if a field has been set.
func (o *TelephonyRequestDataMessageProfile) HasMsgTemplate() bool {
	if o != nil && o.MsgTemplate != nil {
		return true
	}

	return false
}

// SetMsgTemplate gets a reference to the given string and assigns it to the MsgTemplate field.
func (o *TelephonyRequestDataMessageProfile) SetMsgTemplate(v string) {
	o.MsgTemplate = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *TelephonyRequestDataMessageProfile) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataMessageProfile) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *TelephonyRequestDataMessageProfile) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *TelephonyRequestDataMessageProfile) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetOtpExpires returns the OtpExpires field value if set, zero value otherwise.
func (o *TelephonyRequestDataMessageProfile) GetOtpExpires() string {
	if o == nil || o.OtpExpires == nil {
		var ret string
		return ret
	}
	return *o.OtpExpires
}

// GetOtpExpiresOk returns a tuple with the OtpExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataMessageProfile) GetOtpExpiresOk() (*string, bool) {
	if o == nil || o.OtpExpires == nil {
		return nil, false
	}
	return o.OtpExpires, true
}

// HasOtpExpires returns a boolean if a field has been set.
func (o *TelephonyRequestDataMessageProfile) HasOtpExpires() bool {
	if o != nil && o.OtpExpires != nil {
		return true
	}

	return false
}

// SetOtpExpires gets a reference to the given string and assigns it to the OtpExpires field.
func (o *TelephonyRequestDataMessageProfile) SetOtpExpires(v string) {
	o.OtpExpires = &v
}

// GetDeliveryChannel returns the DeliveryChannel field value if set, zero value otherwise.
func (o *TelephonyRequestDataMessageProfile) GetDeliveryChannel() string {
	if o == nil || o.DeliveryChannel == nil {
		var ret string
		return ret
	}
	return *o.DeliveryChannel
}

// GetDeliveryChannelOk returns a tuple with the DeliveryChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataMessageProfile) GetDeliveryChannelOk() (*string, bool) {
	if o == nil || o.DeliveryChannel == nil {
		return nil, false
	}
	return o.DeliveryChannel, true
}

// HasDeliveryChannel returns a boolean if a field has been set.
func (o *TelephonyRequestDataMessageProfile) HasDeliveryChannel() bool {
	if o != nil && o.DeliveryChannel != nil {
		return true
	}

	return false
}

// SetDeliveryChannel gets a reference to the given string and assigns it to the DeliveryChannel field.
func (o *TelephonyRequestDataMessageProfile) SetDeliveryChannel(v string) {
	o.DeliveryChannel = &v
}

// GetOtpCode returns the OtpCode field value if set, zero value otherwise.
func (o *TelephonyRequestDataMessageProfile) GetOtpCode() string {
	if o == nil || o.OtpCode == nil {
		var ret string
		return ret
	}
	return *o.OtpCode
}

// GetOtpCodeOk returns a tuple with the OtpCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataMessageProfile) GetOtpCodeOk() (*string, bool) {
	if o == nil || o.OtpCode == nil {
		return nil, false
	}
	return o.OtpCode, true
}

// HasOtpCode returns a boolean if a field has been set.
func (o *TelephonyRequestDataMessageProfile) HasOtpCode() bool {
	if o != nil && o.OtpCode != nil {
		return true
	}

	return false
}

// SetOtpCode gets a reference to the given string and assigns it to the OtpCode field.
func (o *TelephonyRequestDataMessageProfile) SetOtpCode(v string) {
	o.OtpCode = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *TelephonyRequestDataMessageProfile) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyRequestDataMessageProfile) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *TelephonyRequestDataMessageProfile) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *TelephonyRequestDataMessageProfile) SetLocale(v string) {
	o.Locale = &v
}

func (o TelephonyRequestDataMessageProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MsgTemplate != nil {
		toSerialize["msgTemplate"] = o.MsgTemplate
	}
	if o.PhoneNumber != nil {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if o.OtpExpires != nil {
		toSerialize["otpExpires"] = o.OtpExpires
	}
	if o.DeliveryChannel != nil {
		toSerialize["deliveryChannel"] = o.DeliveryChannel
	}
	if o.OtpCode != nil {
		toSerialize["otpCode"] = o.OtpCode
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyRequestDataMessageProfile) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyRequestDataMessageProfile := _TelephonyRequestDataMessageProfile{}

	err = json.Unmarshal(bytes, &varTelephonyRequestDataMessageProfile)
	if err == nil {
		*o = TelephonyRequestDataMessageProfile(varTelephonyRequestDataMessageProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "msgTemplate")
		delete(additionalProperties, "phoneNumber")
		delete(additionalProperties, "otpExpires")
		delete(additionalProperties, "deliveryChannel")
		delete(additionalProperties, "otpCode")
		delete(additionalProperties, "locale")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTelephonyRequestDataMessageProfile struct {
	value *TelephonyRequestDataMessageProfile
	isSet bool
}

func (v NullableTelephonyRequestDataMessageProfile) Get() *TelephonyRequestDataMessageProfile {
	return v.value
}

func (v *NullableTelephonyRequestDataMessageProfile) Set(val *TelephonyRequestDataMessageProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyRequestDataMessageProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyRequestDataMessageProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyRequestDataMessageProfile(val *TelephonyRequestDataMessageProfile) *NullableTelephonyRequestDataMessageProfile {
	return &NullableTelephonyRequestDataMessageProfile{value: val, isSet: true}
}

func (v NullableTelephonyRequestDataMessageProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyRequestDataMessageProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

