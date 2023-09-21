# MultifactorEnrollmentPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**PolicyRuleConditions**](PolicyRuleConditions.md) |  | [optional] 
**Settings** | Pointer to [**MultifactorEnrollmentPolicySettings**](MultifactorEnrollmentPolicySettings.md) |  | [optional] 

## Methods

### NewMultifactorEnrollmentPolicy

`func NewMultifactorEnrollmentPolicy() *MultifactorEnrollmentPolicy`

NewMultifactorEnrollmentPolicy instantiates a new MultifactorEnrollmentPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultifactorEnrollmentPolicyWithDefaults

`func NewMultifactorEnrollmentPolicyWithDefaults() *MultifactorEnrollmentPolicy`

NewMultifactorEnrollmentPolicyWithDefaults instantiates a new MultifactorEnrollmentPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *MultifactorEnrollmentPolicy) GetConditions() PolicyRuleConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MultifactorEnrollmentPolicy) GetConditionsOk() (*PolicyRuleConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MultifactorEnrollmentPolicy) SetConditions(v PolicyRuleConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *MultifactorEnrollmentPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetSettings

`func (o *MultifactorEnrollmentPolicy) GetSettings() MultifactorEnrollmentPolicySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *MultifactorEnrollmentPolicy) GetSettingsOk() (*MultifactorEnrollmentPolicySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *MultifactorEnrollmentPolicy) SetSettings(v MultifactorEnrollmentPolicySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *MultifactorEnrollmentPolicy) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


