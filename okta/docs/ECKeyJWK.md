# ECKeyJWK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crv** | **string** |  | 
**Kid** | **string** | The unique identifier of the key | 
**Kty** | **string** | The type of public key | 
**Use** | **string** | The intended use for the key. The ECKeyJWK is always &#x60;enc&#x60; because Okta uses it to encrypt requests to Yubico. | 
**X** | **string** | The public x coordinate for the elliptic curve point | 
**Y** | **string** | The public y coordinate for the elliptic curve point | 

## Methods

### NewECKeyJWK

`func NewECKeyJWK(crv string, kid string, kty string, use string, x string, y string, ) *ECKeyJWK`

NewECKeyJWK instantiates a new ECKeyJWK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECKeyJWKWithDefaults

`func NewECKeyJWKWithDefaults() *ECKeyJWK`

NewECKeyJWKWithDefaults instantiates a new ECKeyJWK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *ECKeyJWK) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *ECKeyJWK) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *ECKeyJWK) SetCrv(v string)`

SetCrv sets Crv field to given value.


### GetKid

`func (o *ECKeyJWK) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *ECKeyJWK) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *ECKeyJWK) SetKid(v string)`

SetKid sets Kid field to given value.


### GetKty

`func (o *ECKeyJWK) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *ECKeyJWK) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *ECKeyJWK) SetKty(v string)`

SetKty sets Kty field to given value.


### GetUse

`func (o *ECKeyJWK) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *ECKeyJWK) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *ECKeyJWK) SetUse(v string)`

SetUse sets Use field to given value.


### GetX

`func (o *ECKeyJWK) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *ECKeyJWK) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *ECKeyJWK) SetX(v string)`

SetX sets X field to given value.


### GetY

`func (o *ECKeyJWK) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *ECKeyJWK) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *ECKeyJWK) SetY(v string)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


