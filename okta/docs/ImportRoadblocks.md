# ImportRoadblocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportJobSuspended** | Pointer to **bool** | Indicates whether any user import in the org is currently paused by a roadblock. When the request filters by &#x60;appInstanceId&#x60;, indicates whether that app instance&#39;s user imports are paused. | [optional] [readonly] 
**Roadblocks** | Pointer to [**[]ImportRoadblock**](ImportRoadblock.md) | The import roadblocks that match the request | [optional] [readonly] 

## Methods

### NewImportRoadblocks

`func NewImportRoadblocks() *ImportRoadblocks`

NewImportRoadblocks instantiates a new ImportRoadblocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblocksWithDefaults

`func NewImportRoadblocksWithDefaults() *ImportRoadblocks`

NewImportRoadblocksWithDefaults instantiates a new ImportRoadblocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportJobSuspended

`func (o *ImportRoadblocks) GetImportJobSuspended() bool`

GetImportJobSuspended returns the ImportJobSuspended field if non-nil, zero value otherwise.

### GetImportJobSuspendedOk

`func (o *ImportRoadblocks) GetImportJobSuspendedOk() (*bool, bool)`

GetImportJobSuspendedOk returns a tuple with the ImportJobSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportJobSuspended

`func (o *ImportRoadblocks) SetImportJobSuspended(v bool)`

SetImportJobSuspended sets ImportJobSuspended field to given value.

### HasImportJobSuspended

`func (o *ImportRoadblocks) HasImportJobSuspended() bool`

HasImportJobSuspended returns a boolean if a field has been set.

### GetRoadblocks

`func (o *ImportRoadblocks) GetRoadblocks() []ImportRoadblock`

GetRoadblocks returns the Roadblocks field if non-nil, zero value otherwise.

### GetRoadblocksOk

`func (o *ImportRoadblocks) GetRoadblocksOk() (*[]ImportRoadblock, bool)`

GetRoadblocksOk returns a tuple with the Roadblocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoadblocks

`func (o *ImportRoadblocks) SetRoadblocks(v []ImportRoadblock)`

SetRoadblocks sets Roadblocks field to given value.

### HasRoadblocks

`func (o *ImportRoadblocks) HasRoadblocks() bool`

HasRoadblocks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


