# TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudEventVersion** | Pointer to **string** | The inline hook cloud version | [optional] 
**ContentType** | Pointer to **string** | The inline hook request header content | [optional] 
**EventId** | Pointer to **string** | The individual inline hook request ID | [optional] 
**EventTime** | Pointer to **string** | The time the inline hook request was sent | [optional] 
**EventTypeVersion** | Pointer to **string** | The inline hook version | [optional] 
**Data** | Pointer to [**TokenPayLoadData**](TokenPayLoadData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The token inline hook type is &#x60;com.okta.oauth2.tokens.transform&#x60;. | [optional] 
**Source** | Pointer to **string** | The URL of the token inline hook | [optional] 

## Methods

### NewTokenRequest

`func NewTokenRequest() *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudEventVersion

`func (o *TokenRequest) GetCloudEventVersion() string`

GetCloudEventVersion returns the CloudEventVersion field if non-nil, zero value otherwise.

### GetCloudEventVersionOk

`func (o *TokenRequest) GetCloudEventVersionOk() (*string, bool)`

GetCloudEventVersionOk returns a tuple with the CloudEventVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudEventVersion

`func (o *TokenRequest) SetCloudEventVersion(v string)`

SetCloudEventVersion sets CloudEventVersion field to given value.

### HasCloudEventVersion

`func (o *TokenRequest) HasCloudEventVersion() bool`

HasCloudEventVersion returns a boolean if a field has been set.

### GetContentType

`func (o *TokenRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *TokenRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *TokenRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *TokenRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetEventId

`func (o *TokenRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *TokenRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *TokenRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *TokenRequest) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventTime

`func (o *TokenRequest) GetEventTime() string`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *TokenRequest) GetEventTimeOk() (*string, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *TokenRequest) SetEventTime(v string)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *TokenRequest) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventTypeVersion

`func (o *TokenRequest) GetEventTypeVersion() string`

GetEventTypeVersion returns the EventTypeVersion field if non-nil, zero value otherwise.

### GetEventTypeVersionOk

`func (o *TokenRequest) GetEventTypeVersionOk() (*string, bool)`

GetEventTypeVersionOk returns a tuple with the EventTypeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeVersion

`func (o *TokenRequest) SetEventTypeVersion(v string)`

SetEventTypeVersion sets EventTypeVersion field to given value.

### HasEventTypeVersion

`func (o *TokenRequest) HasEventTypeVersion() bool`

HasEventTypeVersion returns a boolean if a field has been set.

### GetData

`func (o *TokenRequest) GetData() TokenPayLoadData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TokenRequest) GetDataOk() (*TokenPayLoadData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TokenRequest) SetData(v TokenPayLoadData)`

SetData sets Data field to given value.

### HasData

`func (o *TokenRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *TokenRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TokenRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TokenRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TokenRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *TokenRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TokenRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TokenRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TokenRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


