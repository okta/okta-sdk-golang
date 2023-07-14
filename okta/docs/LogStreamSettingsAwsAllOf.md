# LogStreamSettingsAwsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Your AWS account ID | [optional] 
**EventSourceName** | Pointer to **string** | An alphanumeric name (no spaces) to identify this event source in AWS EventBridge | [optional] 
**Region** | Pointer to [**AwsRegion**](AwsRegion.md) |  | [optional] 

## Methods

### NewLogStreamSettingsAwsAllOf

`func NewLogStreamSettingsAwsAllOf() *LogStreamSettingsAwsAllOf`

NewLogStreamSettingsAwsAllOf instantiates a new LogStreamSettingsAwsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSettingsAwsAllOfWithDefaults

`func NewLogStreamSettingsAwsAllOfWithDefaults() *LogStreamSettingsAwsAllOf`

NewLogStreamSettingsAwsAllOfWithDefaults instantiates a new LogStreamSettingsAwsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LogStreamSettingsAwsAllOf) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LogStreamSettingsAwsAllOf) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LogStreamSettingsAwsAllOf) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LogStreamSettingsAwsAllOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEventSourceName

`func (o *LogStreamSettingsAwsAllOf) GetEventSourceName() string`

GetEventSourceName returns the EventSourceName field if non-nil, zero value otherwise.

### GetEventSourceNameOk

`func (o *LogStreamSettingsAwsAllOf) GetEventSourceNameOk() (*string, bool)`

GetEventSourceNameOk returns a tuple with the EventSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceName

`func (o *LogStreamSettingsAwsAllOf) SetEventSourceName(v string)`

SetEventSourceName sets EventSourceName field to given value.

### HasEventSourceName

`func (o *LogStreamSettingsAwsAllOf) HasEventSourceName() bool`

HasEventSourceName returns a boolean if a field has been set.

### GetRegion

`func (o *LogStreamSettingsAwsAllOf) GetRegion() AwsRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogStreamSettingsAwsAllOf) GetRegionOk() (*AwsRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogStreamSettingsAwsAllOf) SetRegion(v AwsRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogStreamSettingsAwsAllOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


