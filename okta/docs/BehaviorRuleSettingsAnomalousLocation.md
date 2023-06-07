# BehaviorRuleSettingsAnomalousLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxEventsUsedForEvaluation** | Pointer to **int32** |  | [optional] [default to 20]
**MinEventsNeededForEvaluation** | Pointer to **int32** |  | [optional] [default to 0]
**Granularity** | [**LocationGranularity**](LocationGranularity.md) |  | 
**RadiusKilometers** | Pointer to **int32** | Required when &#x60;granularity&#x60; is &#x60;LAT_LONG&#x60;. Radius from the provided coordinates in kilometers. | [optional] 

## Methods

### NewBehaviorRuleSettingsAnomalousLocation

`func NewBehaviorRuleSettingsAnomalousLocation(granularity LocationGranularity, ) *BehaviorRuleSettingsAnomalousLocation`

NewBehaviorRuleSettingsAnomalousLocation instantiates a new BehaviorRuleSettingsAnomalousLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRuleSettingsAnomalousLocationWithDefaults

`func NewBehaviorRuleSettingsAnomalousLocationWithDefaults() *BehaviorRuleSettingsAnomalousLocation`

NewBehaviorRuleSettingsAnomalousLocationWithDefaults instantiates a new BehaviorRuleSettingsAnomalousLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxEventsUsedForEvaluation

`func (o *BehaviorRuleSettingsAnomalousLocation) GetMaxEventsUsedForEvaluation() int32`

GetMaxEventsUsedForEvaluation returns the MaxEventsUsedForEvaluation field if non-nil, zero value otherwise.

### GetMaxEventsUsedForEvaluationOk

`func (o *BehaviorRuleSettingsAnomalousLocation) GetMaxEventsUsedForEvaluationOk() (*int32, bool)`

GetMaxEventsUsedForEvaluationOk returns a tuple with the MaxEventsUsedForEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsUsedForEvaluation

`func (o *BehaviorRuleSettingsAnomalousLocation) SetMaxEventsUsedForEvaluation(v int32)`

SetMaxEventsUsedForEvaluation sets MaxEventsUsedForEvaluation field to given value.

### HasMaxEventsUsedForEvaluation

`func (o *BehaviorRuleSettingsAnomalousLocation) HasMaxEventsUsedForEvaluation() bool`

HasMaxEventsUsedForEvaluation returns a boolean if a field has been set.

### GetMinEventsNeededForEvaluation

`func (o *BehaviorRuleSettingsAnomalousLocation) GetMinEventsNeededForEvaluation() int32`

GetMinEventsNeededForEvaluation returns the MinEventsNeededForEvaluation field if non-nil, zero value otherwise.

### GetMinEventsNeededForEvaluationOk

`func (o *BehaviorRuleSettingsAnomalousLocation) GetMinEventsNeededForEvaluationOk() (*int32, bool)`

GetMinEventsNeededForEvaluationOk returns a tuple with the MinEventsNeededForEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEventsNeededForEvaluation

`func (o *BehaviorRuleSettingsAnomalousLocation) SetMinEventsNeededForEvaluation(v int32)`

SetMinEventsNeededForEvaluation sets MinEventsNeededForEvaluation field to given value.

### HasMinEventsNeededForEvaluation

`func (o *BehaviorRuleSettingsAnomalousLocation) HasMinEventsNeededForEvaluation() bool`

HasMinEventsNeededForEvaluation returns a boolean if a field has been set.

### GetGranularity

`func (o *BehaviorRuleSettingsAnomalousLocation) GetGranularity() LocationGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *BehaviorRuleSettingsAnomalousLocation) GetGranularityOk() (*LocationGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *BehaviorRuleSettingsAnomalousLocation) SetGranularity(v LocationGranularity)`

SetGranularity sets Granularity field to given value.


### GetRadiusKilometers

`func (o *BehaviorRuleSettingsAnomalousLocation) GetRadiusKilometers() int32`

GetRadiusKilometers returns the RadiusKilometers field if non-nil, zero value otherwise.

### GetRadiusKilometersOk

`func (o *BehaviorRuleSettingsAnomalousLocation) GetRadiusKilometersOk() (*int32, bool)`

GetRadiusKilometersOk returns a tuple with the RadiusKilometers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusKilometers

`func (o *BehaviorRuleSettingsAnomalousLocation) SetRadiusKilometers(v int32)`

SetRadiusKilometers sets RadiusKilometers field to given value.

### HasRadiusKilometers

`func (o *BehaviorRuleSettingsAnomalousLocation) HasRadiusKilometers() bool`

HasRadiusKilometers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


