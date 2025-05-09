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

// UserProfile Specifies the default and custom profile properties for a user.  The default user profile is based on the [System for Cross-domain Identity Management: Core Schema](https://datatracker.ietf.org/doc/html/rfc7643). The only permitted customizations of the default profile are to update permissions, change whether the `firstName` and `lastName` properties are nullable, and specify a [pattern](https://developer.okta.com/docs/reference/api/schemas/#login-pattern-validation) for `login`. You can use the Profile Editor in the administrator UI or the [Schemas API](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UISchema/#tag/UISchema) to make schema modifications.  You can extend user profiles with custom properties. You must first add the custom property to the user profile schema before you reference it. You can use the Profile Editor in the Admin console or the [Schemas API](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UISchema/#tag/UISchema) to manage schema extensions.  Custom attributes may contain HTML tags. It's the client's responsibility to escape or encode this data before displaying it. Use [best-practices](https://cheatsheetseries.owasp.org/cheatsheets/Cross_Site_Scripting_Prevention_Cheat_Sheet.html) to prevent cross-site scripting.
type UserProfile struct {
	// The city or locality of the user's address (`locality`)
	City NullableString `json:"city,omitempty"`
	// Name of the cost center assigned to a user
	CostCenter NullableString `json:"costCenter,omitempty"`
	// The country name component of the user's address (`country`)
	CountryCode NullableString `json:"countryCode,omitempty"`
	// Name of the user's department
	Department *string `json:"department,omitempty"`
	// Name of the user suitable for display to end users
	DisplayName NullableString `json:"displayName,omitempty"`
	// Name of the user's division
	Division NullableString `json:"division,omitempty"`
	// The primary email address of the user. For validation, see [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3).
	Email *string `json:"email,omitempty"`
	// The organization or company assigned unique identifier for the user
	EmployeeNumber *string `json:"employeeNumber,omitempty"`
	// Given name of the user (`givenName`)
	FirstName NullableString `json:"firstName,omitempty"`
	// Honorific prefix(es) of the user, or title in most Western languages
	HonorificPrefix NullableString `json:"honorificPrefix,omitempty"`
	// Honorific suffix(es) of the user
	HonorificSuffix NullableString `json:"honorificSuffix,omitempty"`
	// The family name of the user (`familyName`)
	LastName NullableString `json:"lastName,omitempty"`
	// The user's default location for purposes of localizing items such as currency, date time format, numerical representations, and so on. A locale value is a concatenation of the ISO 639-1 two-letter language code, an underscore, and the ISO 3166-1 two-letter country code. For example, en_US specifies the language English and country US. This value is `en_US` by default.
	Locale *string `json:"locale,omitempty"`
	// The unique identifier for the user (`username`). For validation, see [Login pattern validation](https://developer.okta.com/docs/reference/api/schemas/#login-pattern-validation). See also [Okta login](https://developer.okta.com/docs/reference/api/users/#okta-login).
	Login *string `json:"login,omitempty"`
	// The `displayName` of the user's manager
	Manager NullableString `json:"manager,omitempty"`
	// The `id` of the user's manager
	ManagerId NullableString `json:"managerId,omitempty"`
	// The middle name of the user
	MiddleName NullableString `json:"middleName,omitempty"`
	// The mobile phone number of the user
	MobilePhone NullableString `json:"mobilePhone,omitempty"`
	// The casual way to address the user in real life
	NickName NullableString `json:"nickName,omitempty"`
	// Name of the the user's organization
	Organization NullableString `json:"organization,omitempty"`
	// Mailing address component of the user's address
	PostalAddress NullableString `json:"postalAddress,omitempty"`
	// The user's preferred written or spoken language
	PreferredLanguage NullableString `json:"preferredLanguage,omitempty"`
	// The primary phone number of the user such as a home number
	PrimaryPhone NullableString `json:"primaryPhone,omitempty"`
	// The URL of the user's online profile. For example, a web page. See [URL](https://datatracker.ietf.org/doc/html/rfc1808).
	ProfileUrl NullableString `json:"profileUrl,omitempty"`
	// The secondary email address of the user typically used for account recovery
	SecondEmail NullableString `json:"secondEmail,omitempty"`
	// The state or region component of the user's address (`region`)
	State NullableString `json:"state,omitempty"`
	// The full street address component of the user's address
	StreetAddress NullableString `json:"streetAddress,omitempty"`
	// The user's time zone
	Timezone NullableString `json:"timezone,omitempty"`
	// The user's title, such as Vice President
	Title NullableString `json:"title,omitempty"`
	// The property used to describe the organization-to-user relationship, such as employee or contractor
	UserType NullableString `json:"userType,omitempty"`
	// The ZIP code or postal code component of the user's address (`postalCode`)
	ZipCode NullableString `json:"zipCode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UserProfile UserProfile

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile() *UserProfile {
	this := UserProfile{}
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetCity() string {
	if o == nil || o.City.Get() == nil {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *UserProfile) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *UserProfile) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *UserProfile) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *UserProfile) UnsetCity() {
	o.City.Unset()
}

// GetCostCenter returns the CostCenter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetCostCenter() string {
	if o == nil || o.CostCenter.Get() == nil {
		var ret string
		return ret
	}
	return *o.CostCenter.Get()
}

// GetCostCenterOk returns a tuple with the CostCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetCostCenterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CostCenter.Get(), o.CostCenter.IsSet()
}

// HasCostCenter returns a boolean if a field has been set.
func (o *UserProfile) HasCostCenter() bool {
	if o != nil && o.CostCenter.IsSet() {
		return true
	}

	return false
}

// SetCostCenter gets a reference to the given NullableString and assigns it to the CostCenter field.
func (o *UserProfile) SetCostCenter(v string) {
	o.CostCenter.Set(&v)
}
// SetCostCenterNil sets the value for CostCenter to be an explicit nil
func (o *UserProfile) SetCostCenterNil() {
	o.CostCenter.Set(nil)
}

// UnsetCostCenter ensures that no value is present for CostCenter, not even an explicit nil
func (o *UserProfile) UnsetCostCenter() {
	o.CostCenter.Unset()
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UserProfile) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *UserProfile) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *UserProfile) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *UserProfile) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *UserProfile) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *UserProfile) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *UserProfile) SetDepartment(v string) {
	o.Department = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserProfile) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *UserProfile) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *UserProfile) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *UserProfile) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetDivision returns the Division field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetDivision() string {
	if o == nil || o.Division.Get() == nil {
		var ret string
		return ret
	}
	return *o.Division.Get()
}

// GetDivisionOk returns a tuple with the Division field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetDivisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Division.Get(), o.Division.IsSet()
}

// HasDivision returns a boolean if a field has been set.
func (o *UserProfile) HasDivision() bool {
	if o != nil && o.Division.IsSet() {
		return true
	}

	return false
}

// SetDivision gets a reference to the given NullableString and assigns it to the Division field.
func (o *UserProfile) SetDivision(v string) {
	o.Division.Set(&v)
}
// SetDivisionNil sets the value for Division to be an explicit nil
func (o *UserProfile) SetDivisionNil() {
	o.Division.Set(nil)
}

// UnsetDivision ensures that no value is present for Division, not even an explicit nil
func (o *UserProfile) UnsetDivision() {
	o.Division.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserProfile) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserProfile) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserProfile) SetEmail(v string) {
	o.Email = &v
}

// GetEmployeeNumber returns the EmployeeNumber field value if set, zero value otherwise.
func (o *UserProfile) GetEmployeeNumber() string {
	if o == nil || o.EmployeeNumber == nil {
		var ret string
		return ret
	}
	return *o.EmployeeNumber
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetEmployeeNumberOk() (*string, bool) {
	if o == nil || o.EmployeeNumber == nil {
		return nil, false
	}
	return o.EmployeeNumber, true
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *UserProfile) HasEmployeeNumber() bool {
	if o != nil && o.EmployeeNumber != nil {
		return true
	}

	return false
}

// SetEmployeeNumber gets a reference to the given string and assigns it to the EmployeeNumber field.
func (o *UserProfile) SetEmployeeNumber(v string) {
	o.EmployeeNumber = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetFirstName() string {
	if o == nil || o.FirstName.Get() == nil {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserProfile) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *UserProfile) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *UserProfile) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *UserProfile) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetHonorificPrefix returns the HonorificPrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetHonorificPrefix() string {
	if o == nil || o.HonorificPrefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.HonorificPrefix.Get()
}

// GetHonorificPrefixOk returns a tuple with the HonorificPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetHonorificPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HonorificPrefix.Get(), o.HonorificPrefix.IsSet()
}

// HasHonorificPrefix returns a boolean if a field has been set.
func (o *UserProfile) HasHonorificPrefix() bool {
	if o != nil && o.HonorificPrefix.IsSet() {
		return true
	}

	return false
}

// SetHonorificPrefix gets a reference to the given NullableString and assigns it to the HonorificPrefix field.
func (o *UserProfile) SetHonorificPrefix(v string) {
	o.HonorificPrefix.Set(&v)
}
// SetHonorificPrefixNil sets the value for HonorificPrefix to be an explicit nil
func (o *UserProfile) SetHonorificPrefixNil() {
	o.HonorificPrefix.Set(nil)
}

// UnsetHonorificPrefix ensures that no value is present for HonorificPrefix, not even an explicit nil
func (o *UserProfile) UnsetHonorificPrefix() {
	o.HonorificPrefix.Unset()
}

// GetHonorificSuffix returns the HonorificSuffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetHonorificSuffix() string {
	if o == nil || o.HonorificSuffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.HonorificSuffix.Get()
}

// GetHonorificSuffixOk returns a tuple with the HonorificSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetHonorificSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HonorificSuffix.Get(), o.HonorificSuffix.IsSet()
}

// HasHonorificSuffix returns a boolean if a field has been set.
func (o *UserProfile) HasHonorificSuffix() bool {
	if o != nil && o.HonorificSuffix.IsSet() {
		return true
	}

	return false
}

// SetHonorificSuffix gets a reference to the given NullableString and assigns it to the HonorificSuffix field.
func (o *UserProfile) SetHonorificSuffix(v string) {
	o.HonorificSuffix.Set(&v)
}
// SetHonorificSuffixNil sets the value for HonorificSuffix to be an explicit nil
func (o *UserProfile) SetHonorificSuffixNil() {
	o.HonorificSuffix.Set(nil)
}

// UnsetHonorificSuffix ensures that no value is present for HonorificSuffix, not even an explicit nil
func (o *UserProfile) UnsetHonorificSuffix() {
	o.HonorificSuffix.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetLastName() string {
	if o == nil || o.LastName.Get() == nil {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *UserProfile) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *UserProfile) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *UserProfile) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *UserProfile) UnsetLastName() {
	o.LastName.Unset()
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *UserProfile) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UserProfile) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *UserProfile) SetLocale(v string) {
	o.Locale = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *UserProfile) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserProfile) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *UserProfile) SetLogin(v string) {
	o.Login = &v
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetManager() string {
	if o == nil || o.Manager.Get() == nil {
		var ret string
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetManagerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *UserProfile) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableString and assigns it to the Manager field.
func (o *UserProfile) SetManager(v string) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *UserProfile) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *UserProfile) UnsetManager() {
	o.Manager.Unset()
}

// GetManagerId returns the ManagerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetManagerId() string {
	if o == nil || o.ManagerId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ManagerId.Get()
}

// GetManagerIdOk returns a tuple with the ManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetManagerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManagerId.Get(), o.ManagerId.IsSet()
}

// HasManagerId returns a boolean if a field has been set.
func (o *UserProfile) HasManagerId() bool {
	if o != nil && o.ManagerId.IsSet() {
		return true
	}

	return false
}

// SetManagerId gets a reference to the given NullableString and assigns it to the ManagerId field.
func (o *UserProfile) SetManagerId(v string) {
	o.ManagerId.Set(&v)
}
// SetManagerIdNil sets the value for ManagerId to be an explicit nil
func (o *UserProfile) SetManagerIdNil() {
	o.ManagerId.Set(nil)
}

// UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
func (o *UserProfile) UnsetManagerId() {
	o.ManagerId.Unset()
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetMiddleName() string {
	if o == nil || o.MiddleName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MiddleName.Get()
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetMiddleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MiddleName.Get(), o.MiddleName.IsSet()
}

// HasMiddleName returns a boolean if a field has been set.
func (o *UserProfile) HasMiddleName() bool {
	if o != nil && o.MiddleName.IsSet() {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given NullableString and assigns it to the MiddleName field.
func (o *UserProfile) SetMiddleName(v string) {
	o.MiddleName.Set(&v)
}
// SetMiddleNameNil sets the value for MiddleName to be an explicit nil
func (o *UserProfile) SetMiddleNameNil() {
	o.MiddleName.Set(nil)
}

// UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
func (o *UserProfile) UnsetMiddleName() {
	o.MiddleName.Unset()
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetMobilePhone() string {
	if o == nil || o.MobilePhone.Get() == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone.Get()
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetMobilePhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobilePhone.Get(), o.MobilePhone.IsSet()
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *UserProfile) HasMobilePhone() bool {
	if o != nil && o.MobilePhone.IsSet() {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given NullableString and assigns it to the MobilePhone field.
func (o *UserProfile) SetMobilePhone(v string) {
	o.MobilePhone.Set(&v)
}
// SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil
func (o *UserProfile) SetMobilePhoneNil() {
	o.MobilePhone.Set(nil)
}

// UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
func (o *UserProfile) UnsetMobilePhone() {
	o.MobilePhone.Unset()
}

// GetNickName returns the NickName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetNickName() string {
	if o == nil || o.NickName.Get() == nil {
		var ret string
		return ret
	}
	return *o.NickName.Get()
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetNickNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NickName.Get(), o.NickName.IsSet()
}

// HasNickName returns a boolean if a field has been set.
func (o *UserProfile) HasNickName() bool {
	if o != nil && o.NickName.IsSet() {
		return true
	}

	return false
}

// SetNickName gets a reference to the given NullableString and assigns it to the NickName field.
func (o *UserProfile) SetNickName(v string) {
	o.NickName.Set(&v)
}
// SetNickNameNil sets the value for NickName to be an explicit nil
func (o *UserProfile) SetNickNameNil() {
	o.NickName.Set(nil)
}

// UnsetNickName ensures that no value is present for NickName, not even an explicit nil
func (o *UserProfile) UnsetNickName() {
	o.NickName.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetOrganization() string {
	if o == nil || o.Organization.Get() == nil {
		var ret string
		return ret
	}
	return *o.Organization.Get()
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetOrganizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Organization.Get(), o.Organization.IsSet()
}

// HasOrganization returns a boolean if a field has been set.
func (o *UserProfile) HasOrganization() bool {
	if o != nil && o.Organization.IsSet() {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given NullableString and assigns it to the Organization field.
func (o *UserProfile) SetOrganization(v string) {
	o.Organization.Set(&v)
}
// SetOrganizationNil sets the value for Organization to be an explicit nil
func (o *UserProfile) SetOrganizationNil() {
	o.Organization.Set(nil)
}

// UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
func (o *UserProfile) UnsetOrganization() {
	o.Organization.Unset()
}

// GetPostalAddress returns the PostalAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetPostalAddress() string {
	if o == nil || o.PostalAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalAddress.Get()
}

// GetPostalAddressOk returns a tuple with the PostalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetPostalAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalAddress.Get(), o.PostalAddress.IsSet()
}

// HasPostalAddress returns a boolean if a field has been set.
func (o *UserProfile) HasPostalAddress() bool {
	if o != nil && o.PostalAddress.IsSet() {
		return true
	}

	return false
}

// SetPostalAddress gets a reference to the given NullableString and assigns it to the PostalAddress field.
func (o *UserProfile) SetPostalAddress(v string) {
	o.PostalAddress.Set(&v)
}
// SetPostalAddressNil sets the value for PostalAddress to be an explicit nil
func (o *UserProfile) SetPostalAddressNil() {
	o.PostalAddress.Set(nil)
}

// UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
func (o *UserProfile) UnsetPostalAddress() {
	o.PostalAddress.Unset()
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetPreferredLanguage() string {
	if o == nil || o.PreferredLanguage.Get() == nil {
		var ret string
		return ret
	}
	return *o.PreferredLanguage.Get()
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetPreferredLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredLanguage.Get(), o.PreferredLanguage.IsSet()
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *UserProfile) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage.IsSet() {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given NullableString and assigns it to the PreferredLanguage field.
func (o *UserProfile) SetPreferredLanguage(v string) {
	o.PreferredLanguage.Set(&v)
}
// SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil
func (o *UserProfile) SetPreferredLanguageNil() {
	o.PreferredLanguage.Set(nil)
}

// UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
func (o *UserProfile) UnsetPreferredLanguage() {
	o.PreferredLanguage.Unset()
}

// GetPrimaryPhone returns the PrimaryPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetPrimaryPhone() string {
	if o == nil || o.PrimaryPhone.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryPhone.Get()
}

// GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetPrimaryPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryPhone.Get(), o.PrimaryPhone.IsSet()
}

// HasPrimaryPhone returns a boolean if a field has been set.
func (o *UserProfile) HasPrimaryPhone() bool {
	if o != nil && o.PrimaryPhone.IsSet() {
		return true
	}

	return false
}

// SetPrimaryPhone gets a reference to the given NullableString and assigns it to the PrimaryPhone field.
func (o *UserProfile) SetPrimaryPhone(v string) {
	o.PrimaryPhone.Set(&v)
}
// SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil
func (o *UserProfile) SetPrimaryPhoneNil() {
	o.PrimaryPhone.Set(nil)
}

// UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
func (o *UserProfile) UnsetPrimaryPhone() {
	o.PrimaryPhone.Unset()
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetProfileUrl() string {
	if o == nil || o.ProfileUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProfileUrl.Get()
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetProfileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProfileUrl.Get(), o.ProfileUrl.IsSet()
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *UserProfile) HasProfileUrl() bool {
	if o != nil && o.ProfileUrl.IsSet() {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given NullableString and assigns it to the ProfileUrl field.
func (o *UserProfile) SetProfileUrl(v string) {
	o.ProfileUrl.Set(&v)
}
// SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil
func (o *UserProfile) SetProfileUrlNil() {
	o.ProfileUrl.Set(nil)
}

// UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
func (o *UserProfile) UnsetProfileUrl() {
	o.ProfileUrl.Unset()
}

// GetSecondEmail returns the SecondEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetSecondEmail() string {
	if o == nil || o.SecondEmail.Get() == nil {
		var ret string
		return ret
	}
	return *o.SecondEmail.Get()
}

// GetSecondEmailOk returns a tuple with the SecondEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetSecondEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondEmail.Get(), o.SecondEmail.IsSet()
}

// HasSecondEmail returns a boolean if a field has been set.
func (o *UserProfile) HasSecondEmail() bool {
	if o != nil && o.SecondEmail.IsSet() {
		return true
	}

	return false
}

// SetSecondEmail gets a reference to the given NullableString and assigns it to the SecondEmail field.
func (o *UserProfile) SetSecondEmail(v string) {
	o.SecondEmail.Set(&v)
}
// SetSecondEmailNil sets the value for SecondEmail to be an explicit nil
func (o *UserProfile) SetSecondEmailNil() {
	o.SecondEmail.Set(nil)
}

// UnsetSecondEmail ensures that no value is present for SecondEmail, not even an explicit nil
func (o *UserProfile) UnsetSecondEmail() {
	o.SecondEmail.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetState() string {
	if o == nil || o.State.Get() == nil {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *UserProfile) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given NullableString and assigns it to the State field.
func (o *UserProfile) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *UserProfile) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *UserProfile) UnsetState() {
	o.State.Unset()
}

// GetStreetAddress returns the StreetAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetStreetAddress() string {
	if o == nil || o.StreetAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.StreetAddress.Get()
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetStreetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StreetAddress.Get(), o.StreetAddress.IsSet()
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *UserProfile) HasStreetAddress() bool {
	if o != nil && o.StreetAddress.IsSet() {
		return true
	}

	return false
}

// SetStreetAddress gets a reference to the given NullableString and assigns it to the StreetAddress field.
func (o *UserProfile) SetStreetAddress(v string) {
	o.StreetAddress.Set(&v)
}
// SetStreetAddressNil sets the value for StreetAddress to be an explicit nil
func (o *UserProfile) SetStreetAddressNil() {
	o.StreetAddress.Set(nil)
}

// UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
func (o *UserProfile) UnsetStreetAddress() {
	o.StreetAddress.Unset()
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetTimezone() string {
	if o == nil || o.Timezone.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserProfile) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableString and assigns it to the Timezone field.
func (o *UserProfile) SetTimezone(v string) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *UserProfile) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *UserProfile) UnsetTimezone() {
	o.Timezone.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *UserProfile) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *UserProfile) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *UserProfile) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *UserProfile) UnsetTitle() {
	o.Title.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetUserType() string {
	if o == nil || o.UserType.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserType.Get()
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetUserTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserType.Get(), o.UserType.IsSet()
}

// HasUserType returns a boolean if a field has been set.
func (o *UserProfile) HasUserType() bool {
	if o != nil && o.UserType.IsSet() {
		return true
	}

	return false
}

// SetUserType gets a reference to the given NullableString and assigns it to the UserType field.
func (o *UserProfile) SetUserType(v string) {
	o.UserType.Set(&v)
}
// SetUserTypeNil sets the value for UserType to be an explicit nil
func (o *UserProfile) SetUserTypeNil() {
	o.UserType.Set(nil)
}

// UnsetUserType ensures that no value is present for UserType, not even an explicit nil
func (o *UserProfile) UnsetUserType() {
	o.UserType.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserProfile) GetZipCode() string {
	if o == nil || o.ZipCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserProfile) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *UserProfile) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *UserProfile) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *UserProfile) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *UserProfile) UnsetZipCode() {
	o.ZipCode.Unset()
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.CostCenter.IsSet() {
		toSerialize["costCenter"] = o.CostCenter.Get()
	}
	if o.CountryCode.IsSet() {
		toSerialize["countryCode"] = o.CountryCode.Get()
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Division.IsSet() {
		toSerialize["division"] = o.Division.Get()
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EmployeeNumber != nil {
		toSerialize["employeeNumber"] = o.EmployeeNumber
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.HonorificPrefix.IsSet() {
		toSerialize["honorificPrefix"] = o.HonorificPrefix.Get()
	}
	if o.HonorificSuffix.IsSet() {
		toSerialize["honorificSuffix"] = o.HonorificSuffix.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if o.ManagerId.IsSet() {
		toSerialize["managerId"] = o.ManagerId.Get()
	}
	if o.MiddleName.IsSet() {
		toSerialize["middleName"] = o.MiddleName.Get()
	}
	if o.MobilePhone.IsSet() {
		toSerialize["mobilePhone"] = o.MobilePhone.Get()
	}
	if o.NickName.IsSet() {
		toSerialize["nickName"] = o.NickName.Get()
	}
	if o.Organization.IsSet() {
		toSerialize["organization"] = o.Organization.Get()
	}
	if o.PostalAddress.IsSet() {
		toSerialize["postalAddress"] = o.PostalAddress.Get()
	}
	if o.PreferredLanguage.IsSet() {
		toSerialize["preferredLanguage"] = o.PreferredLanguage.Get()
	}
	if o.PrimaryPhone.IsSet() {
		toSerialize["primaryPhone"] = o.PrimaryPhone.Get()
	}
	if o.ProfileUrl.IsSet() {
		toSerialize["profileUrl"] = o.ProfileUrl.Get()
	}
	if o.SecondEmail.IsSet() {
		toSerialize["secondEmail"] = o.SecondEmail.Get()
	}
	if o.State.IsSet() {
		toSerialize["state"] = o.State.Get()
	}
	if o.StreetAddress.IsSet() {
		toSerialize["streetAddress"] = o.StreetAddress.Get()
	}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.UserType.IsSet() {
		toSerialize["userType"] = o.UserType.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["zipCode"] = o.ZipCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UserProfile) UnmarshalJSON(bytes []byte) (err error) {
	varUserProfile := _UserProfile{}

	err = json.Unmarshal(bytes, &varUserProfile)
	if err == nil {
		*o = UserProfile(varUserProfile)
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

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

