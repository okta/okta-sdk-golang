# AppServiceAccountForUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the app service account | [optional] 
**Name** | Pointer to **string** | The user-defined name for the app service account | [optional] 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups who own the app service account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users who own the app service account | [optional] 

## Methods

### NewAppServiceAccountForUpdate

`func NewAppServiceAccountForUpdate() *AppServiceAccountForUpdate`

NewAppServiceAccountForUpdate instantiates a new AppServiceAccountForUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAccountForUpdateWithDefaults

`func NewAppServiceAccountForUpdateWithDefaults() *AppServiceAccountForUpdate`

NewAppServiceAccountForUpdateWithDefaults instantiates a new AppServiceAccountForUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AppServiceAccountForUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceAccountForUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceAccountForUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceAccountForUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AppServiceAccountForUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceAccountForUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceAccountForUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppServiceAccountForUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerGroupIds

`func (o *AppServiceAccountForUpdate) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *AppServiceAccountForUpdate) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *AppServiceAccountForUpdate) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *AppServiceAccountForUpdate) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *AppServiceAccountForUpdate) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *AppServiceAccountForUpdate) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *AppServiceAccountForUpdate) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *AppServiceAccountForUpdate) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


