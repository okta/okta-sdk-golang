# CustomTelephonyProviderSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Call** | Pointer to [**CustomTelephonyProviderSettingsCall**](CustomTelephonyProviderSettingsCall.md) |  | [optional] 
**Sms** | Pointer to [**CustomTelephonyProviderSettingsSms**](CustomTelephonyProviderSettingsSms.md) |  | [optional] 

## Methods

### NewCustomTelephonyProviderSettings

`func NewCustomTelephonyProviderSettings() *CustomTelephonyProviderSettings`

NewCustomTelephonyProviderSettings instantiates a new CustomTelephonyProviderSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderSettingsWithDefaults

`func NewCustomTelephonyProviderSettingsWithDefaults() *CustomTelephonyProviderSettings`

NewCustomTelephonyProviderSettingsWithDefaults instantiates a new CustomTelephonyProviderSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCall

`func (o *CustomTelephonyProviderSettings) GetCall() CustomTelephonyProviderSettingsCall`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *CustomTelephonyProviderSettings) GetCallOk() (*CustomTelephonyProviderSettingsCall, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *CustomTelephonyProviderSettings) SetCall(v CustomTelephonyProviderSettingsCall)`

SetCall sets Call field to given value.

### HasCall

`func (o *CustomTelephonyProviderSettings) HasCall() bool`

HasCall returns a boolean if a field has been set.

### GetSms

`func (o *CustomTelephonyProviderSettings) GetSms() CustomTelephonyProviderSettingsSms`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *CustomTelephonyProviderSettings) GetSmsOk() (*CustomTelephonyProviderSettingsSms, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *CustomTelephonyProviderSettings) SetSms(v CustomTelephonyProviderSettingsSms)`

SetSms sets Sms field to given value.

### HasSms

`func (o *CustomTelephonyProviderSettings) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


