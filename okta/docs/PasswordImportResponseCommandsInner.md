# PasswordImportResponseCommandsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The location where you specify the command. For the password import inline hook, there&#39;s only one command, &#x60;com.okta.action.update&#x60;. | [optional] 
**Value** | Pointer to [**PasswordImportResponseCommandsInnerValue**](PasswordImportResponseCommandsInnerValue.md) |  | [optional] 

## Methods

### NewPasswordImportResponseCommandsInner

`func NewPasswordImportResponseCommandsInner() *PasswordImportResponseCommandsInner`

NewPasswordImportResponseCommandsInner instantiates a new PasswordImportResponseCommandsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordImportResponseCommandsInnerWithDefaults

`func NewPasswordImportResponseCommandsInnerWithDefaults() *PasswordImportResponseCommandsInner`

NewPasswordImportResponseCommandsInnerWithDefaults instantiates a new PasswordImportResponseCommandsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PasswordImportResponseCommandsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PasswordImportResponseCommandsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PasswordImportResponseCommandsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PasswordImportResponseCommandsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *PasswordImportResponseCommandsInner) GetValue() PasswordImportResponseCommandsInnerValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PasswordImportResponseCommandsInner) GetValueOk() (*PasswordImportResponseCommandsInnerValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PasswordImportResponseCommandsInner) SetValue(v PasswordImportResponseCommandsInnerValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *PasswordImportResponseCommandsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


