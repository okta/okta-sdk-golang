# DevicePostureCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | User who created the device posture check | [optional] [readonly] 
**CreatedDate** | Pointer to **string** | Time the device posture check was created | [optional] [readonly] 
**Description** | Pointer to **string** | Description of the device posture check | [optional] 
**Id** | Pointer to **string** | The ID of the device posture check | [optional] [readonly] 
**LastUpdate** | Pointer to **string** | Time the device posture check was updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | User who updated the device posture check | [optional] [readonly] 
**MappingType** | Pointer to **string** | Represents how the device posture check is rendered in device assurance policies | [optional] 
**Name** | Pointer to **string** | Display name of the device posture check | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** | OSQuery for the device posture check | [optional] 
**RemediationSettings** | Pointer to [**DevicePostureChecksRemediationSettings**](DevicePostureChecksRemediationSettings.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**VariableName** | Pointer to **string** | Unique name of the device posture check | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewDevicePostureCheck

`func NewDevicePostureCheck() *DevicePostureCheck`

NewDevicePostureCheck instantiates a new DevicePostureCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicePostureCheckWithDefaults

`func NewDevicePostureCheckWithDefaults() *DevicePostureCheck`

NewDevicePostureCheckWithDefaults instantiates a new DevicePostureCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *DevicePostureCheck) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DevicePostureCheck) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DevicePostureCheck) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DevicePostureCheck) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DevicePostureCheck) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DevicePostureCheck) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DevicePostureCheck) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DevicePostureCheck) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDescription

`func (o *DevicePostureCheck) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DevicePostureCheck) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DevicePostureCheck) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DevicePostureCheck) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DevicePostureCheck) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevicePostureCheck) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevicePostureCheck) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DevicePostureCheck) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdate

`func (o *DevicePostureCheck) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *DevicePostureCheck) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *DevicePostureCheck) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *DevicePostureCheck) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *DevicePostureCheck) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *DevicePostureCheck) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *DevicePostureCheck) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *DevicePostureCheck) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetMappingType

`func (o *DevicePostureCheck) GetMappingType() string`

GetMappingType returns the MappingType field if non-nil, zero value otherwise.

### GetMappingTypeOk

`func (o *DevicePostureCheck) GetMappingTypeOk() (*string, bool)`

GetMappingTypeOk returns a tuple with the MappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingType

`func (o *DevicePostureCheck) SetMappingType(v string)`

SetMappingType sets MappingType field to given value.

### HasMappingType

`func (o *DevicePostureCheck) HasMappingType() bool`

HasMappingType returns a boolean if a field has been set.

### GetName

`func (o *DevicePostureCheck) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevicePostureCheck) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevicePostureCheck) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevicePostureCheck) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *DevicePostureCheck) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DevicePostureCheck) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DevicePostureCheck) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DevicePostureCheck) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetQuery

`func (o *DevicePostureCheck) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *DevicePostureCheck) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *DevicePostureCheck) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *DevicePostureCheck) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetRemediationSettings

`func (o *DevicePostureCheck) GetRemediationSettings() DevicePostureChecksRemediationSettings`

GetRemediationSettings returns the RemediationSettings field if non-nil, zero value otherwise.

### GetRemediationSettingsOk

`func (o *DevicePostureCheck) GetRemediationSettingsOk() (*DevicePostureChecksRemediationSettings, bool)`

GetRemediationSettingsOk returns a tuple with the RemediationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationSettings

`func (o *DevicePostureCheck) SetRemediationSettings(v DevicePostureChecksRemediationSettings)`

SetRemediationSettings sets RemediationSettings field to given value.

### HasRemediationSettings

`func (o *DevicePostureCheck) HasRemediationSettings() bool`

HasRemediationSettings returns a boolean if a field has been set.

### GetType

`func (o *DevicePostureCheck) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DevicePostureCheck) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DevicePostureCheck) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DevicePostureCheck) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVariableName

`func (o *DevicePostureCheck) GetVariableName() string`

GetVariableName returns the VariableName field if non-nil, zero value otherwise.

### GetVariableNameOk

`func (o *DevicePostureCheck) GetVariableNameOk() (*string, bool)`

GetVariableNameOk returns a tuple with the VariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableName

`func (o *DevicePostureCheck) SetVariableName(v string)`

SetVariableName sets VariableName field to given value.

### HasVariableName

`func (o *DevicePostureCheck) HasVariableName() bool`

HasVariableName returns a boolean if a field has been set.

### GetLinks

`func (o *DevicePostureCheck) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevicePostureCheck) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevicePostureCheck) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevicePostureCheck) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


