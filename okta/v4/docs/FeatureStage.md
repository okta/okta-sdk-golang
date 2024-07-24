# FeatureStage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **string** | Indicates the release state of the feature | [optional] 
**Value** | Pointer to **string** | Current release stage of the feature | [optional] 

## Methods

### NewFeatureStage

`func NewFeatureStage() *FeatureStage`

NewFeatureStage instantiates a new FeatureStage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureStageWithDefaults

`func NewFeatureStageWithDefaults() *FeatureStage`

NewFeatureStageWithDefaults instantiates a new FeatureStage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *FeatureStage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FeatureStage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FeatureStage) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FeatureStage) HasState() bool`

HasState returns a boolean if a field has been set.

### GetValue

`func (o *FeatureStage) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeatureStage) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeatureStage) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FeatureStage) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


