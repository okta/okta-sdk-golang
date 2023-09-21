# AuthenticatorMethodSignedNonceAllOfSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithms** | Pointer to [**[]AuthenticatorMethodAlgorithm**](AuthenticatorMethodAlgorithm.md) |  | [optional] 
**KeyProtection** | Pointer to [**PushMethodKeyProtection**](PushMethodKeyProtection.md) |  | [optional] 
**ShowSignInWithOV** | Pointer to [**ShowSignInWithOV**](ShowSignInWithOV.md) |  | [optional] 

## Methods

### NewAuthenticatorMethodSignedNonceAllOfSettings

`func NewAuthenticatorMethodSignedNonceAllOfSettings() *AuthenticatorMethodSignedNonceAllOfSettings`

NewAuthenticatorMethodSignedNonceAllOfSettings instantiates a new AuthenticatorMethodSignedNonceAllOfSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodSignedNonceAllOfSettingsWithDefaults

`func NewAuthenticatorMethodSignedNonceAllOfSettingsWithDefaults() *AuthenticatorMethodSignedNonceAllOfSettings`

NewAuthenticatorMethodSignedNonceAllOfSettingsWithDefaults instantiates a new AuthenticatorMethodSignedNonceAllOfSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithms

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetAlgorithms() []AuthenticatorMethodAlgorithm`

GetAlgorithms returns the Algorithms field if non-nil, zero value otherwise.

### GetAlgorithmsOk

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetAlgorithmsOk() (*[]AuthenticatorMethodAlgorithm, bool)`

GetAlgorithmsOk returns a tuple with the Algorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithms

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) SetAlgorithms(v []AuthenticatorMethodAlgorithm)`

SetAlgorithms sets Algorithms field to given value.

### HasAlgorithms

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) HasAlgorithms() bool`

HasAlgorithms returns a boolean if a field has been set.

### GetKeyProtection

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetKeyProtection() PushMethodKeyProtection`

GetKeyProtection returns the KeyProtection field if non-nil, zero value otherwise.

### GetKeyProtectionOk

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetKeyProtectionOk() (*PushMethodKeyProtection, bool)`

GetKeyProtectionOk returns a tuple with the KeyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProtection

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) SetKeyProtection(v PushMethodKeyProtection)`

SetKeyProtection sets KeyProtection field to given value.

### HasKeyProtection

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) HasKeyProtection() bool`

HasKeyProtection returns a boolean if a field has been set.

### GetShowSignInWithOV

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetShowSignInWithOV() ShowSignInWithOV`

GetShowSignInWithOV returns the ShowSignInWithOV field if non-nil, zero value otherwise.

### GetShowSignInWithOVOk

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) GetShowSignInWithOVOk() (*ShowSignInWithOV, bool)`

GetShowSignInWithOVOk returns a tuple with the ShowSignInWithOV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSignInWithOV

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) SetShowSignInWithOV(v ShowSignInWithOV)`

SetShowSignInWithOV sets ShowSignInWithOV field to given value.

### HasShowSignInWithOV

`func (o *AuthenticatorMethodSignedNonceAllOfSettings) HasShowSignInWithOV() bool`

HasShowSignInWithOV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


