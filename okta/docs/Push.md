# Push

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseNumberMatchingChallenge** | Pointer to **bool** | Select whether to use a number matching challenge for a &#x60;push&#x60; factor.  &gt; **Note:** Sending a request with a body is required when you verify a &#x60;push&#x60; factor with a number matching challenge. | [optional] 

## Methods

### NewPush

`func NewPush() *Push`

NewPush instantiates a new Push object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPushWithDefaults

`func NewPushWithDefaults() *Push`

NewPushWithDefaults instantiates a new Push object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseNumberMatchingChallenge

`func (o *Push) GetUseNumberMatchingChallenge() bool`

GetUseNumberMatchingChallenge returns the UseNumberMatchingChallenge field if non-nil, zero value otherwise.

### GetUseNumberMatchingChallengeOk

`func (o *Push) GetUseNumberMatchingChallengeOk() (*bool, bool)`

GetUseNumberMatchingChallengeOk returns a tuple with the UseNumberMatchingChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumberMatchingChallenge

`func (o *Push) SetUseNumberMatchingChallenge(v bool)`

SetUseNumberMatchingChallenge sets UseNumberMatchingChallenge field to given value.

### HasUseNumberMatchingChallenge

`func (o *Push) HasUseNumberMatchingChallenge() bool`

HasUseNumberMatchingChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


