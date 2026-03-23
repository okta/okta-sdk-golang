# CustomTelephonyProviderSettingsSms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TwilioVerifySid** | Pointer to **string** | The Twilio Verify Service SID used for sending verification messages or calls. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Verify API](https://www.twilio.com/docs/verify/api). | [optional] 
**TwilioPhoneNumber** | Pointer to **string** | The Twilio phone number that&#39;s used for sending SMS messages or voice calls. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Programmable Messaging API](https://www.twilio.com/docs/messaging). | [optional] 
**TwilioMessageSid** | Pointer to **string** | The Twilio Messaging Service SID used for sending SMS messages. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Programmable Messaging API](https://www.twilio.com/docs/messaging). | [optional] 
**TelesignService** | Pointer to **string** | The Telesign service identifier used for sending SMS messages. You can find this value in your Telesign console.  The &#x60;telesignVerifyService&#x60; method uses Telesign&#39;s [Verify API](https://developer.telesign.com/enterprise/docs/verify-api-overview). And the &#x60;telesignMessagingService&#x60; method uses Telesign&#39;s [SMS](https://developer.telesign.com/enterprise/docs/voice-overview) and [Messaging](https://developer.telesign.com/enterprise/docs/messaging-overview) APIs. | [optional] 

## Methods

### NewCustomTelephonyProviderSettingsSms

`func NewCustomTelephonyProviderSettingsSms() *CustomTelephonyProviderSettingsSms`

NewCustomTelephonyProviderSettingsSms instantiates a new CustomTelephonyProviderSettingsSms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderSettingsSmsWithDefaults

`func NewCustomTelephonyProviderSettingsSmsWithDefaults() *CustomTelephonyProviderSettingsSms`

NewCustomTelephonyProviderSettingsSmsWithDefaults instantiates a new CustomTelephonyProviderSettingsSms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwilioVerifySid

`func (o *CustomTelephonyProviderSettingsSms) GetTwilioVerifySid() string`

GetTwilioVerifySid returns the TwilioVerifySid field if non-nil, zero value otherwise.

### GetTwilioVerifySidOk

`func (o *CustomTelephonyProviderSettingsSms) GetTwilioVerifySidOk() (*string, bool)`

GetTwilioVerifySidOk returns a tuple with the TwilioVerifySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioVerifySid

`func (o *CustomTelephonyProviderSettingsSms) SetTwilioVerifySid(v string)`

SetTwilioVerifySid sets TwilioVerifySid field to given value.

### HasTwilioVerifySid

`func (o *CustomTelephonyProviderSettingsSms) HasTwilioVerifySid() bool`

HasTwilioVerifySid returns a boolean if a field has been set.

### GetTwilioPhoneNumber

`func (o *CustomTelephonyProviderSettingsSms) GetTwilioPhoneNumber() string`

GetTwilioPhoneNumber returns the TwilioPhoneNumber field if non-nil, zero value otherwise.

### GetTwilioPhoneNumberOk

`func (o *CustomTelephonyProviderSettingsSms) GetTwilioPhoneNumberOk() (*string, bool)`

GetTwilioPhoneNumberOk returns a tuple with the TwilioPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioPhoneNumber

`func (o *CustomTelephonyProviderSettingsSms) SetTwilioPhoneNumber(v string)`

SetTwilioPhoneNumber sets TwilioPhoneNumber field to given value.

### HasTwilioPhoneNumber

`func (o *CustomTelephonyProviderSettingsSms) HasTwilioPhoneNumber() bool`

HasTwilioPhoneNumber returns a boolean if a field has been set.

### GetTwilioMessageSid

`func (o *CustomTelephonyProviderSettingsSms) GetTwilioMessageSid() string`

GetTwilioMessageSid returns the TwilioMessageSid field if non-nil, zero value otherwise.

### GetTwilioMessageSidOk

`func (o *CustomTelephonyProviderSettingsSms) GetTwilioMessageSidOk() (*string, bool)`

GetTwilioMessageSidOk returns a tuple with the TwilioMessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioMessageSid

`func (o *CustomTelephonyProviderSettingsSms) SetTwilioMessageSid(v string)`

SetTwilioMessageSid sets TwilioMessageSid field to given value.

### HasTwilioMessageSid

`func (o *CustomTelephonyProviderSettingsSms) HasTwilioMessageSid() bool`

HasTwilioMessageSid returns a boolean if a field has been set.

### GetTelesignService

`func (o *CustomTelephonyProviderSettingsSms) GetTelesignService() string`

GetTelesignService returns the TelesignService field if non-nil, zero value otherwise.

### GetTelesignServiceOk

`func (o *CustomTelephonyProviderSettingsSms) GetTelesignServiceOk() (*string, bool)`

GetTelesignServiceOk returns a tuple with the TelesignService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelesignService

`func (o *CustomTelephonyProviderSettingsSms) SetTelesignService(v string)`

SetTelesignService sets TelesignService field to given value.

### HasTelesignService

`func (o *CustomTelephonyProviderSettingsSms) HasTelesignService() bool`

HasTelesignService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


