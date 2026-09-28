# OAUTH2CLIENTCREDENTIALS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The client ID that&#39;s used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret that&#39;s used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider. | [optional] 
**Scopes** | Pointer to **[]string** | List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails. | [optional] 
**TokenEndpoint** | Pointer to **string** | The email provider&#39;s specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token | [optional] 
**TokenEndpointAuthMethod** | Pointer to **string** | This method determines how your OAuth 2.0 app sends its credentials (&#x60;client_id&#x60; and &#x60;client_secret&#x60;) to the provider&#39;s server when requesting an access token | [optional] 

## Methods

### NewOAUTH2CLIENTCREDENTIALS

`func NewOAUTH2CLIENTCREDENTIALS() *OAUTH2CLIENTCREDENTIALS`

NewOAUTH2CLIENTCREDENTIALS instantiates a new OAUTH2CLIENTCREDENTIALS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAUTH2CLIENTCREDENTIALSWithDefaults

`func NewOAUTH2CLIENTCREDENTIALSWithDefaults() *OAUTH2CLIENTCREDENTIALS`

NewOAUTH2CLIENTCREDENTIALSWithDefaults instantiates a new OAUTH2CLIENTCREDENTIALS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAUTH2CLIENTCREDENTIALS) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAUTH2CLIENTCREDENTIALS) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAUTH2CLIENTCREDENTIALS) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAUTH2CLIENTCREDENTIALS) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *OAUTH2CLIENTCREDENTIALS) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAUTH2CLIENTCREDENTIALS) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAUTH2CLIENTCREDENTIALS) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAUTH2CLIENTCREDENTIALS) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScopes

`func (o *OAUTH2CLIENTCREDENTIALS) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAUTH2CLIENTCREDENTIALS) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAUTH2CLIENTCREDENTIALS) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAUTH2CLIENTCREDENTIALS) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAUTH2CLIENTCREDENTIALS) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OAUTH2CLIENTCREDENTIALS) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *OAUTH2CLIENTCREDENTIALS) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *OAUTH2CLIENTCREDENTIALS) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *OAUTH2CLIENTCREDENTIALS) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


