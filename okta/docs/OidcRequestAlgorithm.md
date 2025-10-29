# OidcRequestAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to [**OidcRequestSignatureAlgorithm**](OidcRequestSignatureAlgorithm.md) |  | [optional] 

## Methods

### NewOidcRequestAlgorithm

`func NewOidcRequestAlgorithm() *OidcRequestAlgorithm`

NewOidcRequestAlgorithm instantiates a new OidcRequestAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcRequestAlgorithmWithDefaults

`func NewOidcRequestAlgorithmWithDefaults() *OidcRequestAlgorithm`

NewOidcRequestAlgorithmWithDefaults instantiates a new OidcRequestAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *OidcRequestAlgorithm) GetSignature() OidcRequestSignatureAlgorithm`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *OidcRequestAlgorithm) GetSignatureOk() (*OidcRequestSignatureAlgorithm, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *OidcRequestAlgorithm) SetSignature(v OidcRequestSignatureAlgorithm)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *OidcRequestAlgorithm) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


