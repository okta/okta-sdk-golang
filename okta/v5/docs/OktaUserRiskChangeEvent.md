# OktaUserRiskChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentLevel** | **string** | Current risk level of the user | 
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**InitiatingEntity** | Pointer to **string** | The entity that initiated the event | [optional] 
**PreviousLevel** | **string** | Previous risk level of the user | 
**ReasonAdmin** | Pointer to [**CaepDeviceComplianceChangeEventReasonAdmin**](CaepDeviceComplianceChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepDeviceComplianceChangeEventReasonUser**](CaepDeviceComplianceChangeEventReasonUser.md) |  | [optional] 
**Subjects** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewOktaUserRiskChangeEvent

`func NewOktaUserRiskChangeEvent(currentLevel string, eventTimestamp int64, previousLevel string, subjects SecurityEventSubject, ) *OktaUserRiskChangeEvent`

NewOktaUserRiskChangeEvent instantiates a new OktaUserRiskChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaUserRiskChangeEventWithDefaults

`func NewOktaUserRiskChangeEventWithDefaults() *OktaUserRiskChangeEvent`

NewOktaUserRiskChangeEventWithDefaults instantiates a new OktaUserRiskChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentLevel

`func (o *OktaUserRiskChangeEvent) GetCurrentLevel() string`

GetCurrentLevel returns the CurrentLevel field if non-nil, zero value otherwise.

### GetCurrentLevelOk

`func (o *OktaUserRiskChangeEvent) GetCurrentLevelOk() (*string, bool)`

GetCurrentLevelOk returns a tuple with the CurrentLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentLevel

`func (o *OktaUserRiskChangeEvent) SetCurrentLevel(v string)`

SetCurrentLevel sets CurrentLevel field to given value.


### GetEventTimestamp

`func (o *OktaUserRiskChangeEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *OktaUserRiskChangeEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *OktaUserRiskChangeEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetInitiatingEntity

`func (o *OktaUserRiskChangeEvent) GetInitiatingEntity() string`

GetInitiatingEntity returns the InitiatingEntity field if non-nil, zero value otherwise.

### GetInitiatingEntityOk

`func (o *OktaUserRiskChangeEvent) GetInitiatingEntityOk() (*string, bool)`

GetInitiatingEntityOk returns a tuple with the InitiatingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingEntity

`func (o *OktaUserRiskChangeEvent) SetInitiatingEntity(v string)`

SetInitiatingEntity sets InitiatingEntity field to given value.

### HasInitiatingEntity

`func (o *OktaUserRiskChangeEvent) HasInitiatingEntity() bool`

HasInitiatingEntity returns a boolean if a field has been set.

### GetPreviousLevel

`func (o *OktaUserRiskChangeEvent) GetPreviousLevel() string`

GetPreviousLevel returns the PreviousLevel field if non-nil, zero value otherwise.

### GetPreviousLevelOk

`func (o *OktaUserRiskChangeEvent) GetPreviousLevelOk() (*string, bool)`

GetPreviousLevelOk returns a tuple with the PreviousLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLevel

`func (o *OktaUserRiskChangeEvent) SetPreviousLevel(v string)`

SetPreviousLevel sets PreviousLevel field to given value.


### GetReasonAdmin

`func (o *OktaUserRiskChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *OktaUserRiskChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *OktaUserRiskChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *OktaUserRiskChangeEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *OktaUserRiskChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *OktaUserRiskChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *OktaUserRiskChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *OktaUserRiskChangeEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubjects

`func (o *OktaUserRiskChangeEvent) GetSubjects() SecurityEventSubject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *OktaUserRiskChangeEvent) GetSubjectsOk() (*SecurityEventSubject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *OktaUserRiskChangeEvent) SetSubjects(v SecurityEventSubject)`

SetSubjects sets Subjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


