# UserImportRequestExecute

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

## Methods

### NewUserImportRequestExecute

`func NewUserImportRequestExecute() *UserImportRequestExecute`

NewUserImportRequestExecute instantiates a new UserImportRequestExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestExecuteWithDefaults

`func NewUserImportRequestExecuteWithDefaults() *UserImportRequestExecute`

NewUserImportRequestExecuteWithDefaults instantiates a new UserImportRequestExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEventVersion

`func (o *UserImportRequestExecute) GetCloudEventVersion() string`

GetCloudEventVersion returns the CloudEventVersion field if non-nil, zero value otherwise.

### GetCloudEventVersionOk

`func (o *UserImportRequestExecute) GetCloudEventVersionOk() (*string, bool)`

GetCloudEventVersionOk returns a tuple with the CloudEventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEventVersion

`func (o *UserImportRequestExecute) SetCloudEventVersion(v string)`

SetCloudEventVersion sets CloudEventVersion field to given value.

### HasCloudEventVersion

`func (o *UserImportRequestExecute) HasCloudEventVersion() bool`

HasCloudEventVersion returns a boolean if a field has been set.

### GetContentType

`func (o *UserImportRequestExecute) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UserImportRequestExecute) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UserImportRequestExecute) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *UserImportRequestExecute) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEventId

`func (o *UserImportRequestExecute) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *UserImportRequestExecute) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *UserImportRequestExecute) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *UserImportRequestExecute) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *UserImportRequestExecute) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *UserImportRequestExecute) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *UserImportRequestExecute) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *UserImportRequestExecute) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventTypeVersion

`func (o *UserImportRequestExecute) GetEventTypeVersion() string`

GetEventTypeVersion returns the EventTypeVersion field if non-nil, zero value otherwise.

### GetEventTypeVersionOk

`func (o *UserImportRequestExecute) GetEventTypeVersionOk() (*string, bool)`

GetEventTypeVersionOk returns a tuple with the EventTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeVersion

`func (o *UserImportRequestExecute) SetEventTypeVersion(v string)`

SetEventTypeVersion sets EventTypeVersion field to given value.

### HasEventTypeVersion

`func (o *UserImportRequestExecute) HasEventTypeVersion() bool`

HasEventTypeVersion returns a boolean if a field has been set.

### GetData

`func (o *UserImportRequestExecute) GetData() UserImportRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserImportRequestExecute) GetDataOk() (*UserImportRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserImportRequestExecute) SetData(v UserImportRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *UserImportRequestExecute) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *UserImportRequestExecute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *UserImportRequestExecute) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *UserImportRequestExecute) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *UserImportRequestExecute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *UserImportRequestExecute) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserImportRequestExecute) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserImportRequestExecute) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UserImportRequestExecute) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


