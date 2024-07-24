# JsonWebKeyRsa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**E** | Pointer to **string** | The key exponent of a RSA key | [optional] 
**N** | Pointer to **string** | The modulus of the RSA key | [optional] 

## Methods

### NewJsonWebKeyRsa

`func NewJsonWebKeyRsa() *JsonWebKeyRsa`

NewJsonWebKeyRsa instantiates a new JsonWebKeyRsa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyRsaWithDefaults

`func NewJsonWebKeyRsaWithDefaults() *JsonWebKeyRsa`

NewJsonWebKeyRsaWithDefaults instantiates a new JsonWebKeyRsa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetE

`func (o *JsonWebKeyRsa) GetE() string`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *JsonWebKeyRsa) GetEOk() (*string, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *JsonWebKeyRsa) SetE(v string)`

SetE sets E field to given value.

### HasE

`func (o *JsonWebKeyRsa) HasE() bool`

HasE returns a boolean if a field has been set.

### GetN

`func (o *JsonWebKeyRsa) GetN() string`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *JsonWebKeyRsa) GetNOk() (*string, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *JsonWebKeyRsa) SetN(v string)`

SetN sets N field to given value.

### HasN

`func (o *JsonWebKeyRsa) HasN() bool`

HasN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


