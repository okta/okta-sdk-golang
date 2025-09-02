# OrgCreationAdminCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**OrgCreationAdminCredentialsPassword**](OrgCreationAdminCredentialsPassword.md) |  | [optional] 
**RecoveryQuestion** | Pointer to [**RecoveryQuestionCredential**](RecoveryQuestionCredential.md) |  | [optional] 

## Methods

### NewOrgCreationAdminCredentials

`func NewOrgCreationAdminCredentials() *OrgCreationAdminCredentials`

NewOrgCreationAdminCredentials instantiates a new OrgCreationAdminCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgCreationAdminCredentialsWithDefaults

`func NewOrgCreationAdminCredentialsWithDefaults() *OrgCreationAdminCredentials`

NewOrgCreationAdminCredentialsWithDefaults instantiates a new OrgCreationAdminCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *OrgCreationAdminCredentials) GetPassword() OrgCreationAdminCredentialsPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OrgCreationAdminCredentials) GetPasswordOk() (*OrgCreationAdminCredentialsPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OrgCreationAdminCredentials) SetPassword(v OrgCreationAdminCredentialsPassword)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OrgCreationAdminCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRecoveryQuestion

`func (o *OrgCreationAdminCredentials) GetRecoveryQuestion() RecoveryQuestionCredential`

GetRecoveryQuestion returns the RecoveryQuestion field if non-nil, zero value otherwise.

### GetRecoveryQuestionOk

`func (o *OrgCreationAdminCredentials) GetRecoveryQuestionOk() (*RecoveryQuestionCredential, bool)`

GetRecoveryQuestionOk returns a tuple with the RecoveryQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryQuestion

`func (o *OrgCreationAdminCredentials) SetRecoveryQuestion(v RecoveryQuestionCredential)`

SetRecoveryQuestion sets RecoveryQuestion field to given value.

### HasRecoveryQuestion

`func (o *OrgCreationAdminCredentials) HasRecoveryQuestion() bool`

HasRecoveryQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


