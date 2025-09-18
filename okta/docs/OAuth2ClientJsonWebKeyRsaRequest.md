# OAuth2ClientJsonWebKeyRsaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]

## Methods

### NewOAuth2ClientJsonWebKeyRsaRequest

`func NewOAuth2ClientJsonWebKeyRsaRequest() *OAuth2ClientJsonWebKeyRsaRequest`

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

### HasE

`func (o *OAuth2ClientJsonWebKeyRsaRequest) HasE() bool`

HasE returns a boolean if a field has been set.

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

### HasN

`func (o *OAuth2ClientJsonWebKeyRsaRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonWebKeyRsaRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonWebKeyRsaRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonWebKeyRsaRequest) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonWebKeyRsaRequest) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonWebKeyRsaRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonWebKeyRsaRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonWebKeyRsaRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


