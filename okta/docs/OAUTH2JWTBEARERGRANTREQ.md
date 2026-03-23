# OAUTH2JWTBEARERGRANTREQ

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** | The URI of the authorization server that verifies the token. This is typically the token URI of your JWT. | 
**ClientId** | **string** | The client ID that&#39;s used to access your SMTP server. This client ID is obtained when you create an OAuth 2.0 app with your email provider. | 
**Issuer** | **string** | The unique ID of the entity that creates the JWT. This can sometimes be the email address of the user who creates the JWT. Check with your email provider for the correct value. | 
**KeyId** | **string** | The ID of the private key that&#39;s used to sign the JWT | 
**PrivateKey** | **string** | The secret RSA key that&#39;s used to cryptographically sign the JWT | 
**Scopes** | **[]string** | List of OAuth 2.0 scopes for your SMTP server. You must provide a scope that allows your email server to send emails. | 
**SigningAlgorithm** | **string** | The signing algorithm that&#39;s used to sign the JWT | 
**Subject** | **string** | The email address of the user account that the OAuth 2.0 app impersonates to send emails | 
**TokenEndpoint** | **string** | The email provider&#39;s specific URL where the OAuth 2.0 app sends its credentials (or signed JWT) to exchange them for an access token | 

## Methods

### NewOAUTH2JWTBEARERGRANTREQ

`func NewOAUTH2JWTBEARERGRANTREQ(audience string, clientId string, issuer string, keyId string, privateKey string, scopes []string, signingAlgorithm string, subject string, tokenEndpoint string, ) *OAUTH2JWTBEARERGRANTREQ`

NewOAUTH2JWTBEARERGRANTREQ instantiates a new OAUTH2JWTBEARERGRANTREQ object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAUTH2JWTBEARERGRANTREQWithDefaults

`func NewOAUTH2JWTBEARERGRANTREQWithDefaults() *OAUTH2JWTBEARERGRANTREQ`

NewOAUTH2JWTBEARERGRANTREQWithDefaults instantiates a new OAUTH2JWTBEARERGRANTREQ object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *OAUTH2JWTBEARERGRANTREQ) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *OAUTH2JWTBEARERGRANTREQ) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetClientId

`func (o *OAUTH2JWTBEARERGRANTREQ) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAUTH2JWTBEARERGRANTREQ) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetIssuer

`func (o *OAUTH2JWTBEARERGRANTREQ) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *OAUTH2JWTBEARERGRANTREQ) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetKeyId

`func (o *OAUTH2JWTBEARERGRANTREQ) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *OAUTH2JWTBEARERGRANTREQ) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetPrivateKey

`func (o *OAUTH2JWTBEARERGRANTREQ) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *OAUTH2JWTBEARERGRANTREQ) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetScopes

`func (o *OAUTH2JWTBEARERGRANTREQ) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *OAUTH2JWTBEARERGRANTREQ) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetSigningAlgorithm

`func (o *OAUTH2JWTBEARERGRANTREQ) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *OAUTH2JWTBEARERGRANTREQ) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.


### GetSubject

`func (o *OAUTH2JWTBEARERGRANTREQ) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OAUTH2JWTBEARERGRANTREQ) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTokenEndpoint

`func (o *OAUTH2JWTBEARERGRANTREQ) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *OAUTH2JWTBEARERGRANTREQ) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *OAUTH2JWTBEARERGRANTREQ) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


