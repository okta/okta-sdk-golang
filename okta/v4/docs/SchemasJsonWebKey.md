# SchemasJsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**Kid** | Pointer to **string** | The unique identifier of the key | [optional] 
**Kty** | Pointer to **string** | The type of public key | [optional] 
**Status** | Pointer to **string** | The status of the public key | [optional] 
**Use** | Pointer to **string** | The intended use of the public key | [optional] 

## Methods

### NewSchemasJsonWebKey

`func NewSchemasJsonWebKey() *SchemasJsonWebKey`

NewSchemasJsonWebKey instantiates a new SchemasJsonWebKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasJsonWebKeyWithDefaults

`func NewSchemasJsonWebKeyWithDefaults() *SchemasJsonWebKey`

NewSchemasJsonWebKeyWithDefaults instantiates a new SchemasJsonWebKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *SchemasJsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *SchemasJsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *SchemasJsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *SchemasJsonWebKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetKid

`func (o *SchemasJsonWebKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *SchemasJsonWebKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *SchemasJsonWebKey) SetKid(v string)`

SetKid sets Kid field to given value.

### HasKid

`func (o *SchemasJsonWebKey) HasKid() bool`

HasKid returns a boolean if a field has been set.

### GetKty

`func (o *SchemasJsonWebKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *SchemasJsonWebKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *SchemasJsonWebKey) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *SchemasJsonWebKey) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetStatus

`func (o *SchemasJsonWebKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SchemasJsonWebKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SchemasJsonWebKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SchemasJsonWebKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUse

`func (o *SchemasJsonWebKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *SchemasJsonWebKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *SchemasJsonWebKey) SetUse(v string)`

SetUse sets Use field to given value.

### HasUse

`func (o *SchemasJsonWebKey) HasUse() bool`

HasUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


