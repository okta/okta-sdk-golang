# ManagedConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**ManagedConnectionAppInstance**](ManagedConnectionAppInstance.md) |  | 
**AuthorizationServer** | [**CustomAuthorizationServer**](CustomAuthorizationServer.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**Id** | Pointer to **string** | Unique identifier for the managed connection. Only present for managed connections. | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | List of scopes for the connection based on scopeCondition | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 
**Secret** | [**ManagedConnectionVaultedSecret**](ManagedConnectionVaultedSecret.md) |  | 
**ServiceAccount** | [**ManagedConnectionServiceAccount**](ManagedConnectionServiceAccount.md) |  | 

## Methods

### NewManagedConnection

`func NewManagedConnection(app ManagedConnectionAppInstance, authorizationServer CustomAuthorizationServer, connectionType string, secret ManagedConnectionVaultedSecret, serviceAccount ManagedConnectionServiceAccount, ) *ManagedConnection`

NewManagedConnection instantiates a new ManagedConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionWithDefaults

`func NewManagedConnectionWithDefaults() *ManagedConnection`

NewManagedConnectionWithDefaults instantiates a new ManagedConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *ManagedConnection) GetApp() ManagedConnectionAppInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ManagedConnection) GetAppOk() (*ManagedConnectionAppInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ManagedConnection) SetApp(v ManagedConnectionAppInstance)`

SetApp sets App field to given value.


### GetAuthorizationServer

`func (o *ManagedConnection) GetAuthorizationServer() CustomAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *ManagedConnection) GetAuthorizationServerOk() (*CustomAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *ManagedConnection) SetAuthorizationServer(v CustomAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetConnectionType

`func (o *ManagedConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ManagedConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ManagedConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetId

`func (o *ManagedConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *ManagedConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ManagedConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ManagedConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *ManagedConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProtocolType

`func (o *ManagedConnection) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *ManagedConnection) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *ManagedConnection) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *ManagedConnection) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetScopeCondition

`func (o *ManagedConnection) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *ManagedConnection) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *ManagedConnection) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *ManagedConnection) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *ManagedConnection) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ManagedConnection) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ManagedConnection) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ManagedConnection) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *ManagedConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ManagedConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *ManagedConnection) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ManagedConnection) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ManagedConnection) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ManagedConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSecret

`func (o *ManagedConnection) GetSecret() ManagedConnectionVaultedSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ManagedConnection) GetSecretOk() (*ManagedConnectionVaultedSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ManagedConnection) SetSecret(v ManagedConnectionVaultedSecret)`

SetSecret sets Secret field to given value.


### GetServiceAccount

`func (o *ManagedConnection) GetServiceAccount() ManagedConnectionServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ManagedConnection) GetServiceAccountOk() (*ManagedConnectionServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ManagedConnection) SetServiceAccount(v ManagedConnectionServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


