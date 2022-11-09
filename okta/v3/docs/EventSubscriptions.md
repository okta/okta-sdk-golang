# EventSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**EventSubscriptionType**](EventSubscriptionType.md) |  | [optional] 

## Methods

### NewEventSubscriptions

`func NewEventSubscriptions() *EventSubscriptions`

NewEventSubscriptions instantiates a new EventSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventSubscriptionsWithDefaults

`func NewEventSubscriptionsWithDefaults() *EventSubscriptions`

NewEventSubscriptionsWithDefaults instantiates a new EventSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### HasItems

`func (o *EventSubscriptions) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetType

`func (o *EventSubscriptions) GetType() EventSubscriptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventSubscriptions) GetTypeOk() (*EventSubscriptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventSubscriptions) SetType(v EventSubscriptionType)`

SetType sets Type field to given value.

### HasType

`func (o *EventSubscriptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


