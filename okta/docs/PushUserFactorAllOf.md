# PushUserFactorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**FactorResult** | Pointer to [**FactorResultType**](FactorResultType.md) |  | [optional] 
**Profile** | Pointer to [**PushUserFactorProfile**](PushUserFactorProfile.md) |  | [optional] 

## Methods

### NewPushUserFactorAllOf

`func NewPushUserFactorAllOf() *PushUserFactorAllOf`

NewPushUserFactorAllOf instantiates a new PushUserFactorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushUserFactorAllOfWithDefaults

`func NewPushUserFactorAllOfWithDefaults() *PushUserFactorAllOf`

NewPushUserFactorAllOfWithDefaults instantiates a new PushUserFactorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *PushUserFactorAllOf) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PushUserFactorAllOf) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PushUserFactorAllOf) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PushUserFactorAllOf) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorResult

`func (o *PushUserFactorAllOf) GetFactorResult() FactorResultType`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *PushUserFactorAllOf) GetFactorResultOk() (*FactorResultType, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *PushUserFactorAllOf) SetFactorResult(v FactorResultType)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *PushUserFactorAllOf) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetProfile

`func (o *PushUserFactorAllOf) GetProfile() PushUserFactorProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PushUserFactorAllOf) GetProfileOk() (*PushUserFactorProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PushUserFactorAllOf) SetProfile(v PushUserFactorProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PushUserFactorAllOf) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


