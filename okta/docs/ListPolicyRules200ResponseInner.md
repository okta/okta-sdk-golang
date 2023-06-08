# ListPolicyRules200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**System** | Pointer to **bool** |  | [optional] [default to false]
**Type** | Pointer to [**PolicyRuleType**](PolicyRuleType.md) |  | [optional] 
**Actions** | Pointer to [**OktaSignOnPolicyRuleActions**](OktaSignOnPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**OktaSignOnPolicyRuleConditions**](OktaSignOnPolicyRuleConditions.md) |  | [optional] 

## Methods

### NewListPolicyRules200ResponseInner

`func NewListPolicyRules200ResponseInner() *ListPolicyRules200ResponseInner`

NewListPolicyRules200ResponseInner instantiates a new ListPolicyRules200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicyRules200ResponseInnerWithDefaults

`func NewListPolicyRules200ResponseInnerWithDefaults() *ListPolicyRules200ResponseInner`

NewListPolicyRules200ResponseInnerWithDefaults instantiates a new ListPolicyRules200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListPolicyRules200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListPolicyRules200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListPolicyRules200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListPolicyRules200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreatedNil

`func (o *ListPolicyRules200ResponseInner) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ListPolicyRules200ResponseInner) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetId

`func (o *ListPolicyRules200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPolicyRules200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPolicyRules200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListPolicyRules200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListPolicyRules200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListPolicyRules200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListPolicyRules200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListPolicyRules200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *ListPolicyRules200ResponseInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ListPolicyRules200ResponseInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetName

`func (o *ListPolicyRules200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPolicyRules200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPolicyRules200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPolicyRules200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *ListPolicyRules200ResponseInner) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ListPolicyRules200ResponseInner) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ListPolicyRules200ResponseInner) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ListPolicyRules200ResponseInner) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *ListPolicyRules200ResponseInner) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPolicyRules200ResponseInner) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPolicyRules200ResponseInner) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListPolicyRules200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *ListPolicyRules200ResponseInner) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ListPolicyRules200ResponseInner) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ListPolicyRules200ResponseInner) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ListPolicyRules200ResponseInner) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *ListPolicyRules200ResponseInner) GetType() PolicyRuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListPolicyRules200ResponseInner) GetTypeOk() (*PolicyRuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListPolicyRules200ResponseInner) SetType(v PolicyRuleType)`

SetType sets Type field to given value.

### HasType

`func (o *ListPolicyRules200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetActions

`func (o *ListPolicyRules200ResponseInner) GetActions() OktaSignOnPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ListPolicyRules200ResponseInner) GetActionsOk() (*OktaSignOnPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ListPolicyRules200ResponseInner) SetActions(v OktaSignOnPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ListPolicyRules200ResponseInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *ListPolicyRules200ResponseInner) GetConditions() OktaSignOnPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ListPolicyRules200ResponseInner) GetConditionsOk() (*OktaSignOnPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ListPolicyRules200ResponseInner) SetConditions(v OktaSignOnPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ListPolicyRules200ResponseInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


