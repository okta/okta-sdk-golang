# Feature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to [**FeatureStage**](FeatureStage.md) |  | [optional] 
**Status** | Pointer to [**EnabledStatus**](EnabledStatus.md) |  | [optional] 
**Type** | Pointer to [**FeatureType**](FeatureType.md) |  | [optional] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewFeature

`func NewFeature() *Feature`

NewFeature instantiates a new Feature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureWithDefaults

`func NewFeatureWithDefaults() *Feature`

NewFeatureWithDefaults instantiates a new Feature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Feature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Feature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Feature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Feature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Feature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Feature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Feature) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Feature) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Feature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Feature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Feature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Feature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStage

`func (o *Feature) GetStage() FeatureStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *Feature) GetStageOk() (*FeatureStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *Feature) SetStage(v FeatureStage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *Feature) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetStatus

`func (o *Feature) GetStatus() EnabledStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Feature) GetStatusOk() (*EnabledStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Feature) SetStatus(v EnabledStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Feature) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Feature) GetType() FeatureType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Feature) GetTypeOk() (*FeatureType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Feature) SetType(v FeatureType)`

SetType sets Type field to given value.

### HasType

`func (o *Feature) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *Feature) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Feature) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Feature) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Feature) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


