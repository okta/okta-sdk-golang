# Duration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 

## Methods

### NewDuration

`func NewDuration() *Duration`

NewDuration instantiates a new Duration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDurationWithDefaults

`func NewDurationWithDefaults() *Duration`

NewDurationWithDefaults instantiates a new Duration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *Duration) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Duration) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Duration) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Duration) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetUnit

`func (o *Duration) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Duration) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Duration) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Duration) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


