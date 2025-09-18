# BulkGroupDeleteRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIds** | Pointer to **[]string** | Array of external IDs of groups that need to be deleted in Okta | [optional] 

## Methods

### NewBulkGroupDeleteRequestBody

`func NewBulkGroupDeleteRequestBody() *BulkGroupDeleteRequestBody`

NewBulkGroupDeleteRequestBody instantiates a new BulkGroupDeleteRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGroupDeleteRequestBodyWithDefaults

`func NewBulkGroupDeleteRequestBodyWithDefaults() *BulkGroupDeleteRequestBody`

NewBulkGroupDeleteRequestBodyWithDefaults instantiates a new BulkGroupDeleteRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIds

`func (o *BulkGroupDeleteRequestBody) GetExternalIds() []string`

GetExternalIds returns the ExternalIds field if non-nil, zero value otherwise.

### GetExternalIdsOk

`func (o *BulkGroupDeleteRequestBody) GetExternalIdsOk() (*[]string, bool)`

GetExternalIdsOk returns a tuple with the ExternalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIds

`func (o *BulkGroupDeleteRequestBody) SetExternalIds(v []string)`

SetExternalIds sets ExternalIds field to given value.

### HasExternalIds

`func (o *BulkGroupDeleteRequestBody) HasExternalIds() bool`

HasExternalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


