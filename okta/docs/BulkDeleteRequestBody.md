# BulkDeleteRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | Pointer to **string** |  | [optional] 
**Profiles** | Pointer to [**[]IdentitySourceUserProfileForDelete**](IdentitySourceUserProfileForDelete.md) |  | [optional] 

## Methods

### NewBulkDeleteRequestBody

`func NewBulkDeleteRequestBody() *BulkDeleteRequestBody`

NewBulkDeleteRequestBody instantiates a new BulkDeleteRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteRequestBodyWithDefaults

`func NewBulkDeleteRequestBodyWithDefaults() *BulkDeleteRequestBody`

NewBulkDeleteRequestBodyWithDefaults instantiates a new BulkDeleteRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *BulkDeleteRequestBody) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *BulkDeleteRequestBody) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *BulkDeleteRequestBody) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *BulkDeleteRequestBody) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetProfiles

`func (o *BulkDeleteRequestBody) GetProfiles() []IdentitySourceUserProfileForDelete`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BulkDeleteRequestBody) GetProfilesOk() (*[]IdentitySourceUserProfileForDelete, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BulkDeleteRequestBody) SetProfiles(v []IdentitySourceUserProfileForDelete)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BulkDeleteRequestBody) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


