# BulkUpsertRequestBodyProfilesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the entity that needs to be created or updated in Okta | [optional] 
**Profile** | Pointer to [**IdentitySourceUserProfileForUpsert**](IdentitySourceUserProfileForUpsert.md) |  | [optional] 

## Methods

### NewBulkUpsertRequestBodyProfilesInner

`func NewBulkUpsertRequestBodyProfilesInner() *BulkUpsertRequestBodyProfilesInner`

NewBulkUpsertRequestBodyProfilesInner instantiates a new BulkUpsertRequestBodyProfilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpsertRequestBodyProfilesInnerWithDefaults

`func NewBulkUpsertRequestBodyProfilesInnerWithDefaults() *BulkUpsertRequestBodyProfilesInner`

NewBulkUpsertRequestBodyProfilesInnerWithDefaults instantiates a new BulkUpsertRequestBodyProfilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *BulkUpsertRequestBodyProfilesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BulkUpsertRequestBodyProfilesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BulkUpsertRequestBodyProfilesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BulkUpsertRequestBodyProfilesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProfile

`func (o *BulkUpsertRequestBodyProfilesInner) GetProfile() IdentitySourceUserProfileForUpsert`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BulkUpsertRequestBodyProfilesInner) GetProfileOk() (*IdentitySourceUserProfileForUpsert, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BulkUpsertRequestBodyProfilesInner) SetProfile(v IdentitySourceUserProfileForUpsert)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BulkUpsertRequestBodyProfilesInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


