# UserFactorTokenAllOfVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 
**NextPassCode** | Pointer to **int32** | OTP for the next time window | [optional] 

## Methods

### NewUserFactorTokenAllOfVerify

`func NewUserFactorTokenAllOfVerify() *UserFactorTokenAllOfVerify`

NewUserFactorTokenAllOfVerify instantiates a new UserFactorTokenAllOfVerify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorTokenAllOfVerifyWithDefaults

`func NewUserFactorTokenAllOfVerifyWithDefaults() *UserFactorTokenAllOfVerify`

NewUserFactorTokenAllOfVerifyWithDefaults instantiates a new UserFactorTokenAllOfVerify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassCode

`func (o *UserFactorTokenAllOfVerify) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *UserFactorTokenAllOfVerify) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *UserFactorTokenAllOfVerify) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *UserFactorTokenAllOfVerify) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.

### GetNextPassCode

`func (o *UserFactorTokenAllOfVerify) GetNextPassCode() int32`

GetNextPassCode returns the NextPassCode field if non-nil, zero value otherwise.

### GetNextPassCodeOk

`func (o *UserFactorTokenAllOfVerify) GetNextPassCodeOk() (*int32, bool)`

GetNextPassCodeOk returns a tuple with the NextPassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPassCode

`func (o *UserFactorTokenAllOfVerify) SetNextPassCode(v int32)`

SetNextPassCode sets NextPassCode field to given value.

### HasNextPassCode

`func (o *UserFactorTokenAllOfVerify) HasNextPassCode() bool`

HasNextPassCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


