# SubmissionCapabilityEnhanced

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capability** | **string** | API Service capability identifier | 
**SupportedProtocols** | **[]string** | List of supported protocols | 
**Actions** | Pointer to [**[]SubmissionAction**](SubmissionAction.md) | Configuration for the Actions protocol. This parameter is required when &#x60;supportedProtocols&#x60; is set to &#x60;ACTIONS&#x60;. | [optional] 

## Methods

### NewSubmissionCapabilityEnhanced

`func NewSubmissionCapabilityEnhanced(capability string, supportedProtocols []string, ) *SubmissionCapabilityEnhanced`

NewSubmissionCapabilityEnhanced instantiates a new SubmissionCapabilityEnhanced object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmissionCapabilityEnhancedWithDefaults

`func NewSubmissionCapabilityEnhancedWithDefaults() *SubmissionCapabilityEnhanced`

NewSubmissionCapabilityEnhancedWithDefaults instantiates a new SubmissionCapabilityEnhanced object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapability

`func (o *SubmissionCapabilityEnhanced) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *SubmissionCapabilityEnhanced) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *SubmissionCapabilityEnhanced) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetSupportedProtocols

`func (o *SubmissionCapabilityEnhanced) GetSupportedProtocols() []string`

GetSupportedProtocols returns the SupportedProtocols field if non-nil, zero value otherwise.

### GetSupportedProtocolsOk

`func (o *SubmissionCapabilityEnhanced) GetSupportedProtocolsOk() (*[]string, bool)`

GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocols

`func (o *SubmissionCapabilityEnhanced) SetSupportedProtocols(v []string)`

SetSupportedProtocols sets SupportedProtocols field to given value.


### GetActions

`func (o *SubmissionCapabilityEnhanced) GetActions() []SubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *SubmissionCapabilityEnhanced) GetActionsOk() (*[]SubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *SubmissionCapabilityEnhanced) SetActions(v []SubmissionAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *SubmissionCapabilityEnhanced) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


