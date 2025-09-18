# StreamVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | An arbitrary string that Okta as a transmitter must echo back to the Event Receiver in the Verification Event&#39;s payload | [optional] 
**StreamId** | **string** | The ID of the SSF Stream Configuration | 

## Methods

### NewStreamVerificationRequest

`func NewStreamVerificationRequest(streamId string, ) *StreamVerificationRequest`

NewStreamVerificationRequest instantiates a new StreamVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamVerificationRequestWithDefaults

`func NewStreamVerificationRequestWithDefaults() *StreamVerificationRequest`

NewStreamVerificationRequestWithDefaults instantiates a new StreamVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *StreamVerificationRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamVerificationRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamVerificationRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamVerificationRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStreamId

`func (o *StreamVerificationRequest) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StreamVerificationRequest) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StreamVerificationRequest) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


