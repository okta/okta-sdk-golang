# DeviceIntegrations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name of the device integration | [optional] 
**Id** | Pointer to **string** | The ID of the device integration | [optional] [readonly] 
**Metadata** | Pointer to [**DeviceIntegrationsMetadata**](DeviceIntegrationsMetadata.md) |  | [optional] 
**Name** | Pointer to **string** | The namespace of the device integration | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | The status of the device integration | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewDeviceIntegrations

`func NewDeviceIntegrations() *DeviceIntegrations`

NewDeviceIntegrations instantiates a new DeviceIntegrations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceIntegrationsWithDefaults

`func NewDeviceIntegrationsWithDefaults() *DeviceIntegrations`

NewDeviceIntegrationsWithDefaults instantiates a new DeviceIntegrations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DeviceIntegrations) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DeviceIntegrations) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DeviceIntegrations) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DeviceIntegrations) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *DeviceIntegrations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceIntegrations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceIntegrations) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceIntegrations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DeviceIntegrations) GetMetadata() DeviceIntegrationsMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceIntegrations) GetMetadataOk() (*DeviceIntegrationsMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceIntegrations) SetMetadata(v DeviceIntegrationsMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceIntegrations) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *DeviceIntegrations) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceIntegrations) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceIntegrations) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceIntegrations) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *DeviceIntegrations) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DeviceIntegrations) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DeviceIntegrations) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *DeviceIntegrations) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceIntegrations) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceIntegrations) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceIntegrations) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceIntegrations) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *DeviceIntegrations) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceIntegrations) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceIntegrations) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceIntegrations) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


