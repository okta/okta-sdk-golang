# RiskEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **time.Time** | Time stamp at which the event expires (Expressed as a UTC time zone using ISO 8601 format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;). If this optional field is not included, Okta automatically expires the event 24 hours after the event is consumed. | [optional] 
**Subjects** | [**[]RiskEventSubject**](RiskEventSubject.md) |  | 
**Timestamp** | Pointer to **time.Time** | Time stamp at which the event is produced (Expressed as a UTC time zone using ISO 8601 format: yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSS&#39;Z&#39;). | [optional] 

## Methods

### NewRiskEvent

`func NewRiskEvent(subjects []RiskEventSubject, ) *RiskEvent`

NewRiskEvent instantiates a new RiskEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskEventWithDefaults

`func NewRiskEventWithDefaults() *RiskEvent`

NewRiskEventWithDefaults instantiates a new RiskEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *RiskEvent) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *RiskEvent) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *RiskEvent) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *RiskEvent) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSubjects

`func (o *RiskEvent) GetSubjects() []RiskEventSubject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RiskEvent) GetSubjectsOk() (*[]RiskEventSubject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RiskEvent) SetSubjects(v []RiskEventSubject)`

SetSubjects sets Subjects field to given value.


### GetTimestamp

`func (o *RiskEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RiskEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RiskEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *RiskEvent) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


