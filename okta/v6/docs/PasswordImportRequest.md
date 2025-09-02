# PasswordImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**PasswordImportRequestData**](PasswordImportRequestData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The password import inline hook type is &#x60;com.okta.user.credential.password.import&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID and URL of the password import inline hook | [optional] 

## Methods

### NewPasswordImportRequest

`func NewPasswordImportRequest() *PasswordImportRequest`

NewPasswordImportRequest instantiates a new PasswordImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordImportRequestWithDefaults

`func NewPasswordImportRequestWithDefaults() *PasswordImportRequest`

NewPasswordImportRequestWithDefaults instantiates a new PasswordImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PasswordImportRequest) GetData() PasswordImportRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PasswordImportRequest) GetDataOk() (*PasswordImportRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PasswordImportRequest) SetData(v PasswordImportRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *PasswordImportRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *PasswordImportRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PasswordImportRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PasswordImportRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *PasswordImportRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSource

`func (o *PasswordImportRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PasswordImportRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PasswordImportRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PasswordImportRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


