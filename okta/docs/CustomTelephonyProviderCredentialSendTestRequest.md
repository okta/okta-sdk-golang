# CustomTelephonyProviderCredentialSendTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCodeIso2** | Pointer to **string** | The country code for the phone number. Use the [Alpha-2 code from ISO 3166-1](https://www.iso.org/obp/ui/#search) for country codes. | [optional] 
**Factor** | Pointer to **string** | The type of test message to send | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number to which the test message or call is sent | [optional] 

## Methods

### NewCustomTelephonyProviderCredentialSendTestRequest

`func NewCustomTelephonyProviderCredentialSendTestRequest() *CustomTelephonyProviderCredentialSendTestRequest`

NewCustomTelephonyProviderCredentialSendTestRequest instantiates a new CustomTelephonyProviderCredentialSendTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderCredentialSendTestRequestWithDefaults

`func NewCustomTelephonyProviderCredentialSendTestRequestWithDefaults() *CustomTelephonyProviderCredentialSendTestRequest`

NewCustomTelephonyProviderCredentialSendTestRequestWithDefaults instantiates a new CustomTelephonyProviderCredentialSendTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCodeIso2

`func (o *CustomTelephonyProviderCredentialSendTestRequest) GetCountryCodeIso2() string`

GetCountryCodeIso2 returns the CountryCodeIso2 field if non-nil, zero value otherwise.

### GetCountryCodeIso2Ok

`func (o *CustomTelephonyProviderCredentialSendTestRequest) GetCountryCodeIso2Ok() (*string, bool)`

GetCountryCodeIso2Ok returns a tuple with the CountryCodeIso2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeIso2

`func (o *CustomTelephonyProviderCredentialSendTestRequest) SetCountryCodeIso2(v string)`

SetCountryCodeIso2 sets CountryCodeIso2 field to given value.

### HasCountryCodeIso2

`func (o *CustomTelephonyProviderCredentialSendTestRequest) HasCountryCodeIso2() bool`

HasCountryCodeIso2 returns a boolean if a field has been set.

### GetFactor

`func (o *CustomTelephonyProviderCredentialSendTestRequest) GetFactor() string`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *CustomTelephonyProviderCredentialSendTestRequest) GetFactorOk() (*string, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *CustomTelephonyProviderCredentialSendTestRequest) SetFactor(v string)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *CustomTelephonyProviderCredentialSendTestRequest) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomTelephonyProviderCredentialSendTestRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomTelephonyProviderCredentialSendTestRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomTelephonyProviderCredentialSendTestRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomTelephonyProviderCredentialSendTestRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


