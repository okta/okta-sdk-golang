# AuthorizationServerPolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**AuthorizationServerPolicyRuleActions**](AuthorizationServerPolicyRuleActions.md) |  | [optional] 
**Conditions** | Pointer to [**AuthorizationServerPolicyRuleConditions**](AuthorizationServerPolicyRuleConditions.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the rule was created | [optional] [readonly] 
**Id** | Pointer to **string** | Identifier of the rule | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the rule was last modified | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the rule | [optional] 
**Priority** | Pointer to **int32** | Priority of the rule | [optional] 
**Status** | Pointer to **string** | Status of the rule | [optional] 
**System** | Pointer to **bool** | Set to &#x60;true&#x60; for system rules. You can&#39;t delete system rules. | [optional] 
**Type** | Pointer to **string** | Rule type | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewAuthorizationServerPolicyRule

`func NewAuthorizationServerPolicyRule() *AuthorizationServerPolicyRule`

NewAuthorizationServerPolicyRule instantiates a new AuthorizationServerPolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerPolicyRuleWithDefaults

`func NewAuthorizationServerPolicyRuleWithDefaults() *AuthorizationServerPolicyRule`

NewAuthorizationServerPolicyRuleWithDefaults instantiates a new AuthorizationServerPolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *AuthorizationServerPolicyRule) GetActions() AuthorizationServerPolicyRuleActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AuthorizationServerPolicyRule) GetActionsOk() (*AuthorizationServerPolicyRuleActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AuthorizationServerPolicyRule) SetActions(v AuthorizationServerPolicyRuleActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AuthorizationServerPolicyRule) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConditions

`func (o *AuthorizationServerPolicyRule) GetConditions() AuthorizationServerPolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthorizationServerPolicyRule) GetConditionsOk() (*AuthorizationServerPolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthorizationServerPolicyRule) SetConditions(v AuthorizationServerPolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthorizationServerPolicyRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreated

`func (o *AuthorizationServerPolicyRule) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuthorizationServerPolicyRule) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuthorizationServerPolicyRule) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AuthorizationServerPolicyRule) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *AuthorizationServerPolicyRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthorizationServerPolicyRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthorizationServerPolicyRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthorizationServerPolicyRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AuthorizationServerPolicyRule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AuthorizationServerPolicyRule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AuthorizationServerPolicyRule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AuthorizationServerPolicyRule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AuthorizationServerPolicyRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizationServerPolicyRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizationServerPolicyRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizationServerPolicyRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *AuthorizationServerPolicyRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AuthorizationServerPolicyRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AuthorizationServerPolicyRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *AuthorizationServerPolicyRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorizationServerPolicyRule) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorizationServerPolicyRule) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorizationServerPolicyRule) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorizationServerPolicyRule) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *AuthorizationServerPolicyRule) GetSystem() bool`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *AuthorizationServerPolicyRule) GetSystemOk() (*bool, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *AuthorizationServerPolicyRule) SetSystem(v bool)`

SetSystem sets System field to given value.

### HasSystem

`func (o *AuthorizationServerPolicyRule) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *AuthorizationServerPolicyRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizationServerPolicyRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizationServerPolicyRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthorizationServerPolicyRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *AuthorizationServerPolicyRule) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizationServerPolicyRule) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizationServerPolicyRule) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizationServerPolicyRule) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


