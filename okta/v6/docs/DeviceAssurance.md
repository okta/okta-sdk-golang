# DeviceAssurance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedDate** | Pointer to **string** |  | [optional] [readonly] 
**DevicePostureChecks** | Pointer to [**DevicePostureChecks**](DevicePostureChecks.md) |  | [optional] 
**DisplayRemediationMode** | Pointer to **string** | &lt;x-lifecycle-container&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/x-lifecycle-container&gt;Represents the remediation mode of this device assurance policy when users are denied access due to device noncompliance | [optional] 
**GracePeriod** | Pointer to [**GracePeriod**](GracePeriod.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdate** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the device assurance policy | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewDeviceAssurance

`func NewDeviceAssurance() *DeviceAssurance`

NewDeviceAssurance instantiates a new DeviceAssurance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAssuranceWithDefaults

`func NewDeviceAssuranceWithDefaults() *DeviceAssurance`

NewDeviceAssuranceWithDefaults instantiates a new DeviceAssurance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DeviceAssurance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DeviceAssurance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DeviceAssurance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DeviceAssurance) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DeviceAssurance) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DeviceAssurance) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DeviceAssurance) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DeviceAssurance) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDevicePostureChecks

`func (o *DeviceAssurance) GetDevicePostureChecks() DevicePostureChecks`

GetDevicePostureChecks returns the DevicePostureChecks field if non-nil, zero value otherwise.

### GetDevicePostureChecksOk

`func (o *DeviceAssurance) GetDevicePostureChecksOk() (*DevicePostureChecks, bool)`

GetDevicePostureChecksOk returns a tuple with the DevicePostureChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePostureChecks

`func (o *DeviceAssurance) SetDevicePostureChecks(v DevicePostureChecks)`

SetDevicePostureChecks sets DevicePostureChecks field to given value.

### HasDevicePostureChecks

`func (o *DeviceAssurance) HasDevicePostureChecks() bool`

HasDevicePostureChecks returns a boolean if a field has been set.

### GetDisplayRemediationMode

`func (o *DeviceAssurance) GetDisplayRemediationMode() string`

GetDisplayRemediationMode returns the DisplayRemediationMode field if non-nil, zero value otherwise.

### GetDisplayRemediationModeOk

`func (o *DeviceAssurance) GetDisplayRemediationModeOk() (*string, bool)`

GetDisplayRemediationModeOk returns a tuple with the DisplayRemediationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRemediationMode

`func (o *DeviceAssurance) SetDisplayRemediationMode(v string)`

SetDisplayRemediationMode sets DisplayRemediationMode field to given value.

### HasDisplayRemediationMode

`func (o *DeviceAssurance) HasDisplayRemediationMode() bool`

HasDisplayRemediationMode returns a boolean if a field has been set.

### GetGracePeriod

`func (o *DeviceAssurance) GetGracePeriod() GracePeriod`

GetGracePeriod returns the GracePeriod field if non-nil, zero value otherwise.

### GetGracePeriodOk

`func (o *DeviceAssurance) GetGracePeriodOk() (*GracePeriod, bool)`

GetGracePeriodOk returns a tuple with the GracePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGracePeriod

`func (o *DeviceAssurance) SetGracePeriod(v GracePeriod)`

SetGracePeriod sets GracePeriod field to given value.

### HasGracePeriod

`func (o *DeviceAssurance) HasGracePeriod() bool`

HasGracePeriod returns a boolean if a field has been set.

### GetId

`func (o *DeviceAssurance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAssurance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAssurance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceAssurance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdate

`func (o *DeviceAssurance) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *DeviceAssurance) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *DeviceAssurance) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *DeviceAssurance) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *DeviceAssurance) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *DeviceAssurance) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *DeviceAssurance) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *DeviceAssurance) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetName

`func (o *DeviceAssurance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceAssurance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceAssurance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceAssurance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceAssurance) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceAssurance) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceAssurance) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceAssurance) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetLinks

`func (o *DeviceAssurance) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceAssurance) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceAssurance) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceAssurance) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


