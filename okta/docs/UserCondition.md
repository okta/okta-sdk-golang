# UserCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to **[]string** |  | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserCondition

`func NewUserCondition() *UserCondition`

NewUserCondition instantiates a new UserCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConditionWithDefaults

`func NewUserConditionWithDefaults() *UserCondition`

NewUserConditionWithDefaults instantiates a new UserCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *UserCondition) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *UserCondition) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *UserCondition) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *UserCondition) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *UserCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *UserCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *UserCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *UserCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


