# UserIdentificationPasskeySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowSignInButton** | Pointer to **string** | Controls whether to show the **Sign in with a passkey** button on the Sign-In Widget.  &gt; **Note:** You can only set &#x60;showSignInButton&#x60; to &#x60;ALWAYS&#x60; when the Passkey (FIDO2 WebAuthn) authenticator is active. | [optional] 

## Methods

### NewUserIdentificationPasskeySettings

`func NewUserIdentificationPasskeySettings() *UserIdentificationPasskeySettings`

NewUserIdentificationPasskeySettings instantiates a new UserIdentificationPasskeySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentificationPasskeySettingsWithDefaults

`func NewUserIdentificationPasskeySettingsWithDefaults() *UserIdentificationPasskeySettings`

NewUserIdentificationPasskeySettingsWithDefaults instantiates a new UserIdentificationPasskeySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowSignInButton

`func (o *UserIdentificationPasskeySettings) GetShowSignInButton() string`

GetShowSignInButton returns the ShowSignInButton field if non-nil, zero value otherwise.

### GetShowSignInButtonOk

`func (o *UserIdentificationPasskeySettings) GetShowSignInButtonOk() (*string, bool)`

GetShowSignInButtonOk returns a tuple with the ShowSignInButton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSignInButton

`func (o *UserIdentificationPasskeySettings) SetShowSignInButton(v string)`

SetShowSignInButton sets ShowSignInButton field to given value.

### HasShowSignInButton

`func (o *UserIdentificationPasskeySettings) HasShowSignInButton() bool`

HasShowSignInButton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


