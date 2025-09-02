# ResourceServerJsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | The key exponent of a RSA key | [optional] 
**Kid** | Pointer to **string** | The unique identifier of the key | [optional] 
**Kty** | Pointer to **string** | The type of public key | [optional] 
**N** | Pointer to **string** | The modulus of the RSA key | [optional] 
**Status** | Pointer to **string** | The status of the public key | [optional] 
**Use** | Pointer to **string** | The intended use of the public key | [optional] 

## Methods

### NewResourceServerJsonWebKey

`func NewResourceServerJsonWebKey() *ResourceServerJsonWebKey`

NewResourceServerJsonWebKey instantiates a new ResourceServerJsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceServerJsonWebKeyWithDefaults

`func NewResourceServerJsonWebKeyWithDefaults() *ResourceServerJsonWebKey`

NewResourceServerJsonWebKeyWithDefaults instantiates a new ResourceServerJsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *ResourceServerJsonWebKey) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *ResourceServerJsonWebKey) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *ResourceServerJsonWebKey) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *ResourceServerJsonWebKey) HasE() bool`

HasE returns a boolean if a field has been set.

### GetKid

`func (o *ResourceServerJsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *ResourceServerJsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *ResourceServerJsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *ResourceServerJsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *ResourceServerJsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *ResourceServerJsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *ResourceServerJsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *ResourceServerJsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetN

`func (o *ResourceServerJsonWebKey) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ResourceServerJsonWebKey) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ResourceServerJsonWebKey) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *ResourceServerJsonWebKey) HasN() bool`

HasN returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceServerJsonWebKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceServerJsonWebKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceServerJsonWebKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceServerJsonWebKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *ResourceServerJsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *ResourceServerJsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *ResourceServerJsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *ResourceServerJsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


