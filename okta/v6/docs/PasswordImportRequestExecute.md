# PasswordImportRequestExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEventVersion** | Pointer to **string** | The inline hook cloud version | [optional] 
**ContentType** | Pointer to **string** | The inline hook request header content | [optional] 
**EventId** | Pointer to **string** | The individual inline hook request ID | [optional] 
**EventTime** | Pointer to **string** | The time the inline hook request was sent | [optional] 
**EventTypeVersion** | Pointer to **string** | The inline hook version | [optional] 
**Data** | Pointer to [**PasswordImportRequestData**](PasswordImportRequestData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The password import inline hook type is &#x60;com.okta.user.credential.password.import&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID and URL of the password import inline hook | [optional] 

## Methods

### NewPasswordImportRequestExecute

`func NewPasswordImportRequestExecute() *PasswordImportRequestExecute`

NewPasswordImportRequestExecute instantiates a new PasswordImportRequestExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordImportRequestExecuteWithDefaults

`func NewPasswordImportRequestExecuteWithDefaults() *PasswordImportRequestExecute`

NewPasswordImportRequestExecuteWithDefaults instantiates a new PasswordImportRequestExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEventVersion

`func (o *PasswordImportRequestExecute) GetCloudEventVersion() string`

GetCloudEventVersion returns the CloudEventVersion field if non-nil, zero value otherwise.

### GetCloudEventVersionOk

`func (o *PasswordImportRequestExecute) GetCloudEventVersionOk() (*string, bool)`

GetCloudEventVersionOk returns a tuple with the CloudEventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEventVersion

`func (o *PasswordImportRequestExecute) SetCloudEventVersion(v string)`

SetCloudEventVersion sets CloudEventVersion field to given value.

### HasCloudEventVersion

`func (o *PasswordImportRequestExecute) HasCloudEventVersion() bool`

HasCloudEventVersion returns a boolean if a field has been set.

### GetContentType

`func (o *PasswordImportRequestExecute) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PasswordImportRequestExecute) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PasswordImportRequestExecute) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PasswordImportRequestExecute) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEventId

`func (o *PasswordImportRequestExecute) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *PasswordImportRequestExecute) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *PasswordImportRequestExecute) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *PasswordImportRequestExecute) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *PasswordImportRequestExecute) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *PasswordImportRequestExecute) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *PasswordImportRequestExecute) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *PasswordImportRequestExecute) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventTypeVersion

`func (o *PasswordImportRequestExecute) GetEventTypeVersion() string`

GetEventTypeVersion returns the EventTypeVersion field if non-nil, zero value otherwise.

### GetEventTypeVersionOk

`func (o *PasswordImportRequestExecute) GetEventTypeVersionOk() (*string, bool)`

GetEventTypeVersionOk returns a tuple with the EventTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeVersion

`func (o *PasswordImportRequestExecute) SetEventTypeVersion(v string)`

SetEventTypeVersion sets EventTypeVersion field to given value.

### HasEventTypeVersion

`func (o *PasswordImportRequestExecute) HasEventTypeVersion() bool`

HasEventTypeVersion returns a boolean if a field has been set.

### GetData

`func (o *PasswordImportRequestExecute) GetData() PasswordImportRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PasswordImportRequestExecute) GetDataOk() (*PasswordImportRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PasswordImportRequestExecute) SetData(v PasswordImportRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *PasswordImportRequestExecute) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *PasswordImportRequestExecute) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PasswordImportRequestExecute) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PasswordImportRequestExecute) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PasswordImportRequestExecute) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *PasswordImportRequestExecute) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PasswordImportRequestExecute) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PasswordImportRequestExecute) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PasswordImportRequestExecute) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


