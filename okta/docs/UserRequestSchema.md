# UserRequestSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the user in the identity source | [optional] 
**Profile** | Pointer to [**IdentitySourceUserProfileForUpsert**](IdentitySourceUserProfileForUpsert.md) |  | [optional] 

## Methods

### NewUserRequestSchema

`func NewUserRequestSchema() *UserRequestSchema`

NewUserRequestSchema instantiates a new UserRequestSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRequestSchemaWithDefaults

`func NewUserRequestSchemaWithDefaults() *UserRequestSchema`

NewUserRequestSchemaWithDefaults instantiates a new UserRequestSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *UserRequestSchema) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UserRequestSchema) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UserRequestSchema) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UserRequestSchema) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProfile

`func (o *UserRequestSchema) GetProfile() IdentitySourceUserProfileForUpsert`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserRequestSchema) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserRequestSchema) SetProfile(v IdentitySourceUserProfileForUpsert)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserRequestSchema) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


