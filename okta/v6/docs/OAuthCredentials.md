# OAuthCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**OAuthCredentialsClient**](OAuthCredentialsClient.md) |  | [optional] 
**Signing** | Pointer to [**AppleClientSigning**](AppleClientSigning.md) |  | [optional] 

## Methods

### NewOAuthCredentials

`func NewOAuthCredentials() *OAuthCredentials`

NewOAuthCredentials instantiates a new OAuthCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthCredentialsWithDefaults

`func NewOAuthCredentialsWithDefaults() *OAuthCredentials`

NewOAuthCredentialsWithDefaults instantiates a new OAuthCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *OAuthCredentials) GetClient() OAuthCredentialsClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *OAuthCredentials) GetClientOk() (*OAuthCredentialsClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *OAuthCredentials) SetClient(v OAuthCredentialsClient)`

SetClient sets Client field to given value.

### HasClient

`func (o *OAuthCredentials) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetSigning

`func (o *OAuthCredentials) GetSigning() AppleClientSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *OAuthCredentials) GetSigningOk() (*AppleClientSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *OAuthCredentials) SetSigning(v AppleClientSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *OAuthCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


