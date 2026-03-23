# SsoCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | SSO capability identifier | 
**SupportedProtocols** | **[]string** | List of supported SSO protocols | 

## Methods

### NewSsoCapability

`func NewSsoCapability(capability string, supportedProtocols []string, ) *SsoCapability`

NewSsoCapability instantiates a new SsoCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoCapabilityWithDefaults

`func NewSsoCapabilityWithDefaults() *SsoCapability`

NewSsoCapabilityWithDefaults instantiates a new SsoCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *SsoCapability) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *SsoCapability) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *SsoCapability) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetSupportedProtocols

`func (o *SsoCapability) GetSupportedProtocols() []string`

GetSupportedProtocols returns the SupportedProtocols field if non-nil, zero value otherwise.

### GetSupportedProtocolsOk

`func (o *SsoCapability) GetSupportedProtocolsOk() (*[]string, bool)`

GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocols

`func (o *SsoCapability) SetSupportedProtocols(v []string)`

SetSupportedProtocols sets SupportedProtocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


