# OAuthApplicationCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signing** | Pointer to [**ApplicationCredentialsSigning**](ApplicationCredentialsSigning.md) |  | [optional] 
**UserNameTemplate** | Pointer to [**ApplicationCredentialsUsernameTemplate**](ApplicationCredentialsUsernameTemplate.md) |  | [optional] 
**OauthClient** | Pointer to [**ApplicationCredentialsOAuthClient**](ApplicationCredentialsOAuthClient.md) |  | [optional] 

## Methods

### NewOAuthApplicationCredentials

`func NewOAuthApplicationCredentials() *OAuthApplicationCredentials`

NewOAuthApplicationCredentials instantiates a new OAuthApplicationCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthApplicationCredentialsWithDefaults

`func NewOAuthApplicationCredentialsWithDefaults() *OAuthApplicationCredentials`

NewOAuthApplicationCredentialsWithDefaults instantiates a new OAuthApplicationCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigning

`func (o *OAuthApplicationCredentials) GetSigning() ApplicationCredentialsSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *OAuthApplicationCredentials) GetSigningOk() (*ApplicationCredentialsSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *OAuthApplicationCredentials) SetSigning(v ApplicationCredentialsSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *OAuthApplicationCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetUserNameTemplate

`func (o *OAuthApplicationCredentials) GetUserNameTemplate() ApplicationCredentialsUsernameTemplate`

GetUserNameTemplate returns the UserNameTemplate field if non-nil, zero value otherwise.

### GetUserNameTemplateOk

`func (o *OAuthApplicationCredentials) GetUserNameTemplateOk() (*ApplicationCredentialsUsernameTemplate, bool)`

GetUserNameTemplateOk returns a tuple with the UserNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameTemplate

`func (o *OAuthApplicationCredentials) SetUserNameTemplate(v ApplicationCredentialsUsernameTemplate)`

SetUserNameTemplate sets UserNameTemplate field to given value.

### HasUserNameTemplate

`func (o *OAuthApplicationCredentials) HasUserNameTemplate() bool`

HasUserNameTemplate returns a boolean if a field has been set.

### GetOauthClient

`func (o *OAuthApplicationCredentials) GetOauthClient() ApplicationCredentialsOAuthClient`

GetOauthClient returns the OauthClient field if non-nil, zero value otherwise.

### GetOauthClientOk

`func (o *OAuthApplicationCredentials) GetOauthClientOk() (*ApplicationCredentialsOAuthClient, bool)`

GetOauthClientOk returns a tuple with the OauthClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthClient

`func (o *OAuthApplicationCredentials) SetOauthClient(v ApplicationCredentialsOAuthClient)`

SetOauthClient sets OauthClient field to given value.

### HasOauthClient

`func (o *OAuthApplicationCredentials) HasOauthClient() bool`

HasOauthClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


