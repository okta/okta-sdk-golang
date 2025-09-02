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

// ChildOrg struct for ChildOrg
type ChildOrg struct {
	Admin OrgCreationAdmin `json:"admin"`
	// Timestamp when the org was created
	Created *time.Time `json:"created,omitempty"`
	// Edition for the org. `SKU` is the only supported value.
	Edition string `json:"edition"`
	// Org ID
	Id *string `json:"id,omitempty"`
	// Timestamp when the org was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Unique name of the org. This name appears in the HTML `<title>` tag of the new org sign-in page. Only less than 4-width UTF-8 encoded characters are allowed.
	Name string `json:"name"`
	// Settings associated with the created org
	Settings map[string]interface{} `json:"settings,omitempty"`
	// Status of the org. `ACTIVE` is returned after the org is created.
	Status *string `json:"status,omitempty"`
	// Subdomain of the org. Must be unique and include no spaces.
	Subdomain string `json:"subdomain"`
	// API token associated with the child org super admin account. Use this API token to provision resources (such as policies, apps, and groups) on the newly created child org. This token is revoked if the super admin account is deactivated. > **Note:** If this API token expires, sign in to the Admin Console as the super admin user and create a new API token. See [Create an API token](https://developer.okta.com/docs/guides/create-an-api-token/).
	Token *string `json:"token,omitempty"`
	// Type of returned `token`. See [Okta API tokens](https://developer.okta.com/docs/guides/create-an-api-token/main/#okta-api-tokens).
	TokenType *string `json:"tokenType,omitempty"`
	// Default website for the org
	Website *string `json:"website,omitempty"`
	// Specifies available link relations (see [Web Linking](https://www.rfc-editor.org/rfc/rfc8288)) using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification
	Links map[string]interface{} `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChildOrg ChildOrg

// NewChildOrg instantiates a new ChildOrg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChildOrg(admin OrgCreationAdmin, edition string, name string, subdomain string) *ChildOrg {
	this := ChildOrg{}
	this.Admin = admin
	this.Edition = edition
	this.Name = name
	this.Subdomain = subdomain
	return &this
}

// NewChildOrgWithDefaults instantiates a new ChildOrg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChildOrgWithDefaults() *ChildOrg {
	this := ChildOrg{}
	return &this
}

// GetAdmin returns the Admin field value
func (o *ChildOrg) GetAdmin() OrgCreationAdmin {
	if o == nil {
		var ret OrgCreationAdmin
		return ret
	}

	return o.Admin
}

// GetAdminOk returns a tuple with the Admin field value
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetAdminOk() (*OrgCreationAdmin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Admin, true
}

// SetAdmin sets field value
func (o *ChildOrg) SetAdmin(v OrgCreationAdmin) {
	o.Admin = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ChildOrg) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ChildOrg) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *ChildOrg) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEdition returns the Edition field value
func (o *ChildOrg) GetEdition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Edition
}

// GetEditionOk returns a tuple with the Edition field value
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetEditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edition, true
}

// SetEdition sets field value
func (o *ChildOrg) SetEdition(v string) {
	o.Edition = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChildOrg) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChildOrg) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChildOrg) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *ChildOrg) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || o.LastUpdated == nil {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *ChildOrg) HasLastUpdated() bool {
	if o != nil && o.LastUpdated != nil {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *ChildOrg) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value
func (o *ChildOrg) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ChildOrg) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *ChildOrg) GetSettings() map[string]interface{} {
	if o == nil || o.Settings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetSettingsOk() (map[string]interface{}, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *ChildOrg) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given map[string]interface{} and assigns it to the Settings field.
func (o *ChildOrg) SetSettings(v map[string]interface{}) {
	o.Settings = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ChildOrg) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ChildOrg) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ChildOrg) SetStatus(v string) {
	o.Status = &v
}

// GetSubdomain returns the Subdomain field value
func (o *ChildOrg) GetSubdomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetSubdomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subdomain, true
}

// SetSubdomain sets field value
func (o *ChildOrg) SetSubdomain(v string) {
	o.Subdomain = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ChildOrg) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ChildOrg) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ChildOrg) SetToken(v string) {
	o.Token = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *ChildOrg) GetTokenType() string {
	if o == nil || o.TokenType == nil {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetTokenTypeOk() (*string, bool) {
	if o == nil || o.TokenType == nil {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *ChildOrg) HasTokenType() bool {
	if o != nil && o.TokenType != nil {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *ChildOrg) SetTokenType(v string) {
	o.TokenType = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *ChildOrg) GetWebsite() string {
	if o == nil || o.Website == nil {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetWebsiteOk() (*string, bool) {
	if o == nil || o.Website == nil {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *ChildOrg) HasWebsite() bool {
	if o != nil && o.Website != nil {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *ChildOrg) SetWebsite(v string) {
	o.Website = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChildOrg) GetLinks() map[string]interface{} {
	if o == nil || o.Links == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChildOrg) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChildOrg) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *ChildOrg) SetLinks(v map[string]interface{}) {
	o.Links = v
}

func (o ChildOrg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["admin"] = o.Admin
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["edition"] = o.Edition
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdated != nil {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["subdomain"] = o.Subdomain
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.TokenType != nil {
		toSerialize["tokenType"] = o.TokenType
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

func (o *ChildOrg) UnmarshalJSON(bytes []byte) (err error) {
	varChildOrg := _ChildOrg{}

	err = json.Unmarshal(bytes, &varChildOrg)
	if err == nil {
		*o = ChildOrg(varChildOrg)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "admin")
		delete(additionalProperties, "created")
		delete(additionalProperties, "edition")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "settings")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "token")
		delete(additionalProperties, "tokenType")
		delete(additionalProperties, "website")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableChildOrg struct {
	value *ChildOrg
	isSet bool
}

func (v NullableChildOrg) Get() *ChildOrg {
	return v.value
}

func (v *NullableChildOrg) Set(val *ChildOrg) {
	v.value = val
	v.isSet = true
}

func (v NullableChildOrg) IsSet() bool {
	return v.isSet
}

func (v *NullableChildOrg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChildOrg(val *ChildOrg) *NullableChildOrg {
	return &NullableChildOrg{value: val, isSet: true}
}

func (v NullableChildOrg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChildOrg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

