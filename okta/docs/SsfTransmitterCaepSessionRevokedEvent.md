# SsfTransmitterCaepSessionRevokedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestamp** | Pointer to **int64** | The time of the event (UNIX timestamp) | [optional] 
**ReasonAdmin** | Pointer to [**CaepCredentialChangeEventReasonAdmin**](CaepCredentialChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepCredentialChangeEventReasonUser**](CaepCredentialChangeEventReasonUser.md) |  | [optional] 
**Subject** | Pointer to [**SsfTransmitterSecurityEventSubject**](SsfTransmitterSecurityEventSubject.md) |  | [optional] 

## Methods

### NewSsfTransmitterCaepSessionRevokedEvent

`func NewSsfTransmitterCaepSessionRevokedEvent() *SsfTransmitterCaepSessionRevokedEvent`

NewSsfTransmitterCaepSessionRevokedEvent instantiates a new SsfTransmitterCaepSessionRevokedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsfTransmitterCaepSessionRevokedEventWithDefaults

`func NewSsfTransmitterCaepSessionRevokedEventWithDefaults() *SsfTransmitterCaepSessionRevokedEvent`

NewSsfTransmitterCaepSessionRevokedEventWithDefaults instantiates a new SsfTransmitterCaepSessionRevokedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTimestamp

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *SsfTransmitterCaepSessionRevokedEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *SsfTransmitterCaepSessionRevokedEvent) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### GetReasonAdmin

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetReasonAdmin() CaepCredentialChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetReasonAdminOk() (*CaepCredentialChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *SsfTransmitterCaepSessionRevokedEvent) SetReasonAdmin(v CaepCredentialChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *SsfTransmitterCaepSessionRevokedEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetReasonUser() CaepCredentialChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetReasonUserOk() (*CaepCredentialChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *SsfTransmitterCaepSessionRevokedEvent) SetReasonUser(v CaepCredentialChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *SsfTransmitterCaepSessionRevokedEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubject

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetSubject() SsfTransmitterSecurityEventSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SsfTransmitterCaepSessionRevokedEvent) GetSubjectOk() (*SsfTransmitterSecurityEventSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SsfTransmitterCaepSessionRevokedEvent) SetSubject(v SsfTransmitterSecurityEventSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SsfTransmitterCaepSessionRevokedEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


