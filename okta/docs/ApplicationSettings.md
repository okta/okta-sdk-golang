# ApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmOptInStatus** | Pointer to **string** | The Governance Engine opt-in status for the app | [optional] [readonly] 
**IdentityStoreId** | Pointer to **string** | Identifies an additional identity store app, if your app supports it. The &#x60;identityStoreId&#x60; value must be a valid identity store app ID. This identity store app must be created in the same org as your app. | [optional] 
**ImplicitAssignment** | Pointer to **bool** | Controls whether Okta automatically assigns users to the app based on the user&#39;s role or group membership. | [optional] 
**InlineHookId** | Pointer to **string** | Identifier of an inline hook. Inline hooks are outbound calls from Okta to your own custom code, triggered at specific points in Okta process flows. They allow you to integrate custom functionality into those flows. See [Inline hooks](/openapi/okta-management/management/tag/InlineHook/). | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 

## Methods

### NewApplicationSettings

`func NewApplicationSettings() *ApplicationSettings`

NewApplicationSettings instantiates a new ApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSettingsWithDefaults

`func NewApplicationSettingsWithDefaults() *ApplicationSettings`

NewApplicationSettingsWithDefaults instantiates a new ApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmOptInStatus

`func (o *ApplicationSettings) GetEmOptInStatus() string`

GetEmOptInStatus returns the EmOptInStatus field if non-nil, zero value otherwise.

### GetEmOptInStatusOk

`func (o *ApplicationSettings) GetEmOptInStatusOk() (*string, bool)`

GetEmOptInStatusOk returns a tuple with the EmOptInStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmOptInStatus

`func (o *ApplicationSettings) SetEmOptInStatus(v string)`

SetEmOptInStatus sets EmOptInStatus field to given value.

### HasEmOptInStatus

`func (o *ApplicationSettings) HasEmOptInStatus() bool`

HasEmOptInStatus returns a boolean if a field has been set.

### GetIdentityStoreId

`func (o *ApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *ApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *ApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *ApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *ApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *ApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *ApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *ApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *ApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *ApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *ApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *ApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *ApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *ApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


