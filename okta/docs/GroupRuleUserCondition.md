# GroupRuleUserCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to **[]string** | Excluded &#x60;userIds&#x60; when processing rules | [optional] 

## Methods

### NewGroupRuleUserCondition

`func NewGroupRuleUserCondition() *GroupRuleUserCondition`

NewGroupRuleUserCondition instantiates a new GroupRuleUserCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupRuleUserConditionWithDefaults

`func NewGroupRuleUserConditionWithDefaults() *GroupRuleUserCondition`

NewGroupRuleUserConditionWithDefaults instantiates a new GroupRuleUserCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *GroupRuleUserCondition) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *GroupRuleUserCondition) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *GroupRuleUserCondition) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *GroupRuleUserCondition) HasExclude() bool`

HasExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


