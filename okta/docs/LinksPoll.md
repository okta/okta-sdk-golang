# LinksPoll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Poll** | Pointer to [**HrefObject**](HrefObject.md) | Polls the factor resource for status information. Always use the &#x60;poll&#x60; link instead of manually constructing your own URL. | [optional] 

## Methods

### NewLinksPoll

`func NewLinksPoll() *LinksPoll`

NewLinksPoll instantiates a new LinksPoll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksPollWithDefaults

`func NewLinksPollWithDefaults() *LinksPoll`

NewLinksPollWithDefaults instantiates a new LinksPoll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoll

`func (o *LinksPoll) GetPoll() HrefObject`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *LinksPoll) GetPollOk() (*HrefObject, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *LinksPoll) SetPoll(v HrefObject)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *LinksPoll) HasPoll() bool`

HasPoll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


