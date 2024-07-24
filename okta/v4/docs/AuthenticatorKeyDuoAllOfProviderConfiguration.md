# AuthenticatorKeyDuoAllOfProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | The Duo Security API hostname | [optional] 
**IntegrationKey** | Pointer to **string** | The Duo Security integration key | [optional] 
**SecretKey** | Pointer to **string** | The Duo Security secret key | [optional] 
**UserNameTemplate** | Pointer to [**AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate**](AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate.md) |  | [optional] 

## Methods

### NewAuthenticatorKeyDuoAllOfProviderConfiguration

`func NewAuthenticatorKeyDuoAllOfProviderConfiguration() *AuthenticatorKeyDuoAllOfProviderConfiguration`

NewAuthenticatorKeyDuoAllOfProviderConfiguration instantiates a new AuthenticatorKeyDuoAllOfProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorKeyDuoAllOfProviderConfigurationWithDefaults

`func NewAuthenticatorKeyDuoAllOfProviderConfigurationWithDefaults() *AuthenticatorKeyDuoAllOfProviderConfiguration`

NewAuthenticatorKeyDuoAllOfProviderConfigurationWithDefaults instantiates a new AuthenticatorKeyDuoAllOfProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetIntegrationKey

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetIntegrationKey() string`

GetIntegrationKey returns the IntegrationKey field if non-nil, zero value otherwise.

### GetIntegrationKeyOk

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetIntegrationKeyOk() (*string, bool)`

GetIntegrationKeyOk returns a tuple with the IntegrationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationKey

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetIntegrationKey(v string)`

SetIntegrationKey sets IntegrationKey field to given value.

### HasIntegrationKey

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasIntegrationKey() bool`

HasIntegrationKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUserNameTemplate

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetUserNameTemplate() AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate`

GetUserNameTemplate returns the UserNameTemplate field if non-nil, zero value otherwise.

### GetUserNameTemplateOk

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) GetUserNameTemplateOk() (*AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate, bool)`

GetUserNameTemplateOk returns a tuple with the UserNameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameTemplate

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) SetUserNameTemplate(v AuthenticatorKeyDuoAllOfProviderConfigurationUserNameTemplate)`

SetUserNameTemplate sets UserNameTemplate field to given value.

### HasUserNameTemplate

`func (o *AuthenticatorKeyDuoAllOfProviderConfiguration) HasUserNameTemplate() bool`

HasUserNameTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


