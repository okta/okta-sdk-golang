# LogStreamSettingsSplunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The domain name for your Splunk Cloud instance. Don&#39;t include &#x60;http&#x60; or &#x60;https&#x60; in the string. For example: &#x60;acme.splunkcloud.com&#x60; | [optional] 
**Token** | Pointer to **string** | The HEC token for your Splunk Cloud HTTP Event Collector | [optional] 

## Methods

### NewLogStreamSettingsSplunk

`func NewLogStreamSettingsSplunk() *LogStreamSettingsSplunk`

NewLogStreamSettingsSplunk instantiates a new LogStreamSettingsSplunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSettingsSplunkWithDefaults

`func NewLogStreamSettingsSplunkWithDefaults() *LogStreamSettingsSplunk`

NewLogStreamSettingsSplunkWithDefaults instantiates a new LogStreamSettingsSplunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *LogStreamSettingsSplunk) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogStreamSettingsSplunk) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LogStreamSettingsSplunk) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LogStreamSettingsSplunk) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetToken

`func (o *LogStreamSettingsSplunk) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LogStreamSettingsSplunk) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LogStreamSettingsSplunk) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LogStreamSettingsSplunk) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


