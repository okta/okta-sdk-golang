# BulkDeviceUpsertRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profiles** | Pointer to [**[]BulkDeviceUpsertRequestBodyProfilesInner**](BulkDeviceUpsertRequestBodyProfilesInner.md) | Array of device profiles to be uploaded | [optional] 

## Methods

### NewBulkDeviceUpsertRequestBody

`func NewBulkDeviceUpsertRequestBody() *BulkDeviceUpsertRequestBody`

NewBulkDeviceUpsertRequestBody instantiates a new BulkDeviceUpsertRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeviceUpsertRequestBodyWithDefaults

`func NewBulkDeviceUpsertRequestBodyWithDefaults() *BulkDeviceUpsertRequestBody`

NewBulkDeviceUpsertRequestBodyWithDefaults instantiates a new BulkDeviceUpsertRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfiles

`func (o *BulkDeviceUpsertRequestBody) GetProfiles() []BulkDeviceUpsertRequestBodyProfilesInner`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *BulkDeviceUpsertRequestBody) GetProfilesOk() (*[]BulkDeviceUpsertRequestBodyProfilesInner, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *BulkDeviceUpsertRequestBody) SetProfiles(v []BulkDeviceUpsertRequestBodyProfilesInner)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *BulkDeviceUpsertRequestBody) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


