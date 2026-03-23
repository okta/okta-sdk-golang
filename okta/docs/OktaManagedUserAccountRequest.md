# OktaManagedUserAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Okta managed user account | [optional] 
**Name** | **string** | The user-defined name for the Okta managed user account | 
**OktaUserId** | **string** | The ID of the Okta user to manage as a service account. This must be an existing user in your Okta org. | 
**OwnerGroupIds** | Pointer to **[]string** | A list of IDs of the Okta groups who own the Okta managed user account | [optional] 
**OwnerUserIds** | Pointer to **[]string** | A list of IDs of the Okta users who own the Okta managed user account | [optional] 

## Methods

### NewOktaManagedUserAccountRequest

`func NewOktaManagedUserAccountRequest(name string, oktaUserId string, ) *OktaManagedUserAccountRequest`

NewOktaManagedUserAccountRequest instantiates a new OktaManagedUserAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaManagedUserAccountRequestWithDefaults

`func NewOktaManagedUserAccountRequestWithDefaults() *OktaManagedUserAccountRequest`

NewOktaManagedUserAccountRequestWithDefaults instantiates a new OktaManagedUserAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OktaManagedUserAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OktaManagedUserAccountRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OktaManagedUserAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OktaManagedUserAccountRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OktaManagedUserAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OktaManagedUserAccountRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OktaManagedUserAccountRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOktaUserId

`func (o *OktaManagedUserAccountRequest) GetOktaUserId() string`

GetOktaUserId returns the OktaUserId field if non-nil, zero value otherwise.

### GetOktaUserIdOk

`func (o *OktaManagedUserAccountRequest) GetOktaUserIdOk() (*string, bool)`

GetOktaUserIdOk returns a tuple with the OktaUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaUserId

`func (o *OktaManagedUserAccountRequest) SetOktaUserId(v string)`

SetOktaUserId sets OktaUserId field to given value.


### GetOwnerGroupIds

`func (o *OktaManagedUserAccountRequest) GetOwnerGroupIds() []string`

GetOwnerGroupIds returns the OwnerGroupIds field if non-nil, zero value otherwise.

### GetOwnerGroupIdsOk

`func (o *OktaManagedUserAccountRequest) GetOwnerGroupIdsOk() (*[]string, bool)`

GetOwnerGroupIdsOk returns a tuple with the OwnerGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerGroupIds

`func (o *OktaManagedUserAccountRequest) SetOwnerGroupIds(v []string)`

SetOwnerGroupIds sets OwnerGroupIds field to given value.

### HasOwnerGroupIds

`func (o *OktaManagedUserAccountRequest) HasOwnerGroupIds() bool`

HasOwnerGroupIds returns a boolean if a field has been set.

### GetOwnerUserIds

`func (o *OktaManagedUserAccountRequest) GetOwnerUserIds() []string`

GetOwnerUserIds returns the OwnerUserIds field if non-nil, zero value otherwise.

### GetOwnerUserIdsOk

`func (o *OktaManagedUserAccountRequest) GetOwnerUserIdsOk() (*[]string, bool)`

GetOwnerUserIdsOk returns a tuple with the OwnerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserIds

`func (o *OktaManagedUserAccountRequest) SetOwnerUserIds(v []string)`

SetOwnerUserIds sets OwnerUserIds field to given value.

### HasOwnerUserIds

`func (o *OktaManagedUserAccountRequest) HasOwnerUserIds() bool`

HasOwnerUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


