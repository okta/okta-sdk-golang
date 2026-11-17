# ImportRoadblockEntitlementThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether this threshold is enforced | [optional] 
**OrgDefault** | Pointer to **float32** | The org-level default for this threshold | [optional] 
**Value** | Pointer to **float32** | The threshold set for the app instance. &#x60;-1&#x60; means the app instance uses &#x60;orgDefault&#x60;. | [optional] 

## Methods

### NewImportRoadblockEntitlementThreshold

`func NewImportRoadblockEntitlementThreshold() *ImportRoadblockEntitlementThreshold`

NewImportRoadblockEntitlementThreshold instantiates a new ImportRoadblockEntitlementThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockEntitlementThresholdWithDefaults

`func NewImportRoadblockEntitlementThresholdWithDefaults() *ImportRoadblockEntitlementThreshold`

NewImportRoadblockEntitlementThresholdWithDefaults instantiates a new ImportRoadblockEntitlementThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ImportRoadblockEntitlementThreshold) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ImportRoadblockEntitlementThreshold) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ImportRoadblockEntitlementThreshold) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ImportRoadblockEntitlementThreshold) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOrgDefault

`func (o *ImportRoadblockEntitlementThreshold) GetOrgDefault() float32`

GetOrgDefault returns the OrgDefault field if non-nil, zero value otherwise.

### GetOrgDefaultOk

`func (o *ImportRoadblockEntitlementThreshold) GetOrgDefaultOk() (*float32, bool)`

GetOrgDefaultOk returns a tuple with the OrgDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgDefault

`func (o *ImportRoadblockEntitlementThreshold) SetOrgDefault(v float32)`

SetOrgDefault sets OrgDefault field to given value.

### HasOrgDefault

`func (o *ImportRoadblockEntitlementThreshold) HasOrgDefault() bool`

HasOrgDefault returns a boolean if a field has been set.

### GetValue

`func (o *ImportRoadblockEntitlementThreshold) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ImportRoadblockEntitlementThreshold) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ImportRoadblockEntitlementThreshold) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ImportRoadblockEntitlementThreshold) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


