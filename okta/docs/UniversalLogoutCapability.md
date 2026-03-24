# UniversalLogoutCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]SubmissionAction**](SubmissionAction.md) | Configuration for the Actions protocol. This parameter is required when &#x60;supportedProtocols&#x60; is set to &#x60;ACTIONS&#x60;. | [optional] 
**Capability** | **string** | Universal Logout capability identifier | 
**SupportedProtocols** | **[]string** | List of supported logout protocols | 

## Methods

### NewUniversalLogoutCapability

`func NewUniversalLogoutCapability(capability string, supportedProtocols []string, ) *UniversalLogoutCapability`

NewUniversalLogoutCapability instantiates a new UniversalLogoutCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniversalLogoutCapabilityWithDefaults

`func NewUniversalLogoutCapabilityWithDefaults() *UniversalLogoutCapability`

NewUniversalLogoutCapabilityWithDefaults instantiates a new UniversalLogoutCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *UniversalLogoutCapability) GetActions() []SubmissionAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *UniversalLogoutCapability) GetActionsOk() (*[]SubmissionAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *UniversalLogoutCapability) SetActions(v []SubmissionAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *UniversalLogoutCapability) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetCapability

`func (o *UniversalLogoutCapability) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *UniversalLogoutCapability) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *UniversalLogoutCapability) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetSupportedProtocols

`func (o *UniversalLogoutCapability) GetSupportedProtocols() []string`

GetSupportedProtocols returns the SupportedProtocols field if non-nil, zero value otherwise.

### GetSupportedProtocolsOk

`func (o *UniversalLogoutCapability) GetSupportedProtocolsOk() (*[]string, bool)`

GetSupportedProtocolsOk returns a tuple with the SupportedProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedProtocols

`func (o *UniversalLogoutCapability) SetSupportedProtocols(v []string)`

SetSupportedProtocols sets SupportedProtocols field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


