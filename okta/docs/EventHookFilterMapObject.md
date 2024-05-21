# EventHookFilterMapObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**EventHookFilterMapObjectCondition**](EventHookFilterMapObjectCondition.md) |  | [optional] 
**Event** | Pointer to **string** | The filtered event type | [optional] 

## Methods

### NewEventHookFilterMapObject

`func NewEventHookFilterMapObject() *EventHookFilterMapObject`

NewEventHookFilterMapObject instantiates a new EventHookFilterMapObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHookFilterMapObjectWithDefaults

`func NewEventHookFilterMapObjectWithDefaults() *EventHookFilterMapObject`

NewEventHookFilterMapObjectWithDefaults instantiates a new EventHookFilterMapObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *EventHookFilterMapObject) GetCondition() EventHookFilterMapObjectCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *EventHookFilterMapObject) GetConditionOk() (*EventHookFilterMapObjectCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *EventHookFilterMapObject) SetCondition(v EventHookFilterMapObjectCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *EventHookFilterMapObject) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEvent

`func (o *EventHookFilterMapObject) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *EventHookFilterMapObject) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *EventHookFilterMapObject) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *EventHookFilterMapObject) HasEvent() bool`

HasEvent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


