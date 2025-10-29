# BaseContextUserLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**HrefObject**](HrefObject.md) | URL to retrieve the individual user&#39;s group memberships | [optional] 
**Factors** | Pointer to [**HrefObject**](HrefObject.md) | URL to retrieve individual user&#39;s factor enrollments | [optional] 

## Methods

### NewBaseContextUserLinks

`func NewBaseContextUserLinks() *BaseContextUserLinks`

NewBaseContextUserLinks instantiates a new BaseContextUserLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseContextUserLinksWithDefaults

`func NewBaseContextUserLinksWithDefaults() *BaseContextUserLinks`

NewBaseContextUserLinksWithDefaults instantiates a new BaseContextUserLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *BaseContextUserLinks) GetGroups() HrefObject`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BaseContextUserLinks) GetGroupsOk() (*HrefObject, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BaseContextUserLinks) SetGroups(v HrefObject)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BaseContextUserLinks) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetFactors

`func (o *BaseContextUserLinks) GetFactors() HrefObject`

GetFactors returns the Factors field if non-nil, zero value otherwise.

### GetFactorsOk

`func (o *BaseContextUserLinks) GetFactorsOk() (*HrefObject, bool)`

GetFactorsOk returns a tuple with the Factors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactors

`func (o *BaseContextUserLinks) SetFactors(v HrefObject)`

SetFactors sets Factors field to given value.

### HasFactors

`func (o *BaseContextUserLinks) HasFactors() bool`

HasFactors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


