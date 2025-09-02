# IdentitySourceGroupMembershipsUpsertProfileInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupExternalId** | Pointer to **string** | The external ID of the group whose memberships need to be inserted or updated in Okta | [optional] 
**MemberExternalIds** | Pointer to **[]string** | Array of external IDs of member profiles that need to be inserted in this group in Okta | [optional] 

## Methods

### NewIdentitySourceGroupMembershipsUpsertProfileInner

`func NewIdentitySourceGroupMembershipsUpsertProfileInner() *IdentitySourceGroupMembershipsUpsertProfileInner`

NewIdentitySourceGroupMembershipsUpsertProfileInner instantiates a new IdentitySourceGroupMembershipsUpsertProfileInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentitySourceGroupMembershipsUpsertProfileInnerWithDefaults

`func NewIdentitySourceGroupMembershipsUpsertProfileInnerWithDefaults() *IdentitySourceGroupMembershipsUpsertProfileInner`

NewIdentitySourceGroupMembershipsUpsertProfileInnerWithDefaults instantiates a new IdentitySourceGroupMembershipsUpsertProfileInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupExternalId

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetGroupExternalId() string`

GetGroupExternalId returns the GroupExternalId field if non-nil, zero value otherwise.

### GetGroupExternalIdOk

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetGroupExternalIdOk() (*string, bool)`

GetGroupExternalIdOk returns a tuple with the GroupExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupExternalId

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) SetGroupExternalId(v string)`

SetGroupExternalId sets GroupExternalId field to given value.

### HasGroupExternalId

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) HasGroupExternalId() bool`

HasGroupExternalId returns a boolean if a field has been set.

### GetMemberExternalIds

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetMemberExternalIds() []string`

GetMemberExternalIds returns the MemberExternalIds field if non-nil, zero value otherwise.

### GetMemberExternalIdsOk

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) GetMemberExternalIdsOk() (*[]string, bool)`

GetMemberExternalIdsOk returns a tuple with the MemberExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberExternalIds

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) SetMemberExternalIds(v []string)`

SetMemberExternalIds sets MemberExternalIds field to given value.

### HasMemberExternalIds

`func (o *IdentitySourceGroupMembershipsUpsertProfileInner) HasMemberExternalIds() bool`

HasMemberExternalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


