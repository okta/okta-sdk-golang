# LogStreamSettingsSplunkPut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Edition** | **string** | Edition of the Splunk Cloud instance | 
**Host** | **string** | The domain name for your Splunk Cloud instance. Don&#39;t include &#x60;http&#x60; or &#x60;https&#x60; in the string. For example: &#x60;acme.splunkcloud.com&#x60; | 

## Methods

### NewLogStreamSettingsSplunkPut

`func NewLogStreamSettingsSplunkPut(edition string, host string, ) *LogStreamSettingsSplunkPut`

NewLogStreamSettingsSplunkPut instantiates a new LogStreamSettingsSplunkPut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSettingsSplunkPutWithDefaults

`func NewLogStreamSettingsSplunkPutWithDefaults() *LogStreamSettingsSplunkPut`

NewLogStreamSettingsSplunkPutWithDefaults instantiates a new LogStreamSettingsSplunkPut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEdition

`func (o *LogStreamSettingsSplunkPut) GetEdition() string`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### GetEditionOk

`func (o *LogStreamSettingsSplunkPut) GetEditionOk() (*string, bool)`

GetEditionOk returns a tuple with the Edition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdition

`func (o *LogStreamSettingsSplunkPut) SetEdition(v string)`

SetEdition sets Edition field to given value.


### GetHost

`func (o *LogStreamSettingsSplunkPut) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogStreamSettingsSplunkPut) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LogStreamSettingsSplunkPut) SetHost(v string)`

SetHost sets Host field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


