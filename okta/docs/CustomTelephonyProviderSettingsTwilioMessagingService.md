# CustomTelephonyProviderSettingsTwilioMessagingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TwilioMessageSid** | Pointer to **string** | The Twilio Messaging Service SID used for sending SMS messages. You can find this value in your Twilio console.  This method uses Twilio&#39;s [Programmable Messaging API](https://www.twilio.com/docs/messaging). | [optional] 

## Methods

### NewCustomTelephonyProviderSettingsTwilioMessagingService

`func NewCustomTelephonyProviderSettingsTwilioMessagingService() *CustomTelephonyProviderSettingsTwilioMessagingService`

NewCustomTelephonyProviderSettingsTwilioMessagingService instantiates a new CustomTelephonyProviderSettingsTwilioMessagingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomTelephonyProviderSettingsTwilioMessagingServiceWithDefaults

`func NewCustomTelephonyProviderSettingsTwilioMessagingServiceWithDefaults() *CustomTelephonyProviderSettingsTwilioMessagingService`

NewCustomTelephonyProviderSettingsTwilioMessagingServiceWithDefaults instantiates a new CustomTelephonyProviderSettingsTwilioMessagingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwilioMessageSid

`func (o *CustomTelephonyProviderSettingsTwilioMessagingService) GetTwilioMessageSid() string`

GetTwilioMessageSid returns the TwilioMessageSid field if non-nil, zero value otherwise.

### GetTwilioMessageSidOk

`func (o *CustomTelephonyProviderSettingsTwilioMessagingService) GetTwilioMessageSidOk() (*string, bool)`

GetTwilioMessageSidOk returns a tuple with the TwilioMessageSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioMessageSid

`func (o *CustomTelephonyProviderSettingsTwilioMessagingService) SetTwilioMessageSid(v string)`

SetTwilioMessageSid sets TwilioMessageSid field to given value.

### HasTwilioMessageSid

`func (o *CustomTelephonyProviderSettingsTwilioMessagingService) HasTwilioMessageSid() bool`

HasTwilioMessageSid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


