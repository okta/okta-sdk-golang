# PerClientRateLimitSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMode** | [**PerClientRateLimitMode**](PerClientRateLimitMode.md) |  | 
**UseCaseModeOverrides** | Pointer to [**PerClientRateLimitSettingsUseCaseModeOverrides**](PerClientRateLimitSettingsUseCaseModeOverrides.md) |  | [optional] 

## Methods

### NewPerClientRateLimitSettings

`func NewPerClientRateLimitSettings(defaultMode PerClientRateLimitMode, ) *PerClientRateLimitSettings`

NewPerClientRateLimitSettings instantiates a new PerClientRateLimitSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerClientRateLimitSettingsWithDefaults

`func NewPerClientRateLimitSettingsWithDefaults() *PerClientRateLimitSettings`

NewPerClientRateLimitSettingsWithDefaults instantiates a new PerClientRateLimitSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMode

`func (o *PerClientRateLimitSettings) GetDefaultMode() PerClientRateLimitMode`

GetDefaultMode returns the DefaultMode field if non-nil, zero value otherwise.

### GetDefaultModeOk

`func (o *PerClientRateLimitSettings) GetDefaultModeOk() (*PerClientRateLimitMode, bool)`

GetDefaultModeOk returns a tuple with the DefaultMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMode

`func (o *PerClientRateLimitSettings) SetDefaultMode(v PerClientRateLimitMode)`

SetDefaultMode sets DefaultMode field to given value.


### GetUseCaseModeOverrides

`func (o *PerClientRateLimitSettings) GetUseCaseModeOverrides() PerClientRateLimitSettingsUseCaseModeOverrides`

GetUseCaseModeOverrides returns the UseCaseModeOverrides field if non-nil, zero value otherwise.

### GetUseCaseModeOverridesOk

`func (o *PerClientRateLimitSettings) GetUseCaseModeOverridesOk() (*PerClientRateLimitSettingsUseCaseModeOverrides, bool)`

GetUseCaseModeOverridesOk returns a tuple with the UseCaseModeOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCaseModeOverrides

`func (o *PerClientRateLimitSettings) SetUseCaseModeOverrides(v PerClientRateLimitSettingsUseCaseModeOverrides)`

SetUseCaseModeOverrides sets UseCaseModeOverrides field to given value.

### HasUseCaseModeOverrides

`func (o *PerClientRateLimitSettings) HasUseCaseModeOverrides() bool`

HasUseCaseModeOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


