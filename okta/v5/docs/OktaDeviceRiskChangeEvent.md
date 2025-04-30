# OktaDeviceRiskChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentLevel** | **string** | Current risk level of the device | 
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**InitiatingEntity** | Pointer to **string** | The entity that initiated the event | [optional] 
**PreviousLevel** | **string** | Previous risk level of the device | 
**ReasonAdmin** | Pointer to [**CaepDeviceComplianceChangeEventReasonAdmin**](CaepDeviceComplianceChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepDeviceComplianceChangeEventReasonUser**](CaepDeviceComplianceChangeEventReasonUser.md) |  | [optional] 
**Subjects** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewOktaDeviceRiskChangeEvent

`func NewOktaDeviceRiskChangeEvent(currentLevel string, eventTimestamp int64, previousLevel string, subjects SecurityEventSubject, ) *OktaDeviceRiskChangeEvent`

NewOktaDeviceRiskChangeEvent instantiates a new OktaDeviceRiskChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaDeviceRiskChangeEventWithDefaults

`func NewOktaDeviceRiskChangeEventWithDefaults() *OktaDeviceRiskChangeEvent`

NewOktaDeviceRiskChangeEventWithDefaults instantiates a new OktaDeviceRiskChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentLevel

`func (o *OktaDeviceRiskChangeEvent) GetCurrentLevel() string`

GetCurrentLevel returns the CurrentLevel field if non-nil, zero value otherwise.

### GetCurrentLevelOk

`func (o *OktaDeviceRiskChangeEvent) GetCurrentLevelOk() (*string, bool)`

GetCurrentLevelOk returns a tuple with the CurrentLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLevel

`func (o *OktaDeviceRiskChangeEvent) SetCurrentLevel(v string)`

SetCurrentLevel sets CurrentLevel field to given value.


### GetEventTimestamp

`func (o *OktaDeviceRiskChangeEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *OktaDeviceRiskChangeEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *OktaDeviceRiskChangeEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetInitiatingEntity

`func (o *OktaDeviceRiskChangeEvent) GetInitiatingEntity() string`

GetInitiatingEntity returns the InitiatingEntity field if non-nil, zero value otherwise.

### GetInitiatingEntityOk

`func (o *OktaDeviceRiskChangeEvent) GetInitiatingEntityOk() (*string, bool)`

GetInitiatingEntityOk returns a tuple with the InitiatingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingEntity

`func (o *OktaDeviceRiskChangeEvent) SetInitiatingEntity(v string)`

SetInitiatingEntity sets InitiatingEntity field to given value.

### HasInitiatingEntity

`func (o *OktaDeviceRiskChangeEvent) HasInitiatingEntity() bool`

HasInitiatingEntity returns a boolean if a field has been set.

### GetPreviousLevel

`func (o *OktaDeviceRiskChangeEvent) GetPreviousLevel() string`

GetPreviousLevel returns the PreviousLevel field if non-nil, zero value otherwise.

### GetPreviousLevelOk

`func (o *OktaDeviceRiskChangeEvent) GetPreviousLevelOk() (*string, bool)`

GetPreviousLevelOk returns a tuple with the PreviousLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLevel

`func (o *OktaDeviceRiskChangeEvent) SetPreviousLevel(v string)`

SetPreviousLevel sets PreviousLevel field to given value.


### GetReasonAdmin

`func (o *OktaDeviceRiskChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *OktaDeviceRiskChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *OktaDeviceRiskChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *OktaDeviceRiskChangeEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *OktaDeviceRiskChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *OktaDeviceRiskChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *OktaDeviceRiskChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *OktaDeviceRiskChangeEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubjects

`func (o *OktaDeviceRiskChangeEvent) GetSubjects() SecurityEventSubject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *OktaDeviceRiskChangeEvent) GetSubjectsOk() (*SecurityEventSubject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *OktaDeviceRiskChangeEvent) SetSubjects(v SecurityEventSubject)`

SetSubjects sets Subjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


