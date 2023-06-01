# DevicePolicyRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Migrated** | Pointer to **bool** |  | [optional] 
**Platform** | Pointer to [**DevicePolicyRuleConditionPlatform**](DevicePolicyRuleConditionPlatform.md) |  | [optional] 
**Rooted** | Pointer to **bool** |  | [optional] 
**TrustLevel** | Pointer to [**DevicePolicyTrustLevel**](DevicePolicyTrustLevel.md) |  | [optional] 

## Methods

### NewDevicePolicyRuleCondition

`func NewDevicePolicyRuleCondition() *DevicePolicyRuleCondition`

NewDevicePolicyRuleCondition instantiates a new DevicePolicyRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePolicyRuleConditionWithDefaults

`func NewDevicePolicyRuleConditionWithDefaults() *DevicePolicyRuleCondition`

NewDevicePolicyRuleConditionWithDefaults instantiates a new DevicePolicyRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrated

`func (o *DevicePolicyRuleCondition) GetMigrated() bool`

GetMigrated returns the Migrated field if non-nil, zero value otherwise.

### GetMigratedOk

`func (o *DevicePolicyRuleCondition) GetMigratedOk() (*bool, bool)`

GetMigratedOk returns a tuple with the Migrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrated

`func (o *DevicePolicyRuleCondition) SetMigrated(v bool)`

SetMigrated sets Migrated field to given value.

### HasMigrated

`func (o *DevicePolicyRuleCondition) HasMigrated() bool`

HasMigrated returns a boolean if a field has been set.

### GetPlatform

`func (o *DevicePolicyRuleCondition) GetPlatform() DevicePolicyRuleConditionPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DevicePolicyRuleCondition) GetPlatformOk() (*DevicePolicyRuleConditionPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DevicePolicyRuleCondition) SetPlatform(v DevicePolicyRuleConditionPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DevicePolicyRuleCondition) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetRooted

`func (o *DevicePolicyRuleCondition) GetRooted() bool`

GetRooted returns the Rooted field if non-nil, zero value otherwise.

### GetRootedOk

`func (o *DevicePolicyRuleCondition) GetRootedOk() (*bool, bool)`

GetRootedOk returns a tuple with the Rooted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooted

`func (o *DevicePolicyRuleCondition) SetRooted(v bool)`

SetRooted sets Rooted field to given value.

### HasRooted

`func (o *DevicePolicyRuleCondition) HasRooted() bool`

HasRooted returns a boolean if a field has been set.

### GetTrustLevel

`func (o *DevicePolicyRuleCondition) GetTrustLevel() DevicePolicyTrustLevel`

GetTrustLevel returns the TrustLevel field if non-nil, zero value otherwise.

### GetTrustLevelOk

`func (o *DevicePolicyRuleCondition) GetTrustLevelOk() (*DevicePolicyTrustLevel, bool)`

GetTrustLevelOk returns a tuple with the TrustLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustLevel

`func (o *DevicePolicyRuleCondition) SetTrustLevel(v DevicePolicyTrustLevel)`

SetTrustLevel sets TrustLevel field to given value.

### HasTrustLevel

`func (o *DevicePolicyRuleCondition) HasTrustLevel() bool`

HasTrustLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


