# EventHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**EventHookChannel**](EventHookChannel.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Events** | Pointer to [**EventSubscriptions**](EventSubscriptions.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**VerificationStatus** | Pointer to [**EventHookVerificationStatus**](EventHookVerificationStatus.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewEventHook

`func NewEventHook() *EventHook`

NewEventHook instantiates a new EventHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventHookWithDefaults

`func NewEventHookWithDefaults() *EventHook`

NewEventHookWithDefaults instantiates a new EventHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *EventHook) GetChannel() EventHookChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *EventHook) GetChannelOk() (*EventHookChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *EventHook) SetChannel(v EventHookChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *EventHook) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreated

`func (o *EventHook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventHook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EventHook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EventHook) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EventHook) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EventHook) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EventHook) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EventHook) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *EventHook) GetEvents() EventSubscriptions`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventHook) GetEventsOk() (*EventSubscriptions, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventHook) SetEvents(v EventSubscriptions)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *EventHook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetId

`func (o *EventHook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventHook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventHook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EventHook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EventHook) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EventHook) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EventHook) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EventHook) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *EventHook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventHook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventHook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EventHook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *EventHook) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventHook) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventHook) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventHook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *EventHook) GetVerificationStatus() EventHookVerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *EventHook) GetVerificationStatusOk() (*EventHookVerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *EventHook) SetVerificationStatus(v EventHookVerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *EventHook) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetLinks

`func (o *EventHook) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventHook) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventHook) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventHook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


