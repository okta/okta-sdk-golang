# UserImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserImportRequestData**](UserImportRequestData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The user import inline hook type is &#x60;com.okta.import.transform&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID of the user import inline hook | [optional] 

## Methods

### NewUserImportRequest

`func NewUserImportRequest() *UserImportRequest`

NewUserImportRequest instantiates a new UserImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestWithDefaults

`func NewUserImportRequestWithDefaults() *UserImportRequest`

NewUserImportRequestWithDefaults instantiates a new UserImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *UserImportRequest) GetData() UserImportRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserImportRequest) GetDataOk() (*UserImportRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserImportRequest) SetData(v UserImportRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *UserImportRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *UserImportRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *UserImportRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *UserImportRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *UserImportRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *UserImportRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UserImportRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UserImportRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *UserImportRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


