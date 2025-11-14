# IdentityAssertionAppInstanceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**ManagedConnectionAppInstance**](ManagedConnectionAppInstance.md) |  | 
**AuthorizationServer** | Pointer to [**ManagedConnectionAuthorizationServer**](ManagedConnectionAuthorizationServer.md) |  | [optional] 
**ConnectionType** | **string** | Type of connection authentication method | 
**Id** | Pointer to **string** | Unique identifier for the managed connection. Only present for managed connections. | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | List of scopes for the connection based on scopeCondition | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewIdentityAssertionAppInstanceConnection

`func NewIdentityAssertionAppInstanceConnection(app ManagedConnectionAppInstance, connectionType string, ) *IdentityAssertionAppInstanceConnection`

NewIdentityAssertionAppInstanceConnection instantiates a new IdentityAssertionAppInstanceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionAppInstanceConnectionWithDefaults

`func NewIdentityAssertionAppInstanceConnectionWithDefaults() *IdentityAssertionAppInstanceConnection`

NewIdentityAssertionAppInstanceConnectionWithDefaults instantiates a new IdentityAssertionAppInstanceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *IdentityAssertionAppInstanceConnection) GetApp() ManagedConnectionAppInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *IdentityAssertionAppInstanceConnection) GetAppOk() (*ManagedConnectionAppInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *IdentityAssertionAppInstanceConnection) SetApp(v ManagedConnectionAppInstance)`

SetApp sets App field to given value.


### GetAuthorizationServer

`func (o *IdentityAssertionAppInstanceConnection) GetAuthorizationServer() ManagedConnectionAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *IdentityAssertionAppInstanceConnection) GetAuthorizationServerOk() (*ManagedConnectionAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *IdentityAssertionAppInstanceConnection) SetAuthorizationServer(v ManagedConnectionAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.

### HasAuthorizationServer

`func (o *IdentityAssertionAppInstanceConnection) HasAuthorizationServer() bool`

HasAuthorizationServer returns a boolean if a field has been set.

### GetConnectionType

`func (o *IdentityAssertionAppInstanceConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IdentityAssertionAppInstanceConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IdentityAssertionAppInstanceConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetId

`func (o *IdentityAssertionAppInstanceConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityAssertionAppInstanceConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityAssertionAppInstanceConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityAssertionAppInstanceConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *IdentityAssertionAppInstanceConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *IdentityAssertionAppInstanceConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *IdentityAssertionAppInstanceConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *IdentityAssertionAppInstanceConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProtocolType

`func (o *IdentityAssertionAppInstanceConnection) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *IdentityAssertionAppInstanceConnection) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *IdentityAssertionAppInstanceConnection) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *IdentityAssertionAppInstanceConnection) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetScopeCondition

`func (o *IdentityAssertionAppInstanceConnection) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *IdentityAssertionAppInstanceConnection) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *IdentityAssertionAppInstanceConnection) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *IdentityAssertionAppInstanceConnection) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *IdentityAssertionAppInstanceConnection) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityAssertionAppInstanceConnection) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityAssertionAppInstanceConnection) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IdentityAssertionAppInstanceConnection) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *IdentityAssertionAppInstanceConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IdentityAssertionAppInstanceConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IdentityAssertionAppInstanceConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IdentityAssertionAppInstanceConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *IdentityAssertionAppInstanceConnection) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityAssertionAppInstanceConnection) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityAssertionAppInstanceConnection) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityAssertionAppInstanceConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


