# SecurityEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**Subject** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewSecurityEvent

`func NewSecurityEvent(eventTimestamp int64, subject SecurityEventSubject, ) *SecurityEvent`

NewSecurityEvent instantiates a new SecurityEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityEventWithDefaults

`func NewSecurityEventWithDefaults() *SecurityEvent`

NewSecurityEventWithDefaults instantiates a new SecurityEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTimestamp

`func (o *SecurityEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *SecurityEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *SecurityEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetSubject

`func (o *SecurityEvent) GetSubject() SecurityEventSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SecurityEvent) GetSubjectOk() (*SecurityEventSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SecurityEvent) SetSubject(v SecurityEventSubject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


