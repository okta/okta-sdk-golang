# PasswordImportRequestDataAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to **string** | The status of the user credential, either &#x60;UNVERIFIED&#x60; or &#x60;VERIFIED&#x60; | [optional] [default to "UNVERIFIED"]

## Methods

### NewPasswordImportRequestDataAction

`func NewPasswordImportRequestDataAction() *PasswordImportRequestDataAction`

NewPasswordImportRequestDataAction instantiates a new PasswordImportRequestDataAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordImportRequestDataActionWithDefaults

`func NewPasswordImportRequestDataActionWithDefaults() *PasswordImportRequestDataAction`

NewPasswordImportRequestDataActionWithDefaults instantiates a new PasswordImportRequestDataAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *PasswordImportRequestDataAction) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *PasswordImportRequestDataAction) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *PasswordImportRequestDataAction) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *PasswordImportRequestDataAction) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


