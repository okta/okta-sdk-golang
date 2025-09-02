# IdPCsr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the object was created | [optional] [readonly] 
**Csr** | Pointer to **string** | Base64-encoded CSR in DER format | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier for the CSR | [optional] [readonly] 
**Kty** | Pointer to **string** | Cryptographic algorithm family for the CSR&#39;s keypair | [optional] 
**Links** | Pointer to [**IdPCsrLinks**](IdPCsrLinks.md) |  | [optional] 

## Methods

### NewIdPCsr

`func NewIdPCsr() *IdPCsr`

NewIdPCsr instantiates a new IdPCsr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdPCsrWithDefaults

`func NewIdPCsrWithDefaults() *IdPCsr`

NewIdPCsrWithDefaults instantiates a new IdPCsr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *IdPCsr) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IdPCsr) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IdPCsr) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *IdPCsr) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCsr

`func (o *IdPCsr) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *IdPCsr) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *IdPCsr) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *IdPCsr) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetId

`func (o *IdPCsr) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdPCsr) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdPCsr) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdPCsr) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKty

`func (o *IdPCsr) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *IdPCsr) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *IdPCsr) SetKty(v string)`

SetKty sets Kty field to given value.

### HasKty

`func (o *IdPCsr) HasKty() bool`

HasKty returns a boolean if a field has been set.

### GetLinks

`func (o *IdPCsr) GetLinks() IdPCsrLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdPCsr) GetLinksOk() (*IdPCsrLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdPCsr) SetLinks(v IdPCsrLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdPCsr) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


