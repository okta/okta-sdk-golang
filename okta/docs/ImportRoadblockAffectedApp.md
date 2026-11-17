# ImportRoadblockAffectedApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppInstanceId** | Pointer to **string** | ID of the app instance that would lose assignments | [optional] [readonly] 
**AppInstanceLabel** | Pointer to **string** | Label of the app instance that would lose assignments | [optional] [readonly] 
**Assignments** | Pointer to **int32** | Number of assignments to this app used as the denominator for &#x60;unassignmentPercentage&#x60; | [optional] [readonly] 
**Threshold** | Pointer to **int32** | The app-level percentage threshold that this app crossed | [optional] [readonly] 
**UnassignmentPercentage** | Pointer to **float32** | The share of assignments to this app that the import would remove (&#x60;unassignments&#x60; / &#x60;assignments&#x60; × 100) | [optional] [readonly] 
**Unassignments** | Pointer to **int32** | Number of assignments to this app that the import would remove | [optional] [readonly] 

## Methods

### NewImportRoadblockAffectedApp

`func NewImportRoadblockAffectedApp() *ImportRoadblockAffectedApp`

NewImportRoadblockAffectedApp instantiates a new ImportRoadblockAffectedApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportRoadblockAffectedAppWithDefaults

`func NewImportRoadblockAffectedAppWithDefaults() *ImportRoadblockAffectedApp`

NewImportRoadblockAffectedAppWithDefaults instantiates a new ImportRoadblockAffectedApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppInstanceId

`func (o *ImportRoadblockAffectedApp) GetAppInstanceId() string`

GetAppInstanceId returns the AppInstanceId field if non-nil, zero value otherwise.

### GetAppInstanceIdOk

`func (o *ImportRoadblockAffectedApp) GetAppInstanceIdOk() (*string, bool)`

GetAppInstanceIdOk returns a tuple with the AppInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceId

`func (o *ImportRoadblockAffectedApp) SetAppInstanceId(v string)`

SetAppInstanceId sets AppInstanceId field to given value.

### HasAppInstanceId

`func (o *ImportRoadblockAffectedApp) HasAppInstanceId() bool`

HasAppInstanceId returns a boolean if a field has been set.

### GetAppInstanceLabel

`func (o *ImportRoadblockAffectedApp) GetAppInstanceLabel() string`

GetAppInstanceLabel returns the AppInstanceLabel field if non-nil, zero value otherwise.

### GetAppInstanceLabelOk

`func (o *ImportRoadblockAffectedApp) GetAppInstanceLabelOk() (*string, bool)`

GetAppInstanceLabelOk returns a tuple with the AppInstanceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppInstanceLabel

`func (o *ImportRoadblockAffectedApp) SetAppInstanceLabel(v string)`

SetAppInstanceLabel sets AppInstanceLabel field to given value.

### HasAppInstanceLabel

`func (o *ImportRoadblockAffectedApp) HasAppInstanceLabel() bool`

HasAppInstanceLabel returns a boolean if a field has been set.

### GetAssignments

`func (o *ImportRoadblockAffectedApp) GetAssignments() int32`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *ImportRoadblockAffectedApp) GetAssignmentsOk() (*int32, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *ImportRoadblockAffectedApp) SetAssignments(v int32)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *ImportRoadblockAffectedApp) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetThreshold

`func (o *ImportRoadblockAffectedApp) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *ImportRoadblockAffectedApp) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *ImportRoadblockAffectedApp) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *ImportRoadblockAffectedApp) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetUnassignmentPercentage

`func (o *ImportRoadblockAffectedApp) GetUnassignmentPercentage() float32`

GetUnassignmentPercentage returns the UnassignmentPercentage field if non-nil, zero value otherwise.

### GetUnassignmentPercentageOk

`func (o *ImportRoadblockAffectedApp) GetUnassignmentPercentageOk() (*float32, bool)`

GetUnassignmentPercentageOk returns a tuple with the UnassignmentPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignmentPercentage

`func (o *ImportRoadblockAffectedApp) SetUnassignmentPercentage(v float32)`

SetUnassignmentPercentage sets UnassignmentPercentage field to given value.

### HasUnassignmentPercentage

`func (o *ImportRoadblockAffectedApp) HasUnassignmentPercentage() bool`

HasUnassignmentPercentage returns a boolean if a field has been set.

### GetUnassignments

`func (o *ImportRoadblockAffectedApp) GetUnassignments() int32`

GetUnassignments returns the Unassignments field if non-nil, zero value otherwise.

### GetUnassignmentsOk

`func (o *ImportRoadblockAffectedApp) GetUnassignmentsOk() (*int32, bool)`

GetUnassignmentsOk returns a tuple with the Unassignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnassignments

`func (o *ImportRoadblockAffectedApp) SetUnassignments(v int32)`

SetUnassignments sets Unassignments field to given value.

### HasUnassignments

`func (o *ImportRoadblockAffectedApp) HasUnassignments() bool`

HasUnassignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


