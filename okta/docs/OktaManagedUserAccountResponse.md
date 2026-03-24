# OktaManagedUserAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | Timestamp when the Okta managed user account was created | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the Okta managed user account | [optional] 
**Email** | **string** | The email address associated with the Okta user. This parameter is read-only, and it is derived from the Okta user profile. | [readonly] 
**Id** | **string** | The UUID of the Okta managed user account | [readonly] 
**LastUpdated** | Pointer to **time.Time** | Timestamp when the Okta managed user account was last updated | [optional] [readonly] 
**Name** | **string** | The user-defined name for the Okta managed user account | 
**OktaUserId** | **string** | The ID of the Okta user being managed as a service account | 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups who own the Okta managed user account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users who own the Okta managed user account | [optional] 
**Status** | Pointer to **string** | Describes the current status of a service account | [optional] [readonly] 
**StatusDetail** | Pointer to **string** | Describes the detailed status of a service account | [optional] [readonly] 
**Username** | **string** | The username associated with the Okta user. This parameter is read-only, and it is derived from the Okta user profile. | [readonly] 

## Methods

### NewOktaManagedUserAccountResponse

`func NewOktaManagedUserAccountResponse(email string, id string, name string, oktaUserId string, username string, ) *OktaManagedUserAccountResponse`

NewOktaManagedUserAccountResponse instantiates a new OktaManagedUserAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaManagedUserAccountResponseWithDefaults

`func NewOktaManagedUserAccountResponseWithDefaults() *OktaManagedUserAccountResponse`

NewOktaManagedUserAccountResponseWithDefaults instantiates a new OktaManagedUserAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *OktaManagedUserAccountResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *OktaManagedUserAccountResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *OktaManagedUserAccountResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *OktaManagedUserAccountResponse) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *OktaManagedUserAccountResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OktaManagedUserAccountResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OktaManagedUserAccountResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OktaManagedUserAccountResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *OktaManagedUserAccountResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OktaManagedUserAccountResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OktaManagedUserAccountResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *OktaManagedUserAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OktaManagedUserAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OktaManagedUserAccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdated

`func (o *OktaManagedUserAccountResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *OktaManagedUserAccountResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *OktaManagedUserAccountResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *OktaManagedUserAccountResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *OktaManagedUserAccountResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OktaManagedUserAccountResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OktaManagedUserAccountResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOktaUserId

`func (o *OktaManagedUserAccountResponse) GetOktaUserId() string`

GetOktaUserId returns the OktaUserId field if non-nil, zero value otherwise.

### GetOktaUserIdOk

`func (o *OktaManagedUserAccountResponse) GetOktaUserIdOk() (*string, bool)`

GetOktaUserIdOk returns a tuple with the OktaUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaUserId

`func (o *OktaManagedUserAccountResponse) SetOktaUserId(v string)`

SetOktaUserId sets OktaUserId field to given value.


### GetOwnerGroupIds

`func (o *OktaManagedUserAccountResponse) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *OktaManagedUserAccountResponse) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *OktaManagedUserAccountResponse) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *OktaManagedUserAccountResponse) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *OktaManagedUserAccountResponse) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *OktaManagedUserAccountResponse) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *OktaManagedUserAccountResponse) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *OktaManagedUserAccountResponse) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.

### GetStatus

`func (o *OktaManagedUserAccountResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OktaManagedUserAccountResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OktaManagedUserAccountResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OktaManagedUserAccountResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetail

`func (o *OktaManagedUserAccountResponse) GetStatusDetail() string`

GetStatusDetail returns the StatusDetail field if non-nil, zero value otherwise.

### GetStatusDetailOk

`func (o *OktaManagedUserAccountResponse) GetStatusDetailOk() (*string, bool)`

GetStatusDetailOk returns a tuple with the StatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetail

`func (o *OktaManagedUserAccountResponse) SetStatusDetail(v string)`

SetStatusDetail sets StatusDetail field to given value.

### HasStatusDetail

`func (o *OktaManagedUserAccountResponse) HasStatusDetail() bool`

HasStatusDetail returns a boolean if a field has been set.

### GetUsername

`func (o *OktaManagedUserAccountResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OktaManagedUserAccountResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OktaManagedUserAccountResponse) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


