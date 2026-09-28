# UserFactorSignedNonceProfileKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Crv** | Pointer to **string** | EC curve name (present only for EC keys) | [optional] 
**E** | Pointer to **string** | RSA public exponent (present only for RSA keys) | [optional] 
**JwkType** | Pointer to **string** | Purpose of the key | [optional] 
**Kid** | Pointer to **string** | Key ID | [optional] 
**Kty** | Pointer to **string** | Key type | [optional] 
**N** | Pointer to **string** | RSA modulus (present only for RSA keys) | [optional] 
**Use** | Pointer to **string** | Key usage | [optional] 
**X** | Pointer to **string** | EC x-coordinate (present only for EC keys) | [optional] 
**X5c** | Pointer to **[]string** | X.509 certificate chain | [optional] 
**Y** | Pointer to **string** | EC y-coordinate (present only for EC keys) | [optional] 

## Methods

### NewUserFactorSignedNonceProfileKey

`func NewUserFactorSignedNonceProfileKey() *UserFactorSignedNonceProfileKey`

NewUserFactorSignedNonceProfileKey instantiates a new UserFactorSignedNonceProfileKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSignedNonceProfileKeyWithDefaults

`func NewUserFactorSignedNonceProfileKeyWithDefaults() *UserFactorSignedNonceProfileKey`

NewUserFactorSignedNonceProfileKeyWithDefaults instantiates a new UserFactorSignedNonceProfileKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCrv

`func (o *UserFactorSignedNonceProfileKey) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *UserFactorSignedNonceProfileKey) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *UserFactorSignedNonceProfileKey) SetCrv(v string)`

SetCrv sets Crv field to given value.

### HasCrv

`func (o *UserFactorSignedNonceProfileKey) HasCrv() bool`

HasCrv returns a boolean if a field has been set.

### GetE

`func (o *UserFactorSignedNonceProfileKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *UserFactorSignedNonceProfileKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *UserFactorSignedNonceProfileKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *UserFactorSignedNonceProfileKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetJwkType

`func (o *UserFactorSignedNonceProfileKey) GetJwkType() string`

GetJwkType returns the JwkType field if non-nil, zero value otherwise.

### GetJwkTypeOk

`func (o *UserFactorSignedNonceProfileKey) GetJwkTypeOk() (*string, bool)`

GetJwkTypeOk returns a tuple with the JwkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwkType

`func (o *UserFactorSignedNonceProfileKey) SetJwkType(v string)`

SetJwkType sets JwkType field to given value.

### HasJwkType

`func (o *UserFactorSignedNonceProfileKey) HasJwkType() bool`

HasJwkType returns a boolean if a field has been set.

### GetKid

`func (o *UserFactorSignedNonceProfileKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *UserFactorSignedNonceProfileKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *UserFactorSignedNonceProfileKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *UserFactorSignedNonceProfileKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *UserFactorSignedNonceProfileKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *UserFactorSignedNonceProfileKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *UserFactorSignedNonceProfileKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *UserFactorSignedNonceProfileKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *UserFactorSignedNonceProfileKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *UserFactorSignedNonceProfileKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *UserFactorSignedNonceProfileKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *UserFactorSignedNonceProfileKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *UserFactorSignedNonceProfileKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *UserFactorSignedNonceProfileKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *UserFactorSignedNonceProfileKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *UserFactorSignedNonceProfileKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX

`func (o *UserFactorSignedNonceProfileKey) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *UserFactorSignedNonceProfileKey) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *UserFactorSignedNonceProfileKey) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *UserFactorSignedNonceProfileKey) HasX() bool`

HasX returns a boolean if a field has been set.

### GetX5c

`func (o *UserFactorSignedNonceProfileKey) GetX5c() []string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *UserFactorSignedNonceProfileKey) GetX5cOk() (*[]string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *UserFactorSignedNonceProfileKey) SetX5c(v []string)`

SetX5c sets X5c field to given value.

### HasX5c

`func (o *UserFactorSignedNonceProfileKey) HasX5c() bool`

HasX5c returns a boolean if a field has been set.

### GetY

`func (o *UserFactorSignedNonceProfileKey) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *UserFactorSignedNonceProfileKey) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *UserFactorSignedNonceProfileKey) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *UserFactorSignedNonceProfileKey) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


