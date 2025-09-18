# AddGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**OktaUserGroupProfile**](OktaUserGroupProfile.md) |  | [optional] 

## Methods

### NewAddGroupRequest

`func NewAddGroupRequest() *AddGroupRequest`

NewAddGroupRequest instantiates a new AddGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroupRequestWithDefaults

`func NewAddGroupRequestWithDefaults() *AddGroupRequest`

NewAddGroupRequestWithDefaults instantiates a new AddGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *AddGroupRequest) GetProfile() OktaUserGroupProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AddGroupRequest) GetProfileOk() (*OktaUserGroupProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AddGroupRequest) SetProfile(v OktaUserGroupProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AddGroupRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


