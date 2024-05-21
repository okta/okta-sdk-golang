# OINBaseSignOnModeApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStoreId** | Pointer to **string** |  | [optional] 
**ImplicitAssignment** | Pointer to **bool** |  | [optional] 
**InlineHookId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 
**App** | Pointer to **map[string]string** |  | [optional] 
**SignOn** | Pointer to [**OINApplicationSettingsSignOn**](OINApplicationSettingsSignOn.md) |  | [optional] 

## Methods

### NewOINBaseSignOnModeApplicationSettings

`func NewOINBaseSignOnModeApplicationSettings() *OINBaseSignOnModeApplicationSettings`

NewOINBaseSignOnModeApplicationSettings instantiates a new OINBaseSignOnModeApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOINBaseSignOnModeApplicationSettingsWithDefaults

`func NewOINBaseSignOnModeApplicationSettingsWithDefaults() *OINBaseSignOnModeApplicationSettings`

NewOINBaseSignOnModeApplicationSettingsWithDefaults instantiates a new OINBaseSignOnModeApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStoreId

`func (o *OINBaseSignOnModeApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *OINBaseSignOnModeApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *OINBaseSignOnModeApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *OINBaseSignOnModeApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *OINBaseSignOnModeApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *OINBaseSignOnModeApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *OINBaseSignOnModeApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *OINBaseSignOnModeApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *OINBaseSignOnModeApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *OINBaseSignOnModeApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *OINBaseSignOnModeApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *OINBaseSignOnModeApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *OINBaseSignOnModeApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *OINBaseSignOnModeApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *OINBaseSignOnModeApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *OINBaseSignOnModeApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *OINBaseSignOnModeApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *OINBaseSignOnModeApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *OINBaseSignOnModeApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *OINBaseSignOnModeApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetApp

`func (o *OINBaseSignOnModeApplicationSettings) GetApp() map[string]string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OINBaseSignOnModeApplicationSettings) GetAppOk() (*map[string]string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OINBaseSignOnModeApplicationSettings) SetApp(v map[string]string)`

SetApp sets App field to given value.

### HasApp

`func (o *OINBaseSignOnModeApplicationSettings) HasApp() bool`

HasApp returns a boolean if a field has been set.

### SetAppNil

`func (o *OINBaseSignOnModeApplicationSettings) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *OINBaseSignOnModeApplicationSettings) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetSignOn

`func (o *OINBaseSignOnModeApplicationSettings) GetSignOn() OINApplicationSettingsSignOn`

GetSignOn returns the SignOn field if non-nil, zero value otherwise.

### GetSignOnOk

`func (o *OINBaseSignOnModeApplicationSettings) GetSignOnOk() (*OINApplicationSettingsSignOn, bool)`

GetSignOnOk returns a tuple with the SignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOn

`func (o *OINBaseSignOnModeApplicationSettings) SetSignOn(v OINApplicationSettingsSignOn)`

SetSignOn sets SignOn field to given value.

### HasSignOn

`func (o *OINBaseSignOnModeApplicationSettings) HasSignOn() bool`

HasSignOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


