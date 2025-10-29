# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**E** | Pointer to **string** | RSA key value (public exponent) for Key binding | [optional] [readonly] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the certificate expires | [optional] [readonly] 
**Kid** | Pointer to **string** | Unique identifier for the certificate | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the certificate&#39;s keypair. Valid value: &#x60;RSA&#x60; | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**N** | Pointer to **string** | RSA modulus value that is used by both the public and private keys and provides a link between them | [optional] 
**Use** | Pointer to **string** | Acceptable use of the certificate. Valid value: &#x60;sig&#x60; | [optional] [readonly] 
**X5c** | Pointer to **[]string** | X.509 certificate chain that contains a chain of one or more certificates | [optional] [readonly] 
**X5tS256** | Pointer to **string** | X.509 certificate SHA-256 thumbprint, which is the base64url-encoded SHA-256 thumbprint (digest) of the DER encoding of an X.509 certificate | [optional] [readonly] 

## Methods

### NewJsonWebKey

`func NewJsonWebKey() *JsonWebKey`

NewJsonWebKey instantiates a new JsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyWithDefaults

`func NewJsonWebKeyWithDefaults() *JsonWebKey`

NewJsonWebKeyWithDefaults instantiates a new JsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *JsonWebKey) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *JsonWebKey) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *JsonWebKey) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *JsonWebKey) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetE

`func (o *JsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *JsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *JsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *JsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetExpiresAt

`func (o *JsonWebKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *JsonWebKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *JsonWebKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *JsonWebKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetKid

`func (o *JsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *JsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *JsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *JsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLastUpdated

`func (o *JsonWebKey) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *JsonWebKey) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *JsonWebKey) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *JsonWebKey) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetN

`func (o *JsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *JsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *JsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *JsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *JsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *JsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX5c

`func (o *JsonWebKey) GetX5c() []string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *JsonWebKey) GetX5cOk() (*[]string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *JsonWebKey) SetX5c(v []string)`

SetX5c sets X5c field to given value.

### HasX5c

`func (o *JsonWebKey) HasX5c() bool`

HasX5c returns a boolean if a field has been set.

### GetX5tS256

`func (o *JsonWebKey) GetX5tS256() string`

GetX5tS256 returns the X5tS256 field if non-nil, zero value otherwise.

### GetX5tS256Ok

`func (o *JsonWebKey) GetX5tS256Ok() (*string, bool)`

GetX5tS256Ok returns a tuple with the X5tS256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5tS256

`func (o *JsonWebKey) SetX5tS256(v string)`

SetX5tS256 sets X5tS256 field to given value.

### HasX5tS256

`func (o *JsonWebKey) HasX5tS256() bool`

HasX5tS256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


