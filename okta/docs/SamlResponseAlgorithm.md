# SamlResponseAlgorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | Pointer to [**SamlResponseSignatureAlgorithm**](SamlResponseSignatureAlgorithm.md) |  | [optional] 

## Methods

### NewSamlResponseAlgorithm

`func NewSamlResponseAlgorithm() *SamlResponseAlgorithm`

NewSamlResponseAlgorithm instantiates a new SamlResponseAlgorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlResponseAlgorithmWithDefaults

`func NewSamlResponseAlgorithmWithDefaults() *SamlResponseAlgorithm`

NewSamlResponseAlgorithmWithDefaults instantiates a new SamlResponseAlgorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *SamlResponseAlgorithm) GetSignature() SamlResponseSignatureAlgorithm`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *SamlResponseAlgorithm) GetSignatureOk() (*SamlResponseSignatureAlgorithm, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *SamlResponseAlgorithm) SetSignature(v SamlResponseSignatureAlgorithm)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *SamlResponseAlgorithm) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


