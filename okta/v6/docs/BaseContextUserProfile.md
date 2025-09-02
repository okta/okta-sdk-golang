# BaseContextUserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Login** | Pointer to **string** | The username used to identify the user. This is often the user&#39;s email address. | [optional] 
**FirstName** | Pointer to **string** | The first name of the user | [optional] 
**LastName** | Pointer to **string** | The last name of the user | [optional] 
**Locale** | Pointer to **string** | The user&#39;s default location for purposes of localizing items such as currency, date time format, numerical representations, and so on. A locale value is a concatenation of the [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) two-letter language code, an underscore, and the [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) two-letter country code. For example, &#x60;en_US&#x60; specifies the language English and country US. This value is &#x60;en_US&#x60; by default. | [optional] 
**TimeZone** | Pointer to **string** | The user&#39;s timezone | [optional] 

## Methods

### NewBaseContextUserProfile

`func NewBaseContextUserProfile() *BaseContextUserProfile`

NewBaseContextUserProfile instantiates a new BaseContextUserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContextUserProfileWithDefaults

`func NewBaseContextUserProfileWithDefaults() *BaseContextUserProfile`

NewBaseContextUserProfileWithDefaults instantiates a new BaseContextUserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogin

`func (o *BaseContextUserProfile) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *BaseContextUserProfile) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *BaseContextUserProfile) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *BaseContextUserProfile) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetFirstName

`func (o *BaseContextUserProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BaseContextUserProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BaseContextUserProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BaseContextUserProfile) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *BaseContextUserProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BaseContextUserProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BaseContextUserProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BaseContextUserProfile) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLocale

`func (o *BaseContextUserProfile) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *BaseContextUserProfile) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *BaseContextUserProfile) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *BaseContextUserProfile) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetTimeZone

`func (o *BaseContextUserProfile) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *BaseContextUserProfile) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *BaseContextUserProfile) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *BaseContextUserProfile) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


