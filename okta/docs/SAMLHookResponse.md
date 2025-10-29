# SAMLHookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]SAMLHookResponseCommandsInner**](SAMLHookResponseCommandsInner.md) | The &#x60;commands&#x60; object is where you tell Okta to add additional claims to the assertion or to modify the existing assertion statements.  &#x60;commands&#x60; is an array, allowing you to send multiple commands. In each array element, include a &#x60;type&#x60; property and a &#x60;value&#x60; property. The &#x60;type&#x60; property is where you specify which of the supported commands you want to execute, and &#x60;value&#x60; is where you supply an operand for that command. In the case of the SAML assertion inline hook, the &#x60;value&#x60; property is itself a nested object, in which you specify a particular operation, a path to act on, and a value. | [optional] 
**Error** | Pointer to [**SAMLHookResponseError**](SAMLHookResponseError.md) |  | [optional] 

## Methods

### NewSAMLHookResponse

`func NewSAMLHookResponse() *SAMLHookResponse`

NewSAMLHookResponse instantiates a new SAMLHookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLHookResponseWithDefaults

`func NewSAMLHookResponseWithDefaults() *SAMLHookResponse`

NewSAMLHookResponseWithDefaults instantiates a new SAMLHookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *SAMLHookResponse) GetCommands() []SAMLHookResponseCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *SAMLHookResponse) GetCommandsOk() (*[]SAMLHookResponseCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *SAMLHookResponse) SetCommands(v []SAMLHookResponseCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *SAMLHookResponse) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetError

`func (o *SAMLHookResponse) GetError() SAMLHookResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SAMLHookResponse) GetErrorOk() (*SAMLHookResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SAMLHookResponse) SetError(v SAMLHookResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *SAMLHookResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


