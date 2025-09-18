# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassCode

`func (o *Token) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *Token) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *Token) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *Token) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


