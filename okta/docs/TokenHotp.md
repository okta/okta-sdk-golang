# TokenHotp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PassCode** | Pointer to **string** | OTP for the current time window | [optional] 

## Methods

### NewTokenHotp

`func NewTokenHotp() *TokenHotp`

NewTokenHotp instantiates a new TokenHotp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHotpWithDefaults

`func NewTokenHotpWithDefaults() *TokenHotp`

NewTokenHotpWithDefaults instantiates a new TokenHotp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassCode

`func (o *TokenHotp) GetPassCode() string`

GetPassCode returns the PassCode field if non-nil, zero value otherwise.

### GetPassCodeOk

`func (o *TokenHotp) GetPassCodeOk() (*string, bool)`

GetPassCodeOk returns a tuple with the PassCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCode

`func (o *TokenHotp) SetPassCode(v string)`

SetPassCode sets PassCode field to given value.

### HasPassCode

`func (o *TokenHotp) HasPassCode() bool`

HasPassCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


