# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** | The city or locality of the user&#39;s address (&#x60;locality&#x60;) | [optional] 
**CostCenter** | Pointer to **NullableString** | Name of the cost center assigned to a user | [optional] 
**CountryCode** | Pointer to **NullableString** | The country name component of the user&#39;s address (&#x60;country&#x60;) | [optional] 
**Department** | Pointer to **string** | Name of the user&#39;s department | [optional] 
**DisplayName** | Pointer to **NullableString** | Name of the user suitable for display to end users | [optional] 
**Division** | Pointer to **NullableString** | Name of the user&#39;s division | [optional] 
**Email** | Pointer to **string** | The primary email address of the user. For validation, see [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3). | [optional] 
**EmployeeNumber** | Pointer to **string** | The organization or company assigned unique identifier for the user | [optional] 
**FirstName** | Pointer to **NullableString** | Given name of the user (&#x60;givenName&#x60;) | [optional] 
**HonorificPrefix** | Pointer to **NullableString** | Honorific prefix(es) of the user, or title in most Western languages | [optional] 
**HonorificSuffix** | Pointer to **NullableString** | Honorific suffix(es) of the user | [optional] 
**LastName** | Pointer to **NullableString** | The family name of the user (&#x60;familyName&#x60;) | [optional] 
**Locale** | Pointer to **string** | The user&#39;s default location for purposes of localizing items such as currency, date time format, numerical representations, and so on. A locale value is a concatenation of the ISO 639-1 two-letter language code, an underscore, and the ISO 3166-1 two-letter country code. For example, en_US specifies the language English and country US. This value is &#x60;en_US&#x60; by default. | [optional] 
**Login** | Pointer to **string** | The unique identifier for the user (&#x60;username&#x60;). For validation, see [Login pattern validation](https://developer.okta.com/docs/reference/api/schemas/#login-pattern-validation). See also [Okta login](https://developer.okta.com/docs/reference/api/users/#okta-login). | [optional] 
**Manager** | Pointer to **NullableString** | The &#x60;displayName&#x60; of the user&#39;s manager | [optional] 
**ManagerId** | Pointer to **NullableString** | The &#x60;id&#x60; of the user&#39;s manager | [optional] 
**MiddleName** | Pointer to **NullableString** | The middle name of the user | [optional] 
**MobilePhone** | Pointer to **NullableString** | The mobile phone number of the user | [optional] 
**NickName** | Pointer to **NullableString** | The casual way to address the user in real life | [optional] 
**Organization** | Pointer to **NullableString** | Name of the the user&#39;s organization | [optional] 
**PostalAddress** | Pointer to **NullableString** | Mailing address component of the user&#39;s address | [optional] 
**PreferredLanguage** | Pointer to **NullableString** | The user&#39;s preferred written or spoken language | [optional] 
**PrimaryPhone** | Pointer to **NullableString** | The primary phone number of the user such as a home number | [optional] 
**ProfileUrl** | Pointer to **NullableString** | The URL of the user&#39;s online profile. For example, a web page. See [URL](https://datatracker.ietf.org/doc/html/rfc1808). | [optional] 
**SecondEmail** | Pointer to **NullableString** | The secondary email address of the user typically used for account recovery | [optional] 
**State** | Pointer to **NullableString** | The state or region component of the user&#39;s address (&#x60;region&#x60;) | [optional] 
**StreetAddress** | Pointer to **NullableString** | The full street address component of the user&#39;s address | [optional] 
**Timezone** | Pointer to **NullableString** | The user&#39;s time zone | [optional] 
**Title** | Pointer to **NullableString** | The user&#39;s title, such as Vice President | [optional] 
**UserType** | Pointer to **NullableString** | The property used to describe the organization-to-user relationship, such as employee or contractor | [optional] 
**ZipCode** | Pointer to **NullableString** | The ZIP code or postal code component of the user&#39;s address (&#x60;postalCode&#x60;) | [optional] 

## Methods

### NewUserProfile

`func NewUserProfile() *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *UserProfile) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserProfile) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserProfile) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UserProfile) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *UserProfile) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *UserProfile) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCostCenter

`func (o *UserProfile) GetCostCenter() string`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *UserProfile) GetCostCenterOk() (*string, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *UserProfile) SetCostCenter(v string)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *UserProfile) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### SetCostCenterNil

`func (o *UserProfile) SetCostCenterNil(b bool)`

 SetCostCenterNil sets the value for CostCenter to be an explicit nil

### UnsetCostCenter
`func (o *UserProfile) UnsetCostCenter()`

UnsetCostCenter ensures that no value is present for CostCenter, not even an explicit nil
### GetCountryCode

`func (o *UserProfile) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserProfile) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserProfile) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UserProfile) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *UserProfile) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *UserProfile) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetDepartment

`func (o *UserProfile) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UserProfile) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UserProfile) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UserProfile) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserProfile) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserProfile) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserProfile) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserProfile) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *UserProfile) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *UserProfile) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDivision

`func (o *UserProfile) GetDivision() string`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *UserProfile) GetDivisionOk() (*string, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *UserProfile) SetDivision(v string)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *UserProfile) HasDivision() bool`

HasDivision returns a boolean if a field has been set.

### SetDivisionNil

`func (o *UserProfile) SetDivisionNil(b bool)`

 SetDivisionNil sets the value for Division to be an explicit nil

### UnsetDivision
`func (o *UserProfile) UnsetDivision()`

UnsetDivision ensures that no value is present for Division, not even an explicit nil
### GetEmail

`func (o *UserProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *UserProfile) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UserProfile) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UserProfile) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *UserProfile) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *UserProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserProfile) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserProfile) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserProfile) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetHonorificPrefix

`func (o *UserProfile) GetHonorificPrefix() string`

GetHonorificPrefix returns the HonorificPrefix field if non-nil, zero value otherwise.

### GetHonorificPrefixOk

`func (o *UserProfile) GetHonorificPrefixOk() (*string, bool)`

GetHonorificPrefixOk returns a tuple with the HonorificPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorificPrefix

`func (o *UserProfile) SetHonorificPrefix(v string)`

SetHonorificPrefix sets HonorificPrefix field to given value.

### HasHonorificPrefix

`func (o *UserProfile) HasHonorificPrefix() bool`

HasHonorificPrefix returns a boolean if a field has been set.

### SetHonorificPrefixNil

`func (o *UserProfile) SetHonorificPrefixNil(b bool)`

 SetHonorificPrefixNil sets the value for HonorificPrefix to be an explicit nil

### UnsetHonorificPrefix
`func (o *UserProfile) UnsetHonorificPrefix()`

UnsetHonorificPrefix ensures that no value is present for HonorificPrefix, not even an explicit nil
### GetHonorificSuffix

`func (o *UserProfile) GetHonorificSuffix() string`

GetHonorificSuffix returns the HonorificSuffix field if non-nil, zero value otherwise.

### GetHonorificSuffixOk

`func (o *UserProfile) GetHonorificSuffixOk() (*string, bool)`

GetHonorificSuffixOk returns a tuple with the HonorificSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorificSuffix

`func (o *UserProfile) SetHonorificSuffix(v string)`

SetHonorificSuffix sets HonorificSuffix field to given value.

### HasHonorificSuffix

`func (o *UserProfile) HasHonorificSuffix() bool`

HasHonorificSuffix returns a boolean if a field has been set.

### SetHonorificSuffixNil

`func (o *UserProfile) SetHonorificSuffixNil(b bool)`

 SetHonorificSuffixNil sets the value for HonorificSuffix to be an explicit nil

### UnsetHonorificSuffix
`func (o *UserProfile) UnsetHonorificSuffix()`

UnsetHonorificSuffix ensures that no value is present for HonorificSuffix, not even an explicit nil
### GetLastName

`func (o *UserProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserProfile) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserProfile) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserProfile) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLocale

`func (o *UserProfile) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserProfile) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserProfile) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserProfile) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLogin

`func (o *UserProfile) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserProfile) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserProfile) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserProfile) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetManager

`func (o *UserProfile) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *UserProfile) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *UserProfile) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *UserProfile) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *UserProfile) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *UserProfile) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetManagerId

`func (o *UserProfile) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *UserProfile) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *UserProfile) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *UserProfile) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### SetManagerIdNil

`func (o *UserProfile) SetManagerIdNil(b bool)`

 SetManagerIdNil sets the value for ManagerId to be an explicit nil

### UnsetManagerId
`func (o *UserProfile) UnsetManagerId()`

UnsetManagerId ensures that no value is present for ManagerId, not even an explicit nil
### GetMiddleName

`func (o *UserProfile) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UserProfile) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UserProfile) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UserProfile) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *UserProfile) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *UserProfile) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetMobilePhone

`func (o *UserProfile) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *UserProfile) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *UserProfile) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *UserProfile) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *UserProfile) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *UserProfile) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil
### GetNickName

`func (o *UserProfile) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *UserProfile) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *UserProfile) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *UserProfile) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### SetNickNameNil

`func (o *UserProfile) SetNickNameNil(b bool)`

 SetNickNameNil sets the value for NickName to be an explicit nil

### UnsetNickName
`func (o *UserProfile) UnsetNickName()`

UnsetNickName ensures that no value is present for NickName, not even an explicit nil
### GetOrganization

`func (o *UserProfile) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UserProfile) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UserProfile) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *UserProfile) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *UserProfile) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *UserProfile) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetPostalAddress

`func (o *UserProfile) GetPostalAddress() string`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *UserProfile) GetPostalAddressOk() (*string, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *UserProfile) SetPostalAddress(v string)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *UserProfile) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### SetPostalAddressNil

`func (o *UserProfile) SetPostalAddressNil(b bool)`

 SetPostalAddressNil sets the value for PostalAddress to be an explicit nil

### UnsetPostalAddress
`func (o *UserProfile) UnsetPostalAddress()`

UnsetPostalAddress ensures that no value is present for PostalAddress, not even an explicit nil
### GetPreferredLanguage

`func (o *UserProfile) GetPreferredLanguage() string`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *UserProfile) GetPreferredLanguageOk() (*string, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *UserProfile) SetPreferredLanguage(v string)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *UserProfile) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### SetPreferredLanguageNil

`func (o *UserProfile) SetPreferredLanguageNil(b bool)`

 SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil

### UnsetPreferredLanguage
`func (o *UserProfile) UnsetPreferredLanguage()`

UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
### GetPrimaryPhone

`func (o *UserProfile) GetPrimaryPhone() string`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *UserProfile) GetPrimaryPhoneOk() (*string, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *UserProfile) SetPrimaryPhone(v string)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *UserProfile) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### SetPrimaryPhoneNil

`func (o *UserProfile) SetPrimaryPhoneNil(b bool)`

 SetPrimaryPhoneNil sets the value for PrimaryPhone to be an explicit nil

### UnsetPrimaryPhone
`func (o *UserProfile) UnsetPrimaryPhone()`

UnsetPrimaryPhone ensures that no value is present for PrimaryPhone, not even an explicit nil
### GetProfileUrl

`func (o *UserProfile) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *UserProfile) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *UserProfile) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *UserProfile) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### SetProfileUrlNil

`func (o *UserProfile) SetProfileUrlNil(b bool)`

 SetProfileUrlNil sets the value for ProfileUrl to be an explicit nil

### UnsetProfileUrl
`func (o *UserProfile) UnsetProfileUrl()`

UnsetProfileUrl ensures that no value is present for ProfileUrl, not even an explicit nil
### GetSecondEmail

`func (o *UserProfile) GetSecondEmail() string`

GetSecondEmail returns the SecondEmail field if non-nil, zero value otherwise.

### GetSecondEmailOk

`func (o *UserProfile) GetSecondEmailOk() (*string, bool)`

GetSecondEmailOk returns a tuple with the SecondEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondEmail

`func (o *UserProfile) SetSecondEmail(v string)`

SetSecondEmail sets SecondEmail field to given value.

### HasSecondEmail

`func (o *UserProfile) HasSecondEmail() bool`

HasSecondEmail returns a boolean if a field has been set.

### SetSecondEmailNil

`func (o *UserProfile) SetSecondEmailNil(b bool)`

 SetSecondEmailNil sets the value for SecondEmail to be an explicit nil

### UnsetSecondEmail
`func (o *UserProfile) UnsetSecondEmail()`

UnsetSecondEmail ensures that no value is present for SecondEmail, not even an explicit nil
### GetState

`func (o *UserProfile) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserProfile) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserProfile) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserProfile) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *UserProfile) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *UserProfile) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStreetAddress

`func (o *UserProfile) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserProfile) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserProfile) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UserProfile) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### SetStreetAddressNil

`func (o *UserProfile) SetStreetAddressNil(b bool)`

 SetStreetAddressNil sets the value for StreetAddress to be an explicit nil

### UnsetStreetAddress
`func (o *UserProfile) UnsetStreetAddress()`

UnsetStreetAddress ensures that no value is present for StreetAddress, not even an explicit nil
### GetTimezone

`func (o *UserProfile) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserProfile) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserProfile) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserProfile) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *UserProfile) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *UserProfile) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetTitle

`func (o *UserProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserProfile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserProfile) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserProfile) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UserProfile) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UserProfile) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserType

`func (o *UserProfile) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserProfile) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserProfile) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserProfile) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *UserProfile) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *UserProfile) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetZipCode

`func (o *UserProfile) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UserProfile) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UserProfile) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UserProfile) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.

### SetZipCodeNil

`func (o *UserProfile) SetZipCodeNil(b bool)`

 SetZipCodeNil sets the value for ZipCode to be an explicit nil

### UnsetZipCode
`func (o *UserProfile) UnsetZipCode()`

UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


