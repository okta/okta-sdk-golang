# UserImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]UserImportResponseCommandsInner**](UserImportResponseCommandsInner.md) | The &#x60;commands&#x60; object is where you can provide commands to Okta. It is an array that allows you to send multiple commands. Each array element needs to consist of a type-value pair. | [optional] 
**Error** | Pointer to [**UserImportResponseError**](UserImportResponseError.md) |  | [optional] 

## Methods

### NewUserImportResponse

`func NewUserImportResponse() *UserImportResponse`

NewUserImportResponse instantiates a new UserImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportResponseWithDefaults

`func NewUserImportResponseWithDefaults() *UserImportResponse`

NewUserImportResponseWithDefaults instantiates a new UserImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *UserImportResponse) GetCommands() []UserImportResponseCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *UserImportResponse) GetCommandsOk() (*[]UserImportResponseCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *UserImportResponse) SetCommands(v []UserImportResponseCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *UserImportResponse) HasCommands() bool`

HasCommands returns a boolean if a field has been set.

### GetError

`func (o *UserImportResponse) GetError() UserImportResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserImportResponse) GetErrorOk() (*UserImportResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserImportResponse) SetError(v UserImportResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *UserImportResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


