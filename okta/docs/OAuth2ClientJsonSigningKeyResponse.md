# OAuth2ClientJsonSigningKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Algorithm used in the key | 
**Created** | **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [readonly] 
**Id** | **string** | The unique ID of the OAuth Client JSON Web Key | [readonly] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAuth 2.0 client&#39;s JWKS | [optional] 
**Kty** | **string** | Cryptographic algorithm family for the certificate&#39;s key pair | 
**LastUpdated** | **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [readonly] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] 
**Use** | **string** | Acceptable use of the JSON Web Key | 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 

## Methods

### NewOAuth2ClientJsonSigningKeyResponse

`func NewOAuth2ClientJsonSigningKeyResponse(alg string, created string, id string, kty string, lastUpdated string, use string, ) *OAuth2ClientJsonSigningKeyResponse`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


