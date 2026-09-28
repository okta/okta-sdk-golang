# ScriptInvokeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pid** | Pointer to **string** | Process ID of the launched process on the agent host | [optional] 
**StartTime** | Pointer to **time.Time** | ISO-8601 timestamp of when the process was started | [optional] 

## Methods

### NewScriptInvokeResponse

`func NewScriptInvokeResponse() *ScriptInvokeResponse`

NewScriptInvokeResponse instantiates a new ScriptInvokeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptInvokeResponseWithDefaults

`func NewScriptInvokeResponseWithDefaults() *ScriptInvokeResponse`

NewScriptInvokeResponseWithDefaults instantiates a new ScriptInvokeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPid

`func (o *ScriptInvokeResponse) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *ScriptInvokeResponse) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *ScriptInvokeResponse) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *ScriptInvokeResponse) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetStartTime

`func (o *ScriptInvokeResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScriptInvokeResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScriptInvokeResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ScriptInvokeResponse) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


