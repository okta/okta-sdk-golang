# GroupRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**GroupRuleAction**](GroupRuleAction.md) |  | [optional] 
**Conditions** | Pointer to [**GroupRuleConditions**](GroupRuleConditions.md) |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**GroupRuleStatus**](GroupRuleStatus.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewGroupRule

`func NewGroupRule() *GroupRule`

NewGroupRule instantiates a new GroupRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRuleWithDefaults

`func NewGroupRuleWithDefaults() *GroupRule`

NewGroupRuleWithDefaults instantiates a new GroupRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *GroupRule) GetActions() GroupRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GroupRule) GetActionsOk() (*GroupRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GroupRule) SetActions(v GroupRuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GroupRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *GroupRule) GetConditions() GroupRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *GroupRule) GetConditionsOk() (*GroupRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *GroupRule) SetConditions(v GroupRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *GroupRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreated

`func (o *GroupRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GroupRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GroupRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GroupRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GroupRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GroupRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GroupRule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *GroupRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GroupRule) GetStatus() GroupRuleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupRule) GetStatusOk() (*GroupRuleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupRule) SetStatus(v GroupRuleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GroupRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GroupRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GroupRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GroupRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GroupRule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


