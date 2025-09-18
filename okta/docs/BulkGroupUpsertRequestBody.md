# BulkGroupUpsertRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]BulkGroupUpsertRequestBodyProfilesInner**](BulkGroupUpsertRequestBodyProfilesInner.md) | Array of group profiles that needs to be inserted or updated in Okta | [optional] 

## Methods

### NewBulkGroupUpsertRequestBody

`func NewBulkGroupUpsertRequestBody() *BulkGroupUpsertRequestBody`

NewBulkGroupUpsertRequestBody instantiates a new BulkGroupUpsertRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGroupUpsertRequestBodyWithDefaults

`func NewBulkGroupUpsertRequestBodyWithDefaults() *BulkGroupUpsertRequestBody`

NewBulkGroupUpsertRequestBodyWithDefaults instantiates a new BulkGroupUpsertRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *BulkGroupUpsertRequestBody) GetProfiles() []BulkGroupUpsertRequestBodyProfilesInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BulkGroupUpsertRequestBody) GetProfilesOk() (*[]BulkGroupUpsertRequestBodyProfilesInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BulkGroupUpsertRequestBody) SetProfiles(v []BulkGroupUpsertRequestBodyProfilesInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BulkGroupUpsertRequestBody) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


