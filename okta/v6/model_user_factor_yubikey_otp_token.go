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
)

// UserFactorYubikeyOtpToken struct for UserFactorYubikeyOtpToken
type UserFactorYubikeyOtpToken struct {
	// Timestamp when the token was created
	Created *time.Time `json:"created,omitempty"`
	// ID of the token
	Id *string `json:"id,omitempty"`
	// Timestamp when the token was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Timestamp when the token was last verified
	LastVerified *time.Time `json:"lastVerified,omitempty"`
	// Specified profile information for token
	Profile map[string]map[string]interface{} `json:"profile,omitempty"`
	// Token status
	Status *string `json:"status,omitempty"`
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Links *UserFactorLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserFactorYubikeyOtpToken UserFactorYubikeyOtpToken

// NewUserFactorYubikeyOtpToken instantiates a new UserFactorYubikeyOtpToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFactorYubikeyOtpToken() *UserFactorYubikeyOtpToken {
	this := UserFactorYubikeyOtpToken{}
	return &this
}

// NewUserFactorYubikeyOtpTokenWithDefaults instantiates a new UserFactorYubikeyOtpToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFactorYubikeyOtpTokenWithDefaults() *UserFactorYubikeyOtpToken {
	this := UserFactorYubikeyOtpToken{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *UserFactorYubikeyOtpToken) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserFactorYubikeyOtpToken) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *UserFactorYubikeyOtpToken) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastVerified returns the LastVerified field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetLastVerified() time.Time {
	if o == nil || o.LastVerified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastVerified
}

// GetLastVerifiedOk returns a tuple with the LastVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetLastVerifiedOk() (*time.Time, bool) {
	if o == nil || o.LastVerified == nil {
		return nil, false
	}
	return o.LastVerified, true
}

// HasLastVerified returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasLastVerified() bool {
	if o != nil && o.LastVerified != nil {
		return true
	}

	return false
}

// SetLastVerified gets a reference to the given time.Time and assigns it to the LastVerified field.
func (o *UserFactorYubikeyOtpToken) SetLastVerified(v time.Time) {
	o.LastVerified = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetProfile() map[string]map[string]interface{} {
	if o == nil || o.Profile == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetProfileOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]map[string]interface{} and assigns it to the Profile field.
func (o *UserFactorYubikeyOtpToken) SetProfile(v map[string]map[string]interface{}) {
	o.Profile = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserFactorYubikeyOtpToken) SetStatus(v string) {
	o.Status = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || o.Embedded == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Embedded == nil {
		return nil, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasEmbedded() bool {
	if o != nil && o.Embedded != nil {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *UserFactorYubikeyOtpToken) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserFactorYubikeyOtpToken) GetLinks() UserFactorLinks {
	if o == nil || o.Links == nil {
		var ret UserFactorLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFactorYubikeyOtpToken) GetLinksOk() (*UserFactorLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserFactorYubikeyOtpToken) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given UserFactorLinks and assigns it to the Links field.
func (o *UserFactorYubikeyOtpToken) SetLinks(v UserFactorLinks) {
	o.Links = &v
}

func (o UserFactorYubikeyOtpToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.LastVerified != nil {
		toSerialize["lastVerified"] = o.LastVerified
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Embedded != nil {
		toSerialize["_embedded"] = o.Embedded
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserFactorYubikeyOtpToken) UnmarshalJSON(bytes []byte) (err error) {
	varUserFactorYubikeyOtpToken := _UserFactorYubikeyOtpToken{}

	err = json.Unmarshal(bytes, &varUserFactorYubikeyOtpToken)
	if err == nil {
		*o = UserFactorYubikeyOtpToken(varUserFactorYubikeyOtpToken)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastVerified")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_embedded")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserFactorYubikeyOtpToken struct {
	value *UserFactorYubikeyOtpToken
	isSet bool
}

func (v NullableUserFactorYubikeyOtpToken) Get() *UserFactorYubikeyOtpToken {
	return v.value
}

func (v *NullableUserFactorYubikeyOtpToken) Set(val *UserFactorYubikeyOtpToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFactorYubikeyOtpToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFactorYubikeyOtpToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFactorYubikeyOtpToken(val *UserFactorYubikeyOtpToken) *NullableUserFactorYubikeyOtpToken {
	return &NullableUserFactorYubikeyOtpToken{value: val, isSet: true}
}

func (v NullableUserFactorYubikeyOtpToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFactorYubikeyOtpToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

