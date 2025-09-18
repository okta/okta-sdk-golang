# SAMLPayLoadDataAssertionSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameId** | Pointer to **string** | The unique identifier of the user | [optional] 
**NameFormat** | Pointer to **string** | Indicates how to interpret the attribute name | [optional] 
**Confirmation** | Pointer to [**SAMLPayLoadDataAssertionSubjectConfirmation**](SAMLPayLoadDataAssertionSubjectConfirmation.md) |  | [optional] 

## Methods

### NewSAMLPayLoadDataAssertionSubject

`func NewSAMLPayLoadDataAssertionSubject() *SAMLPayLoadDataAssertionSubject`

NewSAMLPayLoadDataAssertionSubject instantiates a new SAMLPayLoadDataAssertionSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataAssertionSubjectWithDefaults

`func NewSAMLPayLoadDataAssertionSubjectWithDefaults() *SAMLPayLoadDataAssertionSubject`

NewSAMLPayLoadDataAssertionSubjectWithDefaults instantiates a new SAMLPayLoadDataAssertionSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameId

`func (o *SAMLPayLoadDataAssertionSubject) GetNameId() string`

GetNameId returns the NameId field if non-nil, zero value otherwise.

### GetNameIdOk

`func (o *SAMLPayLoadDataAssertionSubject) GetNameIdOk() (*string, bool)`

GetNameIdOk returns a tuple with the NameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameId

`func (o *SAMLPayLoadDataAssertionSubject) SetNameId(v string)`

SetNameId sets NameId field to given value.

### HasNameId

`func (o *SAMLPayLoadDataAssertionSubject) HasNameId() bool`

HasNameId returns a boolean if a field has been set.

### GetNameFormat

`func (o *SAMLPayLoadDataAssertionSubject) GetNameFormat() string`

GetNameFormat returns the NameFormat field if non-nil, zero value otherwise.

### GetNameFormatOk

`func (o *SAMLPayLoadDataAssertionSubject) GetNameFormatOk() (*string, bool)`

GetNameFormatOk returns a tuple with the NameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFormat

`func (o *SAMLPayLoadDataAssertionSubject) SetNameFormat(v string)`

SetNameFormat sets NameFormat field to given value.

### HasNameFormat

`func (o *SAMLPayLoadDataAssertionSubject) HasNameFormat() bool`

HasNameFormat returns a boolean if a field has been set.

### GetConfirmation

`func (o *SAMLPayLoadDataAssertionSubject) GetConfirmation() SAMLPayLoadDataAssertionSubjectConfirmation`

GetConfirmation returns the Confirmation field if non-nil, zero value otherwise.

### GetConfirmationOk

`func (o *SAMLPayLoadDataAssertionSubject) GetConfirmationOk() (*SAMLPayLoadDataAssertionSubjectConfirmation, bool)`

GetConfirmationOk returns a tuple with the Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmation

`func (o *SAMLPayLoadDataAssertionSubject) SetConfirmation(v SAMLPayLoadDataAssertionSubjectConfirmation)`

SetConfirmation sets Confirmation field to given value.

### HasConfirmation

`func (o *SAMLPayLoadDataAssertionSubject) HasConfirmation() bool`

HasConfirmation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


