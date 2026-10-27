# DeviceCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | Capability identifier | 
**Reasons** | Pointer to **[]string** | List of possible reason codes for unmet preconditions. It&#39;s only present when the &#x60;status&#x60; is &#x60;NOT_READY&#x60;. | [optional] 
**Status** | **string** | Readiness of the capability | 

## Methods

### NewDeviceCapability

`func NewDeviceCapability(capability string, status string, ) *DeviceCapability`

NewDeviceCapability instantiates a new DeviceCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCapabilityWithDefaults

`func NewDeviceCapabilityWithDefaults() *DeviceCapability`

NewDeviceCapabilityWithDefaults instantiates a new DeviceCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *DeviceCapability) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *DeviceCapability) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *DeviceCapability) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetReasons

`func (o *DeviceCapability) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *DeviceCapability) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *DeviceCapability) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *DeviceCapability) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceCapability) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceCapability) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceCapability) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


