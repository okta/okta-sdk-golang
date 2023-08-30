# PrincipalRateLimitEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** |  | [optional] [readonly] 
**CreatedDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**DefaultConcurrencyPercentage** | Pointer to **int32** |  | [optional] [readonly] 
**DefaultPercentage** | Pointer to **int32** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastUpdate** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**PrincipalId** | **string** |  | 
**PrincipalType** | [**PrincipalType**](PrincipalType.md) |  | 

## Methods

### NewPrincipalRateLimitEntity

`func NewPrincipalRateLimitEntity(principalId string, principalType PrincipalType, ) *PrincipalRateLimitEntity`

NewPrincipalRateLimitEntity instantiates a new PrincipalRateLimitEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalRateLimitEntityWithDefaults

`func NewPrincipalRateLimitEntityWithDefaults() *PrincipalRateLimitEntity`

NewPrincipalRateLimitEntityWithDefaults instantiates a new PrincipalRateLimitEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *PrincipalRateLimitEntity) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrincipalRateLimitEntity) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrincipalRateLimitEntity) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrincipalRateLimitEntity) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDate

`func (o *PrincipalRateLimitEntity) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *PrincipalRateLimitEntity) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *PrincipalRateLimitEntity) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *PrincipalRateLimitEntity) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetDefaultConcurrencyPercentage

`func (o *PrincipalRateLimitEntity) GetDefaultConcurrencyPercentage() int32`

GetDefaultConcurrencyPercentage returns the DefaultConcurrencyPercentage field if non-nil, zero value otherwise.

### GetDefaultConcurrencyPercentageOk

`func (o *PrincipalRateLimitEntity) GetDefaultConcurrencyPercentageOk() (*int32, bool)`

GetDefaultConcurrencyPercentageOk returns a tuple with the DefaultConcurrencyPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConcurrencyPercentage

`func (o *PrincipalRateLimitEntity) SetDefaultConcurrencyPercentage(v int32)`

SetDefaultConcurrencyPercentage sets DefaultConcurrencyPercentage field to given value.

### HasDefaultConcurrencyPercentage

`func (o *PrincipalRateLimitEntity) HasDefaultConcurrencyPercentage() bool`

HasDefaultConcurrencyPercentage returns a boolean if a field has been set.

### GetDefaultPercentage

`func (o *PrincipalRateLimitEntity) GetDefaultPercentage() int32`

GetDefaultPercentage returns the DefaultPercentage field if non-nil, zero value otherwise.

### GetDefaultPercentageOk

`func (o *PrincipalRateLimitEntity) GetDefaultPercentageOk() (*int32, bool)`

GetDefaultPercentageOk returns a tuple with the DefaultPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPercentage

`func (o *PrincipalRateLimitEntity) SetDefaultPercentage(v int32)`

SetDefaultPercentage sets DefaultPercentage field to given value.

### HasDefaultPercentage

`func (o *PrincipalRateLimitEntity) HasDefaultPercentage() bool`

HasDefaultPercentage returns a boolean if a field has been set.

### GetId

`func (o *PrincipalRateLimitEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalRateLimitEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalRateLimitEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrincipalRateLimitEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdate

`func (o *PrincipalRateLimitEntity) GetLastUpdate() time.Time`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *PrincipalRateLimitEntity) GetLastUpdateOk() (*time.Time, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *PrincipalRateLimitEntity) SetLastUpdate(v time.Time)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *PrincipalRateLimitEntity) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *PrincipalRateLimitEntity) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *PrincipalRateLimitEntity) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *PrincipalRateLimitEntity) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *PrincipalRateLimitEntity) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetOrgId

`func (o *PrincipalRateLimitEntity) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *PrincipalRateLimitEntity) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *PrincipalRateLimitEntity) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *PrincipalRateLimitEntity) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetPrincipalId

`func (o *PrincipalRateLimitEntity) GetPrincipalId() string`

GetPrincipalId returns the PrincipalId field if non-nil, zero value otherwise.

### GetPrincipalIdOk

`func (o *PrincipalRateLimitEntity) GetPrincipalIdOk() (*string, bool)`

GetPrincipalIdOk returns a tuple with the PrincipalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalId

`func (o *PrincipalRateLimitEntity) SetPrincipalId(v string)`

SetPrincipalId sets PrincipalId field to given value.


### GetPrincipalType

`func (o *PrincipalRateLimitEntity) GetPrincipalType() PrincipalType`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *PrincipalRateLimitEntity) GetPrincipalTypeOk() (*PrincipalType, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *PrincipalRateLimitEntity) SetPrincipalType(v PrincipalType)`

SetPrincipalType sets PrincipalType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


