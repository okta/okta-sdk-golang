# AutoUpdateSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cron** | Pointer to **string** | The schedule of the update in cron format. The cron settings are limited to only the day of the month or the nth-day-of-the-week configurations. For example, &#x60;0 8 ? * 6#3&#x60; indicates every third Saturday at 8:00 AM. | [optional] 
**Delay** | Pointer to **int32** | Delay in days | [optional] 
**Duration** | Pointer to **int32** | Duration in minutes | [optional] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the update finished (only for a successful or failed update, not for a cancelled update). Null is returned if the job hasn&#39;t finished once yet. | [optional] 
**Timezone** | Pointer to **string** | Timezone of where the scheduled job takes place | [optional] 

## Methods

### NewAutoUpdateSchedule

`func NewAutoUpdateSchedule() *AutoUpdateSchedule`

NewAutoUpdateSchedule instantiates a new AutoUpdateSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoUpdateScheduleWithDefaults

`func NewAutoUpdateScheduleWithDefaults() *AutoUpdateSchedule`

NewAutoUpdateScheduleWithDefaults instantiates a new AutoUpdateSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCron

`func (o *AutoUpdateSchedule) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *AutoUpdateSchedule) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *AutoUpdateSchedule) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *AutoUpdateSchedule) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetDelay

`func (o *AutoUpdateSchedule) GetDelay() int32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *AutoUpdateSchedule) GetDelayOk() (*int32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *AutoUpdateSchedule) SetDelay(v int32)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *AutoUpdateSchedule) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetDuration

`func (o *AutoUpdateSchedule) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AutoUpdateSchedule) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AutoUpdateSchedule) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AutoUpdateSchedule) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AutoUpdateSchedule) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AutoUpdateSchedule) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AutoUpdateSchedule) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AutoUpdateSchedule) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetTimezone

`func (o *AutoUpdateSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *AutoUpdateSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *AutoUpdateSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *AutoUpdateSchedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


