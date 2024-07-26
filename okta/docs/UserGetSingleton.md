# UserGetSingleton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **NullableTime** | The timestamp when the user status transitioned to &#x60;ACTIVE&#x60; | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The timestamp when the user was created | [optional] [readonly] 
**Credentials** | Pointer to [**UserCredentials**](UserCredentials.md) |  | [optional] 
**Id** | Pointer to **string** | The unique key for the user | [optional] [readonly] 
**LastLogin** | Pointer to **NullableTime** | The timestamp of the last login | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The timestamp when the user was last updated | [optional] [readonly] 
**PasswordChanged** | Pointer to **NullableTime** | The timestamp when the user&#39;s password was last updated | [optional] [readonly] 
**Profile** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 
**RealmId** | Pointer to **string** | &lt;div class&#x3D;\&quot;x-lifecycle-container\&quot;&gt;&lt;x-lifecycle class&#x3D;\&quot;ea\&quot;&gt;&lt;/x-lifecycle&gt;&lt;/div&gt;The ID of the Realm in which the user is residing | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the user | [optional] [readonly] 
**StatusChanged** | Pointer to **NullableTime** | The timestamp when the status of the user last changed | [optional] [readonly] 
**TransitioningToStatus** | Pointer to **NullableString** | The target status of an in-progress asynchronous status transition. This property is only returned if the user&#39;s state is transitioning. | [optional] [readonly] 
**Type** | Pointer to [**UserType**](UserType.md) |  | [optional] 
**Embedded** | Pointer to [**UserGetSingletonAllOfEmbedded**](UserGetSingletonAllOfEmbedded.md) |  | [optional] 
**Links** | Pointer to [**UserLinks**](UserLinks.md) |  | [optional] 

## Methods

### NewUserGetSingleton

`func NewUserGetSingleton() *UserGetSingleton`

NewUserGetSingleton instantiates a new UserGetSingleton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetSingletonWithDefaults

`func NewUserGetSingletonWithDefaults() *UserGetSingleton`

NewUserGetSingletonWithDefaults instantiates a new UserGetSingleton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *UserGetSingleton) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *UserGetSingleton) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *UserGetSingleton) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *UserGetSingleton) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### SetActivatedNil

`func (o *UserGetSingleton) SetActivatedNil(b bool)`

 SetActivatedNil sets the value for Activated to be an explicit nil

### UnsetActivated
`func (o *UserGetSingleton) UnsetActivated()`

UnsetActivated ensures that no value is present for Activated, not even an explicit nil
### GetCreated

`func (o *UserGetSingleton) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserGetSingleton) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserGetSingleton) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserGetSingleton) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *UserGetSingleton) GetCredentials() UserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UserGetSingleton) GetCredentialsOk() (*UserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UserGetSingleton) SetCredentials(v UserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UserGetSingleton) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetId

`func (o *UserGetSingleton) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserGetSingleton) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserGetSingleton) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserGetSingleton) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserGetSingleton) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserGetSingleton) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserGetSingleton) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserGetSingleton) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *UserGetSingleton) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *UserGetSingleton) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetLastUpdated

`func (o *UserGetSingleton) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserGetSingleton) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserGetSingleton) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserGetSingleton) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPasswordChanged

`func (o *UserGetSingleton) GetPasswordChanged() time.Time`

GetPasswordChanged returns the PasswordChanged field if non-nil, zero value otherwise.

### GetPasswordChangedOk

`func (o *UserGetSingleton) GetPasswordChangedOk() (*time.Time, bool)`

GetPasswordChangedOk returns a tuple with the PasswordChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChanged

`func (o *UserGetSingleton) SetPasswordChanged(v time.Time)`

SetPasswordChanged sets PasswordChanged field to given value.

### HasPasswordChanged

`func (o *UserGetSingleton) HasPasswordChanged() bool`

HasPasswordChanged returns a boolean if a field has been set.

### SetPasswordChangedNil

`func (o *UserGetSingleton) SetPasswordChangedNil(b bool)`

 SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil

### UnsetPasswordChanged
`func (o *UserGetSingleton) UnsetPasswordChanged()`

UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
### GetProfile

`func (o *UserGetSingleton) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserGetSingleton) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserGetSingleton) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserGetSingleton) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetRealmId

`func (o *UserGetSingleton) GetRealmId() string`

GetRealmId returns the RealmId field if non-nil, zero value otherwise.

### GetRealmIdOk

`func (o *UserGetSingleton) GetRealmIdOk() (*string, bool)`

GetRealmIdOk returns a tuple with the RealmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealmId

`func (o *UserGetSingleton) SetRealmId(v string)`

SetRealmId sets RealmId field to given value.

### HasRealmId

`func (o *UserGetSingleton) HasRealmId() bool`

HasRealmId returns a boolean if a field has been set.

### GetStatus

`func (o *UserGetSingleton) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserGetSingleton) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserGetSingleton) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserGetSingleton) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChanged

`func (o *UserGetSingleton) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *UserGetSingleton) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *UserGetSingleton) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.

### HasStatusChanged

`func (o *UserGetSingleton) HasStatusChanged() bool`

HasStatusChanged returns a boolean if a field has been set.

### SetStatusChangedNil

`func (o *UserGetSingleton) SetStatusChangedNil(b bool)`

 SetStatusChangedNil sets the value for StatusChanged to be an explicit nil

### UnsetStatusChanged
`func (o *UserGetSingleton) UnsetStatusChanged()`

UnsetStatusChanged ensures that no value is present for StatusChanged, not even an explicit nil
### GetTransitioningToStatus

`func (o *UserGetSingleton) GetTransitioningToStatus() string`

GetTransitioningToStatus returns the TransitioningToStatus field if non-nil, zero value otherwise.

### GetTransitioningToStatusOk

`func (o *UserGetSingleton) GetTransitioningToStatusOk() (*string, bool)`

GetTransitioningToStatusOk returns a tuple with the TransitioningToStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitioningToStatus

`func (o *UserGetSingleton) SetTransitioningToStatus(v string)`

SetTransitioningToStatus sets TransitioningToStatus field to given value.

### HasTransitioningToStatus

`func (o *UserGetSingleton) HasTransitioningToStatus() bool`

HasTransitioningToStatus returns a boolean if a field has been set.

### SetTransitioningToStatusNil

`func (o *UserGetSingleton) SetTransitioningToStatusNil(b bool)`

 SetTransitioningToStatusNil sets the value for TransitioningToStatus to be an explicit nil

### UnsetTransitioningToStatus
`func (o *UserGetSingleton) UnsetTransitioningToStatus()`

UnsetTransitioningToStatus ensures that no value is present for TransitioningToStatus, not even an explicit nil
### GetType

`func (o *UserGetSingleton) GetType() UserType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserGetSingleton) GetTypeOk() (*UserType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserGetSingleton) SetType(v UserType)`

SetType sets Type field to given value.

### HasType

`func (o *UserGetSingleton) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEmbedded

`func (o *UserGetSingleton) GetEmbedded() UserGetSingletonAllOfEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UserGetSingleton) GetEmbeddedOk() (*UserGetSingletonAllOfEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UserGetSingleton) SetEmbedded(v UserGetSingletonAllOfEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UserGetSingleton) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *UserGetSingleton) GetLinks() UserLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserGetSingleton) GetLinksOk() (*UserLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserGetSingleton) SetLinks(v UserLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserGetSingleton) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


