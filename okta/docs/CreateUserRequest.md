# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**UserCredentials**](UserCredentials.md) |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**Profile** | [**UserProfile**](UserProfile.md) |  | 
**RealmId** | Pointer to **string** | &lt;div class&#x3D;\&quot;x-lifecycle-container\&quot;&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/div&gt;The ID of the Realm in which the user is residing | [optional] 
**Type** | Pointer to [**CreateUserRequestType**](CreateUserRequestType.md) |  | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(profile UserProfile, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *CreateUserRequest) GetCredentials() UserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateUserRequest) GetCredentialsOk() (*UserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateUserRequest) SetCredentials(v UserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateUserRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGroupIds

`func (o *CreateUserRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *CreateUserRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *CreateUserRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *CreateUserRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetProfile

`func (o *CreateUserRequest) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateUserRequest) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateUserRequest) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.


### GetRealmId

`func (o *CreateUserRequest) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *CreateUserRequest) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *CreateUserRequest) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *CreateUserRequest) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetType

`func (o *CreateUserRequest) GetType() CreateUserRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserRequest) GetTypeOk() (*CreateUserRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserRequest) SetType(v CreateUserRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateUserRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


