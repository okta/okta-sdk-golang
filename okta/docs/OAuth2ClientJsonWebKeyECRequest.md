# OAuth2ClientJsonWebKeyECRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]

## Methods

### NewOAuth2ClientJsonWebKeyECRequest

`func NewOAuth2ClientJsonWebKeyECRequest() *OAuth2ClientJsonWebKeyECRequest`

NewOAuth2ClientJsonWebKeyECRequest instantiates a new OAuth2ClientJsonWebKeyECRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyECRequestWithDefaults

`func NewOAuth2ClientJsonWebKeyECRequestWithDefaults() *OAuth2ClientJsonWebKeyECRequest`

NewOAuth2ClientJsonWebKeyECRequestWithDefaults instantiates a new OAuth2ClientJsonWebKeyECRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKty

`func (o *OAuth2ClientJsonWebKeyECRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *OAuth2ClientJsonWebKeyECRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *OAuth2ClientJsonWebKeyECRequest) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *OAuth2ClientJsonWebKeyECRequest) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetX

`func (o *OAuth2ClientJsonWebKeyECRequest) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *OAuth2ClientJsonWebKeyECRequest) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *OAuth2ClientJsonWebKeyECRequest) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *OAuth2ClientJsonWebKeyECRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *OAuth2ClientJsonWebKeyECRequest) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *OAuth2ClientJsonWebKeyECRequest) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *OAuth2ClientJsonWebKeyECRequest) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *OAuth2ClientJsonWebKeyECRequest) HasY() bool`

HasY returns a boolean if a field has been set.

### GetKid

`func (o *OAuth2ClientJsonWebKeyECRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *OAuth2ClientJsonWebKeyECRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *OAuth2ClientJsonWebKeyECRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *OAuth2ClientJsonWebKeyECRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *OAuth2ClientJsonWebKeyECRequest) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *OAuth2ClientJsonWebKeyECRequest) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *OAuth2ClientJsonWebKeyECRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OAuth2ClientJsonWebKeyECRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OAuth2ClientJsonWebKeyECRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OAuth2ClientJsonWebKeyECRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


