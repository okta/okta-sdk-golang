# MtlsEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sso** | Pointer to [**MtlsSsoEndpoint**](MtlsSsoEndpoint.md) |  | [optional] 

## Methods

### NewMtlsEndpoints

`func NewMtlsEndpoints() *MtlsEndpoints`

NewMtlsEndpoints instantiates a new MtlsEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMtlsEndpointsWithDefaults

`func NewMtlsEndpointsWithDefaults() *MtlsEndpoints`

NewMtlsEndpointsWithDefaults instantiates a new MtlsEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSso

`func (o *MtlsEndpoints) GetSso() MtlsSsoEndpoint`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *MtlsEndpoints) GetSsoOk() (*MtlsSsoEndpoint, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *MtlsEndpoints) SetSso(v MtlsSsoEndpoint)`

SetSso sets Sso field to given value.

### HasSso

`func (o *MtlsEndpoints) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


