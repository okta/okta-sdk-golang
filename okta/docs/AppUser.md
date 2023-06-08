# AppUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] [readonly] 
**Credentials** | Pointer to [**AppUserCredentials**](AppUserCredentials.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] [readonly] 
**PasswordChanged** | Pointer to **time.Time** |  | [optional] [readonly] 
**Profile** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [readonly] 
**StatusChanged** | Pointer to **time.Time** |  | [optional] [readonly] 
**SyncState** | Pointer to **string** |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 

## Methods

### NewAppUser

`func NewAppUser() *AppUser`

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

### HasCreated

`func (o *AppUser) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

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

### HasLastUpdated

`func (o *AppUser) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

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

### HasScope

`func (o *AppUser) HasScope() bool`

HasScope returns a boolean if a field has been set.

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

### HasStatus

`func (o *AppUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

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

### HasStatusChanged

`func (o *AppUser) HasStatusChanged() bool`

HasStatusChanged returns a boolean if a field has been set.

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

`func (o *AppUser) GetLinks() map[string]map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppUser) GetLinksOk() (*map[string]map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppUser) SetLinks(v map[string]map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppUser) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


