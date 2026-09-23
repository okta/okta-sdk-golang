# BulkDeviceUpsertRequestBodyProfilesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | The external ID of the device that needs to be created or updated in Okta | [optional] 
**Profile** | Pointer to [**IdentitySourceDeviceProfileForUpsert**](IdentitySourceDeviceProfileForUpsert.md) |  | [optional] 

## Methods

### NewBulkDeviceUpsertRequestBodyProfilesInner

`func NewBulkDeviceUpsertRequestBodyProfilesInner() *BulkDeviceUpsertRequestBodyProfilesInner`

NewBulkDeviceUpsertRequestBodyProfilesInner instantiates a new BulkDeviceUpsertRequestBodyProfilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeviceUpsertRequestBodyProfilesInnerWithDefaults

`func NewBulkDeviceUpsertRequestBodyProfilesInnerWithDefaults() *BulkDeviceUpsertRequestBodyProfilesInner`

NewBulkDeviceUpsertRequestBodyProfilesInnerWithDefaults instantiates a new BulkDeviceUpsertRequestBodyProfilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProfile

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) GetProfile() IdentitySourceDeviceProfileForUpsert`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) GetProfileOk() (*IdentitySourceDeviceProfileForUpsert, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) SetProfile(v IdentitySourceDeviceProfileForUpsert)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *BulkDeviceUpsertRequestBodyProfilesInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


