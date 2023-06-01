# JsonWebKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**E** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**KeyOps** | Pointer to **[]string** |  | [optional] 
**Kid** | Pointer to **string** |  | [optional] 
**Kty** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**N** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Use** | Pointer to **string** |  | [optional] 
**X5c** | Pointer to **[]string** |  | [optional] 
**X5t** | Pointer to **string** |  | [optional] 
**X5tS256** | Pointer to **string** |  | [optional] 
**X5u** | Pointer to **string** |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

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

### GetAlg

`func (o *JsonWebKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JsonWebKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JsonWebKey) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *JsonWebKey) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

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

### GetKeyOps

`func (o *JsonWebKey) GetKeyOps() []string`

GetKeyOps returns the KeyOps field if non-nil, zero value otherwise.

### GetKeyOpsOk

`func (o *JsonWebKey) GetKeyOpsOk() (*[]string, bool)`

GetKeyOpsOk returns a tuple with the KeyOps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyOps

`func (o *JsonWebKey) SetKeyOps(v []string)`

SetKeyOps sets KeyOps field to given value.

### HasKeyOps

`func (o *JsonWebKey) HasKeyOps() bool`

HasKeyOps returns a boolean if a field has been set.

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

### GetStatus

`func (o *JsonWebKey) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JsonWebKey) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JsonWebKey) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JsonWebKey) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### GetX5t

`func (o *JsonWebKey) GetX5t() string`

GetX5t returns the X5t field if non-nil, zero value otherwise.

### GetX5tOk

`func (o *JsonWebKey) GetX5tOk() (*string, bool)`

GetX5tOk returns a tuple with the X5t field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5t

`func (o *JsonWebKey) SetX5t(v string)`

SetX5t sets X5t field to given value.

### HasX5t

`func (o *JsonWebKey) HasX5t() bool`

HasX5t returns a boolean if a field has been set.

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

### GetX5u

`func (o *JsonWebKey) GetX5u() string`

GetX5u returns the X5u field if non-nil, zero value otherwise.

### GetX5uOk

`func (o *JsonWebKey) GetX5uOk() (*string, bool)`

GetX5uOk returns a tuple with the X5u field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX5u

`func (o *JsonWebKey) SetX5u(v string)`

SetX5u sets X5u field to given value.

### HasX5u

`func (o *JsonWebKey) HasX5u() bool`

HasX5u returns a boolean if a field has been set.

### GetLinks

`func (o *JsonWebKey) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *JsonWebKey) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *JsonWebKey) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *JsonWebKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


