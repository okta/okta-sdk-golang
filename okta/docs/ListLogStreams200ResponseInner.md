# ListLogStreams200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Log Stream was created | [optional] [readonly] 
**Id** | Pointer to **string** | Unique key for the Log Stream | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Log Stream was last updated | [optional] [readonly] 
**Name** | Pointer to **string** | Unique name for the Log Stream | [optional] 
**Status** | Pointer to [**LifecycleStatus**](LifecycleStatus.md) |  | [optional] 
**Type** | Pointer to [**LogStreamType**](LogStreamType.md) |  | [optional] 
**Links** | Pointer to [**LogStreamLinks**](LogStreamLinks.md) |  | [optional] 
**Settings** | Pointer to [**LogStreamSettingsSplunk**](LogStreamSettingsSplunk.md) |  | [optional] 

## Methods

### NewListLogStreams200ResponseInner

`func NewListLogStreams200ResponseInner() *ListLogStreams200ResponseInner`

NewListLogStreams200ResponseInner instantiates a new ListLogStreams200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogStreams200ResponseInnerWithDefaults

`func NewListLogStreams200ResponseInnerWithDefaults() *ListLogStreams200ResponseInner`

NewListLogStreams200ResponseInnerWithDefaults instantiates a new ListLogStreams200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ListLogStreams200ResponseInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListLogStreams200ResponseInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListLogStreams200ResponseInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListLogStreams200ResponseInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *ListLogStreams200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLogStreams200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLogStreams200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListLogStreams200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLogStreams200ResponseInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLogStreams200ResponseInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLogStreams200ResponseInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLogStreams200ResponseInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ListLogStreams200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLogStreams200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLogStreams200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLogStreams200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ListLogStreams200ResponseInner) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLogStreams200ResponseInner) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLogStreams200ResponseInner) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLogStreams200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *ListLogStreams200ResponseInner) GetType() LogStreamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListLogStreams200ResponseInner) GetTypeOk() (*LogStreamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListLogStreams200ResponseInner) SetType(v LogStreamType)`

SetType sets Type field to given value.

### HasType

`func (o *ListLogStreams200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *ListLogStreams200ResponseInner) GetLinks() LogStreamLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListLogStreams200ResponseInner) GetLinksOk() (*LogStreamLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListLogStreams200ResponseInner) SetLinks(v LogStreamLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListLogStreams200ResponseInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetSettings

`func (o *ListLogStreams200ResponseInner) GetSettings() LogStreamSettingsSplunk`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ListLogStreams200ResponseInner) GetSettingsOk() (*LogStreamSettingsSplunk, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ListLogStreams200ResponseInner) SetSettings(v LogStreamSettingsSplunk)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ListLogStreams200ResponseInner) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


