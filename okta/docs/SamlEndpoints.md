# SamlEndpoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acs** | Pointer to [**SamlAcsEndpoint**](SamlAcsEndpoint.md) |  | [optional] 
**Slo** | Pointer to [**SamlSloEndpoint**](SamlSloEndpoint.md) |  | [optional] 
**Sso** | Pointer to [**SamlSsoEndpoint**](SamlSsoEndpoint.md) |  | [optional] 

## Methods

### NewSamlEndpoints

`func NewSamlEndpoints() *SamlEndpoints`

NewSamlEndpoints instantiates a new SamlEndpoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlEndpointsWithDefaults

`func NewSamlEndpointsWithDefaults() *SamlEndpoints`

NewSamlEndpointsWithDefaults instantiates a new SamlEndpoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcs

`func (o *SamlEndpoints) GetAcs() SamlAcsEndpoint`

GetAcs returns the Acs field if non-nil, zero value otherwise.

### GetAcsOk

`func (o *SamlEndpoints) GetAcsOk() (*SamlAcsEndpoint, bool)`

GetAcsOk returns a tuple with the Acs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcs

`func (o *SamlEndpoints) SetAcs(v SamlAcsEndpoint)`

SetAcs sets Acs field to given value.

### HasAcs

`func (o *SamlEndpoints) HasAcs() bool`

HasAcs returns a boolean if a field has been set.

### GetSlo

`func (o *SamlEndpoints) GetSlo() SamlSloEndpoint`

GetSlo returns the Slo field if non-nil, zero value otherwise.

### GetSloOk

`func (o *SamlEndpoints) GetSloOk() (*SamlSloEndpoint, bool)`

GetSloOk returns a tuple with the Slo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlo

`func (o *SamlEndpoints) SetSlo(v SamlSloEndpoint)`

SetSlo sets Slo field to given value.

### HasSlo

`func (o *SamlEndpoints) HasSlo() bool`

HasSlo returns a boolean if a field has been set.

### GetSso

`func (o *SamlEndpoints) GetSso() SamlSsoEndpoint`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SamlEndpoints) GetSsoOk() (*SamlSsoEndpoint, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SamlEndpoints) SetSso(v SamlSsoEndpoint)`

SetSso sets Sso field to given value.

### HasSso

`func (o *SamlEndpoints) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


