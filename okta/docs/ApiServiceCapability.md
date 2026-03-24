# ApiServiceCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | API Service capability identifier | 
**SupportedProtocols** | **[]string** | List of supported protocols | 

## Methods

### NewApiServiceCapability

`func NewApiServiceCapability(capability string, supportedProtocols []string, ) *ApiServiceCapability`

NewApiServiceCapability instantiates a new ApiServiceCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServiceCapabilityWithDefaults

`func NewApiServiceCapabilityWithDefaults() *ApiServiceCapability`

NewApiServiceCapabilityWithDefaults instantiates a new ApiServiceCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *ApiServiceCapability) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *ApiServiceCapability) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *ApiServiceCapability) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetSupportedProtocols

`func (o *ApiServiceCapability) GetSupportedProtocols() []string`

GetSupportedProtocols returns the SupportedProtocols field if non-nil, zero value otherwise.

### GetSupportedProtocolsOk

`func (o *ApiServiceCapability) GetSupportedProtocolsOk() (*[]string, bool)`

GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocols

`func (o *ApiServiceCapability) SetSupportedProtocols(v []string)`

SetSupportedProtocols sets SupportedProtocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


