# CustomTelephonyProviderCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether the provider is enabled and can be used | [optional] 
**Id** | Pointer to **string** | ID of the custom telephony provider | [optional] 
**IsPrimaryProvider** | Pointer to **bool** | Indicates whether the provider is the primary telephony provider | [optional] 
**ProviderCapability** | Pointer to **string** | The types of telephony operations (SMS or Voice) that you use with your telephony provider.  &#x60;ALL&#x60; is the only valid value. It indicates that your provider can handle both SMS messages and voice calls. You&#39;re not required to use both types of telephony operations, but your provider can support both. | [optional] 
**ProviderName** | Pointer to **string** | Name of the telephony provider | [optional] 
**ProviderSettings** | Pointer to [**CustomTelephonyProviderSettings**](CustomTelephonyProviderSettings.md) |  | [optional] 
**ProviderSid** | Pointer to **string** | The account string identifier (SID) for your telephony provider account. Your telephony provider gives you this SID. | [optional] 

## Methods

### NewCustomTelephonyProviderCredentialResponse

`func NewCustomTelephonyProviderCredentialResponse() *CustomTelephonyProviderCredentialResponse`

NewCustomTelephonyProviderCredentialResponse instantiates a new CustomTelephonyProviderCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderCredentialResponseWithDefaults

`func NewCustomTelephonyProviderCredentialResponseWithDefaults() *CustomTelephonyProviderCredentialResponse`

NewCustomTelephonyProviderCredentialResponseWithDefaults instantiates a new CustomTelephonyProviderCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CustomTelephonyProviderCredentialResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CustomTelephonyProviderCredentialResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CustomTelephonyProviderCredentialResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CustomTelephonyProviderCredentialResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *CustomTelephonyProviderCredentialResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomTelephonyProviderCredentialResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomTelephonyProviderCredentialResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomTelephonyProviderCredentialResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrimaryProvider

`func (o *CustomTelephonyProviderCredentialResponse) GetIsPrimaryProvider() bool`

GetIsPrimaryProvider returns the IsPrimaryProvider field if non-nil, zero value otherwise.

### GetIsPrimaryProviderOk

`func (o *CustomTelephonyProviderCredentialResponse) GetIsPrimaryProviderOk() (*bool, bool)`

GetIsPrimaryProviderOk returns a tuple with the IsPrimaryProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimaryProvider

`func (o *CustomTelephonyProviderCredentialResponse) SetIsPrimaryProvider(v bool)`

SetIsPrimaryProvider sets IsPrimaryProvider field to given value.

### HasIsPrimaryProvider

`func (o *CustomTelephonyProviderCredentialResponse) HasIsPrimaryProvider() bool`

HasIsPrimaryProvider returns a boolean if a field has been set.

### GetProviderCapability

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderCapability() string`

GetProviderCapability returns the ProviderCapability field if non-nil, zero value otherwise.

### GetProviderCapabilityOk

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderCapabilityOk() (*string, bool)`

GetProviderCapabilityOk returns a tuple with the ProviderCapability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCapability

`func (o *CustomTelephonyProviderCredentialResponse) SetProviderCapability(v string)`

SetProviderCapability sets ProviderCapability field to given value.

### HasProviderCapability

`func (o *CustomTelephonyProviderCredentialResponse) HasProviderCapability() bool`

HasProviderCapability returns a boolean if a field has been set.

### GetProviderName

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *CustomTelephonyProviderCredentialResponse) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *CustomTelephonyProviderCredentialResponse) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetProviderSettings

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderSettings() CustomTelephonyProviderSettings`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderSettingsOk() (*CustomTelephonyProviderSettings, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *CustomTelephonyProviderCredentialResponse) SetProviderSettings(v CustomTelephonyProviderSettings)`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *CustomTelephonyProviderCredentialResponse) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetProviderSid

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderSid() string`

GetProviderSid returns the ProviderSid field if non-nil, zero value otherwise.

### GetProviderSidOk

`func (o *CustomTelephonyProviderCredentialResponse) GetProviderSidOk() (*string, bool)`

GetProviderSidOk returns a tuple with the ProviderSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSid

`func (o *CustomTelephonyProviderCredentialResponse) SetProviderSid(v string)`

SetProviderSid sets ProviderSid field to given value.

### HasProviderSid

`func (o *CustomTelephonyProviderCredentialResponse) HasProviderSid() bool`

HasProviderSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


