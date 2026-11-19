# ImportRoadblockEntitlementSettingsEntitlementRoadblock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether the entitlement roadblock is enabled for the app instance | [optional] 
**PercentageThreshold** | Pointer to [**ImportRoadblockEntitlementThreshold**](ImportRoadblockEntitlementThreshold.md) |  | [optional] 
**AbsoluteThreshold** | Pointer to [**ImportRoadblockEntitlementThreshold**](ImportRoadblockEntitlementThreshold.md) |  | [optional] 
**BaselineCaptured** | Pointer to **bool** | Indicates whether an entitlement baseline has been captured for the app instance. Changes are only measured against a baseline. | [optional] 

## Methods

### NewImportRoadblockEntitlementSettingsEntitlementRoadblock

`func NewImportRoadblockEntitlementSettingsEntitlementRoadblock() *ImportRoadblockEntitlementSettingsEntitlementRoadblock`

NewImportRoadblockEntitlementSettingsEntitlementRoadblock instantiates a new ImportRoadblockEntitlementSettingsEntitlementRoadblock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockEntitlementSettingsEntitlementRoadblockWithDefaults

`func NewImportRoadblockEntitlementSettingsEntitlementRoadblockWithDefaults() *ImportRoadblockEntitlementSettingsEntitlementRoadblock`

NewImportRoadblockEntitlementSettingsEntitlementRoadblockWithDefaults instantiates a new ImportRoadblockEntitlementSettingsEntitlementRoadblock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPercentageThreshold

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetPercentageThreshold() ImportRoadblockEntitlementThreshold`

GetPercentageThreshold returns the PercentageThreshold field if non-nil, zero value otherwise.

### GetPercentageThresholdOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetPercentageThresholdOk() (*ImportRoadblockEntitlementThreshold, bool)`

GetPercentageThresholdOk returns a tuple with the PercentageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageThreshold

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetPercentageThreshold(v ImportRoadblockEntitlementThreshold)`

SetPercentageThreshold sets PercentageThreshold field to given value.

### HasPercentageThreshold

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasPercentageThreshold() bool`

HasPercentageThreshold returns a boolean if a field has been set.

### GetAbsoluteThreshold

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetAbsoluteThreshold() ImportRoadblockEntitlementThreshold`

GetAbsoluteThreshold returns the AbsoluteThreshold field if non-nil, zero value otherwise.

### GetAbsoluteThresholdOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetAbsoluteThresholdOk() (*ImportRoadblockEntitlementThreshold, bool)`

GetAbsoluteThresholdOk returns a tuple with the AbsoluteThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteThreshold

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetAbsoluteThreshold(v ImportRoadblockEntitlementThreshold)`

SetAbsoluteThreshold sets AbsoluteThreshold field to given value.

### HasAbsoluteThreshold

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasAbsoluteThreshold() bool`

HasAbsoluteThreshold returns a boolean if a field has been set.

### GetBaselineCaptured

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetBaselineCaptured() bool`

GetBaselineCaptured returns the BaselineCaptured field if non-nil, zero value otherwise.

### GetBaselineCapturedOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) GetBaselineCapturedOk() (*bool, bool)`

GetBaselineCapturedOk returns a tuple with the BaselineCaptured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineCaptured

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) SetBaselineCaptured(v bool)`

SetBaselineCaptured sets BaselineCaptured field to given value.

### HasBaselineCaptured

`func (o *ImportRoadblockEntitlementSettingsEntitlementRoadblock) HasBaselineCaptured() bool`

HasBaselineCaptured returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


