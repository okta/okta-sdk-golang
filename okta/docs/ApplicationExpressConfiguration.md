# ApplicationExpressConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledCapabilities** | Pointer to **[]string** | Capabilities currently enabled for the app | [optional] 
**SupportedCapabilities** | Pointer to **[]string** | Capabilities supported by the app | [optional] 

## Methods

### NewApplicationExpressConfiguration

`func NewApplicationExpressConfiguration() *ApplicationExpressConfiguration`

NewApplicationExpressConfiguration instantiates a new ApplicationExpressConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationExpressConfigurationWithDefaults

`func NewApplicationExpressConfigurationWithDefaults() *ApplicationExpressConfiguration`

NewApplicationExpressConfigurationWithDefaults instantiates a new ApplicationExpressConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledCapabilities

`func (o *ApplicationExpressConfiguration) GetEnabledCapabilities() []string`

GetEnabledCapabilities returns the EnabledCapabilities field if non-nil, zero value otherwise.

### GetEnabledCapabilitiesOk

`func (o *ApplicationExpressConfiguration) GetEnabledCapabilitiesOk() (*[]string, bool)`

GetEnabledCapabilitiesOk returns a tuple with the EnabledCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledCapabilities

`func (o *ApplicationExpressConfiguration) SetEnabledCapabilities(v []string)`

SetEnabledCapabilities sets EnabledCapabilities field to given value.

### HasEnabledCapabilities

`func (o *ApplicationExpressConfiguration) HasEnabledCapabilities() bool`

HasEnabledCapabilities returns a boolean if a field has been set.

### GetSupportedCapabilities

`func (o *ApplicationExpressConfiguration) GetSupportedCapabilities() []string`

GetSupportedCapabilities returns the SupportedCapabilities field if non-nil, zero value otherwise.

### GetSupportedCapabilitiesOk

`func (o *ApplicationExpressConfiguration) GetSupportedCapabilitiesOk() (*[]string, bool)`

GetSupportedCapabilitiesOk returns a tuple with the SupportedCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCapabilities

`func (o *ApplicationExpressConfiguration) SetSupportedCapabilities(v []string)`

SetSupportedCapabilities sets SupportedCapabilities field to given value.

### HasSupportedCapabilities

`func (o *ApplicationExpressConfiguration) HasSupportedCapabilities() bool`

HasSupportedCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


