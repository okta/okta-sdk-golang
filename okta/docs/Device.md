# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the device was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the device | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the device was last updated | [optional] [readonly] 
**Profile** | Pointer to [**DeviceProfile**](DeviceProfile.md) |  | [optional] 
**ResourceAlternateId** | Pointer to **string** |  | [optional] [readonly] 
**ResourceDisplayName** | Pointer to [**DeviceDisplayName**](DeviceDisplayName.md) |  | [optional] 
**ResourceId** | Pointer to **string** | Alternate key for the &#x60;id&#x60; | [optional] [readonly] 
**ResourceType** | Pointer to **string** |  | [optional] [readonly] [default to "UDDevice"]
**Status** | Pointer to [**DeviceStatus**](DeviceStatus.md) |  | [optional] 
**Links** | Pointer to [**DeviceLinks**](DeviceLinks.md) |  | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Device) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Device) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Device) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Device) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Device) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Device) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Device) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Device) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *Device) GetProfile() DeviceProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Device) GetProfileOk() (*DeviceProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Device) SetProfile(v DeviceProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Device) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResourceAlternateId

`func (o *Device) GetResourceAlternateId() string`

GetResourceAlternateId returns the ResourceAlternateId field if non-nil, zero value otherwise.

### GetResourceAlternateIdOk

`func (o *Device) GetResourceAlternateIdOk() (*string, bool)`

GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAlternateId

`func (o *Device) SetResourceAlternateId(v string)`

SetResourceAlternateId sets ResourceAlternateId field to given value.

### HasResourceAlternateId

`func (o *Device) HasResourceAlternateId() bool`

HasResourceAlternateId returns a boolean if a field has been set.

### GetResourceDisplayName

`func (o *Device) GetResourceDisplayName() DeviceDisplayName`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *Device) GetResourceDisplayNameOk() (*DeviceDisplayName, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *Device) SetResourceDisplayName(v DeviceDisplayName)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *Device) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### GetResourceId

`func (o *Device) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Device) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Device) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Device) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *Device) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Device) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Device) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *Device) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *Device) GetStatus() DeviceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Device) GetStatusOk() (*DeviceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Device) SetStatus(v DeviceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Device) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *Device) GetLinks() DeviceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Device) GetLinksOk() (*DeviceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Device) SetLinks(v DeviceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Device) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


