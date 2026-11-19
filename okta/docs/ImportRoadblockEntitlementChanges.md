# ImportRoadblockEntitlementChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteThreshold** | Pointer to **int32** | The absolute threshold (number of users) in effect when the roadblock was created. Only returned when the absolute threshold is enabled. | [optional] 
**ChangedUsers** | Pointer to **int32** | Number of users whose entitlements the import would change or remove (&#x60;modifiedUsers&#x60; + &#x60;deletedUsers&#x60;). Compared with &#x60;absoluteThreshold&#x60;. | [optional] 
**ChangePercentage** | Pointer to **float32** | The share of users whose entitlements the import would change or remove (&#x60;changedUsers&#x60; / &#x60;totalUsers&#x60; × 100). Compared with &#x60;percentageThreshold&#x60;. | [optional] 
**DeletedUsers** | Pointer to **int32** | Number of users whose entitlements the import would remove because they&#39;re no longer in the source | [optional] 
**ModifiedUsers** | Pointer to **int32** | Number of users whose entitlements the import would change | [optional] 
**PercentageThreshold** | Pointer to **float32** | The percentage threshold in effect when the roadblock was created. Only returned when the percentage threshold is enabled. | [optional] 
**ThresholdType** | Pointer to **string** | The threshold that the import crossed. If both thresholds were crossed, &#x60;PERCENTAGE&#x60; is returned. | [optional] 
**TotalUsers** | Pointer to **int32** | Number of users with entitlements before the import (the baseline). Used as the denominator for &#x60;changePercentage&#x60;. | [optional] 

## Methods

### NewImportRoadblockEntitlementChanges

`func NewImportRoadblockEntitlementChanges() *ImportRoadblockEntitlementChanges`

NewImportRoadblockEntitlementChanges instantiates a new ImportRoadblockEntitlementChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockEntitlementChangesWithDefaults

`func NewImportRoadblockEntitlementChangesWithDefaults() *ImportRoadblockEntitlementChanges`

NewImportRoadblockEntitlementChangesWithDefaults instantiates a new ImportRoadblockEntitlementChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteThreshold

`func (o *ImportRoadblockEntitlementChanges) GetAbsoluteThreshold() int32`

GetAbsoluteThreshold returns the AbsoluteThreshold field if non-nil, zero value otherwise.

### GetAbsoluteThresholdOk

`func (o *ImportRoadblockEntitlementChanges) GetAbsoluteThresholdOk() (*int32, bool)`

GetAbsoluteThresholdOk returns a tuple with the AbsoluteThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteThreshold

`func (o *ImportRoadblockEntitlementChanges) SetAbsoluteThreshold(v int32)`

SetAbsoluteThreshold sets AbsoluteThreshold field to given value.

### HasAbsoluteThreshold

`func (o *ImportRoadblockEntitlementChanges) HasAbsoluteThreshold() bool`

HasAbsoluteThreshold returns a boolean if a field has been set.

### GetChangedUsers

`func (o *ImportRoadblockEntitlementChanges) GetChangedUsers() int32`

GetChangedUsers returns the ChangedUsers field if non-nil, zero value otherwise.

### GetChangedUsersOk

`func (o *ImportRoadblockEntitlementChanges) GetChangedUsersOk() (*int32, bool)`

GetChangedUsersOk returns a tuple with the ChangedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedUsers

`func (o *ImportRoadblockEntitlementChanges) SetChangedUsers(v int32)`

SetChangedUsers sets ChangedUsers field to given value.

### HasChangedUsers

`func (o *ImportRoadblockEntitlementChanges) HasChangedUsers() bool`

HasChangedUsers returns a boolean if a field has been set.

### GetChangePercentage

`func (o *ImportRoadblockEntitlementChanges) GetChangePercentage() float32`

GetChangePercentage returns the ChangePercentage field if non-nil, zero value otherwise.

### GetChangePercentageOk

`func (o *ImportRoadblockEntitlementChanges) GetChangePercentageOk() (*float32, bool)`

GetChangePercentageOk returns a tuple with the ChangePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangePercentage

`func (o *ImportRoadblockEntitlementChanges) SetChangePercentage(v float32)`

SetChangePercentage sets ChangePercentage field to given value.

### HasChangePercentage

`func (o *ImportRoadblockEntitlementChanges) HasChangePercentage() bool`

HasChangePercentage returns a boolean if a field has been set.

### GetDeletedUsers

`func (o *ImportRoadblockEntitlementChanges) GetDeletedUsers() int32`

GetDeletedUsers returns the DeletedUsers field if non-nil, zero value otherwise.

### GetDeletedUsersOk

`func (o *ImportRoadblockEntitlementChanges) GetDeletedUsersOk() (*int32, bool)`

GetDeletedUsersOk returns a tuple with the DeletedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedUsers

`func (o *ImportRoadblockEntitlementChanges) SetDeletedUsers(v int32)`

SetDeletedUsers sets DeletedUsers field to given value.

### HasDeletedUsers

`func (o *ImportRoadblockEntitlementChanges) HasDeletedUsers() bool`

HasDeletedUsers returns a boolean if a field has been set.

### GetModifiedUsers

`func (o *ImportRoadblockEntitlementChanges) GetModifiedUsers() int32`

GetModifiedUsers returns the ModifiedUsers field if non-nil, zero value otherwise.

### GetModifiedUsersOk

`func (o *ImportRoadblockEntitlementChanges) GetModifiedUsersOk() (*int32, bool)`

GetModifiedUsersOk returns a tuple with the ModifiedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedUsers

`func (o *ImportRoadblockEntitlementChanges) SetModifiedUsers(v int32)`

SetModifiedUsers sets ModifiedUsers field to given value.

### HasModifiedUsers

`func (o *ImportRoadblockEntitlementChanges) HasModifiedUsers() bool`

HasModifiedUsers returns a boolean if a field has been set.

### GetPercentageThreshold

`func (o *ImportRoadblockEntitlementChanges) GetPercentageThreshold() float32`

GetPercentageThreshold returns the PercentageThreshold field if non-nil, zero value otherwise.

### GetPercentageThresholdOk

`func (o *ImportRoadblockEntitlementChanges) GetPercentageThresholdOk() (*float32, bool)`

GetPercentageThresholdOk returns a tuple with the PercentageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageThreshold

`func (o *ImportRoadblockEntitlementChanges) SetPercentageThreshold(v float32)`

SetPercentageThreshold sets PercentageThreshold field to given value.

### HasPercentageThreshold

`func (o *ImportRoadblockEntitlementChanges) HasPercentageThreshold() bool`

HasPercentageThreshold returns a boolean if a field has been set.

### GetThresholdType

`func (o *ImportRoadblockEntitlementChanges) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *ImportRoadblockEntitlementChanges) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *ImportRoadblockEntitlementChanges) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.

### HasThresholdType

`func (o *ImportRoadblockEntitlementChanges) HasThresholdType() bool`

HasThresholdType returns a boolean if a field has been set.

### GetTotalUsers

`func (o *ImportRoadblockEntitlementChanges) GetTotalUsers() int32`

GetTotalUsers returns the TotalUsers field if non-nil, zero value otherwise.

### GetTotalUsersOk

`func (o *ImportRoadblockEntitlementChanges) GetTotalUsersOk() (*int32, bool)`

GetTotalUsersOk returns a tuple with the TotalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsers

`func (o *ImportRoadblockEntitlementChanges) SetTotalUsers(v int32)`

SetTotalUsers sets TotalUsers field to given value.

### HasTotalUsers

`func (o *ImportRoadblockEntitlementChanges) HasTotalUsers() bool`

HasTotalUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


