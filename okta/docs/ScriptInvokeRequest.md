# ScriptInvokeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to **map[string]interface{}** | Optional JSON payload forwarded to the script. Maximum size is 4 KB. | [optional] 
**Type** | **string** | The type of script to invoke | 

## Methods

### NewScriptInvokeRequest

`func NewScriptInvokeRequest(type_ string, ) *ScriptInvokeRequest`

NewScriptInvokeRequest instantiates a new ScriptInvokeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptInvokeRequestWithDefaults

`func NewScriptInvokeRequestWithDefaults() *ScriptInvokeRequest`

NewScriptInvokeRequestWithDefaults instantiates a new ScriptInvokeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *ScriptInvokeRequest) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ScriptInvokeRequest) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ScriptInvokeRequest) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ScriptInvokeRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetType

`func (o *ScriptInvokeRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScriptInvokeRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScriptInvokeRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


