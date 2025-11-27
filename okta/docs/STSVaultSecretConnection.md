# STSVaultSecretConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**Id** | Pointer to **string** | Unique identifier for the managed connection. Only present for managed connections. | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**Secret** | [**ManagedConnectionVaultedSecret**](ManagedConnectionVaultedSecret.md) |  | 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewSTSVaultSecretConnection

`func NewSTSVaultSecretConnection(connectionType string, secret ManagedConnectionVaultedSecret, ) *STSVaultSecretConnection`

NewSTSVaultSecretConnection instantiates a new STSVaultSecretConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTSVaultSecretConnectionWithDefaults

`func NewSTSVaultSecretConnectionWithDefaults() *STSVaultSecretConnection`

NewSTSVaultSecretConnectionWithDefaults instantiates a new STSVaultSecretConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *STSVaultSecretConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *STSVaultSecretConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *STSVaultSecretConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetId

`func (o *STSVaultSecretConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *STSVaultSecretConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *STSVaultSecretConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *STSVaultSecretConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *STSVaultSecretConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *STSVaultSecretConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *STSVaultSecretConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *STSVaultSecretConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProtocolType

`func (o *STSVaultSecretConnection) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *STSVaultSecretConnection) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *STSVaultSecretConnection) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *STSVaultSecretConnection) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetSecret

`func (o *STSVaultSecretConnection) GetSecret() ManagedConnectionVaultedSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *STSVaultSecretConnection) GetSecretOk() (*ManagedConnectionVaultedSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *STSVaultSecretConnection) SetSecret(v ManagedConnectionVaultedSecret)`

SetSecret sets Secret field to given value.


### GetStatus

`func (o *STSVaultSecretConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *STSVaultSecretConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *STSVaultSecretConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *STSVaultSecretConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *STSVaultSecretConnection) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *STSVaultSecretConnection) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *STSVaultSecretConnection) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *STSVaultSecretConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


