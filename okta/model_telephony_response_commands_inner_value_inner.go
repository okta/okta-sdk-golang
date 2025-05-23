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

// TelephonyResponseCommandsInnerValueInner struct for TelephonyResponseCommandsInnerValueInner
type TelephonyResponseCommandsInnerValueInner struct {
	// Status of telephony callout
	Status *string `json:"status,omitempty"`
	// Telephony provider for sms/voice
	Provider *string `json:"provider,omitempty"`
	// Transaction ID for sms/voice
	TransactionId *string `json:"transactionId,omitempty"`
	// Any relevant metadata for the telephony transaction
	TransactionMetadata *string `json:"transactionMetadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelephonyResponseCommandsInnerValueInner TelephonyResponseCommandsInnerValueInner

// NewTelephonyResponseCommandsInnerValueInner instantiates a new TelephonyResponseCommandsInnerValueInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelephonyResponseCommandsInnerValueInner() *TelephonyResponseCommandsInnerValueInner {
	this := TelephonyResponseCommandsInnerValueInner{}
	return &this
}

// NewTelephonyResponseCommandsInnerValueInnerWithDefaults instantiates a new TelephonyResponseCommandsInnerValueInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelephonyResponseCommandsInnerValueInnerWithDefaults() *TelephonyResponseCommandsInnerValueInner {
	this := TelephonyResponseCommandsInnerValueInner{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TelephonyResponseCommandsInnerValueInner) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponseCommandsInnerValueInner) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TelephonyResponseCommandsInnerValueInner) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TelephonyResponseCommandsInnerValueInner) SetStatus(v string) {
	o.Status = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *TelephonyResponseCommandsInnerValueInner) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponseCommandsInnerValueInner) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *TelephonyResponseCommandsInnerValueInner) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *TelephonyResponseCommandsInnerValueInner) SetProvider(v string) {
	o.Provider = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *TelephonyResponseCommandsInnerValueInner) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponseCommandsInnerValueInner) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *TelephonyResponseCommandsInnerValueInner) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *TelephonyResponseCommandsInnerValueInner) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionMetadata returns the TransactionMetadata field value if set, zero value otherwise.
func (o *TelephonyResponseCommandsInnerValueInner) GetTransactionMetadata() string {
	if o == nil || o.TransactionMetadata == nil {
		var ret string
		return ret
	}
	return *o.TransactionMetadata
}

// GetTransactionMetadataOk returns a tuple with the TransactionMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelephonyResponseCommandsInnerValueInner) GetTransactionMetadataOk() (*string, bool) {
	if o == nil || o.TransactionMetadata == nil {
		return nil, false
	}
	return o.TransactionMetadata, true
}

// HasTransactionMetadata returns a boolean if a field has been set.
func (o *TelephonyResponseCommandsInnerValueInner) HasTransactionMetadata() bool {
	if o != nil && o.TransactionMetadata != nil {
		return true
	}

	return false
}

// SetTransactionMetadata gets a reference to the given string and assigns it to the TransactionMetadata field.
func (o *TelephonyResponseCommandsInnerValueInner) SetTransactionMetadata(v string) {
	o.TransactionMetadata = &v
}

func (o TelephonyResponseCommandsInnerValueInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Provider != nil {
		toSerialize["provider"] = o.Provider
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	if o.TransactionMetadata != nil {
		toSerialize["transactionMetadata"] = o.TransactionMetadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelephonyResponseCommandsInnerValueInner) UnmarshalJSON(bytes []byte) (err error) {
	varTelephonyResponseCommandsInnerValueInner := _TelephonyResponseCommandsInnerValueInner{}

	err = json.Unmarshal(bytes, &varTelephonyResponseCommandsInnerValueInner)
	if err == nil {
		*o = TelephonyResponseCommandsInnerValueInner(varTelephonyResponseCommandsInnerValueInner)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "provider")
		delete(additionalProperties, "transactionId")
		delete(additionalProperties, "transactionMetadata")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTelephonyResponseCommandsInnerValueInner struct {
	value *TelephonyResponseCommandsInnerValueInner
	isSet bool
}

func (v NullableTelephonyResponseCommandsInnerValueInner) Get() *TelephonyResponseCommandsInnerValueInner {
	return v.value
}

func (v *NullableTelephonyResponseCommandsInnerValueInner) Set(val *TelephonyResponseCommandsInnerValueInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTelephonyResponseCommandsInnerValueInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTelephonyResponseCommandsInnerValueInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelephonyResponseCommandsInnerValueInner(val *TelephonyResponseCommandsInnerValueInner) *NullableTelephonyResponseCommandsInnerValueInner {
	return &NullableTelephonyResponseCommandsInnerValueInner{value: val, isSet: true}
}

func (v NullableTelephonyResponseCommandsInnerValueInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelephonyResponseCommandsInnerValueInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

