# DeviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the device was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the device | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the device record was last updated. Updates occur when Okta collects and saves device signals during authentication, and when the lifecycle state of the device changes. | [optional] [readonly] 
**Profile** | Pointer to [**DeviceProfile**](DeviceProfile.md) |  | [optional] 
**ResourceAlternateId** | Pointer to **string** |  | [optional] [readonly] 
**ResourceDisplayName** | Pointer to [**DeviceDisplayName**](DeviceDisplayName.md) |  | [optional] 
**ResourceId** | Pointer to **string** | Alternate key for the &#x60;id&#x60; | [optional] [readonly] 
**ResourceType** | Pointer to **string** |  | [optional] [readonly] [default to "UDDevice"]
**Status** | Pointer to **string** | The state object of the device | [optional] 
**Links** | Pointer to [**LinksSelfAndFullUsersLifecycle**](LinksSelfAndFullUsersLifecycle.md) |  | [optional] 
**Embedded** | Pointer to [**DeviceListAllOfEmbedded**](DeviceListAllOfEmbedded.md) |  | [optional] 

## Methods

### NewDeviceList

`func NewDeviceList() *DeviceList`

NewDeviceList instantiates a new DeviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceListWithDefaults

`func NewDeviceListWithDefaults() *DeviceList`

NewDeviceListWithDefaults instantiates a new DeviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DeviceList) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceList) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceList) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceList) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *DeviceList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeviceList) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeviceList) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeviceList) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeviceList) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *DeviceList) GetProfile() DeviceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *DeviceList) GetProfileOk() (*DeviceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *DeviceList) SetProfile(v DeviceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *DeviceList) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResourceAlternateId

`func (o *DeviceList) GetResourceAlternateId() string`

GetResourceAlternateId returns the ResourceAlternateId field if non-nil, zero value otherwise.

### GetResourceAlternateIdOk

`func (o *DeviceList) GetResourceAlternateIdOk() (*string, bool)`

GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAlternateId

`func (o *DeviceList) SetResourceAlternateId(v string)`

SetResourceAlternateId sets ResourceAlternateId field to given value.

### HasResourceAlternateId

`func (o *DeviceList) HasResourceAlternateId() bool`

HasResourceAlternateId returns a boolean if a field has been set.

### GetResourceDisplayName

`func (o *DeviceList) GetResourceDisplayName() DeviceDisplayName`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *DeviceList) GetResourceDisplayNameOk() (*DeviceDisplayName, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *DeviceList) SetResourceDisplayName(v DeviceDisplayName)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *DeviceList) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### GetResourceId

`func (o *DeviceList) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *DeviceList) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *DeviceList) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *DeviceList) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *DeviceList) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *DeviceList) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *DeviceList) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *DeviceList) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceList) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceList) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceList) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *DeviceList) GetLinks() LinksSelfAndFullUsersLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeviceList) GetLinksOk() (*LinksSelfAndFullUsersLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeviceList) SetLinks(v LinksSelfAndFullUsersLifecycle)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeviceList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *DeviceList) GetEmbedded() DeviceListAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DeviceList) GetEmbeddedOk() (*DeviceListAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DeviceList) SetEmbedded(v DeviceListAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DeviceList) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


