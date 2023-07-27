# ProfileMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Properties** | Pointer to [**map[string]ProfileMappingProperty**](ProfileMappingProperty.md) |  | [optional] [readonly] 
**Source** | Pointer to [**ProfileMappingSource**](ProfileMappingSource.md) |  | [optional] 
**Target** | Pointer to [**ProfileMappingSource**](ProfileMappingSource.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewProfileMapping

`func NewProfileMapping() *ProfileMapping`

NewProfileMapping instantiates a new ProfileMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileMappingWithDefaults

`func NewProfileMappingWithDefaults() *ProfileMapping`

NewProfileMappingWithDefaults instantiates a new ProfileMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProfileMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProfileMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProfileMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProfileMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *ProfileMapping) GetProperties() map[string]ProfileMappingProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProfileMapping) GetPropertiesOk() (*map[string]ProfileMappingProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProfileMapping) SetProperties(v map[string]ProfileMappingProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProfileMapping) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSource

`func (o *ProfileMapping) GetSource() ProfileMappingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ProfileMapping) GetSourceOk() (*ProfileMappingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ProfileMapping) SetSource(v ProfileMappingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ProfileMapping) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTarget

`func (o *ProfileMapping) GetTarget() ProfileMappingSource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ProfileMapping) GetTargetOk() (*ProfileMappingSource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ProfileMapping) SetTarget(v ProfileMappingSource)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ProfileMapping) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetLinks

`func (o *ProfileMapping) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProfileMapping) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProfileMapping) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProfileMapping) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


