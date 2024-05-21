# UserFactorWebAuthnProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticatorName** | Pointer to **string** | Human-readable name of the authenticator | [optional] 
**CredentialId** | Pointer to **string** | ID for the Factor credential | [optional] 

## Methods

### NewUserFactorWebAuthnProfile

`func NewUserFactorWebAuthnProfile() *UserFactorWebAuthnProfile`

NewUserFactorWebAuthnProfile instantiates a new UserFactorWebAuthnProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFactorWebAuthnProfileWithDefaults

`func NewUserFactorWebAuthnProfileWithDefaults() *UserFactorWebAuthnProfile`

NewUserFactorWebAuthnProfileWithDefaults instantiates a new UserFactorWebAuthnProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticatorName

`func (o *UserFactorWebAuthnProfile) GetAuthenticatorName() string`

GetAuthenticatorName returns the AuthenticatorName field if non-nil, zero value otherwise.

### GetAuthenticatorNameOk

`func (o *UserFactorWebAuthnProfile) GetAuthenticatorNameOk() (*string, bool)`

GetAuthenticatorNameOk returns a tuple with the AuthenticatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticatorName

`func (o *UserFactorWebAuthnProfile) SetAuthenticatorName(v string)`

SetAuthenticatorName sets AuthenticatorName field to given value.

### HasAuthenticatorName

`func (o *UserFactorWebAuthnProfile) HasAuthenticatorName() bool`

HasAuthenticatorName returns a boolean if a field has been set.

### GetCredentialId

`func (o *UserFactorWebAuthnProfile) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *UserFactorWebAuthnProfile) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *UserFactorWebAuthnProfile) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *UserFactorWebAuthnProfile) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


