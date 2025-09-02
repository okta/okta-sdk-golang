# OAuth2ResourceServerJsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Timestamp when the JSON Web Key was created | [optional] [readonly] 
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Id** | Pointer to **string** | The unique ID of the JSON Web Key | [optional] [readonly] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the Custom Authorization Server&#39;s Public JWKS | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**LastUpdated** | Pointer to **string** | Timestamp when the JSON Web Key was updated | [optional] [readonly] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Status** | Pointer to **string** | The status of the encryption key. You can use only an &#x60;ACTIVE&#x60; key to encrypt tokens issued by the authorization server. | [optional] [default to "ACTIVE"]
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Links** | Pointer to [**OAuthResourceServerKeyLinks**](OAuthResourceServerKeyLinks.md) |  | [optional] 

## Methods

### NewOAuth2ResourceServerJsonWebKey

`func NewOAuth2ResourceServerJsonWebKey() *OAuth2ResourceServerJsonWebKey`

NewOAuth2ResourceServerJsonWebKey instantiates a new OAuth2ResourceServerJsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ResourceServerJsonWebKeyWithDefaults

`func NewOAuth2ResourceServerJsonWebKeyWithDefaults() *OAuth2ResourceServerJsonWebKey`

NewOAuth2ResourceServerJsonWebKeyWithDefaults instantiates a new OAuth2ResourceServerJsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *OAuth2ResourceServerJsonWebKey) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2ResourceServerJsonWebKey) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2ResourceServerJsonWebKey) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OAuth2ResourceServerJsonWebKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetE

`func (o *OAuth2ResourceServerJsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *OAuth2ResourceServerJsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *OAuth2ResourceServerJsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *OAuth2ResourceServerJsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetId

`func (o *OAuth2ResourceServerJsonWebKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ResourceServerJsonWebKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ResourceServerJsonWebKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2ResourceServerJsonWebKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ResourceServerJsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ResourceServerJsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ResourceServerJsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ResourceServerJsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ResourceServerJsonWebKey) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ResourceServerJsonWebKey) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetKty

`func (o *OAuth2ResourceServerJsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2ResourceServerJsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2ResourceServerJsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2ResourceServerJsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OAuth2ResourceServerJsonWebKey) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OAuth2ResourceServerJsonWebKey) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OAuth2ResourceServerJsonWebKey) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OAuth2ResourceServerJsonWebKey) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetN

`func (o *OAuth2ResourceServerJsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *OAuth2ResourceServerJsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *OAuth2ResourceServerJsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *OAuth2ResourceServerJsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStatus

`func (o *OAuth2ResourceServerJsonWebKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ResourceServerJsonWebKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ResourceServerJsonWebKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ResourceServerJsonWebKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *OAuth2ResourceServerJsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *OAuth2ResourceServerJsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *OAuth2ResourceServerJsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *OAuth2ResourceServerJsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2ResourceServerJsonWebKey) GetLinks() OAuthResourceServerKeyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2ResourceServerJsonWebKey) GetLinksOk() (*OAuthResourceServerKeyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2ResourceServerJsonWebKey) SetLinks(v OAuthResourceServerKeyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2ResourceServerJsonWebKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


