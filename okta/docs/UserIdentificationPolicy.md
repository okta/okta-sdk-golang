# UserIdentificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to **NullableString** | Policy conditions aren&#39;t supported for this policy type. | [optional] 

## Methods

### NewUserIdentificationPolicy

`func NewUserIdentificationPolicy() *UserIdentificationPolicy`

NewUserIdentificationPolicy instantiates a new UserIdentificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentificationPolicyWithDefaults

`func NewUserIdentificationPolicyWithDefaults() *UserIdentificationPolicy`

NewUserIdentificationPolicyWithDefaults instantiates a new UserIdentificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *UserIdentificationPolicy) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UserIdentificationPolicy) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UserIdentificationPolicy) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UserIdentificationPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *UserIdentificationPolicy) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *UserIdentificationPolicy) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


