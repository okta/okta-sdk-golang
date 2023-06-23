# LogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpChain** | Pointer to [**[]LogIpAddress**](LogIpAddress.md) |  | [optional] [readonly] 

## Methods

### NewLogRequest

`func NewLogRequest() *LogRequest`

NewLogRequest instantiates a new LogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogRequestWithDefaults

`func NewLogRequestWithDefaults() *LogRequest`

NewLogRequestWithDefaults instantiates a new LogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpChain

`func (o *LogRequest) GetIpChain() []LogIpAddress`

GetIpChain returns the IpChain field if non-nil, zero value otherwise.

### GetIpChainOk

`func (o *LogRequest) GetIpChainOk() (*[]LogIpAddress, bool)`

GetIpChainOk returns a tuple with the IpChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpChain

`func (o *LogRequest) SetIpChain(v []LogIpAddress)`

SetIpChain sets IpChain field to given value.

### HasIpChain

`func (o *LogRequest) HasIpChain() bool`

HasIpChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


