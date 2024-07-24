# TelephonyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TelephonyRequestData**](TelephonyRequestData.md) |  | [optional] 
**EventType** | Pointer to **string** | The type of inline hook. The Telephony inline hook type is &#x60;com.okta.telephony.provider&#x60;. | [optional] 
**RequestType** | Pointer to **string** | The type of inline hook request. For example, &#x60;com.okta.user.telephony.pre-enrollment&#x60;. | [optional] 
**Source** | Pointer to **string** | The ID and URL of the Telephony inline hook | [optional] 

## Methods

### NewTelephonyRequest

`func NewTelephonyRequest() *TelephonyRequest`

NewTelephonyRequest instantiates a new TelephonyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyRequestWithDefaults

`func NewTelephonyRequestWithDefaults() *TelephonyRequest`

NewTelephonyRequestWithDefaults instantiates a new TelephonyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TelephonyRequest) GetData() TelephonyRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TelephonyRequest) GetDataOk() (*TelephonyRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TelephonyRequest) SetData(v TelephonyRequestData)`

SetData sets Data field to given value.

### HasData

`func (o *TelephonyRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEventType

`func (o *TelephonyRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TelephonyRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TelephonyRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TelephonyRequest) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetRequestType

`func (o *TelephonyRequest) GetRequestType() string`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *TelephonyRequest) GetRequestTypeOk() (*string, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *TelephonyRequest) SetRequestType(v string)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *TelephonyRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### GetSource

`func (o *TelephonyRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TelephonyRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TelephonyRequest) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TelephonyRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


