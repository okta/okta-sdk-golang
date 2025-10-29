# OAuth2ClientJsonWebKeyECResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**Created** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was created | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of the OAuth Client JSON Web Key | [optional] [readonly] 
**LastUpdated** | Pointer to **string** | Timestamp when the OAuth 2.0 client JSON Web Key was updated | [optional] [readonly] 
**Links** | Pointer to [**OAuthClientSecretLinks**](OAuthClientSecretLinks.md) |  | [optional] 

## Methods

### NewOAuth2ClientJsonWebKeyECResponse

`func NewOAuth2ClientJsonWebKeyECResponse() *OAuth2ClientJsonWebKeyECResponse`

NewOAuth2ClientJsonWebKeyECResponse instantiates a new OAuth2ClientJsonWebKeyECResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyECResponseWithDefaults

`func NewOAuth2ClientJsonWebKeyECResponseWithDefaults() *OAuth2ClientJsonWebKeyECResponse`

NewOAuth2ClientJsonWebKeyECResponseWithDefaults instantiates a new OAuth2ClientJsonWebKeyECResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKty

`func (o *OAuth2ClientJsonWebKeyECResponse) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2ClientJsonWebKeyECResponse) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2ClientJsonWebKeyECResponse) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetX

`func (o *OAuth2ClientJsonWebKeyECResponse) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *OAuth2ClientJsonWebKeyECResponse) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *OAuth2ClientJsonWebKeyECResponse) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *OAuth2ClientJsonWebKeyECResponse) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *OAuth2ClientJsonWebKeyECResponse) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *OAuth2ClientJsonWebKeyECResponse) HasY() bool`

HasY returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ClientJsonWebKeyECResponse) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonWebKeyECResponse) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonWebKeyECResponse) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonWebKeyECResponse) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonWebKeyECResponse) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonWebKeyECResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonWebKeyECResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonWebKeyECResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreated

`func (o *OAuth2ClientJsonWebKeyECResponse) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OAuth2ClientJsonWebKeyECResponse) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OAuth2ClientJsonWebKeyECResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *OAuth2ClientJsonWebKeyECResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OAuth2ClientJsonWebKeyECResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OAuth2ClientJsonWebKeyECResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *OAuth2ClientJsonWebKeyECResponse) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OAuth2ClientJsonWebKeyECResponse) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OAuth2ClientJsonWebKeyECResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLinks

`func (o *OAuth2ClientJsonWebKeyECResponse) GetLinks() OAuthClientSecretLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetLinksOk() (*OAuthClientSecretLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OAuth2ClientJsonWebKeyECResponse) SetLinks(v OAuthClientSecretLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OAuth2ClientJsonWebKeyECResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


