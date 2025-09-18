# OAuthCredentialsClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The [Unique identifier](https://tools.ietf.org/html/rfc6749#section-2.2) issued by the AS for the Okta IdP instance | [optional] 
**ClientSecret** | Pointer to **string** | The [Client secret](https://tools.ietf.org/html/rfc6749#section-2.3.1) issued by the AS for the Okta IdP instance | [optional] 
**PkceRequired** | Pointer to **bool** | Require Proof Key for Code Exchange (PKCE) for additional verification | [optional] 
**TokenEndpointAuthMethod** | Pointer to **string** | Client authentication methods supported by the token endpoint | [optional] 

## Methods

### NewOAuthCredentialsClient

`func NewOAuthCredentialsClient() *OAuthCredentialsClient`

NewOAuthCredentialsClient instantiates a new OAuthCredentialsClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthCredentialsClientWithDefaults

`func NewOAuthCredentialsClientWithDefaults() *OAuthCredentialsClient`

NewOAuthCredentialsClientWithDefaults instantiates a new OAuthCredentialsClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthCredentialsClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthCredentialsClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthCredentialsClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAuthCredentialsClient) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAuthCredentialsClient) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthCredentialsClient) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthCredentialsClient) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthCredentialsClient) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetPkceRequired

`func (o *OAuthCredentialsClient) GetPkceRequired() bool`

GetPkceRequired returns the PkceRequired field if non-nil, zero value otherwise.

### GetPkceRequiredOk

`func (o *OAuthCredentialsClient) GetPkceRequiredOk() (*bool, bool)`

GetPkceRequiredOk returns a tuple with the PkceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceRequired

`func (o *OAuthCredentialsClient) SetPkceRequired(v bool)`

SetPkceRequired sets PkceRequired field to given value.

### HasPkceRequired

`func (o *OAuthCredentialsClient) HasPkceRequired() bool`

HasPkceRequired returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *OAuthCredentialsClient) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *OAuthCredentialsClient) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *OAuthCredentialsClient) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *OAuthCredentialsClient) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


