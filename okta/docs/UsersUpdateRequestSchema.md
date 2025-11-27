# UsersUpdateRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**IdentitySourceUserProfileForUpsert**](IdentitySourceUserProfileForUpsert.md) |  | [optional] 

## Methods

### NewUsersUpdateRequestSchema

`func NewUsersUpdateRequestSchema() *UsersUpdateRequestSchema`

NewUsersUpdateRequestSchema instantiates a new UsersUpdateRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUpdateRequestSchemaWithDefaults

`func NewUsersUpdateRequestSchemaWithDefaults() *UsersUpdateRequestSchema`

NewUsersUpdateRequestSchemaWithDefaults instantiates a new UsersUpdateRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *UsersUpdateRequestSchema) GetProfile() IdentitySourceUserProfileForUpsert`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UsersUpdateRequestSchema) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UsersUpdateRequestSchema) SetProfile(v IdentitySourceUserProfileForUpsert)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UsersUpdateRequestSchema) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


