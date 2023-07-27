# SchemeApplicationCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signing** | Pointer to [**ApplicationCredentialsSigning**](ApplicationCredentialsSigning.md) |  | [optional] 
**UserNameTemplate** | Pointer to [**ApplicationCredentialsUsernameTemplate**](ApplicationCredentialsUsernameTemplate.md) |  | [optional] 
**Password** | Pointer to [**PasswordCredential**](PasswordCredential.md) |  | [optional] 
**RevealPassword** | Pointer to **bool** |  | [optional] 
**Scheme** | Pointer to [**ApplicationCredentialsScheme**](ApplicationCredentialsScheme.md) |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewSchemeApplicationCredentials

`func NewSchemeApplicationCredentials() *SchemeApplicationCredentials`

NewSchemeApplicationCredentials instantiates a new SchemeApplicationCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemeApplicationCredentialsWithDefaults

`func NewSchemeApplicationCredentialsWithDefaults() *SchemeApplicationCredentials`

NewSchemeApplicationCredentialsWithDefaults instantiates a new SchemeApplicationCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigning

`func (o *SchemeApplicationCredentials) GetSigning() ApplicationCredentialsSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *SchemeApplicationCredentials) GetSigningOk() (*ApplicationCredentialsSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *SchemeApplicationCredentials) SetSigning(v ApplicationCredentialsSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *SchemeApplicationCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetUserNameTemplate

`func (o *SchemeApplicationCredentials) GetUserNameTemplate() ApplicationCredentialsUsernameTemplate`

GetUserNameTemplate returns the UserNameTemplate field if non-nil, zero value otherwise.

### GetUserNameTemplateOk

`func (o *SchemeApplicationCredentials) GetUserNameTemplateOk() (*ApplicationCredentialsUsernameTemplate, bool)`

GetUserNameTemplateOk returns a tuple with the UserNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameTemplate

`func (o *SchemeApplicationCredentials) SetUserNameTemplate(v ApplicationCredentialsUsernameTemplate)`

SetUserNameTemplate sets UserNameTemplate field to given value.

### HasUserNameTemplate

`func (o *SchemeApplicationCredentials) HasUserNameTemplate() bool`

HasUserNameTemplate returns a boolean if a field has been set.

### GetPassword

`func (o *SchemeApplicationCredentials) GetPassword() PasswordCredential`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SchemeApplicationCredentials) GetPasswordOk() (*PasswordCredential, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SchemeApplicationCredentials) SetPassword(v PasswordCredential)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SchemeApplicationCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRevealPassword

`func (o *SchemeApplicationCredentials) GetRevealPassword() bool`

GetRevealPassword returns the RevealPassword field if non-nil, zero value otherwise.

### GetRevealPasswordOk

`func (o *SchemeApplicationCredentials) GetRevealPasswordOk() (*bool, bool)`

GetRevealPasswordOk returns a tuple with the RevealPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevealPassword

`func (o *SchemeApplicationCredentials) SetRevealPassword(v bool)`

SetRevealPassword sets RevealPassword field to given value.

### HasRevealPassword

`func (o *SchemeApplicationCredentials) HasRevealPassword() bool`

HasRevealPassword returns a boolean if a field has been set.

### GetScheme

`func (o *SchemeApplicationCredentials) GetScheme() ApplicationCredentialsScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *SchemeApplicationCredentials) GetSchemeOk() (*ApplicationCredentialsScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *SchemeApplicationCredentials) SetScheme(v ApplicationCredentialsScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *SchemeApplicationCredentials) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetUserName

`func (o *SchemeApplicationCredentials) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SchemeApplicationCredentials) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SchemeApplicationCredentials) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SchemeApplicationCredentials) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


