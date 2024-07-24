# AppAccountContainerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** | The application name | [optional] [readonly] 
**ContainerId** | **string** | The application ID associated with the privileged account | 
**DisplayName** | Pointer to **string** | Human-readable name of the container that owns the privileged resource | [optional] [readonly] 
**GlobalAppId** | Pointer to **string** | The application global ID | [optional] [readonly] 
**PasswordPushSupported** | Pointer to **bool** | Indicates if the application supports password push | [optional] [readonly] 
**ProvisioningEnabled** | Pointer to **bool** | Indicates if provisioning is enabled for this application | [optional] [readonly] 
**Links** | Pointer to [**AppLink**](AppLink.md) |  | [optional] 

## Methods

### NewAppAccountContainerDetails

`func NewAppAccountContainerDetails(containerId string, ) *AppAccountContainerDetails`

NewAppAccountContainerDetails instantiates a new AppAccountContainerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAccountContainerDetailsWithDefaults

`func NewAppAccountContainerDetailsWithDefaults() *AppAccountContainerDetails`

NewAppAccountContainerDetailsWithDefaults instantiates a new AppAccountContainerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *AppAccountContainerDetails) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *AppAccountContainerDetails) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *AppAccountContainerDetails) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *AppAccountContainerDetails) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetContainerId

`func (o *AppAccountContainerDetails) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *AppAccountContainerDetails) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *AppAccountContainerDetails) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.


### GetDisplayName

`func (o *AppAccountContainerDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AppAccountContainerDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AppAccountContainerDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AppAccountContainerDetails) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGlobalAppId

`func (o *AppAccountContainerDetails) GetGlobalAppId() string`

GetGlobalAppId returns the GlobalAppId field if non-nil, zero value otherwise.

### GetGlobalAppIdOk

`func (o *AppAccountContainerDetails) GetGlobalAppIdOk() (*string, bool)`

GetGlobalAppIdOk returns a tuple with the GlobalAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalAppId

`func (o *AppAccountContainerDetails) SetGlobalAppId(v string)`

SetGlobalAppId sets GlobalAppId field to given value.

### HasGlobalAppId

`func (o *AppAccountContainerDetails) HasGlobalAppId() bool`

HasGlobalAppId returns a boolean if a field has been set.

### GetPasswordPushSupported

`func (o *AppAccountContainerDetails) GetPasswordPushSupported() bool`

GetPasswordPushSupported returns the PasswordPushSupported field if non-nil, zero value otherwise.

### GetPasswordPushSupportedOk

`func (o *AppAccountContainerDetails) GetPasswordPushSupportedOk() (*bool, bool)`

GetPasswordPushSupportedOk returns a tuple with the PasswordPushSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPushSupported

`func (o *AppAccountContainerDetails) SetPasswordPushSupported(v bool)`

SetPasswordPushSupported sets PasswordPushSupported field to given value.

### HasPasswordPushSupported

`func (o *AppAccountContainerDetails) HasPasswordPushSupported() bool`

HasPasswordPushSupported returns a boolean if a field has been set.

### GetProvisioningEnabled

`func (o *AppAccountContainerDetails) GetProvisioningEnabled() bool`

GetProvisioningEnabled returns the ProvisioningEnabled field if non-nil, zero value otherwise.

### GetProvisioningEnabledOk

`func (o *AppAccountContainerDetails) GetProvisioningEnabledOk() (*bool, bool)`

GetProvisioningEnabledOk returns a tuple with the ProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningEnabled

`func (o *AppAccountContainerDetails) SetProvisioningEnabled(v bool)`

SetProvisioningEnabled sets ProvisioningEnabled field to given value.

### HasProvisioningEnabled

`func (o *AppAccountContainerDetails) HasProvisioningEnabled() bool`

HasProvisioningEnabled returns a boolean if a field has been set.

### GetLinks

`func (o *AppAccountContainerDetails) GetLinks() AppLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppAccountContainerDetails) GetLinksOk() (*AppLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppAccountContainerDetails) SetLinks(v AppLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppAccountContainerDetails) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


