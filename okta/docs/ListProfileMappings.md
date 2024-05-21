# ListProfileMappings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for profile mapping | [optional] [readonly] 
**Source** | Pointer to [**ProfileMappingSource**](ProfileMappingSource.md) |  | [optional] 
**Target** | Pointer to [**ProfileMappingTarget**](ProfileMappingTarget.md) |  | [optional] 
**Links** | Pointer to [**LinksSelf**](LinksSelf.md) |  | [optional] 

## Methods

### NewListProfileMappings

`func NewListProfileMappings() *ListProfileMappings`

NewListProfileMappings instantiates a new ListProfileMappings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListProfileMappingsWithDefaults

`func NewListProfileMappingsWithDefaults() *ListProfileMappings`

NewListProfileMappingsWithDefaults instantiates a new ListProfileMappings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListProfileMappings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListProfileMappings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListProfileMappings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListProfileMappings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *ListProfileMappings) GetSource() ProfileMappingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListProfileMappings) GetSourceOk() (*ProfileMappingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListProfileMappings) SetSource(v ProfileMappingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ListProfileMappings) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *ListProfileMappings) GetTarget() ProfileMappingTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ListProfileMappings) GetTargetOk() (*ProfileMappingTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ListProfileMappings) SetTarget(v ProfileMappingTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ListProfileMappings) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetLinks

`func (o *ListProfileMappings) GetLinks() LinksSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListProfileMappings) GetLinksOk() (*LinksSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListProfileMappings) SetLinks(v LinksSelf)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListProfileMappings) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


