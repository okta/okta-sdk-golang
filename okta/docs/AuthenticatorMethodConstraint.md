# AuthenticatorMethodConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAuthenticators** | Pointer to [**[]AuthenticatorIdentity**](AuthenticatorIdentity.md) |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthenticatorMethodConstraint

`func NewAuthenticatorMethodConstraint() *AuthenticatorMethodConstraint`

NewAuthenticatorMethodConstraint instantiates a new AuthenticatorMethodConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticatorMethodConstraintWithDefaults

`func NewAuthenticatorMethodConstraintWithDefaults() *AuthenticatorMethodConstraint`

NewAuthenticatorMethodConstraintWithDefaults instantiates a new AuthenticatorMethodConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAuthenticators

`func (o *AuthenticatorMethodConstraint) GetAllowedAuthenticators() []AuthenticatorIdentity`

GetAllowedAuthenticators returns the AllowedAuthenticators field if non-nil, zero value otherwise.

### GetAllowedAuthenticatorsOk

`func (o *AuthenticatorMethodConstraint) GetAllowedAuthenticatorsOk() (*[]AuthenticatorIdentity, bool)`

GetAllowedAuthenticatorsOk returns a tuple with the AllowedAuthenticators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAuthenticators

`func (o *AuthenticatorMethodConstraint) SetAllowedAuthenticators(v []AuthenticatorIdentity)`

SetAllowedAuthenticators sets AllowedAuthenticators field to given value.

### HasAllowedAuthenticators

`func (o *AuthenticatorMethodConstraint) HasAllowedAuthenticators() bool`

HasAllowedAuthenticators returns a boolean if a field has been set.

### GetMethod

`func (o *AuthenticatorMethodConstraint) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AuthenticatorMethodConstraint) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AuthenticatorMethodConstraint) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AuthenticatorMethodConstraint) HasMethod() bool`

HasMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


