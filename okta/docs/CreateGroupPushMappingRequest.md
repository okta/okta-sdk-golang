# CreateGroupPushMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConfig** | Pointer to [**AppConfig**](AppConfig.md) |  | [optional] 
**SourceGroupId** | **string** | The ID of the source group for the group push mapping | 
**Status** | Pointer to **string** | The status of the group push mapping | [optional] [default to "ACTIVE"]
**TargetGroupId** | Pointer to **string** | The ID of the existing target group for the group push mapping. This is used to link to an existing group. Required if &#x60;targetGroupName&#x60; is not provided. | [optional] 
**TargetGroupName** | Pointer to **string** | The name of the target group for the group push mapping. This is used when creating a new downstream group. If the group already exists, it links to the existing group. Required if &#x60;targetGroupId&#x60; is not provided. | [optional] 

## Methods

### NewCreateGroupPushMappingRequest

`func NewCreateGroupPushMappingRequest(sourceGroupId string, ) *CreateGroupPushMappingRequest`

NewCreateGroupPushMappingRequest instantiates a new CreateGroupPushMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGroupPushMappingRequestWithDefaults

`func NewCreateGroupPushMappingRequestWithDefaults() *CreateGroupPushMappingRequest`

NewCreateGroupPushMappingRequestWithDefaults instantiates a new CreateGroupPushMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConfig

`func (o *CreateGroupPushMappingRequest) GetAppConfig() AppConfig`

GetAppConfig returns the AppConfig field if non-nil, zero value otherwise.

### GetAppConfigOk

`func (o *CreateGroupPushMappingRequest) GetAppConfigOk() (*AppConfig, bool)`

GetAppConfigOk returns a tuple with the AppConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConfig

`func (o *CreateGroupPushMappingRequest) SetAppConfig(v AppConfig)`

SetAppConfig sets AppConfig field to given value.

### HasAppConfig

`func (o *CreateGroupPushMappingRequest) HasAppConfig() bool`

HasAppConfig returns a boolean if a field has been set.

### GetSourceGroupId

`func (o *CreateGroupPushMappingRequest) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *CreateGroupPushMappingRequest) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *CreateGroupPushMappingRequest) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.


### GetStatus

`func (o *CreateGroupPushMappingRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateGroupPushMappingRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateGroupPushMappingRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateGroupPushMappingRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetGroupId

`func (o *CreateGroupPushMappingRequest) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *CreateGroupPushMappingRequest) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *CreateGroupPushMappingRequest) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *CreateGroupPushMappingRequest) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### GetTargetGroupName

`func (o *CreateGroupPushMappingRequest) GetTargetGroupName() string`

GetTargetGroupName returns the TargetGroupName field if non-nil, zero value otherwise.

### GetTargetGroupNameOk

`func (o *CreateGroupPushMappingRequest) GetTargetGroupNameOk() (*string, bool)`

GetTargetGroupNameOk returns a tuple with the TargetGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupName

`func (o *CreateGroupPushMappingRequest) SetTargetGroupName(v string)`

SetTargetGroupName sets TargetGroupName field to given value.

### HasTargetGroupName

`func (o *CreateGroupPushMappingRequest) HasTargetGroupName() bool`

HasTargetGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


