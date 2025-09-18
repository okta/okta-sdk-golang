# ListJwk200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [optional] [readonly] 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 

## Methods

### NewListJwk200ResponseInner

`func NewListJwk200ResponseInner() *ListJwk200ResponseInner`

NewListJwk200ResponseInner instantiates a new ListJwk200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListJwk200ResponseInnerWithDefaults

`func NewListJwk200ResponseInnerWithDefaults() *ListJwk200ResponseInner`

NewListJwk200ResponseInnerWithDefaults instantiates a new ListJwk200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListJwk200ResponseInner) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListJwk200ResponseInner) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListJwk200ResponseInner) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListJwk200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ListJwk200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListJwk200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListJwk200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListJwk200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListJwk200ResponseInner) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListJwk200ResponseInner) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListJwk200ResponseInner) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListJwk200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *ListJwk200ResponseInner) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListJwk200ResponseInner) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListJwk200ResponseInner) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListJwk200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetE

`func (o *ListJwk200ResponseInner) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *ListJwk200ResponseInner) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *ListJwk200ResponseInner) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *ListJwk200ResponseInner) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *ListJwk200ResponseInner) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *ListJwk200ResponseInner) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *ListJwk200ResponseInner) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *ListJwk200ResponseInner) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *ListJwk200ResponseInner) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ListJwk200ResponseInner) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ListJwk200ResponseInner) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *ListJwk200ResponseInner) HasN() bool`

HasN returns a boolean if a field has been set.

### GetKid

`func (o *ListJwk200ResponseInner) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *ListJwk200ResponseInner) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *ListJwk200ResponseInner) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *ListJwk200ResponseInner) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *ListJwk200ResponseInner) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *ListJwk200ResponseInner) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *ListJwk200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListJwk200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListJwk200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListJwk200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetX

`func (o *ListJwk200ResponseInner) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ListJwk200ResponseInner) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ListJwk200ResponseInner) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *ListJwk200ResponseInner) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *ListJwk200ResponseInner) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ListJwk200ResponseInner) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ListJwk200ResponseInner) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *ListJwk200ResponseInner) HasY() bool`

HasY returns a boolean if a field has been set.

### GetUse

`func (o *ListJwk200ResponseInner) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *ListJwk200ResponseInner) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *ListJwk200ResponseInner) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *ListJwk200ResponseInner) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


