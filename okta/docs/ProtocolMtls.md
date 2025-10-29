# ProtocolMtls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**MtlsCredentials**](MtlsCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**MtlsEndpoints**](MtlsEndpoints.md) |  | [optional] 
**Type** | Pointer to **string** | Mutual TLS | [optional] 

## Methods

### NewProtocolMtls

`func NewProtocolMtls() *ProtocolMtls`

NewProtocolMtls instantiates a new ProtocolMtls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolMtlsWithDefaults

`func NewProtocolMtlsWithDefaults() *ProtocolMtls`

NewProtocolMtlsWithDefaults instantiates a new ProtocolMtls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ProtocolMtls) GetCredentials() MtlsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ProtocolMtls) GetCredentialsOk() (*MtlsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ProtocolMtls) SetCredentials(v MtlsCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ProtocolMtls) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *ProtocolMtls) GetEndpoints() MtlsEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ProtocolMtls) GetEndpointsOk() (*MtlsEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ProtocolMtls) SetEndpoints(v MtlsEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ProtocolMtls) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetType

`func (o *ProtocolMtls) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtocolMtls) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtocolMtls) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtocolMtls) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


