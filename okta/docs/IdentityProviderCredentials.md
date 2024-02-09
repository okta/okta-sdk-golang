# IdentityProviderCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**IdentityProviderCredentialsClient**](IdentityProviderCredentialsClient.md) |  | [optional] 
**Signing** | Pointer to [**IdentityProviderCredentialsSigning**](IdentityProviderCredentialsSigning.md) |  | [optional] 
**Trust** | Pointer to [**IdentityProviderCredentialsTrust**](IdentityProviderCredentialsTrust.md) |  | [optional] 

## Methods

### NewIdentityProviderCredentials

`func NewIdentityProviderCredentials() *IdentityProviderCredentials`

NewIdentityProviderCredentials instantiates a new IdentityProviderCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderCredentialsWithDefaults

`func NewIdentityProviderCredentialsWithDefaults() *IdentityProviderCredentials`

NewIdentityProviderCredentialsWithDefaults instantiates a new IdentityProviderCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *IdentityProviderCredentials) GetClient() IdentityProviderCredentialsClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *IdentityProviderCredentials) GetClientOk() (*IdentityProviderCredentialsClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *IdentityProviderCredentials) SetClient(v IdentityProviderCredentialsClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *IdentityProviderCredentials) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetSigning

`func (o *IdentityProviderCredentials) GetSigning() IdentityProviderCredentialsSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *IdentityProviderCredentials) GetSigningOk() (*IdentityProviderCredentialsSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *IdentityProviderCredentials) SetSigning(v IdentityProviderCredentialsSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *IdentityProviderCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetTrust

`func (o *IdentityProviderCredentials) GetTrust() IdentityProviderCredentialsTrust`

GetTrust returns the Trust field if non-nil, zero value otherwise.

### GetTrustOk

`func (o *IdentityProviderCredentials) GetTrustOk() (*IdentityProviderCredentialsTrust, bool)`

GetTrustOk returns a tuple with the Trust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrust

`func (o *IdentityProviderCredentials) SetTrust(v IdentityProviderCredentialsTrust)`

SetTrust sets Trust field to given value.

### HasTrust

`func (o *IdentityProviderCredentials) HasTrust() bool`

HasTrust returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


