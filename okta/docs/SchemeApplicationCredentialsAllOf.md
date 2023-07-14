# SchemeApplicationCredentialsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to [**PasswordCredential**](PasswordCredential.md) |  | [optional] 
**RevealPassword** | Pointer to **bool** |  | [optional] 
**Scheme** | Pointer to [**ApplicationCredentialsScheme**](ApplicationCredentialsScheme.md) |  | [optional] 
**Signing** | Pointer to [**ApplicationCredentialsSigning**](ApplicationCredentialsSigning.md) |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewSchemeApplicationCredentialsAllOf

`func NewSchemeApplicationCredentialsAllOf() *SchemeApplicationCredentialsAllOf`

NewSchemeApplicationCredentialsAllOf instantiates a new SchemeApplicationCredentialsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemeApplicationCredentialsAllOfWithDefaults

`func NewSchemeApplicationCredentialsAllOfWithDefaults() *SchemeApplicationCredentialsAllOf`

NewSchemeApplicationCredentialsAllOfWithDefaults instantiates a new SchemeApplicationCredentialsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *SchemeApplicationCredentialsAllOf) GetPassword() PasswordCredential`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SchemeApplicationCredentialsAllOf) GetPasswordOk() (*PasswordCredential, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SchemeApplicationCredentialsAllOf) SetPassword(v PasswordCredential)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SchemeApplicationCredentialsAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRevealPassword

`func (o *SchemeApplicationCredentialsAllOf) GetRevealPassword() bool`

GetRevealPassword returns the RevealPassword field if non-nil, zero value otherwise.

### GetRevealPasswordOk

`func (o *SchemeApplicationCredentialsAllOf) GetRevealPasswordOk() (*bool, bool)`

GetRevealPasswordOk returns a tuple with the RevealPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevealPassword

`func (o *SchemeApplicationCredentialsAllOf) SetRevealPassword(v bool)`

SetRevealPassword sets RevealPassword field to given value.

### HasRevealPassword

`func (o *SchemeApplicationCredentialsAllOf) HasRevealPassword() bool`

HasRevealPassword returns a boolean if a field has been set.

### GetScheme

`func (o *SchemeApplicationCredentialsAllOf) GetScheme() ApplicationCredentialsScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *SchemeApplicationCredentialsAllOf) GetSchemeOk() (*ApplicationCredentialsScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *SchemeApplicationCredentialsAllOf) SetScheme(v ApplicationCredentialsScheme)`

SetScheme sets Scheme field to given value.

### HasScheme

`func (o *SchemeApplicationCredentialsAllOf) HasScheme() bool`

HasScheme returns a boolean if a field has been set.

### GetSigning

`func (o *SchemeApplicationCredentialsAllOf) GetSigning() ApplicationCredentialsSigning`

GetSigning returns the Signing field if non-nil, zero value otherwise.

### GetSigningOk

`func (o *SchemeApplicationCredentialsAllOf) GetSigningOk() (*ApplicationCredentialsSigning, bool)`

GetSigningOk returns a tuple with the Signing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigning

`func (o *SchemeApplicationCredentialsAllOf) SetSigning(v ApplicationCredentialsSigning)`

SetSigning sets Signing field to given value.

### HasSigning

`func (o *SchemeApplicationCredentialsAllOf) HasSigning() bool`

HasSigning returns a boolean if a field has been set.

### GetUserName

`func (o *SchemeApplicationCredentialsAllOf) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SchemeApplicationCredentialsAllOf) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SchemeApplicationCredentialsAllOf) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SchemeApplicationCredentialsAllOf) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


