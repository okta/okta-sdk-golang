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
	"time"
)

// Session struct for Session
type Session struct {
	// Authentication method reference
	Amr []string `json:"amr,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A timestamp when the Session expires
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// A unique key for the Session
	Id *string `json:"id,omitempty"`
	Idp *SessionIdentityProvider `json:"idp,omitempty"`
	// A timestamp when the user last performed multifactor authentication
	LastFactorVerification *time.Time `json:"lastFactorVerification,omitempty"`
	// A timestamp when the user last performed the primary or step-up authentication with a password
	LastPasswordVerification *time.Time `json:"lastPasswordVerification,omitempty"`
	// A unique identifier for the user (username)
	Login *string `json:"login,omitempty"`
	Status *string `json:"status,omitempty"`
	// A unique key for the user
	UserId *string `json:"userId,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Session Session

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession() *Session {
	this := Session{}
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetAmr returns the Amr field value if set, zero value otherwise.
func (o *Session) GetAmr() []string {
	if o == nil || o.Amr == nil {
		var ret []string
		return ret
	}
	return o.Amr
}

// GetAmrOk returns a tuple with the Amr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAmrOk() ([]string, bool) {
	if o == nil || o.Amr == nil {
		return nil, false
	}
	return o.Amr, true
}

// HasAmr returns a boolean if a field has been set.
func (o *Session) HasAmr() bool {
	if o != nil && o.Amr != nil {
		return true
	}

	return false
}

// SetAmr gets a reference to the given []string and assigns it to the Amr field.
func (o *Session) SetAmr(v []string) {
	o.Amr = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Session) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Session) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Session) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Session) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Session) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *Session) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Session) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Session) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Session) SetId(v string) {
	o.Id = &v
}

// GetIdp returns the Idp field value if set, zero value otherwise.
func (o *Session) GetIdp() SessionIdentityProvider {
	if o == nil || o.Idp == nil {
		var ret SessionIdentityProvider
		return ret
	}
	return *o.Idp
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetIdpOk() (*SessionIdentityProvider, bool) {
	if o == nil || o.Idp == nil {
		return nil, false
	}
	return o.Idp, true
}

// HasIdp returns a boolean if a field has been set.
func (o *Session) HasIdp() bool {
	if o != nil && o.Idp != nil {
		return true
	}

	return false
}

// SetIdp gets a reference to the given SessionIdentityProvider and assigns it to the Idp field.
func (o *Session) SetIdp(v SessionIdentityProvider) {
	o.Idp = &v
}

// GetLastFactorVerification returns the LastFactorVerification field value if set, zero value otherwise.
func (o *Session) GetLastFactorVerification() time.Time {
	if o == nil || o.LastFactorVerification == nil {
		var ret time.Time
		return ret
	}
	return *o.LastFactorVerification
}

// GetLastFactorVerificationOk returns a tuple with the LastFactorVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetLastFactorVerificationOk() (*time.Time, bool) {
	if o == nil || o.LastFactorVerification == nil {
		return nil, false
	}
	return o.LastFactorVerification, true
}

// HasLastFactorVerification returns a boolean if a field has been set.
func (o *Session) HasLastFactorVerification() bool {
	if o != nil && o.LastFactorVerification != nil {
		return true
	}

	return false
}

// SetLastFactorVerification gets a reference to the given time.Time and assigns it to the LastFactorVerification field.
func (o *Session) SetLastFactorVerification(v time.Time) {
	o.LastFactorVerification = &v
}

// GetLastPasswordVerification returns the LastPasswordVerification field value if set, zero value otherwise.
func (o *Session) GetLastPasswordVerification() time.Time {
	if o == nil || o.LastPasswordVerification == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPasswordVerification
}

// GetLastPasswordVerificationOk returns a tuple with the LastPasswordVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetLastPasswordVerificationOk() (*time.Time, bool) {
	if o == nil || o.LastPasswordVerification == nil {
		return nil, false
	}
	return o.LastPasswordVerification, true
}

// HasLastPasswordVerification returns a boolean if a field has been set.
func (o *Session) HasLastPasswordVerification() bool {
	if o != nil && o.LastPasswordVerification != nil {
		return true
	}

	return false
}

// SetLastPasswordVerification gets a reference to the given time.Time and assigns it to the LastPasswordVerification field.
func (o *Session) SetLastPasswordVerification(v time.Time) {
	o.LastPasswordVerification = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *Session) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *Session) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *Session) SetLogin(v string) {
	o.Login = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Session) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Session) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Session) SetStatus(v string) {
	o.Status = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Session) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Session) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Session) SetUserId(v string) {
	o.UserId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Session) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Session) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *Session) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amr != nil {
		toSerialize["amr"] = o.Amr
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Idp != nil {
		toSerialize["idp"] = o.Idp
	}
	if o.LastFactorVerification != nil {
		toSerialize["lastFactorVerification"] = o.LastFactorVerification
	}
	if o.LastPasswordVerification != nil {
		toSerialize["lastPasswordVerification"] = o.LastPasswordVerification
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Session) UnmarshalJSON(bytes []byte) (err error) {
	varSession := _Session{}

	err = json.Unmarshal(bytes, &varSession)
	if err == nil {
		*o = Session(varSession)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "amr")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "idp")
		delete(additionalProperties, "lastFactorVerification")
		delete(additionalProperties, "lastPasswordVerification")
		delete(additionalProperties, "login")
		delete(additionalProperties, "status")
		delete(additionalProperties, "userId")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

