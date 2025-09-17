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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"time"
)

// checks if the BaseContextSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseContextSession{}

// BaseContextSession Details of the user session
type BaseContextSession struct {
	// The unique identifier for the user's session
	Id *string `json:"id,omitempty"`
	// The unique identifier for the user
	UserId *string `json:"userId,omitempty"`
	// The username used to identify the user. This is often the user's email address.
	Login *string `json:"login,omitempty"`
	// Timestamp of when the session was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Timestamp of when the session expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Represents the current status of the user's session
	Status *string `json:"status,omitempty"`
	// Timestamp of when the user was last authenticated
	LastPasswordVerification *time.Time `json:"lastPasswordVerification,omitempty"`
	// The authentication method reference
	Amr []string                 `json:"amr,omitempty"`
	Idp *SessionIdentityProvider `json:"idp,omitempty"`
	// Describes whether multifactor authentication was enabled
	MfaActive            *bool `json:"mfaActive,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseContextSession BaseContextSession

// NewBaseContextSession instantiates a new BaseContextSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContextSession() *BaseContextSession {
	this := BaseContextSession{}
	return &this
}

// NewBaseContextSessionWithDefaults instantiates a new BaseContextSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContextSessionWithDefaults() *BaseContextSession {
	this := BaseContextSession{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseContextSession) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseContextSession) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BaseContextSession) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BaseContextSession) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BaseContextSession) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BaseContextSession) SetUserId(v string) {
	o.UserId = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *BaseContextSession) GetLogin() string {
	if o == nil || IsNil(o.Login) {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetLoginOk() (*string, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *BaseContextSession) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *BaseContextSession) SetLogin(v string) {
	o.Login = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BaseContextSession) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BaseContextSession) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *BaseContextSession) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *BaseContextSession) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *BaseContextSession) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *BaseContextSession) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BaseContextSession) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BaseContextSession) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BaseContextSession) SetStatus(v string) {
	o.Status = &v
}

// GetLastPasswordVerification returns the LastPasswordVerification field value if set, zero value otherwise.
func (o *BaseContextSession) GetLastPasswordVerification() time.Time {
	if o == nil || IsNil(o.LastPasswordVerification) {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordVerification
}

// GetLastPasswordVerificationOk returns a tuple with the LastPasswordVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetLastPasswordVerificationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPasswordVerification) {
		return nil, false
	}
	return o.LastPasswordVerification, true
}

// HasLastPasswordVerification returns a boolean if a field has been set.
func (o *BaseContextSession) HasLastPasswordVerification() bool {
	if o != nil && !IsNil(o.LastPasswordVerification) {
		return true
	}

	return false
}

// SetLastPasswordVerification gets a reference to the given time.Time and assigns it to the LastPasswordVerification field.
func (o *BaseContextSession) SetLastPasswordVerification(v time.Time) {
	o.LastPasswordVerification = &v
}

// GetAmr returns the Amr field value if set, zero value otherwise.
func (o *BaseContextSession) GetAmr() []string {
	if o == nil || IsNil(o.Amr) {
		var ret []string
		return ret
	}
	return o.Amr
}

// GetAmrOk returns a tuple with the Amr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetAmrOk() ([]string, bool) {
	if o == nil || IsNil(o.Amr) {
		return nil, false
	}
	return o.Amr, true
}

// HasAmr returns a boolean if a field has been set.
func (o *BaseContextSession) HasAmr() bool {
	if o != nil && !IsNil(o.Amr) {
		return true
	}

	return false
}

// SetAmr gets a reference to the given []string and assigns it to the Amr field.
func (o *BaseContextSession) SetAmr(v []string) {
	o.Amr = v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *BaseContextSession) GetIdp() SessionIdentityProvider {
	if o == nil || IsNil(o.Idp) {
		var ret SessionIdentityProvider
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetIdpOk() (*SessionIdentityProvider, bool) {
	if o == nil || IsNil(o.Idp) {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *BaseContextSession) HasIdp() bool {
	if o != nil && !IsNil(o.Idp) {
		return true
	}

	return false
}

// SetIdp gets a reference to the given SessionIdentityProvider and assigns it to the Idp field.
func (o *BaseContextSession) SetIdp(v SessionIdentityProvider) {
	o.Idp = &v
}

// GetMfaActive returns the MfaActive field value if set, zero value otherwise.
func (o *BaseContextSession) GetMfaActive() bool {
	if o == nil || IsNil(o.MfaActive) {
		var ret bool
		return ret
	}
	return *o.MfaActive
}

// GetMfaActiveOk returns a tuple with the MfaActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextSession) GetMfaActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.MfaActive) {
		return nil, false
	}
	return o.MfaActive, true
}

// HasMfaActive returns a boolean if a field has been set.
func (o *BaseContextSession) HasMfaActive() bool {
	if o != nil && !IsNil(o.MfaActive) {
		return true
	}

	return false
}

// SetMfaActive gets a reference to the given bool and assigns it to the MfaActive field.
func (o *BaseContextSession) SetMfaActive(v bool) {
	o.MfaActive = &v
}

func (o BaseContextSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseContextSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.LastPasswordVerification) {
		toSerialize["lastPasswordVerification"] = o.LastPasswordVerification
	}
	if !IsNil(o.Amr) {
		toSerialize["amr"] = o.Amr
	}
	if !IsNil(o.Idp) {
		toSerialize["idp"] = o.Idp
	}
	if !IsNil(o.MfaActive) {
		toSerialize["mfaActive"] = o.MfaActive
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BaseContextSession) UnmarshalJSON(data []byte) (err error) {
	varBaseContextSession := _BaseContextSession{}

	err = json.Unmarshal(data, &varBaseContextSession)

	if err != nil {
		return err
	}

	*o = BaseContextSession(varBaseContextSession)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "login")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "status")
		delete(additionalProperties, "lastPasswordVerification")
		delete(additionalProperties, "amr")
		delete(additionalProperties, "idp")
		delete(additionalProperties, "mfaActive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseContextSession struct {
	value *BaseContextSession
	isSet bool
}

func (v NullableBaseContextSession) Get() *BaseContextSession {
	return v.value
}

func (v *NullableBaseContextSession) Set(val *BaseContextSession) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContextSession) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContextSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContextSession(val *BaseContextSession) *NullableBaseContextSession {
	return &NullableBaseContextSession{value: val, isSet: true}
}

func (v NullableBaseContextSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContextSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
