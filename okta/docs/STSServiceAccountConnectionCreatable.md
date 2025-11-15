# STSServiceAccountConnectionCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**IdentityAssertionAppInstanceConnectionCreatableApp**](IdentityAssertionAppInstanceConnectionCreatableApp.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ServiceAccount** | [**STSServiceAccountConnectionCreatableServiceAccount**](STSServiceAccountConnectionCreatableServiceAccount.md) |  | 

## Methods

### NewSTSServiceAccountConnectionCreatable

`func NewSTSServiceAccountConnectionCreatable(app IdentityAssertionAppInstanceConnectionCreatableApp, connectionType string, serviceAccount STSServiceAccountConnectionCreatableServiceAccount, ) *STSServiceAccountConnectionCreatable`

NewSTSServiceAccountConnectionCreatable instantiates a new STSServiceAccountConnectionCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTSServiceAccountConnectionCreatableWithDefaults

`func NewSTSServiceAccountConnectionCreatableWithDefaults() *STSServiceAccountConnectionCreatable`

NewSTSServiceAccountConnectionCreatableWithDefaults instantiates a new STSServiceAccountConnectionCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *STSServiceAccountConnectionCreatable) GetApp() IdentityAssertionAppInstanceConnectionCreatableApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *STSServiceAccountConnectionCreatable) GetAppOk() (*IdentityAssertionAppInstanceConnectionCreatableApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *STSServiceAccountConnectionCreatable) SetApp(v IdentityAssertionAppInstanceConnectionCreatableApp)`

SetApp sets App field to given value.


### GetConnectionType

`func (o *STSServiceAccountConnectionCreatable) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *STSServiceAccountConnectionCreatable) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *STSServiceAccountConnectionCreatable) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetProtocolType

`func (o *STSServiceAccountConnectionCreatable) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *STSServiceAccountConnectionCreatable) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *STSServiceAccountConnectionCreatable) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *STSServiceAccountConnectionCreatable) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetServiceAccount

`func (o *STSServiceAccountConnectionCreatable) GetServiceAccount() STSServiceAccountConnectionCreatableServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *STSServiceAccountConnectionCreatable) GetServiceAccountOk() (*STSServiceAccountConnectionCreatableServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *STSServiceAccountConnectionCreatable) SetServiceAccount(v STSServiceAccountConnectionCreatableServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


