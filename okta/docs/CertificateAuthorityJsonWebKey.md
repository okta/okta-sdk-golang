# CertificateAuthorityJsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | The RSA public exponent, base64url-encoded | [optional] [readonly] 
**Kty** | Pointer to **string** | The key type. Always &#x60;RSA&#x60;, because Okta generates RSA key pairs for these authorities. | [optional] [readonly] 
**N** | Pointer to **string** | The RSA modulus, base64url-encoded | [optional] [readonly] 
**Use** | Pointer to **string** | What the key is used for. Always &#x60;sig&#x60;, because a CA key signs certificates. | [optional] [readonly] 

## Methods

### NewCertificateAuthorityJsonWebKey

`func NewCertificateAuthorityJsonWebKey() *CertificateAuthorityJsonWebKey`

NewCertificateAuthorityJsonWebKey instantiates a new CertificateAuthorityJsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateAuthorityJsonWebKeyWithDefaults

`func NewCertificateAuthorityJsonWebKeyWithDefaults() *CertificateAuthorityJsonWebKey`

NewCertificateAuthorityJsonWebKeyWithDefaults instantiates a new CertificateAuthorityJsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *CertificateAuthorityJsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *CertificateAuthorityJsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *CertificateAuthorityJsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *CertificateAuthorityJsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *CertificateAuthorityJsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *CertificateAuthorityJsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *CertificateAuthorityJsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *CertificateAuthorityJsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *CertificateAuthorityJsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *CertificateAuthorityJsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *CertificateAuthorityJsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *CertificateAuthorityJsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *CertificateAuthorityJsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *CertificateAuthorityJsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *CertificateAuthorityJsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *CertificateAuthorityJsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


