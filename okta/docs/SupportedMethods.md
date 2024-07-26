# SupportedMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**SupportedMethodsSettings**](SupportedMethodsSettings.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type of authenticator method | [optional] 

## Methods

### NewSupportedMethods

`func NewSupportedMethods() *SupportedMethods`

NewSupportedMethods instantiates a new SupportedMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedMethodsWithDefaults

`func NewSupportedMethodsWithDefaults() *SupportedMethods`

NewSupportedMethodsWithDefaults instantiates a new SupportedMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *SupportedMethods) GetSettings() SupportedMethodsSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SupportedMethods) GetSettingsOk() (*SupportedMethodsSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SupportedMethods) SetSettings(v SupportedMethodsSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SupportedMethods) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStatus

`func (o *SupportedMethods) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportedMethods) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportedMethods) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportedMethods) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SupportedMethods) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupportedMethods) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupportedMethods) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SupportedMethods) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


