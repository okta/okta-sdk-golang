# ManagedConnectionPatchable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | Optional array of scopes. Required when scopeCondition is INCLUDE_ONLY or EXCLUDE. Not used when scopeCondition is ALL_SCOPES. | [optional] 

## Methods

### NewManagedConnectionPatchable

`func NewManagedConnectionPatchable() *ManagedConnectionPatchable`

NewManagedConnectionPatchable instantiates a new ManagedConnectionPatchable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionPatchableWithDefaults

`func NewManagedConnectionPatchableWithDefaults() *ManagedConnectionPatchable`

NewManagedConnectionPatchableWithDefaults instantiates a new ManagedConnectionPatchable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeCondition

`func (o *ManagedConnectionPatchable) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *ManagedConnectionPatchable) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *ManagedConnectionPatchable) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *ManagedConnectionPatchable) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *ManagedConnectionPatchable) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ManagedConnectionPatchable) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ManagedConnectionPatchable) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ManagedConnectionPatchable) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


