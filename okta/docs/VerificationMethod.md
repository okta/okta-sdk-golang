# VerificationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]AccessPolicyConstraints**](AccessPolicyConstraints.md) |  | [optional] 
**FactorMode** | Pointer to **string** |  | [optional] 
**ReauthenticateIn** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewVerificationMethod

`func NewVerificationMethod() *VerificationMethod`

NewVerificationMethod instantiates a new VerificationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationMethodWithDefaults

`func NewVerificationMethodWithDefaults() *VerificationMethod`

NewVerificationMethodWithDefaults instantiates a new VerificationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *VerificationMethod) GetConstraints() []AccessPolicyConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *VerificationMethod) GetConstraintsOk() (*[]AccessPolicyConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *VerificationMethod) SetConstraints(v []AccessPolicyConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *VerificationMethod) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetFactorMode

`func (o *VerificationMethod) GetFactorMode() string`

GetFactorMode returns the FactorMode field if non-nil, zero value otherwise.

### GetFactorModeOk

`func (o *VerificationMethod) GetFactorModeOk() (*string, bool)`

GetFactorModeOk returns a tuple with the FactorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorMode

`func (o *VerificationMethod) SetFactorMode(v string)`

SetFactorMode sets FactorMode field to given value.

### HasFactorMode

`func (o *VerificationMethod) HasFactorMode() bool`

HasFactorMode returns a boolean if a field has been set.

### GetReauthenticateIn

`func (o *VerificationMethod) GetReauthenticateIn() string`

GetReauthenticateIn returns the ReauthenticateIn field if non-nil, zero value otherwise.

### GetReauthenticateInOk

`func (o *VerificationMethod) GetReauthenticateInOk() (*string, bool)`

GetReauthenticateInOk returns a tuple with the ReauthenticateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticateIn

`func (o *VerificationMethod) SetReauthenticateIn(v string)`

SetReauthenticateIn sets ReauthenticateIn field to given value.

### HasReauthenticateIn

`func (o *VerificationMethod) HasReauthenticateIn() bool`

HasReauthenticateIn returns a boolean if a field has been set.

### GetType

`func (o *VerificationMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VerificationMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VerificationMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VerificationMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


