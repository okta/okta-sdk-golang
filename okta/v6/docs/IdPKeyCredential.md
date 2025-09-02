# IdPKeyCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**E** | Pointer to **string** | The exponent value for the RSA public key | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the object expires | [optional] [readonly] 
**Kid** | Pointer to **string** | Unique identifier for the key | [optional] 
**Kty** | Pointer to **string** | Identifies the cryptographic algorithm family used with the key | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the object was last updated | [optional] [readonly] 
**N** | Pointer to **string** | The modulus value for the RSA public key | [optional] 
**Use** | Pointer to **string** | Intended use of the public key | [optional] 
**X5c** | Pointer to **[]string** | Base64-encoded X.509 certificate chain with DER encoding | [optional] 
**X5tS256** | Pointer to **string** | Base64url-encoded SHA-256 thumbprint of the DER encoding of an X.509 certificate | [optional] 

## Methods

### NewIdPKeyCredential

`func NewIdPKeyCredential() *IdPKeyCredential`

NewIdPKeyCredential instantiates a new IdPKeyCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdPKeyCredentialWithDefaults

`func NewIdPKeyCredentialWithDefaults() *IdPKeyCredential`

NewIdPKeyCredentialWithDefaults instantiates a new IdPKeyCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IdPKeyCredential) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdPKeyCredential) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdPKeyCredential) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdPKeyCredential) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetE

`func (o *IdPKeyCredential) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *IdPKeyCredential) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *IdPKeyCredential) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *IdPKeyCredential) HasE() bool`

HasE returns a boolean if a field has been set.

### GetExpiresAt

`func (o *IdPKeyCredential) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *IdPKeyCredential) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *IdPKeyCredential) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *IdPKeyCredential) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetKid

`func (o *IdPKeyCredential) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *IdPKeyCredential) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *IdPKeyCredential) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *IdPKeyCredential) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *IdPKeyCredential) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *IdPKeyCredential) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *IdPKeyCredential) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *IdPKeyCredential) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLastUpdated

`func (o *IdPKeyCredential) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IdPKeyCredential) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IdPKeyCredential) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *IdPKeyCredential) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetN

`func (o *IdPKeyCredential) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *IdPKeyCredential) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *IdPKeyCredential) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *IdPKeyCredential) HasN() bool`

HasN returns a boolean if a field has been set.

### GetUse

`func (o *IdPKeyCredential) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *IdPKeyCredential) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *IdPKeyCredential) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *IdPKeyCredential) HasUse() bool`

HasUse returns a boolean if a field has been set.

### GetX5c

`func (o *IdPKeyCredential) GetX5c() []string`

GetX5c returns the X5c field if non-nil, zero value otherwise.

### GetX5cOk

`func (o *IdPKeyCredential) GetX5cOk() (*[]string, bool)`

GetX5cOk returns a tuple with the X5c field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5c

`func (o *IdPKeyCredential) SetX5c(v []string)`

SetX5c sets X5c field to given value.

### HasX5c

`func (o *IdPKeyCredential) HasX5c() bool`

HasX5c returns a boolean if a field has been set.

### GetX5tS256

`func (o *IdPKeyCredential) GetX5tS256() string`

GetX5tS256 returns the X5tS256 field if non-nil, zero value otherwise.

### GetX5tS256Ok

`func (o *IdPKeyCredential) GetX5tS256Ok() (*string, bool)`

GetX5tS256Ok returns a tuple with the X5tS256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5tS256

`func (o *IdPKeyCredential) SetX5tS256(v string)`

SetX5tS256 sets X5tS256 field to given value.

### HasX5tS256

`func (o *IdPKeyCredential) HasX5tS256() bool`

HasX5tS256 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


