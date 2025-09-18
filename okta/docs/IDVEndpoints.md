# IDVEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | Pointer to [**IDVAuthorizationEndpoint**](IDVAuthorizationEndpoint.md) |  | [optional] 
**Par** | Pointer to [**IDVParEndpoint**](IDVParEndpoint.md) |  | [optional] 
**Token** | Pointer to [**IDVTokenEndpoint**](IDVTokenEndpoint.md) |  | [optional] 

## Methods

### NewIDVEndpoints

`func NewIDVEndpoints() *IDVEndpoints`

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

### HasAuthorization

`func (o *IDVEndpoints) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

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

### HasPar

`func (o *IDVEndpoints) HasPar() bool`

HasPar returns a boolean if a field has been set.

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

### HasToken

`func (o *IDVEndpoints) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


