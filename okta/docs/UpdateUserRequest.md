# UpdateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**UserCredentials**](UserCredentials.md) |  | [optional] 
**Profile** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 
**RealmId** | Pointer to **string** | &lt;div class&#x3D;\&quot;x-lifecycle-container\&quot;&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/div&gt;The ID of the Realm in which the user is residing | [optional] 

## Methods

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestWithDefaults

`func NewUpdateUserRequestWithDefaults() *UpdateUserRequest`

NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *UpdateUserRequest) GetCredentials() UserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UpdateUserRequest) GetCredentialsOk() (*UserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UpdateUserRequest) SetCredentials(v UserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UpdateUserRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateUserRequest) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateUserRequest) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateUserRequest) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UpdateUserRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRealmId

`func (o *UpdateUserRequest) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *UpdateUserRequest) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *UpdateUserRequest) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *UpdateUserRequest) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


