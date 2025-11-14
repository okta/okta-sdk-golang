# IdentityAssertionAppInstanceConnectionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**IdentityAssertionAppInstanceConnectionCreatableApp**](IdentityAssertionAppInstanceConnectionCreatableApp.md) |  | 
**AuthorizationServer** | [**IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer**](IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ScopeCondition** | Pointer to **string** | Determines how Okta evaluates requested scopes for the connection. | [optional] 
**Scopes** | Pointer to **[]string** | Optional array of scopes | [optional] 

## Methods

### NewIdentityAssertionAppInstanceConnectionCreatable

`func NewIdentityAssertionAppInstanceConnectionCreatable(app IdentityAssertionAppInstanceConnectionCreatableApp, authorizationServer IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer, connectionType string, ) *IdentityAssertionAppInstanceConnectionCreatable`

NewIdentityAssertionAppInstanceConnectionCreatable instantiates a new IdentityAssertionAppInstanceConnectionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityAssertionAppInstanceConnectionCreatableWithDefaults

`func NewIdentityAssertionAppInstanceConnectionCreatableWithDefaults() *IdentityAssertionAppInstanceConnectionCreatable`

NewIdentityAssertionAppInstanceConnectionCreatableWithDefaults instantiates a new IdentityAssertionAppInstanceConnectionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetApp() IdentityAssertionAppInstanceConnectionCreatableApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetAppOk() (*IdentityAssertionAppInstanceConnectionCreatableApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *IdentityAssertionAppInstanceConnectionCreatable) SetApp(v IdentityAssertionAppInstanceConnectionCreatableApp)`

SetApp sets App field to given value.


### GetAuthorizationServer

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetAuthorizationServer() IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetAuthorizationServerOk() (*IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *IdentityAssertionAppInstanceConnectionCreatable) SetAuthorizationServer(v IdentityAssertionAppInstanceConnectionCreatableAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetConnectionType

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *IdentityAssertionAppInstanceConnectionCreatable) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetProtocolType

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *IdentityAssertionAppInstanceConnectionCreatable) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *IdentityAssertionAppInstanceConnectionCreatable) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetScopeCondition

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopeCondition() string`

GetScopeCondition returns the ScopeCondition field if non-nil, zero value otherwise.

### GetScopeConditionOk

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopeConditionOk() (*string, bool)`

GetScopeConditionOk returns a tuple with the ScopeCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeCondition

`func (o *IdentityAssertionAppInstanceConnectionCreatable) SetScopeCondition(v string)`

SetScopeCondition sets ScopeCondition field to given value.

### HasScopeCondition

`func (o *IdentityAssertionAppInstanceConnectionCreatable) HasScopeCondition() bool`

HasScopeCondition returns a boolean if a field has been set.

### GetScopes

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityAssertionAppInstanceConnectionCreatable) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityAssertionAppInstanceConnectionCreatable) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IdentityAssertionAppInstanceConnectionCreatable) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


