# RiscIdentifierChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTimestamp** | **int64** | The time of the event (UNIX timestamp) | 
**NewValue** | Pointer to **string** | The new identifier value | [optional] 
**Subject** | [**SecurityEventSubject**](SecurityEventSubject.md) |  | 

## Methods

### NewRiscIdentifierChangedEvent

`func NewRiscIdentifierChangedEvent(eventTimestamp int64, subject SecurityEventSubject, ) *RiscIdentifierChangedEvent`

NewRiscIdentifierChangedEvent instantiates a new RiscIdentifierChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiscIdentifierChangedEventWithDefaults

`func NewRiscIdentifierChangedEventWithDefaults() *RiscIdentifierChangedEvent`

NewRiscIdentifierChangedEventWithDefaults instantiates a new RiscIdentifierChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTimestamp

`func (o *RiscIdentifierChangedEvent) GetEventTimestamp() int64`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *RiscIdentifierChangedEvent) GetEventTimestampOk() (*int64, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *RiscIdentifierChangedEvent) SetEventTimestamp(v int64)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetNewValue

`func (o *RiscIdentifierChangedEvent) GetNewValue() string`

GetNewValue returns the NewValue field if non-nil, zero value otherwise.

### GetNewValueOk

`func (o *RiscIdentifierChangedEvent) GetNewValueOk() (*string, bool)`

GetNewValueOk returns a tuple with the NewValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewValue

`func (o *RiscIdentifierChangedEvent) SetNewValue(v string)`

SetNewValue sets NewValue field to given value.

### HasNewValue

`func (o *RiscIdentifierChangedEvent) HasNewValue() bool`

HasNewValue returns a boolean if a field has been set.

### GetSubject

`func (o *RiscIdentifierChangedEvent) GetSubject() SecurityEventSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RiscIdentifierChangedEvent) GetSubjectOk() (*SecurityEventSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RiscIdentifierChangedEvent) SetSubject(v SecurityEventSubject)`

SetSubject sets Subject field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


