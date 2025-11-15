# GetJwk200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**Created** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [optional] [readonly] 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 

## Methods

### NewGetJwk200Response

`func NewGetJwk200Response() *GetJwk200Response`

NewGetJwk200Response instantiates a new GetJwk200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJwk200ResponseWithDefaults

`func NewGetJwk200ResponseWithDefaults() *GetJwk200Response`

NewGetJwk200ResponseWithDefaults instantiates a new GetJwk200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *GetJwk200Response) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *GetJwk200Response) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *GetJwk200Response) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *GetJwk200Response) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *GetJwk200Response) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *GetJwk200Response) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *GetJwk200Response) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *GetJwk200Response) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *GetJwk200Response) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *GetJwk200Response) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *GetJwk200Response) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *GetJwk200Response) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *GetJwk200Response) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *GetJwk200Response) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *GetJwk200Response) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *GetJwk200Response) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKid

`func (o *GetJwk200Response) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *GetJwk200Response) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *GetJwk200Response) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *GetJwk200Response) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *GetJwk200Response) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *GetJwk200Response) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *GetJwk200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJwk200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJwk200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJwk200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *GetJwk200Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetJwk200Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetJwk200Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetJwk200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GetJwk200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetJwk200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetJwk200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetJwk200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetJwk200Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetJwk200Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetJwk200Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetJwk200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *GetJwk200Response) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetJwk200Response) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetJwk200Response) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GetJwk200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


