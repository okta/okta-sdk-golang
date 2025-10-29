# BulkGroupUpsertRequestBodyProfilesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the group that needs to be created or updated in Okta | [optional] 
**Profile** | Pointer to [**IdentitySourceGroupProfileForUpsert**](IdentitySourceGroupProfileForUpsert.md) |  | [optional] 

## Methods

### NewBulkGroupUpsertRequestBodyProfilesInner

`func NewBulkGroupUpsertRequestBodyProfilesInner() *BulkGroupUpsertRequestBodyProfilesInner`

NewBulkGroupUpsertRequestBodyProfilesInner instantiates a new BulkGroupUpsertRequestBodyProfilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGroupUpsertRequestBodyProfilesInnerWithDefaults

`func NewBulkGroupUpsertRequestBodyProfilesInnerWithDefaults() *BulkGroupUpsertRequestBodyProfilesInner`

NewBulkGroupUpsertRequestBodyProfilesInnerWithDefaults instantiates a new BulkGroupUpsertRequestBodyProfilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *BulkGroupUpsertRequestBodyProfilesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BulkGroupUpsertRequestBodyProfilesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BulkGroupUpsertRequestBodyProfilesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BulkGroupUpsertRequestBodyProfilesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProfile

`func (o *BulkGroupUpsertRequestBodyProfilesInner) GetProfile() IdentitySourceGroupProfileForUpsert`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BulkGroupUpsertRequestBodyProfilesInner) GetProfileOk() (*IdentitySourceGroupProfileForUpsert, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BulkGroupUpsertRequestBodyProfilesInner) SetProfile(v IdentitySourceGroupProfileForUpsert)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BulkGroupUpsertRequestBodyProfilesInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


