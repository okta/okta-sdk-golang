# CustomTelephonyProviderCredentialUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the custom telephony provider | [optional] 
**ProviderAuthToken** | Pointer to **string** | The authentication token that&#39;s used to authenticate requests to the telephony provider. Your telephony provider gives you this token. | [optional] 
**ProviderSettings** | Pointer to [**CustomTelephonyProviderSettings**](CustomTelephonyProviderSettings.md) |  | [optional] 
**ProviderSid** | Pointer to **string** | The account string identifier (SID) for your telephony provider account. Your telephony provider gives you this SID. | [optional] 

## Methods

### NewCustomTelephonyProviderCredentialUpdateRequest

`func NewCustomTelephonyProviderCredentialUpdateRequest() *CustomTelephonyProviderCredentialUpdateRequest`

NewCustomTelephonyProviderCredentialUpdateRequest instantiates a new CustomTelephonyProviderCredentialUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderCredentialUpdateRequestWithDefaults

`func NewCustomTelephonyProviderCredentialUpdateRequestWithDefaults() *CustomTelephonyProviderCredentialUpdateRequest`

NewCustomTelephonyProviderCredentialUpdateRequestWithDefaults instantiates a new CustomTelephonyProviderCredentialUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomTelephonyProviderCredentialUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomTelephonyProviderCredentialUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProviderAuthToken

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderAuthToken() string`

GetProviderAuthToken returns the ProviderAuthToken field if non-nil, zero value otherwise.

### GetProviderAuthTokenOk

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderAuthTokenOk() (*string, bool)`

GetProviderAuthTokenOk returns a tuple with the ProviderAuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAuthToken

`func (o *CustomTelephonyProviderCredentialUpdateRequest) SetProviderAuthToken(v string)`

SetProviderAuthToken sets ProviderAuthToken field to given value.

### HasProviderAuthToken

`func (o *CustomTelephonyProviderCredentialUpdateRequest) HasProviderAuthToken() bool`

HasProviderAuthToken returns a boolean if a field has been set.

### GetProviderSettings

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSettings() CustomTelephonyProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSettingsOk() (*CustomTelephonyProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *CustomTelephonyProviderCredentialUpdateRequest) SetProviderSettings(v CustomTelephonyProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *CustomTelephonyProviderCredentialUpdateRequest) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetProviderSid

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSid() string`

GetProviderSid returns the ProviderSid field if non-nil, zero value otherwise.

### GetProviderSidOk

`func (o *CustomTelephonyProviderCredentialUpdateRequest) GetProviderSidOk() (*string, bool)`

GetProviderSidOk returns a tuple with the ProviderSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSid

`func (o *CustomTelephonyProviderCredentialUpdateRequest) SetProviderSid(v string)`

SetProviderSid sets ProviderSid field to given value.

### HasProviderSid

`func (o *CustomTelephonyProviderCredentialUpdateRequest) HasProviderSid() bool`

HasProviderSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


