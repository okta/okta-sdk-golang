# OAuth2ClientJsonWebKeyECRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crv** | **string** | The cryptographic curve used with the key | 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**X** | **string** | The public x coordinate for the elliptic curve point | 
**Y** | **string** | The public y coordinate for the elliptic curve point | 

## Methods

### NewOAuth2ClientJsonWebKeyECRequest

`func NewOAuth2ClientJsonWebKeyECRequest(crv string, x string, y string, ) *OAuth2ClientJsonWebKeyECRequest`

NewOAuth2ClientJsonWebKeyECRequest instantiates a new OAuth2ClientJsonWebKeyECRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyECRequestWithDefaults

`func NewOAuth2ClientJsonWebKeyECRequestWithDefaults() *OAuth2ClientJsonWebKeyECRequest`

NewOAuth2ClientJsonWebKeyECRequestWithDefaults instantiates a new OAuth2ClientJsonWebKeyECRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *OAuth2ClientJsonWebKeyECRequest) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *OAuth2ClientJsonWebKeyECRequest) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *OAuth2ClientJsonWebKeyECRequest) SetCrv(v string)`

SetCrv sets Crv field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


