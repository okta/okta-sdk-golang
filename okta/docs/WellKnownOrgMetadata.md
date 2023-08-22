# WellKnownOrgMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique identifier of the Org | [optional] 
**Pipeline** | Pointer to [**PipelineType**](PipelineType.md) |  | [optional] 
**Settings** | Pointer to [**WellKnownOrgMetadataSettings**](WellKnownOrgMetadataSettings.md) |  | [optional] 
**Links** | Pointer to [**WellKnownOrgMetadataLinks**](WellKnownOrgMetadataLinks.md) |  | [optional] 

## Methods

### NewWellKnownOrgMetadata

`func NewWellKnownOrgMetadata() *WellKnownOrgMetadata`

NewWellKnownOrgMetadata instantiates a new WellKnownOrgMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWellKnownOrgMetadataWithDefaults

`func NewWellKnownOrgMetadataWithDefaults() *WellKnownOrgMetadata`

NewWellKnownOrgMetadataWithDefaults instantiates a new WellKnownOrgMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WellKnownOrgMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WellKnownOrgMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WellKnownOrgMetadata) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WellKnownOrgMetadata) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPipeline

`func (o *WellKnownOrgMetadata) GetPipeline() PipelineType`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *WellKnownOrgMetadata) GetPipelineOk() (*PipelineType, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *WellKnownOrgMetadata) SetPipeline(v PipelineType)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *WellKnownOrgMetadata) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetSettings

`func (o *WellKnownOrgMetadata) GetSettings() WellKnownOrgMetadataSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *WellKnownOrgMetadata) GetSettingsOk() (*WellKnownOrgMetadataSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *WellKnownOrgMetadata) SetSettings(v WellKnownOrgMetadataSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *WellKnownOrgMetadata) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLinks

`func (o *WellKnownOrgMetadata) GetLinks() WellKnownOrgMetadataLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WellKnownOrgMetadata) GetLinksOk() (*WellKnownOrgMetadataLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WellKnownOrgMetadata) SetLinks(v WellKnownOrgMetadataLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WellKnownOrgMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


