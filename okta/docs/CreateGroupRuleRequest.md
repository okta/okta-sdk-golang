# CreateGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**GroupRuleAction**](GroupRuleAction.md) |  | [optional] 
**Conditions** | Pointer to [**GroupRuleConditions**](GroupRuleConditions.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the group rule | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateGroupRuleRequest

`func NewCreateGroupRuleRequest() *CreateGroupRuleRequest`

NewCreateGroupRuleRequest instantiates a new CreateGroupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupRuleRequestWithDefaults

`func NewCreateGroupRuleRequestWithDefaults() *CreateGroupRuleRequest`

NewCreateGroupRuleRequestWithDefaults instantiates a new CreateGroupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *CreateGroupRuleRequest) GetActions() GroupRuleAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *CreateGroupRuleRequest) GetActionsOk() (*GroupRuleAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *CreateGroupRuleRequest) SetActions(v GroupRuleAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *CreateGroupRuleRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *CreateGroupRuleRequest) GetConditions() GroupRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateGroupRuleRequest) GetConditionsOk() (*GroupRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateGroupRuleRequest) SetConditions(v GroupRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreateGroupRuleRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetName

`func (o *CreateGroupRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGroupRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGroupRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateGroupRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateGroupRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGroupRuleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGroupRuleRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateGroupRuleRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


