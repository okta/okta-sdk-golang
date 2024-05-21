# BeforeScheduledActionPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**LifecycleAction** | Pointer to [**ScheduledUserLifecycleAction**](ScheduledUserLifecycleAction.md) |  | [optional] 

## Methods

### NewBeforeScheduledActionPolicyRuleCondition

`func NewBeforeScheduledActionPolicyRuleCondition() *BeforeScheduledActionPolicyRuleCondition`

NewBeforeScheduledActionPolicyRuleCondition instantiates a new BeforeScheduledActionPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeforeScheduledActionPolicyRuleConditionWithDefaults

`func NewBeforeScheduledActionPolicyRuleConditionWithDefaults() *BeforeScheduledActionPolicyRuleCondition`

NewBeforeScheduledActionPolicyRuleConditionWithDefaults instantiates a new BeforeScheduledActionPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *BeforeScheduledActionPolicyRuleCondition) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BeforeScheduledActionPolicyRuleCondition) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BeforeScheduledActionPolicyRuleCondition) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BeforeScheduledActionPolicyRuleCondition) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLifecycleAction

`func (o *BeforeScheduledActionPolicyRuleCondition) GetLifecycleAction() ScheduledUserLifecycleAction`

GetLifecycleAction returns the LifecycleAction field if non-nil, zero value otherwise.

### GetLifecycleActionOk

`func (o *BeforeScheduledActionPolicyRuleCondition) GetLifecycleActionOk() (*ScheduledUserLifecycleAction, bool)`

GetLifecycleActionOk returns a tuple with the LifecycleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifecycleAction

`func (o *BeforeScheduledActionPolicyRuleCondition) SetLifecycleAction(v ScheduledUserLifecycleAction)`

SetLifecycleAction sets LifecycleAction field to given value.

### HasLifecycleAction

`func (o *BeforeScheduledActionPolicyRuleCondition) HasLifecycleAction() bool`

HasLifecycleAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


