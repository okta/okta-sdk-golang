# StreamStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The status of the SSF Stream configuration | [optional] 
**StreamId** | Pointer to **string** | The ID of the SSF Stream configuration. This corresponds to the value in the query parameter of the request. | [optional] 

## Methods

### NewStreamStatus

`func NewStreamStatus() *StreamStatus`

NewStreamStatus instantiates a new StreamStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamStatusWithDefaults

`func NewStreamStatusWithDefaults() *StreamStatus`

NewStreamStatusWithDefaults instantiates a new StreamStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *StreamStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StreamStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StreamStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StreamStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStreamId

`func (o *StreamStatus) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StreamStatus) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StreamStatus) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *StreamStatus) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


