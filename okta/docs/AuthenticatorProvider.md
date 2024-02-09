# AuthenticatorProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**AuthenticatorProviderConfiguration**](AuthenticatorProviderConfiguration.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticatorProvider

`func NewAuthenticatorProvider() *AuthenticatorProvider`

NewAuthenticatorProvider instantiates a new AuthenticatorProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorProviderWithDefaults

`func NewAuthenticatorProviderWithDefaults() *AuthenticatorProvider`

NewAuthenticatorProviderWithDefaults instantiates a new AuthenticatorProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *AuthenticatorProvider) GetConfiguration() AuthenticatorProviderConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AuthenticatorProvider) GetConfigurationOk() (*AuthenticatorProviderConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AuthenticatorProvider) SetConfiguration(v AuthenticatorProviderConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *AuthenticatorProvider) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetType

`func (o *AuthenticatorProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthenticatorProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthenticatorProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthenticatorProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


