# SAMLPayloadExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEventVersion** | Pointer to **string** | The inline hook cloud version | [optional] 
**ContentType** | Pointer to **string** | The inline hook request header content | [optional] 
**EventId** | Pointer to **string** | The individual inline hook request ID | [optional] 
**EventTime** | Pointer to **string** | The time the inline hook request was sent | [optional] 
**EventTypeVersion** | Pointer to **string** | The inline hook version | [optional] 
**Data** | Pointer to [**SAMLPayLoadData**](SAMLPayLoadData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The SAML assertion inline hook type is &#x60;com.okta.saml.tokens.transform&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID and URL of the SAML assertion inline hook | [optional] 

## Methods

### NewSAMLPayloadExecute

`func NewSAMLPayloadExecute() *SAMLPayloadExecute`

NewSAMLPayloadExecute instantiates a new SAMLPayloadExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLPayloadExecuteWithDefaults

`func NewSAMLPayloadExecuteWithDefaults() *SAMLPayloadExecute`

NewSAMLPayloadExecuteWithDefaults instantiates a new SAMLPayloadExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEventVersion

`func (o *SAMLPayloadExecute) GetCloudEventVersion() string`

GetCloudEventVersion returns the CloudEventVersion field if non-nil, zero value otherwise.

### GetCloudEventVersionOk

`func (o *SAMLPayloadExecute) GetCloudEventVersionOk() (*string, bool)`

GetCloudEventVersionOk returns a tuple with the CloudEventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEventVersion

`func (o *SAMLPayloadExecute) SetCloudEventVersion(v string)`

SetCloudEventVersion sets CloudEventVersion field to given value.

### HasCloudEventVersion

`func (o *SAMLPayloadExecute) HasCloudEventVersion() bool`

HasCloudEventVersion returns a boolean if a field has been set.

### GetContentType

`func (o *SAMLPayloadExecute) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *SAMLPayloadExecute) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *SAMLPayloadExecute) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *SAMLPayloadExecute) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEventId

`func (o *SAMLPayloadExecute) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SAMLPayloadExecute) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SAMLPayloadExecute) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *SAMLPayloadExecute) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *SAMLPayloadExecute) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SAMLPayloadExecute) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SAMLPayloadExecute) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *SAMLPayloadExecute) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventTypeVersion

`func (o *SAMLPayloadExecute) GetEventTypeVersion() string`

GetEventTypeVersion returns the EventTypeVersion field if non-nil, zero value otherwise.

### GetEventTypeVersionOk

`func (o *SAMLPayloadExecute) GetEventTypeVersionOk() (*string, bool)`

GetEventTypeVersionOk returns a tuple with the EventTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeVersion

`func (o *SAMLPayloadExecute) SetEventTypeVersion(v string)`

SetEventTypeVersion sets EventTypeVersion field to given value.

### HasEventTypeVersion

`func (o *SAMLPayloadExecute) HasEventTypeVersion() bool`

HasEventTypeVersion returns a boolean if a field has been set.

### GetData

`func (o *SAMLPayloadExecute) GetData() SAMLPayLoadData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SAMLPayloadExecute) GetDataOk() (*SAMLPayLoadData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SAMLPayloadExecute) SetData(v SAMLPayLoadData)`

SetData sets Data field to given value.

### HasData

`func (o *SAMLPayloadExecute) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *SAMLPayloadExecute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SAMLPayloadExecute) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SAMLPayloadExecute) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *SAMLPayloadExecute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *SAMLPayloadExecute) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SAMLPayloadExecute) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SAMLPayloadExecute) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *SAMLPayloadExecute) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


