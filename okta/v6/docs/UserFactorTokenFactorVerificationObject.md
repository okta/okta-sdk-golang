# UserFactorTokenFactorVerificationObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPassCode** | Pointer to **string** | OTP for the next time window | [optional] 
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 

## Methods

### NewUserFactorTokenFactorVerificationObject

`func NewUserFactorTokenFactorVerificationObject() *UserFactorTokenFactorVerificationObject`

NewUserFactorTokenFactorVerificationObject instantiates a new UserFactorTokenFactorVerificationObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorTokenFactorVerificationObjectWithDefaults

`func NewUserFactorTokenFactorVerificationObjectWithDefaults() *UserFactorTokenFactorVerificationObject`

NewUserFactorTokenFactorVerificationObjectWithDefaults instantiates a new UserFactorTokenFactorVerificationObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPassCode

`func (o *UserFactorTokenFactorVerificationObject) GetNextPassCode() string`

GetNextPassCode returns the NextPassCode field if non-nil, zero value otherwise.

### GetNextPassCodeOk

`func (o *UserFactorTokenFactorVerificationObject) GetNextPassCodeOk() (*string, bool)`

GetNextPassCodeOk returns a tuple with the NextPassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPassCode

`func (o *UserFactorTokenFactorVerificationObject) SetNextPassCode(v string)`

SetNextPassCode sets NextPassCode field to given value.

### HasNextPassCode

`func (o *UserFactorTokenFactorVerificationObject) HasNextPassCode() bool`

HasNextPassCode returns a boolean if a field has been set.

### GetPassCode

`func (o *UserFactorTokenFactorVerificationObject) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *UserFactorTokenFactorVerificationObject) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *UserFactorTokenFactorVerificationObject) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *UserFactorTokenFactorVerificationObject) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


