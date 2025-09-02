# AuthenticationMethodChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethods** | Pointer to [**[]AuthenticationMethod**](AuthenticationMethod.md) |  | [optional] 
**Next** | Pointer to **[]map[string]interface{}** | The next steps of the authentication method chain. This is an array of &#x60;AuthenticationMethodChain&#x60;. Only supports one item in the array. | [optional] 
**ReauthenticateIn** | Pointer to **string** | Specifies how often the user is prompted for authentication using duration format for the time period. For example, &#x60;PT2H30M&#x60; for two and a half hours. This parameter can&#39;t be set at the same time as the &#x60;reauthenticateIn&#x60; property on the &#x60;verificationMethod&#x60;. | [optional] 

## Methods

### NewAuthenticationMethodChain

`func NewAuthenticationMethodChain() *AuthenticationMethodChain`

NewAuthenticationMethodChain instantiates a new AuthenticationMethodChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodChainWithDefaults

`func NewAuthenticationMethodChainWithDefaults() *AuthenticationMethodChain`

NewAuthenticationMethodChainWithDefaults instantiates a new AuthenticationMethodChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationMethods

`func (o *AuthenticationMethodChain) GetAuthenticationMethods() []AuthenticationMethod`

GetAuthenticationMethods returns the AuthenticationMethods field if non-nil, zero value otherwise.

### GetAuthenticationMethodsOk

`func (o *AuthenticationMethodChain) GetAuthenticationMethodsOk() (*[]AuthenticationMethod, bool)`

GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethods

`func (o *AuthenticationMethodChain) SetAuthenticationMethods(v []AuthenticationMethod)`

SetAuthenticationMethods sets AuthenticationMethods field to given value.

### HasAuthenticationMethods

`func (o *AuthenticationMethodChain) HasAuthenticationMethods() bool`

HasAuthenticationMethods returns a boolean if a field has been set.

### GetNext

`func (o *AuthenticationMethodChain) GetNext() []map[string]interface{}`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *AuthenticationMethodChain) GetNextOk() (*[]map[string]interface{}, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *AuthenticationMethodChain) SetNext(v []map[string]interface{})`

SetNext sets Next field to given value.

### HasNext

`func (o *AuthenticationMethodChain) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetReauthenticateIn

`func (o *AuthenticationMethodChain) GetReauthenticateIn() string`

GetReauthenticateIn returns the ReauthenticateIn field if non-nil, zero value otherwise.

### GetReauthenticateInOk

`func (o *AuthenticationMethodChain) GetReauthenticateInOk() (*string, bool)`

GetReauthenticateInOk returns a tuple with the ReauthenticateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticateIn

`func (o *AuthenticationMethodChain) SetReauthenticateIn(v string)`

SetReauthenticateIn sets ReauthenticateIn field to given value.

### HasReauthenticateIn

`func (o *AuthenticationMethodChain) HasReauthenticateIn() bool`

HasReauthenticateIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


