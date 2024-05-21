# EventHookFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventFilterMap** | Pointer to [**[]EventHookFilterMapObject**](EventHookFilterMapObject.md) | The object that maps the filter to the event type | [optional] 
**Type** | Pointer to **string** | The type of filter. Currently only supports &#x60;EXPRESSION_LANGUAGE&#x60; | [optional] [readonly] 

## Methods

### NewEventHookFilters

`func NewEventHookFilters() *EventHookFilters`

NewEventHookFilters instantiates a new EventHookFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHookFiltersWithDefaults

`func NewEventHookFiltersWithDefaults() *EventHookFilters`

NewEventHookFiltersWithDefaults instantiates a new EventHookFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventFilterMap

`func (o *EventHookFilters) GetEventFilterMap() []EventHookFilterMapObject`

GetEventFilterMap returns the EventFilterMap field if non-nil, zero value otherwise.

### GetEventFilterMapOk

`func (o *EventHookFilters) GetEventFilterMapOk() (*[]EventHookFilterMapObject, bool)`

GetEventFilterMapOk returns a tuple with the EventFilterMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventFilterMap

`func (o *EventHookFilters) SetEventFilterMap(v []EventHookFilterMapObject)`

SetEventFilterMap sets EventFilterMap field to given value.

### HasEventFilterMap

`func (o *EventHookFilters) HasEventFilterMap() bool`

HasEventFilterMap returns a boolean if a field has been set.

### GetType

`func (o *EventHookFilters) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventHookFilters) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventHookFilters) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventHookFilters) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


