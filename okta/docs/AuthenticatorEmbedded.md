# AuthenticatorEmbedded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to [**[]AuthenticatorMethodBase**](AuthenticatorMethodBase.md) |  | [optional] 
**Policies** | Pointer to [**[]Policy**](Policy.md) |  | [optional] 

## Methods

### NewAuthenticatorEmbedded

`func NewAuthenticatorEmbedded() *AuthenticatorEmbedded`

NewAuthenticatorEmbedded instantiates a new AuthenticatorEmbedded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorEmbeddedWithDefaults

`func NewAuthenticatorEmbeddedWithDefaults() *AuthenticatorEmbedded`

NewAuthenticatorEmbeddedWithDefaults instantiates a new AuthenticatorEmbedded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *AuthenticatorEmbedded) GetMethods() []AuthenticatorMethodBase`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *AuthenticatorEmbedded) GetMethodsOk() (*[]AuthenticatorMethodBase, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *AuthenticatorEmbedded) SetMethods(v []AuthenticatorMethodBase)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *AuthenticatorEmbedded) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetPolicies

`func (o *AuthenticatorEmbedded) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *AuthenticatorEmbedded) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *AuthenticatorEmbedded) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *AuthenticatorEmbedded) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


