# Csr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Csr** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Kty** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCsr

`func NewCsr() *Csr`

NewCsr instantiates a new Csr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsrWithDefaults

`func NewCsrWithDefaults() *Csr`

NewCsrWithDefaults instantiates a new Csr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Csr) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Csr) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Csr) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Csr) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCsr

`func (o *Csr) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *Csr) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *Csr) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *Csr) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetId

`func (o *Csr) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Csr) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Csr) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Csr) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKty

`func (o *Csr) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *Csr) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *Csr) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *Csr) HasKty() bool`

HasKty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


