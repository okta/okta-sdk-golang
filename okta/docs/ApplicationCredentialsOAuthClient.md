# ApplicationCredentialsOAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoKeyRotation** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**TokenEndpointAuthMethod** | Pointer to [**OAuthEndpointAuthenticationMethod**](OAuthEndpointAuthenticationMethod.md) |  | [optional] 

## Methods

### NewApplicationCredentialsOAuthClient

`func NewApplicationCredentialsOAuthClient() *ApplicationCredentialsOAuthClient`

NewApplicationCredentialsOAuthClient instantiates a new ApplicationCredentialsOAuthClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCredentialsOAuthClientWithDefaults

`func NewApplicationCredentialsOAuthClientWithDefaults() *ApplicationCredentialsOAuthClient`

NewApplicationCredentialsOAuthClientWithDefaults instantiates a new ApplicationCredentialsOAuthClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoKeyRotation

`func (o *ApplicationCredentialsOAuthClient) GetAutoKeyRotation() bool`

GetAutoKeyRotation returns the AutoKeyRotation field if non-nil, zero value otherwise.

### GetAutoKeyRotationOk

`func (o *ApplicationCredentialsOAuthClient) GetAutoKeyRotationOk() (*bool, bool)`

GetAutoKeyRotationOk returns a tuple with the AutoKeyRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoKeyRotation

`func (o *ApplicationCredentialsOAuthClient) SetAutoKeyRotation(v bool)`

SetAutoKeyRotation sets AutoKeyRotation field to given value.

### HasAutoKeyRotation

`func (o *ApplicationCredentialsOAuthClient) HasAutoKeyRotation() bool`

HasAutoKeyRotation returns a boolean if a field has been set.

### GetClientId

`func (o *ApplicationCredentialsOAuthClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ApplicationCredentialsOAuthClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ApplicationCredentialsOAuthClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ApplicationCredentialsOAuthClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *ApplicationCredentialsOAuthClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApplicationCredentialsOAuthClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApplicationCredentialsOAuthClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ApplicationCredentialsOAuthClient) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethod() OAuthEndpointAuthenticationMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethodOk() (*OAuthEndpointAuthenticationMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationCredentialsOAuthClient) SetTokenEndpointAuthMethod(v OAuthEndpointAuthenticationMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *ApplicationCredentialsOAuthClient) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


