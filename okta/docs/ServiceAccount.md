# ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | **string** | The type of service account | 
**Created** | Pointer to **time.Time** | Timestamp when the service account was created | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the service account | [optional] 
**Id** | Pointer to **string** | The UUID of the service account | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the service account was last updated | [optional] [readonly] 
**Name** | **string** | The user-defined name for the service account | 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups that own the service account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users that own the service account | [optional] 
**Status** | Pointer to **string** | Describes the current status of an app service account | [optional] [readonly] 
**StatusDetail** | Pointer to **string** | Describes the detailed status of an app service account | [optional] [readonly] 

## Methods

### NewServiceAccount

`func NewServiceAccount(accountType string, name string, ) *ServiceAccount`

NewServiceAccount instantiates a new ServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountWithDefaults

`func NewServiceAccountWithDefaults() *ServiceAccount`

NewServiceAccountWithDefaults instantiates a new ServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *ServiceAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ServiceAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ServiceAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetCreated

`func (o *ServiceAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServiceAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ServiceAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ServiceAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ServiceAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ServiceAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ServiceAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *ServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceAccount) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerGroupIds

`func (o *ServiceAccount) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *ServiceAccount) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *ServiceAccount) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *ServiceAccount) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *ServiceAccount) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *ServiceAccount) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *ServiceAccount) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *ServiceAccount) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *ServiceAccount) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *ServiceAccount) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *ServiceAccount) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *ServiceAccount) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


