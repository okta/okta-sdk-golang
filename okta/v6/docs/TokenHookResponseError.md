# TokenHookResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorSummary** | Pointer to **string** | Human-readable summary of the error. If the error object doesn&#39;t include the &#x60;errorSummary&#x60; property defined, the following common default message is returned to the end user: &#x60;The callback service returned an error&#x60;. | [optional] 

## Methods

### NewTokenHookResponseError

`func NewTokenHookResponseError() *TokenHookResponseError`

NewTokenHookResponseError instantiates a new TokenHookResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHookResponseErrorWithDefaults

`func NewTokenHookResponseErrorWithDefaults() *TokenHookResponseError`

NewTokenHookResponseErrorWithDefaults instantiates a new TokenHookResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorSummary

`func (o *TokenHookResponseError) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *TokenHookResponseError) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *TokenHookResponseError) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *TokenHookResponseError) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


