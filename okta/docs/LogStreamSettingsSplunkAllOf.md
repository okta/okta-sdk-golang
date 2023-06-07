# LogStreamSettingsSplunkAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The domain name for your Splunk Cloud instance. Don&#39;t include &#x60;http&#x60; or &#x60;https&#x60; in the string. For example: &#x60;acme.splunkcloud.com&#x60; | [optional] 
**Token** | Pointer to **string** | The HEC token for your Splunk Cloud HTTP Event Collector | [optional] 

## Methods

### NewLogStreamSettingsSplunkAllOf

`func NewLogStreamSettingsSplunkAllOf() *LogStreamSettingsSplunkAllOf`

NewLogStreamSettingsSplunkAllOf instantiates a new LogStreamSettingsSplunkAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSettingsSplunkAllOfWithDefaults

`func NewLogStreamSettingsSplunkAllOfWithDefaults() *LogStreamSettingsSplunkAllOf`

NewLogStreamSettingsSplunkAllOfWithDefaults instantiates a new LogStreamSettingsSplunkAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *LogStreamSettingsSplunkAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *LogStreamSettingsSplunkAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *LogStreamSettingsSplunkAllOf) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *LogStreamSettingsSplunkAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetToken

`func (o *LogStreamSettingsSplunkAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LogStreamSettingsSplunkAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LogStreamSettingsSplunkAllOf) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LogStreamSettingsSplunkAllOf) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


