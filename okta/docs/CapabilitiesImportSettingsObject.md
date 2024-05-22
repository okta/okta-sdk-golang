# CapabilitiesImportSettingsObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**ImportScheduleObject**](ImportScheduleObject.md) |  | [optional] 
**Username** | Pointer to [**ImportUsernameObject**](ImportUsernameObject.md) |  | [optional] 

## Methods

### NewCapabilitiesImportSettingsObject

`func NewCapabilitiesImportSettingsObject() *CapabilitiesImportSettingsObject`

NewCapabilitiesImportSettingsObject instantiates a new CapabilitiesImportSettingsObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilitiesImportSettingsObjectWithDefaults

`func NewCapabilitiesImportSettingsObjectWithDefaults() *CapabilitiesImportSettingsObject`

NewCapabilitiesImportSettingsObjectWithDefaults instantiates a new CapabilitiesImportSettingsObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *CapabilitiesImportSettingsObject) GetSchedule() ImportScheduleObject`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CapabilitiesImportSettingsObject) GetScheduleOk() (*ImportScheduleObject, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CapabilitiesImportSettingsObject) SetSchedule(v ImportScheduleObject)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CapabilitiesImportSettingsObject) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetUsername

`func (o *CapabilitiesImportSettingsObject) GetUsername() ImportUsernameObject`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CapabilitiesImportSettingsObject) GetUsernameOk() (*ImportUsernameObject, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CapabilitiesImportSettingsObject) SetUsername(v ImportUsernameObject)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CapabilitiesImportSettingsObject) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


