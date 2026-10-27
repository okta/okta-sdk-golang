# DeviceCapabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | [**[]DeviceCapability**](DeviceCapability.md) | One entry per capability. An entry is omitted entirely if your org&#39;s plan doesn&#39;t include that capability, or if the device&#39;s platform doesn&#39;t support it. | 

## Methods

### NewDeviceCapabilitiesResponse

`func NewDeviceCapabilitiesResponse(capabilities []DeviceCapability, ) *DeviceCapabilitiesResponse`

NewDeviceCapabilitiesResponse instantiates a new DeviceCapabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCapabilitiesResponseWithDefaults

`func NewDeviceCapabilitiesResponseWithDefaults() *DeviceCapabilitiesResponse`

NewDeviceCapabilitiesResponseWithDefaults instantiates a new DeviceCapabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *DeviceCapabilitiesResponse) GetCapabilities() []DeviceCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DeviceCapabilitiesResponse) GetCapabilitiesOk() (*[]DeviceCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DeviceCapabilitiesResponse) SetCapabilities(v []DeviceCapability)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


