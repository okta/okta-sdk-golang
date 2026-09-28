# GetEmailServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Human-readable name for your SMTP server | [optional] 
**Enabled** | Pointer to **bool** | If &#x60;true&#x60;, all email traffic is routed through your SMTP server | [optional] 
**Host** | Pointer to **string** | Hostname or IP address of your SMTP server | [optional] 
**Id** | Pointer to **string** | ID of your SMTP server | [optional] [readonly] 
**Port** | Pointer to **int32** | Port number of your SMTP server | [optional] 
**Username** | Pointer to **string** | Username that&#39;s used to access your SMTP server | [optional] 
**ClientId** | Pointer to **string** | The client ID that&#39;s used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider. | [optional] 
**ClientSecret** | Pointer to **string** | The client secret that&#39;s used to access your SMTP server. This client secret is obtained when you create an OAuth 2.0 app with your email provider. | [optional] 
**Scopes** | Pointer to **[]string** | List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails. | [optional] 
**TokenEndpoint** | Pointer to **string** | The email provider&#39;s specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token | [optional] 
**TokenEndpointAuthMethod** | Pointer to **string** | This method determines how your OAuth 2.0 app sends its credentials (&#x60;client_id&#x60; and &#x60;client_secret&#x60;) to the provider&#39;s server when requesting an access token | [optional] 
**Audience** | Pointer to **string** | The URI of the authorization server that verifies the token. This is typically the token URI of your JWT. | [optional] 
**Issuer** | Pointer to **string** | The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value. | [optional] 
**KeyId** | Pointer to **string** | The ID of the private key that&#39;s used to sign the JWT | [optional] 
**PrivateKey** | Pointer to **string** | The secret RSA key that&#39;s used to cryptographically sign the JWT | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm that&#39;s used to sign the JWT | [optional] 
**Subject** | Pointer to **string** | The email address of the user account that the OAuth 2.0 app impersonates to send emails | [optional] 

## Methods

### NewGetEmailServer200Response

`func NewGetEmailServer200Response() *GetEmailServer200Response`

NewGetEmailServer200Response instantiates a new GetEmailServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEmailServer200ResponseWithDefaults

`func NewGetEmailServer200ResponseWithDefaults() *GetEmailServer200Response`

NewGetEmailServer200ResponseWithDefaults instantiates a new GetEmailServer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *GetEmailServer200Response) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetEmailServer200Response) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetEmailServer200Response) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetEmailServer200Response) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetEnabled

`func (o *GetEmailServer200Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetEmailServer200Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetEmailServer200Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetEmailServer200Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *GetEmailServer200Response) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetEmailServer200Response) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetEmailServer200Response) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetEmailServer200Response) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *GetEmailServer200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEmailServer200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEmailServer200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetEmailServer200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *GetEmailServer200Response) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetEmailServer200Response) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetEmailServer200Response) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetEmailServer200Response) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *GetEmailServer200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetEmailServer200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetEmailServer200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetEmailServer200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetClientId

`func (o *GetEmailServer200Response) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GetEmailServer200Response) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GetEmailServer200Response) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GetEmailServer200Response) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *GetEmailServer200Response) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *GetEmailServer200Response) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *GetEmailServer200Response) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *GetEmailServer200Response) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScopes

`func (o *GetEmailServer200Response) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetEmailServer200Response) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetEmailServer200Response) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetEmailServer200Response) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *GetEmailServer200Response) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *GetEmailServer200Response) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *GetEmailServer200Response) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *GetEmailServer200Response) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *GetEmailServer200Response) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *GetEmailServer200Response) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *GetEmailServer200Response) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *GetEmailServer200Response) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.

### GetAudience

`func (o *GetEmailServer200Response) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *GetEmailServer200Response) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *GetEmailServer200Response) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *GetEmailServer200Response) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetIssuer

`func (o *GetEmailServer200Response) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetEmailServer200Response) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetEmailServer200Response) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *GetEmailServer200Response) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKeyId

`func (o *GetEmailServer200Response) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *GetEmailServer200Response) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *GetEmailServer200Response) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *GetEmailServer200Response) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GetEmailServer200Response) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GetEmailServer200Response) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GetEmailServer200Response) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GetEmailServer200Response) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *GetEmailServer200Response) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *GetEmailServer200Response) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *GetEmailServer200Response) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *GetEmailServer200Response) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetSubject

`func (o *GetEmailServer200Response) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetEmailServer200Response) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetEmailServer200Response) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetEmailServer200Response) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


