# SsprRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to [**SsprPrimaryRequirement**](SsprPrimaryRequirement.md) |  | [optional] 
**StepUp** | Pointer to [**SsprStepUpRequirement**](SsprStepUpRequirement.md) |  | [optional] 

## Methods

### NewSsprRequirement

`func NewSsprRequirement() *SsprRequirement`

NewSsprRequirement instantiates a new SsprRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsprRequirementWithDefaults

`func NewSsprRequirementWithDefaults() *SsprRequirement`

NewSsprRequirementWithDefaults instantiates a new SsprRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *SsprRequirement) GetPrimary() SsprPrimaryRequirement`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *SsprRequirement) GetPrimaryOk() (*SsprPrimaryRequirement, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *SsprRequirement) SetPrimary(v SsprPrimaryRequirement)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *SsprRequirement) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetStepUp

`func (o *SsprRequirement) GetStepUp() SsprStepUpRequirement`

GetStepUp returns the StepUp field if non-nil, zero value otherwise.

### GetStepUpOk

`func (o *SsprRequirement) GetStepUpOk() (*SsprStepUpRequirement, bool)`

GetStepUpOk returns a tuple with the StepUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepUp

`func (o *SsprRequirement) SetStepUp(v SsprStepUpRequirement)`

SetStepUp sets StepUp field to given value.

### HasStepUp

`func (o *SsprRequirement) HasStepUp() bool`

HasStepUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


