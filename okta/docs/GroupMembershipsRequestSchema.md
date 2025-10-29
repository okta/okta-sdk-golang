# GroupMembershipsRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberExternalIds** | Pointer to **[]string** | A list of app user external IDs to be inserted in this group in Okta | [optional] 

## Methods

### NewGroupMembershipsRequestSchema

`func NewGroupMembershipsRequestSchema() *GroupMembershipsRequestSchema`

NewGroupMembershipsRequestSchema instantiates a new GroupMembershipsRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMembershipsRequestSchemaWithDefaults

`func NewGroupMembershipsRequestSchemaWithDefaults() *GroupMembershipsRequestSchema`

NewGroupMembershipsRequestSchemaWithDefaults instantiates a new GroupMembershipsRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberExternalIds

`func (o *GroupMembershipsRequestSchema) GetMemberExternalIds() []string`

GetMemberExternalIds returns the MemberExternalIds field if non-nil, zero value otherwise.

### GetMemberExternalIdsOk

`func (o *GroupMembershipsRequestSchema) GetMemberExternalIdsOk() (*[]string, bool)`

GetMemberExternalIdsOk returns a tuple with the MemberExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberExternalIds

`func (o *GroupMembershipsRequestSchema) SetMemberExternalIds(v []string)`

SetMemberExternalIds sets MemberExternalIds field to given value.

### HasMemberExternalIds

`func (o *GroupMembershipsRequestSchema) HasMemberExternalIds() bool`

HasMemberExternalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


