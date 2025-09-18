# ExecuteInlineHook200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]UserImportResponseCommandsInner**](UserImportResponseCommandsInner.md) | The &#x60;commands&#x60; object is where you can provide commands to Okta. It is an array that allows you to send multiple commands. Each array element needs to consist of a type-value pair. | [optional] 
**Error** | Pointer to [**UserImportResponseError**](UserImportResponseError.md) |  | [optional] 

## Methods

### NewExecuteInlineHook200Response

`func NewExecuteInlineHook200Response() *ExecuteInlineHook200Response`

NewExecuteInlineHook200Response instantiates a new ExecuteInlineHook200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteInlineHook200ResponseWithDefaults

`func NewExecuteInlineHook200ResponseWithDefaults() *ExecuteInlineHook200Response`

NewExecuteInlineHook200ResponseWithDefaults instantiates a new ExecuteInlineHook200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *ExecuteInlineHook200Response) GetCommands() []UserImportResponseCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ExecuteInlineHook200Response) GetCommandsOk() (*[]UserImportResponseCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ExecuteInlineHook200Response) SetCommands(v []UserImportResponseCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *ExecuteInlineHook200Response) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetError

`func (o *ExecuteInlineHook200Response) GetError() UserImportResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExecuteInlineHook200Response) GetErrorOk() (*UserImportResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExecuteInlineHook200Response) SetError(v UserImportResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ExecuteInlineHook200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


