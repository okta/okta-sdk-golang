# AppUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) |  | 
**Credentials** | Pointer to [**AppUserCredentials**](AppUserCredentials.md) |  | [optional] 
**ExternalId** | Pointer to **string** | The ID of the user in the target app that&#39;s linked to the Okta App User object. This value is the native app-specific identifier or primary key for the user in the target app.  The &#x60;externalId&#x60; is set during import when the user is confirmed (reconciled) or during provisioning when the user has been successfully created in the target app. This value isn&#39;t populated for SSO app assignments (for example, SAML or SWA) because it isn&#39;t synchronized with a target app. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique identifier of the App User object (only required for apps with &#x60;signOnMode&#x60; or authentication schemes that don&#39;t require credentials) | [optional] 
**LastSync** | Pointer to **time.Time** | Timestamp of the last synchronization operation. This value is only updated for apps with the &#x60;IMPORT_PROFILE_UPDATES&#x60; or &#x60;PUSH PROFILE_UPDATES&#x60; feature. | [optional] [readonly] 
**LastUpdated** | [**time.Time**](time.Time.md) |  | 
**PasswordChanged** | Pointer to **NullableTime** | Timestamp when the App User password was last changed | [optional] [readonly] 
**Profile** | Pointer to **map[string]map[string]interface{}** | App user profiles are app-specific and can be customized by the Profile Editor in the Admin Console. SSO apps typically don&#39;t support app user profiles, while apps with user provisioning features have app-specific profiles. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can&#39;t be configured. | [optional] 
**Scope** | **string** | Toggles the assignment between user or group scope | 
**Status** | **string** | Status of an App User | [readonly] 
**StatusChanged** | **time.Time** | Timestamp when the App User status was last changed | [readonly] 
**SyncState** | Pointer to **string** | The synchronization state for the App User. The App User&#39;s &#x60;syncState&#x60; depends on whether the &#x60;PROFILE_MASTERING&#x60; feature is enabled for the app.  &gt; **Note:** User provisioning currently must be configured through the Admin Console. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | Embedded resources related to the App User using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification | [optional] [readonly] 
**Links** | [**LinksAppAndUser**](LinksAppAndUser.md) |  | 

## Methods

### NewAppUser

`func NewAppUser(created time.Time, lastUpdated time.Time, scope string, status string, statusChanged time.Time, links LinksAppAndUser, ) *AppUser`

NewAppUser instantiates a new AppUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserWithDefaults

`func NewAppUserWithDefaults() *AppUser`

NewAppUserWithDefaults instantiates a new AppUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AppUser) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppUser) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppUser) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCredentials

`func (o *AppUser) GetCredentials() AppUserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AppUser) GetCredentialsOk() (*AppUserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AppUser) SetCredentials(v AppUserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AppUser) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetExternalId

`func (o *AppUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AppUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AppUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AppUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *AppUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AppUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastSync

`func (o *AppUser) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AppUser) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AppUser) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AppUser) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AppUser) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AppUser) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AppUser) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetPasswordChanged

`func (o *AppUser) GetPasswordChanged() time.Time`

GetPasswordChanged returns the PasswordChanged field if non-nil, zero value otherwise.

### GetPasswordChangedOk

`func (o *AppUser) GetPasswordChangedOk() (*time.Time, bool)`

GetPasswordChangedOk returns a tuple with the PasswordChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChanged

`func (o *AppUser) SetPasswordChanged(v time.Time)`

SetPasswordChanged sets PasswordChanged field to given value.

### HasPasswordChanged

`func (o *AppUser) HasPasswordChanged() bool`

HasPasswordChanged returns a boolean if a field has been set.

### SetPasswordChangedNil

`func (o *AppUser) SetPasswordChangedNil(b bool)`

 SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil

### UnsetPasswordChanged
`func (o *AppUser) UnsetPasswordChanged()`

UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
### GetProfile

`func (o *AppUser) GetProfile() map[string]map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AppUser) GetProfileOk() (*map[string]map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AppUser) SetProfile(v map[string]map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AppUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetScope

`func (o *AppUser) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppUser) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppUser) SetScope(v string)`

SetScope sets Scope field to given value.


### GetStatus

`func (o *AppUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusChanged

`func (o *AppUser) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *AppUser) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *AppUser) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.


### GetSyncState

`func (o *AppUser) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *AppUser) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *AppUser) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *AppUser) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.

### GetEmbedded

`func (o *AppUser) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AppUser) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AppUser) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AppUser) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *AppUser) GetLinks() LinksAppAndUser`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppUser) GetLinksOk() (*LinksAppAndUser, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppUser) SetLinks(v LinksAppAndUser)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


