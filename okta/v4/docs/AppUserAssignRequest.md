# AppUserAssignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Credentials** | Pointer to [**AppUserCredentials**](AppUserCredentials.md) |  | [optional] 
**ExternalId** | Pointer to **string** | The ID of the user in the target app that&#39;s linked to the Okta Application User object. This value is the native app-specific identifier or primary key for the user in the target app.  The &#x60;externalId&#x60; is set during import when the user is confirmed (reconciled) or during provisioning when the user is created in the target app. This value isn&#39;t populated for SSO app assignments (for example, SAML or SWA) because it isn&#39;t synchronized with a target app. | [optional] [readonly] 
**Id** | **string** | Unique identifier for the Okta User | 
**LastSync** | Pointer to **time.Time** | Timestamp of the last synchronization operation. This value is only updated for apps with the &#x60;IMPORT_PROFILE_UPDATES&#x60; or &#x60;PUSH PROFILE_UPDATES&#x60; feature. | [optional] [readonly] 
**LastUpdated** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**PasswordChanged** | Pointer to **NullableTime** | Timestamp when the Application User password was last changed | [optional] [readonly] 
**Profile** | Pointer to **map[string]interface{}** | Specifies the default and custom profile properties for a user. Properties that are visible in the Admin Console for an app assignment can also be assigned through the API. Some properties are reference properties that are imported from the target app and can&#39;t be configured. See [profile](/openapi/okta-management/management/tag/User/#tag/User/operation/getUser!c&#x3D;200&amp;path&#x3D;profile&amp;t&#x3D;response).  | [optional] 
**Scope** | Pointer to **string** | Indicates if the assignment is direct (&#x60;USER&#x60;) or by group membership (&#x60;GROUP&#x60;). | [optional] 
**Status** | Pointer to **string** | Status of an Application User | [optional] [readonly] 
**StatusChanged** | Pointer to **time.Time** | Timestamp when the Application User status was last changed | [optional] [readonly] 
**SyncState** | Pointer to **string** | The synchronization state for the Application User. The Application User&#39;s &#x60;syncState&#x60; depends on whether the &#x60;PROFILE_MASTERING&#x60; feature is enabled for the app.  &gt; **Note:** User provisioning currently must be configured through the Admin Console. | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | Embedded resources related to the Application User using the [JSON Hypertext Application Language](https://datatracker.ietf.org/doc/html/draft-kelly-json-hal-06) specification | [optional] [readonly] 
**Links** | Pointer to [**LinksAppAndUser**](LinksAppAndUser.md) |  | [optional] 

## Methods

### NewAppUserAssignRequest

`func NewAppUserAssignRequest(id string, ) *AppUserAssignRequest`

NewAppUserAssignRequest instantiates a new AppUserAssignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppUserAssignRequestWithDefaults

`func NewAppUserAssignRequestWithDefaults() *AppUserAssignRequest`

NewAppUserAssignRequestWithDefaults instantiates a new AppUserAssignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *AppUserAssignRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AppUserAssignRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AppUserAssignRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AppUserAssignRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *AppUserAssignRequest) GetCredentials() AppUserCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AppUserAssignRequest) GetCredentialsOk() (*AppUserCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AppUserAssignRequest) SetCredentials(v AppUserCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AppUserAssignRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetExternalId

`func (o *AppUserAssignRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AppUserAssignRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AppUserAssignRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AppUserAssignRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *AppUserAssignRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppUserAssignRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppUserAssignRequest) SetId(v string)`

SetId sets Id field to given value.


### GetLastSync

`func (o *AppUserAssignRequest) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AppUserAssignRequest) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AppUserAssignRequest) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AppUserAssignRequest) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AppUserAssignRequest) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AppUserAssignRequest) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AppUserAssignRequest) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AppUserAssignRequest) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPasswordChanged

`func (o *AppUserAssignRequest) GetPasswordChanged() time.Time`

GetPasswordChanged returns the PasswordChanged field if non-nil, zero value otherwise.

### GetPasswordChangedOk

`func (o *AppUserAssignRequest) GetPasswordChangedOk() (*time.Time, bool)`

GetPasswordChangedOk returns a tuple with the PasswordChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChanged

`func (o *AppUserAssignRequest) SetPasswordChanged(v time.Time)`

SetPasswordChanged sets PasswordChanged field to given value.

### HasPasswordChanged

`func (o *AppUserAssignRequest) HasPasswordChanged() bool`

HasPasswordChanged returns a boolean if a field has been set.

### SetPasswordChangedNil

`func (o *AppUserAssignRequest) SetPasswordChangedNil(b bool)`

 SetPasswordChangedNil sets the value for PasswordChanged to be an explicit nil

### UnsetPasswordChanged
`func (o *AppUserAssignRequest) UnsetPasswordChanged()`

UnsetPasswordChanged ensures that no value is present for PasswordChanged, not even an explicit nil
### GetProfile

`func (o *AppUserAssignRequest) GetProfile() map[string]interface{}`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AppUserAssignRequest) GetProfileOk() (*map[string]interface{}, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AppUserAssignRequest) SetProfile(v map[string]interface{})`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *AppUserAssignRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetScope

`func (o *AppUserAssignRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AppUserAssignRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AppUserAssignRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AppUserAssignRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetStatus

`func (o *AppUserAssignRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppUserAssignRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppUserAssignRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppUserAssignRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChanged

`func (o *AppUserAssignRequest) GetStatusChanged() time.Time`

GetStatusChanged returns the StatusChanged field if non-nil, zero value otherwise.

### GetStatusChangedOk

`func (o *AppUserAssignRequest) GetStatusChangedOk() (*time.Time, bool)`

GetStatusChangedOk returns a tuple with the StatusChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChanged

`func (o *AppUserAssignRequest) SetStatusChanged(v time.Time)`

SetStatusChanged sets StatusChanged field to given value.

### HasStatusChanged

`func (o *AppUserAssignRequest) HasStatusChanged() bool`

HasStatusChanged returns a boolean if a field has been set.

### GetSyncState

`func (o *AppUserAssignRequest) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *AppUserAssignRequest) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *AppUserAssignRequest) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *AppUserAssignRequest) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.

### GetEmbedded

`func (o *AppUserAssignRequest) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *AppUserAssignRequest) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *AppUserAssignRequest) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *AppUserAssignRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *AppUserAssignRequest) GetLinks() LinksAppAndUser`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppUserAssignRequest) GetLinksOk() (*LinksAppAndUser, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppUserAssignRequest) SetLinks(v LinksAppAndUser)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppUserAssignRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


