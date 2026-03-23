# ProvisioningCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]SubmissionAction**](SubmissionAction.md) | Configuration for the Actions protocol. This parameter is required when &#x60;supportedProtocols&#x60; is set to &#x60;ACTIONS&#x60;. | [optional] 
**Capability** | **string** | Provisioning capability identifier | 
**SupportedProtocols** | **[]string** | List of supported provisioning protocols | 

## Methods

### NewProvisioningCapability

`func NewProvisioningCapability(capability string, supportedProtocols []string, ) *ProvisioningCapability`

NewProvisioningCapability instantiates a new ProvisioningCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningCapabilityWithDefaults

`func NewProvisioningCapabilityWithDefaults() *ProvisioningCapability`

NewProvisioningCapabilityWithDefaults instantiates a new ProvisioningCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ProvisioningCapability) GetActions() []SubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ProvisioningCapability) GetActionsOk() (*[]SubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ProvisioningCapability) SetActions(v []SubmissionAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ProvisioningCapability) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCapability

`func (o *ProvisioningCapability) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *ProvisioningCapability) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *ProvisioningCapability) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetSupportedProtocols

`func (o *ProvisioningCapability) GetSupportedProtocols() []string`

GetSupportedProtocols returns the SupportedProtocols field if non-nil, zero value otherwise.

### GetSupportedProtocolsOk

`func (o *ProvisioningCapability) GetSupportedProtocolsOk() (*[]string, bool)`

GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocols

`func (o *ProvisioningCapability) SetSupportedProtocols(v []string)`

SetSupportedProtocols sets SupportedProtocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


