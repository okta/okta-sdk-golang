# UserIdentifierConditionEvaluatorPattern

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchType** | **string** | The type of pattern. For regex, use &#x60;EXPRESSION&#x60;. | 
**Value** | **string** | The regular expression or simple match string | 

## Methods

### NewUserIdentifierConditionEvaluatorPattern

`func NewUserIdentifierConditionEvaluatorPattern(matchType string, value string, ) *UserIdentifierConditionEvaluatorPattern`

NewUserIdentifierConditionEvaluatorPattern instantiates a new UserIdentifierConditionEvaluatorPattern object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentifierConditionEvaluatorPatternWithDefaults

`func NewUserIdentifierConditionEvaluatorPatternWithDefaults() *UserIdentifierConditionEvaluatorPattern`

NewUserIdentifierConditionEvaluatorPatternWithDefaults instantiates a new UserIdentifierConditionEvaluatorPattern object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchType

`func (o *UserIdentifierConditionEvaluatorPattern) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *UserIdentifierConditionEvaluatorPattern) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *UserIdentifierConditionEvaluatorPattern) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetValue

`func (o *UserIdentifierConditionEvaluatorPattern) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserIdentifierConditionEvaluatorPattern) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserIdentifierConditionEvaluatorPattern) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


