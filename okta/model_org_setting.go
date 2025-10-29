/*
Okta Admin Management

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

// checks if the OrgSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSetting{}

// OrgSetting struct for OrgSetting
type OrgSetting struct {
	// Primary address of the organization associated with the org
	Address1 *string `json:"address1,omitempty"`
	// Secondary address of the organization associated with the org
	Address2 *string `json:"address2,omitempty"`
	// City of the organization associated with the org
	City *string `json:"city,omitempty"`
	// Name of org
	CompanyName *string `json:"companyName,omitempty"`
	// County of the organization associated with the org
	Country *string `json:"country,omitempty"`
	// When org was created
	Created *time.Time `json:"created,omitempty"`
	// Support link of org
	EndUserSupportHelpURL *string `json:"endUserSupportHelpURL,omitempty"`
	// Expiration of org
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// Org ID
	Id *string `json:"id,omitempty"`
	// When org was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// Phone number of the organization associated with the org
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// Postal code of the organization associated with the org
	PostalCode *string `json:"postalCode,omitempty"`
	// State of the organization associated with the org
	State *string `json:"state,omitempty"`
	// Status of org
	Status *string `json:"status,omitempty"`
	// Subdomain of org
	Subdomain *string `json:"subdomain,omitempty"`
	// Support help phone of the organization associated with the org
	SupportPhoneNumber *string `json:"supportPhoneNumber,omitempty"`
	// Website of the organization associated with the org
	Website              *string                 `json:"website,omitempty"`
	Links                *OrgGeneralSettingLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSetting OrgSetting

// NewOrgSetting instantiates a new OrgSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSetting() *OrgSetting {
	this := OrgSetting{}
	return &this
}

// NewOrgSettingWithDefaults instantiates a new OrgSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingWithDefaults() *OrgSetting {
	this := OrgSetting{}
	return &this
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *OrgSetting) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *OrgSetting) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *OrgSetting) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *OrgSetting) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *OrgSetting) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *OrgSetting) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *OrgSetting) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *OrgSetting) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *OrgSetting) SetCity(v string) {
	o.City = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *OrgSetting) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *OrgSetting) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *OrgSetting) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *OrgSetting) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *OrgSetting) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *OrgSetting) SetCountry(v string) {
	o.Country = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrgSetting) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrgSetting) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OrgSetting) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEndUserSupportHelpURL returns the EndUserSupportHelpURL field value if set, zero value otherwise.
func (o *OrgSetting) GetEndUserSupportHelpURL() string {
	if o == nil || IsNil(o.EndUserSupportHelpURL) {
		var ret string
		return ret
	}
	return *o.EndUserSupportHelpURL
}

// GetEndUserSupportHelpURLOk returns a tuple with the EndUserSupportHelpURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetEndUserSupportHelpURLOk() (*string, bool) {
	if o == nil || IsNil(o.EndUserSupportHelpURL) {
		return nil, false
	}
	return o.EndUserSupportHelpURL, true
}

// HasEndUserSupportHelpURL returns a boolean if a field has been set.
func (o *OrgSetting) HasEndUserSupportHelpURL() bool {
	if o != nil && !IsNil(o.EndUserSupportHelpURL) {
		return true
	}

	return false
}

// SetEndUserSupportHelpURL gets a reference to the given string and assigns it to the EndUserSupportHelpURL field.
func (o *OrgSetting) SetEndUserSupportHelpURL(v string) {
	o.EndUserSupportHelpURL = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *OrgSetting) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *OrgSetting) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *OrgSetting) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgSetting) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgSetting) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgSetting) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *OrgSetting) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *OrgSetting) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *OrgSetting) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *OrgSetting) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *OrgSetting) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *OrgSetting) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *OrgSetting) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *OrgSetting) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *OrgSetting) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *OrgSetting) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *OrgSetting) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *OrgSetting) SetState(v string) {
	o.State = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrgSetting) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrgSetting) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrgSetting) SetStatus(v string) {
	o.Status = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *OrgSetting) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *OrgSetting) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *OrgSetting) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetSupportPhoneNumber returns the SupportPhoneNumber field value if set, zero value otherwise.
func (o *OrgSetting) GetSupportPhoneNumber() string {
	if o == nil || IsNil(o.SupportPhoneNumber) {
		var ret string
		return ret
	}
	return *o.SupportPhoneNumber
}

// GetSupportPhoneNumberOk returns a tuple with the SupportPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetSupportPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SupportPhoneNumber) {
		return nil, false
	}
	return o.SupportPhoneNumber, true
}

// HasSupportPhoneNumber returns a boolean if a field has been set.
func (o *OrgSetting) HasSupportPhoneNumber() bool {
	if o != nil && !IsNil(o.SupportPhoneNumber) {
		return true
	}

	return false
}

// SetSupportPhoneNumber gets a reference to the given string and assigns it to the SupportPhoneNumber field.
func (o *OrgSetting) SetSupportPhoneNumber(v string) {
	o.SupportPhoneNumber = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *OrgSetting) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *OrgSetting) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *OrgSetting) SetWebsite(v string) {
	o.Website = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OrgSetting) GetLinks() OrgGeneralSettingLinks {
	if o == nil || IsNil(o.Links) {
		var ret OrgGeneralSettingLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSetting) GetLinksOk() (*OrgGeneralSettingLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OrgSetting) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OrgGeneralSettingLinks and assigns it to the Links field.
func (o *OrgSetting) SetLinks(v OrgGeneralSettingLinks) {
	o.Links = &v
}

func (o OrgSetting) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.EndUserSupportHelpURL) {
		toSerialize["endUserSupportHelpURL"] = o.EndUserSupportHelpURL
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postalCode"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.SupportPhoneNumber) {
		toSerialize["supportPhoneNumber"] = o.SupportPhoneNumber
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSetting) UnmarshalJSON(data []byte) (err error) {
	varOrgSetting := _OrgSetting{}

	err = json.Unmarshal(data, &varOrgSetting)

	if err != nil {
		return err
	}

	*o = OrgSetting(varOrgSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "address1")
		delete(additionalProperties, "address2")
		delete(additionalProperties, "city")
		delete(additionalProperties, "companyName")
		delete(additionalProperties, "country")
		delete(additionalProperties, "created")
		delete(additionalProperties, "endUserSupportHelpURL")
		delete(additionalProperties, "expiresAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "phoneNumber")
		delete(additionalProperties, "postalCode")
		delete(additionalProperties, "state")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subdomain")
		delete(additionalProperties, "supportPhoneNumber")
		delete(additionalProperties, "website")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSetting struct {
	value *OrgSetting
	isSet bool
}

func (v NullableOrgSetting) Get() *OrgSetting {
	return v.value
}

func (v *NullableOrgSetting) Set(val *OrgSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSetting(val *OrgSetting) *NullableOrgSetting {
	return &NullableOrgSetting{value: val, isSet: true}
}

func (v NullableOrgSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
