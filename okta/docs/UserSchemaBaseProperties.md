# UserSchemaBaseProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | City or locality component of the user&#39;s address (&#x60;locality&#x60;) | [optional] 
**CostCenter** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Name of a cost center assigned to the user | [optional] 
**CountryCode** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Country name component of the user&#39;s address (&#x60;country&#x60;.) This property uses [ISO 3166-1 alpha 2 \&quot;short\&quot; code format](https://tools.ietf.org/html/draft-ietf-scim-core-schema-22#ref-ISO3166). | [optional] 
**Department** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Name of the user&#39;s department | [optional] 
**DisplayName** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Name of the user, suitable for display to end users | [optional] 
**Division** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Name of the user&#39;s division | [optional] 
**Email** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Primary email address of the user. This property is formatted according to [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3). | [optional] 
**EmployeeNumber** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Organization or company assigned unique identifier for the user | [optional] 
**FirstName** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Given name of the user (&#x60;givenName&#x60;) | [optional] 
**HonorificPrefix** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Honorific prefix(es) of the user or title in most Western languages | [optional] 
**HonorificSuffix** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Honorific suffix(es) of the user | [optional] 
**LastName** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Family name of the user (&#x60;familyName&#x60;) | [optional] 
**Locale** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | User&#39;s default location for purposes of localizing items such as currency, date time format, numerical representations, and so on.  A locale value is a concatenation of the ISO 639-1 two-letter language code, an underscore, and the ISO 3166-1 two-letter country code. For example: &#x60;en_US&#x60; specifies the language English and country US. This value is &#x60;en_US&#x60; by default. | [optional] 
**Login** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Unique identifier for the user (&#x60;userName&#x60;)  The login property is validated according to its pattern attribute, which is a string. By default, the attribute is null. When the attribute is null, the username is required to be formatted as an email address as defined by [RFC 6531 Section 3.3](http://tools.ietf.org/html/rfc6531#section-3.3). The pattern can be set through the API to one of the following forms. (The Admin Console provides access to the same forms.)   * A login pattern of &#x60;\&quot;.+\&quot;&#x60; indicates that there is no restriction on usernames. Any non-empty, unique value is permitted, and the minimum length of five isn&#39;t enforced. In this case, usernames don&#39;t need to include the &#x60;@&#x60; character. If a name does include &#x60;@&#x60;, the portion ahead of the &#x60;@&#x60; can be used for logging in, provided it identifies a unique user within the org.   * A login pattern of the form &#x60;\&quot;[...]+\&quot;&#x60; indicates that usernames must only contain characters from the set given between the brackets. The enclosing brackets and final &#x60;+&#x60; are required for this form. Character ranges can be indicated using hyphens. To include the hyphen itself in the allowed set, the hyphen must appear first. Any characters in the set except the hyphen, a-z, A-Z, and 0-9 must be preceded by a backslash (&#x60;\\&#x60;). For example, &#x60;\&quot;[a-z13579\\.]+\&quot;&#x60; would restrict usernames to lowercase letters, odd digits, and periods, while &#x60;\&quot;[-a-zA-Z0-9]+\&quot;&#x60; would allow basic alphanumeric characters and hyphens. | [optional] 
**Manager** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | The &#x60;displayName&#x60; of the user&#39;s manager | [optional] 
**ManagerId** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | The &#x60;id&#x60; of the user&#39;s manager | [optional] 
**MiddleName** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Middle name(s) of the user | [optional] 
**MobilePhone** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Mobile phone number of the user | [optional] 
**NickName** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Casual way to address the user in real life | [optional] 
**Organization** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Name of the user&#39;s organization | [optional] 
**PostalAddress** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Mailing address component of the user&#39;s address | [optional] 
**PreferredLanguage** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | User&#39;s preferred written or spoken languages. This property is formatted according to [RFC 7231 Section 5.3.5](https://tools.ietf.org/html/rfc7231#section-5.3.5). | [optional] 
**PrimaryPhone** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Primary phone number of the user, such as home number | [optional] 
**ProfileUrl** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | URL of the user&#39;s online profile (for example, a web page.) This property is formatted according to the [Relative Uniform Resource Locators specification](https://tools.ietf.org/html/draft-ietf-scim-core-schema-22#ref-ISO3166). | [optional] 
**SecondEmail** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Secondary email address of the user typically used for account recovery. This property is formatted according to [RFC 5322 Section 3.2.3](https://datatracker.ietf.org/doc/html/rfc5322#section-3.2.3). | [optional] 
**State** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | State or region component of the user&#39;s address (&#x60;region&#x60;) | [optional] 
**StreetAddress** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Full street address component of the user&#39;s address | [optional] 
**Timezone** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | User&#39;s time zone. This property is formatted according to the [IANA Time Zone database format](https://tools.ietf.org/html/rfc6557). | [optional] 
**Title** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | User&#39;s title, such as \&quot;Vice President\&quot; | [optional] 
**UserType** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | Used to describe the organization to the user relationship such as \&quot;Employee\&quot; or \&quot;Contractor\&quot;.  **Note:** The &#x60;userType&#x60; field is an arbitrary string value and isn&#39;t related to the newer [User Types](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/UserType/) feature. | [optional] 
**ZipCode** | Pointer to [**UserSchemaAttribute**](UserSchemaAttribute.md) | ZIP code or postal code component of the user&#39;s address (&#x60;postalCode&#x60;) | [optional] 

## Methods

### NewUserSchemaBaseProperties

`func NewUserSchemaBaseProperties() *UserSchemaBaseProperties`

NewUserSchemaBaseProperties instantiates a new UserSchemaBaseProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSchemaBasePropertiesWithDefaults

`func NewUserSchemaBasePropertiesWithDefaults() *UserSchemaBaseProperties`

NewUserSchemaBasePropertiesWithDefaults instantiates a new UserSchemaBaseProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *UserSchemaBaseProperties) GetCity() UserSchemaAttribute`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UserSchemaBaseProperties) GetCityOk() (*UserSchemaAttribute, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UserSchemaBaseProperties) SetCity(v UserSchemaAttribute)`

SetCity sets City field to given value.

### HasCity

`func (o *UserSchemaBaseProperties) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCostCenter

`func (o *UserSchemaBaseProperties) GetCostCenter() UserSchemaAttribute`

GetCostCenter returns the CostCenter field if non-nil, zero value otherwise.

### GetCostCenterOk

`func (o *UserSchemaBaseProperties) GetCostCenterOk() (*UserSchemaAttribute, bool)`

GetCostCenterOk returns a tuple with the CostCenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCenter

`func (o *UserSchemaBaseProperties) SetCostCenter(v UserSchemaAttribute)`

SetCostCenter sets CostCenter field to given value.

### HasCostCenter

`func (o *UserSchemaBaseProperties) HasCostCenter() bool`

HasCostCenter returns a boolean if a field has been set.

### GetCountryCode

`func (o *UserSchemaBaseProperties) GetCountryCode() UserSchemaAttribute`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserSchemaBaseProperties) GetCountryCodeOk() (*UserSchemaAttribute, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserSchemaBaseProperties) SetCountryCode(v UserSchemaAttribute)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UserSchemaBaseProperties) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDepartment

`func (o *UserSchemaBaseProperties) GetDepartment() UserSchemaAttribute`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UserSchemaBaseProperties) GetDepartmentOk() (*UserSchemaAttribute, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UserSchemaBaseProperties) SetDepartment(v UserSchemaAttribute)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UserSchemaBaseProperties) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserSchemaBaseProperties) GetDisplayName() UserSchemaAttribute`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserSchemaBaseProperties) GetDisplayNameOk() (*UserSchemaAttribute, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserSchemaBaseProperties) SetDisplayName(v UserSchemaAttribute)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserSchemaBaseProperties) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDivision

`func (o *UserSchemaBaseProperties) GetDivision() UserSchemaAttribute`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *UserSchemaBaseProperties) GetDivisionOk() (*UserSchemaAttribute, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *UserSchemaBaseProperties) SetDivision(v UserSchemaAttribute)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *UserSchemaBaseProperties) HasDivision() bool`

HasDivision returns a boolean if a field has been set.

### GetEmail

`func (o *UserSchemaBaseProperties) GetEmail() UserSchemaAttribute`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSchemaBaseProperties) GetEmailOk() (*UserSchemaAttribute, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSchemaBaseProperties) SetEmail(v UserSchemaAttribute)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSchemaBaseProperties) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmployeeNumber

`func (o *UserSchemaBaseProperties) GetEmployeeNumber() UserSchemaAttribute`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UserSchemaBaseProperties) GetEmployeeNumberOk() (*UserSchemaAttribute, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UserSchemaBaseProperties) SetEmployeeNumber(v UserSchemaAttribute)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *UserSchemaBaseProperties) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### GetFirstName

`func (o *UserSchemaBaseProperties) GetFirstName() UserSchemaAttribute`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserSchemaBaseProperties) GetFirstNameOk() (*UserSchemaAttribute, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserSchemaBaseProperties) SetFirstName(v UserSchemaAttribute)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserSchemaBaseProperties) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetHonorificPrefix

`func (o *UserSchemaBaseProperties) GetHonorificPrefix() UserSchemaAttribute`

GetHonorificPrefix returns the HonorificPrefix field if non-nil, zero value otherwise.

### GetHonorificPrefixOk

`func (o *UserSchemaBaseProperties) GetHonorificPrefixOk() (*UserSchemaAttribute, bool)`

GetHonorificPrefixOk returns a tuple with the HonorificPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorificPrefix

`func (o *UserSchemaBaseProperties) SetHonorificPrefix(v UserSchemaAttribute)`

SetHonorificPrefix sets HonorificPrefix field to given value.

### HasHonorificPrefix

`func (o *UserSchemaBaseProperties) HasHonorificPrefix() bool`

HasHonorificPrefix returns a boolean if a field has been set.

### GetHonorificSuffix

`func (o *UserSchemaBaseProperties) GetHonorificSuffix() UserSchemaAttribute`

GetHonorificSuffix returns the HonorificSuffix field if non-nil, zero value otherwise.

### GetHonorificSuffixOk

`func (o *UserSchemaBaseProperties) GetHonorificSuffixOk() (*UserSchemaAttribute, bool)`

GetHonorificSuffixOk returns a tuple with the HonorificSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHonorificSuffix

`func (o *UserSchemaBaseProperties) SetHonorificSuffix(v UserSchemaAttribute)`

SetHonorificSuffix sets HonorificSuffix field to given value.

### HasHonorificSuffix

`func (o *UserSchemaBaseProperties) HasHonorificSuffix() bool`

HasHonorificSuffix returns a boolean if a field has been set.

### GetLastName

`func (o *UserSchemaBaseProperties) GetLastName() UserSchemaAttribute`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserSchemaBaseProperties) GetLastNameOk() (*UserSchemaAttribute, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserSchemaBaseProperties) SetLastName(v UserSchemaAttribute)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserSchemaBaseProperties) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLocale

`func (o *UserSchemaBaseProperties) GetLocale() UserSchemaAttribute`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *UserSchemaBaseProperties) GetLocaleOk() (*UserSchemaAttribute, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *UserSchemaBaseProperties) SetLocale(v UserSchemaAttribute)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *UserSchemaBaseProperties) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLogin

`func (o *UserSchemaBaseProperties) GetLogin() UserSchemaAttribute`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserSchemaBaseProperties) GetLoginOk() (*UserSchemaAttribute, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserSchemaBaseProperties) SetLogin(v UserSchemaAttribute)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserSchemaBaseProperties) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetManager

`func (o *UserSchemaBaseProperties) GetManager() UserSchemaAttribute`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *UserSchemaBaseProperties) GetManagerOk() (*UserSchemaAttribute, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *UserSchemaBaseProperties) SetManager(v UserSchemaAttribute)`

SetManager sets Manager field to given value.

### HasManager

`func (o *UserSchemaBaseProperties) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetManagerId

`func (o *UserSchemaBaseProperties) GetManagerId() UserSchemaAttribute`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *UserSchemaBaseProperties) GetManagerIdOk() (*UserSchemaAttribute, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *UserSchemaBaseProperties) SetManagerId(v UserSchemaAttribute)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *UserSchemaBaseProperties) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetMiddleName

`func (o *UserSchemaBaseProperties) GetMiddleName() UserSchemaAttribute`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *UserSchemaBaseProperties) GetMiddleNameOk() (*UserSchemaAttribute, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *UserSchemaBaseProperties) SetMiddleName(v UserSchemaAttribute)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *UserSchemaBaseProperties) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### GetMobilePhone

`func (o *UserSchemaBaseProperties) GetMobilePhone() UserSchemaAttribute`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *UserSchemaBaseProperties) GetMobilePhoneOk() (*UserSchemaAttribute, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *UserSchemaBaseProperties) SetMobilePhone(v UserSchemaAttribute)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *UserSchemaBaseProperties) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetNickName

`func (o *UserSchemaBaseProperties) GetNickName() UserSchemaAttribute`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *UserSchemaBaseProperties) GetNickNameOk() (*UserSchemaAttribute, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *UserSchemaBaseProperties) SetNickName(v UserSchemaAttribute)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *UserSchemaBaseProperties) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetOrganization

`func (o *UserSchemaBaseProperties) GetOrganization() UserSchemaAttribute`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UserSchemaBaseProperties) GetOrganizationOk() (*UserSchemaAttribute, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UserSchemaBaseProperties) SetOrganization(v UserSchemaAttribute)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *UserSchemaBaseProperties) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPostalAddress

`func (o *UserSchemaBaseProperties) GetPostalAddress() UserSchemaAttribute`

GetPostalAddress returns the PostalAddress field if non-nil, zero value otherwise.

### GetPostalAddressOk

`func (o *UserSchemaBaseProperties) GetPostalAddressOk() (*UserSchemaAttribute, bool)`

GetPostalAddressOk returns a tuple with the PostalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalAddress

`func (o *UserSchemaBaseProperties) SetPostalAddress(v UserSchemaAttribute)`

SetPostalAddress sets PostalAddress field to given value.

### HasPostalAddress

`func (o *UserSchemaBaseProperties) HasPostalAddress() bool`

HasPostalAddress returns a boolean if a field has been set.

### GetPreferredLanguage

`func (o *UserSchemaBaseProperties) GetPreferredLanguage() UserSchemaAttribute`

GetPreferredLanguage returns the PreferredLanguage field if non-nil, zero value otherwise.

### GetPreferredLanguageOk

`func (o *UserSchemaBaseProperties) GetPreferredLanguageOk() (*UserSchemaAttribute, bool)`

GetPreferredLanguageOk returns a tuple with the PreferredLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLanguage

`func (o *UserSchemaBaseProperties) SetPreferredLanguage(v UserSchemaAttribute)`

SetPreferredLanguage sets PreferredLanguage field to given value.

### HasPreferredLanguage

`func (o *UserSchemaBaseProperties) HasPreferredLanguage() bool`

HasPreferredLanguage returns a boolean if a field has been set.

### GetPrimaryPhone

`func (o *UserSchemaBaseProperties) GetPrimaryPhone() UserSchemaAttribute`

GetPrimaryPhone returns the PrimaryPhone field if non-nil, zero value otherwise.

### GetPrimaryPhoneOk

`func (o *UserSchemaBaseProperties) GetPrimaryPhoneOk() (*UserSchemaAttribute, bool)`

GetPrimaryPhoneOk returns a tuple with the PrimaryPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPhone

`func (o *UserSchemaBaseProperties) SetPrimaryPhone(v UserSchemaAttribute)`

SetPrimaryPhone sets PrimaryPhone field to given value.

### HasPrimaryPhone

`func (o *UserSchemaBaseProperties) HasPrimaryPhone() bool`

HasPrimaryPhone returns a boolean if a field has been set.

### GetProfileUrl

`func (o *UserSchemaBaseProperties) GetProfileUrl() UserSchemaAttribute`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *UserSchemaBaseProperties) GetProfileUrlOk() (*UserSchemaAttribute, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *UserSchemaBaseProperties) SetProfileUrl(v UserSchemaAttribute)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *UserSchemaBaseProperties) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetSecondEmail

`func (o *UserSchemaBaseProperties) GetSecondEmail() UserSchemaAttribute`

GetSecondEmail returns the SecondEmail field if non-nil, zero value otherwise.

### GetSecondEmailOk

`func (o *UserSchemaBaseProperties) GetSecondEmailOk() (*UserSchemaAttribute, bool)`

GetSecondEmailOk returns a tuple with the SecondEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondEmail

`func (o *UserSchemaBaseProperties) SetSecondEmail(v UserSchemaAttribute)`

SetSecondEmail sets SecondEmail field to given value.

### HasSecondEmail

`func (o *UserSchemaBaseProperties) HasSecondEmail() bool`

HasSecondEmail returns a boolean if a field has been set.

### GetState

`func (o *UserSchemaBaseProperties) GetState() UserSchemaAttribute`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserSchemaBaseProperties) GetStateOk() (*UserSchemaAttribute, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserSchemaBaseProperties) SetState(v UserSchemaAttribute)`

SetState sets State field to given value.

### HasState

`func (o *UserSchemaBaseProperties) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreetAddress

`func (o *UserSchemaBaseProperties) GetStreetAddress() UserSchemaAttribute`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *UserSchemaBaseProperties) GetStreetAddressOk() (*UserSchemaAttribute, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *UserSchemaBaseProperties) SetStreetAddress(v UserSchemaAttribute)`

SetStreetAddress sets StreetAddress field to given value.

### HasStreetAddress

`func (o *UserSchemaBaseProperties) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.

### GetTimezone

`func (o *UserSchemaBaseProperties) GetTimezone() UserSchemaAttribute`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserSchemaBaseProperties) GetTimezoneOk() (*UserSchemaAttribute, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserSchemaBaseProperties) SetTimezone(v UserSchemaAttribute)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserSchemaBaseProperties) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTitle

`func (o *UserSchemaBaseProperties) GetTitle() UserSchemaAttribute`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserSchemaBaseProperties) GetTitleOk() (*UserSchemaAttribute, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserSchemaBaseProperties) SetTitle(v UserSchemaAttribute)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserSchemaBaseProperties) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUserType

`func (o *UserSchemaBaseProperties) GetUserType() UserSchemaAttribute`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserSchemaBaseProperties) GetUserTypeOk() (*UserSchemaAttribute, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserSchemaBaseProperties) SetUserType(v UserSchemaAttribute)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserSchemaBaseProperties) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetZipCode

`func (o *UserSchemaBaseProperties) GetZipCode() UserSchemaAttribute`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *UserSchemaBaseProperties) GetZipCodeOk() (*UserSchemaAttribute, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *UserSchemaBaseProperties) SetZipCode(v UserSchemaAttribute)`

SetZipCode sets ZipCode field to given value.

### HasZipCode

`func (o *UserSchemaBaseProperties) HasZipCode() bool`

HasZipCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


