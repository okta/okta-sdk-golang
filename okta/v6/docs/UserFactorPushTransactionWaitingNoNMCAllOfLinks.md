# UserFactorPushTransactionWaitingNoNMCAllOfLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Poll** | Pointer to [**HrefObject**](HrefObject.md) | Polls the factor resource for status information. Always use the &#x60;poll&#x60; link instead of manually constructing your own URL. | [optional] 
**Cancel** | Pointer to [**HrefObject**](HrefObject.md) | Cancels a &#x60;push&#x60; factor challenge with a &#x60;WAITING&#x60; status | [optional] 

## Methods

### NewUserFactorPushTransactionWaitingNoNMCAllOfLinks

`func NewUserFactorPushTransactionWaitingNoNMCAllOfLinks() *UserFactorPushTransactionWaitingNoNMCAllOfLinks`

NewUserFactorPushTransactionWaitingNoNMCAllOfLinks instantiates a new UserFactorPushTransactionWaitingNoNMCAllOfLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorPushTransactionWaitingNoNMCAllOfLinksWithDefaults

`func NewUserFactorPushTransactionWaitingNoNMCAllOfLinksWithDefaults() *UserFactorPushTransactionWaitingNoNMCAllOfLinks`

NewUserFactorPushTransactionWaitingNoNMCAllOfLinksWithDefaults instantiates a new UserFactorPushTransactionWaitingNoNMCAllOfLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoll

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetPoll() HrefObject`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetPollOk() (*HrefObject, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) SetPoll(v HrefObject)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetCancel

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetCancel() HrefObject`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) GetCancelOk() (*HrefObject, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) SetCancel(v HrefObject)`

SetCancel sets Cancel field to given value.

### HasCancel

`func (o *UserFactorPushTransactionWaitingNoNMCAllOfLinks) HasCancel() bool`

HasCancel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


