# ApplicationLayoutRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to [**ApplicationLayoutRuleCondition**](ApplicationLayoutRuleCondition.md) |  | [optional] 

## Methods

### NewApplicationLayoutRule

`func NewApplicationLayoutRule() *ApplicationLayoutRule`

NewApplicationLayoutRule instantiates a new ApplicationLayoutRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationLayoutRuleWithDefaults

`func NewApplicationLayoutRuleWithDefaults() *ApplicationLayoutRule`

NewApplicationLayoutRuleWithDefaults instantiates a new ApplicationLayoutRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *ApplicationLayoutRule) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *ApplicationLayoutRule) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *ApplicationLayoutRule) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *ApplicationLayoutRule) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetCondition

`func (o *ApplicationLayoutRule) GetCondition() ApplicationLayoutRuleCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ApplicationLayoutRule) GetConditionOk() (*ApplicationLayoutRuleCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ApplicationLayoutRule) SetCondition(v ApplicationLayoutRuleCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ApplicationLayoutRule) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


