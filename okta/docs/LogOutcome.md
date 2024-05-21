# LogOutcome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] [readonly] 
**Result** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewLogOutcome

`func NewLogOutcome() *LogOutcome`

NewLogOutcome instantiates a new LogOutcome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogOutcomeWithDefaults

`func NewLogOutcomeWithDefaults() *LogOutcome`

NewLogOutcomeWithDefaults instantiates a new LogOutcome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *LogOutcome) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LogOutcome) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LogOutcome) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *LogOutcome) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetResult

`func (o *LogOutcome) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LogOutcome) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LogOutcome) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *LogOutcome) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


