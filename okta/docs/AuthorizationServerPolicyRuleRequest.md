# AuthorizationServerPolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AuthorizationServerPolicyRuleActions**](AuthorizationServerPolicyRuleActions.md) |  | [optional] 
**Conditions** | [**AuthorizationServerPolicyRuleConditions**](AuthorizationServerPolicyRuleConditions.md) |  | 
**Created** | Pointer to **time.Time** | Timestamp when the rule was created | [optional] [readonly] 
**Id** | Pointer to **string** | Identifier of the rule | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the rule was last modified | [optional] [readonly] 
**Name** | **string** | Name of the rule | 
**Priority** | Pointer to **int32** | Priority of the rule | [optional] 
**Status** | Pointer to **string** | Status of the rule | [optional] 
**System** | Pointer to **bool** | Set to &#x60;true&#x60; for system rules. You can&#39;t delete system rules. | [optional] 
**Type** | **string** | Rule type | 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRuleRequest

`func NewAuthorizationServerPolicyRuleRequest(conditions AuthorizationServerPolicyRuleConditions, name string, type_ string, ) *AuthorizationServerPolicyRuleRequest`

NewAuthorizationServerPolicyRuleRequest instantiates a new AuthorizationServerPolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleRequestWithDefaults

`func NewAuthorizationServerPolicyRuleRequestWithDefaults() *AuthorizationServerPolicyRuleRequest`

NewAuthorizationServerPolicyRuleRequestWithDefaults instantiates a new AuthorizationServerPolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuthorizationServerPolicyRuleRequest) GetActions() AuthorizationServerPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthorizationServerPolicyRuleRequest) GetActionsOk() (*AuthorizationServerPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthorizationServerPolicyRuleRequest) SetActions(v AuthorizationServerPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthorizationServerPolicyRuleRequest) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *AuthorizationServerPolicyRuleRequest) GetConditions() AuthorizationServerPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizationServerPolicyRuleRequest) GetConditionsOk() (*AuthorizationServerPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizationServerPolicyRuleRequest) SetConditions(v AuthorizationServerPolicyRuleConditions)`

SetConditions sets Conditions field to given value.


### GetCreated

`func (o *AuthorizationServerPolicyRuleRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorizationServerPolicyRuleRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorizationServerPolicyRuleRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorizationServerPolicyRuleRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AuthorizationServerPolicyRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationServerPolicyRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationServerPolicyRuleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationServerPolicyRuleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthorizationServerPolicyRuleRequest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthorizationServerPolicyRuleRequest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthorizationServerPolicyRuleRequest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthorizationServerPolicyRuleRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AuthorizationServerPolicyRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizationServerPolicyRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizationServerPolicyRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPriority

`func (o *AuthorizationServerPolicyRuleRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AuthorizationServerPolicyRuleRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AuthorizationServerPolicyRuleRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AuthorizationServerPolicyRuleRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizationServerPolicyRuleRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationServerPolicyRuleRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationServerPolicyRuleRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationServerPolicyRuleRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *AuthorizationServerPolicyRuleRequest) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AuthorizationServerPolicyRuleRequest) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AuthorizationServerPolicyRuleRequest) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AuthorizationServerPolicyRuleRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *AuthorizationServerPolicyRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationServerPolicyRuleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationServerPolicyRuleRequest) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *AuthorizationServerPolicyRuleRequest) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizationServerPolicyRuleRequest) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizationServerPolicyRuleRequest) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizationServerPolicyRuleRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


