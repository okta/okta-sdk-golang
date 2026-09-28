# OAUTH2JWTBEARERGRANT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | The URI of the authorization server that verifies the token. This is typically the token URI of your JWT. | [optional] 
**ClientId** | Pointer to **string** | The client ID that&#39;s used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider. | [optional] 
**Issuer** | Pointer to **string** | The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value. | [optional] 
**KeyId** | Pointer to **string** | The ID of the private key that&#39;s used to sign the JWT | [optional] 
**PrivateKey** | Pointer to **string** | The secret RSA key that&#39;s used to cryptographically sign the JWT | [optional] 
**Scopes** | Pointer to **[]string** | List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails. | [optional] 
**SigningAlgorithm** | Pointer to **string** | The signing algorithm that&#39;s used to sign the JWT | [optional] 
**Subject** | Pointer to **string** | The email address of the user account that the OAuth 2.0 app impersonates to send emails | [optional] 
**TokenEndpoint** | Pointer to **string** | The email provider&#39;s specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token | [optional] 

## Methods

### NewOAUTH2JWTBEARERGRANT

`func NewOAUTH2JWTBEARERGRANT() *OAUTH2JWTBEARERGRANT`

NewOAUTH2JWTBEARERGRANT instantiates a new OAUTH2JWTBEARERGRANT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAUTH2JWTBEARERGRANTWithDefaults

`func NewOAUTH2JWTBEARERGRANTWithDefaults() *OAUTH2JWTBEARERGRANT`

NewOAUTH2JWTBEARERGRANTWithDefaults instantiates a new OAUTH2JWTBEARERGRANT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *OAUTH2JWTBEARERGRANT) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *OAUTH2JWTBEARERGRANT) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *OAUTH2JWTBEARERGRANT) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *OAUTH2JWTBEARERGRANT) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetClientId

`func (o *OAUTH2JWTBEARERGRANT) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAUTH2JWTBEARERGRANT) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAUTH2JWTBEARERGRANT) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *OAUTH2JWTBEARERGRANT) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetIssuer

`func (o *OAUTH2JWTBEARERGRANT) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAUTH2JWTBEARERGRANT) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAUTH2JWTBEARERGRANT) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *OAUTH2JWTBEARERGRANT) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetKeyId

`func (o *OAUTH2JWTBEARERGRANT) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *OAUTH2JWTBEARERGRANT) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *OAUTH2JWTBEARERGRANT) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *OAUTH2JWTBEARERGRANT) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetPrivateKey

`func (o *OAUTH2JWTBEARERGRANT) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *OAUTH2JWTBEARERGRANT) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *OAUTH2JWTBEARERGRANT) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *OAUTH2JWTBEARERGRANT) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetScopes

`func (o *OAUTH2JWTBEARERGRANT) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAUTH2JWTBEARERGRANT) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAUTH2JWTBEARERGRANT) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *OAUTH2JWTBEARERGRANT) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *OAUTH2JWTBEARERGRANT) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *OAUTH2JWTBEARERGRANT) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *OAUTH2JWTBEARERGRANT) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *OAUTH2JWTBEARERGRANT) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### GetSubject

`func (o *OAUTH2JWTBEARERGRANT) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAUTH2JWTBEARERGRANT) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAUTH2JWTBEARERGRANT) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OAUTH2JWTBEARERGRANT) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *OAUTH2JWTBEARERGRANT) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAUTH2JWTBEARERGRANT) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAUTH2JWTBEARERGRANT) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *OAUTH2JWTBEARERGRANT) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


