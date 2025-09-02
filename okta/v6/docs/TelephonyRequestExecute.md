# TelephonyRequestExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEventVersion** | Pointer to **string** | The inline hook cloud version | [optional] 
**ContentType** | Pointer to **string** | The inline hook request header content | [optional] 
**EventId** | Pointer to **string** | The individual inline hook request ID | [optional] 
**EventTime** | Pointer to **string** | The time the inline hook request was sent | [optional] 
**EventTypeVersion** | Pointer to **string** | The inline hook version | [optional] 
**Data** | Pointer to [**TelephonyRequestData**](TelephonyRequestData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The telephony inline hook type is &#x60;com.okta.telephony.provider&#x60;. | [optional] 
**RequestType** | Pointer to **string** | The type of inline hook request. For example, &#x60;com.okta.user.telephony.pre-enrollment&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID and URL of the telephony inline hook | [optional] 

## Methods

### NewTelephonyRequestExecute

`func NewTelephonyRequestExecute() *TelephonyRequestExecute`

NewTelephonyRequestExecute instantiates a new TelephonyRequestExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyRequestExecuteWithDefaults

`func NewTelephonyRequestExecuteWithDefaults() *TelephonyRequestExecute`

NewTelephonyRequestExecuteWithDefaults instantiates a new TelephonyRequestExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEventVersion

`func (o *TelephonyRequestExecute) GetCloudEventVersion() string`

GetCloudEventVersion returns the CloudEventVersion field if non-nil, zero value otherwise.

### GetCloudEventVersionOk

`func (o *TelephonyRequestExecute) GetCloudEventVersionOk() (*string, bool)`

GetCloudEventVersionOk returns a tuple with the CloudEventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEventVersion

`func (o *TelephonyRequestExecute) SetCloudEventVersion(v string)`

SetCloudEventVersion sets CloudEventVersion field to given value.

### HasCloudEventVersion

`func (o *TelephonyRequestExecute) HasCloudEventVersion() bool`

HasCloudEventVersion returns a boolean if a field has been set.

### GetContentType

`func (o *TelephonyRequestExecute) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *TelephonyRequestExecute) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *TelephonyRequestExecute) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *TelephonyRequestExecute) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEventId

`func (o *TelephonyRequestExecute) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TelephonyRequestExecute) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TelephonyRequestExecute) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *TelephonyRequestExecute) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *TelephonyRequestExecute) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *TelephonyRequestExecute) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *TelephonyRequestExecute) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *TelephonyRequestExecute) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventTypeVersion

`func (o *TelephonyRequestExecute) GetEventTypeVersion() string`

GetEventTypeVersion returns the EventTypeVersion field if non-nil, zero value otherwise.

### GetEventTypeVersionOk

`func (o *TelephonyRequestExecute) GetEventTypeVersionOk() (*string, bool)`

GetEventTypeVersionOk returns a tuple with the EventTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeVersion

`func (o *TelephonyRequestExecute) SetEventTypeVersion(v string)`

SetEventTypeVersion sets EventTypeVersion field to given value.

### HasEventTypeVersion

`func (o *TelephonyRequestExecute) HasEventTypeVersion() bool`

HasEventTypeVersion returns a boolean if a field has been set.

### GetData

`func (o *TelephonyRequestExecute) GetData() TelephonyRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TelephonyRequestExecute) GetDataOk() (*TelephonyRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TelephonyRequestExecute) SetData(v TelephonyRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *TelephonyRequestExecute) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *TelephonyRequestExecute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TelephonyRequestExecute) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TelephonyRequestExecute) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TelephonyRequestExecute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetRequestType

`func (o *TelephonyRequestExecute) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TelephonyRequestExecute) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TelephonyRequestExecute) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TelephonyRequestExecute) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetSource

`func (o *TelephonyRequestExecute) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TelephonyRequestExecute) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TelephonyRequestExecute) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TelephonyRequestExecute) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


