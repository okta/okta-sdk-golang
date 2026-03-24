# OSAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp when the OS account was created | [readonly] 
**DeviceId** | **string** | Unique identifier of the device this OS account belongs to | [readonly] 
**Id** | **string** | Unique identifier for the OS account | [readonly] 
**LastUpdated** | **time.Time** | Timestamp when the OS account was last updated | [readonly] 
**Platform** | **string** | OS platform for OS accounts (desktop platforms only) | 
**Links** | [**OSAccountLinks**](OSAccountLinks.md) |  | 

## Methods

### NewOSAccount

`func NewOSAccount(created time.Time, deviceId string, id string, lastUpdated time.Time, platform string, links OSAccountLinks, ) *OSAccount`

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


