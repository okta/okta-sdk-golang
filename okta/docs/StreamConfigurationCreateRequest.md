# StreamConfigurationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delivery** | [**StreamConfigurationDelivery**](StreamConfigurationDelivery.md) |  | 
**EventsRequested** | **[]string** | The events (mapped by the array of event type URIs) that the receiver wants to receive | 
**Format** | Pointer to **string** | The Subject Identifier format expected for any SET transmitted. | [optional] 

## Methods

### NewStreamConfigurationCreateRequest

`func NewStreamConfigurationCreateRequest(delivery StreamConfigurationDelivery, eventsRequested []string, ) *StreamConfigurationCreateRequest`

NewStreamConfigurationCreateRequest instantiates a new StreamConfigurationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamConfigurationCreateRequestWithDefaults

`func NewStreamConfigurationCreateRequestWithDefaults() *StreamConfigurationCreateRequest`

NewStreamConfigurationCreateRequestWithDefaults instantiates a new StreamConfigurationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelivery

`func (o *StreamConfigurationCreateRequest) GetDelivery() StreamConfigurationDelivery`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *StreamConfigurationCreateRequest) GetDeliveryOk() (*StreamConfigurationDelivery, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *StreamConfigurationCreateRequest) SetDelivery(v StreamConfigurationDelivery)`

SetDelivery sets Delivery field to given value.


### GetEventsRequested

`func (o *StreamConfigurationCreateRequest) GetEventsRequested() []string`

GetEventsRequested returns the EventsRequested field if non-nil, zero value otherwise.

### GetEventsRequestedOk

`func (o *StreamConfigurationCreateRequest) GetEventsRequestedOk() (*[]string, bool)`

GetEventsRequestedOk returns a tuple with the EventsRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsRequested

`func (o *StreamConfigurationCreateRequest) SetEventsRequested(v []string)`

SetEventsRequested sets EventsRequested field to given value.


### GetFormat

`func (o *StreamConfigurationCreateRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *StreamConfigurationCreateRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *StreamConfigurationCreateRequest) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *StreamConfigurationCreateRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


