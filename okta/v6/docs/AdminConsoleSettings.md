# AdminConsoleSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionIdleTimeoutMinutes** | Pointer to **int32** | The maximum idle time before the Okta Admin Console session expires. Must be no more than 12 hours. | [optional] [default to 15]
**SessionMaxLifetimeMinutes** | Pointer to **int32** | The absolute maximum session lifetime of the Okta Admin Console. Must be no more than 7 days. | [optional] [default to 720]

## Methods

### NewAdminConsoleSettings

`func NewAdminConsoleSettings() *AdminConsoleSettings`

NewAdminConsoleSettings instantiates a new AdminConsoleSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminConsoleSettingsWithDefaults

`func NewAdminConsoleSettingsWithDefaults() *AdminConsoleSettings`

NewAdminConsoleSettingsWithDefaults instantiates a new AdminConsoleSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionIdleTimeoutMinutes

`func (o *AdminConsoleSettings) GetSessionIdleTimeoutMinutes() int32`

GetSessionIdleTimeoutMinutes returns the SessionIdleTimeoutMinutes field if non-nil, zero value otherwise.

### GetSessionIdleTimeoutMinutesOk

`func (o *AdminConsoleSettings) GetSessionIdleTimeoutMinutesOk() (*int32, bool)`

GetSessionIdleTimeoutMinutesOk returns a tuple with the SessionIdleTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIdleTimeoutMinutes

`func (o *AdminConsoleSettings) SetSessionIdleTimeoutMinutes(v int32)`

SetSessionIdleTimeoutMinutes sets SessionIdleTimeoutMinutes field to given value.

### HasSessionIdleTimeoutMinutes

`func (o *AdminConsoleSettings) HasSessionIdleTimeoutMinutes() bool`

HasSessionIdleTimeoutMinutes returns a boolean if a field has been set.

### GetSessionMaxLifetimeMinutes

`func (o *AdminConsoleSettings) GetSessionMaxLifetimeMinutes() int32`

GetSessionMaxLifetimeMinutes returns the SessionMaxLifetimeMinutes field if non-nil, zero value otherwise.

### GetSessionMaxLifetimeMinutesOk

`func (o *AdminConsoleSettings) GetSessionMaxLifetimeMinutesOk() (*int32, bool)`

GetSessionMaxLifetimeMinutesOk returns a tuple with the SessionMaxLifetimeMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionMaxLifetimeMinutes

`func (o *AdminConsoleSettings) SetSessionMaxLifetimeMinutes(v int32)`

SetSessionMaxLifetimeMinutes sets SessionMaxLifetimeMinutes field to given value.

### HasSessionMaxLifetimeMinutes

`func (o *AdminConsoleSettings) HasSessionMaxLifetimeMinutes() bool`

HasSessionMaxLifetimeMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


