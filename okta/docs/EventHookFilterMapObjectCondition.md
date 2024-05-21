# EventHookFilterMapObjectCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | The Okta Expression language statement that filters the event type | [optional] 
**Version** | Pointer to **NullableString** | Internal field | [optional] [readonly] 

## Methods

### NewEventHookFilterMapObjectCondition

`func NewEventHookFilterMapObjectCondition() *EventHookFilterMapObjectCondition`

NewEventHookFilterMapObjectCondition instantiates a new EventHookFilterMapObjectCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHookFilterMapObjectConditionWithDefaults

`func NewEventHookFilterMapObjectConditionWithDefaults() *EventHookFilterMapObjectCondition`

NewEventHookFilterMapObjectConditionWithDefaults instantiates a new EventHookFilterMapObjectCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *EventHookFilterMapObjectCondition) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *EventHookFilterMapObjectCondition) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *EventHookFilterMapObjectCondition) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *EventHookFilterMapObjectCondition) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetVersion

`func (o *EventHookFilterMapObjectCondition) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EventHookFilterMapObjectCondition) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EventHookFilterMapObjectCondition) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EventHookFilterMapObjectCondition) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *EventHookFilterMapObjectCondition) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *EventHookFilterMapObjectCondition) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


