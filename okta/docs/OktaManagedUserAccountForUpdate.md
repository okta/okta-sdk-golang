# OktaManagedUserAccountForUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Okta managed user account | [optional] 
**Name** | Pointer to **string** | The user-defined name for the Okta managed user account | [optional] 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups who own the Okta managed user account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users who own the Okta managed user account | [optional] 

## Methods

### NewOktaManagedUserAccountForUpdate

`func NewOktaManagedUserAccountForUpdate() *OktaManagedUserAccountForUpdate`

NewOktaManagedUserAccountForUpdate instantiates a new OktaManagedUserAccountForUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaManagedUserAccountForUpdateWithDefaults

`func NewOktaManagedUserAccountForUpdateWithDefaults() *OktaManagedUserAccountForUpdate`

NewOktaManagedUserAccountForUpdateWithDefaults instantiates a new OktaManagedUserAccountForUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OktaManagedUserAccountForUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OktaManagedUserAccountForUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OktaManagedUserAccountForUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OktaManagedUserAccountForUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OktaManagedUserAccountForUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OktaManagedUserAccountForUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OktaManagedUserAccountForUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OktaManagedUserAccountForUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerGroupIds

`func (o *OktaManagedUserAccountForUpdate) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *OktaManagedUserAccountForUpdate) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *OktaManagedUserAccountForUpdate) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *OktaManagedUserAccountForUpdate) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *OktaManagedUserAccountForUpdate) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *OktaManagedUserAccountForUpdate) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *OktaManagedUserAccountForUpdate) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *OktaManagedUserAccountForUpdate) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


