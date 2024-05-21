# EventHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | [**EventHookChannel**](EventHookChannel.md) |  | 
**Created** | Pointer to **time.Time** | Timestamp of the event hook creation | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The ID of the user who created the event hook | [optional] [readonly] 
**Description** | Pointer to **NullableString** | Description of the event hook | [optional] 
**Events** | [**EventSubscriptions**](EventSubscriptions.md) |  | 
**Id** | Pointer to **string** | Unique key for the event hook | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Date of the last event hook update | [optional] [readonly] 
**Name** | **string** | Display name for the event hook | 
**Status** | Pointer to **string** | Status of the event hook | [optional] [readonly] 
**VerificationStatus** | Pointer to **string** | Verification status of the event hook. &#x60;UNVERIFIED&#x60; event hooks won&#39;t receive any events. | [optional] [readonly] 
**Links** | Pointer to [**EventHookLinks**](EventHookLinks.md) |  | [optional] 

## Methods

### NewEventHook

`func NewEventHook(channel EventHookChannel, events EventSubscriptions, name string, ) *EventHook`

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

### GetDescription

`func (o *EventHook) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventHook) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventHook) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventHook) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EventHook) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EventHook) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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


### GetStatus

`func (o *EventHook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventHook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventHook) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventHook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerificationStatus

`func (o *EventHook) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *EventHook) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *EventHook) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *EventHook) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### GetLinks

`func (o *EventHook) GetLinks() EventHookLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EventHook) GetLinksOk() (*EventHookLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EventHook) SetLinks(v EventHookLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EventHook) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


