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

// SubmissionResponse struct for SubmissionResponse
type SubmissionResponse struct {
	// List of org-level variables for the customer per-tenant configuration. For example, a `subdomain` variable can be used in the ACS URL: `https://${org.subdomain}.example.com/saml/login`
	Config []SubmissionResponseConfigInner `json:"config,omitempty"`
	// A general description of your application and the benefits provided to your customers
	Description *string `json:"description,omitempty"`
	// OIN Integration ID
	Id *string `json:"id,omitempty"`
	// Timestamp when the OIN Integration was last published
	LastPublished *string `json:"lastPublished,omitempty"`
	// Timestamp when the OIN Integration instance was last updated
	LastUpdated *string `json:"lastUpdated,omitempty"`
	// ID of the user who made the last update
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// URL to an uploaded application logo. This logo appears next to your app integration name in the OIN catalog. You must first [Upload an OIN Integration logo](/openapi/okta-management/management/tag/YourOinIntegrations/#tag/YourOinIntegrations/operation/uploadSubmissionLogo) to obtain the logo URL before you can specify this value.
	Logo *string `json:"logo,omitempty"`
	// The app integration name. This is the main title used for your integration in the OIN catalog.
	Name *string `json:"name,omitempty"`
	Sso *Sso `json:"sso,omitempty"`
	// Status of the OIN Integration submission
	Status *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubmissionResponse SubmissionResponse

// NewSubmissionResponse instantiates a new SubmissionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmissionResponse() *SubmissionResponse {
	this := SubmissionResponse{}
	return &this
}

// NewSubmissionResponseWithDefaults instantiates a new SubmissionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmissionResponseWithDefaults() *SubmissionResponse {
	this := SubmissionResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SubmissionResponse) GetConfig() []SubmissionResponseConfigInner {
	if o == nil || o.Config == nil {
		var ret []SubmissionResponseConfigInner
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetConfigOk() ([]SubmissionResponseConfigInner, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SubmissionResponse) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given []SubmissionResponseConfigInner and assigns it to the Config field.
func (o *SubmissionResponse) SetConfig(v []SubmissionResponseConfigInner) {
	o.Config = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SubmissionResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SubmissionResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SubmissionResponse) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubmissionResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubmissionResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SubmissionResponse) SetId(v string) {
	o.Id = &v
}

// GetLastPublished returns the LastPublished field value if set, zero value otherwise.
func (o *SubmissionResponse) GetLastPublished() string {
	if o == nil || o.LastPublished == nil {
		var ret string
		return ret
	}
	return *o.LastPublished
}

// GetLastPublishedOk returns a tuple with the LastPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetLastPublishedOk() (*string, bool) {
	if o == nil || o.LastPublished == nil {
		return nil, false
	}
	return o.LastPublished, true
}

// HasLastPublished returns a boolean if a field has been set.
func (o *SubmissionResponse) HasLastPublished() bool {
	if o != nil && o.LastPublished != nil {
		return true
	}

	return false
}

// SetLastPublished gets a reference to the given string and assigns it to the LastPublished field.
func (o *SubmissionResponse) SetLastPublished(v string) {
	o.LastPublished = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *SubmissionResponse) GetLastUpdated() string {
	if o == nil || o.LastUpdated == nil {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetLastUpdatedOk() (*string, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *SubmissionResponse) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *SubmissionResponse) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *SubmissionResponse) GetLastUpdatedBy() string {
	if o == nil || o.LastUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *SubmissionResponse) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *SubmissionResponse) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *SubmissionResponse) GetLogo() string {
	if o == nil || o.Logo == nil {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetLogoOk() (*string, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *SubmissionResponse) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *SubmissionResponse) SetLogo(v string) {
	o.Logo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubmissionResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubmissionResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubmissionResponse) SetName(v string) {
	o.Name = &v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *SubmissionResponse) GetSso() Sso {
	if o == nil || o.Sso == nil {
		var ret Sso
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetSsoOk() (*Sso, bool) {
	if o == nil || o.Sso == nil {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *SubmissionResponse) HasSso() bool {
	if o != nil && o.Sso != nil {
		return true
	}

	return false
}

// SetSso gets a reference to the given Sso and assigns it to the Sso field.
func (o *SubmissionResponse) SetSso(v Sso) {
	o.Sso = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SubmissionResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmissionResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SubmissionResponse) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SubmissionResponse) SetStatus(v string) {
	o.Status = &v
}

func (o SubmissionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastPublished != nil {
		toSerialize["lastPublished"] = o.LastPublished
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Sso != nil {
		toSerialize["sso"] = o.Sso
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SubmissionResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSubmissionResponse := _SubmissionResponse{}

	err = json.Unmarshal(bytes, &varSubmissionResponse)
	if err == nil {
		*o = SubmissionResponse(varSubmissionResponse)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "config")
		delete(additionalProperties, "description")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastPublished")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "name")
		delete(additionalProperties, "sso")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSubmissionResponse struct {
	value *SubmissionResponse
	isSet bool
}

func (v NullableSubmissionResponse) Get() *SubmissionResponse {
	return v.value
}

func (v *NullableSubmissionResponse) Set(val *SubmissionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmissionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmissionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmissionResponse(val *SubmissionResponse) *NullableSubmissionResponse {
	return &NullableSubmissionResponse{value: val, isSet: true}
}

func (v NullableSubmissionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmissionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

