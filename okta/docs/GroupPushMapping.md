# GroupPushMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConfig** | Pointer to [**AppConfig**](AppConfig.md) |  | [optional] 
**Created** | Pointer to **time.Time** | Timestamp when the group push mapping was created | [optional] [readonly] 
**ErrorSummary** | Pointer to **string** | The error message summary if the latest push failed | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the group push mapping | [optional] [readonly] 
**LastPush** | Pointer to **time.Time** | Timestamp when the group push mapping was pushed | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the group push mapping was last updated | [optional] [readonly] 
**SourceGroupId** | Pointer to **string** | The ID of the source group for the group push mapping | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the group push mapping | [optional] 
**TargetGroupId** | Pointer to **string** | The ID of the target group for the group push mapping | [optional] [readonly] 
**Links** | Pointer to [**GroupPushMappingLinks**](GroupPushMappingLinks.md) |  | [optional] 

## Methods

### NewGroupPushMapping

`func NewGroupPushMapping() *GroupPushMapping`

NewGroupPushMapping instantiates a new GroupPushMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPushMappingWithDefaults

`func NewGroupPushMappingWithDefaults() *GroupPushMapping`

NewGroupPushMappingWithDefaults instantiates a new GroupPushMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConfig

`func (o *GroupPushMapping) GetAppConfig() AppConfig`

GetAppConfig returns the AppConfig field if non-nil, zero value otherwise.

### GetAppConfigOk

`func (o *GroupPushMapping) GetAppConfigOk() (*AppConfig, bool)`

GetAppConfigOk returns a tuple with the AppConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConfig

`func (o *GroupPushMapping) SetAppConfig(v AppConfig)`

SetAppConfig sets AppConfig field to given value.

### HasAppConfig

`func (o *GroupPushMapping) HasAppConfig() bool`

HasAppConfig returns a boolean if a field has been set.

### GetCreated

`func (o *GroupPushMapping) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupPushMapping) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupPushMapping) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GroupPushMapping) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetErrorSummary

`func (o *GroupPushMapping) GetErrorSummary() string`

GetErrorSummary returns the ErrorSummary field if non-nil, zero value otherwise.

### GetErrorSummaryOk

`func (o *GroupPushMapping) GetErrorSummaryOk() (*string, bool)`

GetErrorSummaryOk returns a tuple with the ErrorSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorSummary

`func (o *GroupPushMapping) SetErrorSummary(v string)`

SetErrorSummary sets ErrorSummary field to given value.

### HasErrorSummary

`func (o *GroupPushMapping) HasErrorSummary() bool`

HasErrorSummary returns a boolean if a field has been set.

### GetId

`func (o *GroupPushMapping) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupPushMapping) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupPushMapping) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupPushMapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastPush

`func (o *GroupPushMapping) GetLastPush() time.Time`

GetLastPush returns the LastPush field if non-nil, zero value otherwise.

### GetLastPushOk

`func (o *GroupPushMapping) GetLastPushOk() (*time.Time, bool)`

GetLastPushOk returns a tuple with the LastPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPush

`func (o *GroupPushMapping) SetLastPush(v time.Time)`

SetLastPush sets LastPush field to given value.

### HasLastPush

`func (o *GroupPushMapping) HasLastPush() bool`

HasLastPush returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GroupPushMapping) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GroupPushMapping) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GroupPushMapping) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GroupPushMapping) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSourceGroupId

`func (o *GroupPushMapping) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *GroupPushMapping) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *GroupPushMapping) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *GroupPushMapping) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### GetStatus

`func (o *GroupPushMapping) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupPushMapping) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupPushMapping) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GroupPushMapping) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *GroupPushMapping) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *GroupPushMapping) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *GroupPushMapping) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *GroupPushMapping) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetLinks

`func (o *GroupPushMapping) GetLinks() GroupPushMappingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupPushMapping) GetLinksOk() (*GroupPushMappingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupPushMapping) SetLinks(v GroupPushMappingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupPushMapping) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


