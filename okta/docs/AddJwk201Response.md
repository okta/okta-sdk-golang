# AddJwk201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [optional] [readonly] 
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client JSON Web Key | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [optional] [readonly] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAuth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 

## Methods

### NewAddJwk201Response

`func NewAddJwk201Response() *AddJwk201Response`

NewAddJwk201Response instantiates a new AddJwk201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJwk201ResponseWithDefaults

`func NewAddJwk201ResponseWithDefaults() *AddJwk201Response`

NewAddJwk201ResponseWithDefaults instantiates a new AddJwk201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AddJwk201Response) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AddJwk201Response) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AddJwk201Response) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AddJwk201Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetE

`func (o *AddJwk201Response) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AddJwk201Response) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AddJwk201Response) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AddJwk201Response) HasE() bool`

HasE returns a boolean if a field has been set.

### GetId

`func (o *AddJwk201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddJwk201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddJwk201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddJwk201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKty

`func (o *AddJwk201Response) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AddJwk201Response) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AddJwk201Response) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *AddJwk201Response) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddJwk201Response) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddJwk201Response) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddJwk201Response) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddJwk201Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetN

`func (o *AddJwk201Response) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AddJwk201Response) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AddJwk201Response) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AddJwk201Response) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *AddJwk201Response) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AddJwk201Response) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AddJwk201Response) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AddJwk201Response) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKid

`func (o *AddJwk201Response) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AddJwk201Response) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AddJwk201Response) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AddJwk201Response) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *AddJwk201Response) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *AddJwk201Response) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *AddJwk201Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddJwk201Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddJwk201Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddJwk201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *AddJwk201Response) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AddJwk201Response) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AddJwk201Response) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AddJwk201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


