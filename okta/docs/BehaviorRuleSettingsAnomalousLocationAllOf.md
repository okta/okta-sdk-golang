# BehaviorRuleSettingsAnomalousLocationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Granularity** | [**LocationGranularity**](LocationGranularity.md) |  | 
**RadiusKilometers** | Pointer to **int32** | Required when &#x60;granularity&#x60; is &#x60;LAT_LONG&#x60;. Radius from the provided coordinates in kilometers. | [optional] 

## Methods

### NewBehaviorRuleSettingsAnomalousLocationAllOf

`func NewBehaviorRuleSettingsAnomalousLocationAllOf(granularity LocationGranularity, ) *BehaviorRuleSettingsAnomalousLocationAllOf`

NewBehaviorRuleSettingsAnomalousLocationAllOf instantiates a new BehaviorRuleSettingsAnomalousLocationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBehaviorRuleSettingsAnomalousLocationAllOfWithDefaults

`func NewBehaviorRuleSettingsAnomalousLocationAllOfWithDefaults() *BehaviorRuleSettingsAnomalousLocationAllOf`

NewBehaviorRuleSettingsAnomalousLocationAllOfWithDefaults instantiates a new BehaviorRuleSettingsAnomalousLocationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularity

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetGranularity() LocationGranularity`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetGranularityOk() (*LocationGranularity, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) SetGranularity(v LocationGranularity)`

SetGranularity sets Granularity field to given value.


### GetRadiusKilometers

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetRadiusKilometers() int32`

GetRadiusKilometers returns the RadiusKilometers field if non-nil, zero value otherwise.

### GetRadiusKilometersOk

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) GetRadiusKilometersOk() (*int32, bool)`

GetRadiusKilometersOk returns a tuple with the RadiusKilometers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusKilometers

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) SetRadiusKilometers(v int32)`

SetRadiusKilometers sets RadiusKilometers field to given value.

### HasRadiusKilometers

`func (o *BehaviorRuleSettingsAnomalousLocationAllOf) HasRadiusKilometers() bool`

HasRadiusKilometers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


