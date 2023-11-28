# ReplaceLogStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name for the Log Stream object | 
**Type** | [**LogStreamType**](LogStreamType.md) |  | 
**Settings** | [**LogStreamSettingsSplunkPut**](LogStreamSettingsSplunkPut.md) |  | 

## Methods

### NewReplaceLogStreamRequest

`func NewReplaceLogStreamRequest(name string, type_ LogStreamType, settings LogStreamSettingsSplunkPut, ) *ReplaceLogStreamRequest`

NewReplaceLogStreamRequest instantiates a new ReplaceLogStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplaceLogStreamRequestWithDefaults

`func NewReplaceLogStreamRequestWithDefaults() *ReplaceLogStreamRequest`

NewReplaceLogStreamRequestWithDefaults instantiates a new ReplaceLogStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ReplaceLogStreamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplaceLogStreamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplaceLogStreamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ReplaceLogStreamRequest) GetType() LogStreamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReplaceLogStreamRequest) GetTypeOk() (*LogStreamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReplaceLogStreamRequest) SetType(v LogStreamType)`

SetType sets Type field to given value.


### GetSettings

`func (o *ReplaceLogStreamRequest) GetSettings() LogStreamSettingsSplunkPut`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ReplaceLogStreamRequest) GetSettingsOk() (*LogStreamSettingsSplunkPut, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ReplaceLogStreamRequest) SetSettings(v LogStreamSettingsSplunkPut)`

SetSettings sets Settings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


