# UserTypeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | **[]string** | The user types to exclude | 
**Include** | **[]string** | The user types to include | 

## Methods

### NewUserTypeCondition

`func NewUserTypeCondition(exclude []string, include []string, ) *UserTypeCondition`

NewUserTypeCondition instantiates a new UserTypeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTypeConditionWithDefaults

`func NewUserTypeConditionWithDefaults() *UserTypeCondition`

NewUserTypeConditionWithDefaults instantiates a new UserTypeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *UserTypeCondition) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *UserTypeCondition) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *UserTypeCondition) SetExclude(v []string)`

SetExclude sets Exclude field to given value.


### GetInclude

`func (o *UserTypeCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *UserTypeCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *UserTypeCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


