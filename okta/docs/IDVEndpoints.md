# IDVEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | [**IDVAuthorizationEndpoint**](IDVAuthorizationEndpoint.md) |  | 
**Jwks** | [**OidcJwksEndpoint**](OidcJwksEndpoint.md) |  | 
**Par** | [**IDVParEndpoint**](IDVParEndpoint.md) |  | 
**Token** | [**IDVTokenEndpoint**](IDVTokenEndpoint.md) |  | 

## Methods

### NewIDVEndpoints

`func NewIDVEndpoints(authorization IDVAuthorizationEndpoint, jwks OidcJwksEndpoint, par IDVParEndpoint, token IDVTokenEndpoint, ) *IDVEndpoints`

NewIDVEndpoints instantiates a new IDVEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIDVEndpointsWithDefaults

`func NewIDVEndpointsWithDefaults() *IDVEndpoints`

NewIDVEndpointsWithDefaults instantiates a new IDVEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorization

`func (o *IDVEndpoints) GetAuthorization() IDVAuthorizationEndpoint`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *IDVEndpoints) GetAuthorizationOk() (*IDVAuthorizationEndpoint, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *IDVEndpoints) SetAuthorization(v IDVAuthorizationEndpoint)`

SetAuthorization sets Authorization field to given value.


### GetJwks

`func (o *IDVEndpoints) GetJwks() OidcJwksEndpoint`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### GetJwksOk

`func (o *IDVEndpoints) GetJwksOk() (*OidcJwksEndpoint, bool)`

GetJwksOk returns a tuple with the Jwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwks

`func (o *IDVEndpoints) SetJwks(v OidcJwksEndpoint)`

SetJwks sets Jwks field to given value.


### GetPar

`func (o *IDVEndpoints) GetPar() IDVParEndpoint`

GetPar returns the Par field if non-nil, zero value otherwise.

### GetParOk

`func (o *IDVEndpoints) GetParOk() (*IDVParEndpoint, bool)`

GetParOk returns a tuple with the Par field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPar

`func (o *IDVEndpoints) SetPar(v IDVParEndpoint)`

SetPar sets Par field to given value.


### GetToken

`func (o *IDVEndpoints) GetToken() IDVTokenEndpoint`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IDVEndpoints) GetTokenOk() (*IDVTokenEndpoint, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IDVEndpoints) SetToken(v IDVTokenEndpoint)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


