# ServiceAccountForUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the service account | [optional] 
**Name** | Pointer to **string** | The human-readable name for the service account | [optional] 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups who own the service account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users who own the service account | [optional] 

## Methods

### NewServiceAccountForUpdate

`func NewServiceAccountForUpdate() *ServiceAccountForUpdate`

NewServiceAccountForUpdate instantiates a new ServiceAccountForUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountForUpdateWithDefaults

`func NewServiceAccountForUpdateWithDefaults() *ServiceAccountForUpdate`

NewServiceAccountForUpdateWithDefaults instantiates a new ServiceAccountForUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceAccountForUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccountForUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccountForUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccountForUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccountForUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccountForUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccountForUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceAccountForUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerGroupIds

`func (o *ServiceAccountForUpdate) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *ServiceAccountForUpdate) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *ServiceAccountForUpdate) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *ServiceAccountForUpdate) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *ServiceAccountForUpdate) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *ServiceAccountForUpdate) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *ServiceAccountForUpdate) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *ServiceAccountForUpdate) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


