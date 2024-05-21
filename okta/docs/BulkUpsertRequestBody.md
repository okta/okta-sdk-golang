# BulkUpsertRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** |  | [optional] 
**Profiles** | Pointer to [**[]IdentitySourceUserProfileForUpsert**](IdentitySourceUserProfileForUpsert.md) |  | [optional] 

## Methods

### NewBulkUpsertRequestBody

`func NewBulkUpsertRequestBody() *BulkUpsertRequestBody`

NewBulkUpsertRequestBody instantiates a new BulkUpsertRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpsertRequestBodyWithDefaults

`func NewBulkUpsertRequestBodyWithDefaults() *BulkUpsertRequestBody`

NewBulkUpsertRequestBodyWithDefaults instantiates a new BulkUpsertRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *BulkUpsertRequestBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BulkUpsertRequestBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BulkUpsertRequestBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BulkUpsertRequestBody) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetProfiles

`func (o *BulkUpsertRequestBody) GetProfiles() []IdentitySourceUserProfileForUpsert`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BulkUpsertRequestBody) GetProfilesOk() (*[]IdentitySourceUserProfileForUpsert, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BulkUpsertRequestBody) SetProfiles(v []IdentitySourceUserProfileForUpsert)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BulkUpsertRequestBody) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


