# JsonWebKeyEC

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | Pointer to **string** | The public x coordinate for the elliptic curve point | [optional] 
**Y** | Pointer to **string** | The public y coordinate for the elliptic curve point | [optional] 

## Methods

### NewJsonWebKeyEC

`func NewJsonWebKeyEC() *JsonWebKeyEC`

NewJsonWebKeyEC instantiates a new JsonWebKeyEC object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonWebKeyECWithDefaults

`func NewJsonWebKeyECWithDefaults() *JsonWebKeyEC`

NewJsonWebKeyECWithDefaults instantiates a new JsonWebKeyEC object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *JsonWebKeyEC) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *JsonWebKeyEC) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *JsonWebKeyEC) SetX(v string)`

SetX sets X field to given value.

### HasX

`func (o *JsonWebKeyEC) HasX() bool`

HasX returns a boolean if a field has been set.

### GetY

`func (o *JsonWebKeyEC) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *JsonWebKeyEC) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *JsonWebKeyEC) SetY(v string)`

SetY sets Y field to given value.

### HasY

`func (o *JsonWebKeyEC) HasY() bool`

HasY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


