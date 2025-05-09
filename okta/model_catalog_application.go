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

// CatalogApplication struct for CatalogApplication
type CatalogApplication struct {
	Category *string `json:"category,omitempty"`
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Features []string `json:"features,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	SignOnModes []string `json:"signOnModes,omitempty"`
	Status *string `json:"status,omitempty"`
	VerificationStatus *string `json:"verificationStatus,omitempty"`
	Website *string `json:"website,omitempty"`
	Links *LinksSelf `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogApplication CatalogApplication

// NewCatalogApplication instantiates a new CatalogApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogApplication() *CatalogApplication {
	this := CatalogApplication{}
	return &this
}

// NewCatalogApplicationWithDefaults instantiates a new CatalogApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogApplicationWithDefaults() *CatalogApplication {
	this := CatalogApplication{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *CatalogApplication) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *CatalogApplication) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *CatalogApplication) SetCategory(v string) {
	o.Category = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CatalogApplication) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CatalogApplication) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CatalogApplication) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CatalogApplication) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CatalogApplication) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CatalogApplication) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *CatalogApplication) GetFeatures() []string {
	if o == nil || o.Features == nil {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetFeaturesOk() ([]string, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CatalogApplication) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *CatalogApplication) SetFeatures(v []string) {
	o.Features = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CatalogApplication) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CatalogApplication) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CatalogApplication) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *CatalogApplication) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *CatalogApplication) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *CatalogApplication) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CatalogApplication) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CatalogApplication) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CatalogApplication) SetName(v string) {
	o.Name = &v
}

// GetSignOnModes returns the SignOnModes field value if set, zero value otherwise.
func (o *CatalogApplication) GetSignOnModes() []string {
	if o == nil || o.SignOnModes == nil {
		var ret []string
		return ret
	}
	return o.SignOnModes
}

// GetSignOnModesOk returns a tuple with the SignOnModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetSignOnModesOk() ([]string, bool) {
	if o == nil || o.SignOnModes == nil {
		return nil, false
	}
	return o.SignOnModes, true
}

// HasSignOnModes returns a boolean if a field has been set.
func (o *CatalogApplication) HasSignOnModes() bool {
	if o != nil && o.SignOnModes != nil {
		return true
	}

	return false
}

// SetSignOnModes gets a reference to the given []string and assigns it to the SignOnModes field.
func (o *CatalogApplication) SetSignOnModes(v []string) {
	o.SignOnModes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CatalogApplication) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CatalogApplication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CatalogApplication) SetStatus(v string) {
	o.Status = &v
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise.
func (o *CatalogApplication) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetVerificationStatusOk() (*string, bool) {
	if o == nil || o.VerificationStatus == nil {
		return nil, false
	}
	return o.VerificationStatus, true
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *CatalogApplication) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus != nil {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given string and assigns it to the VerificationStatus field.
func (o *CatalogApplication) SetVerificationStatus(v string) {
	o.VerificationStatus = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *CatalogApplication) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *CatalogApplication) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *CatalogApplication) SetWebsite(v string) {
	o.Website = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CatalogApplication) GetLinks() LinksSelf {
	if o == nil || o.Links == nil {
		var ret LinksSelf
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogApplication) GetLinksOk() (*LinksSelf, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CatalogApplication) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksSelf and assigns it to the Links field.
func (o *CatalogApplication) SetLinks(v LinksSelf) {
	o.Links = &v
}

func (o CatalogApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SignOnModes != nil {
		toSerialize["signOnModes"] = o.SignOnModes
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.VerificationStatus != nil {
		toSerialize["verificationStatus"] = o.VerificationStatus
	}
	if o.Website != nil {
		toSerialize["website"] = o.Website
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CatalogApplication) UnmarshalJSON(bytes []byte) (err error) {
	varCatalogApplication := _CatalogApplication{}

	err = json.Unmarshal(bytes, &varCatalogApplication)
	if err == nil {
		*o = CatalogApplication(varCatalogApplication)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "category")
		delete(additionalProperties, "description")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "features")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "signOnModes")
		delete(additionalProperties, "status")
		delete(additionalProperties, "verificationStatus")
		delete(additionalProperties, "website")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCatalogApplication struct {
	value *CatalogApplication
	isSet bool
}

func (v NullableCatalogApplication) Get() *CatalogApplication {
	return v.value
}

func (v *NullableCatalogApplication) Set(val *CatalogApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogApplication(val *CatalogApplication) *NullableCatalogApplication {
	return &NullableCatalogApplication{value: val, isSet: true}
}

func (v NullableCatalogApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

