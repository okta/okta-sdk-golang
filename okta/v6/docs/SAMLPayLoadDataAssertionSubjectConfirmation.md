# SAMLPayLoadDataAssertionSubjectConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Used to indicate how the authorization server confirmed the SAML assertion | [optional] 
**Data** | Pointer to [**SAMLPayLoadDataAssertionSubjectConfirmationData**](SAMLPayLoadDataAssertionSubjectConfirmationData.md) |  | [optional] 

## Methods

### NewSAMLPayLoadDataAssertionSubjectConfirmation

`func NewSAMLPayLoadDataAssertionSubjectConfirmation() *SAMLPayLoadDataAssertionSubjectConfirmation`

NewSAMLPayLoadDataAssertionSubjectConfirmation instantiates a new SAMLPayLoadDataAssertionSubjectConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayLoadDataAssertionSubjectConfirmationWithDefaults

`func NewSAMLPayLoadDataAssertionSubjectConfirmationWithDefaults() *SAMLPayLoadDataAssertionSubjectConfirmation`

NewSAMLPayLoadDataAssertionSubjectConfirmationWithDefaults instantiates a new SAMLPayLoadDataAssertionSubjectConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetData

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetData() SAMLPayLoadDataAssertionSubjectConfirmationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) GetDataOk() (*SAMLPayLoadDataAssertionSubjectConfirmationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) SetData(v SAMLPayLoadDataAssertionSubjectConfirmationData)`

SetData sets Data field to given value.

### HasData

`func (o *SAMLPayLoadDataAssertionSubjectConfirmation) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


