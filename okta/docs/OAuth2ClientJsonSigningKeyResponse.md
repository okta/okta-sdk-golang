# OAuth2ClientJsonSigningKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Algorithm used in the key | [optional] 
**Created** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [optional] [readonly] 
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client JSON Web Key | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [optional] [readonly] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAuth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 
**Crv** | Pointer to **string** | Identifies the cryptographic curve used with the key | [optional] 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 

## Methods

### NewOAuth2ClientJsonSigningKeyResponse

`func NewOAuth2ClientJsonSigningKeyResponse() *OAuth2ClientJsonSigningKeyResponse`

NewOAuth2ClientJsonSigningKeyResponse instantiates a new OAuth2ClientJsonSigningKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonSigningKeyResponseWithDefaults

`func NewOAuth2ClientJsonSigningKeyResponseWithDefaults() *OAuth2ClientJsonSigningKeyResponse`

NewOAuth2ClientJsonSigningKeyResponseWithDefaults instantiates a new OAuth2ClientJsonSigningKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *OAuth2ClientJsonSigningKeyResponse) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *OAuth2ClientJsonSigningKeyResponse) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *OAuth2ClientJsonSigningKeyResponse) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetCreated

`func (o *OAuth2ClientJsonSigningKeyResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2ClientJsonSigningKeyResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OAuth2ClientJsonSigningKeyResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetE

`func (o *OAuth2ClientJsonSigningKeyResponse) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OAuth2ClientJsonSigningKeyResponse) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OAuth2ClientJsonSigningKeyResponse) HasE() bool`

HasE returns a boolean if a field has been set.

### GetId

`func (o *OAuth2ClientJsonSigningKeyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ClientJsonSigningKeyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2ClientJsonSigningKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKty

`func (o *OAuth2ClientJsonSigningKeyResponse) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2ClientJsonSigningKeyResponse) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2ClientJsonSigningKeyResponse) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OAuth2ClientJsonSigningKeyResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OAuth2ClientJsonSigningKeyResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OAuth2ClientJsonSigningKeyResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetN

`func (o *OAuth2ClientJsonSigningKeyResponse) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OAuth2ClientJsonSigningKeyResponse) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *OAuth2ClientJsonSigningKeyResponse) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ClientJsonSigningKeyResponse) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonSigningKeyResponse) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonSigningKeyResponse) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonSigningKeyResponse) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonSigningKeyResponse) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonSigningKeyResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonSigningKeyResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonSigningKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *OAuth2ClientJsonSigningKeyResponse) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *OAuth2ClientJsonSigningKeyResponse) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *OAuth2ClientJsonSigningKeyResponse) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2ClientJsonSigningKeyResponse) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2ClientJsonSigningKeyResponse) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2ClientJsonSigningKeyResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCrv

`func (o *OAuth2ClientJsonSigningKeyResponse) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *OAuth2ClientJsonSigningKeyResponse) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *OAuth2ClientJsonSigningKeyResponse) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetX

`func (o *OAuth2ClientJsonSigningKeyResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *OAuth2ClientJsonSigningKeyResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *OAuth2ClientJsonSigningKeyResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *OAuth2ClientJsonSigningKeyResponse) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *OAuth2ClientJsonSigningKeyResponse) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *OAuth2ClientJsonSigningKeyResponse) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *OAuth2ClientJsonSigningKeyResponse) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


