# IdentitySourceGroupMembershipsDeleteProfileInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupExternalId** | Pointer to **string** | The external ID of the group whose memberships need to be deleted in Okta | [optional] 
**MemberExternalIds** | Pointer to **[]string** | Array of external IDs of member profiles that need to be inserted in this group in Okta | [optional] 

## Methods

### NewIdentitySourceGroupMembershipsDeleteProfileInner

`func NewIdentitySourceGroupMembershipsDeleteProfileInner() *IdentitySourceGroupMembershipsDeleteProfileInner`

NewIdentitySourceGroupMembershipsDeleteProfileInner instantiates a new IdentitySourceGroupMembershipsDeleteProfileInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourceGroupMembershipsDeleteProfileInnerWithDefaults

`func NewIdentitySourceGroupMembershipsDeleteProfileInnerWithDefaults() *IdentitySourceGroupMembershipsDeleteProfileInner`

NewIdentitySourceGroupMembershipsDeleteProfileInnerWithDefaults instantiates a new IdentitySourceGroupMembershipsDeleteProfileInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupExternalId

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) GetGroupExternalId() string`

GetGroupExternalId returns the GroupExternalId field if non-nil, zero value otherwise.

### GetGroupExternalIdOk

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) GetGroupExternalIdOk() (*string, bool)`

GetGroupExternalIdOk returns a tuple with the GroupExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupExternalId

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) SetGroupExternalId(v string)`

SetGroupExternalId sets GroupExternalId field to given value.

### HasGroupExternalId

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) HasGroupExternalId() bool`

HasGroupExternalId returns a boolean if a field has been set.

### GetMemberExternalIds

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) GetMemberExternalIds() []string`

GetMemberExternalIds returns the MemberExternalIds field if non-nil, zero value otherwise.

### GetMemberExternalIdsOk

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) GetMemberExternalIdsOk() (*[]string, bool)`

GetMemberExternalIdsOk returns a tuple with the MemberExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberExternalIds

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) SetMemberExternalIds(v []string)`

SetMemberExternalIds sets MemberExternalIds field to given value.

### HasMemberExternalIds

`func (o *IdentitySourceGroupMembershipsDeleteProfileInner) HasMemberExternalIds() bool`

HasMemberExternalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


