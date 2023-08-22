# ProfileSettingObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**EnabledStatus**](EnabledStatus.md) |  | [optional] 

## Methods

### NewProfileSettingObject

`func NewProfileSettingObject() *ProfileSettingObject`

NewProfileSettingObject instantiates a new ProfileSettingObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileSettingObjectWithDefaults

`func NewProfileSettingObjectWithDefaults() *ProfileSettingObject`

NewProfileSettingObjectWithDefaults instantiates a new ProfileSettingObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProfileSettingObject) GetStatus() EnabledStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProfileSettingObject) GetStatusOk() (*EnabledStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProfileSettingObject) SetStatus(v EnabledStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProfileSettingObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


