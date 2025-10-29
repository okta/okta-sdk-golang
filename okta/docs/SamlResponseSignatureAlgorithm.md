# SamlResponseSignatureAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** | Specifies whether to verify responses from the IdP | [optional] 

## Methods

### NewSamlResponseSignatureAlgorithm

`func NewSamlResponseSignatureAlgorithm() *SamlResponseSignatureAlgorithm`

NewSamlResponseSignatureAlgorithm instantiates a new SamlResponseSignatureAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlResponseSignatureAlgorithmWithDefaults

`func NewSamlResponseSignatureAlgorithmWithDefaults() *SamlResponseSignatureAlgorithm`

NewSamlResponseSignatureAlgorithmWithDefaults instantiates a new SamlResponseSignatureAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *SamlResponseSignatureAlgorithm) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *SamlResponseSignatureAlgorithm) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *SamlResponseSignatureAlgorithm) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *SamlResponseSignatureAlgorithm) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetScope

`func (o *SamlResponseSignatureAlgorithm) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SamlResponseSignatureAlgorithm) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SamlResponseSignatureAlgorithm) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SamlResponseSignatureAlgorithm) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


