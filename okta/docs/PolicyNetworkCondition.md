# PolicyNetworkCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connection** | Pointer to [**PolicyNetworkConnection**](PolicyNetworkConnection.md) |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyNetworkCondition

`func NewPolicyNetworkCondition() *PolicyNetworkCondition`

NewPolicyNetworkCondition instantiates a new PolicyNetworkCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyNetworkConditionWithDefaults

`func NewPolicyNetworkConditionWithDefaults() *PolicyNetworkCondition`

NewPolicyNetworkConditionWithDefaults instantiates a new PolicyNetworkCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnection

`func (o *PolicyNetworkCondition) GetConnection() PolicyNetworkConnection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *PolicyNetworkCondition) GetConnectionOk() (*PolicyNetworkConnection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *PolicyNetworkCondition) SetConnection(v PolicyNetworkConnection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *PolicyNetworkCondition) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetExclude

`func (o *PolicyNetworkCondition) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *PolicyNetworkCondition) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *PolicyNetworkCondition) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *PolicyNetworkCondition) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### GetInclude

`func (o *PolicyNetworkCondition) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *PolicyNetworkCondition) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *PolicyNetworkCondition) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *PolicyNetworkCondition) HasInclude() bool`

HasInclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


