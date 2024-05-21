# ReplaceLogStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Unique name for the Log Stream object | 
**Type** | **string** | Specifies the streaming provider used  Supported providers:   * &#x60;aws_eventbridge&#x60; ([AWS EventBridge](https://aws.amazon.com/eventbridge))   * &#x60;splunk_cloud_logstreaming&#x60; ([Splunk Cloud](https://www.splunk.com/en_us/software/splunk-cloud-platform.html))  Select the provider type to see provider-specific configurations in the &#x60;settings&#x60; property: | 
**Settings** | [**LogStreamSettingsSplunkPut**](LogStreamSettingsSplunkPut.md) |  | 

## Methods

### NewReplaceLogStreamRequest

`func NewReplaceLogStreamRequest(name string, type_ string, settings LogStreamSettingsSplunkPut, ) *ReplaceLogStreamRequest`

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

`func (o *ReplaceLogStreamRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReplaceLogStreamRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReplaceLogStreamRequest) SetType(v string)`

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


