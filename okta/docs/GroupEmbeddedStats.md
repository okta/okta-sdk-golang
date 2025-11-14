# GroupEmbeddedStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsersCount** | Pointer to **int32** | Number of users in the group | [optional] 
**AppsCount** | Pointer to **int32** | Number of apps associated with the group | [optional] 
**GroupPushMappingsCount** | Pointer to **int32** | Number of group push mappings associated with the group | [optional] 
**HasAdminPrivlege** | Pointer to **bool** | Indicates if the group has admin privileges via a group-level role assignment | [optional] 

## Methods

### NewGroupEmbeddedStats

`func NewGroupEmbeddedStats() *GroupEmbeddedStats`

NewGroupEmbeddedStats instantiates a new GroupEmbeddedStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupEmbeddedStatsWithDefaults

`func NewGroupEmbeddedStatsWithDefaults() *GroupEmbeddedStats`

NewGroupEmbeddedStatsWithDefaults instantiates a new GroupEmbeddedStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsersCount

`func (o *GroupEmbeddedStats) GetUsersCount() int32`

GetUsersCount returns the UsersCount field if non-nil, zero value otherwise.

### GetUsersCountOk

`func (o *GroupEmbeddedStats) GetUsersCountOk() (*int32, bool)`

GetUsersCountOk returns a tuple with the UsersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersCount

`func (o *GroupEmbeddedStats) SetUsersCount(v int32)`

SetUsersCount sets UsersCount field to given value.

### HasUsersCount

`func (o *GroupEmbeddedStats) HasUsersCount() bool`

HasUsersCount returns a boolean if a field has been set.

### GetAppsCount

`func (o *GroupEmbeddedStats) GetAppsCount() int32`

GetAppsCount returns the AppsCount field if non-nil, zero value otherwise.

### GetAppsCountOk

`func (o *GroupEmbeddedStats) GetAppsCountOk() (*int32, bool)`

GetAppsCountOk returns a tuple with the AppsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsCount

`func (o *GroupEmbeddedStats) SetAppsCount(v int32)`

SetAppsCount sets AppsCount field to given value.

### HasAppsCount

`func (o *GroupEmbeddedStats) HasAppsCount() bool`

HasAppsCount returns a boolean if a field has been set.

### GetGroupPushMappingsCount

`func (o *GroupEmbeddedStats) GetGroupPushMappingsCount() int32`

GetGroupPushMappingsCount returns the GroupPushMappingsCount field if non-nil, zero value otherwise.

### GetGroupPushMappingsCountOk

`func (o *GroupEmbeddedStats) GetGroupPushMappingsCountOk() (*int32, bool)`

GetGroupPushMappingsCountOk returns a tuple with the GroupPushMappingsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPushMappingsCount

`func (o *GroupEmbeddedStats) SetGroupPushMappingsCount(v int32)`

SetGroupPushMappingsCount sets GroupPushMappingsCount field to given value.

### HasGroupPushMappingsCount

`func (o *GroupEmbeddedStats) HasGroupPushMappingsCount() bool`

HasGroupPushMappingsCount returns a boolean if a field has been set.

### GetHasAdminPrivlege

`func (o *GroupEmbeddedStats) GetHasAdminPrivlege() bool`

GetHasAdminPrivlege returns the HasAdminPrivlege field if non-nil, zero value otherwise.

### GetHasAdminPrivlegeOk

`func (o *GroupEmbeddedStats) GetHasAdminPrivlegeOk() (*bool, bool)`

GetHasAdminPrivlegeOk returns a tuple with the HasAdminPrivlege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAdminPrivlege

`func (o *GroupEmbeddedStats) SetHasAdminPrivlege(v bool)`

SetHasAdminPrivlege sets HasAdminPrivlege field to given value.

### HasHasAdminPrivlege

`func (o *GroupEmbeddedStats) HasHasAdminPrivlege() bool`

HasHasAdminPrivlege returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


