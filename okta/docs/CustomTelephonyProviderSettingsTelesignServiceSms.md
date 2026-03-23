# CustomTelephonyProviderSettingsTelesignServiceSms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TelesignService** | Pointer to **string** | The Telesign service identifier used for sending SMS messages. You can find this value in your Telesign console.  The &#x60;telesignVerifyService&#x60; method uses Telesign&#39;s [Verify API](https://developer.telesign.com/enterprise/docs/verify-api-overview). And the &#x60;telesignMessagingService&#x60; method uses Telesign&#39;s [SMS](https://developer.telesign.com/enterprise/docs/voice-overview) and [Messaging](https://developer.telesign.com/enterprise/docs/messaging-overview) APIs. | [optional] 

## Methods

### NewCustomTelephonyProviderSettingsTelesignServiceSms

`func NewCustomTelephonyProviderSettingsTelesignServiceSms() *CustomTelephonyProviderSettingsTelesignServiceSms`

NewCustomTelephonyProviderSettingsTelesignServiceSms instantiates a new CustomTelephonyProviderSettingsTelesignServiceSms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderSettingsTelesignServiceSmsWithDefaults

`func NewCustomTelephonyProviderSettingsTelesignServiceSmsWithDefaults() *CustomTelephonyProviderSettingsTelesignServiceSms`

NewCustomTelephonyProviderSettingsTelesignServiceSmsWithDefaults instantiates a new CustomTelephonyProviderSettingsTelesignServiceSms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTelesignService

`func (o *CustomTelephonyProviderSettingsTelesignServiceSms) GetTelesignService() string`

GetTelesignService returns the TelesignService field if non-nil, zero value otherwise.

### GetTelesignServiceOk

`func (o *CustomTelephonyProviderSettingsTelesignServiceSms) GetTelesignServiceOk() (*string, bool)`

GetTelesignServiceOk returns a tuple with the TelesignService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelesignService

`func (o *CustomTelephonyProviderSettingsTelesignServiceSms) SetTelesignService(v string)`

SetTelesignService sets TelesignService field to given value.

### HasTelesignService

`func (o *CustomTelephonyProviderSettingsTelesignServiceSms) HasTelesignService() bool`

HasTelesignService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


