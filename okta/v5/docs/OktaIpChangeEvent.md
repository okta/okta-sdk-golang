# OktaIpChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentIpAddress** | **string** | Current IP address of the subject | 
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**InitiatingEntity** | Pointer to **string** | The entity that initiated the event | [optional] 
**PreviousIpAddress** | **string** | Previous IP address of the subject | 
**ReasonAdmin** | Pointer to [**CaepDeviceComplianceChangeEventReasonAdmin**](CaepDeviceComplianceChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepDeviceComplianceChangeEventReasonUser**](CaepDeviceComplianceChangeEventReasonUser.md) |  | [optional] 
**Subjects** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewOktaIpChangeEvent

`func NewOktaIpChangeEvent(currentIpAddress string, eventTimestamp int64, previousIpAddress string, subjects SecurityEventSubject, ) *OktaIpChangeEvent`

NewOktaIpChangeEvent instantiates a new OktaIpChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaIpChangeEventWithDefaults

`func NewOktaIpChangeEventWithDefaults() *OktaIpChangeEvent`

NewOktaIpChangeEventWithDefaults instantiates a new OktaIpChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentIpAddress

`func (o *OktaIpChangeEvent) GetCurrentIpAddress() string`

GetCurrentIpAddress returns the CurrentIpAddress field if non-nil, zero value otherwise.

### GetCurrentIpAddressOk

`func (o *OktaIpChangeEvent) GetCurrentIpAddressOk() (*string, bool)`

GetCurrentIpAddressOk returns a tuple with the CurrentIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIpAddress

`func (o *OktaIpChangeEvent) SetCurrentIpAddress(v string)`

SetCurrentIpAddress sets CurrentIpAddress field to given value.


### GetEventTimestamp

`func (o *OktaIpChangeEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *OktaIpChangeEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *OktaIpChangeEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetInitiatingEntity

`func (o *OktaIpChangeEvent) GetInitiatingEntity() string`

GetInitiatingEntity returns the InitiatingEntity field if non-nil, zero value otherwise.

### GetInitiatingEntityOk

`func (o *OktaIpChangeEvent) GetInitiatingEntityOk() (*string, bool)`

GetInitiatingEntityOk returns a tuple with the InitiatingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingEntity

`func (o *OktaIpChangeEvent) SetInitiatingEntity(v string)`

SetInitiatingEntity sets InitiatingEntity field to given value.

### HasInitiatingEntity

`func (o *OktaIpChangeEvent) HasInitiatingEntity() bool`

HasInitiatingEntity returns a boolean if a field has been set.

### GetPreviousIpAddress

`func (o *OktaIpChangeEvent) GetPreviousIpAddress() string`

GetPreviousIpAddress returns the PreviousIpAddress field if non-nil, zero value otherwise.

### GetPreviousIpAddressOk

`func (o *OktaIpChangeEvent) GetPreviousIpAddressOk() (*string, bool)`

GetPreviousIpAddressOk returns a tuple with the PreviousIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousIpAddress

`func (o *OktaIpChangeEvent) SetPreviousIpAddress(v string)`

SetPreviousIpAddress sets PreviousIpAddress field to given value.


### GetReasonAdmin

`func (o *OktaIpChangeEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *OktaIpChangeEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *OktaIpChangeEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *OktaIpChangeEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *OktaIpChangeEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *OktaIpChangeEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *OktaIpChangeEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *OktaIpChangeEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubjects

`func (o *OktaIpChangeEvent) GetSubjects() SecurityEventSubject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *OktaIpChangeEvent) GetSubjectsOk() (*SecurityEventSubject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *OktaIpChangeEvent) SetSubjects(v SecurityEventSubject)`

SetSubjects sets Subjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


