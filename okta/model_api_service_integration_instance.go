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

// APIServiceIntegrationInstance struct for APIServiceIntegrationInstance
type APIServiceIntegrationInstance struct {
	// The URL to the API service integration configuration guide
	ConfigGuideUrl *string `json:"configGuideUrl,omitempty"`
	// Timestamp when the API Service Integration instance was created
	CreatedAt *string `json:"createdAt,omitempty"`
	// The user ID of the API Service Integration instance creator
	CreatedBy *string `json:"createdBy,omitempty"`
	// The list of Okta management scopes granted to the API Service Integration instance. See [Okta management OAuth 2.0 scopes](/oauth2/#okta-admin-management).
	GrantedScopes []string `json:"grantedScopes,omitempty"`
	// The ID of the API Service Integration instance
	Id *string `json:"id,omitempty"`
	// The name of the API service integration that corresponds with the `type` property. This is the full name of the API service integration listed in the Okta Integration Network (OIN) catalog.
	Name *string `json:"name,omitempty"`
	// The type of the API service integration. This string is an underscore-concatenated, lowercased API service integration name. For example, `my_api_log_integration`.
	Type *string `json:"type,omitempty"`
	Links *APIServiceIntegrationLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _APIServiceIntegrationInstance APIServiceIntegrationInstance

// NewAPIServiceIntegrationInstance instantiates a new APIServiceIntegrationInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServiceIntegrationInstance() *APIServiceIntegrationInstance {
	this := APIServiceIntegrationInstance{}
	return &this
}

// NewAPIServiceIntegrationInstanceWithDefaults instantiates a new APIServiceIntegrationInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServiceIntegrationInstanceWithDefaults() *APIServiceIntegrationInstance {
	this := APIServiceIntegrationInstance{}
	return &this
}

// GetConfigGuideUrl returns the ConfigGuideUrl field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetConfigGuideUrl() string {
	if o == nil || o.ConfigGuideUrl == nil {
		var ret string
		return ret
	}
	return *o.ConfigGuideUrl
}

// GetConfigGuideUrlOk returns a tuple with the ConfigGuideUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetConfigGuideUrlOk() (*string, bool) {
	if o == nil || o.ConfigGuideUrl == nil {
		return nil, false
	}
	return o.ConfigGuideUrl, true
}

// HasConfigGuideUrl returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasConfigGuideUrl() bool {
	if o != nil && o.ConfigGuideUrl != nil {
		return true
	}

	return false
}

// SetConfigGuideUrl gets a reference to the given string and assigns it to the ConfigGuideUrl field.
func (o *APIServiceIntegrationInstance) SetConfigGuideUrl(v string) {
	o.ConfigGuideUrl = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *APIServiceIntegrationInstance) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *APIServiceIntegrationInstance) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetGrantedScopes returns the GrantedScopes field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetGrantedScopes() []string {
	if o == nil || o.GrantedScopes == nil {
		var ret []string
		return ret
	}
	return o.GrantedScopes
}

// GetGrantedScopesOk returns a tuple with the GrantedScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetGrantedScopesOk() ([]string, bool) {
	if o == nil || o.GrantedScopes == nil {
		return nil, false
	}
	return o.GrantedScopes, true
}

// HasGrantedScopes returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasGrantedScopes() bool {
	if o != nil && o.GrantedScopes != nil {
		return true
	}

	return false
}

// SetGrantedScopes gets a reference to the given []string and assigns it to the GrantedScopes field.
func (o *APIServiceIntegrationInstance) SetGrantedScopes(v []string) {
	o.GrantedScopes = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APIServiceIntegrationInstance) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *APIServiceIntegrationInstance) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *APIServiceIntegrationInstance) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *APIServiceIntegrationInstance) GetLinks() APIServiceIntegrationLinks {
	if o == nil || o.Links == nil {
		var ret APIServiceIntegrationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServiceIntegrationInstance) GetLinksOk() (*APIServiceIntegrationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *APIServiceIntegrationInstance) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given APIServiceIntegrationLinks and assigns it to the Links field.
func (o *APIServiceIntegrationInstance) SetLinks(v APIServiceIntegrationLinks) {
	o.Links = &v
}

func (o APIServiceIntegrationInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigGuideUrl != nil {
		toSerialize["configGuideUrl"] = o.ConfigGuideUrl
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.GrantedScopes != nil {
		toSerialize["grantedScopes"] = o.GrantedScopes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *APIServiceIntegrationInstance) UnmarshalJSON(bytes []byte) (err error) {
	varAPIServiceIntegrationInstance := _APIServiceIntegrationInstance{}

	err = json.Unmarshal(bytes, &varAPIServiceIntegrationInstance)
	if err == nil {
		*o = APIServiceIntegrationInstance(varAPIServiceIntegrationInstance)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "configGuideUrl")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "grantedScopes")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAPIServiceIntegrationInstance struct {
	value *APIServiceIntegrationInstance
	isSet bool
}

func (v NullableAPIServiceIntegrationInstance) Get() *APIServiceIntegrationInstance {
	return v.value
}

func (v *NullableAPIServiceIntegrationInstance) Set(val *APIServiceIntegrationInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServiceIntegrationInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServiceIntegrationInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServiceIntegrationInstance(val *APIServiceIntegrationInstance) *NullableAPIServiceIntegrationInstance {
	return &NullableAPIServiceIntegrationInstance{value: val, isSet: true}
}

func (v NullableAPIServiceIntegrationInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServiceIntegrationInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

