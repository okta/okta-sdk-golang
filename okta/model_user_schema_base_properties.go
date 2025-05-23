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

// UserSchemaBaseProperties struct for UserSchemaBaseProperties
type UserSchemaBaseProperties struct {
	City *UserSchemaAttribute `json:"city,omitempty"`
	CostCenter *UserSchemaAttribute `json:"costCenter,omitempty"`
	CountryCode *UserSchemaAttribute `json:"countryCode,omitempty"`
	Department *UserSchemaAttribute `json:"department,omitempty"`
	DisplayName *UserSchemaAttribute `json:"displayName,omitempty"`
	Division *UserSchemaAttribute `json:"division,omitempty"`
	Email *UserSchemaAttribute `json:"email,omitempty"`
	EmployeeNumber *UserSchemaAttribute `json:"employeeNumber,omitempty"`
	FirstName *UserSchemaAttribute `json:"firstName,omitempty"`
	HonorificPrefix *UserSchemaAttribute `json:"honorificPrefix,omitempty"`
	HonorificSuffix *UserSchemaAttribute `json:"honorificSuffix,omitempty"`
	LastName *UserSchemaAttribute `json:"lastName,omitempty"`
	Locale *UserSchemaAttribute `json:"locale,omitempty"`
	Login *UserSchemaAttribute `json:"login,omitempty"`
	Manager *UserSchemaAttribute `json:"manager,omitempty"`
	ManagerId *UserSchemaAttribute `json:"managerId,omitempty"`
	MiddleName *UserSchemaAttribute `json:"middleName,omitempty"`
	MobilePhone *UserSchemaAttribute `json:"mobilePhone,omitempty"`
	NickName *UserSchemaAttribute `json:"nickName,omitempty"`
	Organization *UserSchemaAttribute `json:"organization,omitempty"`
	PostalAddress *UserSchemaAttribute `json:"postalAddress,omitempty"`
	PreferredLanguage *UserSchemaAttribute `json:"preferredLanguage,omitempty"`
	PrimaryPhone *UserSchemaAttribute `json:"primaryPhone,omitempty"`
	ProfileUrl *UserSchemaAttribute `json:"profileUrl,omitempty"`
	SecondEmail *UserSchemaAttribute `json:"secondEmail,omitempty"`
	State *UserSchemaAttribute `json:"state,omitempty"`
	StreetAddress *UserSchemaAttribute `json:"streetAddress,omitempty"`
	Timezone *UserSchemaAttribute `json:"timezone,omitempty"`
	Title *UserSchemaAttribute `json:"title,omitempty"`
	UserType *UserSchemaAttribute `json:"userType,omitempty"`
	ZipCode *UserSchemaAttribute `json:"zipCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserSchemaBaseProperties UserSchemaBaseProperties

// NewUserSchemaBaseProperties instantiates a new UserSchemaBaseProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSchemaBaseProperties() *UserSchemaBaseProperties {
	this := UserSchemaBaseProperties{}
	return &this
}

// NewUserSchemaBasePropertiesWithDefaults instantiates a new UserSchemaBaseProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSchemaBasePropertiesWithDefaults() *UserSchemaBaseProperties {
	this := UserSchemaBaseProperties{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetCity() UserSchemaAttribute {
	if o == nil || o.City == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetCityOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given UserSchemaAttribute and assigns it to the City field.
func (o *UserSchemaBaseProperties) SetCity(v UserSchemaAttribute) {
	o.City = &v
}

// GetCostCenter returns the CostCenter field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetCostCenter() UserSchemaAttribute {
	if o == nil || o.CostCenter == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.CostCenter
}

// GetCostCenterOk returns a tuple with the CostCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetCostCenterOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.CostCenter == nil {
		return nil, false
	}
	return o.CostCenter, true
}

// HasCostCenter returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasCostCenter() bool {
	if o != nil && o.CostCenter != nil {
		return true
	}

	return false
}

// SetCostCenter gets a reference to the given UserSchemaAttribute and assigns it to the CostCenter field.
func (o *UserSchemaBaseProperties) SetCostCenter(v UserSchemaAttribute) {
	o.CostCenter = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetCountryCode() UserSchemaAttribute {
	if o == nil || o.CountryCode == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetCountryCodeOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.CountryCode == nil {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasCountryCode() bool {
	if o != nil && o.CountryCode != nil {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given UserSchemaAttribute and assigns it to the CountryCode field.
func (o *UserSchemaBaseProperties) SetCountryCode(v UserSchemaAttribute) {
	o.CountryCode = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetDepartment() UserSchemaAttribute {
	if o == nil || o.Department == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetDepartmentOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given UserSchemaAttribute and assigns it to the Department field.
func (o *UserSchemaBaseProperties) SetDepartment(v UserSchemaAttribute) {
	o.Department = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetDisplayName() UserSchemaAttribute {
	if o == nil || o.DisplayName == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetDisplayNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given UserSchemaAttribute and assigns it to the DisplayName field.
func (o *UserSchemaBaseProperties) SetDisplayName(v UserSchemaAttribute) {
	o.DisplayName = &v
}

// GetDivision returns the Division field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetDivision() UserSchemaAttribute {
	if o == nil || o.Division == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Division
}

// GetDivisionOk returns a tuple with the Division field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetDivisionOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Division == nil {
		return nil, false
	}
	return o.Division, true
}

// HasDivision returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasDivision() bool {
	if o != nil && o.Division != nil {
		return true
	}

	return false
}

// SetDivision gets a reference to the given UserSchemaAttribute and assigns it to the Division field.
func (o *UserSchemaBaseProperties) SetDivision(v UserSchemaAttribute) {
	o.Division = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetEmail() UserSchemaAttribute {
	if o == nil || o.Email == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetEmailOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given UserSchemaAttribute and assigns it to the Email field.
func (o *UserSchemaBaseProperties) SetEmail(v UserSchemaAttribute) {
	o.Email = &v
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetEmployeeNumber() UserSchemaAttribute {
	if o == nil || o.EmployeeNumber == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.EmployeeNumber
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetEmployeeNumberOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.EmployeeNumber == nil {
		return nil, false
	}
	return o.EmployeeNumber, true
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasEmployeeNumber() bool {
	if o != nil && o.EmployeeNumber != nil {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given UserSchemaAttribute and assigns it to the EmployeeNumber field.
func (o *UserSchemaBaseProperties) SetEmployeeNumber(v UserSchemaAttribute) {
	o.EmployeeNumber = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetFirstName() UserSchemaAttribute {
	if o == nil || o.FirstName == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetFirstNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given UserSchemaAttribute and assigns it to the FirstName field.
func (o *UserSchemaBaseProperties) SetFirstName(v UserSchemaAttribute) {
	o.FirstName = &v
}

// GetHonorificPrefix returns the HonorificPrefix field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetHonorificPrefix() UserSchemaAttribute {
	if o == nil || o.HonorificPrefix == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.HonorificPrefix
}

// GetHonorificPrefixOk returns a tuple with the HonorificPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetHonorificPrefixOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.HonorificPrefix == nil {
		return nil, false
	}
	return o.HonorificPrefix, true
}

// HasHonorificPrefix returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasHonorificPrefix() bool {
	if o != nil && o.HonorificPrefix != nil {
		return true
	}

	return false
}

// SetHonorificPrefix gets a reference to the given UserSchemaAttribute and assigns it to the HonorificPrefix field.
func (o *UserSchemaBaseProperties) SetHonorificPrefix(v UserSchemaAttribute) {
	o.HonorificPrefix = &v
}

// GetHonorificSuffix returns the HonorificSuffix field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetHonorificSuffix() UserSchemaAttribute {
	if o == nil || o.HonorificSuffix == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.HonorificSuffix
}

// GetHonorificSuffixOk returns a tuple with the HonorificSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetHonorificSuffixOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.HonorificSuffix == nil {
		return nil, false
	}
	return o.HonorificSuffix, true
}

// HasHonorificSuffix returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasHonorificSuffix() bool {
	if o != nil && o.HonorificSuffix != nil {
		return true
	}

	return false
}

// SetHonorificSuffix gets a reference to the given UserSchemaAttribute and assigns it to the HonorificSuffix field.
func (o *UserSchemaBaseProperties) SetHonorificSuffix(v UserSchemaAttribute) {
	o.HonorificSuffix = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetLastName() UserSchemaAttribute {
	if o == nil || o.LastName == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetLastNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given UserSchemaAttribute and assigns it to the LastName field.
func (o *UserSchemaBaseProperties) SetLastName(v UserSchemaAttribute) {
	o.LastName = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetLocale() UserSchemaAttribute {
	if o == nil || o.Locale == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetLocaleOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given UserSchemaAttribute and assigns it to the Locale field.
func (o *UserSchemaBaseProperties) SetLocale(v UserSchemaAttribute) {
	o.Locale = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetLogin() UserSchemaAttribute {
	if o == nil || o.Login == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetLoginOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given UserSchemaAttribute and assigns it to the Login field.
func (o *UserSchemaBaseProperties) SetLogin(v UserSchemaAttribute) {
	o.Login = &v
}

// GetManager returns the Manager field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetManager() UserSchemaAttribute {
	if o == nil || o.Manager == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetManagerOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Manager == nil {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasManager() bool {
	if o != nil && o.Manager != nil {
		return true
	}

	return false
}

// SetManager gets a reference to the given UserSchemaAttribute and assigns it to the Manager field.
func (o *UserSchemaBaseProperties) SetManager(v UserSchemaAttribute) {
	o.Manager = &v
}

// GetManagerId returns the ManagerId field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetManagerId() UserSchemaAttribute {
	if o == nil || o.ManagerId == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.ManagerId
}

// GetManagerIdOk returns a tuple with the ManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetManagerIdOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.ManagerId == nil {
		return nil, false
	}
	return o.ManagerId, true
}

// HasManagerId returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasManagerId() bool {
	if o != nil && o.ManagerId != nil {
		return true
	}

	return false
}

// SetManagerId gets a reference to the given UserSchemaAttribute and assigns it to the ManagerId field.
func (o *UserSchemaBaseProperties) SetManagerId(v UserSchemaAttribute) {
	o.ManagerId = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetMiddleName() UserSchemaAttribute {
	if o == nil || o.MiddleName == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetMiddleNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.MiddleName == nil {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasMiddleName() bool {
	if o != nil && o.MiddleName != nil {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given UserSchemaAttribute and assigns it to the MiddleName field.
func (o *UserSchemaBaseProperties) SetMiddleName(v UserSchemaAttribute) {
	o.MiddleName = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetMobilePhone() UserSchemaAttribute {
	if o == nil || o.MobilePhone == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetMobilePhoneOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.MobilePhone == nil {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasMobilePhone() bool {
	if o != nil && o.MobilePhone != nil {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given UserSchemaAttribute and assigns it to the MobilePhone field.
func (o *UserSchemaBaseProperties) SetMobilePhone(v UserSchemaAttribute) {
	o.MobilePhone = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetNickName() UserSchemaAttribute {
	if o == nil || o.NickName == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetNickNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.NickName == nil {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasNickName() bool {
	if o != nil && o.NickName != nil {
		return true
	}

	return false
}

// SetNickName gets a reference to the given UserSchemaAttribute and assigns it to the NickName field.
func (o *UserSchemaBaseProperties) SetNickName(v UserSchemaAttribute) {
	o.NickName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetOrganization() UserSchemaAttribute {
	if o == nil || o.Organization == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetOrganizationOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given UserSchemaAttribute and assigns it to the Organization field.
func (o *UserSchemaBaseProperties) SetOrganization(v UserSchemaAttribute) {
	o.Organization = &v
}

// GetPostalAddress returns the PostalAddress field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetPostalAddress() UserSchemaAttribute {
	if o == nil || o.PostalAddress == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.PostalAddress
}

// GetPostalAddressOk returns a tuple with the PostalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetPostalAddressOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.PostalAddress == nil {
		return nil, false
	}
	return o.PostalAddress, true
}

// HasPostalAddress returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasPostalAddress() bool {
	if o != nil && o.PostalAddress != nil {
		return true
	}

	return false
}

// SetPostalAddress gets a reference to the given UserSchemaAttribute and assigns it to the PostalAddress field.
func (o *UserSchemaBaseProperties) SetPostalAddress(v UserSchemaAttribute) {
	o.PostalAddress = &v
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetPreferredLanguage() UserSchemaAttribute {
	if o == nil || o.PreferredLanguage == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetPreferredLanguageOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.PreferredLanguage == nil {
		return nil, false
	}
	return o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage != nil {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given UserSchemaAttribute and assigns it to the PreferredLanguage field.
func (o *UserSchemaBaseProperties) SetPreferredLanguage(v UserSchemaAttribute) {
	o.PreferredLanguage = &v
}

// GetPrimaryPhone returns the PrimaryPhone field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetPrimaryPhone() UserSchemaAttribute {
	if o == nil || o.PrimaryPhone == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.PrimaryPhone
}

// GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetPrimaryPhoneOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.PrimaryPhone == nil {
		return nil, false
	}
	return o.PrimaryPhone, true
}

// HasPrimaryPhone returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasPrimaryPhone() bool {
	if o != nil && o.PrimaryPhone != nil {
		return true
	}

	return false
}

// SetPrimaryPhone gets a reference to the given UserSchemaAttribute and assigns it to the PrimaryPhone field.
func (o *UserSchemaBaseProperties) SetPrimaryPhone(v UserSchemaAttribute) {
	o.PrimaryPhone = &v
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetProfileUrl() UserSchemaAttribute {
	if o == nil || o.ProfileUrl == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.ProfileUrl
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetProfileUrlOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.ProfileUrl == nil {
		return nil, false
	}
	return o.ProfileUrl, true
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl != nil {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given UserSchemaAttribute and assigns it to the ProfileUrl field.
func (o *UserSchemaBaseProperties) SetProfileUrl(v UserSchemaAttribute) {
	o.ProfileUrl = &v
}

// GetSecondEmail returns the SecondEmail field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetSecondEmail() UserSchemaAttribute {
	if o == nil || o.SecondEmail == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.SecondEmail
}

// GetSecondEmailOk returns a tuple with the SecondEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetSecondEmailOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.SecondEmail == nil {
		return nil, false
	}
	return o.SecondEmail, true
}

// HasSecondEmail returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasSecondEmail() bool {
	if o != nil && o.SecondEmail != nil {
		return true
	}

	return false
}

// SetSecondEmail gets a reference to the given UserSchemaAttribute and assigns it to the SecondEmail field.
func (o *UserSchemaBaseProperties) SetSecondEmail(v UserSchemaAttribute) {
	o.SecondEmail = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetState() UserSchemaAttribute {
	if o == nil || o.State == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetStateOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given UserSchemaAttribute and assigns it to the State field.
func (o *UserSchemaBaseProperties) SetState(v UserSchemaAttribute) {
	o.State = &v
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetStreetAddress() UserSchemaAttribute {
	if o == nil || o.StreetAddress == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetStreetAddressOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.StreetAddress == nil {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasStreetAddress() bool {
	if o != nil && o.StreetAddress != nil {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given UserSchemaAttribute and assigns it to the StreetAddress field.
func (o *UserSchemaBaseProperties) SetStreetAddress(v UserSchemaAttribute) {
	o.StreetAddress = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetTimezone() UserSchemaAttribute {
	if o == nil || o.Timezone == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetTimezoneOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given UserSchemaAttribute and assigns it to the Timezone field.
func (o *UserSchemaBaseProperties) SetTimezone(v UserSchemaAttribute) {
	o.Timezone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetTitle() UserSchemaAttribute {
	if o == nil || o.Title == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetTitleOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given UserSchemaAttribute and assigns it to the Title field.
func (o *UserSchemaBaseProperties) SetTitle(v UserSchemaAttribute) {
	o.Title = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetUserType() UserSchemaAttribute {
	if o == nil || o.UserType == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetUserTypeOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserSchemaAttribute and assigns it to the UserType field.
func (o *UserSchemaBaseProperties) SetUserType(v UserSchemaAttribute) {
	o.UserType = &v
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise.
func (o *UserSchemaBaseProperties) GetZipCode() UserSchemaAttribute {
	if o == nil || o.ZipCode == nil {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetZipCodeOk() (*UserSchemaAttribute, bool) {
	if o == nil || o.ZipCode == nil {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasZipCode() bool {
	if o != nil && o.ZipCode != nil {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given UserSchemaAttribute and assigns it to the ZipCode field.
func (o *UserSchemaBaseProperties) SetZipCode(v UserSchemaAttribute) {
	o.ZipCode = &v
}

func (o UserSchemaBaseProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.CostCenter != nil {
		toSerialize["costCenter"] = o.CostCenter
	}
	if o.CountryCode != nil {
		toSerialize["countryCode"] = o.CountryCode
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Division != nil {
		toSerialize["division"] = o.Division
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EmployeeNumber != nil {
		toSerialize["employeeNumber"] = o.EmployeeNumber
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.HonorificPrefix != nil {
		toSerialize["honorificPrefix"] = o.HonorificPrefix
	}
	if o.HonorificSuffix != nil {
		toSerialize["honorificSuffix"] = o.HonorificSuffix
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Manager != nil {
		toSerialize["manager"] = o.Manager
	}
	if o.ManagerId != nil {
		toSerialize["managerId"] = o.ManagerId
	}
	if o.MiddleName != nil {
		toSerialize["middleName"] = o.MiddleName
	}
	if o.MobilePhone != nil {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if o.NickName != nil {
		toSerialize["nickName"] = o.NickName
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.PostalAddress != nil {
		toSerialize["postalAddress"] = o.PostalAddress
	}
	if o.PreferredLanguage != nil {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if o.PrimaryPhone != nil {
		toSerialize["primaryPhone"] = o.PrimaryPhone
	}
	if o.ProfileUrl != nil {
		toSerialize["profileUrl"] = o.ProfileUrl
	}
	if o.SecondEmail != nil {
		toSerialize["secondEmail"] = o.SecondEmail
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StreetAddress != nil {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.UserType != nil {
		toSerialize["userType"] = o.UserType
	}
	if o.ZipCode != nil {
		toSerialize["zipCode"] = o.ZipCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserSchemaBaseProperties) UnmarshalJSON(bytes []byte) (err error) {
	varUserSchemaBaseProperties := _UserSchemaBaseProperties{}

	err = json.Unmarshal(bytes, &varUserSchemaBaseProperties)
	if err == nil {
		*o = UserSchemaBaseProperties(varUserSchemaBaseProperties)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "costCenter")
		delete(additionalProperties, "countryCode")
		delete(additionalProperties, "department")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "division")
		delete(additionalProperties, "email")
		delete(additionalProperties, "employeeNumber")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "honorificPrefix")
		delete(additionalProperties, "honorificSuffix")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "locale")
		delete(additionalProperties, "login")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "managerId")
		delete(additionalProperties, "middleName")
		delete(additionalProperties, "mobilePhone")
		delete(additionalProperties, "nickName")
		delete(additionalProperties, "organization")
		delete(additionalProperties, "postalAddress")
		delete(additionalProperties, "preferredLanguage")
		delete(additionalProperties, "primaryPhone")
		delete(additionalProperties, "profileUrl")
		delete(additionalProperties, "secondEmail")
		delete(additionalProperties, "state")
		delete(additionalProperties, "streetAddress")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "title")
		delete(additionalProperties, "userType")
		delete(additionalProperties, "zipCode")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableUserSchemaBaseProperties struct {
	value *UserSchemaBaseProperties
	isSet bool
}

func (v NullableUserSchemaBaseProperties) Get() *UserSchemaBaseProperties {
	return v.value
}

func (v *NullableUserSchemaBaseProperties) Set(val *UserSchemaBaseProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSchemaBaseProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSchemaBaseProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSchemaBaseProperties(val *UserSchemaBaseProperties) *NullableUserSchemaBaseProperties {
	return &NullableUserSchemaBaseProperties{value: val, isSet: true}
}

func (v NullableUserSchemaBaseProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSchemaBaseProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

