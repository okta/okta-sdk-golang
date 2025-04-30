# TelephonyRequestDataMessageProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgTemplate** | Pointer to **string** | Default or Okta org configured sms or voice message template | [optional] 
**PhoneNumber** | Pointer to **string** | The Okta&#39;s user&#39;s phone number | [optional] 
**OtpExpires** | Pointer to **string** | The time when OTP expires | [optional] 
**DeliveryChannel** | Pointer to **string** | The channel for OTP delivery - SMS or voice | [optional] 
**OtpCode** | Pointer to **string** | The OTP code requested by the Okta user | [optional] 
**Locale** | Pointer to **string** | The locale associated with the Okta user | [optional] 

## Methods

### NewTelephonyRequestDataMessageProfile

`func NewTelephonyRequestDataMessageProfile() *TelephonyRequestDataMessageProfile`

NewTelephonyRequestDataMessageProfile instantiates a new TelephonyRequestDataMessageProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyRequestDataMessageProfileWithDefaults

`func NewTelephonyRequestDataMessageProfileWithDefaults() *TelephonyRequestDataMessageProfile`

NewTelephonyRequestDataMessageProfileWithDefaults instantiates a new TelephonyRequestDataMessageProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgTemplate

`func (o *TelephonyRequestDataMessageProfile) GetMsgTemplate() string`

GetMsgTemplate returns the MsgTemplate field if non-nil, zero value otherwise.

### GetMsgTemplateOk

`func (o *TelephonyRequestDataMessageProfile) GetMsgTemplateOk() (*string, bool)`

GetMsgTemplateOk returns a tuple with the MsgTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgTemplate

`func (o *TelephonyRequestDataMessageProfile) SetMsgTemplate(v string)`

SetMsgTemplate sets MsgTemplate field to given value.

### HasMsgTemplate

`func (o *TelephonyRequestDataMessageProfile) HasMsgTemplate() bool`

HasMsgTemplate returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *TelephonyRequestDataMessageProfile) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *TelephonyRequestDataMessageProfile) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *TelephonyRequestDataMessageProfile) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *TelephonyRequestDataMessageProfile) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetOtpExpires

`func (o *TelephonyRequestDataMessageProfile) GetOtpExpires() string`

GetOtpExpires returns the OtpExpires field if non-nil, zero value otherwise.

### GetOtpExpiresOk

`func (o *TelephonyRequestDataMessageProfile) GetOtpExpiresOk() (*string, bool)`

GetOtpExpiresOk returns a tuple with the OtpExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpExpires

`func (o *TelephonyRequestDataMessageProfile) SetOtpExpires(v string)`

SetOtpExpires sets OtpExpires field to given value.

### HasOtpExpires

`func (o *TelephonyRequestDataMessageProfile) HasOtpExpires() bool`

HasOtpExpires returns a boolean if a field has been set.

### GetDeliveryChannel

`func (o *TelephonyRequestDataMessageProfile) GetDeliveryChannel() string`

GetDeliveryChannel returns the DeliveryChannel field if non-nil, zero value otherwise.

### GetDeliveryChannelOk

`func (o *TelephonyRequestDataMessageProfile) GetDeliveryChannelOk() (*string, bool)`

GetDeliveryChannelOk returns a tuple with the DeliveryChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryChannel

`func (o *TelephonyRequestDataMessageProfile) SetDeliveryChannel(v string)`

SetDeliveryChannel sets DeliveryChannel field to given value.

### HasDeliveryChannel

`func (o *TelephonyRequestDataMessageProfile) HasDeliveryChannel() bool`

HasDeliveryChannel returns a boolean if a field has been set.

### GetOtpCode

`func (o *TelephonyRequestDataMessageProfile) GetOtpCode() string`

GetOtpCode returns the OtpCode field if non-nil, zero value otherwise.

### GetOtpCodeOk

`func (o *TelephonyRequestDataMessageProfile) GetOtpCodeOk() (*string, bool)`

GetOtpCodeOk returns a tuple with the OtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpCode

`func (o *TelephonyRequestDataMessageProfile) SetOtpCode(v string)`

SetOtpCode sets OtpCode field to given value.

### HasOtpCode

`func (o *TelephonyRequestDataMessageProfile) HasOtpCode() bool`

HasOtpCode returns a boolean if a field has been set.

### GetLocale

`func (o *TelephonyRequestDataMessageProfile) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *TelephonyRequestDataMessageProfile) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *TelephonyRequestDataMessageProfile) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *TelephonyRequestDataMessageProfile) HasLocale() bool`

HasLocale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


