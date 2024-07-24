# GetSsfStreams200Response

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

### NewGetSsfStreams200Response

`func NewGetSsfStreams200Response(delivery StreamConfigurationDelivery, eventsRequested []string, ) *GetSsfStreams200Response`

NewGetSsfStreams200Response instantiates a new GetSsfStreams200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSsfStreams200ResponseWithDefaults

`func NewGetSsfStreams200ResponseWithDefaults() *GetSsfStreams200Response`

NewGetSsfStreams200ResponseWithDefaults instantiates a new GetSsfStreams200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *GetSsfStreams200Response) GetAud() StreamConfigurationAud`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *GetSsfStreams200Response) GetAudOk() (*StreamConfigurationAud, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *GetSsfStreams200Response) SetAud(v StreamConfigurationAud)`

SetAud sets Aud field to given value.

### HasAud

`func (o *GetSsfStreams200Response) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetDelivery

`func (o *GetSsfStreams200Response) GetDelivery() StreamConfigurationDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *GetSsfStreams200Response) GetDeliveryOk() (*StreamConfigurationDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *GetSsfStreams200Response) SetDelivery(v StreamConfigurationDelivery)`

SetDelivery sets Delivery field to given value.


### GetEventsDelivered

`func (o *GetSsfStreams200Response) GetEventsDelivered() []string`

GetEventsDelivered returns the EventsDelivered field if non-nil, zero value otherwise.

### GetEventsDeliveredOk

`func (o *GetSsfStreams200Response) GetEventsDeliveredOk() (*[]string, bool)`

GetEventsDeliveredOk returns a tuple with the EventsDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsDelivered

`func (o *GetSsfStreams200Response) SetEventsDelivered(v []string)`

SetEventsDelivered sets EventsDelivered field to given value.

### HasEventsDelivered

`func (o *GetSsfStreams200Response) HasEventsDelivered() bool`

HasEventsDelivered returns a boolean if a field has been set.

### GetEventsRequested

`func (o *GetSsfStreams200Response) GetEventsRequested() []string`

GetEventsRequested returns the EventsRequested field if non-nil, zero value otherwise.

### GetEventsRequestedOk

`func (o *GetSsfStreams200Response) GetEventsRequestedOk() (*[]string, bool)`

GetEventsRequestedOk returns a tuple with the EventsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRequested

`func (o *GetSsfStreams200Response) SetEventsRequested(v []string)`

SetEventsRequested sets EventsRequested field to given value.


### GetEventsSupported

`func (o *GetSsfStreams200Response) GetEventsSupported() []string`

GetEventsSupported returns the EventsSupported field if non-nil, zero value otherwise.

### GetEventsSupportedOk

`func (o *GetSsfStreams200Response) GetEventsSupportedOk() (*[]string, bool)`

GetEventsSupportedOk returns a tuple with the EventsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsSupported

`func (o *GetSsfStreams200Response) SetEventsSupported(v []string)`

SetEventsSupported sets EventsSupported field to given value.

### HasEventsSupported

`func (o *GetSsfStreams200Response) HasEventsSupported() bool`

HasEventsSupported returns a boolean if a field has been set.

### GetFormat

`func (o *GetSsfStreams200Response) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *GetSsfStreams200Response) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *GetSsfStreams200Response) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *GetSsfStreams200Response) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetIss

`func (o *GetSsfStreams200Response) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *GetSsfStreams200Response) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *GetSsfStreams200Response) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *GetSsfStreams200Response) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetMinVerificationInterval

`func (o *GetSsfStreams200Response) GetMinVerificationInterval() int32`

GetMinVerificationInterval returns the MinVerificationInterval field if non-nil, zero value otherwise.

### GetMinVerificationIntervalOk

`func (o *GetSsfStreams200Response) GetMinVerificationIntervalOk() (*int32, bool)`

GetMinVerificationIntervalOk returns a tuple with the MinVerificationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinVerificationInterval

`func (o *GetSsfStreams200Response) SetMinVerificationInterval(v int32)`

SetMinVerificationInterval sets MinVerificationInterval field to given value.

### HasMinVerificationInterval

`func (o *GetSsfStreams200Response) HasMinVerificationInterval() bool`

HasMinVerificationInterval returns a boolean if a field has been set.

### SetMinVerificationIntervalNil

`func (o *GetSsfStreams200Response) SetMinVerificationIntervalNil(b bool)`

 SetMinVerificationIntervalNil sets the value for MinVerificationInterval to be an explicit nil

### UnsetMinVerificationInterval
`func (o *GetSsfStreams200Response) UnsetMinVerificationInterval()`

UnsetMinVerificationInterval ensures that no value is present for MinVerificationInterval, not even an explicit nil
### GetStreamId

`func (o *GetSsfStreams200Response) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *GetSsfStreams200Response) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *GetSsfStreams200Response) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *GetSsfStreams200Response) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


