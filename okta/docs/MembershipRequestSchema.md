# MembershipRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberExternalId** | Pointer to **string** | The external ID of the user to be added as a member of the group in Okta | [optional] 

## Methods

### NewMembershipRequestSchema

`func NewMembershipRequestSchema() *MembershipRequestSchema`

NewMembershipRequestSchema instantiates a new MembershipRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMembershipRequestSchemaWithDefaults

`func NewMembershipRequestSchemaWithDefaults() *MembershipRequestSchema`

NewMembershipRequestSchemaWithDefaults instantiates a new MembershipRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberExternalId

`func (o *MembershipRequestSchema) GetMemberExternalId() string`

GetMemberExternalId returns the MemberExternalId field if non-nil, zero value otherwise.

### GetMemberExternalIdOk

`func (o *MembershipRequestSchema) GetMemberExternalIdOk() (*string, bool)`

GetMemberExternalIdOk returns a tuple with the MemberExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberExternalId

`func (o *MembershipRequestSchema) SetMemberExternalId(v string)`

SetMemberExternalId sets MemberExternalId field to given value.

### HasMemberExternalId

`func (o *MembershipRequestSchema) HasMemberExternalId() bool`

HasMemberExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


