# SlackApplicationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityStoreId** | Pointer to **string** |  | [optional] 
**ImplicitAssignment** | Pointer to **bool** |  | [optional] 
**InlineHookId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to [**ApplicationSettingsNotes**](ApplicationSettingsNotes.md) |  | [optional] 
**Notifications** | Pointer to [**ApplicationSettingsNotifications**](ApplicationSettingsNotifications.md) |  | [optional] 
**App** | [**SlackApplicationSettingsApplication**](SlackApplicationSettingsApplication.md) |  | 
**SignOn** | Pointer to [**OINApplicationSettingsSignOn**](OINApplicationSettingsSignOn.md) |  | [optional] 

## Methods

### NewSlackApplicationSettings

`func NewSlackApplicationSettings(app SlackApplicationSettingsApplication, ) *SlackApplicationSettings`

NewSlackApplicationSettings instantiates a new SlackApplicationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlackApplicationSettingsWithDefaults

`func NewSlackApplicationSettingsWithDefaults() *SlackApplicationSettings`

NewSlackApplicationSettingsWithDefaults instantiates a new SlackApplicationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityStoreId

`func (o *SlackApplicationSettings) GetIdentityStoreId() string`

GetIdentityStoreId returns the IdentityStoreId field if non-nil, zero value otherwise.

### GetIdentityStoreIdOk

`func (o *SlackApplicationSettings) GetIdentityStoreIdOk() (*string, bool)`

GetIdentityStoreIdOk returns a tuple with the IdentityStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStoreId

`func (o *SlackApplicationSettings) SetIdentityStoreId(v string)`

SetIdentityStoreId sets IdentityStoreId field to given value.

### HasIdentityStoreId

`func (o *SlackApplicationSettings) HasIdentityStoreId() bool`

HasIdentityStoreId returns a boolean if a field has been set.

### GetImplicitAssignment

`func (o *SlackApplicationSettings) GetImplicitAssignment() bool`

GetImplicitAssignment returns the ImplicitAssignment field if non-nil, zero value otherwise.

### GetImplicitAssignmentOk

`func (o *SlackApplicationSettings) GetImplicitAssignmentOk() (*bool, bool)`

GetImplicitAssignmentOk returns a tuple with the ImplicitAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitAssignment

`func (o *SlackApplicationSettings) SetImplicitAssignment(v bool)`

SetImplicitAssignment sets ImplicitAssignment field to given value.

### HasImplicitAssignment

`func (o *SlackApplicationSettings) HasImplicitAssignment() bool`

HasImplicitAssignment returns a boolean if a field has been set.

### GetInlineHookId

`func (o *SlackApplicationSettings) GetInlineHookId() string`

GetInlineHookId returns the InlineHookId field if non-nil, zero value otherwise.

### GetInlineHookIdOk

`func (o *SlackApplicationSettings) GetInlineHookIdOk() (*string, bool)`

GetInlineHookIdOk returns a tuple with the InlineHookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineHookId

`func (o *SlackApplicationSettings) SetInlineHookId(v string)`

SetInlineHookId sets InlineHookId field to given value.

### HasInlineHookId

`func (o *SlackApplicationSettings) HasInlineHookId() bool`

HasInlineHookId returns a boolean if a field has been set.

### GetNotes

`func (o *SlackApplicationSettings) GetNotes() ApplicationSettingsNotes`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SlackApplicationSettings) GetNotesOk() (*ApplicationSettingsNotes, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SlackApplicationSettings) SetNotes(v ApplicationSettingsNotes)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SlackApplicationSettings) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetNotifications

`func (o *SlackApplicationSettings) GetNotifications() ApplicationSettingsNotifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *SlackApplicationSettings) GetNotificationsOk() (*ApplicationSettingsNotifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *SlackApplicationSettings) SetNotifications(v ApplicationSettingsNotifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *SlackApplicationSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetApp

`func (o *SlackApplicationSettings) GetApp() SlackApplicationSettingsApplication`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SlackApplicationSettings) GetAppOk() (*SlackApplicationSettingsApplication, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SlackApplicationSettings) SetApp(v SlackApplicationSettingsApplication)`

SetApp sets App field to given value.


### GetSignOn

`func (o *SlackApplicationSettings) GetSignOn() OINApplicationSettingsSignOn`

GetSignOn returns the SignOn field if non-nil, zero value otherwise.

### GetSignOnOk

`func (o *SlackApplicationSettings) GetSignOnOk() (*OINApplicationSettingsSignOn, bool)`

GetSignOnOk returns a tuple with the SignOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOn

`func (o *SlackApplicationSettings) SetSignOn(v OINApplicationSettingsSignOn)`

SetSignOn sets SignOn field to given value.

### HasSignOn

`func (o *SlackApplicationSettings) HasSignOn() bool`

HasSignOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


