# BulkGroupMembershipsUpsertRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memberships** | Pointer to [**[]IdentitySourceGroupMembershipsUpsertProfileInner**](IdentitySourceGroupMembershipsUpsertProfileInner.md) | Array of group memberships that need to be inserted or updated in Okta | [optional] 

## Methods

### NewBulkGroupMembershipsUpsertRequestBody

`func NewBulkGroupMembershipsUpsertRequestBody() *BulkGroupMembershipsUpsertRequestBody`

NewBulkGroupMembershipsUpsertRequestBody instantiates a new BulkGroupMembershipsUpsertRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGroupMembershipsUpsertRequestBodyWithDefaults

`func NewBulkGroupMembershipsUpsertRequestBodyWithDefaults() *BulkGroupMembershipsUpsertRequestBody`

NewBulkGroupMembershipsUpsertRequestBodyWithDefaults instantiates a new BulkGroupMembershipsUpsertRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberships

`func (o *BulkGroupMembershipsUpsertRequestBody) GetMemberships() []IdentitySourceGroupMembershipsUpsertProfileInner`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *BulkGroupMembershipsUpsertRequestBody) GetMembershipsOk() (*[]IdentitySourceGroupMembershipsUpsertProfileInner, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *BulkGroupMembershipsUpsertRequestBody) SetMemberships(v []IdentitySourceGroupMembershipsUpsertProfileInner)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *BulkGroupMembershipsUpsertRequestBody) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


