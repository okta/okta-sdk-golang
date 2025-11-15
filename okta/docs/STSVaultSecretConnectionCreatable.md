# STSVaultSecretConnectionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionType** | **string** | Type of connection authentication method | 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**Secret** | [**STSVaultSecretConnectionCreatableSecret**](STSVaultSecretConnectionCreatableSecret.md) |  | 

## Methods

### NewSTSVaultSecretConnectionCreatable

`func NewSTSVaultSecretConnectionCreatable(connectionType string, secret STSVaultSecretConnectionCreatableSecret, ) *STSVaultSecretConnectionCreatable`

NewSTSVaultSecretConnectionCreatable instantiates a new STSVaultSecretConnectionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTSVaultSecretConnectionCreatableWithDefaults

`func NewSTSVaultSecretConnectionCreatableWithDefaults() *STSVaultSecretConnectionCreatable`

NewSTSVaultSecretConnectionCreatableWithDefaults instantiates a new STSVaultSecretConnectionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionType

`func (o *STSVaultSecretConnectionCreatable) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *STSVaultSecretConnectionCreatable) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *STSVaultSecretConnectionCreatable) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetProtocolType

`func (o *STSVaultSecretConnectionCreatable) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *STSVaultSecretConnectionCreatable) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *STSVaultSecretConnectionCreatable) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *STSVaultSecretConnectionCreatable) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetSecret

`func (o *STSVaultSecretConnectionCreatable) GetSecret() STSVaultSecretConnectionCreatableSecret`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *STSVaultSecretConnectionCreatable) GetSecretOk() (*STSVaultSecretConnectionCreatableSecret, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *STSVaultSecretConnectionCreatable) SetSecret(v STSVaultSecretConnectionCreatableSecret)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


