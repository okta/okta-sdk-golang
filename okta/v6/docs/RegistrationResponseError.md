# RegistrationResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorSummary** | Pointer to **string** | Human-readable summary of one or more errors | [optional] 
**ErrorCauses** | Pointer to [**[]RegistrationResponseErrorErrorCausesInner**](RegistrationResponseErrorErrorCausesInner.md) |  | [optional] 

## Methods

### NewRegistrationResponseError

`func NewRegistrationResponseError() *RegistrationResponseError`

NewRegistrationResponseError instantiates a new RegistrationResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationResponseErrorWithDefaults

`func NewRegistrationResponseErrorWithDefaults() *RegistrationResponseError`

NewRegistrationResponseErrorWithDefaults instantiates a new RegistrationResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorSummary

`func (o *RegistrationResponseError) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *RegistrationResponseError) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *RegistrationResponseError) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *RegistrationResponseError) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetErrorCauses

`func (o *RegistrationResponseError) GetErrorCauses() []RegistrationResponseErrorErrorCausesInner`

GetErrorCauses returns the ErrorCauses field if non-nil, zero value otherwise.

### GetErrorCausesOk

`func (o *RegistrationResponseError) GetErrorCausesOk() (*[]RegistrationResponseErrorErrorCausesInner, bool)`

GetErrorCausesOk returns a tuple with the ErrorCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCauses

`func (o *RegistrationResponseError) SetErrorCauses(v []RegistrationResponseErrorErrorCausesInner)`

SetErrorCauses sets ErrorCauses field to given value.

### HasErrorCauses

`func (o *RegistrationResponseError) HasErrorCauses() bool`

HasErrorCauses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


