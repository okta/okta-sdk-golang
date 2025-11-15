# UserResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The timestamp when the user was created in the identity source | [optional] [readonly] 
**ExternalId** | Pointer to **string** | The external ID of the user in the identity source | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the user in the identity source | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The timestamp when the user was last updated in the identity source | [optional] [readonly] 
**Profile** | Pointer to [**IdentitySourceUserProfileForUpsert**](IdentitySourceUserProfileForUpsert.md) |  | [optional] 

## Methods

### NewUserResponseSchema

`func NewUserResponseSchema() *UserResponseSchema`

NewUserResponseSchema instantiates a new UserResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResponseSchemaWithDefaults

`func NewUserResponseSchemaWithDefaults() *UserResponseSchema`

NewUserResponseSchemaWithDefaults instantiates a new UserResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserResponseSchema) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserResponseSchema) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserResponseSchema) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserResponseSchema) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExternalId

`func (o *UserResponseSchema) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UserResponseSchema) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UserResponseSchema) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UserResponseSchema) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *UserResponseSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResponseSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResponseSchema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserResponseSchema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserResponseSchema) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserResponseSchema) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserResponseSchema) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserResponseSchema) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetProfile

`func (o *UserResponseSchema) GetProfile() IdentitySourceUserProfileForUpsert`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserResponseSchema) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserResponseSchema) SetProfile(v IdentitySourceUserProfileForUpsert)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserResponseSchema) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


