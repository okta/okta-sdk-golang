# DevicePolicyRuleConditionPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMDMFrameworks** | Pointer to [**[]DevicePolicyMDMFramework**](DevicePolicyMDMFramework.md) |  | [optional] 
**Types** | Pointer to [**[]DevicePolicyPlatformType**](DevicePolicyPlatformType.md) |  | [optional] 

## Methods

### NewDevicePolicyRuleConditionPlatform

`func NewDevicePolicyRuleConditionPlatform() *DevicePolicyRuleConditionPlatform`

NewDevicePolicyRuleConditionPlatform instantiates a new DevicePolicyRuleConditionPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePolicyRuleConditionPlatformWithDefaults

`func NewDevicePolicyRuleConditionPlatformWithDefaults() *DevicePolicyRuleConditionPlatform`

NewDevicePolicyRuleConditionPlatformWithDefaults instantiates a new DevicePolicyRuleConditionPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMDMFrameworks

`func (o *DevicePolicyRuleConditionPlatform) GetSupportedMDMFrameworks() []DevicePolicyMDMFramework`

GetSupportedMDMFrameworks returns the SupportedMDMFrameworks field if non-nil, zero value otherwise.

### GetSupportedMDMFrameworksOk

`func (o *DevicePolicyRuleConditionPlatform) GetSupportedMDMFrameworksOk() (*[]DevicePolicyMDMFramework, bool)`

GetSupportedMDMFrameworksOk returns a tuple with the SupportedMDMFrameworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMDMFrameworks

`func (o *DevicePolicyRuleConditionPlatform) SetSupportedMDMFrameworks(v []DevicePolicyMDMFramework)`

SetSupportedMDMFrameworks sets SupportedMDMFrameworks field to given value.

### HasSupportedMDMFrameworks

`func (o *DevicePolicyRuleConditionPlatform) HasSupportedMDMFrameworks() bool`

HasSupportedMDMFrameworks returns a boolean if a field has been set.

### GetTypes

`func (o *DevicePolicyRuleConditionPlatform) GetTypes() []DevicePolicyPlatformType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *DevicePolicyRuleConditionPlatform) GetTypesOk() (*[]DevicePolicyPlatformType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *DevicePolicyRuleConditionPlatform) SetTypes(v []DevicePolicyPlatformType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *DevicePolicyRuleConditionPlatform) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


