# LogStream

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

## Methods

### NewLogStream

`func NewLogStream() *LogStream`

NewLogStream instantiates a new LogStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogStreamWithDefaults

`func NewLogStreamWithDefaults() *LogStream`

NewLogStreamWithDefaults instantiates a new LogStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *LogStream) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LogStream) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LogStream) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LogStream) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *LogStream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogStream) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogStream) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogStream) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *LogStream) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LogStream) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LogStream) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *LogStream) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *LogStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LogStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LogStream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LogStream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *LogStream) GetStatus() LifecycleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LogStream) GetStatusOk() (*LifecycleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LogStream) SetStatus(v LifecycleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LogStream) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *LogStream) GetType() LogStreamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogStream) GetTypeOk() (*LogStreamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogStream) SetType(v LogStreamType)`

SetType sets Type field to given value.

### HasType

`func (o *LogStream) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinks

`func (o *LogStream) GetLinks() LogStreamLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *LogStream) GetLinksOk() (*LogStreamLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *LogStream) SetLinks(v LogStreamLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *LogStream) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


