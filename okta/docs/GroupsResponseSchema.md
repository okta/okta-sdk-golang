# GroupsResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the identity source group | [optional] 
**Id** | Pointer to **string** | The Okta group ID of the identity source group | [optional] [readonly] 
**Profile** | Pointer to [**GroupsResponseSchemaProfile**](GroupsResponseSchemaProfile.md) |  | [optional] 

## Methods

### NewGroupsResponseSchema

`func NewGroupsResponseSchema() *GroupsResponseSchema`

NewGroupsResponseSchema instantiates a new GroupsResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsResponseSchemaWithDefaults

`func NewGroupsResponseSchemaWithDefaults() *GroupsResponseSchema`

NewGroupsResponseSchemaWithDefaults instantiates a new GroupsResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *GroupsResponseSchema) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GroupsResponseSchema) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GroupsResponseSchema) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GroupsResponseSchema) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *GroupsResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupsResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupsResponseSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupsResponseSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfile

`func (o *GroupsResponseSchema) GetProfile() GroupsResponseSchemaProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GroupsResponseSchema) GetProfileOk() (*GroupsResponseSchemaProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GroupsResponseSchema) SetProfile(v GroupsResponseSchemaProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GroupsResponseSchema) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


