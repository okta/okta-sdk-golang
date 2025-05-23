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

// TrustedOrigin struct for TrustedOrigin
type TrustedOrigin struct {
	// Timestamp when the Trusted Origin was created
	Created *time.Time `json:"created,omitempty"`
	// The ID of the user who created the Trusted Origin
	CreatedBy *string `json:"createdBy,omitempty"`
	// Unique identifier for the Trusted Origin
	Id *string `json:"id,omitempty"`
	// Timestamp when the Trusted Origin was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The ID of the user who last updated the Trusted Origin
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// Unique name for the Trusted Origin
	Name *string `json:"name,omitempty"`
	// Unique origin URL for the Trusted Origin. The supported schemes for this attribute are HTTP, HTTPS, FTP, Ionic 2, and Capacitor.
	Origin *string `json:"origin,omitempty"`
	// Array of Scope types that this Trusted Origin is used for
	Scopes []TrustedOriginScope `json:"scopes,omitempty"`
	Status *string `json:"status,omitempty"`
	Links *LinksSelfAndLifecycle `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TrustedOrigin TrustedOrigin

// NewTrustedOrigin instantiates a new TrustedOrigin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrustedOrigin() *TrustedOrigin {
	this := TrustedOrigin{}
	return &this
}

// NewTrustedOriginWithDefaults instantiates a new TrustedOrigin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrustedOriginWithDefaults() *TrustedOrigin {
	this := TrustedOrigin{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *TrustedOrigin) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *TrustedOrigin) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *TrustedOrigin) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *TrustedOrigin) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *TrustedOrigin) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *TrustedOrigin) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TrustedOrigin) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TrustedOrigin) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TrustedOrigin) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *TrustedOrigin) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *TrustedOrigin) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *TrustedOrigin) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *TrustedOrigin) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *TrustedOrigin) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *TrustedOrigin) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TrustedOrigin) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TrustedOrigin) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TrustedOrigin) SetName(v string) {
	o.Name = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *TrustedOrigin) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *TrustedOrigin) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *TrustedOrigin) SetOrigin(v string) {
	o.Origin = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *TrustedOrigin) GetScopes() []TrustedOriginScope {
	if o == nil || o.Scopes == nil {
		var ret []TrustedOriginScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetScopesOk() ([]TrustedOriginScope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *TrustedOrigin) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []TrustedOriginScope and assigns it to the Scopes field.
func (o *TrustedOrigin) SetScopes(v []TrustedOriginScope) {
	o.Scopes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TrustedOrigin) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TrustedOrigin) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TrustedOrigin) SetStatus(v string) {
	o.Status = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TrustedOrigin) GetLinks() LinksSelfAndLifecycle {
	if o == nil || o.Links == nil {
		var ret LinksSelfAndLifecycle
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrustedOrigin) GetLinksOk() (*LinksSelfAndLifecycle, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TrustedOrigin) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelfAndLifecycle and assigns it to the Links field.
func (o *TrustedOrigin) SetLinks(v LinksSelfAndLifecycle) {
	o.Links = &v
}

func (o TrustedOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TrustedOrigin) UnmarshalJSON(bytes []byte) (err error) {
	varTrustedOrigin := _TrustedOrigin{}

	err = json.Unmarshal(bytes, &varTrustedOrigin)
	if err == nil {
		*o = TrustedOrigin(varTrustedOrigin)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "name")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "scopes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableTrustedOrigin struct {
	value *TrustedOrigin
	isSet bool
}

func (v NullableTrustedOrigin) Get() *TrustedOrigin {
	return v.value
}

func (v *NullableTrustedOrigin) Set(val *TrustedOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustedOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustedOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustedOrigin(val *TrustedOrigin) *NullableTrustedOrigin {
	return &NullableTrustedOrigin{value: val, isSet: true}
}

func (v NullableTrustedOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustedOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

