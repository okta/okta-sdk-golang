# SsprPrimaryRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodConstraints** | Pointer to [**[]AuthenticatorMethodConstraint**](AuthenticatorMethodConstraint.md) | Constraints on the values specified in the &#x60;methods&#x60; array. Specifying a constraint limits methods to specific authenticator(s). Currently, Google OTP is the only accepted constraint. | [optional] 
**Methods** | Pointer to **[]string** | Authenticator methods allowed for the initial authentication step of password recovery. Method &#x60;otp&#x60; requires a constraint limiting it to a Google authenticator. | [optional] 

## Methods

### NewSsprPrimaryRequirement

`func NewSsprPrimaryRequirement() *SsprPrimaryRequirement`

NewSsprPrimaryRequirement instantiates a new SsprPrimaryRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprPrimaryRequirementWithDefaults

`func NewSsprPrimaryRequirementWithDefaults() *SsprPrimaryRequirement`

NewSsprPrimaryRequirementWithDefaults instantiates a new SsprPrimaryRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodConstraints

`func (o *SsprPrimaryRequirement) GetMethodConstraints() []AuthenticatorMethodConstraint`

GetMethodConstraints returns the MethodConstraints field if non-nil, zero value otherwise.

### GetMethodConstraintsOk

`func (o *SsprPrimaryRequirement) GetMethodConstraintsOk() (*[]AuthenticatorMethodConstraint, bool)`

GetMethodConstraintsOk returns a tuple with the MethodConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodConstraints

`func (o *SsprPrimaryRequirement) SetMethodConstraints(v []AuthenticatorMethodConstraint)`

SetMethodConstraints sets MethodConstraints field to given value.

### HasMethodConstraints

`func (o *SsprPrimaryRequirement) HasMethodConstraints() bool`

HasMethodConstraints returns a boolean if a field has been set.

### GetMethods

`func (o *SsprPrimaryRequirement) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *SsprPrimaryRequirement) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *SsprPrimaryRequirement) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *SsprPrimaryRequirement) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


