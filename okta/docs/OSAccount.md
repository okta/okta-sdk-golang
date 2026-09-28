# OSAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp when the OS account was created | [readonly] 
**DeviceId** | **string** | Unique identifier of the device this OS account belongs to | [readonly] 
**Id** | **string** | Unique identifier for the OS account | [readonly] 
**LastSeenAt** | Pointer to **NullableTime** | Timestamp when the OS account was last seen | [optional] 
**LastUpdated** | **time.Time** | Timestamp when the OS account was last updated | [readonly] 
**Platform** | **string** | OS platform for OS accounts (desktop platforms only) | 
**ResourceAlternateId** | Pointer to **NullableString** |  | [optional] [readonly] 
**ResourceDisplayName** | Pointer to [**OSAccountDisplayName**](OSAccountDisplayName.md) |  | [optional] 
**ResourceId** | Pointer to **string** | Alternate key for the &#x60;id&#x60; | [optional] [readonly] 
**ResourceType** | Pointer to **string** |  | [optional] [readonly] [default to "DOSAccount"]
**Status** | **string** | Status of the OS account | 
**Embedded** | Pointer to [**OSAccountEmbedded**](OSAccountEmbedded.md) |  | [optional] 
**Links** | [**OSAccountLinks**](OSAccountLinks.md) |  | 

## Methods

### NewOSAccount

`func NewOSAccount(created time.Time, deviceId string, id string, lastUpdated time.Time, platform string, status string, links OSAccountLinks, ) *OSAccount`

NewOSAccount instantiates a new OSAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOSAccountWithDefaults

`func NewOSAccountWithDefaults() *OSAccount`

NewOSAccountWithDefaults instantiates a new OSAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *OSAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OSAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OSAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetDeviceId

`func (o *OSAccount) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *OSAccount) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *OSAccount) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetId

`func (o *OSAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OSAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OSAccount) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeenAt

`func (o *OSAccount) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *OSAccount) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *OSAccount) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *OSAccount) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### SetLastSeenAtNil

`func (o *OSAccount) SetLastSeenAtNil(b bool)`

 SetLastSeenAtNil sets the value for LastSeenAt to be an explicit nil

### UnsetLastSeenAt
`func (o *OSAccount) UnsetLastSeenAt()`

UnsetLastSeenAt ensures that no value is present for LastSeenAt, not even an explicit nil
### GetLastUpdated

`func (o *OSAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OSAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OSAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetPlatform

`func (o *OSAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *OSAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *OSAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetResourceAlternateId

`func (o *OSAccount) GetResourceAlternateId() string`

GetResourceAlternateId returns the ResourceAlternateId field if non-nil, zero value otherwise.

### GetResourceAlternateIdOk

`func (o *OSAccount) GetResourceAlternateIdOk() (*string, bool)`

GetResourceAlternateIdOk returns a tuple with the ResourceAlternateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAlternateId

`func (o *OSAccount) SetResourceAlternateId(v string)`

SetResourceAlternateId sets ResourceAlternateId field to given value.

### HasResourceAlternateId

`func (o *OSAccount) HasResourceAlternateId() bool`

HasResourceAlternateId returns a boolean if a field has been set.

### SetResourceAlternateIdNil

`func (o *OSAccount) SetResourceAlternateIdNil(b bool)`

 SetResourceAlternateIdNil sets the value for ResourceAlternateId to be an explicit nil

### UnsetResourceAlternateId
`func (o *OSAccount) UnsetResourceAlternateId()`

UnsetResourceAlternateId ensures that no value is present for ResourceAlternateId, not even an explicit nil
### GetResourceDisplayName

`func (o *OSAccount) GetResourceDisplayName() OSAccountDisplayName`

GetResourceDisplayName returns the ResourceDisplayName field if non-nil, zero value otherwise.

### GetResourceDisplayNameOk

`func (o *OSAccount) GetResourceDisplayNameOk() (*OSAccountDisplayName, bool)`

GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDisplayName

`func (o *OSAccount) SetResourceDisplayName(v OSAccountDisplayName)`

SetResourceDisplayName sets ResourceDisplayName field to given value.

### HasResourceDisplayName

`func (o *OSAccount) HasResourceDisplayName() bool`

HasResourceDisplayName returns a boolean if a field has been set.

### GetResourceId

`func (o *OSAccount) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *OSAccount) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *OSAccount) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *OSAccount) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *OSAccount) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *OSAccount) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *OSAccount) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *OSAccount) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetStatus

`func (o *OSAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OSAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OSAccount) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetEmbedded

`func (o *OSAccount) GetEmbedded() OSAccountEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *OSAccount) GetEmbeddedOk() (*OSAccountEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *OSAccount) SetEmbedded(v OSAccountEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *OSAccount) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *OSAccount) GetLinks() OSAccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OSAccount) GetLinksOk() (*OSAccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OSAccount) SetLinks(v OSAccountLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


