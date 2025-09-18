# AddJwkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** | Algorithm used in the key | [optional] 
**Use** | Pointer to **string** | Acceptable use of the JSON Web Key | [optional] 
**Kid** | Pointer to **NullableString** | Unique identifier of the JSON Web Key in the OAUth 2.0 client&#39;s JWKS | [optional] 
**Status** | Pointer to **string** | Status of the OAuth 2.0 client JSON Web Key | [optional] [default to "ACTIVE"]
**E** | Pointer to **string** | RSA key value (exponent) for key binding | [optional] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s key pair | [optional] 
**N** | Pointer to **string** | RSA key value (modulus) for key binding | [optional] 
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 

## Methods

### NewAddJwkRequest

`func NewAddJwkRequest() *AddJwkRequest`

NewAddJwkRequest instantiates a new AddJwkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddJwkRequestWithDefaults

`func NewAddJwkRequestWithDefaults() *AddJwkRequest`

NewAddJwkRequestWithDefaults instantiates a new AddJwkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *AddJwkRequest) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *AddJwkRequest) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *AddJwkRequest) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *AddJwkRequest) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetUse

`func (o *AddJwkRequest) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *AddJwkRequest) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *AddJwkRequest) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *AddJwkRequest) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetKid

`func (o *AddJwkRequest) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *AddJwkRequest) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *AddJwkRequest) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *AddJwkRequest) HasKid() bool`

HasKid returns a boolean if a field has been set.

### SetKidNil

`func (o *AddJwkRequest) SetKidNil(b bool)`

 SetKidNil sets the value for Kid to be an explicit nil

### UnsetKid
`func (o *AddJwkRequest) UnsetKid()`

UnsetKid ensures that no value is present for Kid, not even an explicit nil
### GetStatus

`func (o *AddJwkRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddJwkRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddJwkRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddJwkRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetE

`func (o *AddJwkRequest) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *AddJwkRequest) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *AddJwkRequest) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *AddJwkRequest) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKty

`func (o *AddJwkRequest) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *AddJwkRequest) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *AddJwkRequest) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *AddJwkRequest) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *AddJwkRequest) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AddJwkRequest) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AddJwkRequest) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *AddJwkRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetX

`func (o *AddJwkRequest) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *AddJwkRequest) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *AddJwkRequest) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *AddJwkRequest) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *AddJwkRequest) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *AddJwkRequest) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *AddJwkRequest) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *AddJwkRequest) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


