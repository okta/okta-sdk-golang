# OAuth2ClientJsonWebKeyRsaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | **string** | RSA key value (exponent) for key binding | 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | **string** | RSA key value (modulus) for key binding | 

## Methods

### NewOAuth2ClientJsonWebKeyRsaRequest

`func NewOAuth2ClientJsonWebKeyRsaRequest(e string, n string, ) *OAuth2ClientJsonWebKeyRsaRequest`

NewOAuth2ClientJsonWebKeyRsaRequest instantiates a new OAuth2ClientJsonWebKeyRsaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyRsaRequestWithDefaults

`func NewOAuth2ClientJsonWebKeyRsaRequestWithDefaults() *OAuth2ClientJsonWebKeyRsaRequest`

NewOAuth2ClientJsonWebKeyRsaRequestWithDefaults instantiates a new OAuth2ClientJsonWebKeyRsaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OAuth2ClientJsonWebKeyRsaRequest) SetE(v string)`

SetE sets E field to given value.


### GetKty

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2ClientJsonWebKeyRsaRequest) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2ClientJsonWebKeyRsaRequest) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OAuth2ClientJsonWebKeyRsaRequest) SetN(v string)`

SetN sets N field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


