# AssuranceMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]AccessPolicyConstraints**](AccessPolicyConstraints.md) | Specifies constraints for the authenticator. Constraints are logically evaluated such that only one constraint object needs to be satisfied. But, within a constraint object, each constraint property must be satisfied. | [optional] 
**FactorMode** | Pointer to **string** |  | [optional] 
**InactivityPeriod** | Pointer to **string** | The inactivity duration after which the user must re-authenticate. Use the ISO 8601 period format (for example, PT2H). | [optional] 
**ReauthenticateIn** | Pointer to **string** | The duration after which the user must re-authenticate, regardless of user activity. Keep in mind that the re-authentication intervals for constraints take precedent over this value. Use the ISO 8601 period format for recurring time intervals (for example, PT2H, PT0S, PT43800H, and so on). | [optional] 

## Methods

### NewAssuranceMethod

`func NewAssuranceMethod() *AssuranceMethod`

NewAssuranceMethod instantiates a new AssuranceMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssuranceMethodWithDefaults

`func NewAssuranceMethodWithDefaults() *AssuranceMethod`

NewAssuranceMethodWithDefaults instantiates a new AssuranceMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *AssuranceMethod) GetConstraints() []AccessPolicyConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *AssuranceMethod) GetConstraintsOk() (*[]AccessPolicyConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *AssuranceMethod) SetConstraints(v []AccessPolicyConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *AssuranceMethod) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetFactorMode

`func (o *AssuranceMethod) GetFactorMode() string`

GetFactorMode returns the FactorMode field if non-nil, zero value otherwise.

### GetFactorModeOk

`func (o *AssuranceMethod) GetFactorModeOk() (*string, bool)`

GetFactorModeOk returns a tuple with the FactorMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorMode

`func (o *AssuranceMethod) SetFactorMode(v string)`

SetFactorMode sets FactorMode field to given value.

### HasFactorMode

`func (o *AssuranceMethod) HasFactorMode() bool`

HasFactorMode returns a boolean if a field has been set.

### GetInactivityPeriod

`func (o *AssuranceMethod) GetInactivityPeriod() string`

GetInactivityPeriod returns the InactivityPeriod field if non-nil, zero value otherwise.

### GetInactivityPeriodOk

`func (o *AssuranceMethod) GetInactivityPeriodOk() (*string, bool)`

GetInactivityPeriodOk returns a tuple with the InactivityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityPeriod

`func (o *AssuranceMethod) SetInactivityPeriod(v string)`

SetInactivityPeriod sets InactivityPeriod field to given value.

### HasInactivityPeriod

`func (o *AssuranceMethod) HasInactivityPeriod() bool`

HasInactivityPeriod returns a boolean if a field has been set.

### GetReauthenticateIn

`func (o *AssuranceMethod) GetReauthenticateIn() string`

GetReauthenticateIn returns the ReauthenticateIn field if non-nil, zero value otherwise.

### GetReauthenticateInOk

`func (o *AssuranceMethod) GetReauthenticateInOk() (*string, bool)`

GetReauthenticateInOk returns a tuple with the ReauthenticateIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReauthenticateIn

`func (o *AssuranceMethod) SetReauthenticateIn(v string)`

SetReauthenticateIn sets ReauthenticateIn field to given value.

### HasReauthenticateIn

`func (o *AssuranceMethod) HasReauthenticateIn() bool`

HasReauthenticateIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


