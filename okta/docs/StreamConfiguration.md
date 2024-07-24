# StreamConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to [**StreamConfigurationAud**](StreamConfigurationAud.md) |  | [optional] 
**Delivery** | [**StreamConfigurationDelivery**](StreamConfigurationDelivery.md) |  | 
**EventsDelivered** | Pointer to **[]string** | The events (mapped by the array of event type URIs) that the transmitter actually delivers to the SSF Stream.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter. | [optional] 
**EventsRequested** | **[]string** | The events (mapped by the array of event type URIs) that the receiver wants to receive | 
**EventsSupported** | Pointer to **[]string** | An array of event type URIs that the transmitter supports.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter. | [optional] 
**Format** | Pointer to **string** | The Subject Identifier format expected for any SET transmitted. | [optional] 
**Iss** | Pointer to **string** | The issuer used in Security Event Tokens (SETs). This value is set as &#x60;iss&#x60; in the claim.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter. | [optional] 
**MinVerificationInterval** | Pointer to **NullableInt32** | The minimum amount of time, in seconds, between two verification requests.  A read-only parameter that is set by the transmitter. If this parameter is included in the request, the value must match the expected value from the transmitter. | [optional] 
**StreamId** | Pointer to **string** | The ID of the SSF Stream configuration | [optional] 

## Methods

### NewStreamConfiguration

`func NewStreamConfiguration(delivery StreamConfigurationDelivery, eventsRequested []string, ) *StreamConfiguration`

NewStreamConfiguration instantiates a new StreamConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConfigurationWithDefaults

`func NewStreamConfigurationWithDefaults() *StreamConfiguration`

NewStreamConfigurationWithDefaults instantiates a new StreamConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *StreamConfiguration) GetAud() StreamConfigurationAud`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *StreamConfiguration) GetAudOk() (*StreamConfigurationAud, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *StreamConfiguration) SetAud(v StreamConfigurationAud)`

SetAud sets Aud field to given value.

### HasAud

`func (o *StreamConfiguration) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetDelivery

`func (o *StreamConfiguration) GetDelivery() StreamConfigurationDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *StreamConfiguration) GetDeliveryOk() (*StreamConfigurationDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *StreamConfiguration) SetDelivery(v StreamConfigurationDelivery)`

SetDelivery sets Delivery field to given value.


### GetEventsDelivered

`func (o *StreamConfiguration) GetEventsDelivered() []string`

GetEventsDelivered returns the EventsDelivered field if non-nil, zero value otherwise.

### GetEventsDeliveredOk

`func (o *StreamConfiguration) GetEventsDeliveredOk() (*[]string, bool)`

GetEventsDeliveredOk returns a tuple with the EventsDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDelivered

`func (o *StreamConfiguration) SetEventsDelivered(v []string)`

SetEventsDelivered sets EventsDelivered field to given value.

### HasEventsDelivered

`func (o *StreamConfiguration) HasEventsDelivered() bool`

HasEventsDelivered returns a boolean if a field has been set.

### GetEventsRequested

`func (o *StreamConfiguration) GetEventsRequested() []string`

GetEventsRequested returns the EventsRequested field if non-nil, zero value otherwise.

### GetEventsRequestedOk

`func (o *StreamConfiguration) GetEventsRequestedOk() (*[]string, bool)`

GetEventsRequestedOk returns a tuple with the EventsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRequested

`func (o *StreamConfiguration) SetEventsRequested(v []string)`

SetEventsRequested sets EventsRequested field to given value.


### GetEventsSupported

`func (o *StreamConfiguration) GetEventsSupported() []string`

GetEventsSupported returns the EventsSupported field if non-nil, zero value otherwise.

### GetEventsSupportedOk

`func (o *StreamConfiguration) GetEventsSupportedOk() (*[]string, bool)`

GetEventsSupportedOk returns a tuple with the EventsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsSupported

`func (o *StreamConfiguration) SetEventsSupported(v []string)`

SetEventsSupported sets EventsSupported field to given value.

### HasEventsSupported

`func (o *StreamConfiguration) HasEventsSupported() bool`

HasEventsSupported returns a boolean if a field has been set.

### GetFormat

`func (o *StreamConfiguration) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StreamConfiguration) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StreamConfiguration) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *StreamConfiguration) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIss

`func (o *StreamConfiguration) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *StreamConfiguration) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *StreamConfiguration) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *StreamConfiguration) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetMinVerificationInterval

`func (o *StreamConfiguration) GetMinVerificationInterval() int32`

GetMinVerificationInterval returns the MinVerificationInterval field if non-nil, zero value otherwise.

### GetMinVerificationIntervalOk

`func (o *StreamConfiguration) GetMinVerificationIntervalOk() (*int32, bool)`

GetMinVerificationIntervalOk returns a tuple with the MinVerificationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVerificationInterval

`func (o *StreamConfiguration) SetMinVerificationInterval(v int32)`

SetMinVerificationInterval sets MinVerificationInterval field to given value.

### HasMinVerificationInterval

`func (o *StreamConfiguration) HasMinVerificationInterval() bool`

HasMinVerificationInterval returns a boolean if a field has been set.

### SetMinVerificationIntervalNil

`func (o *StreamConfiguration) SetMinVerificationIntervalNil(b bool)`

 SetMinVerificationIntervalNil sets the value for MinVerificationInterval to be an explicit nil

### UnsetMinVerificationInterval
`func (o *StreamConfiguration) UnsetMinVerificationInterval()`

UnsetMinVerificationInterval ensures that no value is present for MinVerificationInterval, not even an explicit nil
### GetStreamId

`func (o *StreamConfiguration) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StreamConfiguration) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StreamConfiguration) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *StreamConfiguration) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


