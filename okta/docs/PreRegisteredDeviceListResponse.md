# PreRegisteredDeviceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the device was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the device | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the device record was last updated | [optional] [readonly] 
**Profile** | Pointer to [**DeviceProfile**](DeviceProfile.md) |  | [optional] 
**RegistrationGrants** | Pointer to [**[]RegistrationGrant**](RegistrationGrant.md) | Registration grants associated with this device | [optional] 
**ResourceId** | Pointer to **string** | Alternate key for the &#x60;id&#x60; | [optional] [readonly] 
**ResourceType** | Pointer to **string** |  | [optional] [readonly] [default to "UDDevice"]
**Status** | Pointer to **string** | The state object of the device | [optional] 
**Links** | Pointer to [**LinksSelfAndLifecycle**](LinksSelfAndLifecycle.md) |  | [optional] 

## Methods

### NewPreRegisteredDeviceListResponse

`func NewPreRegisteredDeviceListResponse() *PreRegisteredDeviceListResponse`

NewPreRegisteredDeviceListResponse instantiates a new PreRegisteredDeviceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreRegisteredDeviceListResponseWithDefaults

`func NewPreRegisteredDeviceListResponseWithDefaults() *PreRegisteredDeviceListResponse`

NewPreRegisteredDeviceListResponseWithDefaults instantiates a new PreRegisteredDeviceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PreRegisteredDeviceListResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PreRegisteredDeviceListResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PreRegisteredDeviceListResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PreRegisteredDeviceListResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *PreRegisteredDeviceListResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreRegisteredDeviceListResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreRegisteredDeviceListResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PreRegisteredDeviceListResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PreRegisteredDeviceListResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PreRegisteredDeviceListResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PreRegisteredDeviceListResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PreRegisteredDeviceListResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *PreRegisteredDeviceListResponse) GetProfile() DeviceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PreRegisteredDeviceListResponse) GetProfileOk() (*DeviceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PreRegisteredDeviceListResponse) SetProfile(v DeviceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *PreRegisteredDeviceListResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRegistrationGrants

`func (o *PreRegisteredDeviceListResponse) GetRegistrationGrants() []RegistrationGrant`

GetRegistrationGrants returns the RegistrationGrants field if non-nil, zero value otherwise.

### GetRegistrationGrantsOk

`func (o *PreRegisteredDeviceListResponse) GetRegistrationGrantsOk() (*[]RegistrationGrant, bool)`

GetRegistrationGrantsOk returns a tuple with the RegistrationGrants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationGrants

`func (o *PreRegisteredDeviceListResponse) SetRegistrationGrants(v []RegistrationGrant)`

SetRegistrationGrants sets RegistrationGrants field to given value.

### HasRegistrationGrants

`func (o *PreRegisteredDeviceListResponse) HasRegistrationGrants() bool`

HasRegistrationGrants returns a boolean if a field has been set.

### GetResourceId

`func (o *PreRegisteredDeviceListResponse) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PreRegisteredDeviceListResponse) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PreRegisteredDeviceListResponse) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *PreRegisteredDeviceListResponse) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *PreRegisteredDeviceListResponse) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PreRegisteredDeviceListResponse) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PreRegisteredDeviceListResponse) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *PreRegisteredDeviceListResponse) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *PreRegisteredDeviceListResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PreRegisteredDeviceListResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PreRegisteredDeviceListResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PreRegisteredDeviceListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *PreRegisteredDeviceListResponse) GetLinks() LinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PreRegisteredDeviceListResponse) GetLinksOk() (*LinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PreRegisteredDeviceListResponse) SetLinks(v LinksSelfAndLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PreRegisteredDeviceListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


