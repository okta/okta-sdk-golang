# AppServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerGlobalName** | Pointer to **string** | The key name of the app in the Okta Integration Network (OIN) | [optional] [readonly] 
**ContainerInstanceName** | Pointer to **string** | The app instance label | [optional] [readonly] 
**ContainerOrn** | **string** | The [ORN](/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the relevant resource.  Use the specific app ORN format (&#x60;orn:{partition}:idp:{yourOrgId}:apps:{appType}:{appId}&#x60;) to identify an Okta app instance in your org. | 
**Created** | Pointer to **time.Time** | Timestamp when the app service account was created | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the app service account | [optional] 
**Id** | Pointer to **string** | The UUID of the app service account | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the app service account was last updated | [optional] [readonly] 
**Name** | **string** | The user-defined name for the app service account | 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups who own the app service account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users who own the app service account | [optional] 
**Password** | Pointer to **string** | The app service account password. Required for apps that don&#39;t have provisioning enabled or don&#39;t support password synchronization. | [optional] 
**Status** | Pointer to **string** | Describes the current status of an app service account | [optional] [readonly] 
**StatusDetail** | Pointer to **string** | Describes the detailed status of an app service account | [optional] [readonly] 
**Username** | **string** | The username that serves as the direct link to your managed app account. Ensure that this value precisely matches the identifier of the target app account. | 

## Methods

### NewAppServiceAccount

`func NewAppServiceAccount(containerOrn string, name string, username string, ) *AppServiceAccount`

NewAppServiceAccount instantiates a new AppServiceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAccountWithDefaults

`func NewAppServiceAccountWithDefaults() *AppServiceAccount`

NewAppServiceAccountWithDefaults instantiates a new AppServiceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerGlobalName

`func (o *AppServiceAccount) GetContainerGlobalName() string`

GetContainerGlobalName returns the ContainerGlobalName field if non-nil, zero value otherwise.

### GetContainerGlobalNameOk

`func (o *AppServiceAccount) GetContainerGlobalNameOk() (*string, bool)`

GetContainerGlobalNameOk returns a tuple with the ContainerGlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerGlobalName

`func (o *AppServiceAccount) SetContainerGlobalName(v string)`

SetContainerGlobalName sets ContainerGlobalName field to given value.

### HasContainerGlobalName

`func (o *AppServiceAccount) HasContainerGlobalName() bool`

HasContainerGlobalName returns a boolean if a field has been set.

### GetContainerInstanceName

`func (o *AppServiceAccount) GetContainerInstanceName() string`

GetContainerInstanceName returns the ContainerInstanceName field if non-nil, zero value otherwise.

### GetContainerInstanceNameOk

`func (o *AppServiceAccount) GetContainerInstanceNameOk() (*string, bool)`

GetContainerInstanceNameOk returns a tuple with the ContainerInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerInstanceName

`func (o *AppServiceAccount) SetContainerInstanceName(v string)`

SetContainerInstanceName sets ContainerInstanceName field to given value.

### HasContainerInstanceName

`func (o *AppServiceAccount) HasContainerInstanceName() bool`

HasContainerInstanceName returns a boolean if a field has been set.

### GetContainerOrn

`func (o *AppServiceAccount) GetContainerOrn() string`

GetContainerOrn returns the ContainerOrn field if non-nil, zero value otherwise.

### GetContainerOrnOk

`func (o *AppServiceAccount) GetContainerOrnOk() (*string, bool)`

GetContainerOrnOk returns a tuple with the ContainerOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerOrn

`func (o *AppServiceAccount) SetContainerOrn(v string)`

SetContainerOrn sets ContainerOrn field to given value.


### GetCreated

`func (o *AppServiceAccount) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppServiceAccount) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppServiceAccount) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppServiceAccount) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *AppServiceAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServiceAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServiceAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServiceAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AppServiceAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppServiceAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppServiceAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppServiceAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AppServiceAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AppServiceAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AppServiceAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AppServiceAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *AppServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServiceAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServiceAccount) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerGroupIds

`func (o *AppServiceAccount) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *AppServiceAccount) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *AppServiceAccount) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *AppServiceAccount) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *AppServiceAccount) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *AppServiceAccount) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *AppServiceAccount) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *AppServiceAccount) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.

### GetPassword

`func (o *AppServiceAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AppServiceAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AppServiceAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AppServiceAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetStatus

`func (o *AppServiceAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppServiceAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppServiceAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppServiceAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *AppServiceAccount) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *AppServiceAccount) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *AppServiceAccount) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *AppServiceAccount) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetUsername

`func (o *AppServiceAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AppServiceAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AppServiceAccount) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


