# CaepCredentialChangeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | **string** | The type of action done towards the credential | 
**CredentialType** | **string** | The credential type of the changed credential. It will one of the supported enum values or any other credential type supported mutually by the Transmitter and the Receiver. | 
**EventTimestamp** | Pointer to **int64** | The time of the event (UNIX timestamp) | [optional] 
**Fido2Aaguid** | Pointer to **string** | FIDO2 Authenticator Attestation GUID | [optional] 
**FriendlyName** | Pointer to **string** | Credential friendly name | [optional] 
**InitiatingEntity** | Pointer to **string** | The entity that initiated the event | [optional] 
**ReasonAdmin** | Pointer to [**CaepCredentialChangeEventReasonAdmin**](CaepCredentialChangeEventReasonAdmin.md) |  | [optional] 
**ReasonUser** | Pointer to [**CaepCredentialChangeEventReasonUser**](CaepCredentialChangeEventReasonUser.md) |  | [optional] 
**Subject** | Pointer to [**SecurityEventSubject**](SecurityEventSubject.md) |  | [optional] 

## Methods

### NewCaepCredentialChangeEvent

`func NewCaepCredentialChangeEvent(changeType string, credentialType string, ) *CaepCredentialChangeEvent`

NewCaepCredentialChangeEvent instantiates a new CaepCredentialChangeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaepCredentialChangeEventWithDefaults

`func NewCaepCredentialChangeEventWithDefaults() *CaepCredentialChangeEvent`

NewCaepCredentialChangeEventWithDefaults instantiates a new CaepCredentialChangeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *CaepCredentialChangeEvent) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *CaepCredentialChangeEvent) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *CaepCredentialChangeEvent) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.


### GetCredentialType

`func (o *CaepCredentialChangeEvent) GetCredentialType() string`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *CaepCredentialChangeEvent) GetCredentialTypeOk() (*string, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *CaepCredentialChangeEvent) SetCredentialType(v string)`

SetCredentialType sets CredentialType field to given value.


### GetEventTimestamp

`func (o *CaepCredentialChangeEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *CaepCredentialChangeEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *CaepCredentialChangeEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.

### HasEventTimestamp

`func (o *CaepCredentialChangeEvent) HasEventTimestamp() bool`

HasEventTimestamp returns a boolean if a field has been set.

### GetFido2Aaguid

`func (o *CaepCredentialChangeEvent) GetFido2Aaguid() string`

GetFido2Aaguid returns the Fido2Aaguid field if non-nil, zero value otherwise.

### GetFido2AaguidOk

`func (o *CaepCredentialChangeEvent) GetFido2AaguidOk() (*string, bool)`

GetFido2AaguidOk returns a tuple with the Fido2Aaguid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2Aaguid

`func (o *CaepCredentialChangeEvent) SetFido2Aaguid(v string)`

SetFido2Aaguid sets Fido2Aaguid field to given value.

### HasFido2Aaguid

`func (o *CaepCredentialChangeEvent) HasFido2Aaguid() bool`

HasFido2Aaguid returns a boolean if a field has been set.

### GetFriendlyName

`func (o *CaepCredentialChangeEvent) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *CaepCredentialChangeEvent) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *CaepCredentialChangeEvent) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *CaepCredentialChangeEvent) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### GetInitiatingEntity

`func (o *CaepCredentialChangeEvent) GetInitiatingEntity() string`

GetInitiatingEntity returns the InitiatingEntity field if non-nil, zero value otherwise.

### GetInitiatingEntityOk

`func (o *CaepCredentialChangeEvent) GetInitiatingEntityOk() (*string, bool)`

GetInitiatingEntityOk returns a tuple with the InitiatingEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatingEntity

`func (o *CaepCredentialChangeEvent) SetInitiatingEntity(v string)`

SetInitiatingEntity sets InitiatingEntity field to given value.

### HasInitiatingEntity

`func (o *CaepCredentialChangeEvent) HasInitiatingEntity() bool`

HasInitiatingEntity returns a boolean if a field has been set.

### GetReasonAdmin

`func (o *CaepCredentialChangeEvent) GetReasonAdmin() CaepCredentialChangeEventReasonAdmin`

GetReasonAdmin returns the ReasonAdmin field if non-nil, zero value otherwise.

### GetReasonAdminOk

`func (o *CaepCredentialChangeEvent) GetReasonAdminOk() (*CaepCredentialChangeEventReasonAdmin, bool)`

GetReasonAdminOk returns a tuple with the ReasonAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonAdmin

`func (o *CaepCredentialChangeEvent) SetReasonAdmin(v CaepCredentialChangeEventReasonAdmin)`

SetReasonAdmin sets ReasonAdmin field to given value.

### HasReasonAdmin

`func (o *CaepCredentialChangeEvent) HasReasonAdmin() bool`

HasReasonAdmin returns a boolean if a field has been set.

### GetReasonUser

`func (o *CaepCredentialChangeEvent) GetReasonUser() CaepCredentialChangeEventReasonUser`

GetReasonUser returns the ReasonUser field if non-nil, zero value otherwise.

### GetReasonUserOk

`func (o *CaepCredentialChangeEvent) GetReasonUserOk() (*CaepCredentialChangeEventReasonUser, bool)`

GetReasonUserOk returns a tuple with the ReasonUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonUser

`func (o *CaepCredentialChangeEvent) SetReasonUser(v CaepCredentialChangeEventReasonUser)`

SetReasonUser sets ReasonUser field to given value.

### HasReasonUser

`func (o *CaepCredentialChangeEvent) HasReasonUser() bool`

HasReasonUser returns a boolean if a field has been set.

### GetSubject

`func (o *CaepCredentialChangeEvent) GetSubject() SecurityEventSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CaepCredentialChangeEvent) GetSubjectOk() (*SecurityEventSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CaepCredentialChangeEvent) SetSubject(v SecurityEventSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CaepCredentialChangeEvent) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


