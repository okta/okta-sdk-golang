# CaepDeviceComplianceChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentStatus** | **string** | Current device compliance status | 
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**InitiatingEntity** | Pointer to **string** | The entity that initiated the event | [optional] 
**PreviousStatus** | **string** | Previous device compliance status | 
**ReasonAdmin** | Pointer to [**CaepDeviceComplianceChangeEventReasonAdmin**](CaepDeviceComplianceChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepDeviceComplianceChangeEventReasonUser**](CaepDeviceComplianceChangeEventReasonUser.md) |  | [optional] 
**Subjects** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewCaepDeviceComplianceChangeEvent

`func NewCaepDeviceComplianceChangeEvent(currentStatus string, eventTimestamp int64, previousStatus string, subjects SecurityEventSubject, ) *CaepDeviceComplianceChangeEvent`

NewCaepDeviceComplianceChangeEvent instantiates a new CaepDeviceComplianceChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaepDeviceComplianceChangeEventWithDefaults

`func NewCaepDeviceComplianceChangeEventWithDefaults() *CaepDeviceComplianceChangeEvent`

NewCaepDeviceComplianceChangeEventWithDefaults instantiates a new CaepDeviceComplianceChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentStatus

`func (o *CaepDeviceComplianceChangeEvent) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *CaepDeviceComplianceChangeEvent) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *CaepDeviceComplianceChangeEvent) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.


### GetEventTimestamp

`func (o *CaepDeviceComplianceChangeEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CaepDeviceComplianceChangeEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CaepDeviceComplianceChangeEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetInitiatingEntity

`func (o *CaepDeviceComplianceChangeEvent) GetInitiatingEntity() string`

GetInitiatingEntity returns the InitiatingEntity field if non-nil, zero value otherwise.

### GetInitiatingEntityOk

`func (o *CaepDeviceComplianceChangeEvent) GetInitiatingEntityOk() (*string, bool)`

GetInitiatingEntityOk returns a tuple with the InitiatingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingEntity

`func (o *CaepDeviceComplianceChangeEvent) SetInitiatingEntity(v string)`

SetInitiatingEntity sets InitiatingEntity field to given value.

### HasInitiatingEntity

`func (o *CaepDeviceComplianceChangeEvent) HasInitiatingEntity() bool`

HasInitiatingEntity returns a boolean if a field has been set.

### GetPreviousStatus

`func (o *CaepDeviceComplianceChangeEvent) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *CaepDeviceComplianceChangeEvent) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *CaepDeviceComplianceChangeEvent) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.


### GetReasonAdmin

`func (o *CaepDeviceComplianceChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *CaepDeviceComplianceChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *CaepDeviceComplianceChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *CaepDeviceComplianceChangeEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *CaepDeviceComplianceChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *CaepDeviceComplianceChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *CaepDeviceComplianceChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *CaepDeviceComplianceChangeEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubjects

`func (o *CaepDeviceComplianceChangeEvent) GetSubjects() SecurityEventSubject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CaepDeviceComplianceChangeEvent) GetSubjectsOk() (*SecurityEventSubject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CaepDeviceComplianceChangeEvent) SetSubjects(v SecurityEventSubject)`

SetSubjects sets Subjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


