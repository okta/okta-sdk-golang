# ManagedConnectionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**IdentityAssertionAppInstanceConnectionCreatableApp**](IdentityAssertionAppInstanceConnectionCreatableApp.md) |  | 
**AuthorizationServer** | [**IdentityAssertionCustomASConnectionCreatableAuthorizationServer**](IdentityAssertionCustomASConnectionCreatableAuthorizationServer.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | Optional array of scopes | [optional] 
**Secret** | [**STSVaultSecretConnectionCreatableSecret**](STSVaultSecretConnectionCreatableSecret.md) |  | 
**ServiceAccount** | [**STSServiceAccountConnectionCreatableServiceAccount**](STSServiceAccountConnectionCreatableServiceAccount.md) |  | 

## Methods

### NewManagedConnectionCreatable

`func NewManagedConnectionCreatable(app IdentityAssertionAppInstanceConnectionCreatableApp, authorizationServer IdentityAssertionCustomASConnectionCreatableAuthorizationServer, connectionType string, secret STSVaultSecretConnectionCreatableSecret, serviceAccount STSServiceAccountConnectionCreatableServiceAccount, ) *ManagedConnectionCreatable`

NewManagedConnectionCreatable instantiates a new ManagedConnectionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionCreatableWithDefaults

`func NewManagedConnectionCreatableWithDefaults() *ManagedConnectionCreatable`

NewManagedConnectionCreatableWithDefaults instantiates a new ManagedConnectionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *ManagedConnectionCreatable) GetApp() IdentityAssertionAppInstanceConnectionCreatableApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *ManagedConnectionCreatable) GetAppOk() (*IdentityAssertionAppInstanceConnectionCreatableApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *ManagedConnectionCreatable) SetApp(v IdentityAssertionAppInstanceConnectionCreatableApp)`

SetApp sets App field to given value.


### GetAuthorizationServer

`func (o *ManagedConnectionCreatable) GetAuthorizationServer() IdentityAssertionCustomASConnectionCreatableAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *ManagedConnectionCreatable) GetAuthorizationServerOk() (*IdentityAssertionCustomASConnectionCreatableAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *ManagedConnectionCreatable) SetAuthorizationServer(v IdentityAssertionCustomASConnectionCreatableAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetConnectionType

`func (o *ManagedConnectionCreatable) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *ManagedConnectionCreatable) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *ManagedConnectionCreatable) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetProtocolType

`func (o *ManagedConnectionCreatable) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *ManagedConnectionCreatable) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *ManagedConnectionCreatable) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *ManagedConnectionCreatable) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetScopeCondition

`func (o *ManagedConnectionCreatable) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *ManagedConnectionCreatable) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *ManagedConnectionCreatable) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *ManagedConnectionCreatable) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *ManagedConnectionCreatable) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ManagedConnectionCreatable) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ManagedConnectionCreatable) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ManagedConnectionCreatable) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSecret

`func (o *ManagedConnectionCreatable) GetSecret() STSVaultSecretConnectionCreatableSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ManagedConnectionCreatable) GetSecretOk() (*STSVaultSecretConnectionCreatableSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ManagedConnectionCreatable) SetSecret(v STSVaultSecretConnectionCreatableSecret)`

SetSecret sets Secret field to given value.


### GetServiceAccount

`func (o *ManagedConnectionCreatable) GetServiceAccount() STSServiceAccountConnectionCreatableServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *ManagedConnectionCreatable) GetServiceAccountOk() (*STSServiceAccountConnectionCreatableServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *ManagedConnectionCreatable) SetServiceAccount(v STSServiceAccountConnectionCreatableServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


