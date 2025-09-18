# OidcRequestSignatureAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | Specifies whether to digitally sign authorization requests to the IdP | [optional] 

## Methods

### NewOidcRequestSignatureAlgorithm

`func NewOidcRequestSignatureAlgorithm() *OidcRequestSignatureAlgorithm`

NewOidcRequestSignatureAlgorithm instantiates a new OidcRequestSignatureAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcRequestSignatureAlgorithmWithDefaults

`func NewOidcRequestSignatureAlgorithmWithDefaults() *OidcRequestSignatureAlgorithm`

NewOidcRequestSignatureAlgorithmWithDefaults instantiates a new OidcRequestSignatureAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *OidcRequestSignatureAlgorithm) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *OidcRequestSignatureAlgorithm) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *OidcRequestSignatureAlgorithm) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *OidcRequestSignatureAlgorithm) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetScope

`func (o *OidcRequestSignatureAlgorithm) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OidcRequestSignatureAlgorithm) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OidcRequestSignatureAlgorithm) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OidcRequestSignatureAlgorithm) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


