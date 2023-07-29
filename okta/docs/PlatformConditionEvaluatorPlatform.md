# PlatformConditionEvaluatorPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | Pointer to [**PlatformConditionEvaluatorPlatformOperatingSystem**](PlatformConditionEvaluatorPlatformOperatingSystem.md) |  | [optional] 
**Type** | Pointer to [**PolicyPlatformType**](PolicyPlatformType.md) |  | [optional] 

## Methods

### NewPlatformConditionEvaluatorPlatform

`func NewPlatformConditionEvaluatorPlatform() *PlatformConditionEvaluatorPlatform`

NewPlatformConditionEvaluatorPlatform instantiates a new PlatformConditionEvaluatorPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformConditionEvaluatorPlatformWithDefaults

`func NewPlatformConditionEvaluatorPlatformWithDefaults() *PlatformConditionEvaluatorPlatform`

NewPlatformConditionEvaluatorPlatformWithDefaults instantiates a new PlatformConditionEvaluatorPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *PlatformConditionEvaluatorPlatform) GetOs() PlatformConditionEvaluatorPlatformOperatingSystem`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *PlatformConditionEvaluatorPlatform) GetOsOk() (*PlatformConditionEvaluatorPlatformOperatingSystem, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *PlatformConditionEvaluatorPlatform) SetOs(v PlatformConditionEvaluatorPlatformOperatingSystem)`

SetOs sets Os field to given value.

### HasOs

`func (o *PlatformConditionEvaluatorPlatform) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetType

`func (o *PlatformConditionEvaluatorPlatform) GetType() PolicyPlatformType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlatformConditionEvaluatorPlatform) GetTypeOk() (*PolicyPlatformType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlatformConditionEvaluatorPlatform) SetType(v PolicyPlatformType)`

SetType sets Type field to given value.

### HasType

`func (o *PlatformConditionEvaluatorPlatform) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


