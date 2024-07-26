# PasswordImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | Pointer to [**[]PasswordImportResponseCommandsInner**](PasswordImportResponseCommandsInner.md) | The &#x60;commands&#x60; object specifies whether Okta accepts the end user&#39;s sign-in credentials as valid or not. For the password import inline hook, you typically only return one &#x60;commands&#x60; object with one array element in it. | [optional] 

## Methods

### NewPasswordImportResponse

`func NewPasswordImportResponse() *PasswordImportResponse`

NewPasswordImportResponse instantiates a new PasswordImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordImportResponseWithDefaults

`func NewPasswordImportResponseWithDefaults() *PasswordImportResponse`

NewPasswordImportResponseWithDefaults instantiates a new PasswordImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *PasswordImportResponse) GetCommands() []PasswordImportResponseCommandsInner`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *PasswordImportResponse) GetCommandsOk() (*[]PasswordImportResponseCommandsInner, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *PasswordImportResponse) SetCommands(v []PasswordImportResponseCommandsInner)`

SetCommands sets Commands field to given value.

### HasCommands

`func (o *PasswordImportResponse) HasCommands() bool`

HasCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


