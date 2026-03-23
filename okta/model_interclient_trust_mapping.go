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
)

// checks if the InterclientTrustMapping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InterclientTrustMapping{}

// InterclientTrustMapping struct for InterclientTrustMapping
type InterclientTrustMapping struct {
	// The app ID of the target app
	AppInstanceId *string `json:"appInstanceId,omitempty"`
	// Timestamp when the interclient trust mapping was created
	Created *string `json:"created,omitempty"`
	// The unique ID of the interclient trust mapping
	Id *string `json:"id,omitempty"`
	// Timestamp when the interclient trust mapping was updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// ID of the user who created the interclient trust mapping
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// ID of the org
	OrgId *string `json:"orgId,omitempty"`
	// The app ID of the allowed app
	TrustedAppInstanceId *string `json:"trustedAppInstanceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InterclientTrustMapping InterclientTrustMapping

// NewInterclientTrustMapping instantiates a new InterclientTrustMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterclientTrustMapping() *InterclientTrustMapping {
	this := InterclientTrustMapping{}
	return &this
}

// NewInterclientTrustMappingWithDefaults instantiates a new InterclientTrustMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterclientTrustMappingWithDefaults() *InterclientTrustMapping {
	this := InterclientTrustMapping{}
	return &this
}

// GetAppInstanceId returns the AppInstanceId field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetAppInstanceId() string {
	if o == nil || IsNil(o.AppInstanceId) {
		var ret string
		return ret
	}
	return *o.AppInstanceId
}

// GetAppInstanceIdOk returns a tuple with the AppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppInstanceId) {
		return nil, false
	}
	return o.AppInstanceId, true
}

// HasAppInstanceId returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasAppInstanceId() bool {
	if o != nil && !IsNil(o.AppInstanceId) {
		return true
	}

	return false
}

// SetAppInstanceId gets a reference to the given string and assigns it to the AppInstanceId field.
func (o *InterclientTrustMapping) SetAppInstanceId(v string) {
	o.AppInstanceId = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *InterclientTrustMapping) SetCreated(v string) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InterclientTrustMapping) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *InterclientTrustMapping) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *InterclientTrustMapping) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *InterclientTrustMapping) SetOrgId(v string) {
	o.OrgId = &v
}

// GetTrustedAppInstanceId returns the TrustedAppInstanceId field value if set, zero value otherwise.
func (o *InterclientTrustMapping) GetTrustedAppInstanceId() string {
	if o == nil || IsNil(o.TrustedAppInstanceId) {
		var ret string
		return ret
	}
	return *o.TrustedAppInstanceId
}

// GetTrustedAppInstanceIdOk returns a tuple with the TrustedAppInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterclientTrustMapping) GetTrustedAppInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrustedAppInstanceId) {
		return nil, false
	}
	return o.TrustedAppInstanceId, true
}

// HasTrustedAppInstanceId returns a boolean if a field has been set.
func (o *InterclientTrustMapping) HasTrustedAppInstanceId() bool {
	if o != nil && !IsNil(o.TrustedAppInstanceId) {
		return true
	}

	return false
}

// SetTrustedAppInstanceId gets a reference to the given string and assigns it to the TrustedAppInstanceId field.
func (o *InterclientTrustMapping) SetTrustedAppInstanceId(v string) {
	o.TrustedAppInstanceId = &v
}

func (o InterclientTrustMapping) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterclientTrustMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppInstanceId) {
		toSerialize["appInstanceId"] = o.AppInstanceId
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if !IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !IsNil(o.TrustedAppInstanceId) {
		toSerialize["trustedAppInstanceId"] = o.TrustedAppInstanceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InterclientTrustMapping) UnmarshalJSON(data []byte) (err error) {
	varInterclientTrustMapping := _InterclientTrustMapping{}

	err = json.Unmarshal(data, &varInterclientTrustMapping)

	if err != nil {
		return err
	}

	*o = InterclientTrustMapping(varInterclientTrustMapping)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "appInstanceId")
		delete(additionalProperties, "created")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "orgId")
		delete(additionalProperties, "trustedAppInstanceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterclientTrustMapping struct {
	value *InterclientTrustMapping
	isSet bool
}

func (v NullableInterclientTrustMapping) Get() *InterclientTrustMapping {
	return v.value
}

func (v *NullableInterclientTrustMapping) Set(val *InterclientTrustMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableInterclientTrustMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableInterclientTrustMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterclientTrustMapping(val *InterclientTrustMapping) *NullableInterclientTrustMapping {
	return &NullableInterclientTrustMapping{value: val, isSet: true}
}

func (v NullableInterclientTrustMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterclientTrustMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
