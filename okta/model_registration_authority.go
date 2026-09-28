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
	"time"
)

// checks if the RegistrationAuthority type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationAuthority{}

// RegistrationAuthority A registration authority (RA) configuration — the SCEP endpoint that mobile device management software enrolls devices against
type RegistrationAuthority struct {
	// The authority key identifier of the RA certificate — the `ski` of the CA certificate that signed it. After a rollover or a parallel renewal, this is the renewed CA's `ski`.
	Aki *string `json:"aki,omitempty"`
	// The stable identifier of the certificate authority this configuration belongs to
	AuthorityInstanceId *string `json:"authorityInstanceId,omitempty"`
	// Timestamp when the CA certificate bound to this configuration expires. It's resolved from the configuration's own `aki`, so configurations under one authority can report different expiries after a renewal.
	CaExpirationDate *time.Time `json:"caExpirationDate,omitempty"`
	// How the configuration validates an enrollment challenge. `STATIC` uses a fixed shared secret, as Jamf Pro does. `DYNAMIC` uses a per-enrollment secret that Okta issues. `DELEGATED` hands validation to Microsoft Intune.
	ChallengeType *string                          `json:"challengeType,omitempty"`
	ConfigInfo    *RegistrationAuthorityConfigInfo `json:"configInfo,omitempty"`
	// Timestamp when the object was created
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// The ID of the configuration. It's the `raId` segment of the SCEP enrollment URL, and the value you pass to a renewal.
	Id *string `json:"id,omitempty"`
	// Timestamp when the object was last updated
	LastUpdatedDate *time.Time `json:"lastUpdatedDate,omitempty"`
	// The name of the configuration
	Name *string `json:"name,omitempty"`
	// The enrollment protocol the configuration serves
	Protocol *string `json:"protocol,omitempty"`
	// The SCEP enrollment URL to put into your MDM software's profile. It's built from `authorityInstanceId` rather than from a certificate, so a renewal never changes it.
	ScepUrl *string `json:"scepUrl,omitempty"`
	// The subject key identifier of the RA certificate, in uppercase hexadecimal
	Ski *string `json:"ski,omitempty"`
	// The ID of the configuration this one was cloned from, on a configuration created by a parallel renewal. Absent on every other configuration.
	SourceConfigId *string `json:"sourceConfigId,omitempty"`
	// The status of the configuration. Only active configurations are returned.
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegistrationAuthority RegistrationAuthority

// NewRegistrationAuthority instantiates a new RegistrationAuthority object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationAuthority() *RegistrationAuthority {
	this := RegistrationAuthority{}
	return &this
}

// NewRegistrationAuthorityWithDefaults instantiates a new RegistrationAuthority object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationAuthorityWithDefaults() *RegistrationAuthority {
	this := RegistrationAuthority{}
	return &this
}

// GetAki returns the Aki field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetAki() string {
	if o == nil || IsNil(o.Aki) {
		var ret string
		return ret
	}
	return *o.Aki
}

// GetAkiOk returns a tuple with the Aki field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetAkiOk() (*string, bool) {
	if o == nil || IsNil(o.Aki) {
		return nil, false
	}
	return o.Aki, true
}

// HasAki returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasAki() bool {
	if o != nil && !IsNil(o.Aki) {
		return true
	}

	return false
}

// SetAki gets a reference to the given string and assigns it to the Aki field.
func (o *RegistrationAuthority) SetAki(v string) {
	o.Aki = &v
}

// GetAuthorityInstanceId returns the AuthorityInstanceId field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetAuthorityInstanceId() string {
	if o == nil || IsNil(o.AuthorityInstanceId) {
		var ret string
		return ret
	}
	return *o.AuthorityInstanceId
}

// GetAuthorityInstanceIdOk returns a tuple with the AuthorityInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetAuthorityInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorityInstanceId) {
		return nil, false
	}
	return o.AuthorityInstanceId, true
}

// HasAuthorityInstanceId returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasAuthorityInstanceId() bool {
	if o != nil && !IsNil(o.AuthorityInstanceId) {
		return true
	}

	return false
}

// SetAuthorityInstanceId gets a reference to the given string and assigns it to the AuthorityInstanceId field.
func (o *RegistrationAuthority) SetAuthorityInstanceId(v string) {
	o.AuthorityInstanceId = &v
}

// GetCaExpirationDate returns the CaExpirationDate field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetCaExpirationDate() time.Time {
	if o == nil || IsNil(o.CaExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.CaExpirationDate
}

// GetCaExpirationDateOk returns a tuple with the CaExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetCaExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CaExpirationDate) {
		return nil, false
	}
	return o.CaExpirationDate, true
}

// HasCaExpirationDate returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasCaExpirationDate() bool {
	if o != nil && !IsNil(o.CaExpirationDate) {
		return true
	}

	return false
}

// SetCaExpirationDate gets a reference to the given time.Time and assigns it to the CaExpirationDate field.
func (o *RegistrationAuthority) SetCaExpirationDate(v time.Time) {
	o.CaExpirationDate = &v
}

// GetChallengeType returns the ChallengeType field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetChallengeType() string {
	if o == nil || IsNil(o.ChallengeType) {
		var ret string
		return ret
	}
	return *o.ChallengeType
}

// GetChallengeTypeOk returns a tuple with the ChallengeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetChallengeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ChallengeType) {
		return nil, false
	}
	return o.ChallengeType, true
}

// HasChallengeType returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasChallengeType() bool {
	if o != nil && !IsNil(o.ChallengeType) {
		return true
	}

	return false
}

// SetChallengeType gets a reference to the given string and assigns it to the ChallengeType field.
func (o *RegistrationAuthority) SetChallengeType(v string) {
	o.ChallengeType = &v
}

// GetConfigInfo returns the ConfigInfo field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetConfigInfo() RegistrationAuthorityConfigInfo {
	if o == nil || IsNil(o.ConfigInfo) {
		var ret RegistrationAuthorityConfigInfo
		return ret
	}
	return *o.ConfigInfo
}

// GetConfigInfoOk returns a tuple with the ConfigInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetConfigInfoOk() (*RegistrationAuthorityConfigInfo, bool) {
	if o == nil || IsNil(o.ConfigInfo) {
		return nil, false
	}
	return o.ConfigInfo, true
}

// HasConfigInfo returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasConfigInfo() bool {
	if o != nil && !IsNil(o.ConfigInfo) {
		return true
	}

	return false
}

// SetConfigInfo gets a reference to the given RegistrationAuthorityConfigInfo and assigns it to the ConfigInfo field.
func (o *RegistrationAuthority) SetConfigInfo(v RegistrationAuthorityConfigInfo) {
	o.ConfigInfo = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *RegistrationAuthority) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegistrationAuthority) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedDate returns the LastUpdatedDate field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetLastUpdatedDate() time.Time {
	if o == nil || IsNil(o.LastUpdatedDate) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDate
}

// GetLastUpdatedDateOk returns a tuple with the LastUpdatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetLastUpdatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdatedDate) {
		return nil, false
	}
	return o.LastUpdatedDate, true
}

// HasLastUpdatedDate returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasLastUpdatedDate() bool {
	if o != nil && !IsNil(o.LastUpdatedDate) {
		return true
	}

	return false
}

// SetLastUpdatedDate gets a reference to the given time.Time and assigns it to the LastUpdatedDate field.
func (o *RegistrationAuthority) SetLastUpdatedDate(v time.Time) {
	o.LastUpdatedDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegistrationAuthority) SetName(v string) {
	o.Name = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *RegistrationAuthority) SetProtocol(v string) {
	o.Protocol = &v
}

// GetScepUrl returns the ScepUrl field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetScepUrl() string {
	if o == nil || IsNil(o.ScepUrl) {
		var ret string
		return ret
	}
	return *o.ScepUrl
}

// GetScepUrlOk returns a tuple with the ScepUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetScepUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ScepUrl) {
		return nil, false
	}
	return o.ScepUrl, true
}

// HasScepUrl returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasScepUrl() bool {
	if o != nil && !IsNil(o.ScepUrl) {
		return true
	}

	return false
}

// SetScepUrl gets a reference to the given string and assigns it to the ScepUrl field.
func (o *RegistrationAuthority) SetScepUrl(v string) {
	o.ScepUrl = &v
}

// GetSki returns the Ski field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetSki() string {
	if o == nil || IsNil(o.Ski) {
		var ret string
		return ret
	}
	return *o.Ski
}

// GetSkiOk returns a tuple with the Ski field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetSkiOk() (*string, bool) {
	if o == nil || IsNil(o.Ski) {
		return nil, false
	}
	return o.Ski, true
}

// HasSki returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasSki() bool {
	if o != nil && !IsNil(o.Ski) {
		return true
	}

	return false
}

// SetSki gets a reference to the given string and assigns it to the Ski field.
func (o *RegistrationAuthority) SetSki(v string) {
	o.Ski = &v
}

// GetSourceConfigId returns the SourceConfigId field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetSourceConfigId() string {
	if o == nil || IsNil(o.SourceConfigId) {
		var ret string
		return ret
	}
	return *o.SourceConfigId
}

// GetSourceConfigIdOk returns a tuple with the SourceConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetSourceConfigIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceConfigId) {
		return nil, false
	}
	return o.SourceConfigId, true
}

// HasSourceConfigId returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasSourceConfigId() bool {
	if o != nil && !IsNil(o.SourceConfigId) {
		return true
	}

	return false
}

// SetSourceConfigId gets a reference to the given string and assigns it to the SourceConfigId field.
func (o *RegistrationAuthority) SetSourceConfigId(v string) {
	o.SourceConfigId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RegistrationAuthority) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationAuthority) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RegistrationAuthority) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RegistrationAuthority) SetStatus(v string) {
	o.Status = &v
}

func (o RegistrationAuthority) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationAuthority) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aki) {
		toSerialize["aki"] = o.Aki
	}
	if !IsNil(o.AuthorityInstanceId) {
		toSerialize["authorityInstanceId"] = o.AuthorityInstanceId
	}
	if !IsNil(o.CaExpirationDate) {
		toSerialize["caExpirationDate"] = o.CaExpirationDate
	}
	if !IsNil(o.ChallengeType) {
		toSerialize["challengeType"] = o.ChallengeType
	}
	if !IsNil(o.ConfigInfo) {
		toSerialize["configInfo"] = o.ConfigInfo
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdatedDate) {
		toSerialize["lastUpdatedDate"] = o.LastUpdatedDate
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.ScepUrl) {
		toSerialize["scepUrl"] = o.ScepUrl
	}
	if !IsNil(o.Ski) {
		toSerialize["ski"] = o.Ski
	}
	if !IsNil(o.SourceConfigId) {
		toSerialize["sourceConfigId"] = o.SourceConfigId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegistrationAuthority) UnmarshalJSON(data []byte) (err error) {
	varRegistrationAuthority := _RegistrationAuthority{}

	err = json.Unmarshal(data, &varRegistrationAuthority)

	if err != nil {
		return err
	}

	*o = RegistrationAuthority(varRegistrationAuthority)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "aki")
		delete(additionalProperties, "authorityInstanceId")
		delete(additionalProperties, "caExpirationDate")
		delete(additionalProperties, "challengeType")
		delete(additionalProperties, "configInfo")
		delete(additionalProperties, "createdDate")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdatedDate")
		delete(additionalProperties, "name")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "scepUrl")
		delete(additionalProperties, "ski")
		delete(additionalProperties, "sourceConfigId")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegistrationAuthority struct {
	value *RegistrationAuthority
	isSet bool
}

func (v NullableRegistrationAuthority) Get() *RegistrationAuthority {
	return v.value
}

func (v *NullableRegistrationAuthority) Set(val *RegistrationAuthority) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationAuthority) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationAuthority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationAuthority(val *RegistrationAuthority) *NullableRegistrationAuthority {
	return &NullableRegistrationAuthority{value: val, isSet: true}
}

func (v NullableRegistrationAuthority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationAuthority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
