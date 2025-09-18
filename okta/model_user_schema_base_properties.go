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

API version: 5.1.0
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
)

// checks if the UserSchemaBaseProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSchemaBaseProperties{}

// UserSchemaBaseProperties struct for UserSchemaBaseProperties
type UserSchemaBaseProperties struct {
	// City or locality component of the user's address (`locality`)
	City *UserSchemaAttribute `json:"city,omitempty"`
	// Name of a cost center assigned to the user
	CostCenter *UserSchemaAttribute `json:"costCenter,omitempty"`
	// Country name component of the user's address (`country`.) This property uses [ISO 3166-1 alpha 2 \"short\" code format](https://tools.ietf.org/html/draft-ietf-scim-core-schema-22#ref-ISO3166).
	CountryCode *UserSchemaAttribute `json:"countryCode,omitempty"`
	// Name of the user's department
	Department *UserSchemaAttribute `json:"department,omitempty"`
	// Name of the user, suitable for display to end users
	DisplayName *UserSchemaAttribute `json:"displayName,omitempty"`
	// Name of the user's division
	Division *UserSchemaAttribute `json:"division,omitempty"`
	// Primary email address of the user. This property is formatted according to [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3).
	Email *UserSchemaAttribute `json:"email,omitempty"`
	// Organization or company assigned unique identifier for the user
	EmployeeNumber *UserSchemaAttribute `json:"employeeNumber,omitempty"`
	// Given name of the user (`givenName`)
	FirstName *UserSchemaAttribute `json:"firstName,omitempty"`
	// Honorific prefix(es) of the user or title in most Western languages
	HonorificPrefix *UserSchemaAttribute `json:"honorificPrefix,omitempty"`
	// Honorific suffix(es) of the user
	HonorificSuffix *UserSchemaAttribute `json:"honorificSuffix,omitempty"`
	// Family name of the user (`familyName`)
	LastName *UserSchemaAttribute `json:"lastName,omitempty"`
	// User's default location for purposes of localizing items such as currency, date time format, numerical representations, and so on.  A locale value is a concatenation of the ISO 639-1 two-letter language code, an underscore, and the ISO 3166-1 two-letter country code. For example: `en_US` specifies the language English and country US. This value is `en_US` by default.
	Locale *UserSchemaAttribute `json:"locale,omitempty"`
	// Unique identifier for the user (`userName`)  The login property is validated according to its pattern attribute, which is a string. By default, the attribute is null. When the attribute is null, the username is required to be formatted as an email address as defined by [RFC 6531 Section 3.3](http://tools.ietf.org/html/rfc6531#section-3.3). The pattern can be set through the API to one of the following forms. (The Admin Console provides access to the same forms.)   * A login pattern of `\".+\"` indicates that there is no restriction on usernames. Any non-empty, unique value is permitted, and the minimum length of five isn't enforced. In this case, usernames don't need to include the `@` character. If a name does include `@`, the portion ahead of the `@` can be used for logging in, provided it identifies a unique user within the org.   * A login pattern of the form `\"[...]+\"` indicates that usernames must only contain characters from the set given between the brackets. The enclosing brackets and final `+` are required for this form. Character ranges can be indicated using hyphens. To include the hyphen itself in the allowed set, the hyphen must appear first. Any characters in the set except the hyphen, a-z, A-Z, and 0-9 must be preceded by a backslash (`\\`). For example, `\"[a-z13579\\.]+\"` would restrict usernames to lowercase letters, odd digits, and periods, while `\"[-a-zA-Z0-9]+\"` would allow basic alphanumeric characters and hyphens.
	Login *UserSchemaAttribute `json:"login,omitempty"`
	// The `displayName` of the user's manager
	Manager *UserSchemaAttribute `json:"manager,omitempty"`
	// The `id` of the user's manager
	ManagerId *UserSchemaAttribute `json:"managerId,omitempty"`
	// Middle name(s) of the user
	MiddleName *UserSchemaAttribute `json:"middleName,omitempty"`
	// Mobile phone number of the user
	MobilePhone *UserSchemaAttribute `json:"mobilePhone,omitempty"`
	// Casual way to address the user in real life
	NickName *UserSchemaAttribute `json:"nickName,omitempty"`
	// Name of the user's organization
	Organization *UserSchemaAttribute `json:"organization,omitempty"`
	// Mailing address component of the user's address
	PostalAddress *UserSchemaAttribute `json:"postalAddress,omitempty"`
	// User's preferred written or spoken languages. This property is formatted according to [RFC 7231 Section 5.3.5](https://tools.ietf.org/html/rfc7231#section-5.3.5).
	PreferredLanguage *UserSchemaAttribute `json:"preferredLanguage,omitempty"`
	// Primary phone number of the user, such as home number
	PrimaryPhone *UserSchemaAttribute `json:"primaryPhone,omitempty"`
	// URL of the user's online profile (for example, a web page.) This property is formatted according to the [Relative Uniform Resource Locators specification](https://tools.ietf.org/html/draft-ietf-scim-core-schema-22#ref-ISO3166).
	ProfileUrl *UserSchemaAttribute `json:"profileUrl,omitempty"`
	// Secondary email address of the user typically used for account recovery. This property is formatted according to [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3).
	SecondEmail *UserSchemaAttribute `json:"secondEmail,omitempty"`
	// State or region component of the user's address (`region`)
	State *UserSchemaAttribute `json:"state,omitempty"`
	// Full street address component of the user's address
	StreetAddress *UserSchemaAttribute `json:"streetAddress,omitempty"`
	// User's time zone. This property is formatted according to the [IANA Time Zone database format](https://tools.ietf.org/html/rfc6557).
	Timezone *UserSchemaAttribute `json:"timezone,omitempty"`
	// User's title, such as \"Vice President\"
	Title *UserSchemaAttribute `json:"title,omitempty"`
	// Used to describe the organization to the user relationship such as \"Employee\" or \"Contractor\".  **Note:** The `userType` field is an arbitrary string value and isn't related to the newer [User Types](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/) feature.
	UserType *UserSchemaAttribute `json:"userType,omitempty"`
	// ZIP code or postal code component of the user's address (`postalCode`)
	ZipCode              *UserSchemaAttribute `json:"zipCode,omitempty"`
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
	if o == nil || IsNil(o.City) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetCityOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasCity() bool {
	if o != nil && !IsNil(o.City) {
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
	if o == nil || IsNil(o.CostCenter) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.CostCenter
}

// GetCostCenterOk returns a tuple with the CostCenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetCostCenterOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.CostCenter) {
		return nil, false
	}
	return o.CostCenter, true
}

// HasCostCenter returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasCostCenter() bool {
	if o != nil && !IsNil(o.CostCenter) {
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
	if o == nil || IsNil(o.CountryCode) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetCountryCodeOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
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
	if o == nil || IsNil(o.Department) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetDepartmentOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
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
	if o == nil || IsNil(o.DisplayName) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetDisplayNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
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
	if o == nil || IsNil(o.Division) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Division
}

// GetDivisionOk returns a tuple with the Division field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetDivisionOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Division) {
		return nil, false
	}
	return o.Division, true
}

// HasDivision returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasDivision() bool {
	if o != nil && !IsNil(o.Division) {
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
	if o == nil || IsNil(o.Email) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetEmailOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
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
	if o == nil || IsNil(o.EmployeeNumber) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.EmployeeNumber
}

// GetEmployeeNumberOk returns a tuple with the EmployeeNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetEmployeeNumberOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.EmployeeNumber) {
		return nil, false
	}
	return o.EmployeeNumber, true
}

// HasEmployeeNumber returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasEmployeeNumber() bool {
	if o != nil && !IsNil(o.EmployeeNumber) {
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
	if o == nil || IsNil(o.FirstName) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetFirstNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
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
	if o == nil || IsNil(o.HonorificPrefix) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.HonorificPrefix
}

// GetHonorificPrefixOk returns a tuple with the HonorificPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetHonorificPrefixOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.HonorificPrefix) {
		return nil, false
	}
	return o.HonorificPrefix, true
}

// HasHonorificPrefix returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasHonorificPrefix() bool {
	if o != nil && !IsNil(o.HonorificPrefix) {
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
	if o == nil || IsNil(o.HonorificSuffix) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.HonorificSuffix
}

// GetHonorificSuffixOk returns a tuple with the HonorificSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetHonorificSuffixOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.HonorificSuffix) {
		return nil, false
	}
	return o.HonorificSuffix, true
}

// HasHonorificSuffix returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasHonorificSuffix() bool {
	if o != nil && !IsNil(o.HonorificSuffix) {
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
	if o == nil || IsNil(o.LastName) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetLastNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
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
	if o == nil || IsNil(o.Locale) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetLocaleOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
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
	if o == nil || IsNil(o.Login) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetLoginOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Login) {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasLogin() bool {
	if o != nil && !IsNil(o.Login) {
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
	if o == nil || IsNil(o.Manager) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Manager
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetManagerOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Manager) {
		return nil, false
	}
	return o.Manager, true
}

// HasManager returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasManager() bool {
	if o != nil && !IsNil(o.Manager) {
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
	if o == nil || IsNil(o.ManagerId) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.ManagerId
}

// GetManagerIdOk returns a tuple with the ManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetManagerIdOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.ManagerId) {
		return nil, false
	}
	return o.ManagerId, true
}

// HasManagerId returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasManagerId() bool {
	if o != nil && !IsNil(o.ManagerId) {
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
	if o == nil || IsNil(o.MiddleName) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetMiddleNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
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
	if o == nil || IsNil(o.MobilePhone) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetMobilePhoneOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.MobilePhone) {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasMobilePhone() bool {
	if o != nil && !IsNil(o.MobilePhone) {
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
	if o == nil || IsNil(o.NickName) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetNickNameOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
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
	if o == nil || IsNil(o.Organization) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetOrganizationOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
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
	if o == nil || IsNil(o.PostalAddress) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.PostalAddress
}

// GetPostalAddressOk returns a tuple with the PostalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetPostalAddressOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.PostalAddress) {
		return nil, false
	}
	return o.PostalAddress, true
}

// HasPostalAddress returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasPostalAddress() bool {
	if o != nil && !IsNil(o.PostalAddress) {
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
	if o == nil || IsNil(o.PreferredLanguage) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetPreferredLanguageOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.PreferredLanguage) {
		return nil, false
	}
	return o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasPreferredLanguage() bool {
	if o != nil && !IsNil(o.PreferredLanguage) {
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
	if o == nil || IsNil(o.PrimaryPhone) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.PrimaryPhone
}

// GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetPrimaryPhoneOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.PrimaryPhone) {
		return nil, false
	}
	return o.PrimaryPhone, true
}

// HasPrimaryPhone returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasPrimaryPhone() bool {
	if o != nil && !IsNil(o.PrimaryPhone) {
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
	if o == nil || IsNil(o.ProfileUrl) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.ProfileUrl
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetProfileUrlOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.ProfileUrl) {
		return nil, false
	}
	return o.ProfileUrl, true
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasProfileUrl() bool {
	if o != nil && !IsNil(o.ProfileUrl) {
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
	if o == nil || IsNil(o.SecondEmail) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.SecondEmail
}

// GetSecondEmailOk returns a tuple with the SecondEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetSecondEmailOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.SecondEmail) {
		return nil, false
	}
	return o.SecondEmail, true
}

// HasSecondEmail returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasSecondEmail() bool {
	if o != nil && !IsNil(o.SecondEmail) {
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
	if o == nil || IsNil(o.State) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetStateOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasState() bool {
	if o != nil && !IsNil(o.State) {
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
	if o == nil || IsNil(o.StreetAddress) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.StreetAddress
}

// GetStreetAddressOk returns a tuple with the StreetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetStreetAddressOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.StreetAddress) {
		return nil, false
	}
	return o.StreetAddress, true
}

// HasStreetAddress returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasStreetAddress() bool {
	if o != nil && !IsNil(o.StreetAddress) {
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
	if o == nil || IsNil(o.Timezone) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetTimezoneOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
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
	if o == nil || IsNil(o.Title) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetTitleOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
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
	if o == nil || IsNil(o.UserType) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetUserTypeOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
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
	if o == nil || IsNil(o.ZipCode) {
		var ret UserSchemaAttribute
		return ret
	}
	return *o.ZipCode
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSchemaBaseProperties) GetZipCodeOk() (*UserSchemaAttribute, bool) {
	if o == nil || IsNil(o.ZipCode) {
		return nil, false
	}
	return o.ZipCode, true
}

// HasZipCode returns a boolean if a field has been set.
func (o *UserSchemaBaseProperties) HasZipCode() bool {
	if o != nil && !IsNil(o.ZipCode) {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given UserSchemaAttribute and assigns it to the ZipCode field.
func (o *UserSchemaBaseProperties) SetZipCode(v UserSchemaAttribute) {
	o.ZipCode = &v
}

func (o UserSchemaBaseProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSchemaBaseProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.CostCenter) {
		toSerialize["costCenter"] = o.CostCenter
	}
	if !IsNil(o.CountryCode) {
		toSerialize["countryCode"] = o.CountryCode
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Division) {
		toSerialize["division"] = o.Division
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmployeeNumber) {
		toSerialize["employeeNumber"] = o.EmployeeNumber
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.HonorificPrefix) {
		toSerialize["honorificPrefix"] = o.HonorificPrefix
	}
	if !IsNil(o.HonorificSuffix) {
		toSerialize["honorificSuffix"] = o.HonorificSuffix
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Login) {
		toSerialize["login"] = o.Login
	}
	if !IsNil(o.Manager) {
		toSerialize["manager"] = o.Manager
	}
	if !IsNil(o.ManagerId) {
		toSerialize["managerId"] = o.ManagerId
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.MobilePhone) {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.PostalAddress) {
		toSerialize["postalAddress"] = o.PostalAddress
	}
	if !IsNil(o.PreferredLanguage) {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if !IsNil(o.PrimaryPhone) {
		toSerialize["primaryPhone"] = o.PrimaryPhone
	}
	if !IsNil(o.ProfileUrl) {
		toSerialize["profileUrl"] = o.ProfileUrl
	}
	if !IsNil(o.SecondEmail) {
		toSerialize["secondEmail"] = o.SecondEmail
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.StreetAddress) {
		toSerialize["streetAddress"] = o.StreetAddress
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	if !IsNil(o.ZipCode) {
		toSerialize["zipCode"] = o.ZipCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UserSchemaBaseProperties) UnmarshalJSON(data []byte) (err error) {
	varUserSchemaBaseProperties := _UserSchemaBaseProperties{}

	err = json.Unmarshal(data, &varUserSchemaBaseProperties)

	if err != nil {
		return err
	}

	*o = UserSchemaBaseProperties(varUserSchemaBaseProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
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
