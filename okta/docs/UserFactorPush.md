# UserFactorPush

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | Timestamp indicating when the Factor verification attempt expires | [optional] [readonly] 
**FactorResult** | Pointer to [**UserFactorResultType**](UserFactorResultType.md) |  | [optional] 
**Profile** | Pointer to [**UserFactorPushProfile**](UserFactorPushProfile.md) |  | [optional] 

## Methods

### NewUserFactorPush

`func NewUserFactorPush() *UserFactorPush`

NewUserFactorPush instantiates a new UserFactorPush object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorPushWithDefaults

`func NewUserFactorPushWithDefaults() *UserFactorPush`

NewUserFactorPushWithDefaults instantiates a new UserFactorPush object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *UserFactorPush) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UserFactorPush) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UserFactorPush) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UserFactorPush) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFactorResult

`func (o *UserFactorPush) GetFactorResult() UserFactorResultType`

GetFactorResult returns the FactorResult field if non-nil, zero value otherwise.

### GetFactorResultOk

`func (o *UserFactorPush) GetFactorResultOk() (*UserFactorResultType, bool)`

GetFactorResultOk returns a tuple with the FactorResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorResult

`func (o *UserFactorPush) SetFactorResult(v UserFactorResultType)`

SetFactorResult sets FactorResult field to given value.

### HasFactorResult

`func (o *UserFactorPush) HasFactorResult() bool`

HasFactorResult returns a boolean if a field has been set.

### GetProfile

`func (o *UserFactorPush) GetProfile() UserFactorPushProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserFactorPush) GetProfileOk() (*UserFactorPushProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserFactorPush) SetProfile(v UserFactorPushProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserFactorPush) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


