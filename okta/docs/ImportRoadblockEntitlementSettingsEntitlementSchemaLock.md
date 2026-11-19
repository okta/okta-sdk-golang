# ImportRoadblockEntitlementSettingsEntitlementSchemaLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locked** | Pointer to **bool** | Indicates whether the entitlement schema is locked | [optional] 
**LockedAt** | Pointer to **time.Time** | Timestamp when the entitlement schema was locked | [optional] 
**LockedBy** | Pointer to **string** | ID of the user who locked the entitlement schema | [optional] 

## Methods

### NewImportRoadblockEntitlementSettingsEntitlementSchemaLock

`func NewImportRoadblockEntitlementSettingsEntitlementSchemaLock() *ImportRoadblockEntitlementSettingsEntitlementSchemaLock`

NewImportRoadblockEntitlementSettingsEntitlementSchemaLock instantiates a new ImportRoadblockEntitlementSettingsEntitlementSchemaLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockEntitlementSettingsEntitlementSchemaLockWithDefaults

`func NewImportRoadblockEntitlementSettingsEntitlementSchemaLockWithDefaults() *ImportRoadblockEntitlementSettingsEntitlementSchemaLock`

NewImportRoadblockEntitlementSettingsEntitlementSchemaLockWithDefaults instantiates a new ImportRoadblockEntitlementSettingsEntitlementSchemaLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocked

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedAt

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedAt() time.Time`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedAtOk() (*time.Time, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) SetLockedAt(v time.Time)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetLockedBy

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *ImportRoadblockEntitlementSettingsEntitlementSchemaLock) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


