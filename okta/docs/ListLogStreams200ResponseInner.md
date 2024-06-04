# ListLogStreams200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **time.Time** | Timestamp when the Log Stream object was created | [readonly] 
**Id** | **string** | Unique identifier for the Log Stream | [readonly] 
**LastUpdated** | **time.Time** | Timestamp when the Log Stream object was last updated | [readonly] 
**Name** | **string** | Unique name for the Log Stream object | 
**Status** | **string** | Lifecycle status of the Log Stream object | [readonly] 
**Type** | **string** | Specifies the streaming provider used  Supported providers:   * &#x60;aws_eventbridge&#x60; ([AWS EventBridge](https://aws.amazon.com/eventbridge))   * &#x60;splunk_cloud_logstreaming&#x60; ([Splunk Cloud](https://www.splunk.com/en_us/software/splunk-cloud-platform.html))  Select the provider type to see provider-specific configurations in the &#x60;settings&#x60; property: | 
**Links** | [**LogStreamLinksSelfAndLifecycle**](LogStreamLinksSelfAndLifecycle.md) |  | 
**Settings** | [**LogStreamSettingsSplunk**](LogStreamSettingsSplunk.md) |  | 

## Methods

### NewListLogStreams200ResponseInner

`func NewListLogStreams200ResponseInner(created time.Time, id string, lastUpdated time.Time, name string, status string, type_ string, links LogStreamLinksSelfAndLifecycle, settings LogStreamSettingsSplunk, ) *ListLogStreams200ResponseInner`

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


### GetStatus

`func (o *ListLogStreams200ResponseInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLogStreams200ResponseInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLogStreams200ResponseInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ListLogStreams200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListLogStreams200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListLogStreams200ResponseInner) SetType(v string)`

SetType sets Type field to given value.


### GetLinks

`func (o *ListLogStreams200ResponseInner) GetLinks() LogStreamLinksSelfAndLifecycle`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListLogStreams200ResponseInner) GetLinksOk() (*LogStreamLinksSelfAndLifecycle, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListLogStreams200ResponseInner) SetLinks(v LogStreamLinksSelfAndLifecycle)`

SetLinks sets Links field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


