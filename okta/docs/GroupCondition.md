# GroupCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | **[]string** | Groups to be excluded | 
**Include** | **[]string** | Groups to be included | 

## Methods

### NewGroupCondition

`func NewGroupCondition(exclude []string, include []string, ) *GroupCondition`

NewGroupCondition instantiates a new GroupCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupConditionWithDefaults

`func NewGroupConditionWithDefaults() *GroupCondition`

NewGroupConditionWithDefaults instantiates a new GroupCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *GroupCondition) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *GroupCondition) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *GroupCondition) SetExclude(v []string)`

SetExclude sets Exclude field to given value.


### GetInclude

`func (o *GroupCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *GroupCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *GroupCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


