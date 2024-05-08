# UserFactorSecurityQuestionProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answer** | Pointer to **string** | Answer to the question | [optional] 
**Question** | Pointer to **string** | Unique key for the question | [optional] 
**QuestionText** | Pointer to **string** | Human-readable text displayed to the user | [optional] [readonly] 

## Methods

### NewUserFactorSecurityQuestionProfile

`func NewUserFactorSecurityQuestionProfile() *UserFactorSecurityQuestionProfile`

NewUserFactorSecurityQuestionProfile instantiates a new UserFactorSecurityQuestionProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorSecurityQuestionProfileWithDefaults

`func NewUserFactorSecurityQuestionProfileWithDefaults() *UserFactorSecurityQuestionProfile`

NewUserFactorSecurityQuestionProfileWithDefaults instantiates a new UserFactorSecurityQuestionProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswer

`func (o *UserFactorSecurityQuestionProfile) GetAnswer() string`

GetAnswer returns the Answer field if non-nil, zero value otherwise.

### GetAnswerOk

`func (o *UserFactorSecurityQuestionProfile) GetAnswerOk() (*string, bool)`

GetAnswerOk returns a tuple with the Answer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswer

`func (o *UserFactorSecurityQuestionProfile) SetAnswer(v string)`

SetAnswer sets Answer field to given value.

### HasAnswer

`func (o *UserFactorSecurityQuestionProfile) HasAnswer() bool`

HasAnswer returns a boolean if a field has been set.

### GetQuestion

`func (o *UserFactorSecurityQuestionProfile) GetQuestion() string`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *UserFactorSecurityQuestionProfile) GetQuestionOk() (*string, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *UserFactorSecurityQuestionProfile) SetQuestion(v string)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *UserFactorSecurityQuestionProfile) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetQuestionText

`func (o *UserFactorSecurityQuestionProfile) GetQuestionText() string`

GetQuestionText returns the QuestionText field if non-nil, zero value otherwise.

### GetQuestionTextOk

`func (o *UserFactorSecurityQuestionProfile) GetQuestionTextOk() (*string, bool)`

GetQuestionTextOk returns a tuple with the QuestionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionText

`func (o *UserFactorSecurityQuestionProfile) SetQuestionText(v string)`

SetQuestionText sets QuestionText field to given value.

### HasQuestionText

`func (o *UserFactorSecurityQuestionProfile) HasQuestionText() bool`

HasQuestionText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


