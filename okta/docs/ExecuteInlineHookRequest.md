# ExecuteInlineHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEventVersion** | Pointer to **string** | The inline hook cloud version | [optional] 
**ContentType** | Pointer to **string** | The inline hook request header content | [optional] 
**EventId** | Pointer to **string** | The individual inline hook request ID | [optional] 
**EventTime** | Pointer to **string** | The time the inline hook request was sent | [optional] 
**EventTypeVersion** | Pointer to **string** | The inline hook version | [optional] 
**Data** | Pointer to [**UserImportRequestData**](UserImportRequestData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The user import inline hook type is &#x60;com.okta.import.transform&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID of the user import inline hook | [optional] 
**RequestType** | Pointer to **string** | The type of inline hook request. For example, &#x60;com.okta.user.telephony.pre-enrollment&#x60;. | [optional] 

## Methods

### NewExecuteInlineHookRequest

`func NewExecuteInlineHookRequest() *ExecuteInlineHookRequest`

NewExecuteInlineHookRequest instantiates a new ExecuteInlineHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteInlineHookRequestWithDefaults

`func NewExecuteInlineHookRequestWithDefaults() *ExecuteInlineHookRequest`

NewExecuteInlineHookRequestWithDefaults instantiates a new ExecuteInlineHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEventVersion

`func (o *ExecuteInlineHookRequest) GetCloudEventVersion() string`

GetCloudEventVersion returns the CloudEventVersion field if non-nil, zero value otherwise.

### GetCloudEventVersionOk

`func (o *ExecuteInlineHookRequest) GetCloudEventVersionOk() (*string, bool)`

GetCloudEventVersionOk returns a tuple with the CloudEventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEventVersion

`func (o *ExecuteInlineHookRequest) SetCloudEventVersion(v string)`

SetCloudEventVersion sets CloudEventVersion field to given value.

### HasCloudEventVersion

`func (o *ExecuteInlineHookRequest) HasCloudEventVersion() bool`

HasCloudEventVersion returns a boolean if a field has been set.

### GetContentType

`func (o *ExecuteInlineHookRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ExecuteInlineHookRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ExecuteInlineHookRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ExecuteInlineHookRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEventId

`func (o *ExecuteInlineHookRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ExecuteInlineHookRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ExecuteInlineHookRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *ExecuteInlineHookRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *ExecuteInlineHookRequest) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ExecuteInlineHookRequest) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ExecuteInlineHookRequest) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *ExecuteInlineHookRequest) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventTypeVersion

`func (o *ExecuteInlineHookRequest) GetEventTypeVersion() string`

GetEventTypeVersion returns the EventTypeVersion field if non-nil, zero value otherwise.

### GetEventTypeVersionOk

`func (o *ExecuteInlineHookRequest) GetEventTypeVersionOk() (*string, bool)`

GetEventTypeVersionOk returns a tuple with the EventTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeVersion

`func (o *ExecuteInlineHookRequest) SetEventTypeVersion(v string)`

SetEventTypeVersion sets EventTypeVersion field to given value.

### HasEventTypeVersion

`func (o *ExecuteInlineHookRequest) HasEventTypeVersion() bool`

HasEventTypeVersion returns a boolean if a field has been set.

### GetData

`func (o *ExecuteInlineHookRequest) GetData() UserImportRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ExecuteInlineHookRequest) GetDataOk() (*UserImportRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ExecuteInlineHookRequest) SetData(v UserImportRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *ExecuteInlineHookRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *ExecuteInlineHookRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ExecuteInlineHookRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ExecuteInlineHookRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ExecuteInlineHookRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *ExecuteInlineHookRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExecuteInlineHookRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExecuteInlineHookRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ExecuteInlineHookRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetRequestType

`func (o *ExecuteInlineHookRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *ExecuteInlineHookRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *ExecuteInlineHookRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *ExecuteInlineHookRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


