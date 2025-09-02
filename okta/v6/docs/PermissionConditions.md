# PermissionConditions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | Pointer to **map[string]map[string]interface{}** | Exclude attributes with specific values for the permission | [optional] 
**Include** | Pointer to **map[string]map[string]interface{}** | Include attributes with specific values for the permission | [optional] 

## Methods

### NewPermissionConditions

`func NewPermissionConditions() *PermissionConditions`

NewPermissionConditions instantiates a new PermissionConditions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionConditionsWithDefaults

`func NewPermissionConditionsWithDefaults() *PermissionConditions`

NewPermissionConditionsWithDefaults instantiates a new PermissionConditions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *PermissionConditions) GetExclude() map[string]map[string]interface{}`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PermissionConditions) GetExcludeOk() (*map[string]map[string]interface{}, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PermissionConditions) SetExclude(v map[string]map[string]interface{})`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PermissionConditions) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### SetExcludeNil

`func (o *PermissionConditions) SetExcludeNil(b bool)`

 SetExcludeNil sets the value for Exclude to be an explicit nil

### UnsetExclude
`func (o *PermissionConditions) UnsetExclude()`

UnsetExclude ensures that no value is present for Exclude, not even an explicit nil
### GetInclude

`func (o *PermissionConditions) GetInclude() map[string]map[string]interface{}`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *PermissionConditions) GetIncludeOk() (*map[string]map[string]interface{}, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *PermissionConditions) SetInclude(v map[string]map[string]interface{})`

SetInclude sets Include field to given value.

### HasInclude

`func (o *PermissionConditions) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### SetIncludeNil

`func (o *PermissionConditions) SetIncludeNil(b bool)`

 SetIncludeNil sets the value for Include to be an explicit nil

### UnsetInclude
`func (o *PermissionConditions) UnsetInclude()`

UnsetInclude ensures that no value is present for Include, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


