# OAuth2ResourceServerJsonWebKeyRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON web key in the custom authorization server&#39;s public JWKS | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Status** | Pointer to **string** | Status of the JSON Web Key | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 

## Methods

### NewOAuth2ResourceServerJsonWebKeyRequestBody

`func NewOAuth2ResourceServerJsonWebKeyRequestBody() *OAuth2ResourceServerJsonWebKeyRequestBody`

NewOAuth2ResourceServerJsonWebKeyRequestBody instantiates a new OAuth2ResourceServerJsonWebKeyRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ResourceServerJsonWebKeyRequestBodyWithDefaults

`func NewOAuth2ResourceServerJsonWebKeyRequestBodyWithDefaults() *OAuth2ResourceServerJsonWebKeyRequestBody`

NewOAuth2ResourceServerJsonWebKeyRequestBodyWithDefaults instantiates a new OAuth2ResourceServerJsonWebKeyRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetKty

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *OAuth2ResourceServerJsonWebKeyRequestBody) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


