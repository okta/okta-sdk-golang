# OAUTH2CLIENTCREDENTIALSREQ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The client ID that&#39;s used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider. | 
**ClientSecret** | **string** | The client secret that&#39;s used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider. | 
**Scopes** | **[]string** | List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails. | 
**TokenEndpoint** | **string** | The email provider&#39;s specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token | 
**TokenEndpointAuthMethod** | **string** | This method determines how your OAuth 2.0 app sends its credentials (&#x60;client_id&#x60; and &#x60;client_secret&#x60;) to the provider&#39;s server when requesting an access token. | 

## Methods

### NewOAUTH2CLIENTCREDENTIALSREQ

`func NewOAUTH2CLIENTCREDENTIALSREQ(clientId string, clientSecret string, scopes []string, tokenEndpoint string, tokenEndpointAuthMethod string, ) *OAUTH2CLIENTCREDENTIALSREQ`

NewOAUTH2CLIENTCREDENTIALSREQ instantiates a new OAUTH2CLIENTCREDENTIALSREQ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAUTH2CLIENTCREDENTIALSREQWithDefaults

`func NewOAUTH2CLIENTCREDENTIALSREQWithDefaults() *OAUTH2CLIENTCREDENTIALSREQ`

NewOAUTH2CLIENTCREDENTIALSREQWithDefaults instantiates a new OAUTH2CLIENTCREDENTIALSREQ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAUTH2CLIENTCREDENTIALSREQ) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAUTH2CLIENTCREDENTIALSREQ) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetScopes

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAUTH2CLIENTCREDENTIALSREQ) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTokenEndpoint

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAUTH2CLIENTCREDENTIALSREQ) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetTokenEndpointAuthMethod

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *OAUTH2CLIENTCREDENTIALSREQ) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *OAUTH2CLIENTCREDENTIALSREQ) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


