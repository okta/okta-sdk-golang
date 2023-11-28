# RealmAssignmentRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**Actions**](Actions.md) |  | [optional] 
**Conditions** | Pointer to [**Conditions**](Conditions.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsDefault** | Pointer to **bool** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewRealmAssignmentRule

`func NewRealmAssignmentRule() *RealmAssignmentRule`

NewRealmAssignmentRule instantiates a new RealmAssignmentRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmAssignmentRuleWithDefaults

`func NewRealmAssignmentRuleWithDefaults() *RealmAssignmentRule`

NewRealmAssignmentRuleWithDefaults instantiates a new RealmAssignmentRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *RealmAssignmentRule) GetActions() Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RealmAssignmentRule) GetActionsOk() (*Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RealmAssignmentRule) SetActions(v Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *RealmAssignmentRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *RealmAssignmentRule) GetConditions() Conditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *RealmAssignmentRule) GetConditionsOk() (*Conditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *RealmAssignmentRule) SetConditions(v Conditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *RealmAssignmentRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreated

`func (o *RealmAssignmentRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RealmAssignmentRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RealmAssignmentRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *RealmAssignmentRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *RealmAssignmentRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RealmAssignmentRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RealmAssignmentRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RealmAssignmentRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *RealmAssignmentRule) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RealmAssignmentRule) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RealmAssignmentRule) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *RealmAssignmentRule) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetLastUpdated

`func (o *RealmAssignmentRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RealmAssignmentRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RealmAssignmentRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *RealmAssignmentRule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *RealmAssignmentRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RealmAssignmentRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RealmAssignmentRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RealmAssignmentRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *RealmAssignmentRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RealmAssignmentRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RealmAssignmentRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RealmAssignmentRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *RealmAssignmentRule) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RealmAssignmentRule) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RealmAssignmentRule) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RealmAssignmentRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *RealmAssignmentRule) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RealmAssignmentRule) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RealmAssignmentRule) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RealmAssignmentRule) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


