# ApplicationCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signing** | Pointer to [**ApplicationCredentialsSigning**](ApplicationCredentialsSigning.md) |  | [optional] 
**UserNameTemplate** | Pointer to [**ApplicationCredentialsUsernameTemplate**](ApplicationCredentialsUsernameTemplate.md) |  | [optional] 

## Methods

### NewApplicationCredentials

`func NewApplicationCredentials() *ApplicationCredentials`

NewApplicationCredentials instantiates a new ApplicationCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCredentialsWithDefaults

`func NewApplicationCredentialsWithDefaults() *ApplicationCredentials`

NewApplicationCredentialsWithDefaults instantiates a new ApplicationCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigning

`func (o *ApplicationCredentials) GetSigning() ApplicationCredentialsSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *ApplicationCredentials) GetSigningOk() (*ApplicationCredentialsSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *ApplicationCredentials) SetSigning(v ApplicationCredentialsSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *ApplicationCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetUserNameTemplate

`func (o *ApplicationCredentials) GetUserNameTemplate() ApplicationCredentialsUsernameTemplate`

GetUserNameTemplate returns the UserNameTemplate field if non-nil, zero value otherwise.

### GetUserNameTemplateOk

`func (o *ApplicationCredentials) GetUserNameTemplateOk() (*ApplicationCredentialsUsernameTemplate, bool)`

GetUserNameTemplateOk returns a tuple with the UserNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameTemplate

`func (o *ApplicationCredentials) SetUserNameTemplate(v ApplicationCredentialsUsernameTemplate)`

SetUserNameTemplate sets UserNameTemplate field to given value.

### HasUserNameTemplate

`func (o *ApplicationCredentials) HasUserNameTemplate() bool`

HasUserNameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


