# CaepSessionRevokedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentIp** | Pointer to **string** | Current IP of the session | [optional] 
**CurrentUserAgent** | Pointer to **string** | Current User Agent of the session | [optional] 
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**InitiatingEntity** | Pointer to **string** | The entity that initiated the event | [optional] 
**LastKnownIp** | Pointer to **string** | Last known IP of the session | [optional] 
**LastKnownUserAgent** | Pointer to **string** | Last known User Agent of the session | [optional] 
**ReasonAdmin** | Pointer to [**CaepDeviceComplianceChangeEventReasonAdmin**](CaepDeviceComplianceChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepDeviceComplianceChangeEventReasonUser**](CaepDeviceComplianceChangeEventReasonUser.md) |  | [optional] 
**Subjects** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewCaepSessionRevokedEvent

`func NewCaepSessionRevokedEvent(eventTimestamp int64, subjects SecurityEventSubject, ) *CaepSessionRevokedEvent`

NewCaepSessionRevokedEvent instantiates a new CaepSessionRevokedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaepSessionRevokedEventWithDefaults

`func NewCaepSessionRevokedEventWithDefaults() *CaepSessionRevokedEvent`

NewCaepSessionRevokedEventWithDefaults instantiates a new CaepSessionRevokedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentIp

`func (o *CaepSessionRevokedEvent) GetCurrentIp() string`

GetCurrentIp returns the CurrentIp field if non-nil, zero value otherwise.

### GetCurrentIpOk

`func (o *CaepSessionRevokedEvent) GetCurrentIpOk() (*string, bool)`

GetCurrentIpOk returns a tuple with the CurrentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentIp

`func (o *CaepSessionRevokedEvent) SetCurrentIp(v string)`

SetCurrentIp sets CurrentIp field to given value.

### HasCurrentIp

`func (o *CaepSessionRevokedEvent) HasCurrentIp() bool`

HasCurrentIp returns a boolean if a field has been set.

### GetCurrentUserAgent

`func (o *CaepSessionRevokedEvent) GetCurrentUserAgent() string`

GetCurrentUserAgent returns the CurrentUserAgent field if non-nil, zero value otherwise.

### GetCurrentUserAgentOk

`func (o *CaepSessionRevokedEvent) GetCurrentUserAgentOk() (*string, bool)`

GetCurrentUserAgentOk returns a tuple with the CurrentUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserAgent

`func (o *CaepSessionRevokedEvent) SetCurrentUserAgent(v string)`

SetCurrentUserAgent sets CurrentUserAgent field to given value.

### HasCurrentUserAgent

`func (o *CaepSessionRevokedEvent) HasCurrentUserAgent() bool`

HasCurrentUserAgent returns a boolean if a field has been set.

### GetEventTimestamp

`func (o *CaepSessionRevokedEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CaepSessionRevokedEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CaepSessionRevokedEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetInitiatingEntity

`func (o *CaepSessionRevokedEvent) GetInitiatingEntity() string`

GetInitiatingEntity returns the InitiatingEntity field if non-nil, zero value otherwise.

### GetInitiatingEntityOk

`func (o *CaepSessionRevokedEvent) GetInitiatingEntityOk() (*string, bool)`

GetInitiatingEntityOk returns a tuple with the InitiatingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingEntity

`func (o *CaepSessionRevokedEvent) SetInitiatingEntity(v string)`

SetInitiatingEntity sets InitiatingEntity field to given value.

### HasInitiatingEntity

`func (o *CaepSessionRevokedEvent) HasInitiatingEntity() bool`

HasInitiatingEntity returns a boolean if a field has been set.

### GetLastKnownIp

`func (o *CaepSessionRevokedEvent) GetLastKnownIp() string`

GetLastKnownIp returns the LastKnownIp field if non-nil, zero value otherwise.

### GetLastKnownIpOk

`func (o *CaepSessionRevokedEvent) GetLastKnownIpOk() (*string, bool)`

GetLastKnownIpOk returns a tuple with the LastKnownIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownIp

`func (o *CaepSessionRevokedEvent) SetLastKnownIp(v string)`

SetLastKnownIp sets LastKnownIp field to given value.

### HasLastKnownIp

`func (o *CaepSessionRevokedEvent) HasLastKnownIp() bool`

HasLastKnownIp returns a boolean if a field has been set.

### GetLastKnownUserAgent

`func (o *CaepSessionRevokedEvent) GetLastKnownUserAgent() string`

GetLastKnownUserAgent returns the LastKnownUserAgent field if non-nil, zero value otherwise.

### GetLastKnownUserAgentOk

`func (o *CaepSessionRevokedEvent) GetLastKnownUserAgentOk() (*string, bool)`

GetLastKnownUserAgentOk returns a tuple with the LastKnownUserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownUserAgent

`func (o *CaepSessionRevokedEvent) SetLastKnownUserAgent(v string)`

SetLastKnownUserAgent sets LastKnownUserAgent field to given value.

### HasLastKnownUserAgent

`func (o *CaepSessionRevokedEvent) HasLastKnownUserAgent() bool`

HasLastKnownUserAgent returns a boolean if a field has been set.

### GetReasonAdmin

`func (o *CaepSessionRevokedEvent) GetReasonAdmin() CaepDeviceComplianceChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *CaepSessionRevokedEvent) GetReasonAdminOk() (*CaepDeviceComplianceChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *CaepSessionRevokedEvent) SetReasonAdmin(v CaepDeviceComplianceChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *CaepSessionRevokedEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *CaepSessionRevokedEvent) GetReasonUser() CaepDeviceComplianceChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *CaepSessionRevokedEvent) GetReasonUserOk() (*CaepDeviceComplianceChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *CaepSessionRevokedEvent) SetReasonUser(v CaepDeviceComplianceChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *CaepSessionRevokedEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubjects

`func (o *CaepSessionRevokedEvent) GetSubjects() SecurityEventSubject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *CaepSessionRevokedEvent) GetSubjectsOk() (*SecurityEventSubject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *CaepSessionRevokedEvent) SetSubjects(v SecurityEventSubject)`

SetSubjects sets Subjects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


