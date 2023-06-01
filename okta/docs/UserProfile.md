# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** |  | [optional] 
**CostCenter** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Division** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmployeeNumber** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**HonorificPrefix** | Pointer to **string** |  | [optional] 
**HonorificSuffix** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to **string** | The language specified as an [IETF BCP 47 language tag](https://datatracker.ietf.org/doc/html/rfc5646). | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **string** |  | [optional] 
**ManagerId** | Pointer to **string** |  | [optional] 
**MiddleName** | Pointer to **string** |  | [optional] 
**MobilePhone** | Pointer to **NullableString** |  | [optional] 
**NickName** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**PostalAddress** | Pointer to **NullableString** |  | [optional] 
**PreferredLanguage** | Pointer to **string** |  | [optional] 
**PrimaryPhone** | Pointer to **NullableString** |  | [optional] 
**ProfileUrl** | Pointer to **string** |  | [optional] 
**SecondEmail** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**StreetAddress** | Pointer to **NullableString** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UserType** | Pointer to **string** |  | [optional] 
**ZipCode** | Pointer to **NullableString** |  | [optional] 

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


