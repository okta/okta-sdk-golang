# DeactivateOAuth2ClientJsonWebKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Algorithm used in the key | 
**Created** | **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [readonly] 
**Id** | **string** | The unique ID of the OAuth Client JSON Web Key | [readonly] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAuth 2.0 client&#39;s JWKS | [optional] 
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Kty** | **string** | Cryptographic algorithm family for the certificate&#39;s key pair | 
**X** | **string** | The public x coordinate for the elliptic curve point | 
**Y** | **string** | The public y coordinate for the elliptic curve point | 
**LastUpdated** | **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [readonly] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] 
**Use** | **string** | Acceptable use of the JSON Web Key | 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 
**Crv** | **string** | The cryptographic curve used with the key | 

## Methods

### NewDeactivateOAuth2ClientJsonWebKey200Response

`func NewDeactivateOAuth2ClientJsonWebKey200Response(alg string, created string, id string, kty string, x string, y string, lastUpdated string, use string, crv string, ) *DeactivateOAuth2ClientJsonWebKey200Response`

NewDeactivateOAuth2ClientJsonWebKey200Response instantiates a new DeactivateOAuth2ClientJsonWebKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeactivateOAuth2ClientJsonWebKey200ResponseWithDefaults

`func NewDeactivateOAuth2ClientJsonWebKey200ResponseWithDefaults() *DeactivateOAuth2ClientJsonWebKey200Response`

NewDeactivateOAuth2ClientJsonWebKey200ResponseWithDefaults instantiates a new DeactivateOAuth2ClientJsonWebKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetCreated

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetId

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetId(v string)`

SetId sets Id field to given value.


### GetKid

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *DeactivateOAuth2ClientJsonWebKey200Response) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetE

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) HasE() bool`

HasE returns a boolean if a field has been set.

### GetN

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKty

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetKty(v string)`

SetKty sets Kty field to given value.


### GetX

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetX(v string)`

SetX sets X field to given value.


### GetY

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetY(v string)`

SetY sets Y field to given value.


### GetLastUpdated

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.


### GetStatus

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetUse(v string)`

SetUse sets Use field to given value.


### GetLinks

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCrv

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *DeactivateOAuth2ClientJsonWebKey200Response) SetCrv(v string)`

SetCrv sets Crv field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


