# DevicePostureChecksRemediationSettingsMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultI18nKey** | Pointer to **string** | Default i18n key for the message. This property is only relevant if type is set to &#x60;BUILTIN&#x60;. If type is set to &#x60;CUSTOM&#x60;, this field is ignored. | [optional] 
**CustomText** | Pointer to **string** | Custom text for the message | [optional] 

## Methods

### NewDevicePostureChecksRemediationSettingsMessage

`func NewDevicePostureChecksRemediationSettingsMessage() *DevicePostureChecksRemediationSettingsMessage`

NewDevicePostureChecksRemediationSettingsMessage instantiates a new DevicePostureChecksRemediationSettingsMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePostureChecksRemediationSettingsMessageWithDefaults

`func NewDevicePostureChecksRemediationSettingsMessageWithDefaults() *DevicePostureChecksRemediationSettingsMessage`

NewDevicePostureChecksRemediationSettingsMessageWithDefaults instantiates a new DevicePostureChecksRemediationSettingsMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultI18nKey

`func (o *DevicePostureChecksRemediationSettingsMessage) GetDefaultI18nKey() string`

GetDefaultI18nKey returns the DefaultI18nKey field if non-nil, zero value otherwise.

### GetDefaultI18nKeyOk

`func (o *DevicePostureChecksRemediationSettingsMessage) GetDefaultI18nKeyOk() (*string, bool)`

GetDefaultI18nKeyOk returns a tuple with the DefaultI18nKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultI18nKey

`func (o *DevicePostureChecksRemediationSettingsMessage) SetDefaultI18nKey(v string)`

SetDefaultI18nKey sets DefaultI18nKey field to given value.

### HasDefaultI18nKey

`func (o *DevicePostureChecksRemediationSettingsMessage) HasDefaultI18nKey() bool`

HasDefaultI18nKey returns a boolean if a field has been set.

### GetCustomText

`func (o *DevicePostureChecksRemediationSettingsMessage) GetCustomText() string`

GetCustomText returns the CustomText field if non-nil, zero value otherwise.

### GetCustomTextOk

`func (o *DevicePostureChecksRemediationSettingsMessage) GetCustomTextOk() (*string, bool)`

GetCustomTextOk returns a tuple with the CustomText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomText

`func (o *DevicePostureChecksRemediationSettingsMessage) SetCustomText(v string)`

SetCustomText sets CustomText field to given value.

### HasCustomText

`func (o *DevicePostureChecksRemediationSettingsMessage) HasCustomText() bool`

HasCustomText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


