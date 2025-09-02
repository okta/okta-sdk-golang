# SecurityEventTokenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Describes the error &gt; **Note:** SET claim fields with underscores (snake case) are presented in camelcase. For example, &#x60;previous_status&#x60; appears as &#x60;previousStatus&#x60;.  | [optional] 
**Err** | Pointer to **string** | A code that describes the category of the error | [optional] 

## Methods

### NewSecurityEventTokenError

`func NewSecurityEventTokenError() *SecurityEventTokenError`

NewSecurityEventTokenError instantiates a new SecurityEventTokenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventTokenErrorWithDefaults

`func NewSecurityEventTokenErrorWithDefaults() *SecurityEventTokenError`

NewSecurityEventTokenErrorWithDefaults instantiates a new SecurityEventTokenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityEventTokenError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityEventTokenError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityEventTokenError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityEventTokenError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErr

`func (o *SecurityEventTokenError) GetErr() string`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *SecurityEventTokenError) GetErrOk() (*string, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *SecurityEventTokenError) SetErr(v string)`

SetErr sets Err field to given value.

### HasErr

`func (o *SecurityEventTokenError) HasErr() bool`

HasErr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


