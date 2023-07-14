# LogStreamSettingsAws

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Your AWS account ID | [optional] 
**EventSourceName** | Pointer to **string** | An alphanumeric name (no spaces) to identify this event source in AWS EventBridge | [optional] 
**Region** | Pointer to [**AwsRegion**](AwsRegion.md) |  | [optional] 

## Methods

### NewLogStreamSettingsAws

`func NewLogStreamSettingsAws() *LogStreamSettingsAws`

NewLogStreamSettingsAws instantiates a new LogStreamSettingsAws object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamSettingsAwsWithDefaults

`func NewLogStreamSettingsAwsWithDefaults() *LogStreamSettingsAws`

NewLogStreamSettingsAwsWithDefaults instantiates a new LogStreamSettingsAws object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LogStreamSettingsAws) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LogStreamSettingsAws) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LogStreamSettingsAws) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LogStreamSettingsAws) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEventSourceName

`func (o *LogStreamSettingsAws) GetEventSourceName() string`

GetEventSourceName returns the EventSourceName field if non-nil, zero value otherwise.

### GetEventSourceNameOk

`func (o *LogStreamSettingsAws) GetEventSourceNameOk() (*string, bool)`

GetEventSourceNameOk returns a tuple with the EventSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventSourceName

`func (o *LogStreamSettingsAws) SetEventSourceName(v string)`

SetEventSourceName sets EventSourceName field to given value.

### HasEventSourceName

`func (o *LogStreamSettingsAws) HasEventSourceName() bool`

HasEventSourceName returns a boolean if a field has been set.

### GetRegion

`func (o *LogStreamSettingsAws) GetRegion() AwsRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LogStreamSettingsAws) GetRegionOk() (*AwsRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LogStreamSettingsAws) SetRegion(v AwsRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LogStreamSettingsAws) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


