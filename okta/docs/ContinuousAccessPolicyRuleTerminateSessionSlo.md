# ContinuousAccessPolicyRuleTerminateSessionSlo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppSelectionMode** | Pointer to **string** | This property defines the session to terminate - everyone, no one, or a specific app instance. | [optional] 
**AppInstanceIds** | Pointer to **[]string** | This property defines the app instance access to terminate. Only include this property when &#x60;appSelectionMode&#x60; is set to &#x60;SPECIFIC&#x60;. | [optional] 

## Methods

### NewContinuousAccessPolicyRuleTerminateSessionSlo

`func NewContinuousAccessPolicyRuleTerminateSessionSlo() *ContinuousAccessPolicyRuleTerminateSessionSlo`

NewContinuousAccessPolicyRuleTerminateSessionSlo instantiates a new ContinuousAccessPolicyRuleTerminateSessionSlo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContinuousAccessPolicyRuleTerminateSessionSloWithDefaults

`func NewContinuousAccessPolicyRuleTerminateSessionSloWithDefaults() *ContinuousAccessPolicyRuleTerminateSessionSlo`

NewContinuousAccessPolicyRuleTerminateSessionSloWithDefaults instantiates a new ContinuousAccessPolicyRuleTerminateSessionSlo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppSelectionMode

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppSelectionMode() string`

GetAppSelectionMode returns the AppSelectionMode field if non-nil, zero value otherwise.

### GetAppSelectionModeOk

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppSelectionModeOk() (*string, bool)`

GetAppSelectionModeOk returns a tuple with the AppSelectionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSelectionMode

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) SetAppSelectionMode(v string)`

SetAppSelectionMode sets AppSelectionMode field to given value.

### HasAppSelectionMode

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) HasAppSelectionMode() bool`

HasAppSelectionMode returns a boolean if a field has been set.

### GetAppInstanceIds

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppInstanceIds() []string`

GetAppInstanceIds returns the AppInstanceIds field if non-nil, zero value otherwise.

### GetAppInstanceIdsOk

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) GetAppInstanceIdsOk() (*[]string, bool)`

GetAppInstanceIdsOk returns a tuple with the AppInstanceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceIds

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) SetAppInstanceIds(v []string)`

SetAppInstanceIds sets AppInstanceIds field to given value.

### HasAppInstanceIds

`func (o *ContinuousAccessPolicyRuleTerminateSessionSlo) HasAppInstanceIds() bool`

HasAppInstanceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


