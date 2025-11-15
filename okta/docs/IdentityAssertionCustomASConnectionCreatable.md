# IdentityAssertionCustomASConnectionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationServer** | [**IdentityAssertionCustomASConnectionCreatableAuthorizationServer**](IdentityAssertionCustomASConnectionCreatableAuthorizationServer.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | Optional array of scopes | [optional] 

## Methods

### NewIdentityAssertionCustomASConnectionCreatable

`func NewIdentityAssertionCustomASConnectionCreatable(authorizationServer IdentityAssertionCustomASConnectionCreatableAuthorizationServer, connectionType string, ) *IdentityAssertionCustomASConnectionCreatable`

NewIdentityAssertionCustomASConnectionCreatable instantiates a new IdentityAssertionCustomASConnectionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionCustomASConnectionCreatableWithDefaults

`func NewIdentityAssertionCustomASConnectionCreatableWithDefaults() *IdentityAssertionCustomASConnectionCreatable`

NewIdentityAssertionCustomASConnectionCreatableWithDefaults instantiates a new IdentityAssertionCustomASConnectionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationServer

`func (o *IdentityAssertionCustomASConnectionCreatable) GetAuthorizationServer() IdentityAssertionCustomASConnectionCreatableAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *IdentityAssertionCustomASConnectionCreatable) GetAuthorizationServerOk() (*IdentityAssertionCustomASConnectionCreatableAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *IdentityAssertionCustomASConnectionCreatable) SetAuthorizationServer(v IdentityAssertionCustomASConnectionCreatableAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetConnectionType

`func (o *IdentityAssertionCustomASConnectionCreatable) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IdentityAssertionCustomASConnectionCreatable) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IdentityAssertionCustomASConnectionCreatable) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetProtocolType

`func (o *IdentityAssertionCustomASConnectionCreatable) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *IdentityAssertionCustomASConnectionCreatable) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *IdentityAssertionCustomASConnectionCreatable) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *IdentityAssertionCustomASConnectionCreatable) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetScopeCondition

`func (o *IdentityAssertionCustomASConnectionCreatable) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *IdentityAssertionCustomASConnectionCreatable) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *IdentityAssertionCustomASConnectionCreatable) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *IdentityAssertionCustomASConnectionCreatable) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *IdentityAssertionCustomASConnectionCreatable) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityAssertionCustomASConnectionCreatable) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityAssertionCustomASConnectionCreatable) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IdentityAssertionCustomASConnectionCreatable) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


