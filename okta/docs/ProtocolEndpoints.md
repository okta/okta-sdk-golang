# ProtocolEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acs** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**Authorization** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**Jwks** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**Metadata** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**Slo** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**Sso** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**Token** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 
**UserInfo** | Pointer to [**ProtocolEndpoint**](ProtocolEndpoint.md) |  | [optional] 

## Methods

### NewProtocolEndpoints

`func NewProtocolEndpoints() *ProtocolEndpoints`

NewProtocolEndpoints instantiates a new ProtocolEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolEndpointsWithDefaults

`func NewProtocolEndpointsWithDefaults() *ProtocolEndpoints`

NewProtocolEndpointsWithDefaults instantiates a new ProtocolEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcs

`func (o *ProtocolEndpoints) GetAcs() ProtocolEndpoint`

GetAcs returns the Acs field if non-nil, zero value otherwise.

### GetAcsOk

`func (o *ProtocolEndpoints) GetAcsOk() (*ProtocolEndpoint, bool)`

GetAcsOk returns a tuple with the Acs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcs

`func (o *ProtocolEndpoints) SetAcs(v ProtocolEndpoint)`

SetAcs sets Acs field to given value.

### HasAcs

`func (o *ProtocolEndpoints) HasAcs() bool`

HasAcs returns a boolean if a field has been set.

### GetAuthorization

`func (o *ProtocolEndpoints) GetAuthorization() ProtocolEndpoint`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *ProtocolEndpoints) GetAuthorizationOk() (*ProtocolEndpoint, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *ProtocolEndpoints) SetAuthorization(v ProtocolEndpoint)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *ProtocolEndpoints) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### GetJwks

`func (o *ProtocolEndpoints) GetJwks() ProtocolEndpoint`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *ProtocolEndpoints) GetJwksOk() (*ProtocolEndpoint, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *ProtocolEndpoints) SetJwks(v ProtocolEndpoint)`

SetJwks sets Jwks field to given value.

### HasJwks

`func (o *ProtocolEndpoints) HasJwks() bool`

HasJwks returns a boolean if a field has been set.

### GetMetadata

`func (o *ProtocolEndpoints) GetMetadata() ProtocolEndpoint`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProtocolEndpoints) GetMetadataOk() (*ProtocolEndpoint, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProtocolEndpoints) SetMetadata(v ProtocolEndpoint)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProtocolEndpoints) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSlo

`func (o *ProtocolEndpoints) GetSlo() ProtocolEndpoint`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *ProtocolEndpoints) GetSloOk() (*ProtocolEndpoint, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *ProtocolEndpoints) SetSlo(v ProtocolEndpoint)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *ProtocolEndpoints) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetSso

`func (o *ProtocolEndpoints) GetSso() ProtocolEndpoint`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *ProtocolEndpoints) GetSsoOk() (*ProtocolEndpoint, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *ProtocolEndpoints) SetSso(v ProtocolEndpoint)`

SetSso sets Sso field to given value.

### HasSso

`func (o *ProtocolEndpoints) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetToken

`func (o *ProtocolEndpoints) GetToken() ProtocolEndpoint`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ProtocolEndpoints) GetTokenOk() (*ProtocolEndpoint, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ProtocolEndpoints) SetToken(v ProtocolEndpoint)`

SetToken sets Token field to given value.

### HasToken

`func (o *ProtocolEndpoints) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUserInfo

`func (o *ProtocolEndpoints) GetUserInfo() ProtocolEndpoint`

GetUserInfo returns the UserInfo field if non-nil, zero value otherwise.

### GetUserInfoOk

`func (o *ProtocolEndpoints) GetUserInfoOk() (*ProtocolEndpoint, bool)`

GetUserInfoOk returns a tuple with the UserInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfo

`func (o *ProtocolEndpoints) SetUserInfo(v ProtocolEndpoint)`

SetUserInfo sets UserInfo field to given value.

### HasUserInfo

`func (o *ProtocolEndpoints) HasUserInfo() bool`

HasUserInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


