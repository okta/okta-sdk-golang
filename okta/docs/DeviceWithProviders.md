# DeviceWithProviders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the device was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the device record. This identifier must be unique across all devices. If two or more physical devices report a non-unique device profile attribute (such as a truncated &#x60;profile.udid&#x60;), Okta may map them to the same device record (shared &#x60;id&#x60;), which can cause unexpected behavior. | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the device record was last updated. Updates occur when Okta collects and saves device signals during authentication, and when the lifecycle state of the device changes. | [optional] [readonly] 
**Profile** | Pointer to [**DeviceProfile**](DeviceProfile.md) |  | [optional] 
**RegistrationGrants** | Pointer to [**[]RegistrationGrant**](RegistrationGrant.md) | Registration grants associated with this device. Present only when the device was pre-registered. Omitted (not an empty array) if no grants exist for the device. | [optional] 
**ResourceAlternateId** | Pointer to **string** |  | [optional] [readonly] 
**ResourceDisplayName** | Pointer to [**DeviceDisplayName**](DeviceDisplayName.md) |  | [optional] 
**ResourceId** | Pointer to **string** | Alternate key for the &#x60;id&#x60; | [optional] [readonly] 
**ResourceType** | Pointer to **string** |  | [optional] [readonly] [default to "UDDevice"]
**Status** | Pointer to **string** | The state object of the device | [optional] 
**Links** | Pointer to [**LinksSelfAndFullUsersLifecycle**](LinksSelfAndFullUsersLifecycle.md) |  | [optional] 
**Providers** | Pointer to [**[]DeviceProvider**](DeviceProvider.md) | List of providers for the device when the &#x60;expand&#x3D;providers&#x60; query parameter is specified | [optional] [readonly] 

## Methods

### NewDeviceWithProviders

`func NewDeviceWithProviders() *DeviceWithProviders`

NewDeviceWithProviders instantiates a new DeviceWithProviders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithProvidersWithDefaults

`func NewDeviceWithProvidersWithDefaults() *DeviceWithProviders`

NewDeviceWithProvidersWithDefaults instantiates a new DeviceWithProviders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeviceWithProviders) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceWithProviders) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceWithProviders) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceWithProviders) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *DeviceWithProviders) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceWithProviders) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceWithProviders) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceWithProviders) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeviceWithProviders) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceWithProviders) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceWithProviders) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeviceWithProviders) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *DeviceWithProviders) GetProfile() DeviceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DeviceWithProviders) GetProfileOk() (*DeviceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DeviceWithProviders) SetProfile(v DeviceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *DeviceWithProviders) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRegistrationGrants

`func (o *DeviceWithProviders) GetRegistrationGrants() []RegistrationGrant`

GetRegistrationGrants returns the RegistrationGrants field if non-nil, zero value otherwise.

### GetRegistrationGrantsOk

`func (o *DeviceWithProviders) GetRegistrationGrantsOk() (*[]RegistrationGrant, bool)`

GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationGrants

`func (o *DeviceWithProviders) SetRegistrationGrants(v []RegistrationGrant)`

SetRegistrationGrants sets RegistrationGrants field to given value.

### HasRegistrationGrants

`func (o *DeviceWithProviders) HasRegistrationGrants() bool`

HasRegistrationGrants returns a boolean if a field has been set.

### GetResourceAlternateId

`func (o *DeviceWithProviders) GetResourceAlternateId() string`

GetResourceAlternateId returns the ResourceAlternateId field if non-nil, zero value otherwise.

### GetResourceAlternateIdOk

`func (o *DeviceWithProviders) GetResourceAlternateIdOk() (*string, bool)`

GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAlternateId

`func (o *DeviceWithProviders) SetResourceAlternateId(v string)`

SetResourceAlternateId sets ResourceAlternateId field to given value.

### HasResourceAlternateId

`func (o *DeviceWithProviders) HasResourceAlternateId() bool`

HasResourceAlternateId returns a boolean if a field has been set.

### GetResourceDisplayName

`func (o *DeviceWithProviders) GetResourceDisplayName() DeviceDisplayName`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *DeviceWithProviders) GetResourceDisplayNameOk() (*DeviceDisplayName, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *DeviceWithProviders) SetResourceDisplayName(v DeviceDisplayName)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *DeviceWithProviders) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### GetResourceId

`func (o *DeviceWithProviders) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DeviceWithProviders) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DeviceWithProviders) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DeviceWithProviders) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *DeviceWithProviders) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DeviceWithProviders) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DeviceWithProviders) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DeviceWithProviders) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceWithProviders) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceWithProviders) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceWithProviders) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceWithProviders) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *DeviceWithProviders) GetLinks() LinksSelfAndFullUsersLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceWithProviders) GetLinksOk() (*LinksSelfAndFullUsersLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceWithProviders) SetLinks(v LinksSelfAndFullUsersLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceWithProviders) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProviders

`func (o *DeviceWithProviders) GetProviders() []DeviceProvider`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *DeviceWithProviders) GetProvidersOk() (*[]DeviceProvider, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *DeviceWithProviders) SetProviders(v []DeviceProvider)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *DeviceWithProviders) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


