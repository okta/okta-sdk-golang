# IdentityAssertionCustomASConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationServer** | [**CustomAuthorizationServer**](CustomAuthorizationServer.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**Id** | Pointer to **string** | Unique identifier for the managed connection. Only present for managed connections. | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | List of scopes for the connection based on scopeCondition | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewIdentityAssertionCustomASConnection

`func NewIdentityAssertionCustomASConnection(authorizationServer CustomAuthorizationServer, connectionType string, ) *IdentityAssertionCustomASConnection`

NewIdentityAssertionCustomASConnection instantiates a new IdentityAssertionCustomASConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionCustomASConnectionWithDefaults

`func NewIdentityAssertionCustomASConnectionWithDefaults() *IdentityAssertionCustomASConnection`

NewIdentityAssertionCustomASConnectionWithDefaults instantiates a new IdentityAssertionCustomASConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationServer

`func (o *IdentityAssertionCustomASConnection) GetAuthorizationServer() CustomAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *IdentityAssertionCustomASConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *IdentityAssertionCustomASConnection) SetAuthorizationServer(v CustomAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetConnectionType

`func (o *IdentityAssertionCustomASConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IdentityAssertionCustomASConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IdentityAssertionCustomASConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetId

`func (o *IdentityAssertionCustomASConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAssertionCustomASConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAssertionCustomASConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAssertionCustomASConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *IdentityAssertionCustomASConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *IdentityAssertionCustomASConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *IdentityAssertionCustomASConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *IdentityAssertionCustomASConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProtocolType

`func (o *IdentityAssertionCustomASConnection) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *IdentityAssertionCustomASConnection) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *IdentityAssertionCustomASConnection) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *IdentityAssertionCustomASConnection) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetScopeCondition

`func (o *IdentityAssertionCustomASConnection) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *IdentityAssertionCustomASConnection) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *IdentityAssertionCustomASConnection) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *IdentityAssertionCustomASConnection) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *IdentityAssertionCustomASConnection) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityAssertionCustomASConnection) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityAssertionCustomASConnection) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IdentityAssertionCustomASConnection) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityAssertionCustomASConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityAssertionCustomASConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityAssertionCustomASConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityAssertionCustomASConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *IdentityAssertionCustomASConnection) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityAssertionCustomASConnection) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityAssertionCustomASConnection) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityAssertionCustomASConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


