# ProtocolIdVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**IDVCredentials**](IDVCredentials.md) |  | [optional] 
**Endpoints** | Pointer to [**IDVEndpoints**](IDVEndpoints.md) |  | [optional] 
**Scopes** | Pointer to **[]string** | IdP-defined permission bundles to request delegated access from the user. &gt; **Note:** The [identity provider type](https://developer.okta.com/docs/api/openapi/okta-management/management/tag/IdentityProvider/#tag/IdentityProvider/operation/createIdentityProvider!path&#x3D;type&amp;t&#x3D;request) table lists the scopes that are supported for each IdP. | [optional] 
**Type** | Pointer to **string** | ID verification protocol | [optional] 

## Methods

### NewProtocolIdVerification

`func NewProtocolIdVerification() *ProtocolIdVerification`

NewProtocolIdVerification instantiates a new ProtocolIdVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtocolIdVerificationWithDefaults

`func NewProtocolIdVerificationWithDefaults() *ProtocolIdVerification`

NewProtocolIdVerificationWithDefaults instantiates a new ProtocolIdVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ProtocolIdVerification) GetCredentials() IDVCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ProtocolIdVerification) GetCredentialsOk() (*IDVCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ProtocolIdVerification) SetCredentials(v IDVCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ProtocolIdVerification) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetEndpoints

`func (o *ProtocolIdVerification) GetEndpoints() IDVEndpoints`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *ProtocolIdVerification) GetEndpointsOk() (*IDVEndpoints, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *ProtocolIdVerification) SetEndpoints(v IDVEndpoints)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *ProtocolIdVerification) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetScopes

`func (o *ProtocolIdVerification) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProtocolIdVerification) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProtocolIdVerification) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProtocolIdVerification) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetType

`func (o *ProtocolIdVerification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProtocolIdVerification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProtocolIdVerification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ProtocolIdVerification) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


