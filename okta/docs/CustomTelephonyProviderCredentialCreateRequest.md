# CustomTelephonyProviderCredentialCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderAuthToken** | Pointer to **string** | The authentication token that&#39;s used to authenticate requests to the telephony provider. Your telephony provider gives you this token. | [optional] 
**ProviderCapability** | Pointer to **string** | The types of telephony operations (SMS or Voice) that you use with your telephony provider.  &#x60;ALL&#x60; is the only valid value. It indicates that your provider can handle both SMS messages and voice calls. You&#39;re not required to use both types of telephony operations, but your provider can support both. | [optional] 
**ProviderName** | Pointer to **string** | The name of the telephony provider | [optional] 
**ProviderSettings** | Pointer to [**CustomTelephonyProviderSettings**](CustomTelephonyProviderSettings.md) |  | [optional] 
**ProviderSid** | Pointer to **string** | The account string identifier (SID) for your telephony provider account. Your telephony provider gives you this SID. | [optional] 

## Methods

### NewCustomTelephonyProviderCredentialCreateRequest

`func NewCustomTelephonyProviderCredentialCreateRequest() *CustomTelephonyProviderCredentialCreateRequest`

NewCustomTelephonyProviderCredentialCreateRequest instantiates a new CustomTelephonyProviderCredentialCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderCredentialCreateRequestWithDefaults

`func NewCustomTelephonyProviderCredentialCreateRequestWithDefaults() *CustomTelephonyProviderCredentialCreateRequest`

NewCustomTelephonyProviderCredentialCreateRequestWithDefaults instantiates a new CustomTelephonyProviderCredentialCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderAuthToken

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderAuthToken() string`

GetProviderAuthToken returns the ProviderAuthToken field if non-nil, zero value otherwise.

### GetProviderAuthTokenOk

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderAuthTokenOk() (*string, bool)`

GetProviderAuthTokenOk returns a tuple with the ProviderAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAuthToken

`func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderAuthToken(v string)`

SetProviderAuthToken sets ProviderAuthToken field to given value.

### HasProviderAuthToken

`func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderAuthToken() bool`

HasProviderAuthToken returns a boolean if a field has been set.

### GetProviderCapability

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderCapability() string`

GetProviderCapability returns the ProviderCapability field if non-nil, zero value otherwise.

### GetProviderCapabilityOk

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderCapabilityOk() (*string, bool)`

GetProviderCapabilityOk returns a tuple with the ProviderCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCapability

`func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderCapability(v string)`

SetProviderCapability sets ProviderCapability field to given value.

### HasProviderCapability

`func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderCapability() bool`

HasProviderCapability returns a boolean if a field has been set.

### GetProviderName

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSettings() CustomTelephonyProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSettingsOk() (*CustomTelephonyProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderSettings(v CustomTelephonyProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetProviderSid

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSid() string`

GetProviderSid returns the ProviderSid field if non-nil, zero value otherwise.

### GetProviderSidOk

`func (o *CustomTelephonyProviderCredentialCreateRequest) GetProviderSidOk() (*string, bool)`

GetProviderSidOk returns a tuple with the ProviderSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSid

`func (o *CustomTelephonyProviderCredentialCreateRequest) SetProviderSid(v string)`

SetProviderSid sets ProviderSid field to given value.

### HasProviderSid

`func (o *CustomTelephonyProviderCredentialCreateRequest) HasProviderSid() bool`

HasProviderSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


