# OidcJwksEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Binding** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the endpoint to the JWK Set | [optional] 

## Methods

### NewOidcJwksEndpoint

`func NewOidcJwksEndpoint() *OidcJwksEndpoint`

NewOidcJwksEndpoint instantiates a new OidcJwksEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcJwksEndpointWithDefaults

`func NewOidcJwksEndpointWithDefaults() *OidcJwksEndpoint`

NewOidcJwksEndpointWithDefaults instantiates a new OidcJwksEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinding

`func (o *OidcJwksEndpoint) GetBinding() string`

GetBinding returns the Binding field if non-nil, zero value otherwise.

### GetBindingOk

`func (o *OidcJwksEndpoint) GetBindingOk() (*string, bool)`

GetBindingOk returns a tuple with the Binding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinding

`func (o *OidcJwksEndpoint) SetBinding(v string)`

SetBinding sets Binding field to given value.

### HasBinding

`func (o *OidcJwksEndpoint) HasBinding() bool`

HasBinding returns a boolean if a field has been set.

### GetUrl

`func (o *OidcJwksEndpoint) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OidcJwksEndpoint) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OidcJwksEndpoint) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OidcJwksEndpoint) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


