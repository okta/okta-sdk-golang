# DeviceAccessPolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assurance** | Pointer to [**DevicePolicyRuleConditionAssurance**](DevicePolicyRuleConditionAssurance.md) |  | [optional] 
**Managed** | Pointer to **bool** | Indicates if the device is managed. A device is considered managed if it&#39;s part of a device management system. | [optional] 
**Registered** | Pointer to **bool** | Indicates if the device is registered. A device is registered if the User enrolls with Okta Verify that&#39;s installed on the device. When the &#x60;managed&#x60; property is passed, you must also include the &#x60;registered&#x60; property and set it to &#x60;true&#x60;.  | [optional] 

## Methods

### NewDeviceAccessPolicyRuleCondition

`func NewDeviceAccessPolicyRuleCondition() *DeviceAccessPolicyRuleCondition`

NewDeviceAccessPolicyRuleCondition instantiates a new DeviceAccessPolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAccessPolicyRuleConditionWithDefaults

`func NewDeviceAccessPolicyRuleConditionWithDefaults() *DeviceAccessPolicyRuleCondition`

NewDeviceAccessPolicyRuleConditionWithDefaults instantiates a new DeviceAccessPolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssurance

`func (o *DeviceAccessPolicyRuleCondition) GetAssurance() DevicePolicyRuleConditionAssurance`

GetAssurance returns the Assurance field if non-nil, zero value otherwise.

### GetAssuranceOk

`func (o *DeviceAccessPolicyRuleCondition) GetAssuranceOk() (*DevicePolicyRuleConditionAssurance, bool)`

GetAssuranceOk returns a tuple with the Assurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssurance

`func (o *DeviceAccessPolicyRuleCondition) SetAssurance(v DevicePolicyRuleConditionAssurance)`

SetAssurance sets Assurance field to given value.

### HasAssurance

`func (o *DeviceAccessPolicyRuleCondition) HasAssurance() bool`

HasAssurance returns a boolean if a field has been set.

### GetManaged

`func (o *DeviceAccessPolicyRuleCondition) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeviceAccessPolicyRuleCondition) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeviceAccessPolicyRuleCondition) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *DeviceAccessPolicyRuleCondition) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetRegistered

`func (o *DeviceAccessPolicyRuleCondition) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *DeviceAccessPolicyRuleCondition) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *DeviceAccessPolicyRuleCondition) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *DeviceAccessPolicyRuleCondition) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


