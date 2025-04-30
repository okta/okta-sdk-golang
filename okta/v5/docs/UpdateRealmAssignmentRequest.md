# UpdateRealmAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**Actions**](Actions.md) |  | [optional] 
**Conditions** | Pointer to [**Conditions**](Conditions.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateRealmAssignmentRequest

`func NewUpdateRealmAssignmentRequest() *UpdateRealmAssignmentRequest`

NewUpdateRealmAssignmentRequest instantiates a new UpdateRealmAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRealmAssignmentRequestWithDefaults

`func NewUpdateRealmAssignmentRequestWithDefaults() *UpdateRealmAssignmentRequest`

NewUpdateRealmAssignmentRequestWithDefaults instantiates a new UpdateRealmAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *UpdateRealmAssignmentRequest) GetActions() Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UpdateRealmAssignmentRequest) GetActionsOk() (*Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UpdateRealmAssignmentRequest) SetActions(v Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UpdateRealmAssignmentRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *UpdateRealmAssignmentRequest) GetConditions() Conditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UpdateRealmAssignmentRequest) GetConditionsOk() (*Conditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UpdateRealmAssignmentRequest) SetConditions(v Conditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UpdateRealmAssignmentRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetName

`func (o *UpdateRealmAssignmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRealmAssignmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRealmAssignmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRealmAssignmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *UpdateRealmAssignmentRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *UpdateRealmAssignmentRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *UpdateRealmAssignmentRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *UpdateRealmAssignmentRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


