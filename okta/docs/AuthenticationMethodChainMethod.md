# AuthenticationMethodChainMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chains** | Pointer to [**[]AuthenticationMethodChain**](AuthenticationMethodChain.md) | Authentication method chains. Only supports 5 items in the array. Each chain can support maximum 3 steps. | [optional] 
**ReauthenticateIn** | Pointer to **string** | Specifies how often the user is prompted for authentication using duration format for the time period. For example, &#x60;PT2H30M&#x60; for two and a half hours. Don&#39;t set this parameter if you&#39;re setting the &#x60;reauthenticateIn&#x60; parameter in &#x60;chains&#x60;. | [optional] 

## Methods

### NewAuthenticationMethodChainMethod

`func NewAuthenticationMethodChainMethod() *AuthenticationMethodChainMethod`

NewAuthenticationMethodChainMethod instantiates a new AuthenticationMethodChainMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodChainMethodWithDefaults

`func NewAuthenticationMethodChainMethodWithDefaults() *AuthenticationMethodChainMethod`

NewAuthenticationMethodChainMethodWithDefaults instantiates a new AuthenticationMethodChainMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChains

`func (o *AuthenticationMethodChainMethod) GetChains() []AuthenticationMethodChain`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *AuthenticationMethodChainMethod) GetChainsOk() (*[]AuthenticationMethodChain, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *AuthenticationMethodChainMethod) SetChains(v []AuthenticationMethodChain)`

SetChains sets Chains field to given value.

### HasChains

`func (o *AuthenticationMethodChainMethod) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetReauthenticateIn

`func (o *AuthenticationMethodChainMethod) GetReauthenticateIn() string`

GetReauthenticateIn returns the ReauthenticateIn field if non-nil, zero value otherwise.

### GetReauthenticateInOk

`func (o *AuthenticationMethodChainMethod) GetReauthenticateInOk() (*string, bool)`

GetReauthenticateInOk returns a tuple with the ReauthenticateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticateIn

`func (o *AuthenticationMethodChainMethod) SetReauthenticateIn(v string)`

SetReauthenticateIn sets ReauthenticateIn field to given value.

### HasReauthenticateIn

`func (o *AuthenticationMethodChainMethod) HasReauthenticateIn() bool`

HasReauthenticateIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


