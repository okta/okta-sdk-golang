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
	"fmt"
)

// PrivilegedResource Base class for PrivilegedResourceRequest and PrivilegedResourceResponse
type PrivilegedResource struct {
	// Timestamp when the object was created
	Created *time.Time `json:"created,omitempty"`
	CredentialSyncInfo *CredentialSyncInfo `json:"credentialSyncInfo,omitempty"`
	// ID of the privileged resource
	Id *string `json:"id,omitempty"`
	// Timestamp when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The type of the resource
	ResourceType string `json:"resourceType"`
	// Current status of the privileged resource
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrivilegedResource PrivilegedResource

// NewPrivilegedResource instantiates a new PrivilegedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivilegedResource(resourceType string) *PrivilegedResource {
	this := PrivilegedResource{}
	this.ResourceType = resourceType
	return &this
}

// NewPrivilegedResourceWithDefaults instantiates a new PrivilegedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivilegedResourceWithDefaults() *PrivilegedResource {
	this := PrivilegedResource{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PrivilegedResource) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResource) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PrivilegedResource) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PrivilegedResource) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCredentialSyncInfo returns the CredentialSyncInfo field value if set, zero value otherwise.
func (o *PrivilegedResource) GetCredentialSyncInfo() CredentialSyncInfo {
	if o == nil || o.CredentialSyncInfo == nil {
		var ret CredentialSyncInfo
		return ret
	}
	return *o.CredentialSyncInfo
}

// GetCredentialSyncInfoOk returns a tuple with the CredentialSyncInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResource) GetCredentialSyncInfoOk() (*CredentialSyncInfo, bool) {
	if o == nil || o.CredentialSyncInfo == nil {
		return nil, false
	}
	return o.CredentialSyncInfo, true
}

// HasCredentialSyncInfo returns a boolean if a field has been set.
func (o *PrivilegedResource) HasCredentialSyncInfo() bool {
	if o != nil && o.CredentialSyncInfo != nil {
		return true
	}

	return false
}

// SetCredentialSyncInfo gets a reference to the given CredentialSyncInfo and assigns it to the CredentialSyncInfo field.
func (o *PrivilegedResource) SetCredentialSyncInfo(v CredentialSyncInfo) {
	o.CredentialSyncInfo = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PrivilegedResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResource) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PrivilegedResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PrivilegedResource) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *PrivilegedResource) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResource) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *PrivilegedResource) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *PrivilegedResource) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetResourceType returns the ResourceType field value
func (o *PrivilegedResource) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *PrivilegedResource) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *PrivilegedResource) SetResourceType(v string) {
	o.ResourceType = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PrivilegedResource) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivilegedResource) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PrivilegedResource) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PrivilegedResource) SetStatus(v string) {
	o.Status = &v
}

func (o PrivilegedResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.CredentialSyncInfo != nil {
		toSerialize["credentialSyncInfo"] = o.CredentialSyncInfo
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrivilegedResource) UnmarshalJSON(bytes []byte) (err error) {
	varPrivilegedResource := _PrivilegedResource{}

	err = json.Unmarshal(bytes, &varPrivilegedResource)
	if err == nil {
		*o = PrivilegedResource(varPrivilegedResource)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "created")
		delete(additionalProperties, "credentialSyncInfo")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrivilegedResource struct {
	value *PrivilegedResource
	isSet bool
}

func (v NullablePrivilegedResource) Get() *PrivilegedResource {
	return v.value
}

func (v *NullablePrivilegedResource) Set(val *PrivilegedResource) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivilegedResource) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivilegedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivilegedResource(val *PrivilegedResource) *NullablePrivilegedResource {
	return &NullablePrivilegedResource{value: val, isSet: true}
}

func (v NullablePrivilegedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivilegedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

