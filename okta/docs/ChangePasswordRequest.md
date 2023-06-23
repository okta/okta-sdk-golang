# ChangePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | Pointer to [**PasswordCredential**](PasswordCredential.md) |  | [optional] 
**OldPassword** | Pointer to [**PasswordCredential**](PasswordCredential.md) |  | [optional] 

## Methods

### NewChangePasswordRequest

`func NewChangePasswordRequest() *ChangePasswordRequest`

NewChangePasswordRequest instantiates a new ChangePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangePasswordRequestWithDefaults

`func NewChangePasswordRequestWithDefaults() *ChangePasswordRequest`

NewChangePasswordRequestWithDefaults instantiates a new ChangePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *ChangePasswordRequest) GetNewPassword() PasswordCredential`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ChangePasswordRequest) GetNewPasswordOk() (*PasswordCredential, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ChangePasswordRequest) SetNewPassword(v PasswordCredential)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *ChangePasswordRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetOldPassword

`func (o *ChangePasswordRequest) GetOldPassword() PasswordCredential`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *ChangePasswordRequest) GetOldPasswordOk() (*PasswordCredential, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *ChangePasswordRequest) SetOldPassword(v PasswordCredential)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *ChangePasswordRequest) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


