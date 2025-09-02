# InlineHookCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to [**InlineHookChannel**](InlineHookChannel.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Date of the inline hook creation | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for the inline hook | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Date of the last inline hook update | [optional] [readonly] 
**Name** | Pointer to **string** | The display name of the inline hook | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | One of the inline hook types | [optional] 
**Version** | Pointer to **string** | Version of the inline hook type. The currently supported version is &#x60;1.0.0&#x60;. | [optional] [readonly] 
**Links** | Pointer to [**InlineHookLinksCreate**](InlineHookLinksCreate.md) |  | [optional] 

## Methods

### NewInlineHookCreateResponse

`func NewInlineHookCreateResponse() *InlineHookCreateResponse`

NewInlineHookCreateResponse instantiates a new InlineHookCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineHookCreateResponseWithDefaults

`func NewInlineHookCreateResponseWithDefaults() *InlineHookCreateResponse`

NewInlineHookCreateResponseWithDefaults instantiates a new InlineHookCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *InlineHookCreateResponse) GetChannel() InlineHookChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineHookCreateResponse) GetChannelOk() (*InlineHookChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineHookCreateResponse) SetChannel(v InlineHookChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineHookCreateResponse) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetCreated

`func (o *InlineHookCreateResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InlineHookCreateResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InlineHookCreateResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *InlineHookCreateResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *InlineHookCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineHookCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineHookCreateResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineHookCreateResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *InlineHookCreateResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InlineHookCreateResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InlineHookCreateResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *InlineHookCreateResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *InlineHookCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineHookCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineHookCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineHookCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *InlineHookCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineHookCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineHookCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineHookCreateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *InlineHookCreateResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineHookCreateResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineHookCreateResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InlineHookCreateResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *InlineHookCreateResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InlineHookCreateResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InlineHookCreateResponse) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InlineHookCreateResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLinks

`func (o *InlineHookCreateResponse) GetLinks() InlineHookLinksCreate`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InlineHookCreateResponse) GetLinksOk() (*InlineHookLinksCreate, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InlineHookCreateResponse) SetLinks(v InlineHookLinksCreate)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InlineHookCreateResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


