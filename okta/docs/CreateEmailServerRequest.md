# CreateEmailServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Human-readable name for your SMTP server | 
**AuthType** | **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt; &lt;x-lifecycle class&#x3D;\&quot;oie\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;The authentication type that&#39;s used by your SMTP server | 
**Enabled** | **bool** | If &#x60;true&#x60;, all email traffic is routed through your SMTP server | 
**Host** | **string** | Hostname or IP address of your SMTP server | 
**Id** | Pointer to **string** | ID of your SMTP server | [optional] [readonly] 
**Port** | **int32** | Port number of your SMTP server | 
**Username** | **string** | Username that&#39;s used to access your SMTP server | 
**ClientId** | **string** | The client ID that&#39;s used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider. | 
**ClientSecret** | **string** | The client secret that&#39;s used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider. | 
**Scopes** | **[]string** | List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails. | 
**TokenEndpoint** | **string** | The email provider&#39;s specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token | 
**TokenEndpointAuthMethod** | **string** | This method determines how your OAuth 2.0 app sends its credentials (&#x60;client_id&#x60; and &#x60;client_secret&#x60;) to the provider&#39;s server when requesting an access token. | 
**Audience** | **string** | The URI of the authorization server that verifies the token. This is typically the token URI of your JWT. | 
**Issuer** | **string** | The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value. | 
**KeyId** | **string** | The ID of the private key that&#39;s used to sign the JWT | 
**PrivateKey** | **string** | The secret RSA key that&#39;s used to cryptographically sign the JWT | 
**SigningAlgorithm** | **string** | The signing algorithm that&#39;s used to sign the JWT | 
**Subject** | **string** | The email address of the user account that the OAuth 2.0 app impersonates to send emails | 

## Methods

### NewCreateEmailServerRequest

`func NewCreateEmailServerRequest(alias string, authType string, enabled bool, host string, port int32, username string, clientId string, clientSecret string, scopes []string, tokenEndpoint string, tokenEndpointAuthMethod string, audience string, issuer string, keyId string, privateKey string, signingAlgorithm string, subject string, ) *CreateEmailServerRequest`

NewCreateEmailServerRequest instantiates a new CreateEmailServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmailServerRequestWithDefaults

`func NewCreateEmailServerRequestWithDefaults() *CreateEmailServerRequest`

NewCreateEmailServerRequestWithDefaults instantiates a new CreateEmailServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *CreateEmailServerRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreateEmailServerRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreateEmailServerRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetAuthType

`func (o *CreateEmailServerRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *CreateEmailServerRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *CreateEmailServerRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.


### GetEnabled

`func (o *CreateEmailServerRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateEmailServerRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateEmailServerRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHost

`func (o *CreateEmailServerRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CreateEmailServerRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CreateEmailServerRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetId

`func (o *CreateEmailServerRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateEmailServerRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateEmailServerRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateEmailServerRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *CreateEmailServerRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateEmailServerRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateEmailServerRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *CreateEmailServerRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateEmailServerRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateEmailServerRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetClientId

`func (o *CreateEmailServerRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CreateEmailServerRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CreateEmailServerRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *CreateEmailServerRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CreateEmailServerRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CreateEmailServerRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetScopes

`func (o *CreateEmailServerRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CreateEmailServerRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CreateEmailServerRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTokenEndpoint

`func (o *CreateEmailServerRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *CreateEmailServerRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *CreateEmailServerRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetTokenEndpointAuthMethod

`func (o *CreateEmailServerRequest) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *CreateEmailServerRequest) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *CreateEmailServerRequest) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetAudience

`func (o *CreateEmailServerRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CreateEmailServerRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CreateEmailServerRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetIssuer

`func (o *CreateEmailServerRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *CreateEmailServerRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *CreateEmailServerRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetKeyId

`func (o *CreateEmailServerRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *CreateEmailServerRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *CreateEmailServerRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetPrivateKey

`func (o *CreateEmailServerRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreateEmailServerRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreateEmailServerRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetSigningAlgorithm

`func (o *CreateEmailServerRequest) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CreateEmailServerRequest) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CreateEmailServerRequest) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.


### GetSubject

`func (o *CreateEmailServerRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateEmailServerRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateEmailServerRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


