# EventSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NullableEventHookFilters**](EventHookFilters.md) |  | [optional] 
**Items** | **[]string** | The subscribed event types that trigger the event hook. When you register an event hook you need to specify which events you want to subscribe to. To see the list of event types currently eligible for use in event hooks, use the [Event Types catalog](/docs/reference/api/event-types/#catalog) and search with the parameter &#x60;event-hook-eligible&#x60;. | 
**Type** | **string** | The events object type. Currently supports &#x60;EVENT_TYPE&#x60;. | 

## Methods

### NewEventSubscriptions

`func NewEventSubscriptions(items []string, type_ string, ) *EventSubscriptions`

NewEventSubscriptions instantiates a new EventSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionsWithDefaults

`func NewEventSubscriptionsWithDefaults() *EventSubscriptions`

NewEventSubscriptionsWithDefaults instantiates a new EventSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *EventSubscriptions) GetFilter() EventHookFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *EventSubscriptions) GetFilterOk() (*EventHookFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *EventSubscriptions) SetFilter(v EventHookFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *EventSubscriptions) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *EventSubscriptions) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *EventSubscriptions) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetItems

`func (o *EventSubscriptions) GetItems() []string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EventSubscriptions) GetItemsOk() (*[]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EventSubscriptions) SetItems(v []string)`

SetItems sets Items field to given value.


### GetType

`func (o *EventSubscriptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSubscriptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSubscriptions) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


