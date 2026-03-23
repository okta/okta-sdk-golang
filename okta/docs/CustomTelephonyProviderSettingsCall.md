# CustomTelephonyProviderSettingsCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TwilioVerifySid** | Pointer to **string** | The Twilio Verify Service SID used for sending verification messages or calls. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Verify API](https://www.twilio.com/docs/verify/api). | [optional] 
**TwilioPhoneNumber** | Pointer to **string** | The Twilio phone number that&#39;s used for sending SMS messages or voice calls. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Programmable Messaging API](https://www.twilio.com/docs/messaging). | [optional] 
**TwilioCallerId** | Pointer to **string** | The Twilio caller ID that&#39;s used for making calls. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Programmable Voice API](https://www.twilio.com/docs/voice). | [optional] 
**TelesignService** | Pointer to **string** | The Telesign service identifier used for sending making calls. You can find this value in your Telesign console.  The &#x60;telesignVerifyService&#x60; method uses Telesign&#39;s [Verify API](https://developer.telesign.com/enterprise/docs/verify-api-overview). And the &#x60;telesignVoiceService&#x60; method uses Telesign&#39;s [Voice API](https://developer.telesign.com/enterprise/docs/voice-overview). | [optional] 

## Methods

### NewCustomTelephonyProviderSettingsCall

`func NewCustomTelephonyProviderSettingsCall() *CustomTelephonyProviderSettingsCall`

NewCustomTelephonyProviderSettingsCall instantiates a new CustomTelephonyProviderSettingsCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderSettingsCallWithDefaults

`func NewCustomTelephonyProviderSettingsCallWithDefaults() *CustomTelephonyProviderSettingsCall`

NewCustomTelephonyProviderSettingsCallWithDefaults instantiates a new CustomTelephonyProviderSettingsCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwilioVerifySid

`func (o *CustomTelephonyProviderSettingsCall) GetTwilioVerifySid() string`

GetTwilioVerifySid returns the TwilioVerifySid field if non-nil, zero value otherwise.

### GetTwilioVerifySidOk

`func (o *CustomTelephonyProviderSettingsCall) GetTwilioVerifySidOk() (*string, bool)`

GetTwilioVerifySidOk returns a tuple with the TwilioVerifySid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioVerifySid

`func (o *CustomTelephonyProviderSettingsCall) SetTwilioVerifySid(v string)`

SetTwilioVerifySid sets TwilioVerifySid field to given value.

### HasTwilioVerifySid

`func (o *CustomTelephonyProviderSettingsCall) HasTwilioVerifySid() bool`

HasTwilioVerifySid returns a boolean if a field has been set.

### GetTwilioPhoneNumber

`func (o *CustomTelephonyProviderSettingsCall) GetTwilioPhoneNumber() string`

GetTwilioPhoneNumber returns the TwilioPhoneNumber field if non-nil, zero value otherwise.

### GetTwilioPhoneNumberOk

`func (o *CustomTelephonyProviderSettingsCall) GetTwilioPhoneNumberOk() (*string, bool)`

GetTwilioPhoneNumberOk returns a tuple with the TwilioPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioPhoneNumber

`func (o *CustomTelephonyProviderSettingsCall) SetTwilioPhoneNumber(v string)`

SetTwilioPhoneNumber sets TwilioPhoneNumber field to given value.

### HasTwilioPhoneNumber

`func (o *CustomTelephonyProviderSettingsCall) HasTwilioPhoneNumber() bool`

HasTwilioPhoneNumber returns a boolean if a field has been set.

### GetTwilioCallerId

`func (o *CustomTelephonyProviderSettingsCall) GetTwilioCallerId() string`

GetTwilioCallerId returns the TwilioCallerId field if non-nil, zero value otherwise.

### GetTwilioCallerIdOk

`func (o *CustomTelephonyProviderSettingsCall) GetTwilioCallerIdOk() (*string, bool)`

GetTwilioCallerIdOk returns a tuple with the TwilioCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioCallerId

`func (o *CustomTelephonyProviderSettingsCall) SetTwilioCallerId(v string)`

SetTwilioCallerId sets TwilioCallerId field to given value.

### HasTwilioCallerId

`func (o *CustomTelephonyProviderSettingsCall) HasTwilioCallerId() bool`

HasTwilioCallerId returns a boolean if a field has been set.

### GetTelesignService

`func (o *CustomTelephonyProviderSettingsCall) GetTelesignService() string`

GetTelesignService returns the TelesignService field if non-nil, zero value otherwise.

### GetTelesignServiceOk

`func (o *CustomTelephonyProviderSettingsCall) GetTelesignServiceOk() (*string, bool)`

GetTelesignServiceOk returns a tuple with the TelesignService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelesignService

`func (o *CustomTelephonyProviderSettingsCall) SetTelesignService(v string)`

SetTelesignService sets TelesignService field to given value.

### HasTelesignService

`func (o *CustomTelephonyProviderSettingsCall) HasTelesignService() bool`

HasTelesignService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


