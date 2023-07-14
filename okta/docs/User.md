# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Credentials** | Pointer to [**UserCredentials**](UserCredentials.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**PasswordChanged** | Pointer to **NullableTime** |  | [optional] [readonly] 
**Profile** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 
**Status** | Pointer to [**UserStatus**](UserStatus.md) |  | [optional] 
**StatusChanged** | Pointer to **NullableTime** |  | [optional] [readonly] 
**TransitioningToStatus** | Pointer to [**UserStatus**](UserStatus.md) |  | [optional] 
**Type** | Pointer to [**UserType**](UserType.md) |  | [optional] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *User) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *User) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *User) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *User) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### SetActivatedNil

`func (o *User) SetActivatedNil(b bool)`

 SetActivatedNil sets the value for Activated to be an explicit nil

### UnsetActivated
`func (o *User) UnsetActivated()`

UnsetActivated ensures that no value is present for Activated, not even an explicit nil
### GetCreated

`func (o *User) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *User) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *User) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *User) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *User) GetCredentials() UserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *User) GetCredentialsOk() (*UserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *User) SetCredentials(v UserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *User) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *User) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *User) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLastUpdated

`func (o *User) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *User) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *User) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *User) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPasswordChanged

`func (o *User) GetPasswordChanged() time.Time`

GetPasswordChanged returns the PasswordChanged field if non-nil, zero value otherwise.

### GetPasswordChangedOk

`func (o *User) GetPasswordChangedOk() (*time.Time, bool)`

GetPasswordChangedOk returns a tuple with the PasswordChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChanged

`func (o *User) SetPasswordChanged(v time.Time)`

SetPasswordChanged sets PasswordChanged field to given value.

### HasPasswordChanged

`func (o *User) HasPasswordChanged() bool`

HasPasswordChanged returns a boolean if a field has been set.

### SetPasswordChangedNil

`func (o *User) SetPasswordChangedNil(b bool)`

 SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil

### UnsetPasswordChanged
`func (o *User) UnsetPasswordChanged()`

UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
### GetProfile

`func (o *User) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *User) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *User) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *User) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() UserStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*UserStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v UserStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChanged

`func (o *User) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *User) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *User) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.

### HasStatusChanged

`func (o *User) HasStatusChanged() bool`

HasStatusChanged returns a boolean if a field has been set.

### SetStatusChangedNil

`func (o *User) SetStatusChangedNil(b bool)`

 SetStatusChangedNil sets the value for StatusChanged to be an explicit nil

### UnsetStatusChanged
`func (o *User) UnsetStatusChanged()`

UnsetStatusChanged ensures that no value is present for StatusChanged, not even an explicit nil
### GetTransitioningToStatus

`func (o *User) GetTransitioningToStatus() UserStatus`

GetTransitioningToStatus returns the TransitioningToStatus field if non-nil, zero value otherwise.

### GetTransitioningToStatusOk

`func (o *User) GetTransitioningToStatusOk() (*UserStatus, bool)`

GetTransitioningToStatusOk returns a tuple with the TransitioningToStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitioningToStatus

`func (o *User) SetTransitioningToStatus(v UserStatus)`

SetTransitioningToStatus sets TransitioningToStatus field to given value.

### HasTransitioningToStatus

`func (o *User) HasTransitioningToStatus() bool`

HasTransitioningToStatus returns a boolean if a field has been set.

### GetType

`func (o *User) GetType() UserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*UserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v UserType)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *User) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *User) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *User) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *User) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *User) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *User) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *User) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *User) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


