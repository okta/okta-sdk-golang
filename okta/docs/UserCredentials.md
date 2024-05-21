# UserCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**PasswordCredential**](PasswordCredential.md) |  | [optional] 
**Provider** | Pointer to [**AuthenticationProvider**](AuthenticationProvider.md) |  | [optional] 
**RecoveryQuestion** | Pointer to [**RecoveryQuestionCredential**](RecoveryQuestionCredential.md) |  | [optional] 

## Methods

### NewUserCredentials

`func NewUserCredentials() *UserCredentials`

NewUserCredentials instantiates a new UserCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCredentialsWithDefaults

`func NewUserCredentialsWithDefaults() *UserCredentials`

NewUserCredentialsWithDefaults instantiates a new UserCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UserCredentials) GetPassword() PasswordCredential`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCredentials) GetPasswordOk() (*PasswordCredential, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCredentials) SetPassword(v PasswordCredential)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProvider

`func (o *UserCredentials) GetProvider() AuthenticationProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UserCredentials) GetProviderOk() (*AuthenticationProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UserCredentials) SetProvider(v AuthenticationProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *UserCredentials) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetRecoveryQuestion

`func (o *UserCredentials) GetRecoveryQuestion() RecoveryQuestionCredential`

GetRecoveryQuestion returns the RecoveryQuestion field if non-nil, zero value otherwise.

### GetRecoveryQuestionOk

`func (o *UserCredentials) GetRecoveryQuestionOk() (*RecoveryQuestionCredential, bool)`

GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryQuestion

`func (o *UserCredentials) SetRecoveryQuestion(v RecoveryQuestionCredential)`

SetRecoveryQuestion sets RecoveryQuestion field to given value.

### HasRecoveryQuestion

`func (o *UserCredentials) HasRecoveryQuestion() bool`

HasRecoveryQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


