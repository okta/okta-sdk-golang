# UserCredentialsWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**PasswordCredential**](PasswordCredential.md) |  | [optional] 
**Provider** | Pointer to [**AuthenticationProviderWritable**](AuthenticationProviderWritable.md) |  | [optional] 
**RecoveryQuestion** | Pointer to [**RecoveryQuestionCredential**](RecoveryQuestionCredential.md) |  | [optional] 

## Methods

### NewUserCredentialsWritable

`func NewUserCredentialsWritable() *UserCredentialsWritable`

NewUserCredentialsWritable instantiates a new UserCredentialsWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCredentialsWritableWithDefaults

`func NewUserCredentialsWritableWithDefaults() *UserCredentialsWritable`

NewUserCredentialsWritableWithDefaults instantiates a new UserCredentialsWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UserCredentialsWritable) GetPassword() PasswordCredential`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCredentialsWritable) GetPasswordOk() (*PasswordCredential, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCredentialsWritable) SetPassword(v PasswordCredential)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCredentialsWritable) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProvider

`func (o *UserCredentialsWritable) GetProvider() AuthenticationProviderWritable`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserCredentialsWritable) GetProviderOk() (*AuthenticationProviderWritable, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserCredentialsWritable) SetProvider(v AuthenticationProviderWritable)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserCredentialsWritable) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRecoveryQuestion

`func (o *UserCredentialsWritable) GetRecoveryQuestion() RecoveryQuestionCredential`

GetRecoveryQuestion returns the RecoveryQuestion field if non-nil, zero value otherwise.

### GetRecoveryQuestionOk

`func (o *UserCredentialsWritable) GetRecoveryQuestionOk() (*RecoveryQuestionCredential, bool)`

GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryQuestion

`func (o *UserCredentialsWritable) SetRecoveryQuestion(v RecoveryQuestionCredential)`

SetRecoveryQuestion sets RecoveryQuestion field to given value.

### HasRecoveryQuestion

`func (o *UserCredentialsWritable) HasRecoveryQuestion() bool`

HasRecoveryQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


