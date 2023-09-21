# AuthenticatorMethodWebAuthnAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserVerification** | Pointer to [**UserVerificationEnum**](UserVerificationEnum.md) |  | [optional] 
**Attachment** | Pointer to [**WebAuthnAttachment**](WebAuthnAttachment.md) |  | [optional] 

## Methods

### NewAuthenticatorMethodWebAuthnAllOfSettings

`func NewAuthenticatorMethodWebAuthnAllOfSettings() *AuthenticatorMethodWebAuthnAllOfSettings`

NewAuthenticatorMethodWebAuthnAllOfSettings instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults

`func NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults() *AuthenticatorMethodWebAuthnAllOfSettings`

NewAuthenticatorMethodWebAuthnAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodWebAuthnAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserVerification

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerification() UserVerificationEnum`

GetUserVerification returns the UserVerification field if non-nil, zero value otherwise.

### GetUserVerificationOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetUserVerificationOk() (*UserVerificationEnum, bool)`

GetUserVerificationOk returns a tuple with the UserVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerification

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetUserVerification(v UserVerificationEnum)`

SetUserVerification sets UserVerification field to given value.

### HasUserVerification

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasUserVerification() bool`

HasUserVerification returns a boolean if a field has been set.

### GetAttachment

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachment() WebAuthnAttachment`

GetAttachment returns the Attachment field if non-nil, zero value otherwise.

### GetAttachmentOk

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) GetAttachmentOk() (*WebAuthnAttachment, bool)`

GetAttachmentOk returns a tuple with the Attachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachment

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) SetAttachment(v WebAuthnAttachment)`

SetAttachment sets Attachment field to given value.

### HasAttachment

`func (o *AuthenticatorMethodWebAuthnAllOfSettings) HasAttachment() bool`

HasAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


