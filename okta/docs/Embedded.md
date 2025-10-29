# Embedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Algorithm used in the key | [optional] [readonly] 
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] [readonly] 
**Kid** | Pointer to **string** | Unique identifier for the certificate | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s keypair | [optional] [readonly] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] [readonly] 
**Use** | Pointer to **NullableString** | Acceptable use of the certificate | [optional] [readonly] 

## Methods

### NewEmbedded

`func NewEmbedded() *Embedded`

NewEmbedded instantiates a new Embedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedWithDefaults

`func NewEmbeddedWithDefaults() *Embedded`

NewEmbeddedWithDefaults instantiates a new Embedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *Embedded) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *Embedded) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *Embedded) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *Embedded) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetE

`func (o *Embedded) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *Embedded) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *Embedded) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *Embedded) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKid

`func (o *Embedded) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *Embedded) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *Embedded) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *Embedded) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *Embedded) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *Embedded) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *Embedded) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *Embedded) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *Embedded) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *Embedded) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *Embedded) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *Embedded) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *Embedded) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *Embedded) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *Embedded) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *Embedded) HasUse() bool`

HasUse returns a boolean if a field has been set.

### SetUseNil

`func (o *Embedded) SetUseNil(b bool)`

 SetUseNil sets the value for Use to be an explicit nil

### UnsetUse
`func (o *Embedded) UnsetUse()`

UnsetUse ensures that no value is present for Use, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


