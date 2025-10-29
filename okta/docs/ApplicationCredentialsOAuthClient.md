# ApplicationCredentialsOAuthClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoKeyRotation** | Pointer to **bool** | Requested key rotation mode | [optional] [default to true]
**ClientId** | Pointer to **string** | Unique identifier for the OAuth 2.0 client app  &gt; **Notes:** &gt; * If you don&#39;t specify the &#x60;client_id&#x60;, this immutable property is populated with the [Application instance ID](/openapi/okta-management/management/tag/Application/#tag/Application/operation/getApplication!c&#x3D;200&amp;path&#x3D;4/id&amp;t&#x3D;response). &gt; * The &#x60;client_id&#x60; must consist of alphanumeric characters or the following special characters: &#x60;$-_.+!*&#39;(),&#x60;. &gt; * You can&#39;t use the reserved word &#x60;ALL_CLIENTS&#x60;. | [optional] 
**ClientSecret** | Pointer to **string** | OAuth 2.0 client secret string (used for confidential clients)  &gt; **Notes:** If a &#x60;client_secret&#x60; isn&#39;t provided on creation, and the &#x60;token_endpoint_auth_method&#x60; requires one, Okta generates a random &#x60;client_secret&#x60; for the client app. &gt; The &#x60;client_secret&#x60; is only shown when an OAuth 2.0 client app is created or updated (and only if the &#x60;token_endpoint_auth_method&#x60; requires a client secret). | [optional] 
**PkceRequired** | Pointer to **bool** | Requires Proof Key for Code Exchange (PKCE) for additional verification. If &#x60;token_endpoint_auth_method&#x60; is &#x60;none&#x60;, then &#x60;pkce_required&#x60; must be &#x60;true&#x60;. The default is &#x60;true&#x60; for browser and native app types. | [optional] [default to true]
**TokenEndpointAuthMethod** | Pointer to **string** | Requested authentication method for the token endpoint | [optional] [default to "client_secret_basic"]

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

### GetPkceRequired

`func (o *ApplicationCredentialsOAuthClient) GetPkceRequired() bool`

GetPkceRequired returns the PkceRequired field if non-nil, zero value otherwise.

### GetPkceRequiredOk

`func (o *ApplicationCredentialsOAuthClient) GetPkceRequiredOk() (*bool, bool)`

GetPkceRequiredOk returns a tuple with the PkceRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkceRequired

`func (o *ApplicationCredentialsOAuthClient) SetPkceRequired(v bool)`

SetPkceRequired sets PkceRequired field to given value.

### HasPkceRequired

`func (o *ApplicationCredentialsOAuthClient) HasPkceRequired() bool`

HasPkceRequired returns a boolean if a field has been set.

### GetTokenEndpointAuthMethod

`func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethod() string`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ApplicationCredentialsOAuthClient) GetTokenEndpointAuthMethodOk() (*string, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ApplicationCredentialsOAuthClient) SetTokenEndpointAuthMethod(v string)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *ApplicationCredentialsOAuthClient) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


