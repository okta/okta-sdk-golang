# RegistrationInlineHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | The type of inline hook. The registration inline hook type is &#x60;com.okta.user.pre-registration&#x60;. | [optional] 
**RequestType** | Pointer to **string** | The type of registration hook. Use either &#x60;self.service.registration&#x60; or &#x60;progressive.profile&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID of the registration inline hook | [optional] 

## Methods

### NewRegistrationInlineHookRequest

`func NewRegistrationInlineHookRequest() *RegistrationInlineHookRequest`

NewRegistrationInlineHookRequest instantiates a new RegistrationInlineHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrationInlineHookRequestWithDefaults

`func NewRegistrationInlineHookRequestWithDefaults() *RegistrationInlineHookRequest`

NewRegistrationInlineHookRequestWithDefaults instantiates a new RegistrationInlineHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *RegistrationInlineHookRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *RegistrationInlineHookRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *RegistrationInlineHookRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *RegistrationInlineHookRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetRequestType

`func (o *RegistrationInlineHookRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *RegistrationInlineHookRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *RegistrationInlineHookRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *RegistrationInlineHookRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetSource

`func (o *RegistrationInlineHookRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RegistrationInlineHookRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RegistrationInlineHookRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RegistrationInlineHookRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


