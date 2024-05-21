# ImportScheduleObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullImport** | Pointer to [**ImportScheduleObjectFullImport**](ImportScheduleObjectFullImport.md) |  | [optional] 
**IncrementalImport** | Pointer to [**ImportScheduleObjectIncrementalImport**](ImportScheduleObjectIncrementalImport.md) |  | [optional] 
**Status** | Pointer to **string** | Setting status | [optional] 

## Methods

### NewImportScheduleObject

`func NewImportScheduleObject() *ImportScheduleObject`

NewImportScheduleObject instantiates a new ImportScheduleObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportScheduleObjectWithDefaults

`func NewImportScheduleObjectWithDefaults() *ImportScheduleObject`

NewImportScheduleObjectWithDefaults instantiates a new ImportScheduleObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullImport

`func (o *ImportScheduleObject) GetFullImport() ImportScheduleObjectFullImport`

GetFullImport returns the FullImport field if non-nil, zero value otherwise.

### GetFullImportOk

`func (o *ImportScheduleObject) GetFullImportOk() (*ImportScheduleObjectFullImport, bool)`

GetFullImportOk returns a tuple with the FullImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullImport

`func (o *ImportScheduleObject) SetFullImport(v ImportScheduleObjectFullImport)`

SetFullImport sets FullImport field to given value.

### HasFullImport

`func (o *ImportScheduleObject) HasFullImport() bool`

HasFullImport returns a boolean if a field has been set.

### GetIncrementalImport

`func (o *ImportScheduleObject) GetIncrementalImport() ImportScheduleObjectIncrementalImport`

GetIncrementalImport returns the IncrementalImport field if non-nil, zero value otherwise.

### GetIncrementalImportOk

`func (o *ImportScheduleObject) GetIncrementalImportOk() (*ImportScheduleObjectIncrementalImport, bool)`

GetIncrementalImportOk returns a tuple with the IncrementalImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalImport

`func (o *ImportScheduleObject) SetIncrementalImport(v ImportScheduleObjectIncrementalImport)`

SetIncrementalImport sets IncrementalImport field to given value.

### HasIncrementalImport

`func (o *ImportScheduleObject) HasIncrementalImport() bool`

HasIncrementalImport returns a boolean if a field has been set.

### GetStatus

`func (o *ImportScheduleObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportScheduleObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportScheduleObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ImportScheduleObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


