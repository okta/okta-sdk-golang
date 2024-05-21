# AuthorizationServerCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signing** | Pointer to [**AuthorizationServerCredentialsSigningConfig**](AuthorizationServerCredentialsSigningConfig.md) |  | [optional] 

## Methods

### NewAuthorizationServerCredentials

`func NewAuthorizationServerCredentials() *AuthorizationServerCredentials`

NewAuthorizationServerCredentials instantiates a new AuthorizationServerCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationServerCredentialsWithDefaults

`func NewAuthorizationServerCredentialsWithDefaults() *AuthorizationServerCredentials`

NewAuthorizationServerCredentialsWithDefaults instantiates a new AuthorizationServerCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigning

`func (o *AuthorizationServerCredentials) GetSigning() AuthorizationServerCredentialsSigningConfig`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *AuthorizationServerCredentials) GetSigningOk() (*AuthorizationServerCredentialsSigningConfig, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *AuthorizationServerCredentials) SetSigning(v AuthorizationServerCredentialsSigningConfig)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *AuthorizationServerCredentials) HasSigning() bool`

HasSigning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


