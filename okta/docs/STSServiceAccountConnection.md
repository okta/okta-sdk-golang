# STSServiceAccountConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | [**ManagedConnectionAppInstance**](ManagedConnectionAppInstance.md) |  | 
**ConnectionType** | **string** | Type of connection authentication method | 
**Id** | Pointer to **string** | Unique identifier for the managed connection. Only present for managed connections. | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**ProtocolType** | Pointer to **string** | The authentication protocol type used for the connection | [optional] 
**ServiceAccount** | [**ManagedConnectionServiceAccount**](ManagedConnectionServiceAccount.md) |  | 
**Status** | Pointer to **string** | The status of the connection | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewSTSServiceAccountConnection

`func NewSTSServiceAccountConnection(app ManagedConnectionAppInstance, connectionType string, serviceAccount ManagedConnectionServiceAccount, ) *STSServiceAccountConnection`

NewSTSServiceAccountConnection instantiates a new STSServiceAccountConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSTSServiceAccountConnectionWithDefaults

`func NewSTSServiceAccountConnectionWithDefaults() *STSServiceAccountConnection`

NewSTSServiceAccountConnectionWithDefaults instantiates a new STSServiceAccountConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *STSServiceAccountConnection) GetApp() ManagedConnectionAppInstance`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *STSServiceAccountConnection) GetAppOk() (*ManagedConnectionAppInstance, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *STSServiceAccountConnection) SetApp(v ManagedConnectionAppInstance)`

SetApp sets App field to given value.


### GetConnectionType

`func (o *STSServiceAccountConnection) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *STSServiceAccountConnection) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *STSServiceAccountConnection) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.


### GetId

`func (o *STSServiceAccountConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *STSServiceAccountConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *STSServiceAccountConnection) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *STSServiceAccountConnection) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *STSServiceAccountConnection) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *STSServiceAccountConnection) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *STSServiceAccountConnection) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *STSServiceAccountConnection) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetProtocolType

`func (o *STSServiceAccountConnection) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *STSServiceAccountConnection) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *STSServiceAccountConnection) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.

### HasProtocolType

`func (o *STSServiceAccountConnection) HasProtocolType() bool`

HasProtocolType returns a boolean if a field has been set.

### GetServiceAccount

`func (o *STSServiceAccountConnection) GetServiceAccount() ManagedConnectionServiceAccount`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *STSServiceAccountConnection) GetServiceAccountOk() (*ManagedConnectionServiceAccount, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *STSServiceAccountConnection) SetServiceAccount(v ManagedConnectionServiceAccount)`

SetServiceAccount sets ServiceAccount field to given value.


### GetStatus

`func (o *STSServiceAccountConnection) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *STSServiceAccountConnection) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *STSServiceAccountConnection) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *STSServiceAccountConnection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *STSServiceAccountConnection) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *STSServiceAccountConnection) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *STSServiceAccountConnection) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *STSServiceAccountConnection) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


