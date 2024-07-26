# UserFactorActivatePush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | Timestamp when the Factor verification attempt expires | [optional] [readonly] 
**FactorResult** | Pointer to **string** | Result of a Factor verification | [optional] 

## Methods

### NewUserFactorActivatePush

`func NewUserFactorActivatePush() *UserFactorActivatePush`

NewUserFactorActivatePush instantiates a new UserFactorActivatePush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorActivatePushWithDefaults

`func NewUserFactorActivatePushWithDefaults() *UserFactorActivatePush`

NewUserFactorActivatePushWithDefaults instantiates a new UserFactorActivatePush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *UserFactorActivatePush) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserFactorActivatePush) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserFactorActivatePush) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UserFactorActivatePush) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorResult

`func (o *UserFactorActivatePush) GetFactorResult() string`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *UserFactorActivatePush) GetFactorResultOk() (*string, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *UserFactorActivatePush) SetFactorResult(v string)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *UserFactorActivatePush) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


