# UserImportRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**UserImportRequestDataAction**](UserImportRequestDataAction.md) |  | [optional] 
**AppUser** | Pointer to [**UserImportRequestDataAppUser**](UserImportRequestDataAppUser.md) |  | [optional] 
**Context** | Pointer to [**UserImportRequestDataContext**](UserImportRequestDataContext.md) |  | [optional] 
**User** | Pointer to [**UserImportRequestDataUser**](UserImportRequestDataUser.md) |  | [optional] 

## Methods

### NewUserImportRequestData

`func NewUserImportRequestData() *UserImportRequestData`

NewUserImportRequestData instantiates a new UserImportRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportRequestDataWithDefaults

`func NewUserImportRequestDataWithDefaults() *UserImportRequestData`

NewUserImportRequestDataWithDefaults instantiates a new UserImportRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UserImportRequestData) GetAction() UserImportRequestDataAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UserImportRequestData) GetActionOk() (*UserImportRequestDataAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UserImportRequestData) SetAction(v UserImportRequestDataAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *UserImportRequestData) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAppUser

`func (o *UserImportRequestData) GetAppUser() UserImportRequestDataAppUser`

GetAppUser returns the AppUser field if non-nil, zero value otherwise.

### GetAppUserOk

`func (o *UserImportRequestData) GetAppUserOk() (*UserImportRequestDataAppUser, bool)`

GetAppUserOk returns a tuple with the AppUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppUser

`func (o *UserImportRequestData) SetAppUser(v UserImportRequestDataAppUser)`

SetAppUser sets AppUser field to given value.

### HasAppUser

`func (o *UserImportRequestData) HasAppUser() bool`

HasAppUser returns a boolean if a field has been set.

### GetContext

`func (o *UserImportRequestData) GetContext() UserImportRequestDataContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserImportRequestData) GetContextOk() (*UserImportRequestDataContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserImportRequestData) SetContext(v UserImportRequestDataContext)`

SetContext sets Context field to given value.

### HasContext

`func (o *UserImportRequestData) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetUser

`func (o *UserImportRequestData) GetUser() UserImportRequestDataUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserImportRequestData) GetUserOk() (*UserImportRequestDataUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserImportRequestData) SetUser(v UserImportRequestDataUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UserImportRequestData) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


