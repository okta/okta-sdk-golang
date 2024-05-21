# AuthenticatorProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPort** | Pointer to **int32** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**SharedSecret** | Pointer to **string** |  | [optional] 
**UserNameTemplate** | Pointer to [**AuthenticatorProviderConfigurationUserNameTemplate**](AuthenticatorProviderConfigurationUserNameTemplate.md) |  | [optional] 

## Methods

### NewAuthenticatorProviderConfiguration

`func NewAuthenticatorProviderConfiguration() *AuthenticatorProviderConfiguration`

NewAuthenticatorProviderConfiguration instantiates a new AuthenticatorProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorProviderConfigurationWithDefaults

`func NewAuthenticatorProviderConfigurationWithDefaults() *AuthenticatorProviderConfiguration`

NewAuthenticatorProviderConfigurationWithDefaults instantiates a new AuthenticatorProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthPort

`func (o *AuthenticatorProviderConfiguration) GetAuthPort() int32`

GetAuthPort returns the AuthPort field if non-nil, zero value otherwise.

### GetAuthPortOk

`func (o *AuthenticatorProviderConfiguration) GetAuthPortOk() (*int32, bool)`

GetAuthPortOk returns a tuple with the AuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPort

`func (o *AuthenticatorProviderConfiguration) SetAuthPort(v int32)`

SetAuthPort sets AuthPort field to given value.

### HasAuthPort

`func (o *AuthenticatorProviderConfiguration) HasAuthPort() bool`

HasAuthPort returns a boolean if a field has been set.

### GetHostName

`func (o *AuthenticatorProviderConfiguration) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AuthenticatorProviderConfiguration) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AuthenticatorProviderConfiguration) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AuthenticatorProviderConfiguration) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInstanceId

`func (o *AuthenticatorProviderConfiguration) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AuthenticatorProviderConfiguration) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AuthenticatorProviderConfiguration) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *AuthenticatorProviderConfiguration) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetSharedSecret

`func (o *AuthenticatorProviderConfiguration) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *AuthenticatorProviderConfiguration) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *AuthenticatorProviderConfiguration) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *AuthenticatorProviderConfiguration) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### GetUserNameTemplate

`func (o *AuthenticatorProviderConfiguration) GetUserNameTemplate() AuthenticatorProviderConfigurationUserNameTemplate`

GetUserNameTemplate returns the UserNameTemplate field if non-nil, zero value otherwise.

### GetUserNameTemplateOk

`func (o *AuthenticatorProviderConfiguration) GetUserNameTemplateOk() (*AuthenticatorProviderConfigurationUserNameTemplate, bool)`

GetUserNameTemplateOk returns a tuple with the UserNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameTemplate

`func (o *AuthenticatorProviderConfiguration) SetUserNameTemplate(v AuthenticatorProviderConfigurationUserNameTemplate)`

SetUserNameTemplate sets UserNameTemplate field to given value.

### HasUserNameTemplate

`func (o *AuthenticatorProviderConfiguration) HasUserNameTemplate() bool`

HasUserNameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


