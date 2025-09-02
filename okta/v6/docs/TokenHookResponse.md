# TokenHookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]TokenHookResponseCommandsInner**](TokenHookResponseCommandsInner.md) | You can use the &#x60;commands&#x60; object to provide commands to Okta. It&#39;s where you can tell Okta to add more claims to the token. The &#x60;commands&#x60; object is an array, allowing you to send multiple commands. In each array element, there needs to be a &#x60;type&#x60; property and &#x60;value&#x60; property. The &#x60;type&#x60; property is where you specify which of the supported commands you want to execute, and &#x60;value&#x60; is where you supply an operand for that command. In the case of the token hook type, the &#x60;value&#x60; property is itself a nested object in which you specify a particular operation, a path to act on, and a value. | [optional] 
**Error** | Pointer to [**TokenHookResponseError**](TokenHookResponseError.md) |  | [optional] 

## Methods

### NewTokenHookResponse

`func NewTokenHookResponse() *TokenHookResponse`

NewTokenHookResponse instantiates a new TokenHookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenHookResponseWithDefaults

`func NewTokenHookResponseWithDefaults() *TokenHookResponse`

NewTokenHookResponseWithDefaults instantiates a new TokenHookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *TokenHookResponse) GetCommands() []TokenHookResponseCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *TokenHookResponse) GetCommandsOk() (*[]TokenHookResponseCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *TokenHookResponse) SetCommands(v []TokenHookResponseCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *TokenHookResponse) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetError

`func (o *TokenHookResponse) GetError() TokenHookResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TokenHookResponse) GetErrorOk() (*TokenHookResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TokenHookResponse) SetError(v TokenHookResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *TokenHookResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


