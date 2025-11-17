# OAuth2ClientJsonWebKeyECResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crv** | **string** | The cryptographic curve used with the key | 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**X** | **string** | The public x coordinate for the elliptic curve point | 
**Y** | **string** | The public y coordinate for the elliptic curve point | 

## Methods

### NewOAuth2ClientJsonWebKeyECResponse

`func NewOAuth2ClientJsonWebKeyECResponse(crv string, x string, y string, ) *OAuth2ClientJsonWebKeyECResponse`

NewOAuth2ClientJsonWebKeyECResponse instantiates a new OAuth2ClientJsonWebKeyECResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientJsonWebKeyECResponseWithDefaults

`func NewOAuth2ClientJsonWebKeyECResponseWithDefaults() *OAuth2ClientJsonWebKeyECResponse`

NewOAuth2ClientJsonWebKeyECResponseWithDefaults instantiates a new OAuth2ClientJsonWebKeyECResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *OAuth2ClientJsonWebKeyECResponse) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *OAuth2ClientJsonWebKeyECResponse) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *OAuth2ClientJsonWebKeyECResponse) SetCrv(v string)`

SetCrv sets Crv field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


