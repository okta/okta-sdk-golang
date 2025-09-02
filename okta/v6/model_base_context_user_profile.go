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
)

// BaseContextUserProfile struct for BaseContextUserProfile
type BaseContextUserProfile struct {
	// The username used to identify the user. This is often the user's email address.
	Login *string `json:"login,omitempty"`
	// The first name of the user
	FirstName *string `json:"firstName,omitempty"`
	// The last name of the user
	LastName *string `json:"lastName,omitempty"`
	// The user's default location for purposes of localizing items such as currency, date time format, numerical representations, and so on. A locale value is a concatenation of the [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) two-letter language code, an underscore, and the [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter country code. For example, `en_US` specifies the language English and country US. This value is `en_US` by default.
	Locale *string `json:"locale,omitempty"`
	// The user's timezone
	TimeZone *string `json:"timeZone,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BaseContextUserProfile BaseContextUserProfile

// NewBaseContextUserProfile instantiates a new BaseContextUserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseContextUserProfile() *BaseContextUserProfile {
	this := BaseContextUserProfile{}
	return &this
}

// NewBaseContextUserProfileWithDefaults instantiates a new BaseContextUserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseContextUserProfileWithDefaults() *BaseContextUserProfile {
	this := BaseContextUserProfile{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *BaseContextUserProfile) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserProfile) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *BaseContextUserProfile) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *BaseContextUserProfile) SetLogin(v string) {
	o.Login = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BaseContextUserProfile) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserProfile) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BaseContextUserProfile) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BaseContextUserProfile) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BaseContextUserProfile) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserProfile) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BaseContextUserProfile) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BaseContextUserProfile) SetLastName(v string) {
	o.LastName = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *BaseContextUserProfile) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserProfile) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *BaseContextUserProfile) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *BaseContextUserProfile) SetLocale(v string) {
	o.Locale = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *BaseContextUserProfile) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseContextUserProfile) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *BaseContextUserProfile) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *BaseContextUserProfile) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o BaseContextUserProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseContextUserProfile) UnmarshalJSON(bytes []byte) (err error) {
	varBaseContextUserProfile := _BaseContextUserProfile{}

	err = json.Unmarshal(bytes, &varBaseContextUserProfile)
	if err == nil {
		*o = BaseContextUserProfile(varBaseContextUserProfile)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "timeZone")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableBaseContextUserProfile struct {
	value *BaseContextUserProfile
	isSet bool
}

func (v NullableBaseContextUserProfile) Get() *BaseContextUserProfile {
	return v.value
}

func (v *NullableBaseContextUserProfile) Set(val *BaseContextUserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseContextUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseContextUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseContextUserProfile(val *BaseContextUserProfile) *NullableBaseContextUserProfile {
	return &NullableBaseContextUserProfile{value: val, isSet: true}
}

func (v NullableBaseContextUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseContextUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

